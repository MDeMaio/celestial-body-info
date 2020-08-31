// Code generated by protoc-gen-go. DO NOT EDIT.
// source: planet/planetpb/planet.proto

package planetpb

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

// We can define our planet and all of its corresponding data here, we can split this up later if we want.
type Planet struct {
	PlanetId             string   `protobuf:"bytes,1,opt,name=planet_id,json=planetId,proto3" json:"planet_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Facts                []*Facts `protobuf:"bytes,3,rep,name=facts,proto3" json:"facts,omitempty"`
	Image                string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Planet) Reset()         { *m = Planet{} }
func (m *Planet) String() string { return proto.CompactTextString(m) }
func (*Planet) ProtoMessage()    {}
func (*Planet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{0}
}

func (m *Planet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Planet.Unmarshal(m, b)
}
func (m *Planet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Planet.Marshal(b, m, deterministic)
}
func (m *Planet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Planet.Merge(m, src)
}
func (m *Planet) XXX_Size() int {
	return xxx_messageInfo_Planet.Size(m)
}
func (m *Planet) XXX_DiscardUnknown() {
	xxx_messageInfo_Planet.DiscardUnknown(m)
}

var xxx_messageInfo_Planet proto.InternalMessageInfo

func (m *Planet) GetPlanetId() string {
	if m != nil {
		return m.PlanetId
	}
	return ""
}

func (m *Planet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Planet) GetFacts() []*Facts {
	if m != nil {
		return m.Facts
	}
	return nil
}

func (m *Planet) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type Facts struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Fact                 string   `protobuf:"bytes,2,opt,name=fact,proto3" json:"fact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Facts) Reset()         { *m = Facts{} }
func (m *Facts) String() string { return proto.CompactTextString(m) }
func (*Facts) ProtoMessage()    {}
func (*Facts) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{1}
}

func (m *Facts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Facts.Unmarshal(m, b)
}
func (m *Facts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Facts.Marshal(b, m, deterministic)
}
func (m *Facts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Facts.Merge(m, src)
}
func (m *Facts) XXX_Size() int {
	return xxx_messageInfo_Facts.Size(m)
}
func (m *Facts) XXX_DiscardUnknown() {
	xxx_messageInfo_Facts.DiscardUnknown(m)
}

var xxx_messageInfo_Facts proto.InternalMessageInfo

func (m *Facts) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Facts) GetFact() string {
	if m != nil {
		return m.Fact
	}
	return ""
}

type ReadPlanetRequest struct {
	PlanetId             string   `protobuf:"bytes,1,opt,name=planet_id,json=planetId,proto3" json:"planet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPlanetRequest) Reset()         { *m = ReadPlanetRequest{} }
func (m *ReadPlanetRequest) String() string { return proto.CompactTextString(m) }
func (*ReadPlanetRequest) ProtoMessage()    {}
func (*ReadPlanetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{2}
}

func (m *ReadPlanetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPlanetRequest.Unmarshal(m, b)
}
func (m *ReadPlanetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPlanetRequest.Marshal(b, m, deterministic)
}
func (m *ReadPlanetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPlanetRequest.Merge(m, src)
}
func (m *ReadPlanetRequest) XXX_Size() int {
	return xxx_messageInfo_ReadPlanetRequest.Size(m)
}
func (m *ReadPlanetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPlanetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPlanetRequest proto.InternalMessageInfo

func (m *ReadPlanetRequest) GetPlanetId() string {
	if m != nil {
		return m.PlanetId
	}
	return ""
}

type ReadPlanetResponse struct {
	Planet               *Planet  `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPlanetResponse) Reset()         { *m = ReadPlanetResponse{} }
func (m *ReadPlanetResponse) String() string { return proto.CompactTextString(m) }
func (*ReadPlanetResponse) ProtoMessage()    {}
func (*ReadPlanetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{3}
}

func (m *ReadPlanetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPlanetResponse.Unmarshal(m, b)
}
func (m *ReadPlanetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPlanetResponse.Marshal(b, m, deterministic)
}
func (m *ReadPlanetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPlanetResponse.Merge(m, src)
}
func (m *ReadPlanetResponse) XXX_Size() int {
	return xxx_messageInfo_ReadPlanetResponse.Size(m)
}
func (m *ReadPlanetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPlanetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPlanetResponse proto.InternalMessageInfo

func (m *ReadPlanetResponse) GetPlanet() *Planet {
	if m != nil {
		return m.Planet
	}
	return nil
}

type ListPlanetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPlanetRequest) Reset()         { *m = ListPlanetRequest{} }
func (m *ListPlanetRequest) String() string { return proto.CompactTextString(m) }
func (*ListPlanetRequest) ProtoMessage()    {}
func (*ListPlanetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{4}
}

func (m *ListPlanetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlanetRequest.Unmarshal(m, b)
}
func (m *ListPlanetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlanetRequest.Marshal(b, m, deterministic)
}
func (m *ListPlanetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlanetRequest.Merge(m, src)
}
func (m *ListPlanetRequest) XXX_Size() int {
	return xxx_messageInfo_ListPlanetRequest.Size(m)
}
func (m *ListPlanetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlanetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlanetRequest proto.InternalMessageInfo

type ListPlanetResponse struct {
	Planet               []*Planet `protobuf:"bytes,1,rep,name=planet,proto3" json:"planet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListPlanetResponse) Reset()         { *m = ListPlanetResponse{} }
func (m *ListPlanetResponse) String() string { return proto.CompactTextString(m) }
func (*ListPlanetResponse) ProtoMessage()    {}
func (*ListPlanetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7ae48ba9b9a6fb, []int{5}
}

func (m *ListPlanetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlanetResponse.Unmarshal(m, b)
}
func (m *ListPlanetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlanetResponse.Marshal(b, m, deterministic)
}
func (m *ListPlanetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlanetResponse.Merge(m, src)
}
func (m *ListPlanetResponse) XXX_Size() int {
	return xxx_messageInfo_ListPlanetResponse.Size(m)
}
func (m *ListPlanetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlanetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlanetResponse proto.InternalMessageInfo

func (m *ListPlanetResponse) GetPlanet() []*Planet {
	if m != nil {
		return m.Planet
	}
	return nil
}

func init() {
	proto.RegisterType((*Planet)(nil), "planet.Planet")
	proto.RegisterType((*Facts)(nil), "planet.Facts")
	proto.RegisterType((*ReadPlanetRequest)(nil), "planet.ReadPlanetRequest")
	proto.RegisterType((*ReadPlanetResponse)(nil), "planet.ReadPlanetResponse")
	proto.RegisterType((*ListPlanetRequest)(nil), "planet.ListPlanetRequest")
	proto.RegisterType((*ListPlanetResponse)(nil), "planet.ListPlanetResponse")
}

func init() {
	proto.RegisterFile("planet/planetpb/planet.proto", fileDescriptor_4e7ae48ba9b9a6fb)
}

var fileDescriptor_4e7ae48ba9b9a6fb = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xa6, 0x09, 0xed, 0x94, 0x0a, 0x1d, 0x3d, 0xc4, 0xea, 0xa1, 0x44, 0x90, 0x9e, 0xaa,
	0xc6, 0xab, 0x27, 0x0b, 0x05, 0xc1, 0x83, 0xc4, 0x9b, 0x17, 0xd9, 0x36, 0xa3, 0x2c, 0xb4, 0x49,
	0xec, 0xae, 0xbe, 0x8b, 0x6f, 0x2b, 0xbb, 0xb3, 0x4b, 0x2a, 0x89, 0x3d, 0x65, 0xe6, 0x9b, 0xfd,
	0x7e, 0x26, 0x0c, 0x5c, 0xd4, 0x1b, 0x51, 0x92, 0xbe, 0xe6, 0x4f, 0xbd, 0x72, 0xc5, 0xbc, 0xde,
	0x55, 0xba, 0xc2, 0x98, 0xbb, 0x54, 0x43, 0xfc, 0x6c, 0x2b, 0x3c, 0x87, 0x01, 0x63, 0x6f, 0xb2,
	0x48, 0x82, 0x69, 0x30, 0x1b, 0xe4, 0x7d, 0x06, 0x1e, 0x0b, 0x44, 0xe8, 0x95, 0x62, 0x4b, 0xc9,
	0x91, 0xc5, 0x6d, 0x8d, 0x97, 0x10, 0xbd, 0x8b, 0xb5, 0x56, 0x49, 0x38, 0x0d, 0x67, 0xc3, 0x6c,
	0x34, 0x77, 0x06, 0x4b, 0x03, 0xe6, 0x3c, 0xc3, 0x53, 0x88, 0xe4, 0x56, 0x7c, 0x50, 0xd2, 0xb3,
	0x4c, 0x6e, 0xd2, 0x5b, 0x88, 0x96, 0x7e, 0xac, 0xa5, 0xde, 0x90, 0x33, 0xe4, 0xc6, 0xb8, 0x19,
	0xb6, 0x77, 0x33, 0x75, 0x7a, 0x03, 0xe3, 0x9c, 0x44, 0xc1, 0x61, 0x73, 0xfa, 0xfc, 0x22, 0x75,
	0x38, 0x73, 0x7a, 0x0f, 0xb8, 0xcf, 0x50, 0x75, 0x55, 0x2a, 0xc2, 0x2b, 0x70, 0xab, 0xdb, 0xf7,
	0xc3, 0xec, 0xd8, 0xc7, 0x76, 0xef, 0xfc, 0x8f, 0x39, 0x81, 0xf1, 0x93, 0x54, 0xfa, 0x8f, 0x9f,
	0x91, 0xdc, 0x07, 0x3b, 0x24, 0xc3, 0xff, 0x25, 0xb3, 0x9f, 0x00, 0x46, 0x0c, 0xbd, 0xd0, 0xee,
	0x5b, 0xae, 0x09, 0x17, 0x00, 0x4d, 0x44, 0x3c, 0xf3, 0xbc, 0xd6, 0xa2, 0x93, 0x49, 0xd7, 0xc8,
	0xd9, 0x2f, 0x00, 0x9a, 0x50, 0x8d, 0x48, 0x2b, 0x7d, 0x23, 0xd2, 0xde, 0xe1, 0x01, 0x5e, 0xfb,
	0xfe, 0x50, 0x56, 0xb1, 0x3d, 0x91, 0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xe2, 0xe6,
	0x09, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlanetServiceClient is the client API for PlanetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlanetServiceClient interface {
	ReadPlanet(ctx context.Context, in *ReadPlanetRequest, opts ...grpc.CallOption) (*ReadPlanetResponse, error)
	ListPlanet(ctx context.Context, in *ListPlanetRequest, opts ...grpc.CallOption) (*ListPlanetResponse, error)
}

type planetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlanetServiceClient(cc grpc.ClientConnInterface) PlanetServiceClient {
	return &planetServiceClient{cc}
}

func (c *planetServiceClient) ReadPlanet(ctx context.Context, in *ReadPlanetRequest, opts ...grpc.CallOption) (*ReadPlanetResponse, error) {
	out := new(ReadPlanetResponse)
	err := c.cc.Invoke(ctx, "/planet.PlanetService/ReadPlanet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planetServiceClient) ListPlanet(ctx context.Context, in *ListPlanetRequest, opts ...grpc.CallOption) (*ListPlanetResponse, error) {
	out := new(ListPlanetResponse)
	err := c.cc.Invoke(ctx, "/planet.PlanetService/ListPlanet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanetServiceServer is the server API for PlanetService service.
type PlanetServiceServer interface {
	ReadPlanet(context.Context, *ReadPlanetRequest) (*ReadPlanetResponse, error)
	ListPlanet(context.Context, *ListPlanetRequest) (*ListPlanetResponse, error)
}

// UnimplementedPlanetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlanetServiceServer struct {
}

func (*UnimplementedPlanetServiceServer) ReadPlanet(ctx context.Context, req *ReadPlanetRequest) (*ReadPlanetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPlanet not implemented")
}
func (*UnimplementedPlanetServiceServer) ListPlanet(ctx context.Context, req *ListPlanetRequest) (*ListPlanetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlanet not implemented")
}

func RegisterPlanetServiceServer(s *grpc.Server, srv PlanetServiceServer) {
	s.RegisterService(&_PlanetService_serviceDesc, srv)
}

func _PlanetService_ReadPlanet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPlanetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).ReadPlanet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planet.PlanetService/ReadPlanet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).ReadPlanet(ctx, req.(*ReadPlanetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanetService_ListPlanet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlanetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).ListPlanet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planet.PlanetService/ListPlanet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).ListPlanet(ctx, req.(*ListPlanetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlanetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "planet.PlanetService",
	HandlerType: (*PlanetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadPlanet",
			Handler:    _PlanetService_ReadPlanet_Handler,
		},
		{
			MethodName: "ListPlanet",
			Handler:    _PlanetService_ListPlanet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planet/planetpb/planet.proto",
}
