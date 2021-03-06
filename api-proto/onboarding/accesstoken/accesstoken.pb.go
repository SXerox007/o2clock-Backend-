// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api-proto/onboarding/accesstoken/accesstoken.proto

package accesstokenpb

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

type AccessTokenRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenRequest) Reset()         { *m = AccessTokenRequest{} }
func (m *AccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*AccessTokenRequest) ProtoMessage()    {}
func (*AccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf525f607908d194, []int{0}
}

func (m *AccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenRequest.Unmarshal(m, b)
}
func (m *AccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *AccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenRequest.Merge(m, src)
}
func (m *AccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_AccessTokenRequest.Size(m)
}
func (m *AccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenRequest proto.InternalMessageInfo

func (m *AccessTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type AccessTokenResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenResponse) Reset()         { *m = AccessTokenResponse{} }
func (m *AccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*AccessTokenResponse) ProtoMessage()    {}
func (*AccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf525f607908d194, []int{1}
}

func (m *AccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenResponse.Unmarshal(m, b)
}
func (m *AccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *AccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenResponse.Merge(m, src)
}
func (m *AccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_AccessTokenResponse.Size(m)
}
func (m *AccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenResponse proto.InternalMessageInfo

func (m *AccessTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AccessTokenResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessTokenRequest)(nil), "accesstokenpb.AccessTokenRequest")
	proto.RegisterType((*AccessTokenResponse)(nil), "accesstokenpb.AccessTokenResponse")
}

func init() {
	proto.RegisterFile("api-proto/onboarding/accesstoken/accesstoken.proto", fileDescriptor_cf525f607908d194)
}

var fileDescriptor_cf525f607908d194 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0x2c, 0xc8, 0xd4,
	0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0xcf, 0x4b, 0xca, 0x4f, 0x2c, 0x4a, 0xc9, 0xcc, 0x4b,
	0xd7, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0x2e, 0xc9, 0xcf, 0x4e, 0xcd, 0x43, 0x66, 0xeb, 0x81,
	0x15, 0x0a, 0xf1, 0x22, 0x09, 0x15, 0x24, 0x49, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43,
	0x14, 0x2b, 0x99, 0x73, 0x09, 0x39, 0x82, 0x95, 0x87, 0x80, 0x94, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x29, 0x72, 0xf1, 0x40, 0x0c, 0x89, 0x07, 0x9b, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x9d, 0x88, 0x50, 0xa9, 0xe4, 0xcc, 0x25, 0x8c, 0xa2, 0xb1, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x15, 0xaa, 0x09,
	0xc6, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d,
	0x02, 0xb3, 0x8d, 0xa6, 0x30, 0xa2, 0x58, 0x1f, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x54,
	0xc7, 0x25, 0xee, 0x9c, 0x91, 0x9a, 0x9c, 0x8d, 0x45, 0x4a, 0x51, 0x0f, 0xc5, 0x77, 0x7a, 0x98,
	0x8e, 0x97, 0x52, 0xc2, 0xa7, 0x04, 0xe2, 0x4c, 0x25, 0x99, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x89,
	0x09, 0x89, 0xe8, 0x97, 0x19, 0xea, 0x97, 0x16, 0xa7, 0x16, 0x21, 0x87, 0x63, 0x12, 0x1b, 0x38,
	0x6c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x4a, 0x27, 0xeb, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessTokenServiceClient is the client API for AccessTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessTokenServiceClient interface {
	CheckAccessTokenService(ctx context.Context, in *AccessTokenRequest, opts ...grpc.CallOption) (*AccessTokenResponse, error)
}

type accessTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessTokenServiceClient(cc *grpc.ClientConn) AccessTokenServiceClient {
	return &accessTokenServiceClient{cc}
}

func (c *accessTokenServiceClient) CheckAccessTokenService(ctx context.Context, in *AccessTokenRequest, opts ...grpc.CallOption) (*AccessTokenResponse, error) {
	out := new(AccessTokenResponse)
	err := c.cc.Invoke(ctx, "/accesstokenpb.AccessTokenService/CheckAccessTokenService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessTokenServiceServer is the server API for AccessTokenService service.
type AccessTokenServiceServer interface {
	CheckAccessTokenService(context.Context, *AccessTokenRequest) (*AccessTokenResponse, error)
}

func RegisterAccessTokenServiceServer(s *grpc.Server, srv AccessTokenServiceServer) {
	s.RegisterService(&_AccessTokenService_serviceDesc, srv)
}

func _AccessTokenService_CheckAccessTokenService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).CheckAccessTokenService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accesstokenpb.AccessTokenService/CheckAccessTokenService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).CheckAccessTokenService(ctx, req.(*AccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accesstokenpb.AccessTokenService",
	HandlerType: (*AccessTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAccessTokenService",
			Handler:    _AccessTokenService_CheckAccessTokenService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api-proto/onboarding/accesstoken/accesstoken.proto",
}
