// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserRequest struct {
	Id                   int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Designation          string                 `protobuf:"bytes,4,opt,name=designation,proto3" json:"designation,omitempty"`
	Team                 *UserRequest_Team      `protobuf:"bytes,5,opt,name=team,proto3" json:"team,omitempty"`
	Addresses            []*UserRequest_Address `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetDesignation() string {
	if m != nil {
		return m.Designation
	}
	return ""
}

func (m *UserRequest) GetTeam() *UserRequest_Team {
	if m != nil {
		return m.Team
	}
	return nil
}

func (m *UserRequest) GetAddresses() []*UserRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type UserRequest_Team struct {
	TeamId               int32    `protobuf:"varint,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	TeamName             string   `protobuf:"bytes,2,opt,name=teamName,proto3" json:"teamName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest_Team) Reset()         { *m = UserRequest_Team{} }
func (m *UserRequest_Team) String() string { return proto.CompactTextString(m) }
func (*UserRequest_Team) ProtoMessage()    {}
func (*UserRequest_Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0, 0}
}

func (m *UserRequest_Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest_Team.Unmarshal(m, b)
}
func (m *UserRequest_Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest_Team.Marshal(b, m, deterministic)
}
func (m *UserRequest_Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest_Team.Merge(m, src)
}
func (m *UserRequest_Team) XXX_Size() int {
	return xxx_messageInfo_UserRequest_Team.Size(m)
}
func (m *UserRequest_Team) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest_Team.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest_Team proto.InternalMessageInfo

func (m *UserRequest_Team) GetTeamId() int32 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *UserRequest_Team) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

type UserRequest_Address struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest_Address) Reset()         { *m = UserRequest_Address{} }
func (m *UserRequest_Address) String() string { return proto.CompactTextString(m) }
func (*UserRequest_Address) ProtoMessage()    {}
func (*UserRequest_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0, 1}
}

func (m *UserRequest_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest_Address.Unmarshal(m, b)
}
func (m *UserRequest_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest_Address.Marshal(b, m, deterministic)
}
func (m *UserRequest_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest_Address.Merge(m, src)
}
func (m *UserRequest_Address) XXX_Size() int {
	return xxx_messageInfo_UserRequest_Address.Size(m)
}
func (m *UserRequest_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest_Address.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest_Address proto.InternalMessageInfo

func (m *UserRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *UserRequest_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type UserFilter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFilter) Reset()         { *m = UserFilter{} }
func (m *UserFilter) String() string { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()    {}
func (*UserFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFilter.Unmarshal(m, b)
}
func (m *UserFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFilter.Marshal(b, m, deterministic)
}
func (m *UserFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFilter.Merge(m, src)
}
func (m *UserFilter) XXX_Size() int {
	return xxx_messageInfo_UserFilter.Size(m)
}
func (m *UserFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFilter.DiscardUnknown(m)
}

var xxx_messageInfo_UserFilter proto.InternalMessageInfo

func (m *UserFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type UserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
	proto.RegisterType((*UserRequest_Team)(nil), "user.UserRequest.Team")
	proto.RegisterType((*UserRequest_Address)(nil), "user.UserRequest.Address")
	proto.RegisterType((*UserFilter)(nil), "user.UserFilter")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0x6d, 0xd2, 0xf4, 0x6b, 0x72, 0xb9, 0xdc, 0x3b, 0x48, 0x59, 0xf3, 0x14, 0xf2, 0x20, 0xc1,
	0x87, 0x22, 0x2d, 0xa2, 0xf8, 0x26, 0x82, 0xe2, 0x83, 0x3e, 0x04, 0xfd, 0x01, 0x6b, 0x32, 0xc8,
	0x62, 0x93, 0xd4, 0x9d, 0x0d, 0xd2, 0x9f, 0xea, 0xbf, 0x91, 0xdd, 0x24, 0xb6, 0x90, 0xb7, 0x73,
	0x26, 0x67, 0xce, 0xcc, 0x99, 0x2c, 0x40, 0xc3, 0xa4, 0x57, 0x3b, 0x5d, 0x9b, 0x1a, 0x03, 0x8b,
	0x93, 0x6f, 0x1f, 0xc2, 0x57, 0x26, 0x9d, 0xd1, 0x67, 0x43, 0x6c, 0xf0, 0x2f, 0xf8, 0xaa, 0x10,
	0x5e, 0xec, 0xa5, 0x93, 0xcc, 0x57, 0x05, 0x22, 0x04, 0x95, 0x2c, 0x49, 0xf8, 0xb1, 0x97, 0x2e,
	0x32, 0x87, 0xf1, 0x04, 0x26, 0x54, 0x4a, 0xb5, 0x15, 0x63, 0x57, 0x6c, 0x09, 0xc6, 0x10, 0x16,
	0xc4, 0xea, 0xbd, 0x92, 0x46, 0xd5, 0x95, 0x08, 0xdc, 0xb7, 0xe3, 0x12, 0x9e, 0x43, 0x60, 0x48,
	0x96, 0x62, 0x12, 0x7b, 0x69, 0xb8, 0x5e, 0xae, 0xdc, 0x32, 0x47, 0xc3, 0x57, 0x2f, 0x24, 0xcb,
	0xcc, 0x69, 0xf0, 0x0a, 0x16, 0xb2, 0x28, 0x34, 0x31, 0x13, 0x8b, 0x69, 0x3c, 0x4e, 0xc3, 0xf5,
	0xe9, 0xb0, 0xe1, 0xb6, 0x95, 0x64, 0x07, 0x6d, 0x74, 0x03, 0x81, 0xb5, 0xc1, 0x25, 0x4c, 0xad,
	0xd1, 0x63, 0x1f, 0xa6, 0x63, 0x18, 0xc1, 0xdc, 0xa2, 0xe7, 0x43, 0xa8, 0x5f, 0x1e, 0x3d, 0xc1,
	0xac, 0x73, 0xb4, 0xb9, 0x73, 0x65, 0xf6, 0xae, 0x79, 0x91, 0x39, 0x6c, 0x73, 0xb3, 0x91, 0xa6,
	0xef, 0x6b, 0x09, 0x0a, 0x98, 0xe5, 0x75, 0x53, 0x19, 0xbd, 0xef, 0xee, 0xd1, 0xd3, 0xe4, 0x0c,
	0xc0, 0x2e, 0x7b, 0xaf, 0xb6, 0x86, 0xb4, 0xd5, 0x7d, 0xd0, 0xfe, 0xab, 0xd6, 0x45, 0x67, 0xda,
	0xd3, 0xe4, 0x1a, 0xfe, 0xb4, 0xa1, 0x78, 0x57, 0x57, 0x4c, 0x83, 0x7f, 0x20, 0x60, 0xc6, 0x4d,
	0x9e, 0x13, 0xb3, 0x9b, 0x3c, 0xcf, 0x7a, 0xba, 0xd6, 0x10, 0xd8, 0x4e, 0xdc, 0xc0, 0xfc, 0x81,
	0x8c, 0x85, 0x8c, 0xff, 0x0e, 0x67, 0x6a, 0x27, 0x47, 0xff, 0x07, 0x87, 0x4b, 0x46, 0x17, 0x1e,
	0x5e, 0x02, 0xdc, 0x69, 0x92, 0x86, 0x9c, 0xc5, 0x50, 0x14, 0xe1, 0x71, 0xa9, 0xdd, 0x2d, 0x19,
	0xbd, 0x4d, 0xdd, 0xf3, 0xd9, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xe4, 0x0b, 0xca, 0x4c,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error)
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_User_serviceDesc.Streams[0], "/user.User/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_GetUsersClient interface {
	Recv() (*UserRequest, error)
	grpc.ClientStream
}

type userGetUsersClient struct {
	grpc.ClientStream
}

func (x *userGetUsersClient) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUsers(*UserFilter, User_GetUsersServer) error
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUsers(req *UserFilter, srv User_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).GetUsers(m, &userGetUsersServer{stream})
}

type User_GetUsersServer interface {
	Send(*UserRequest) error
	grpc.ServerStream
}

type userGetUsersServer struct {
	grpc.ServerStream
}

func (x *userGetUsersServer) Send(m *UserRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _User_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
