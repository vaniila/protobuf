// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test_objects.proto

package jsonpb

import proto "github.com/vaniila/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/vaniila/protobuf/types"
import google_protobuf1 "github.com/vaniila/protobuf/types"
import google_protobuf2 "github.com/vaniila/protobuf/types"
import google_protobuf3 "github.com/vaniila/protobuf/types"
import google_protobuf4 "github.com/vaniila/protobuf/types"

// skipping weak import gogoproto "gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Widget_Color int32

const (
	Widget_RED   Widget_Color = 0
	Widget_GREEN Widget_Color = 1
	Widget_BLUE  Widget_Color = 2
)

var Widget_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Widget_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Widget_Color) Enum() *Widget_Color {
	p := new(Widget_Color)
	*p = x
	return p
}
func (x Widget_Color) String() string {
	return proto.EnumName(Widget_Color_name, int32(x))
}
func (x *Widget_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Widget_Color_value, data, "Widget_Color")
	if err != nil {
		return err
	}
	*x = Widget_Color(value)
	return nil
}
func (Widget_Color) EnumDescriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{3, 0} }

// Test message for holding primitive types.
type Simple struct {
	OBool            *bool    `protobuf:"varint,1,opt,name=o_bool,json=oBool" json:"o_bool,omitempty"`
	OInt32           *int32   `protobuf:"varint,2,opt,name=o_int32,json=oInt32" json:"o_int32,omitempty"`
	OInt64           *int64   `protobuf:"varint,3,opt,name=o_int64,json=oInt64" json:"o_int64,omitempty"`
	OUint32          *uint32  `protobuf:"varint,4,opt,name=o_uint32,json=oUint32" json:"o_uint32,omitempty"`
	OUint64          *uint64  `protobuf:"varint,5,opt,name=o_uint64,json=oUint64" json:"o_uint64,omitempty"`
	OSint32          *int32   `protobuf:"zigzag32,6,opt,name=o_sint32,json=oSint32" json:"o_sint32,omitempty"`
	OSint64          *int64   `protobuf:"zigzag64,7,opt,name=o_sint64,json=oSint64" json:"o_sint64,omitempty"`
	OFloat           *float32 `protobuf:"fixed32,8,opt,name=o_float,json=oFloat" json:"o_float,omitempty"`
	ODouble          *float64 `protobuf:"fixed64,9,opt,name=o_double,json=oDouble" json:"o_double,omitempty"`
	OString          *string  `protobuf:"bytes,10,opt,name=o_string,json=oString" json:"o_string,omitempty"`
	OBytes           []byte   `protobuf:"bytes,11,opt,name=o_bytes,json=oBytes" json:"o_bytes,omitempty"`
	OCastBytes       Bytes    `protobuf:"bytes,12,opt,name=o_cast_bytes,json=oCastBytes,casttype=Bytes" json:"o_cast_bytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Simple) Reset()                    { *m = Simple{} }
func (m *Simple) String() string            { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()               {}
func (*Simple) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{0} }

func (m *Simple) GetOBool() bool {
	if m != nil && m.OBool != nil {
		return *m.OBool
	}
	return false
}

func (m *Simple) GetOInt32() int32 {
	if m != nil && m.OInt32 != nil {
		return *m.OInt32
	}
	return 0
}

func (m *Simple) GetOInt64() int64 {
	if m != nil && m.OInt64 != nil {
		return *m.OInt64
	}
	return 0
}

func (m *Simple) GetOUint32() uint32 {
	if m != nil && m.OUint32 != nil {
		return *m.OUint32
	}
	return 0
}

func (m *Simple) GetOUint64() uint64 {
	if m != nil && m.OUint64 != nil {
		return *m.OUint64
	}
	return 0
}

func (m *Simple) GetOSint32() int32 {
	if m != nil && m.OSint32 != nil {
		return *m.OSint32
	}
	return 0
}

func (m *Simple) GetOSint64() int64 {
	if m != nil && m.OSint64 != nil {
		return *m.OSint64
	}
	return 0
}

func (m *Simple) GetOFloat() float32 {
	if m != nil && m.OFloat != nil {
		return *m.OFloat
	}
	return 0
}

func (m *Simple) GetODouble() float64 {
	if m != nil && m.ODouble != nil {
		return *m.ODouble
	}
	return 0
}

func (m *Simple) GetOString() string {
	if m != nil && m.OString != nil {
		return *m.OString
	}
	return ""
}

func (m *Simple) GetOBytes() []byte {
	if m != nil {
		return m.OBytes
	}
	return nil
}

func (m *Simple) GetOCastBytes() Bytes {
	if m != nil {
		return m.OCastBytes
	}
	return nil
}

// Test message for holding special non-finites primitives.
type NonFinites struct {
	FNan             *float32 `protobuf:"fixed32,1,opt,name=f_nan,json=fNan" json:"f_nan,omitempty"`
	FPinf            *float32 `protobuf:"fixed32,2,opt,name=f_pinf,json=fPinf" json:"f_pinf,omitempty"`
	FNinf            *float32 `protobuf:"fixed32,3,opt,name=f_ninf,json=fNinf" json:"f_ninf,omitempty"`
	DNan             *float64 `protobuf:"fixed64,4,opt,name=d_nan,json=dNan" json:"d_nan,omitempty"`
	DPinf            *float64 `protobuf:"fixed64,5,opt,name=d_pinf,json=dPinf" json:"d_pinf,omitempty"`
	DNinf            *float64 `protobuf:"fixed64,6,opt,name=d_ninf,json=dNinf" json:"d_ninf,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NonFinites) Reset()                    { *m = NonFinites{} }
func (m *NonFinites) String() string            { return proto.CompactTextString(m) }
func (*NonFinites) ProtoMessage()               {}
func (*NonFinites) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{1} }

func (m *NonFinites) GetFNan() float32 {
	if m != nil && m.FNan != nil {
		return *m.FNan
	}
	return 0
}

func (m *NonFinites) GetFPinf() float32 {
	if m != nil && m.FPinf != nil {
		return *m.FPinf
	}
	return 0
}

func (m *NonFinites) GetFNinf() float32 {
	if m != nil && m.FNinf != nil {
		return *m.FNinf
	}
	return 0
}

func (m *NonFinites) GetDNan() float64 {
	if m != nil && m.DNan != nil {
		return *m.DNan
	}
	return 0
}

func (m *NonFinites) GetDPinf() float64 {
	if m != nil && m.DPinf != nil {
		return *m.DPinf
	}
	return 0
}

func (m *NonFinites) GetDNinf() float64 {
	if m != nil && m.DNinf != nil {
		return *m.DNinf
	}
	return 0
}

// Test message for holding repeated primitives.
type Repeats struct {
	RBool            []bool    `protobuf:"varint,1,rep,name=r_bool,json=rBool" json:"r_bool,omitempty"`
	RInt32           []int32   `protobuf:"varint,2,rep,name=r_int32,json=rInt32" json:"r_int32,omitempty"`
	RInt64           []int64   `protobuf:"varint,3,rep,name=r_int64,json=rInt64" json:"r_int64,omitempty"`
	RUint32          []uint32  `protobuf:"varint,4,rep,name=r_uint32,json=rUint32" json:"r_uint32,omitempty"`
	RUint64          []uint64  `protobuf:"varint,5,rep,name=r_uint64,json=rUint64" json:"r_uint64,omitempty"`
	RSint32          []int32   `protobuf:"zigzag32,6,rep,name=r_sint32,json=rSint32" json:"r_sint32,omitempty"`
	RSint64          []int64   `protobuf:"zigzag64,7,rep,name=r_sint64,json=rSint64" json:"r_sint64,omitempty"`
	RFloat           []float32 `protobuf:"fixed32,8,rep,name=r_float,json=rFloat" json:"r_float,omitempty"`
	RDouble          []float64 `protobuf:"fixed64,9,rep,name=r_double,json=rDouble" json:"r_double,omitempty"`
	RString          []string  `protobuf:"bytes,10,rep,name=r_string,json=rString" json:"r_string,omitempty"`
	RBytes           [][]byte  `protobuf:"bytes,11,rep,name=r_bytes,json=rBytes" json:"r_bytes,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Repeats) Reset()                    { *m = Repeats{} }
func (m *Repeats) String() string            { return proto.CompactTextString(m) }
func (*Repeats) ProtoMessage()               {}
func (*Repeats) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{2} }

func (m *Repeats) GetRBool() []bool {
	if m != nil {
		return m.RBool
	}
	return nil
}

func (m *Repeats) GetRInt32() []int32 {
	if m != nil {
		return m.RInt32
	}
	return nil
}

func (m *Repeats) GetRInt64() []int64 {
	if m != nil {
		return m.RInt64
	}
	return nil
}

func (m *Repeats) GetRUint32() []uint32 {
	if m != nil {
		return m.RUint32
	}
	return nil
}

func (m *Repeats) GetRUint64() []uint64 {
	if m != nil {
		return m.RUint64
	}
	return nil
}

func (m *Repeats) GetRSint32() []int32 {
	if m != nil {
		return m.RSint32
	}
	return nil
}

func (m *Repeats) GetRSint64() []int64 {
	if m != nil {
		return m.RSint64
	}
	return nil
}

func (m *Repeats) GetRFloat() []float32 {
	if m != nil {
		return m.RFloat
	}
	return nil
}

func (m *Repeats) GetRDouble() []float64 {
	if m != nil {
		return m.RDouble
	}
	return nil
}

func (m *Repeats) GetRString() []string {
	if m != nil {
		return m.RString
	}
	return nil
}

func (m *Repeats) GetRBytes() [][]byte {
	if m != nil {
		return m.RBytes
	}
	return nil
}

// Test message for holding enums and nested messages.
type Widget struct {
	Color            *Widget_Color  `protobuf:"varint,1,opt,name=color,enum=jsonpb.Widget_Color" json:"color,omitempty"`
	RColor           []Widget_Color `protobuf:"varint,2,rep,name=r_color,json=rColor,enum=jsonpb.Widget_Color" json:"r_color,omitempty"`
	Simple           *Simple        `protobuf:"bytes,10,opt,name=simple" json:"simple,omitempty"`
	RSimple          []*Simple      `protobuf:"bytes,11,rep,name=r_simple,json=rSimple" json:"r_simple,omitempty"`
	Repeats          *Repeats       `protobuf:"bytes,20,opt,name=repeats" json:"repeats,omitempty"`
	RRepeats         []*Repeats     `protobuf:"bytes,21,rep,name=r_repeats,json=rRepeats" json:"r_repeats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Widget) Reset()                    { *m = Widget{} }
func (m *Widget) String() string            { return proto.CompactTextString(m) }
func (*Widget) ProtoMessage()               {}
func (*Widget) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{3} }

func (m *Widget) GetColor() Widget_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Widget_RED
}

func (m *Widget) GetRColor() []Widget_Color {
	if m != nil {
		return m.RColor
	}
	return nil
}

func (m *Widget) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Widget) GetRSimple() []*Simple {
	if m != nil {
		return m.RSimple
	}
	return nil
}

func (m *Widget) GetRepeats() *Repeats {
	if m != nil {
		return m.Repeats
	}
	return nil
}

func (m *Widget) GetRRepeats() []*Repeats {
	if m != nil {
		return m.RRepeats
	}
	return nil
}

type Maps struct {
	MInt64Str        map[int64]string `protobuf:"bytes,1,rep,name=m_int64_str,json=mInt64Str" json:"m_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MBoolSimple      map[bool]*Simple `protobuf:"bytes,2,rep,name=m_bool_simple,json=mBoolSimple" json:"m_bool_simple,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Maps) Reset()                    { *m = Maps{} }
func (m *Maps) String() string            { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()               {}
func (*Maps) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{4} }

func (m *Maps) GetMInt64Str() map[int64]string {
	if m != nil {
		return m.MInt64Str
	}
	return nil
}

func (m *Maps) GetMBoolSimple() map[bool]*Simple {
	if m != nil {
		return m.MBoolSimple
	}
	return nil
}

type MsgWithOneof struct {
	// Types that are valid to be assigned to Union:
	//	*MsgWithOneof_Title
	//	*MsgWithOneof_Salary
	//	*MsgWithOneof_Country
	//	*MsgWithOneof_HomeAddress
	Union            isMsgWithOneof_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MsgWithOneof) Reset()                    { *m = MsgWithOneof{} }
func (m *MsgWithOneof) String() string            { return proto.CompactTextString(m) }
func (*MsgWithOneof) ProtoMessage()               {}
func (*MsgWithOneof) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{5} }

type isMsgWithOneof_Union interface {
	isMsgWithOneof_Union()
}

type MsgWithOneof_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title,oneof"`
}
type MsgWithOneof_Salary struct {
	Salary int64 `protobuf:"varint,2,opt,name=salary,oneof"`
}
type MsgWithOneof_Country struct {
	Country string `protobuf:"bytes,3,opt,name=Country,oneof"`
}
type MsgWithOneof_HomeAddress struct {
	HomeAddress string `protobuf:"bytes,4,opt,name=home_address,json=homeAddress,oneof"`
}

func (*MsgWithOneof_Title) isMsgWithOneof_Union()       {}
func (*MsgWithOneof_Salary) isMsgWithOneof_Union()      {}
func (*MsgWithOneof_Country) isMsgWithOneof_Union()     {}
func (*MsgWithOneof_HomeAddress) isMsgWithOneof_Union() {}

func (m *MsgWithOneof) GetUnion() isMsgWithOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *MsgWithOneof) GetTitle() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Title); ok {
		return x.Title
	}
	return ""
}

func (m *MsgWithOneof) GetSalary() int64 {
	if x, ok := m.GetUnion().(*MsgWithOneof_Salary); ok {
		return x.Salary
	}
	return 0
}

func (m *MsgWithOneof) GetCountry() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Country); ok {
		return x.Country
	}
	return ""
}

func (m *MsgWithOneof) GetHomeAddress() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_HomeAddress); ok {
		return x.HomeAddress
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MsgWithOneof) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MsgWithOneof_OneofMarshaler, _MsgWithOneof_OneofUnmarshaler, _MsgWithOneof_OneofSizer, []interface{}{
		(*MsgWithOneof_Title)(nil),
		(*MsgWithOneof_Salary)(nil),
		(*MsgWithOneof_Country)(nil),
		(*MsgWithOneof_HomeAddress)(nil),
	}
}

func _MsgWithOneof_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Title)
	case *MsgWithOneof_Salary:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Country)
	case *MsgWithOneof_HomeAddress:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.HomeAddress)
	case nil:
	default:
		return fmt.Errorf("MsgWithOneof.Union has unexpected type %T", x)
	}
	return nil
}

func _MsgWithOneof_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MsgWithOneof)
	switch tag {
	case 1: // union.title
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Title{x}
		return true, err
	case 2: // union.salary
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &MsgWithOneof_Salary{int64(x)}
		return true, err
	case 3: // union.Country
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Country{x}
		return true, err
	case 4: // union.home_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_HomeAddress{x}
		return true, err
	default:
		return false, nil
	}
}

func _MsgWithOneof_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Title)))
		n += len(x.Title)
	case *MsgWithOneof_Salary:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Country)))
		n += len(x.Country)
	case *MsgWithOneof_HomeAddress:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.HomeAddress)))
		n += len(x.HomeAddress)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Real struct {
	Value                        *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Real) Reset()                    { *m = Real{} }
func (m *Real) String() string            { return proto.CompactTextString(m) }
func (*Real) ProtoMessage()               {}
func (*Real) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{6} }

var extRange_Real = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Real) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Real
}

func (m *Real) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Complex struct {
	Imaginary                    *float64 `protobuf:"fixed64,1,opt,name=imaginary" json:"imaginary,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Complex) Reset()                    { *m = Complex{} }
func (m *Complex) String() string            { return proto.CompactTextString(m) }
func (*Complex) ProtoMessage()               {}
func (*Complex) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{7} }

var extRange_Complex = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Complex) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Complex
}

func (m *Complex) GetImaginary() float64 {
	if m != nil && m.Imaginary != nil {
		return *m.Imaginary
	}
	return 0
}

var E_Complex_RealExtension = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*Complex)(nil),
	Field:         123,
	Name:          "jsonpb.Complex.real_extension",
	Tag:           "bytes,123,opt,name=real_extension,json=realExtension",
	Filename:      "test_objects.proto",
}

type KnownTypes struct {
	An               *google_protobuf.Any          `protobuf:"bytes,14,opt,name=an" json:"an,omitempty"`
	Dur              *google_protobuf1.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	St               *google_protobuf2.Struct      `protobuf:"bytes,12,opt,name=st" json:"st,omitempty"`
	Ts               *google_protobuf3.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Lv               *google_protobuf2.ListValue   `protobuf:"bytes,15,opt,name=lv" json:"lv,omitempty"`
	Val              *google_protobuf2.Value       `protobuf:"bytes,16,opt,name=val" json:"val,omitempty"`
	Dbl              *google_protobuf4.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt              *google_protobuf4.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64              *google_protobuf4.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64              *google_protobuf4.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32              *google_protobuf4.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32              *google_protobuf4.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool             *google_protobuf4.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str              *google_protobuf4.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes            *google_protobuf4.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *KnownTypes) Reset()                    { *m = KnownTypes{} }
func (m *KnownTypes) String() string            { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()               {}
func (*KnownTypes) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{8} }

func (m *KnownTypes) GetAn() *google_protobuf.Any {
	if m != nil {
		return m.An
	}
	return nil
}

func (m *KnownTypes) GetDur() *google_protobuf1.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *KnownTypes) GetSt() *google_protobuf2.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *KnownTypes) GetTs() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetLv() *google_protobuf2.ListValue {
	if m != nil {
		return m.Lv
	}
	return nil
}

func (m *KnownTypes) GetVal() *google_protobuf2.Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *KnownTypes) GetDbl() *google_protobuf4.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *google_protobuf4.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *google_protobuf4.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *google_protobuf4.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *google_protobuf4.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *google_protobuf4.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *google_protobuf4.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *google_protobuf4.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *google_protobuf4.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*string)(nil),
	Field:         124,
	Name:          "jsonpb.name",
	Tag:           "bytes,124,opt,name=name",
	Filename:      "test_objects.proto",
}

func init() {
	proto.RegisterType((*Simple)(nil), "jsonpb.Simple")
	proto.RegisterType((*NonFinites)(nil), "jsonpb.NonFinites")
	proto.RegisterType((*Repeats)(nil), "jsonpb.Repeats")
	proto.RegisterType((*Widget)(nil), "jsonpb.Widget")
	proto.RegisterType((*Maps)(nil), "jsonpb.Maps")
	proto.RegisterType((*MsgWithOneof)(nil), "jsonpb.MsgWithOneof")
	proto.RegisterType((*Real)(nil), "jsonpb.Real")
	proto.RegisterType((*Complex)(nil), "jsonpb.Complex")
	proto.RegisterType((*KnownTypes)(nil), "jsonpb.KnownTypes")
	proto.RegisterEnum("jsonpb.Widget_Color", Widget_Color_name, Widget_Color_value)
	proto.RegisterExtension(E_Complex_RealExtension)
	proto.RegisterExtension(E_Name)
}

func init() { proto.RegisterFile("test_objects.proto", fileDescriptorTestObjects) }

var fileDescriptorTestObjects = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xcf, 0x72, 0x13, 0xc7,
	0x13, 0x66, 0x77, 0xb5, 0xfa, 0xd3, 0x32, 0x46, 0xbf, 0xc1, 0xc0, 0xa2, 0x1f, 0x09, 0x2a, 0x85,
	0x10, 0x05, 0x82, 0xa8, 0xc8, 0x2a, 0x55, 0x8a, 0xe4, 0x82, 0xb1, 0x09, 0xa9, 0x80, 0x93, 0x1a,
	0x43, 0xc8, 0x4d, 0xb5, 0xf2, 0xae, 0xc4, 0x92, 0xd5, 0x8c, 0x6a, 0x66, 0xd6, 0xa0, 0x4a, 0x0e,
	0x3e, 0xe7, 0x98, 0xca, 0x33, 0xe4, 0x11, 0x72, 0xc8, 0x63, 0xe4, 0x01, 0xf2, 0x20, 0x39, 0xa5,
	0xba, 0x67, 0x57, 0x6b, 0x2c, 0x74, 0xf2, 0x76, 0xf7, 0xf7, 0x7d, 0x9e, 0x99, 0xaf, 0x67, 0x5a,
	0xc0, 0x4c, 0xac, 0xcd, 0x58, 0x4e, 0x5e, 0xc7, 0xc7, 0x46, 0xf7, 0x17, 0x4a, 0x1a, 0xc9, 0xaa,
	0xaf, 0xb5, 0x14, 0x8b, 0x49, 0xfb, 0xfa, 0x4c, 0xca, 0x59, 0x1a, 0xdf, 0xa7, 0xec, 0x24, 0x9b,
	0xde, 0x0f, 0xc5, 0xd2, 0x42, 0xda, 0x1f, 0x9e, 0x2f, 0x45, 0x99, 0x0a, 0x4d, 0x22, 0x45, 0x5e,
	0xbf, 0x71, 0xbe, 0xae, 0x8d, 0xca, 0x8e, 0x4d, 0x5e, 0xbd, 0x79, 0xbe, 0x6a, 0x92, 0x79, 0xac,
	0x4d, 0x38, 0x5f, 0x6c, 0x92, 0x7f, 0xa3, 0xc2, 0xc5, 0x22, 0x56, 0xf9, 0x0a, 0xdb, 0x3b, 0x33,
	0x39, 0x93, 0xf4, 0x79, 0x1f, 0xbf, 0x6c, 0xb6, 0xfb, 0xb7, 0x0b, 0xd5, 0xa3, 0x64, 0xbe, 0x48,
	0x63, 0x76, 0x05, 0xaa, 0x72, 0x3c, 0x91, 0x32, 0x0d, 0x9c, 0x8e, 0xd3, 0xab, 0x73, 0x5f, 0xee,
	0x49, 0x99, 0xb2, 0x6b, 0x50, 0x93, 0xe3, 0x44, 0x98, 0xdd, 0x41, 0xe0, 0x76, 0x9c, 0x9e, 0xcf,
	0xab, 0xf2, 0x1b, 0x8c, 0x56, 0x85, 0xd1, 0x30, 0xf0, 0x3a, 0x4e, 0xcf, 0xb3, 0x85, 0xd1, 0x90,
	0x5d, 0x87, 0xba, 0x1c, 0x67, 0x96, 0x52, 0xe9, 0x38, 0xbd, 0x8b, 0xbc, 0x26, 0x5f, 0x50, 0x58,
	0x96, 0x46, 0xc3, 0xc0, 0xef, 0x38, 0xbd, 0x4a, 0x5e, 0x2a, 0x58, 0xda, 0xb2, 0xaa, 0x1d, 0xa7,
	0xf7, 0x3f, 0x5e, 0x93, 0x47, 0x67, 0x58, 0xda, 0xb2, 0x6a, 0x1d, 0xa7, 0xc7, 0xf2, 0xd2, 0x68,
	0x68, 0x17, 0x31, 0x4d, 0x65, 0x68, 0x82, 0x7a, 0xc7, 0xe9, 0xb9, 0xbc, 0x2a, 0x1f, 0x63, 0x64,
	0x39, 0x91, 0xcc, 0x26, 0x69, 0x1c, 0x34, 0x3a, 0x4e, 0xcf, 0xe1, 0x35, 0xb9, 0x4f, 0x61, 0x2e,
	0x67, 0x54, 0x22, 0x66, 0x01, 0x74, 0x9c, 0x5e, 0x03, 0xe5, 0x28, 0xb4, 0x72, 0x93, 0xa5, 0x89,
	0x75, 0xd0, 0xec, 0x38, 0xbd, 0x2d, 0x5e, 0x95, 0x7b, 0x18, 0xb1, 0xbb, 0xb0, 0x25, 0xc7, 0xc7,
	0xa1, 0x36, 0x79, 0x75, 0x0b, 0xab, 0x7b, 0x8d, 0x7f, 0xff, 0xb9, 0xe9, 0x13, 0x80, 0x83, 0x7c,
	0x14, 0x6a, 0x43, 0xdf, 0xdd, 0xdf, 0x1c, 0x80, 0x43, 0x29, 0x1e, 0x27, 0x22, 0x41, 0xee, 0x65,
	0xf0, 0xa7, 0x63, 0x11, 0x0a, 0x3a, 0x57, 0x97, 0x57, 0xa6, 0x87, 0xa1, 0xc0, 0xd3, 0x9e, 0x8e,
	0x17, 0x89, 0x98, 0xd2, 0xa9, 0xba, 0xdc, 0x9f, 0x7e, 0x9f, 0x88, 0xa9, 0x4d, 0x0b, 0x4c, 0x7b,
	0x79, 0xfa, 0x10, 0xd3, 0x97, 0xc1, 0x8f, 0x48, 0xa2, 0x42, 0x5b, 0xa9, 0x44, 0xb9, 0x44, 0x64,
	0x25, 0x7c, 0xca, 0xfa, 0x51, 0x21, 0x11, 0x59, 0x89, 0x6a, 0x9e, 0x46, 0x89, 0xee, 0x1f, 0x2e,
	0xd4, 0x78, 0xbc, 0x88, 0x43, 0xa3, 0x11, 0xa2, 0x0a, 0xab, 0x3d, 0xb4, 0x5a, 0x15, 0x56, 0xab,
	0x95, 0xd5, 0x1e, 0x5a, 0xad, 0x56, 0x56, 0xab, 0x95, 0xd5, 0x1e, 0x5a, 0xad, 0x56, 0x56, 0xab,
	0xd2, 0x6a, 0x0f, 0xad, 0x56, 0xa5, 0xd5, 0xaa, 0xb4, 0xda, 0x43, 0xab, 0x55, 0x69, 0xb5, 0x2a,
	0xad, 0xf6, 0xd0, 0x6a, 0x75, 0x74, 0x86, 0xb5, 0xb2, 0xda, 0x43, 0xab, 0x55, 0x69, 0xb5, 0x5a,
	0x59, 0xed, 0xa1, 0xd5, 0x6a, 0x65, 0xb5, 0x2a, 0xad, 0xf6, 0xd0, 0x6a, 0x55, 0x5a, 0xad, 0x4a,
	0xab, 0x3d, 0xb4, 0x5a, 0x95, 0x56, 0xab, 0x95, 0xd5, 0x1e, 0x5a, 0xad, 0xac, 0x7b, 0x7f, 0xba,
	0x50, 0x7d, 0x99, 0x44, 0xb3, 0xd8, 0xb0, 0x3b, 0xe0, 0x1f, 0xcb, 0x54, 0x2a, 0x72, 0x6e, 0x7b,
	0xb0, 0xd3, 0xb7, 0xb7, 0xbc, 0x6f, 0xcb, 0xfd, 0x47, 0x58, 0xe3, 0x16, 0xc2, 0xee, 0xa1, 0x9e,
	0x45, 0xe3, 0xe1, 0x6d, 0x42, 0x57, 0x15, 0xfd, 0x65, 0xb7, 0xa1, 0xaa, 0xe9, 0xde, 0x51, 0x0b,
	0x36, 0x07, 0xdb, 0x05, 0xda, 0xde, 0x46, 0x9e, 0x57, 0xd9, 0xa7, 0xf6, 0x40, 0x08, 0x89, 0xeb,
	0x5c, 0x47, 0xe2, 0x01, 0xe5, 0xd0, 0x9a, 0xb2, 0x06, 0x07, 0x3b, 0xa4, 0x79, 0xa9, 0x40, 0xe6,
	0xbe, 0xf3, 0xa2, 0xce, 0x3e, 0x83, 0x86, 0x1a, 0x17, 0xe0, 0x2b, 0x24, 0xbb, 0x06, 0xae, 0xab,
	0xfc, 0xab, 0xfb, 0x31, 0xf8, 0x76, 0xd1, 0x35, 0xf0, 0xf8, 0xc1, 0x7e, 0xeb, 0x02, 0x6b, 0x80,
	0xff, 0x35, 0x3f, 0x38, 0x38, 0x6c, 0x39, 0xac, 0x0e, 0x95, 0xbd, 0xa7, 0x2f, 0x0e, 0x5a, 0x6e,
	0xf7, 0x77, 0x17, 0x2a, 0xcf, 0xc2, 0x85, 0x66, 0x5f, 0x42, 0x73, 0x6e, 0xdb, 0x05, 0xcf, 0x9e,
	0x7a, 0xac, 0x39, 0xf8, 0x7f, 0xa1, 0x8f, 0x90, 0xfe, 0x33, 0xea, 0x9f, 0x23, 0xa3, 0x0e, 0x84,
	0x51, 0x4b, 0xde, 0x98, 0x17, 0x31, 0x7b, 0x08, 0x17, 0xe7, 0xd4, 0x9b, 0xc5, 0xae, 0x5d, 0xa2,
	0x7f, 0xf0, 0x2e, 0x1d, 0xfb, 0xd5, 0x6e, 0xdb, 0x0a, 0x34, 0xe7, 0x65, 0xa6, 0xfd, 0x15, 0x6c,
	0xbf, 0xab, 0xcf, 0x5a, 0xe0, 0xfd, 0x14, 0x2f, 0xc9, 0x46, 0x8f, 0xe3, 0x27, 0xdb, 0x01, 0xff,
	0x24, 0x4c, 0xb3, 0x98, 0xae, 0x5f, 0x83, 0xdb, 0xe0, 0x81, 0xfb, 0x85, 0xd3, 0x3e, 0x84, 0xd6,
	0x79, 0xf9, 0xb3, 0xfc, 0xba, 0xe5, 0xdf, 0x3a, 0xcb, 0x5f, 0x37, 0xa5, 0xd4, 0xeb, 0xfe, 0xea,
	0xc0, 0xd6, 0x33, 0x3d, 0x7b, 0x99, 0x98, 0x57, 0xdf, 0x89, 0x58, 0x4e, 0xd9, 0x55, 0xf0, 0x4d,
	0x62, 0xd2, 0x98, 0xe4, 0x1a, 0x4f, 0x2e, 0x70, 0x1b, 0xb2, 0x00, 0xaa, 0x3a, 0x4c, 0x43, 0xb5,
	0x24, 0x4d, 0xef, 0xc9, 0x05, 0x9e, 0xc7, 0xac, 0x0d, 0xb5, 0x47, 0x32, 0xc3, 0x95, 0xd0, 0xb3,
	0x80, 0x9c, 0x22, 0xc1, 0x3e, 0x82, 0xad, 0x57, 0x72, 0x1e, 0x8f, 0xc3, 0x28, 0x52, 0xb1, 0xd6,
	0xf4, 0x42, 0x20, 0xa0, 0x89, 0xd9, 0x87, 0x36, 0xb9, 0x57, 0x03, 0x3f, 0x13, 0x89, 0x14, 0xdd,
	0xdb, 0x50, 0xe1, 0x71, 0x98, 0x96, 0xdb, 0x77, 0xec, 0x1b, 0x41, 0xc1, 0x9d, 0x7a, 0x3d, 0x6a,
	0x9d, 0x9e, 0x9e, 0x9e, 0xba, 0xdd, 0x37, 0xf8, 0x1f, 0x71, 0x27, 0x6f, 0xd9, 0x0d, 0x68, 0x24,
	0xf3, 0x70, 0x96, 0x08, 0x5c, 0x99, 0x85, 0x97, 0x89, 0x92, 0x32, 0xd8, 0x87, 0x6d, 0x15, 0x87,
	0xe9, 0x38, 0x7e, 0x6b, 0x62, 0xa1, 0x13, 0x29, 0xd8, 0x56, 0xd9, 0x52, 0x61, 0x1a, 0xfc, 0xfc,
	0x6e, 0x4f, 0xe6, 0xf2, 0xfc, 0x22, 0x92, 0x0e, 0x0a, 0x4e, 0xf7, 0x2f, 0x1f, 0xe0, 0x5b, 0x21,
	0xdf, 0x88, 0xe7, 0xcb, 0x45, 0xac, 0xd9, 0x2d, 0x70, 0x43, 0x11, 0x6c, 0x13, 0x75, 0xa7, 0x6f,
	0x47, 0x5c, 0xbf, 0x18, 0x71, 0xfd, 0x87, 0x62, 0xc9, 0xdd, 0x50, 0xb0, 0xbb, 0xe0, 0x45, 0x99,
	0xbd, 0xa5, 0xcd, 0xc1, 0xf5, 0x35, 0xd8, 0x7e, 0x3e, 0x68, 0x39, 0xa2, 0xd8, 0x27, 0xe0, 0x6a,
	0x43, 0x0f, 0x78, 0x73, 0x70, 0x6d, 0x0d, 0x7b, 0x44, 0x43, 0x97, 0xbb, 0x1a, 0x6f, 0xbf, 0x6b,
	0x74, 0xee, 0x6f, 0x7b, 0x0d, 0xf8, 0xbc, 0x98, 0xbf, 0xdc, 0x35, 0x1a, 0xb1, 0xe9, 0x49, 0x70,
	0x69, 0x03, 0xf6, 0x69, 0xa2, 0xcd, 0x0f, 0x78, 0xc2, 0xdc, 0x4d, 0x4f, 0x58, 0x0f, 0xbc, 0x93,
	0x30, 0x0d, 0x5a, 0x04, 0xbe, 0xba, 0x06, 0xb6, 0x40, 0x84, 0xb0, 0x3e, 0x78, 0xd1, 0x24, 0x25,
	0xcf, 0x9b, 0x83, 0x1b, 0xeb, 0xfb, 0xa2, 0x47, 0x2e, 0xc7, 0x47, 0x93, 0x94, 0xdd, 0x03, 0x6f,
	0x9a, 0x1a, 0x6a, 0x01, 0xbc, 0x70, 0xe7, 0xf1, 0xf4, 0x5c, 0xe6, 0xf0, 0x69, 0x6a, 0x10, 0x9e,
	0xe4, 0x83, 0xf8, 0x7d, 0x70, 0xba, 0x42, 0x39, 0x3c, 0x19, 0x0d, 0x71, 0x35, 0xd9, 0x68, 0x48,
	0x53, 0xe5, 0x7d, 0xab, 0x79, 0x71, 0x16, 0x9f, 0x8d, 0x86, 0x24, 0xbf, 0x3b, 0xa0, 0x89, 0xbd,
	0x41, 0x7e, 0x77, 0x50, 0xc8, 0xef, 0x0e, 0x48, 0x7e, 0x77, 0x40, 0x63, 0x7c, 0x93, 0xfc, 0x0a,
	0x9f, 0x11, 0xbe, 0x42, 0x23, 0xac, 0xb1, 0xe1, 0xd0, 0xf1, 0x0e, 0x5b, 0x38, 0xe1, 0x50, 0x1f,
	0x5f, 0x23, 0xd8, 0xa0, 0x6f, 0xc7, 0x42, 0xae, 0xaf, 0x8d, 0x62, 0x9f, 0x83, 0x5f, 0xfe, 0x12,
	0x78, 0xdf, 0x06, 0x68, 0x5c, 0x58, 0x82, 0x45, 0x3e, 0xe8, 0x40, 0x45, 0x84, 0xf3, 0xf8, 0x5c,
	0xe3, 0xff, 0x42, 0x2f, 0x0c, 0x55, 0x7e, 0xf4, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xa6,
	0x21, 0x68, 0x3e, 0x0a, 0x00, 0x00,
}
