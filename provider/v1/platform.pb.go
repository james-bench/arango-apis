// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platform.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/arangodb-managed/apis/common/v1"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	io "io"
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

// Provider represents a specific cloud provider such as AWS or GCP.
type Provider struct {
	// System identifier of the provider.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the provider
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Provider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List of providers.
type ProviderList struct {
	Items                []*Provider `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProviderList) Reset()         { *m = ProviderList{} }
func (m *ProviderList) String() string { return proto.CompactTextString(m) }
func (*ProviderList) ProtoMessage()    {}
func (*ProviderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{1}
}
func (m *ProviderList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderList.Merge(m, src)
}
func (m *ProviderList) XXX_Size() int {
	return m.Size()
}
func (m *ProviderList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderList.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderList proto.InternalMessageInfo

func (m *ProviderList) GetItems() []*Provider {
	if m != nil {
		return m.Items
	}
	return nil
}

// Region represents a geographical region in which deployments are run.
type Region struct {
	// System identifier of the region.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Identifier of the provider that hosts this region.
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// Name of the region
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Is this region available for creating new deployments?
	Available            bool     `protobuf:"varint,4,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Region) Reset()         { *m = Region{} }
func (m *Region) String() string { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()    {}
func (*Region) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{2}
}
func (m *Region) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Region) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Region.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Region) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Region.Merge(m, src)
}
func (m *Region) XXX_Size() int {
	return m.Size()
}
func (m *Region) XXX_DiscardUnknown() {
	xxx_messageInfo_Region.DiscardUnknown(m)
}

var xxx_messageInfo_Region proto.InternalMessageInfo

func (m *Region) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Region) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

// List of regions.
type RegionList struct {
	Items                []*Region `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegionList) Reset()         { *m = RegionList{} }
func (m *RegionList) String() string { return proto.CompactTextString(m) }
func (*RegionList) ProtoMessage()    {}
func (*RegionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{3}
}
func (m *RegionList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegionList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionList.Merge(m, src)
}
func (m *RegionList) XXX_Size() int {
	return m.Size()
}
func (m *RegionList) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionList.DiscardUnknown(m)
}

var xxx_messageInfo_RegionList proto.InternalMessageInfo

func (m *RegionList) GetItems() []*Region {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Provider)(nil), "arangodb.cloud.platform.v1.Provider")
	proto.RegisterType((*ProviderList)(nil), "arangodb.cloud.platform.v1.ProviderList")
	proto.RegisterType((*Region)(nil), "arangodb.cloud.platform.v1.Region")
	proto.RegisterType((*RegionList)(nil), "arangodb.cloud.platform.v1.RegionList")
}

func init() { proto.RegisterFile("platform.proto", fileDescriptor_918f3d50bfb447e4) }

var fileDescriptor_918f3d50bfb447e4 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0xc7, 0x9d, 0xed, 0x5a, 0xba, 0xbf, 0xd5, 0x0a, 0x73, 0x90, 0x25, 0xac, 0xbb, 0xeb, 0x68,
	0x75, 0x0f, 0x9a, 0x21, 0xf5, 0xa0, 0x78, 0xf0, 0x20, 0x62, 0xa9, 0x08, 0x96, 0x78, 0xf3, 0x52,
	0x26, 0x99, 0x31, 0x0e, 0x26, 0x33, 0x21, 0x99, 0x06, 0xa1, 0xf4, 0xa0, 0xf8, 0x06, 0x22, 0xf8,
	0x1a, 0xbe, 0x85, 0x47, 0xc1, 0x17, 0x90, 0xd5, 0x07, 0x91, 0x4c, 0x26, 0xbb, 0x6a, 0xd9, 0x36,
	0xb7, 0xe1, 0xf7, 0xef, 0xfb, 0xd9, 0xef, 0x77, 0x03, 0xdb, 0x79, 0xca, 0xcc, 0x6b, 0x5d, 0x64,
	0x7e, 0x5e, 0x68, 0xa3, 0xb1, 0xc7, 0x0a, 0xa6, 0x12, 0xcd, 0x23, 0x3f, 0x4e, 0xf5, 0x11, 0xf7,
	0x97, 0xed, 0x2a, 0xf0, 0xae, 0xc6, 0x3a, 0xcb, 0xb4, 0xa2, 0x55, 0x40, 0x9b, 0x57, 0xb3, 0xe3,
	0x8d, 0x13, 0xad, 0x93, 0x54, 0x50, 0x96, 0x4b, 0xca, 0x94, 0xd2, 0x86, 0x19, 0xa9, 0x55, 0xd9,
	0x74, 0x89, 0x0f, 0x5b, 0x07, 0x85, 0xae, 0x24, 0x17, 0x05, 0xde, 0x86, 0x9e, 0xe4, 0x23, 0x34,
	0x43, 0xf3, 0x41, 0xd8, 0x93, 0x1c, 0x63, 0xe8, 0x2b, 0x96, 0x89, 0x51, 0xcf, 0x56, 0xec, 0x9b,
	0x3c, 0x83, 0x4b, 0xed, 0xfc, 0x73, 0x59, 0x1a, 0xfc, 0x10, 0x2e, 0x4a, 0x23, 0xb2, 0x72, 0x84,
	0x66, 0x1b, 0xf3, 0xe1, 0xee, 0x4d, 0x7f, 0x3d, 0xa1, 0xdf, 0x2e, 0x86, 0xcd, 0x0a, 0x79, 0x0b,
	0x9b, 0xa1, 0x48, 0xa4, 0x56, 0xa7, 0x94, 0xa7, 0x30, 0xcc, 0xdd, 0xf0, 0xa1, 0xe4, 0x0e, 0x00,
	0xda, 0xd2, 0xfe, 0x0a, 0x6d, 0x63, 0x85, 0x86, 0xc7, 0x30, 0x60, 0x15, 0x93, 0x29, 0x8b, 0x52,
	0x31, 0xea, 0xcf, 0xd0, 0x7c, 0x2b, 0x5c, 0x15, 0xc8, 0x53, 0x80, 0x46, 0xcc, 0x62, 0x3f, 0xf8,
	0x17, 0x9b, 0x9c, 0x85, 0xdd, 0xac, 0x39, 0xe8, 0xdd, 0xaf, 0x7d, 0xb8, 0x72, 0xe0, 0xba, 0x2f,
	0x45, 0x51, 0xc9, 0x58, 0xe0, 0x8f, 0x08, 0x2e, 0xd7, 0x67, 0xdb, 0x1f, 0x58, 0xe2, 0x9d, 0xff,
	0x0f, 0xba, 0x48, 0xaa, 0xc0, 0xaf, 0x07, 0x5f, 0xe4, 0x36, 0x03, 0x6f, 0xde, 0xc5, 0xae, 0x7a,
	0x81, 0x90, 0x0f, 0x3f, 0x7e, 0x7f, 0xea, 0x8d, 0xb1, 0x67, 0x73, 0x6c, 0xc7, 0xea, 0xb0, 0xf3,
	0xa5, 0xe8, 0x7b, 0x04, 0xc3, 0x3d, 0xb1, 0xa4, 0xc0, 0x37, 0xd6, 0x43, 0xec, 0x3f, 0x69, 0x11,
	0x3a, 0x25, 0x46, 0x6e, 0x5b, 0xf9, 0xeb, 0x78, 0xba, 0x5e, 0x9e, 0x1e, 0x4b, 0x7e, 0x82, 0x3f,
	0x23, 0x18, 0xd6, 0xc0, 0x8d, 0x69, 0x9d, 0x8d, 0xb8, 0x75, 0x7e, 0x00, 0xd6, 0x86, 0xfb, 0x96,
	0x23, 0xc0, 0xf4, 0x2c, 0x8e, 0x58, 0x2b, 0x23, 0xde, 0x99, 0x43, 0xc9, 0x4f, 0x68, 0xe1, 0x38,
	0x8e, 0x61, 0xb0, 0x27, 0x1c, 0x55, 0x37, 0x63, 0x3a, 0xfc, 0x27, 0xc8, 0x8e, 0xc5, 0x99, 0xe2,
	0x6b, 0xa7, 0x70, 0x9c, 0xae, 0x35, 0xe5, 0xf1, 0xa3, 0x6f, 0x8b, 0x09, 0xfa, 0xbe, 0x98, 0xa0,
	0x9f, 0x8b, 0x09, 0xfa, 0xf2, 0x6b, 0x72, 0xe1, 0xd5, 0x9d, 0x44, 0x9a, 0x37, 0x47, 0x51, 0x2d,
	0x4c, 0x5b, 0x99, 0xbb, 0x19, 0x53, 0x2c, 0x11, 0xbc, 0xbe, 0x55, 0xfe, 0x7d, 0x2c, 0xda, 0xb4,
	0xdf, 0xea, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x4f, 0x77, 0xf7, 0x0f, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlatformServiceClient is the client API for PlatformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformServiceClient interface {
	// Fetch all providers that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListProviders(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*ProviderList, error)
	// Fetch a provider by its id.
	// Required permissions:
	// - None
	GetProvider(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Provider, error)
	// Fetch all regions provided by the provided identified by the given context ID.
	// Required permissions:
	// - None
	ListRegions(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*RegionList, error)
	// Fetch a region by its id.
	// Required permissions:
	// - None
	GetRegion(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Region, error)
}

type platformServiceClient struct {
	cc *grpc.ClientConn
}

func NewPlatformServiceClient(cc *grpc.ClientConn) PlatformServiceClient {
	return &platformServiceClient{cc}
}

func (c *platformServiceClient) ListProviders(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*ProviderList, error) {
	out := new(ProviderList)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.platform.v1.PlatformService/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetProvider(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Provider, error) {
	out := new(Provider)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.platform.v1.PlatformService/GetProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) ListRegions(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*RegionList, error) {
	out := new(RegionList)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.platform.v1.PlatformService/ListRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetRegion(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.platform.v1.PlatformService/GetRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServiceServer is the server API for PlatformService service.
type PlatformServiceServer interface {
	// Fetch all providers that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListProviders(context.Context, *v1.ListOptions) (*ProviderList, error)
	// Fetch a provider by its id.
	// Required permissions:
	// - None
	GetProvider(context.Context, *v1.IDOptions) (*Provider, error)
	// Fetch all regions provided by the provided identified by the given context ID.
	// Required permissions:
	// - None
	ListRegions(context.Context, *v1.ListOptions) (*RegionList, error)
	// Fetch a region by its id.
	// Required permissions:
	// - None
	GetRegion(context.Context, *v1.IDOptions) (*Region, error)
}

func RegisterPlatformServiceServer(s *grpc.Server, srv PlatformServiceServer) {
	s.RegisterService(&_PlatformService_serviceDesc, srv)
}

func _PlatformService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.platform.v1.PlatformService/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ListProviders(ctx, req.(*v1.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.platform.v1.PlatformService/GetProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetProvider(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.platform.v1.PlatformService/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ListRegions(ctx, req.(*v1.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.platform.v1.PlatformService/GetRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetRegion(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlatformService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.platform.v1.PlatformService",
	HandlerType: (*PlatformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProviders",
			Handler:    _PlatformService_ListProviders_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _PlatformService_GetProvider_Handler,
		},
		{
			MethodName: "ListRegions",
			Handler:    _PlatformService_ListRegions_Handler,
		},
		{
			MethodName: "GetRegion",
			Handler:    _PlatformService_GetRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProviderList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPlatform(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Region) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Region) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.ProviderId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.ProviderId)))
		i += copy(dAtA[i:], m.ProviderId)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPlatform(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Available {
		dAtA[i] = 0x20
		i++
		if m.Available {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RegionList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegionList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPlatform(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPlatform(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProviderList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovPlatform(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Region) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	l = len(m.ProviderId)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlatform(uint64(l))
	}
	if m.Available {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegionList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovPlatform(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlatform(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlatform(x uint64) (n int) {
	return sovPlatform(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlatform
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlatform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlatform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProviderList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlatform
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProviderList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Provider{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlatform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlatform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Region) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlatform
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Region: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Region: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Available = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlatform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlatform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegionList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlatform
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlatform
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Region{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlatform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlatform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlatform(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlatform
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlatform
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPlatform
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlatform
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPlatform(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPlatform = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlatform   = fmt.Errorf("proto: integer overflow")
)
