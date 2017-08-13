// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

/*
	Package types is a generated protocol buffer package.

	It is generated from these files:
		permission.proto

	It has these top-level messages:
		Permission
*/
package types

import proto "github.com/vaniila/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

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

type Permission struct {
	Perm uint64 `protobuf:"fixed64,1,opt,name=perm,proto3" json:"perm,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptorPermission, []int{0} }

func (m *Permission) GetPerm() uint64 {
	if m != nil {
		return m.Perm
	}
	return 0
}

func (m *Permission) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Permission)(nil), "google.protobuf.Permission")
}
func (this *Permission) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Permission)
	if !ok {
		that2, ok := that.(Permission)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Perm != that1.Perm {
		if this.Perm < that1.Perm {
			return -1
		}
		return 1
	}
	if this.Name != that1.Name {
		if this.Name < that1.Name {
			return -1
		}
		return 1
	}
	return 0
}
func (this *Permission) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Permission)
	if !ok {
		that2, ok := that.(Permission)
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
	if this.Perm != that1.Perm {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *Permission) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&types.Permission{")
	s = append(s, "Perm: "+fmt.Sprintf("%#v", this.Perm)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPermission(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Permission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permission) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Perm != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Permission(dAtA, i, uint64(m.Perm))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeFixed64Permission(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Permission(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPermission(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPermission(r randyPermission, easy bool) *Permission {
	this := &Permission{}
	this.Perm = uint64(uint64(r.Uint32()))
	this.Name = string(randStringPermission(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPermission interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePermission(r randyPermission) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPermission(r randyPermission) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RunePermission(r)
	}
	return string(tmps)
}
func randUnrecognizedPermission(r randyPermission, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPermission(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPermission(dAtA []byte, r randyPermission, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePermission(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePermission(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Permission) Size() (n int) {
	var l int
	_ = l
	if m.Perm != 0 {
		n += 9
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	return n
}

func sovPermission(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPermission(x uint64) (n int) {
	return sovPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Permission) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Permission{`,
		`Perm:` + fmt.Sprintf("%v", this.Perm) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPermission(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Permission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: Permission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			m.Perm = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Perm = uint64(dAtA[iNdEx-8])
			m.Perm |= uint64(dAtA[iNdEx-7]) << 8
			m.Perm |= uint64(dAtA[iNdEx-6]) << 16
			m.Perm |= uint64(dAtA[iNdEx-5]) << 24
			m.Perm |= uint64(dAtA[iNdEx-4]) << 32
			m.Perm |= uint64(dAtA[iNdEx-3]) << 40
			m.Perm |= uint64(dAtA[iNdEx-2]) << 48
			m.Perm |= uint64(dAtA[iNdEx-1]) << 56
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
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
func skipPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
				return 0, ErrInvalidLengthPermission
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPermission
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
				next, err := skipPermission(dAtA[start:])
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
	ErrInvalidLengthPermission = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermission   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("permission.proto", fileDescriptorPermission) }

var fileDescriptorPermission = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0x85, 0xf0, 0x92, 0x4a, 0xd3, 0x94, 0x4c, 0xb8, 0xb8, 0x02, 0xe0, 0x8a,
	0x84, 0x84, 0xb8, 0x58, 0x40, 0x5a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xd8, 0x82, 0xc0, 0x6c, 0x90,
	0x58, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xed, 0xd4, 0xc6,
	0x78, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28,
	0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x20, 0xf1, 0xc7, 0x72,
	0x8c, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0x7a, 0x68, 0xb6, 0x3b, 0xf1, 0x23, 0xec, 0x0e, 0x00, 0x89,
	0x05, 0x30, 0x46, 0xb1, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x2f, 0x62, 0x62, 0x76, 0x0f, 0x70, 0x5a,
	0xc5, 0x24, 0xe7, 0x0e, 0xd1, 0x11, 0x00, 0xd5, 0xa1, 0x17, 0x9e, 0x9a, 0x93, 0xe3, 0x9d, 0x97,
	0x5f, 0x9e, 0x17, 0x02, 0x52, 0x96, 0xc4, 0x06, 0x36, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x85, 0xfd, 0xc0, 0x38, 0xea, 0x00, 0x00, 0x00,
}
