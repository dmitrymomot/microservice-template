// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/examplesrv/service.proto

package examplesrv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Error struct {
	Code                 int32                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              *any.Any                    `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Validation           map[string]*ValidationError `protobuf:"bytes,5,rep,name=validation,proto3" json:"validation,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ff6b146f380be9, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Error) GetValidation() map[string]*ValidationError {
	if m != nil {
		return m.Validation
	}
	return nil
}

type ValidationError struct {
	Errors               []string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationError) Reset()         { *m = ValidationError{} }
func (m *ValidationError) String() string { return proto.CompactTextString(m) }
func (*ValidationError) ProtoMessage()    {}
func (*ValidationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ff6b146f380be9, []int{1}
}

func (m *ValidationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationError.Unmarshal(m, b)
}
func (m *ValidationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationError.Marshal(b, m, deterministic)
}
func (m *ValidationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationError.Merge(m, src)
}
func (m *ValidationError) XXX_Size() int {
	return xxx_messageInfo_ValidationError.Size(m)
}
func (m *ValidationError) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationError.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationError proto.InternalMessageInfo

func (m *ValidationError) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Req struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ff6b146f380be9, []int{2}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type Resp struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ff6b146f380be9, []int{3}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Resp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "examplesrv.Error")
	proto.RegisterMapType((map[string]*ValidationError)(nil), "examplesrv.Error.ValidationEntry")
	proto.RegisterType((*ValidationError)(nil), "examplesrv.ValidationError")
	proto.RegisterType((*Req)(nil), "examplesrv.Req")
	proto.RegisterType((*Resp)(nil), "examplesrv.Resp")
}

func init() { proto.RegisterFile("pb/examplesrv/service.proto", fileDescriptor_49ff6b146f380be9) }

var fileDescriptor_49ff6b146f380be9 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0x1d, 0xdb, 0x20, 0x3c, 0x0e, 0x60, 0x63, 0x74, 0xc2, 0x65, 0xee, 0xe2, 0x8c, 0x49,
	0x17, 0x67, 0x62, 0x8c, 0x37, 0x62, 0xf8, 0x07, 0x6a, 0xe2, 0xc1, 0x5b, 0x81, 0xe7, 0xb2, 0x58,
	0xd6, 0xd1, 0x8e, 0xc5, 0xdd, 0xfd, 0xc3, 0xcd, 0x3a, 0x26, 0x43, 0x6e, 0xef, 0xf5, 0xfb, 0x5e,
	0x7f, 0x5f, 0x5f, 0x61, 0x96, 0x2f, 0x23, 0xfc, 0xe6, 0x9b, 0x5c, 0xa0, 0x56, 0x65, 0xa4, 0x51,
	0x95, 0xe9, 0x0a, 0x69, 0xae, 0x64, 0x21, 0x09, 0x1c, 0x94, 0xe9, 0x75, 0x22, 0x65, 0x22, 0x30,
	0x32, 0xca, 0x72, 0xf7, 0x19, 0xf1, 0xac, 0x6a, 0x6c, 0xc1, 0x4f, 0x0f, 0xdc, 0x85, 0x52, 0x52,
	0x11, 0x02, 0xce, 0x4a, 0xae, 0xd1, 0xb3, 0x7c, 0x2b, 0x74, 0x99, 0xa9, 0x89, 0x07, 0x83, 0x0d,
	0x6a, 0xcd, 0x13, 0xf4, 0x7a, 0xbe, 0x15, 0x0e, 0x59, 0xdb, 0x12, 0x0a, 0x83, 0x35, 0x16, 0x3c,
	0x15, 0xda, 0x73, 0x7c, 0x2b, 0x1c, 0xc5, 0x17, 0xb4, 0x81, 0xd0, 0x16, 0x42, 0xe7, 0x59, 0xc5,
	0x5a, 0x13, 0x99, 0x03, 0x94, 0x5c, 0xa4, 0x6b, 0x5e, 0xa4, 0x32, 0xf3, 0x5c, 0xdf, 0x0e, 0x47,
	0xf1, 0x0d, 0x3d, 0x64, 0xa4, 0x26, 0x04, 0x7d, 0xff, 0xf3, 0x2c, 0xb2, 0x42, 0x55, 0xac, 0x33,
	0x34, 0xfd, 0x80, 0xf1, 0x3f, 0x99, 0x4c, 0xc0, 0xfe, 0xc2, 0xca, 0x44, 0x1e, 0xb2, 0xba, 0x24,
	0x0f, 0xe0, 0x96, 0x5c, 0xec, 0x9a, 0xbc, 0xa3, 0x78, 0xd6, 0x45, 0x74, 0xa6, 0x6b, 0x18, 0x6b,
	0x9c, 0x2f, 0xbd, 0x67, 0x2b, 0xb8, 0x3b, 0xba, 0xdb, 0xec, 0xe3, 0x12, 0xfa, 0x58, 0x17, 0xda,
	0xb3, 0x7c, 0x3b, 0x1c, 0xb2, 0x7d, 0x17, 0x5c, 0x81, 0xcd, 0x70, 0x5b, 0xa3, 0x75, 0xa1, 0x5a,
	0xb4, 0x2e, 0x54, 0x30, 0x07, 0x87, 0xa1, 0xce, 0x4f, 0x15, 0x72, 0x0b, 0xae, 0x19, 0xde, 0x87,
	0x3a, 0x3f, 0x79, 0x37, 0x6b, 0xf4, 0xf8, 0x09, 0x06, 0x6f, 0xcd, 0x2f, 0x92, 0x7b, 0x70, 0x5e,
	0xb9, 0x10, 0x64, 0xdc, 0x35, 0x33, 0xdc, 0x4e, 0x27, 0xc7, 0x07, 0x3a, 0x0f, 0xce, 0x96, 0x7d,
	0xb3, 0xf4, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x9c, 0x27, 0x76, 0x12, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Call(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Call(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/examplesrv.Service/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Call(context.Context, *Req) (*Resp, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Call(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examplesrv.Service/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Call(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examplesrv.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Service_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/examplesrv/service.proto",
}
