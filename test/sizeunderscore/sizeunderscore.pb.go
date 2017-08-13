// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sizeunderscore.proto

/*
	Package sizeunderscore is a generated protocol buffer package.

	It is generated from these files:
		sizeunderscore.proto

	It has these top-level messages:
		SizeMessage
*/
package sizeunderscore

import proto "github.com/vaniila/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/vaniila/protobuf/gogoproto"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SizeMessage struct {
	Size_            *int64  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Equal_           *bool   `protobuf:"varint,2,opt,name=Equal" json:"Equal,omitempty"`
	String_          *string `protobuf:"bytes,3,opt,name=String" json:"String,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SizeMessage) Reset()                    { *m = SizeMessage{} }
func (m *SizeMessage) String() string            { return proto.CompactTextString(m) }
func (*SizeMessage) ProtoMessage()               {}
func (*SizeMessage) Descriptor() ([]byte, []int) { return fileDescriptorSizeunderscore, []int{0} }

func (m *SizeMessage) GetSize_() int64 {
	if m != nil && m.Size_ != nil {
		return *m.Size_
	}
	return 0
}

func (m *SizeMessage) GetEqual_() bool {
	if m != nil && m.Equal_ != nil {
		return *m.Equal_
	}
	return false
}

func (m *SizeMessage) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMessage)(nil), "sizeunderscore.SizeMessage")
}
func (this *SizeMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SizeMessage)
	if !ok {
		that2, ok := that.(SizeMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Size_ != nil && that1.Size_ != nil {
		if *this.Size_ != *that1.Size_ {
			return false
		}
	} else if this.Size_ != nil {
		return false
	} else if that1.Size_ != nil {
		return false
	}
	if this.Equal_ != nil && that1.Equal_ != nil {
		if *this.Equal_ != *that1.Equal_ {
			return false
		}
	} else if this.Equal_ != nil {
		return false
	} else if that1.Equal_ != nil {
		return false
	}
	if this.String_ != nil && that1.String_ != nil {
		if *this.String_ != *that1.String_ {
			return false
		}
	} else if this.String_ != nil {
		return false
	} else if that1.String_ != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *SizeMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SizeMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Size_ != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSizeunderscore(dAtA, i, uint64(*m.Size_))
	}
	if m.Equal_ != nil {
		dAtA[i] = 0x10
		i++
		if *m.Equal_ {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.String_ != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSizeunderscore(dAtA, i, uint64(len(*m.String_)))
		i += copy(dAtA[i:], *m.String_)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Sizeunderscore(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Sizeunderscore(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSizeunderscore(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSizeMessage(r randySizeunderscore, easy bool) *SizeMessage {
	this := &SizeMessage{}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Size_ = &v1
	}
	if r.Intn(10) != 0 {
		v2 := bool(bool(r.Intn(2) == 0))
		this.Equal_ = &v2
	}
	if r.Intn(10) != 0 {
		v3 := string(randStringSizeunderscore(r))
		this.String_ = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSizeunderscore(r, 4)
	}
	return this
}

type randySizeunderscore interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSizeunderscore(r randySizeunderscore) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSizeunderscore(r randySizeunderscore) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneSizeunderscore(r)
	}
	return string(tmps)
}
func randUnrecognizedSizeunderscore(r randySizeunderscore, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSizeunderscore(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSizeunderscore(dAtA []byte, r randySizeunderscore, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSizeunderscore(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SizeMessage) Size() (n int) {
	var l int
	_ = l
	if m.Size_ != nil {
		n += 1 + sovSizeunderscore(uint64(*m.Size_))
	}
	if m.Equal_ != nil {
		n += 2
	}
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovSizeunderscore(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSizeunderscore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSizeunderscore(x uint64) (n int) {
	return sovSizeunderscore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SizeMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSizeunderscore
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
			return fmt.Errorf("proto: SizeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SizeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Size_ = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Equal_", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
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
			b := bool(v != 0)
			m.Equal_ = &b
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
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
				return ErrInvalidLengthSizeunderscore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSizeunderscore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSizeunderscore
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
func skipSizeunderscore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSizeunderscore
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
					return 0, ErrIntOverflowSizeunderscore
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
					return 0, ErrIntOverflowSizeunderscore
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
				return 0, ErrInvalidLengthSizeunderscore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSizeunderscore
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
				next, err := skipSizeunderscore(dAtA[start:])
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
	ErrInvalidLengthSizeunderscore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSizeunderscore   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sizeunderscore.proto", fileDescriptorSizeunderscore) }

var fileDescriptorSizeunderscore = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xce, 0xac, 0x4a,
	0x2d, 0xcd, 0x4b, 0x49, 0x2d, 0x2a, 0x4e, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x95, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x2f, 0x4b, 0xcc, 0xcb, 0xcc, 0xcc, 0x49, 0xd4, 0x07, 0xab, 0x4c, 0x2a, 0x4d, 0xd3, 0x4f,
	0xcf, 0x4f, 0xcf, 0x07, 0x73, 0xc0, 0x2c, 0x88, 0x09, 0x4a, 0xfe, 0x5c, 0xdc, 0xc1, 0x99, 0x55,
	0xa9, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0x23, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0x56, 0xd7, 0xc2, 0xd2, 0xc4, 0x1c,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0xb8, 0xa4, 0x28,
	0x33, 0x2f, 0x5d, 0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x73, 0x92, 0xf8, 0xf1, 0x50,
	0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xfa, 0xab,
	0x04, 0xc3, 0x00, 0x00, 0x00,
}
