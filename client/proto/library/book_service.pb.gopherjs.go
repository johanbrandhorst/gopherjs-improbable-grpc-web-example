// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: proto/library/book_service.proto

/*
	Package library is a generated protocol buffer package.

	Package library exposes a list of books
	over a gRPC API.

	It is generated from these files:
		proto/library/book_service.proto

	It has these top-level messages:
		Publisher
		Book
		GetBookRequest
		QueryBooksRequest
		Collection
		BookMessage
		BookResponse
*/
package library

import jspb "github.com/johanbrandhorst/protobuf/jspb"
import google_protobuf "github.com/johanbrandhorst/protobuf/ptypes/timestamp"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// BookType describes the different types
// a book in the library can be.
type BookType int

const (
	// Hardcover is a book with a hard back.
	BookType_HARDCOVER BookType = 0
	// Paperback is a book with a soft back.
	BookType_PAPERBACK BookType = 1
	// Audiobook is an audio recording of the book.
	BookType_AUDIOBOOK BookType = 2
)

var BookType_name = map[int]string{
	0: "HARDCOVER",
	1: "PAPERBACK",
	2: "AUDIOBOOK",
}
var BookType_value = map[string]int{
	"HARDCOVER": 0,
	"PAPERBACK": 1,
	"AUDIOBOOK": 2,
}

func (x BookType) String() string {
	return BookType_name[int(x)]
}

// Publisher describes a Book Publisher.
type Publisher struct {
	// Name is the name of the Publisher.
	Name string
}

// GetName gets the Name of the Publisher.
func (m *Publisher) GetName() (x string) {
	if m == nil {
		return x
	}
	return m.Name
}

// MarshalToWriter marshals Publisher to the provided writer.
func (m *Publisher) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Name) > 0 {
		writer.WriteString(1, m.Name)
	}

	return
}

// Marshal marshals Publisher to a slice of bytes.
func (m *Publisher) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Publisher from the provided reader.
func (m *Publisher) UnmarshalFromReader(reader jspb.Reader) *Publisher {
	for reader.Next() {
		if m == nil {
			m = &Publisher{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Name = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Publisher from a slice of bytes.
func (m *Publisher) Unmarshal(rawBytes []byte) (*Publisher, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Book represents a book in the library.
type Book struct {
	// Isbn is the ISBN number of the book.
	Isbn int64
	// Title is the title of the book.
	Title string
	// Author is the author of the book.
	Author string
	// BookType is the type of the book.
	BookType BookType
	// PublishingMethod is the publishing method
	// used for this Book.
	//
	// Types that are valid to be assigned to PublishingMethod:
	//	*Book_SelfPublished
	//	*Book_Publisher
	PublishingMethod isBook_PublishingMethod
	// PublicationDate is the time of publication of the book.
	PublicationDate *google_protobuf.Timestamp
}

// isBook_PublishingMethod is used to distinguish types assignable to PublishingMethod
type isBook_PublishingMethod interface{ isBook_PublishingMethod() }

// Book_SelfPublished is assignable to PublishingMethod
type Book_SelfPublished struct {
	// SelfPublished means this book was
	// self published.
	SelfPublished bool
}

// Book_Publisher is assignable to PublishingMethod
type Book_Publisher struct {
	// Publisher means this book was published
	// through a Publisher.
	Publisher *Publisher
}

func (*Book_SelfPublished) isBook_PublishingMethod() {}
func (*Book_Publisher) isBook_PublishingMethod()     {}

// GetPublishingMethod gets the PublishingMethod of the Book.
func (m *Book) GetPublishingMethod() (x isBook_PublishingMethod) {
	if m == nil {
		return x
	}
	return m.PublishingMethod
}

// GetIsbn gets the Isbn of the Book.
func (m *Book) GetIsbn() (x int64) {
	if m == nil {
		return x
	}
	return m.Isbn
}

// GetTitle gets the Title of the Book.
func (m *Book) GetTitle() (x string) {
	if m == nil {
		return x
	}
	return m.Title
}

// GetAuthor gets the Author of the Book.
func (m *Book) GetAuthor() (x string) {
	if m == nil {
		return x
	}
	return m.Author
}

// GetBookType gets the BookType of the Book.
func (m *Book) GetBookType() (x BookType) {
	if m == nil {
		return x
	}
	return m.BookType
}

// GetSelfPublished gets the SelfPublished of the Book.
func (m *Book) GetSelfPublished() (x bool) {
	if v, ok := m.GetPublishingMethod().(*Book_SelfPublished); ok {
		return v.SelfPublished
	}
	return x
}

// GetPublisher gets the Publisher of the Book.
func (m *Book) GetPublisher() (x *Publisher) {
	if v, ok := m.GetPublishingMethod().(*Book_Publisher); ok {
		return v.Publisher
	}
	return x
}

// GetPublicationDate gets the PublicationDate of the Book.
func (m *Book) GetPublicationDate() (x *google_protobuf.Timestamp) {
	if m == nil {
		return x
	}
	return m.PublicationDate
}

// MarshalToWriter marshals Book to the provided writer.
func (m *Book) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	switch t := m.PublishingMethod.(type) {
	case *Book_SelfPublished:
		if t.SelfPublished {
			writer.WriteBool(5, t.SelfPublished)
		}
	case *Book_Publisher:
		if t.Publisher != nil {
			writer.WriteMessage(6, func() {
				t.Publisher.MarshalToWriter(writer)
			})
		}
	}

	if m.Isbn != 0 {
		writer.WriteInt64(1, m.Isbn)
	}

	if len(m.Title) > 0 {
		writer.WriteString(2, m.Title)
	}

	if len(m.Author) > 0 {
		writer.WriteString(3, m.Author)
	}

	if int(m.BookType) != 0 {
		writer.WriteEnum(4, int(m.BookType))
	}

	if m.PublicationDate != nil {
		writer.WriteMessage(7, func() {
			m.PublicationDate.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals Book to a slice of bytes.
func (m *Book) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Book from the provided reader.
func (m *Book) UnmarshalFromReader(reader jspb.Reader) *Book {
	for reader.Next() {
		if m == nil {
			m = &Book{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Isbn = reader.ReadInt64()
		case 2:
			m.Title = reader.ReadString()
		case 3:
			m.Author = reader.ReadString()
		case 4:
			m.BookType = BookType(reader.ReadEnum())
		case 5:
			m.PublishingMethod = &Book_SelfPublished{
				SelfPublished: reader.ReadBool(),
			}
		case 6:
			reader.ReadMessage(func() {
				m.PublishingMethod = &Book_Publisher{
					Publisher: new(Publisher).UnmarshalFromReader(reader),
				}
			})
		case 7:
			reader.ReadMessage(func() {
				m.PublicationDate = m.PublicationDate.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Book from a slice of bytes.
func (m *Book) Unmarshal(rawBytes []byte) (*Book, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// GetBookRequest is the input to the GetBook method.
type GetBookRequest struct {
	// Isbn is the ISBN with which
	// to match against the ISBN of a book in the library.
	Isbn int64
}

// GetIsbn gets the Isbn of the GetBookRequest.
func (m *GetBookRequest) GetIsbn() (x int64) {
	if m == nil {
		return x
	}
	return m.Isbn
}

// MarshalToWriter marshals GetBookRequest to the provided writer.
func (m *GetBookRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Isbn != 0 {
		writer.WriteInt64(1, m.Isbn)
	}

	return
}

// Marshal marshals GetBookRequest to a slice of bytes.
func (m *GetBookRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a GetBookRequest from the provided reader.
func (m *GetBookRequest) UnmarshalFromReader(reader jspb.Reader) *GetBookRequest {
	for reader.Next() {
		if m == nil {
			m = &GetBookRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Isbn = reader.ReadInt64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a GetBookRequest from a slice of bytes.
func (m *GetBookRequest) Unmarshal(rawBytes []byte) (*GetBookRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// QueryBooksRequest is the input to the QueryBooks method.
type QueryBooksRequest struct {
	// AuthorPrefix is the prefix with which
	// to match against the author of a book in the library.
	AuthorPrefix string
}

// GetAuthorPrefix gets the AuthorPrefix of the QueryBooksRequest.
func (m *QueryBooksRequest) GetAuthorPrefix() (x string) {
	if m == nil {
		return x
	}
	return m.AuthorPrefix
}

// MarshalToWriter marshals QueryBooksRequest to the provided writer.
func (m *QueryBooksRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.AuthorPrefix) > 0 {
		writer.WriteString(1, m.AuthorPrefix)
	}

	return
}

// Marshal marshals QueryBooksRequest to a slice of bytes.
func (m *QueryBooksRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a QueryBooksRequest from the provided reader.
func (m *QueryBooksRequest) UnmarshalFromReader(reader jspb.Reader) *QueryBooksRequest {
	for reader.Next() {
		if m == nil {
			m = &QueryBooksRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.AuthorPrefix = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a QueryBooksRequest from a slice of bytes.
func (m *QueryBooksRequest) Unmarshal(rawBytes []byte) (*QueryBooksRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Collection is a collection of books
type Collection struct {
	// Books is a list of books
	Books []*Book
}

// GetBooks gets the Books of the Collection.
func (m *Collection) GetBooks() (x []*Book) {
	if m == nil {
		return x
	}
	return m.Books
}

// MarshalToWriter marshals Collection to the provided writer.
func (m *Collection) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.Books {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals Collection to a slice of bytes.
func (m *Collection) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Collection from the provided reader.
func (m *Collection) UnmarshalFromReader(reader jspb.Reader) *Collection {
	for reader.Next() {
		if m == nil {
			m = &Collection{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.Books = append(m.Books, new(Book).UnmarshalFromReader(reader))
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Collection from a slice of bytes.
func (m *Collection) Unmarshal(rawBytes []byte) (*Collection, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// BookMessage is used to discuss books
type BookMessage struct {
	// Types that are valid to be assigned to Content:
	//	*BookMessage_Name
	//	*BookMessage_Message
	Content isBookMessage_Content
}

// isBookMessage_Content is used to distinguish types assignable to Content
type isBookMessage_Content interface{ isBookMessage_Content() }

// BookMessage_Name is assignable to Content
type BookMessage_Name struct {
	// Name is the name of the person who is sending this message.
	// It should be sent as the first message on the stream.
	Name string
}

// BookMessage_Message is assignable to Content
type BookMessage_Message struct {
	// Message is any message the user wishes to send.
	Message string
}

func (*BookMessage_Name) isBookMessage_Content()    {}
func (*BookMessage_Message) isBookMessage_Content() {}

// GetContent gets the Content of the BookMessage.
func (m *BookMessage) GetContent() (x isBookMessage_Content) {
	if m == nil {
		return x
	}
	return m.Content
}

// GetName gets the Name of the BookMessage.
func (m *BookMessage) GetName() (x string) {
	if v, ok := m.GetContent().(*BookMessage_Name); ok {
		return v.Name
	}
	return x
}

// GetMessage gets the Message of the BookMessage.
func (m *BookMessage) GetMessage() (x string) {
	if v, ok := m.GetContent().(*BookMessage_Message); ok {
		return v.Message
	}
	return x
}

// MarshalToWriter marshals BookMessage to the provided writer.
func (m *BookMessage) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	switch t := m.Content.(type) {
	case *BookMessage_Name:
		if len(t.Name) > 0 {
			writer.WriteString(1, t.Name)
		}
	case *BookMessage_Message:
		if len(t.Message) > 0 {
			writer.WriteString(2, t.Message)
		}
	}

	return
}

// Marshal marshals BookMessage to a slice of bytes.
func (m *BookMessage) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a BookMessage from the provided reader.
func (m *BookMessage) UnmarshalFromReader(reader jspb.Reader) *BookMessage {
	for reader.Next() {
		if m == nil {
			m = &BookMessage{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Content = &BookMessage_Name{
				Name: reader.ReadString(),
			}
		case 2:
			m.Content = &BookMessage_Message{
				Message: reader.ReadString(),
			}
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a BookMessage from a slice of bytes.
func (m *BookMessage) Unmarshal(rawBytes []byte) (*BookMessage, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// BookResponse is used to discuss books
type BookResponse struct {
	// Message is a message from a user.
	Message string
}

// GetMessage gets the Message of the BookResponse.
func (m *BookResponse) GetMessage() (x string) {
	if m == nil {
		return x
	}
	return m.Message
}

// MarshalToWriter marshals BookResponse to the provided writer.
func (m *BookResponse) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Message) > 0 {
		writer.WriteString(2, m.Message)
	}

	return
}

// Marshal marshals BookResponse to a slice of bytes.
func (m *BookResponse) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a BookResponse from the provided reader.
func (m *BookResponse) UnmarshalFromReader(reader jspb.Reader) *BookResponse {
	for reader.Next() {
		if m == nil {
			m = &BookResponse{}
		}

		switch reader.GetFieldNumber() {
		case 2:
			m.Message = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a BookResponse from a slice of bytes.
func (m *BookResponse) Unmarshal(rawBytes []byte) (*BookResponse, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb.GrpcWebPackageIsVersion3

// Client API for BookService service

// BookService exposes GetBook and QueryBooks,
// which allow querying of the library.
type BookServiceClient interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb.CallOption) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb.CallOption) (BookService_QueryBooksClient, error)
	// MakeCollection takes a stream of books and returns a Book collection.
	MakeCollection(ctx context.Context, opts ...grpcweb.CallOption) (BookService_MakeCollectionClient, error)
	// BookChat allows discussion about books
	BookChat(ctx context.Context, opts ...grpcweb.CallOption) (BookService_BookChatClient, error)
}

type bookServiceClient struct {
	client *grpcweb.Client
}

// NewBookServiceClient creates a new gRPC-Web client.
func NewBookServiceClient(hostname string, opts ...grpcweb.DialOption) BookServiceClient {
	return &bookServiceClient{
		client: grpcweb.NewClient(hostname, "library.BookService", opts...),
	}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb.CallOption) (*Book, error) {
	resp, err := c.client.RPCCall(ctx, "GetBook", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Book).Unmarshal(resp)
}

func (c *bookServiceClient) QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb.CallOption) (BookService_QueryBooksClient, error) {
	srv, err := c.client.NewClientStream(ctx, false, true, "QueryBooks", opts...)
	if err != nil {
		return nil, err
	}

	err = srv.SendMsg(in.Marshal())
	if err != nil {
		return nil, err
	}

	return &bookServiceQueryBooksClient{srv}, nil
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
	grpcweb.ClientStream
}

type bookServiceQueryBooksClient struct {
	grpcweb.ClientStream
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	resp, err := x.RecvMsg()
	if err != nil {
		return nil, err
	}

	return new(Book).Unmarshal(resp)
}

func (c *bookServiceClient) MakeCollection(ctx context.Context, opts ...grpcweb.CallOption) (BookService_MakeCollectionClient, error) {
	srv, err := c.client.NewClientStream(ctx, true, false, "MakeCollection", opts...)
	if err != nil {
		return nil, err
	}

	return &bookServiceMakeCollectionClient{srv}, nil
}

type BookService_MakeCollectionClient interface {
	Send(*Book) error
	CloseAndRecv() (*Collection, error)
	grpcweb.ClientStream
}

type bookServiceMakeCollectionClient struct {
	grpcweb.ClientStream
}

func (x *bookServiceMakeCollectionClient) Send(req *Book) error {
	return x.SendMsg(req.Marshal())
}

func (x *bookServiceMakeCollectionClient) CloseAndRecv() (*Collection, error) {
	err := x.CloseSend()
	if err != nil {
		return nil, err
	}

	resp, err := x.RecvMsg()
	if err != nil {
		return nil, err
	}

	return new(Collection).Unmarshal(resp)
}

func (c *bookServiceClient) BookChat(ctx context.Context, opts ...grpcweb.CallOption) (BookService_BookChatClient, error) {
	srv, err := c.client.NewClientStream(ctx, true, true, "BookChat", opts...)
	if err != nil {
		return nil, err
	}

	return &bookServiceBookChatClient{srv}, nil
}

type BookService_BookChatClient interface {
	Send(*BookMessage) error
	Recv() (*BookResponse, error)
	grpcweb.ClientStream
}

type bookServiceBookChatClient struct {
	grpcweb.ClientStream
}

func (x *bookServiceBookChatClient) Send(req *BookMessage) error {
	return x.SendMsg(req.Marshal())
}

func (x *bookServiceBookChatClient) Recv() (*BookResponse, error) {
	resp, err := x.RecvMsg()
	if err != nil {
		return nil, err
	}

	return new(BookResponse).Unmarshal(resp)
}
