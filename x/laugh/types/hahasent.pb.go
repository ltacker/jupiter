// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: laugh/hahasent.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Hahasent struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *Hahasent) Reset()         { *m = Hahasent{} }
func (m *Hahasent) String() string { return proto.CompactTextString(m) }
func (*Hahasent) ProtoMessage()    {}
func (*Hahasent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9a0dfc6031f93c, []int{0}
}
func (m *Hahasent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hahasent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Hahasent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Hahasent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hahasent.Merge(m, src)
}
func (m *Hahasent) XXX_Size() int {
	return m.Size()
}
func (m *Hahasent) XXX_DiscardUnknown() {
	xxx_messageInfo_Hahasent.DiscardUnknown(m)
}

var xxx_messageInfo_Hahasent proto.InternalMessageInfo

func (m *Hahasent) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Hahasent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hahasent) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Hahasent)(nil), "ltacker.jupiter.laugh.Hahasent")
}

func init() { proto.RegisterFile("laugh/hahasent.proto", fileDescriptor_4c9a0dfc6031f93c) }

var fileDescriptor_4c9a0dfc6031f93c = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x2c, 0x4d,
	0xcf, 0xd0, 0xcf, 0x48, 0xcc, 0x48, 0x2c, 0x4e, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0xcd, 0x29, 0x49, 0x4c, 0xce, 0x4e, 0x2d, 0xd2, 0xcb, 0x2a, 0x2d, 0xc8, 0x2c, 0x49,
	0x2d, 0xd2, 0x03, 0xab, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20,
	0x8a, 0x95, 0x3c, 0xb8, 0x38, 0x3c, 0xa0, 0xda, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13,
	0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x3e, 0x2e, 0xa6,
	0xcc, 0x14, 0x09, 0x26, 0xb0, 0x20, 0x53, 0x66, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a, 0x45,
	0x89, 0x04, 0x33, 0x58, 0x04, 0xcc, 0x76, 0x72, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0x8d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xa8,
	0xdb, 0xf4, 0xa1, 0x6e, 0xd3, 0xaf, 0xd0, 0x87, 0xf8, 0xa1, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0xec, 0x28, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x3d, 0x74, 0x7a, 0xd9, 0x00,
	0x00, 0x00,
}

func (m *Hahasent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hahasent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hahasent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintHahasent(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintHahasent(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintHahasent(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHahasent(dAtA []byte, offset int, v uint64) int {
	offset -= sovHahasent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Hahasent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovHahasent(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovHahasent(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovHahasent(uint64(l))
	}
	return n
}

func sovHahasent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHahasent(x uint64) (n int) {
	return sovHahasent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Hahasent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHahasent
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
			return fmt.Errorf("proto: Hahasent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hahasent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHahasent
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
				return ErrInvalidLengthHahasent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHahasent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHahasent
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
				return ErrInvalidLengthHahasent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHahasent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHahasent
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
				return ErrInvalidLengthHahasent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHahasent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHahasent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHahasent
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
func skipHahasent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHahasent
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
					return 0, ErrIntOverflowHahasent
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
					return 0, ErrIntOverflowHahasent
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
				return 0, ErrInvalidLengthHahasent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHahasent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHahasent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHahasent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHahasent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHahasent = fmt.Errorf("proto: unexpected end of group")
)