// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainaccounts/nameservice/cmp_host_result.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CmpHostResult struct {
	Index   string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Request string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *CmpHostResult) Reset()         { *m = CmpHostResult{} }
func (m *CmpHostResult) String() string { return proto.CompactTextString(m) }
func (*CmpHostResult) ProtoMessage()    {}
func (*CmpHostResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fcc3963923c8e3, []int{0}
}
func (m *CmpHostResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CmpHostResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CmpHostResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CmpHostResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmpHostResult.Merge(m, src)
}
func (m *CmpHostResult) XXX_Size() int {
	return m.Size()
}
func (m *CmpHostResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CmpHostResult.DiscardUnknown(m)
}

var xxx_messageInfo_CmpHostResult proto.InternalMessageInfo

func (m *CmpHostResult) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *CmpHostResult) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *CmpHostResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*CmpHostResult)(nil), "cosmos.interchainaccounts.nameservice.CmpHostResult")
}

func init() {
	proto.RegisterFile("interchainaccounts/nameservice/cmp_host_result.proto", fileDescriptor_53fcc3963923c8e3)
}

var fileDescriptor_53fcc3963923c8e3 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0xcc, 0x2b, 0x49,
	0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0x4b, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0xcf,
	0x4b, 0xcc, 0x4d, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xce, 0x2d, 0x88, 0xcf,
	0xc8, 0x2f, 0x2e, 0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x4d, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xc3, 0xd4, 0xac, 0x87, 0xa4, 0x59, 0x29,
	0x9c, 0x8b, 0xd7, 0x39, 0xb7, 0xc0, 0x23, 0xbf, 0xb8, 0x24, 0x08, 0xac, 0x5b, 0x48, 0x84, 0x8b,
	0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x92,
	0xe0, 0x62, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x60, 0x02, 0x8b, 0xc3, 0xb8, 0x42,
	0x62, 0x5c, 0x6c, 0x10, 0x7b, 0x25, 0x98, 0xc1, 0x12, 0x50, 0x9e, 0x53, 0xf0, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x43, 0x1c, 0xa9, 0x8f, 0x70, 0xa4, 0x2e, 0xdc, 0x8b, 0x15, 0x28, 0x9e, 0x2c,
	0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0x83, 0xa2, 0xfc, 0x13, 0x01, 0x00, 0x00,
}

func (m *CmpHostResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CmpHostResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CmpHostResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintCmpHostResult(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintCmpHostResult(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCmpHostResult(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCmpHostResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovCmpHostResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CmpHostResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCmpHostResult(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovCmpHostResult(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovCmpHostResult(uint64(l))
	}
	return n
}

func sovCmpHostResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCmpHostResult(x uint64) (n int) {
	return sovCmpHostResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CmpHostResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCmpHostResult
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
			return fmt.Errorf("proto: CmpHostResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CmpHostResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCmpHostResult
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
				return ErrInvalidLengthCmpHostResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCmpHostResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCmpHostResult
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
				return ErrInvalidLengthCmpHostResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCmpHostResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCmpHostResult
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
				return ErrInvalidLengthCmpHostResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCmpHostResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCmpHostResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCmpHostResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCmpHostResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCmpHostResult
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
					return 0, ErrIntOverflowCmpHostResult
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
					return 0, ErrIntOverflowCmpHostResult
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
				return 0, ErrInvalidLengthCmpHostResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCmpHostResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCmpHostResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCmpHostResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCmpHostResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCmpHostResult = fmt.Errorf("proto: unexpected end of group")
)
