// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: support.proto

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

// Plan represents a specific support plan such as Bronze, Silver or Gold.
type Plan struct {
	// System identifier of the plan.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the plan.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// If set, this plan is the default support plan.
	IsDefault bool `protobuf:"varint,3,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	// Human readable description of the plan
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc54fa3a6fab31, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// List of plans.
type PlanList struct {
	Items                []*Plan  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanList) Reset()         { *m = PlanList{} }
func (m *PlanList) String() string { return proto.CompactTextString(m) }
func (*PlanList) ProtoMessage()    {}
func (*PlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc54fa3a6fab31, []int{1}
}
func (m *PlanList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlanList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanList.Merge(m, src)
}
func (m *PlanList) XXX_Size() int {
	return m.Size()
}
func (m *PlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanList.DiscardUnknown(m)
}

var xxx_messageInfo_PlanList proto.InternalMessageInfo

func (m *PlanList) GetItems() []*Plan {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Plan)(nil), "arangodb.cloud.support.v1.Plan")
	proto.RegisterType((*PlanList)(nil), "arangodb.cloud.support.v1.PlanList")
}

func init() { proto.RegisterFile("support.proto", fileDescriptor_61fc54fa3a6fab31) }

var fileDescriptor_61fc54fa3a6fab31 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xef, 0xa4, 0xbd, 0xf7, 0xb6, 0x53, 0x6e, 0x17, 0x03, 0x57, 0x63, 0x68, 0x63, 0x48,
	0x11, 0x0a, 0x62, 0x42, 0x2b, 0x2e, 0x5d, 0x28, 0x05, 0x11, 0x04, 0xa5, 0xdd, 0xb9, 0x91, 0x69,
	0x66, 0x8c, 0x83, 0xc9, 0x4c, 0xc8, 0x4c, 0xb2, 0x11, 0x37, 0xbe, 0x82, 0x1b, 0x1f, 0xc9, 0xa5,
	0xe0, 0x0b, 0x48, 0xf5, 0x35, 0x04, 0x99, 0x49, 0x0b, 0x45, 0x2c, 0xdd, 0x1d, 0xce, 0xf9, 0xcf,
	0xff, 0xcd, 0xf9, 0x07, 0xfe, 0x93, 0x45, 0x96, 0x89, 0x5c, 0x05, 0x59, 0x2e, 0x94, 0x40, 0x5b,
	0x38, 0xc7, 0x3c, 0x16, 0x64, 0x1a, 0x44, 0x89, 0x28, 0x48, 0xb0, 0x98, 0x96, 0x03, 0x67, 0x23,
	0x12, 0x69, 0x2a, 0x78, 0x58, 0x0e, 0xc2, 0xaa, 0xaa, 0x56, 0x9c, 0x4e, 0x2c, 0x44, 0x9c, 0xd0,
	0x10, 0x67, 0x2c, 0xc4, 0x9c, 0x0b, 0x85, 0x15, 0x13, 0x5c, 0x56, 0x53, 0xff, 0x16, 0xd6, 0x2f,
	0x12, 0xcc, 0x51, 0x1b, 0x5a, 0x8c, 0xd8, 0xc0, 0x03, 0xfd, 0xe6, 0xd8, 0x62, 0x04, 0x21, 0x58,
	0xe7, 0x38, 0xa5, 0xb6, 0x65, 0x3a, 0xa6, 0x46, 0x5d, 0x08, 0x99, 0xbc, 0x22, 0xf4, 0x1a, 0x17,
	0x89, 0xb2, 0x6b, 0x1e, 0xe8, 0x37, 0xc6, 0x4d, 0x26, 0x47, 0x55, 0x03, 0x79, 0xb0, 0x45, 0xa8,
	0x8c, 0x72, 0x96, 0x69, 0x80, 0x5d, 0x37, 0x9b, 0xcb, 0x2d, 0xff, 0x08, 0x36, 0x34, 0xec, 0x8c,
	0x49, 0x85, 0x0e, 0xe0, 0x6f, 0xa6, 0x68, 0x2a, 0x6d, 0xe0, 0xd5, 0xfa, 0xad, 0xe1, 0x76, 0xb0,
	0xf2, 0xb2, 0x40, 0xef, 0x8c, 0x2b, 0xf5, 0xf0, 0x13, 0xc0, 0xf6, 0xa4, 0x1a, 0x4d, 0x68, 0x5e,
	0xb2, 0x88, 0xa2, 0x12, 0x36, 0xb5, 0xa3, 0x56, 0x49, 0xb4, 0xf3, 0xdd, 0x67, 0x9e, 0x45, 0x39,
	0x08, 0xb4, 0xe8, 0xdc, 0x3c, 0x44, 0x3a, 0xbd, 0x35, 0x38, 0xad, 0xf5, 0xbb, 0x0f, 0xaf, 0x1f,
	0x8f, 0xd6, 0x26, 0xfa, 0x6f, 0xb2, 0x9b, 0x2b, 0x74, 0xbe, 0x99, 0x41, 0x49, 0xf8, 0xf7, 0x84,
	0x1a, 0x2c, 0xea, 0xad, 0xa6, 0x9e, 0x8e, 0x16, 0xcc, 0x75, 0x27, 0xfa, 0xbe, 0xe1, 0x75, 0x90,
	0xf3, 0x23, 0x2f, 0xbc, 0x63, 0xe4, 0xfe, 0xf8, 0xf0, 0x79, 0xe6, 0x82, 0x97, 0x99, 0x0b, 0xde,
	0x66, 0x2e, 0x78, 0x7a, 0x77, 0x7f, 0x5d, 0xee, 0xc6, 0x4c, 0xdd, 0x14, 0x53, 0x8d, 0x0c, 0x17,
	0x80, 0xbd, 0x14, 0x73, 0x1c, 0x53, 0xa2, 0x8d, 0xe4, 0x92, 0xd3, 0xf4, 0x8f, 0xf9, 0xf5, 0xfd,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x4e, 0x8f, 0x35, 0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SupportServiceClient is the client API for SupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupportServiceClient interface {
	// Fetch all support plans that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListPlans(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*PlanList, error)
	// Fetch a support plan by its id.
	// Required permissions:
	// - None
	GetPlan(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Plan, error)
}

type supportServiceClient struct {
	cc *grpc.ClientConn
}

func NewSupportServiceClient(cc *grpc.ClientConn) SupportServiceClient {
	return &supportServiceClient{cc}
}

func (c *supportServiceClient) ListPlans(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*PlanList, error) {
	out := new(PlanList)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.support.v1.SupportService/ListPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportServiceClient) GetPlan(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Plan, error) {
	out := new(Plan)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.support.v1.SupportService/GetPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportServiceServer is the server API for SupportService service.
type SupportServiceServer interface {
	// Fetch all support plans that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListPlans(context.Context, *v1.ListOptions) (*PlanList, error)
	// Fetch a support plan by its id.
	// Required permissions:
	// - None
	GetPlan(context.Context, *v1.IDOptions) (*Plan, error)
}

func RegisterSupportServiceServer(s *grpc.Server, srv SupportServiceServer) {
	s.RegisterService(&_SupportService_serviceDesc, srv)
}

func _SupportService_ListPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportServiceServer).ListPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.support.v1.SupportService/ListPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportServiceServer).ListPlans(ctx, req.(*v1.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportService_GetPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportServiceServer).GetPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.support.v1.SupportService/GetPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportServiceServer).GetPlan(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.support.v1.SupportService",
	HandlerType: (*SupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPlans",
			Handler:    _SupportService_ListPlans_Handler,
		},
		{
			MethodName: "GetPlan",
			Handler:    _SupportService_GetPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "support.proto",
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSupport(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSupport(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.IsDefault {
		dAtA[i] = 0x18
		i++
		if m.IsDefault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSupport(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PlanList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlanList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSupport(dAtA, i, uint64(msg.Size()))
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

func encodeVarintSupport(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSupport(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSupport(uint64(l))
	}
	if m.IsDefault {
		n += 2
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSupport(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PlanList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovSupport(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSupport(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSupport(x uint64) (n int) {
	return sovSupport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSupport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupport
			}
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
					return ErrIntOverflowSupport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSupport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsDefault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSupport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSupport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSupport
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
func (m *PlanList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PlanList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlanList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSupport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Plan{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSupport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSupport
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
func skipSupport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupport
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
					return 0, ErrIntOverflowSupport
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
					return 0, ErrIntOverflowSupport
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
			if length < 0 {
				return 0, ErrInvalidLengthSupport
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSupport
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSupport
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
				next, err := skipSupport(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSupport
				}
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
	ErrInvalidLengthSupport = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupport   = fmt.Errorf("proto: integer overflow")
)