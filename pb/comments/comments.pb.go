// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comments.proto

package comments

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/tuneinc/truss/deftree/googlethirdparty"

import (
	context "golang.org/x/net/context"
	newcontext "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PostReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReq) Reset()         { *m = PostReq{} }
func (m *PostReq) String() string { return proto.CompactTextString(m) }
func (*PostReq) ProtoMessage()    {}
func (*PostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_comments_6675062b8b2123c9, []int{0}
}
func (m *PostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReq.Unmarshal(m, b)
}
func (m *PostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReq.Marshal(b, m, deterministic)
}
func (dst *PostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReq.Merge(dst, src)
}
func (m *PostReq) XXX_Size() int {
	return xxx_messageInfo_PostReq.Size(m)
}
func (m *PostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReq.DiscardUnknown(m)
}

var xxx_messageInfo_PostReq proto.InternalMessageInfo

func (m *PostReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PostReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PostResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResp) Reset()         { *m = PostResp{} }
func (m *PostResp) String() string { return proto.CompactTextString(m) }
func (*PostResp) ProtoMessage()    {}
func (*PostResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_comments_6675062b8b2123c9, []int{1}
}
func (m *PostResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResp.Unmarshal(m, b)
}
func (m *PostResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResp.Marshal(b, m, deterministic)
}
func (dst *PostResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResp.Merge(dst, src)
}
func (m *PostResp) XXX_Size() int {
	return xxx_messageInfo_PostResp.Size(m)
}
func (m *PostResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResp.DiscardUnknown(m)
}

var xxx_messageInfo_PostResp proto.InternalMessageInfo

func (m *PostResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PostResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_comments_6675062b8b2123c9, []int{2}
}
func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (dst *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(dst, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetResp struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []*Comment `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_comments_6675062b8b2123c9, []int{3}
}
func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (dst *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(dst, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetResp) GetData() []*Comment {
	if m != nil {
		return m.Data
	}
	return nil
}

type Comment struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_comments_6675062b8b2123c9, []int{4}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (dst *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(dst, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*PostReq)(nil), "comments.PostReq")
	proto.RegisterType((*PostResp)(nil), "comments.PostResp")
	proto.RegisterType((*GetReq)(nil), "comments.GetReq")
	proto.RegisterType((*GetResp)(nil), "comments.GetResp")
	proto.RegisterType((*Comment)(nil), "comments.Comment")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookCommentsClient is the client API for BookComments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookCommentsClient interface {
	Post(ctx context.Context, in *PostReq, opts ...grpc.CallOption) (*PostResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type bookCommentsClient struct {
	cc *grpc.ClientConn
}

func NewBookCommentsClient(cc *grpc.ClientConn) BookCommentsClient {
	return &bookCommentsClient{cc}
}

func (c *bookCommentsClient) Post(ctx context.Context, in *PostReq, opts ...grpc.CallOption) (*PostResp, error) {
	out := new(PostResp)
	err := c.cc.Invoke(ctx, "/comments.BookComments/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookCommentsClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/comments.BookComments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookCommentsServer is the server API for BookComments service.
type BookCommentsServer interface {
	Post(newcontext.Context, *PostReq) (*PostResp, error)
	Get(newcontext.Context, *GetReq) (*GetResp, error)
}

func RegisterBookCommentsServer(s *grpc.Server, srv BookCommentsServer) {
	s.RegisterService(&_BookComments_serviceDesc, srv)
}

func _BookComments_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookCommentsServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comments.BookComments/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookCommentsServer).Post(ctx, req.(*PostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookComments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookCommentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comments.BookComments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookCommentsServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookComments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comments.BookComments",
	HandlerType: (*BookCommentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _BookComments_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookComments_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments.proto",
}

func init() { proto.RegisterFile("comments.proto", fileDescriptor_comments_6675062b8b2123c9) }

var fileDescriptor_comments_6675062b8b2123c9 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4e, 0xea, 0x50,
	0x10, 0x4d, 0x5b, 0x1e, 0x85, 0x79, 0xef, 0xf1, 0x60, 0xde, 0xa6, 0x21, 0x2e, 0xc8, 0x4d, 0x4c,
	0x88, 0x0b, 0xaa, 0xb0, 0xd3, 0x15, 0xb0, 0x60, 0x6b, 0xba, 0x70, 0x5f, 0x7a, 0xc7, 0xd2, 0x68,
	0xef, 0xd4, 0xde, 0xc1, 0xc4, 0xad, 0x3f, 0xe0, 0xc2, 0x4f, 0xf3, 0x17, 0xfc, 0x10, 0x43, 0x29,
	0x82, 0xc4, 0x85, 0xbb, 0x33, 0x67, 0xee, 0x39, 0x93, 0x73, 0x2e, 0x74, 0x12, 0xce, 0x73, 0x32,
	0x62, 0x47, 0x45, 0xc9, 0xc2, 0xd8, 0xda, 0xcd, 0xfd, 0x79, 0x9a, 0xc9, 0x6a, 0xbd, 0x1c, 0x25,
	0x9c, 0x87, 0xb2, 0x36, 0x94, 0x99, 0x24, 0x94, 0x72, 0x6d, 0x6d, 0xa8, 0xe9, 0x56, 0x4a, 0xa2,
	0x30, 0x65, 0x4e, 0xef, 0x49, 0x56, 0x59, 0xa9, 0x8b, 0xb8, 0x94, 0xa7, 0x30, 0x36, 0x86, 0x25,
	0x96, 0x8c, 0x4d, 0x6d, 0xa7, 0x26, 0xe0, 0x5f, 0xb3, 0x95, 0x88, 0x1e, 0xb0, 0x03, 0x6e, 0xa6,
	0x03, 0x67, 0xe0, 0x0c, 0xbd, 0xc8, 0xcd, 0x34, 0x06, 0xe0, 0x27, 0x6c, 0x84, 0x8c, 0x04, 0xee,
	0xc0, 0x19, 0xb6, 0xa3, 0xdd, 0xa8, 0xce, 0xa1, 0xb5, 0x15, 0xd9, 0x02, 0x11, 0x1a, 0x09, 0x6b,
	0xaa, 0x74, 0xbf, 0xa2, 0x0a, 0x63, 0x17, 0xbc, 0xdc, 0xa6, 0xb5, 0x6a, 0x03, 0x55, 0x00, 0xcd,
	0x05, 0x7d, 0x77, 0x45, 0xdd, 0x80, 0x5f, 0x6d, 0x7e, 0x6a, 0x85, 0xa7, 0xd0, 0xd0, 0xb1, 0xc4,
	0x81, 0x37, 0xf0, 0x86, 0xbf, 0xc7, 0xbd, 0xd1, 0x67, 0x3f, 0xf3, 0x2d, 0x88, 0xaa, 0xb5, 0x9a,
	0x82, 0x5f, 0x13, 0x87, 0x41, 0x9c, 0x2f, 0x41, 0xf0, 0x04, 0xda, 0x49, 0x49, 0xb1, 0x90, 0x9e,
	0xee, 0x42, 0xee, 0x89, 0xf1, 0x8b, 0x03, 0x7f, 0x66, 0xcc, 0x77, 0xb5, 0x8f, 0xc5, 0x19, 0x34,
	0x36, 0xb9, 0xf1, 0xe0, 0x68, 0x5d, 0x5e, 0x1f, 0x8f, 0x29, 0x5b, 0xa8, 0xff, 0xcf, 0x6f, 0xef,
	0xaf, 0xee, 0x5f, 0xd5, 0x0a, 0x1f, 0x2f, 0xc2, 0x82, 0xad, 0x5c, 0x3a, 0x67, 0x78, 0x05, 0xde,
	0x82, 0x04, 0xbb, 0xfb, 0xf7, 0xdb, 0x62, 0xfa, 0xbd, 0x23, 0xc6, 0x16, 0xea, 0x5f, 0x65, 0xd0,
	0x46, 0x7f, 0x63, 0x90, 0x92, 0x2c, 0x9b, 0xd5, 0xa7, 0x4d, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x69, 0xdc, 0xdd, 0x15, 0x02, 0x00, 0x00,
}