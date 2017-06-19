// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"crypto/tls"
	"flag"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/johanbrandhorst/grpcweb-example/client/compiled"
	"github.com/johanbrandhorst/grpcweb-example/server"
	"github.com/johanbrandhorst/grpcweb-example/server/proto/library"
)

var logger *logrus.Logger
var host = flag.String("host", "", "host to get LetsEncrypt certificate for")

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.Kitchen,
		DisableSorting:  true,
	})
	// Should only be done from init functions
	grpclog.SetLogger(logger)
}

func main() {
	flag.Parse()

	gs := grpc.NewServer()
	library.RegisterBookServiceServer(gs, &server.BookService{})
	wrappedServer := grpcweb.WrapServer(gs)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		if wrappedServer.IsGrpcWebRequest(req) {
			wrappedServer.ServeHttp(resp, req)
		} else {
			// Serve the GopherJS client
			http.FileServer(&assetfs.AssetFS{
				Asset:     compiled.Asset,
				AssetDir:  compiled.AssetDir,
				AssetInfo: compiled.AssetInfo,
			}).ServeHTTP(resp, req)
		}
	}

	// Serve on localhost with localhost certs if no host provided
	if *host == "" {
		httpServer := http.Server{
			Handler: http.HandlerFunc(handler),
			Addr:    "localhost:10000",
		}
		logger.Info("Serving on https://localhost:10000")
		logger.Fatal(httpServer.ListenAndServeTLS("./insecure/localhost.crt", "./insecure/localhost.key"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// Create server for redirecting HTTP to HTTPS
	httpSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + req.Host + req.URL.String()
			http.Redirect(w, req, url, http.StatusMovedPermanently)
		}),
	}
	go func() {
		logger.Fatal(httpSrv.ListenAndServe())
	}()

	// Create auto-certificate https server
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(*host),
	}
	httpsSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":https",
		TLSConfig: &tls.Config{
			GetCertificate:           m.GetCertificate,
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
		Handler: hstsHandler(mux.ServeHTTP),
	}
	logger.Info("Serving on https://0.0.0.0:443, authenticating for https://", *host)
	logger.Fatal(httpsSrv.ListenAndServeTLS("", ""))
}

// hstsHandler wraps an http.HandlerFunc such that it sets the HSTS header.
func hstsHandler(fn http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		fn(w, r)
	})
}