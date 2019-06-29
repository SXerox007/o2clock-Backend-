// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api-proto/onboarding/forgotpassword/forgotpassword.proto

package forgotpasswordpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ForgotPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordRequest) Reset()         { *m = ForgotPasswordRequest{} }
func (m *ForgotPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordRequest) ProtoMessage()    {}
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d84b57999c33dad, []int{0}
}

func (m *ForgotPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordRequest.Unmarshal(m, b)
}
func (m *ForgotPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordRequest.Merge(m, src)
}
func (m *ForgotPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordRequest.Size(m)
}
func (m *ForgotPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordRequest proto.InternalMessageInfo

func (m *ForgotPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ForgotPasswordRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ForgotPasswordResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	VerifyCode           string   `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordResponse) Reset()         { *m = ForgotPasswordResponse{} }
func (m *ForgotPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordResponse) ProtoMessage()    {}
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d84b57999c33dad, []int{1}
}

func (m *ForgotPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordResponse.Unmarshal(m, b)
}
func (m *ForgotPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordResponse.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordResponse.Merge(m, src)
}
func (m *ForgotPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordResponse.Size(m)
}
func (m *ForgotPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordResponse proto.InternalMessageInfo

func (m *ForgotPasswordResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ForgotPasswordResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ForgotPasswordResponse) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

func init() {
	proto.RegisterType((*ForgotPasswordRequest)(nil), "forgotpasswordpb.ForgotPasswordRequest")
	proto.RegisterType((*ForgotPasswordResponse)(nil), "forgotpasswordpb.ForgotPasswordResponse")
}

func init() {
	proto.RegisterFile("api-proto/onboarding/forgotpassword/forgotpassword.proto", fileDescriptor_2d84b57999c33dad)
}

var fileDescriptor_2d84b57999c33dad = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcf, 0x4a, 0x33, 0x31,
	0x10, 0x67, 0xfb, 0x7d, 0x55, 0x8c, 0x17, 0x09, 0x2a, 0xeb, 0x22, 0x58, 0xf6, 0x62, 0x2f, 0x6e,
	0x50, 0x2f, 0xde, 0x0b, 0x9e, 0x65, 0xc5, 0xb3, 0x64, 0x9b, 0x69, 0x0c, 0xb4, 0x99, 0x98, 0x49,
	0xb7, 0x78, 0xf5, 0x09, 0x04, 0xdf, 0xc2, 0xd7, 0xf1, 0x15, 0x7c, 0x10, 0x31, 0x69, 0x11, 0x83,
	0xe0, 0x2d, 0xbf, 0xbf, 0x64, 0x66, 0xd8, 0x95, 0x74, 0xe6, 0xcc, 0x79, 0x0c, 0x28, 0xd0, 0x76,
	0x28, 0xbd, 0x32, 0x56, 0x8b, 0x19, 0x7a, 0x8d, 0xc1, 0x49, 0xa2, 0x15, 0x7a, 0x95, 0xc1, 0x26,
	0xda, 0xf9, 0xde, 0x4f, 0xd6, 0x75, 0xd5, 0xb1, 0x46, 0xd4, 0x73, 0x10, 0xd2, 0x19, 0x21, 0xad,
	0xc5, 0x20, 0x83, 0x41, 0x4b, 0xc9, 0x5f, 0x4f, 0xd8, 0xc1, 0x75, 0x4c, 0xdc, 0xac, 0x13, 0x2d,
	0x3c, 0x2e, 0x81, 0x02, 0xdf, 0x67, 0x43, 0x58, 0x48, 0x33, 0x2f, 0x8b, 0x51, 0x31, 0xde, 0x69,
	0x13, 0xf8, 0x62, 0xdd, 0x03, 0x5a, 0x28, 0x07, 0x89, 0x8d, 0xa0, 0xd6, 0xec, 0x30, 0x2f, 0x21,
	0x87, 0x96, 0x80, 0x97, 0x6c, 0x7b, 0x01, 0x44, 0x52, 0xc3, 0xba, 0x67, 0x03, 0x39, 0x67, 0xff,
	0xa7, 0xa8, 0x52, 0xd1, 0xb0, 0x8d, 0x6f, 0x7e, 0xc2, 0x76, 0x7b, 0xf0, 0x66, 0xf6, 0x74, 0x1f,
	0xa5, 0x7f, 0x31, 0xc1, 0x12, 0x35, 0x41, 0x05, 0x17, 0x6f, 0x45, 0xfe, 0xdd, 0x5b, 0xf0, 0xbd,
	0x99, 0x02, 0x7f, 0x29, 0xd8, 0xd1, 0xb7, 0x82, 0x2b, 0xaf, 0xee, 0x08, 0xfc, 0x46, 0x3d, 0x6d,
	0xf2, 0xb5, 0x34, 0xbf, 0x4e, 0x5d, 0x8d, 0xff, 0x36, 0xa6, 0xc9, 0xea, 0xd1, 0xf3, 0xfb, 0xc7,
	0xeb, 0xa0, 0xe2, 0xa5, 0xe8, 0xcf, 0xc5, 0x92, 0xc0, 0x8b, 0xec, 0x30, 0xdd, 0x56, 0xdc, 0xf0,
	0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x1a, 0x1e, 0x88, 0xcd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ForgotPasswordServiceClient is the client API for ForgotPasswordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForgotPasswordServiceClient interface {
	ForgotPassowrdUserService(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
}

type forgotPasswordServiceClient struct {
	cc *grpc.ClientConn
}

func NewForgotPasswordServiceClient(cc *grpc.ClientConn) ForgotPasswordServiceClient {
	return &forgotPasswordServiceClient{cc}
}

func (c *forgotPasswordServiceClient) ForgotPassowrdUserService(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/forgotpasswordpb.ForgotPasswordService/ForgotPassowrdUserService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForgotPasswordServiceServer is the server API for ForgotPasswordService service.
type ForgotPasswordServiceServer interface {
	ForgotPassowrdUserService(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
}

func RegisterForgotPasswordServiceServer(s *grpc.Server, srv ForgotPasswordServiceServer) {
	s.RegisterService(&_ForgotPasswordService_serviceDesc, srv)
}

func _ForgotPasswordService_ForgotPassowrdUserService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForgotPasswordServiceServer).ForgotPassowrdUserService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forgotpasswordpb.ForgotPasswordService/ForgotPassowrdUserService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForgotPasswordServiceServer).ForgotPassowrdUserService(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ForgotPasswordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forgotpasswordpb.ForgotPasswordService",
	HandlerType: (*ForgotPasswordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForgotPassowrdUserService",
			Handler:    _ForgotPasswordService_ForgotPassowrdUserService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api-proto/onboarding/forgotpassword/forgotpassword.proto",
}
