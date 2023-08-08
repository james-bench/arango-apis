// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credits.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/arangodb-managed/apis/common/v1"
	_ "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Request for listing credit bundles
type ListCreditBundlesRequest struct {
	// ID of the organization for which credit bundles are listed.
	// This is a required field.
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	// If set, include expired bundles.
	IncludeExpired       bool     `protobuf:"varint,2,opt,name=include_expired,json=includeExpired,proto3" json:"include_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCreditBundlesRequest) Reset()         { *m = ListCreditBundlesRequest{} }
func (m *ListCreditBundlesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCreditBundlesRequest) ProtoMessage()    {}
func (*ListCreditBundlesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0331d70383bb145, []int{0}
}
func (m *ListCreditBundlesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListCreditBundlesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListCreditBundlesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListCreditBundlesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCreditBundlesRequest.Merge(m, src)
}
func (m *ListCreditBundlesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListCreditBundlesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCreditBundlesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCreditBundlesRequest proto.InternalMessageInfo

func (m *ListCreditBundlesRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *ListCreditBundlesRequest) GetIncludeExpired() bool {
	if m != nil {
		return m.IncludeExpired
	}
	return false
}

func init() {
	proto.RegisterType((*ListCreditBundlesRequest)(nil), "arangodb.cloud.credits.v1.ListCreditBundlesRequest")
}

func init() { proto.RegisterFile("credits.proto", fileDescriptor_a0331d70383bb145) }

var fileDescriptor_a0331d70383bb145 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xbf, 0x4a, 0xc3, 0x40,
	0x1c, 0xf6, 0x3a, 0x88, 0x06, 0xac, 0x98, 0x41, 0x6a, 0x91, 0x58, 0xba, 0x58, 0x10, 0xef, 0x88,
	0x1d, 0x1c, 0xc4, 0xa5, 0xd2, 0x41, 0x70, 0xaa, 0x9b, 0x4b, 0xb9, 0xe4, 0xce, 0xf3, 0x20, 0x77,
	0xbf, 0x33, 0x77, 0x09, 0x56, 0x71, 0xf1, 0x15, 0x5c, 0x9c, 0x7d, 0x07, 0xdf, 0xc1, 0x51, 0xf0,
	0x05, 0xa4, 0xfa, 0x20, 0xd2, 0x24, 0xc5, 0x60, 0xe9, 0xf6, 0xe3, 0xfb, 0xf2, 0xfd, 0xc9, 0x77,
	0xde, 0x46, 0x9c, 0x72, 0x26, 0x9d, 0xc5, 0x26, 0x05, 0x07, 0xfe, 0x0e, 0x4d, 0xa9, 0x16, 0xc0,
	0x22, 0x1c, 0x27, 0x90, 0x31, 0x3c, 0x67, 0xf3, 0xb0, 0xbd, 0x1d, 0x83, 0x52, 0xa0, 0x49, 0x1e,
	0x92, 0xf2, 0x2a, 0x25, 0xed, 0x13, 0x21, 0xdd, 0x4d, 0x16, 0xe1, 0x18, 0x14, 0x11, 0x90, 0x50,
	0x2d, 0x48, 0x41, 0x44, 0xd9, 0x35, 0x31, 0x6e, 0x62, 0xb8, 0x25, 0x4e, 0x2a, 0x6e, 0x1d, 0x55,
	0xe6, 0xef, 0xaa, 0xc4, 0xbb, 0x02, 0x40, 0x24, 0x9c, 0x50, 0x23, 0x09, 0xd5, 0x1a, 0x1c, 0x75,
	0x12, 0x74, 0xd5, 0xa6, 0x9b, 0x78, 0xad, 0x0b, 0x69, 0xdd, 0x59, 0x51, 0x62, 0x90, 0x69, 0x96,
	0x70, 0x3b, 0xe2, 0xb7, 0x19, 0xb7, 0xce, 0xdf, 0xf7, 0x36, 0x21, 0x15, 0x54, 0xcb, 0xfb, 0x42,
	0x32, 0x96, 0xac, 0x85, 0x3a, 0xa8, 0xb7, 0x3e, 0x6a, 0xd6, 0xe1, 0x73, 0x36, 0xfb, 0x50, 0xea,
	0x38, 0xc9, 0x18, 0x1f, 0xf3, 0x3b, 0x23, 0x53, 0xce, 0x5a, 0x8d, 0x0e, 0xea, 0xad, 0x8d, 0x9a,
	0x15, 0x3c, 0x2c, 0xd1, 0xa3, 0x37, 0xe4, 0x35, 0xcb, 0x28, 0x7b, 0xc9, 0xd3, 0x5c, 0xc6, 0xdc,
	0x7f, 0x45, 0xde, 0xd6, 0x42, 0x03, 0xbf, 0x8f, 0x97, 0xae, 0x84, 0x97, 0xf5, 0x6d, 0xef, 0x2d,
	0x88, 0xca, 0x11, 0xf3, 0x10, 0x0f, 0x95, 0x71, 0x93, 0xee, 0xf1, 0xd3, 0xe7, 0xcf, 0x73, 0x23,
	0xf4, 0x49, 0x31, 0x46, 0x65, 0x39, 0x5b, 0xfb, 0xe1, 0xdf, 0x7f, 0x3e, 0x56, 0x64, 0x54, 0x06,
	0x0c, 0x4e, 0xdf, 0xa7, 0x01, 0xfa, 0x98, 0x06, 0xe8, 0x6b, 0x1a, 0xa0, 0x97, 0xef, 0x60, 0xe5,
	0xea, 0xa0, 0xf6, 0x24, 0xf3, 0xd4, 0x43, 0x45, 0x35, 0x15, 0x9c, 0xcd, 0xdc, 0x6d, 0xcd, 0x3e,
	0x5a, 0x2d, 0xb6, 0xee, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x04, 0x50, 0x96, 0x63, 0x0a, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CreditsServiceClient is the client API for CreditsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreditsServiceClient interface {
	// List credit bundles for an organization.
	// Required permissions:
	// - bundle.credits.list
	ListCreditBundles(ctx context.Context, in *ListCreditBundlesRequest, opts ...grpc.CallOption) (*v1.Empty, error)
}

type creditsServiceClient struct {
	cc *grpc.ClientConn
}

func NewCreditsServiceClient(cc *grpc.ClientConn) CreditsServiceClient {
	return &creditsServiceClient{cc}
}

func (c *creditsServiceClient) ListCreditBundles(ctx context.Context, in *ListCreditBundlesRequest, opts ...grpc.CallOption) (*v1.Empty, error) {
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.credits.v1.CreditsService/ListCreditBundles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreditsServiceServer is the server API for CreditsService service.
type CreditsServiceServer interface {
	// List credit bundles for an organization.
	// Required permissions:
	// - bundle.credits.list
	ListCreditBundles(context.Context, *ListCreditBundlesRequest) (*v1.Empty, error)
}

// UnimplementedCreditsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCreditsServiceServer struct {
}

func (*UnimplementedCreditsServiceServer) ListCreditBundles(ctx context.Context, req *ListCreditBundlesRequest) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCreditBundles not implemented")
}

func RegisterCreditsServiceServer(s *grpc.Server, srv CreditsServiceServer) {
	s.RegisterService(&_CreditsService_serviceDesc, srv)
}

func _CreditsService_ListCreditBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCreditBundlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditsServiceServer).ListCreditBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.credits.v1.CreditsService/ListCreditBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditsServiceServer).ListCreditBundles(ctx, req.(*ListCreditBundlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreditsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.credits.v1.CreditsService",
	HandlerType: (*CreditsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCreditBundles",
			Handler:    _CreditsService_ListCreditBundles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credits.proto",
}

func (m *ListCreditBundlesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListCreditBundlesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListCreditBundlesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IncludeExpired {
		i--
		if m.IncludeExpired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.OrganizationId) > 0 {
		i -= len(m.OrganizationId)
		copy(dAtA[i:], m.OrganizationId)
		i = encodeVarintCredits(dAtA, i, uint64(len(m.OrganizationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredits(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredits(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListCreditBundlesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrganizationId)
	if l > 0 {
		n += 1 + l + sovCredits(uint64(l))
	}
	if m.IncludeExpired {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCredits(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredits(x uint64) (n int) {
	return sovCredits(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListCreditBundlesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredits
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
			return fmt.Errorf("proto: ListCreditBundlesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListCreditBundlesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredits
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
				return ErrInvalidLengthCredits
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncludeExpired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredits
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
			m.IncludeExpired = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCredits(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredits
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
func skipCredits(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredits
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
					return 0, ErrIntOverflowCredits
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCredits
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
				return 0, ErrInvalidLengthCredits
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredits
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredits
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredits        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredits          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredits = fmt.Errorf("proto: unexpected end of group")
)
