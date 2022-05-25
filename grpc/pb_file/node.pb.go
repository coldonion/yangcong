// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package node is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	RecallRequest
	RecallReply
	User
	Meta
	LikeRequest
	LikeReply
	DislikeRequest
	DislikeReply
	AllRequest
	AllReply
*/
package pb_file

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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

type RecallRequest struct {
	UserId int32   `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	X      float32 `protobuf:"fixed32,2,opt,name=X" json:"X,omitempty"`
	Y      float32 `protobuf:"fixed32,3,opt,name=Y" json:"Y,omitempty"`
	Gender int32   `protobuf:"varint,4,opt,name=Gender" json:"Gender,omitempty"`
	Radius int32   `protobuf:"varint,5,opt,name=Radius" json:"Radius,omitempty"`
}

func (m *RecallRequest) Reset()                    { *m = RecallRequest{} }
func (m *RecallRequest) String() string            { return proto.CompactTextString(m) }
func (*RecallRequest) ProtoMessage()               {}
func (*RecallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecallRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RecallRequest) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RecallRequest) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *RecallRequest) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *RecallRequest) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

type RecallReply struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
	Me    *User   `protobuf:"bytes,2,opt,name=Me" json:"Me,omitempty"`
	Meta  *Meta   `protobuf:"bytes,3,opt,name=Meta" json:"Meta,omitempty"`
}

func (m *RecallReply) Reset()                    { *m = RecallReply{} }
func (m *RecallReply) String() string            { return proto.CompactTextString(m) }
func (*RecallReply) ProtoMessage()               {}
func (*RecallReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RecallReply) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *RecallReply) GetMe() *User {
	if m != nil {
		return m.Me
	}
	return nil
}

func (m *RecallReply) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type User struct {
	UserId   int32   `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Tags     uint32  `protobuf:"varint,2,opt,name=Tags" json:"Tags,omitempty"`
	Gender   int32   `protobuf:"varint,3,opt,name=Gender" json:"Gender,omitempty"`
	Pop      float32 `protobuf:"fixed32,4,opt,name=Pop" json:"Pop,omitempty"`
	Distance float32 `protobuf:"fixed32,5,opt,name=Distance" json:"Distance,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetTags() uint32 {
	if m != nil {
		return m.Tags
	}
	return 0
}

func (m *User) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *User) GetPop() float32 {
	if m != nil {
		return m.Pop
	}
	return 0
}

func (m *User) GetDistance() float32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

type Meta struct {
	Code    string `protobuf:"bytes,1,opt,name=Code" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *Meta) Reset()                    { *m = Meta{} }
func (m *Meta) String() string            { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()               {}
func (*Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Meta) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Meta) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LikeRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *LikeRequest) Reset()                    { *m = LikeRequest{} }
func (m *LikeRequest) String() string            { return proto.CompactTextString(m) }
func (*LikeRequest) ProtoMessage()               {}
func (*LikeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LikeRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type LikeReply struct {
	Ids  map[int32]string `protobuf:"bytes,1,rep,name=Ids" json:"Ids,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Meta *Meta            `protobuf:"bytes,2,opt,name=Meta" json:"Meta,omitempty"`
}

func (m *LikeReply) Reset()                    { *m = LikeReply{} }
func (m *LikeReply) String() string            { return proto.CompactTextString(m) }
func (*LikeReply) ProtoMessage()               {}
func (*LikeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LikeReply) GetIds() map[int32]string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *LikeReply) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DislikeRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *DislikeRequest) Reset()                    { *m = DislikeRequest{} }
func (m *DislikeRequest) String() string            { return proto.CompactTextString(m) }
func (*DislikeRequest) ProtoMessage()               {}
func (*DislikeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DislikeRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DislikeReply struct {
	Ids  map[int32]string `protobuf:"bytes,1,rep,name=Ids" json:"Ids,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Meta *Meta            `protobuf:"bytes,2,opt,name=Meta" json:"Meta,omitempty"`
}

func (m *DislikeReply) Reset()                    { *m = DislikeReply{} }
func (m *DislikeReply) String() string            { return proto.CompactTextString(m) }
func (*DislikeReply) ProtoMessage()               {}
func (*DislikeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DislikeReply) GetIds() map[int32]string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *DislikeReply) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type AllRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *AllRequest) Reset()                    { *m = AllRequest{} }
func (m *AllRequest) String() string            { return proto.CompactTextString(m) }
func (*AllRequest) ProtoMessage()               {}
func (*AllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AllReply struct {
	Ids  map[int32]string `protobuf:"bytes,1,rep,name=Ids" json:"Ids,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Meta *Meta            `protobuf:"bytes,2,opt,name=Meta" json:"Meta,omitempty"`
}

func (m *AllReply) Reset()                    { *m = AllReply{} }
func (m *AllReply) String() string            { return proto.CompactTextString(m) }
func (*AllReply) ProtoMessage()               {}
func (*AllReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AllReply) GetIds() map[int32]string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *AllReply) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*RecallRequest)(nil), "node.RecallRequest")
	proto.RegisterType((*RecallReply)(nil), "node.RecallReply")
	proto.RegisterType((*User)(nil), "node.User")
	proto.RegisterType((*Meta)(nil), "node.Meta")
	proto.RegisterType((*LikeRequest)(nil), "node.LikeRequest")
	proto.RegisterType((*LikeReply)(nil), "node.LikeReply")
	proto.RegisterType((*DislikeRequest)(nil), "node.DislikeRequest")
	proto.RegisterType((*DislikeReply)(nil), "node.DislikeReply")
	proto.RegisterType((*AllRequest)(nil), "node.AllRequest")
	proto.RegisterType((*AllReply)(nil), "node.AllReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeService service

type NodeServiceClient interface {
	Recall(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*RecallReply, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Recall(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*RecallReply, error) {
	out := new(RecallReply)
	err := grpc.Invoke(ctx, "/node.NodeService/Recall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeService service

type NodeServiceServer interface {
	Recall(context.Context, *RecallRequest) (*RecallReply, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Recall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Recall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeService/Recall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Recall(ctx, req.(*RecallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recall",
			Handler:    _NodeService_Recall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

// Client API for RelationshipService service

type RelationshipServiceClient interface {
	GetLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeReply, error)
	GetDislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*DislikeReply, error)
	GetAll(ctx context.Context, in *AllRequest, opts ...grpc.CallOption) (*AllReply, error)
}

type relationshipServiceClient struct {
	cc *grpc.ClientConn
}

func NewRelationshipServiceClient(cc *grpc.ClientConn) RelationshipServiceClient {
	return &relationshipServiceClient{cc}
}

func (c *relationshipServiceClient) GetLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeReply, error) {
	out := new(LikeReply)
	err := grpc.Invoke(ctx, "/node.RelationshipService/GetLike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) GetDislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*DislikeReply, error) {
	out := new(DislikeReply)
	err := grpc.Invoke(ctx, "/node.RelationshipService/GetDislike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) GetAll(ctx context.Context, in *AllRequest, opts ...grpc.CallOption) (*AllReply, error) {
	out := new(AllReply)
	err := grpc.Invoke(ctx, "/node.RelationshipService/GetAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RelationshipService service

type RelationshipServiceServer interface {
	GetLike(context.Context, *LikeRequest) (*LikeReply, error)
	GetDislike(context.Context, *DislikeRequest) (*DislikeReply, error)
	GetAll(context.Context, *AllRequest) (*AllReply, error)
}

func RegisterRelationshipServiceServer(s *grpc.Server, srv RelationshipServiceServer) {
	s.RegisterService(&_RelationshipService_serviceDesc, srv)
}

func _RelationshipService_GetLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).GetLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.RelationshipService/GetLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).GetLike(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_GetDislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DislikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).GetDislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.RelationshipService/GetDislike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).GetDislike(ctx, req.(*DislikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.RelationshipService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).GetAll(ctx, req.(*AllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RelationshipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.RelationshipService",
	HandlerType: (*RelationshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLike",
			Handler:    _RelationshipService_GetLike_Handler,
		},
		{
			MethodName: "GetDislike",
			Handler:    _RelationshipService_GetDislike_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RelationshipService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5f, 0x6b, 0xd3, 0x50,
	0x14, 0xe7, 0x26, 0x69, 0xd7, 0x9e, 0x6e, 0x73, 0x3b, 0x1b, 0x1a, 0x2a, 0x48, 0x09, 0x0a, 0x75,
	0xb0, 0x22, 0x55, 0x86, 0xf8, 0x22, 0xc5, 0x49, 0x29, 0x58, 0x91, 0xab, 0xc2, 0xf6, 0x18, 0x9b,
	0xc3, 0x0c, 0x0d, 0x49, 0xcc, 0xbd, 0x1d, 0xf6, 0x43, 0x88, 0x4f, 0x7e, 0x07, 0x3f, 0xa6, 0xdc,
	0x7b, 0x9a, 0x90, 0x4c, 0x65, 0x7b, 0xea, 0xdb, 0xf9, 0x9d, 0x73, 0x6e, 0xce, 0xef, 0x77, 0xfe,
	0x04, 0x20, 0xcd, 0x22, 0x1a, 0xe5, 0x45, 0xa6, 0x33, 0xf4, 0x8c, 0x1d, 0x28, 0xd8, 0x93, 0xb4,
	0x08, 0x93, 0x44, 0xd2, 0xb7, 0x15, 0x29, 0x8d, 0xf7, 0xa1, 0xfd, 0x59, 0x51, 0x31, 0x8b, 0x7c,
	0x31, 0x10, 0xc3, 0x96, 0xdc, 0x20, 0xdc, 0x05, 0x71, 0xe1, 0x3b, 0x03, 0x31, 0x74, 0xa4, 0xb8,
	0x30, 0xe8, 0xd2, 0x77, 0x19, 0x5d, 0x9a, 0x37, 0x53, 0x4a, 0x23, 0x2a, 0x7c, 0x8f, 0xdf, 0x30,
	0x32, 0x7e, 0x19, 0x46, 0xf1, 0x4a, 0xf9, 0x2d, 0xf6, 0x33, 0x0a, 0x96, 0xd0, 0x2b, 0x8b, 0xe6,
	0xc9, 0x1a, 0x07, 0xd0, 0x32, 0x45, 0x94, 0x2f, 0x06, 0xee, 0xb0, 0x37, 0x86, 0x91, 0x65, 0x69,
	0x5c, 0x92, 0x03, 0xd8, 0x07, 0x67, 0x4e, 0xb6, 0x7a, 0x33, 0xec, 0xcc, 0x09, 0x1f, 0x81, 0x37,
	0x27, 0x1d, 0x5a, 0x36, 0x55, 0xd4, 0x78, 0xa4, 0xf5, 0x07, 0xdf, 0xc1, 0x33, 0xb9, 0xff, 0x15,
	0x86, 0xe0, 0x7d, 0x0a, 0xaf, 0x94, 0xfd, 0xfa, 0x9e, 0xb4, 0x76, 0x4d, 0x90, 0xdb, 0x10, 0x74,
	0x00, 0xee, 0x87, 0x2c, 0xb7, 0x2a, 0x1d, 0x69, 0x4c, 0xec, 0x43, 0xe7, 0x3c, 0x56, 0x3a, 0x4c,
	0x17, 0x64, 0x45, 0x3a, 0xb2, 0xc2, 0xc1, 0x0b, 0x66, 0x66, 0x2a, 0xbc, 0xc9, 0x22, 0xb2, 0x75,
	0xbb, 0xd2, 0xda, 0xe8, 0xc3, 0xce, 0x9c, 0x94, 0x0a, 0xaf, 0x58, 0x56, 0x57, 0x96, 0x30, 0x78,
	0x02, 0xbd, 0x77, 0xf1, 0x92, 0x6e, 0x99, 0x47, 0xf0, 0x53, 0x40, 0x97, 0xf3, 0x4c, 0x0b, 0x4f,
	0xc0, 0x9d, 0x45, 0x65, 0x03, 0x7d, 0xee, 0x41, 0x15, 0x1d, 0xcd, 0x22, 0xf5, 0x36, 0xd5, 0xc5,
	0x5a, 0x9a, 0xa4, 0xaa, 0x61, 0xce, 0xbf, 0x1b, 0xd6, 0x3f, 0x83, 0x4e, 0xf9, 0xc0, 0x08, 0x5e,
	0xd2, 0x7a, 0x53, 0xda, 0x98, 0x78, 0x0c, 0xad, 0xeb, 0x30, 0x59, 0x95, 0xb4, 0x19, 0xbc, 0x72,
	0x5e, 0x8a, 0x60, 0x08, 0xfb, 0xe7, 0xb1, 0x4a, 0xee, 0xc0, 0xfd, 0x97, 0x80, 0xdd, 0x2a, 0xd5,
	0xd0, 0x3f, 0xad, 0xd3, 0x7f, 0xc8, 0x8c, 0xea, 0x09, 0x5b, 0x52, 0xf0, 0x18, 0x60, 0x72, 0xeb,
	0x25, 0x04, 0x3f, 0x04, 0x74, 0x26, 0xe5, 0xee, 0x3e, 0xad, 0x33, 0x7f, 0xc0, 0x4c, 0xca, 0xe0,
	0x76, 0x58, 0x8f, 0x5f, 0x43, 0xef, 0x7d, 0x16, 0xd1, 0x47, 0x2a, 0xae, 0xe3, 0x05, 0xe1, 0x33,
	0x68, 0xf3, 0x71, 0xe1, 0x11, 0x97, 0x68, 0xdc, 0x77, 0xff, 0xb0, 0xe9, 0xcc, 0x93, 0xf5, 0xf8,
	0xb7, 0x80, 0x23, 0x49, 0x49, 0xa8, 0xe3, 0x2c, 0x55, 0x5f, 0xe3, 0xbc, 0xfc, 0xd2, 0x29, 0xec,
	0x4c, 0x49, 0x9b, 0x35, 0xc2, 0xc3, 0xfa, 0x4a, 0xf1, 0x87, 0xee, 0xdd, 0xd8, 0x32, 0x3c, 0x03,
	0x98, 0x92, 0xde, 0x8c, 0x0d, 0x8f, 0x6f, 0x4c, 0x91, 0x1f, 0xe1, 0xdf, 0xb3, 0xc5, 0x13, 0x73,
	0x6c, 0x7a, 0x92, 0x24, 0x78, 0x50, 0xeb, 0x1f, 0xe7, 0xef, 0x37, 0x3b, 0xfa, 0xa5, 0x6d, 0xff,
	0x5d, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x71, 0x82, 0xe0, 0x9b, 0xc9, 0x04, 0x00, 0x00,
}
