// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cockroachdb/errors/errorspb/tags.proto

package errorspb

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TagsPayload is the error payload for a WithContext
// marker.
// See errors/contexttags/withcontext.go and the RFC on
// error handling for details.
type TagsPayload struct {
	Tags                 []TagPayload `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TagsPayload) Reset()         { *m = TagsPayload{} }
func (m *TagsPayload) String() string { return proto.CompactTextString(m) }
func (*TagsPayload) ProtoMessage()    {}
func (*TagsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d76365188c23fe2, []int{0}
}
func (m *TagsPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TagsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TagsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagsPayload.Merge(m, src)
}
func (m *TagsPayload) XXX_Size() int {
	return m.Size()
}
func (m *TagsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TagsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TagsPayload proto.InternalMessageInfo

type TagPayload struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagPayload) Reset()         { *m = TagPayload{} }
func (m *TagPayload) String() string { return proto.CompactTextString(m) }
func (*TagPayload) ProtoMessage()    {}
func (*TagPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d76365188c23fe2, []int{1}
}
func (m *TagPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TagPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TagPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagPayload.Merge(m, src)
}
func (m *TagPayload) XXX_Size() int {
	return m.Size()
}
func (m *TagPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TagPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TagPayload proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TagsPayload)(nil), "cockroach.errorspb.TagsPayload")
	proto.RegisterType((*TagPayload)(nil), "cockroach.errorspb.TagPayload")
}

func init() {
	proto.RegisterFile("github.com/cockroachdb/errors/errorspb/tags.proto", fileDescriptor_5d76365188c23fe2)
}

var fileDescriptor_5d76365188c23fe2 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce,
	0x48, 0x49, 0xd2, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0x86, 0x52, 0x05, 0x49, 0xfa, 0x25, 0x89,
	0xe9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x70, 0x75, 0x7a, 0x30, 0x69, 0x29,
	0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb4, 0x3e, 0x88, 0x05, 0x51, 0xa9, 0xe4, 0xce, 0xc5, 0x1d,
	0x92, 0x98, 0x5e, 0x1c, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0x22, 0x64, 0xc1, 0xc5, 0x02, 0x32,
	0x46, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x4e, 0x0f, 0xd3, 0x1c, 0xbd, 0x90, 0xc4, 0x74,
	0xa8, 0x6a, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xc0, 0x3a, 0x94, 0x4c, 0xb8, 0xb8, 0x10,
	0x32, 0x42, 0x02, 0x5c, 0xcc, 0x25, 0x89, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20,
	0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0x71,
	0x52, 0x3a, 0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x6f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x28, 0x0e, 0x98, 0x85, 0x49,
	0x6c, 0x60, 0x97, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x3f, 0x04, 0x8e, 0x08, 0x01,
	0x00, 0x00,
}

func (m *TagsPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TagsPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, msg := range m.Tags {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTags(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TagPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TagPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tag) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTags(dAtA, i, uint64(len(m.Tag)))
		i += copy(dAtA[i:], m.Tag)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTags(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeVarintTags(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TagsPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, e := range m.Tags {
			l = e.Size()
			n += 1 + l + sovTags(uint64(l))
		}
	}
	return n
}

func (m *TagPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovTags(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTags(uint64(l))
	}
	return n
}

func sovTags(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTags(x uint64) (n int) {
	return sovTags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TagsPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTags
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
			return fmt.Errorf("proto: TagsPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagsPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTags
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
				return ErrInvalidLengthTags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, TagPayload{})
			if err := m.Tags[len(m.Tags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTags
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTags
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
func (m *TagPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTags
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
			return fmt.Errorf("proto: TagPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTags
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
				return ErrInvalidLengthTags
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTags
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
				return ErrInvalidLengthTags
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTags
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTags
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
func skipTags(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTags
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
					return 0, ErrIntOverflowTags
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
					return 0, ErrIntOverflowTags
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
				return 0, ErrInvalidLengthTags
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTags
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTags
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
				next, err := skipTags(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTags
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
	ErrInvalidLengthTags = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTags   = fmt.Errorf("proto: integer overflow")
)
