// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mixer/v1/attributes.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Attributes represents a set of typed name/value pairs. Many of Mixer's
// API either consume and/or return attributes.
//
// Istio uses attributes to control the runtime behavior of services running in the service mesh.
// Attributes are named and typed pieces of metadata describing ingress and egress traffic and the
// environment this traffic occurs in. An Istio attribute carries a specific piece
// of information such as the error code of an API request, the latency of an API request, or the
// original IP address of a TCP connection. For example:
//
// ```yaml
// request.path: xyz/abc
// request.size: 234
// request.time: 12:34:56.789 04/17/2017
// source.ip: 192.168.0.1
// target.service: example
// ```
//
// A given Istio deployment has a fixed vocabulary of attributes that it understands.
// The specific vocabulary is determined by the set of attribute producers being used
// in the deployment. The primary attribute producer in Istio is Envoy, although
// specialized Mixer adapters and services can also generate attributes.
//
// The common baseline set of attributes available in most Istio deployments is defined
// [here](https://istio.io/docs/reference/config/policy-and-telemetry/attribute-vocabulary/).
//
// Attributes are strongly typed. The supported attribute types are defined by
// [ValueType](https://github.com/istio/api/blob/master/policy/v1beta1/value_type.proto).
// Each type of value is encoded into one of the so-called transport types present
// in this message.
//
// Defines a map of attributes in uncompressed format.
// Following places may use this message:
// 1) Configure Istio/Proxy with static per-proxy attributes, such as source.uid.
// 2) Service IDL definition to extract api attributes for active requests.
// 3) Forward attributes from client proxy to server proxy for HTTP requests.
type Attributes struct {
	// A map of attribute name to its value.
	Attributes           map[string]*Attributes_AttributeValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}
func (*Attributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6504964367320bd3, []int{0}
}

func (m *Attributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes.Unmarshal(m, b)
}
func (m *Attributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes.Marshal(b, m, deterministic)
}
func (m *Attributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes.Merge(m, src)
}
func (m *Attributes) XXX_Size() int {
	return xxx_messageInfo_Attributes.Size(m)
}
func (m *Attributes) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes proto.InternalMessageInfo

func (m *Attributes) GetAttributes() map[string]*Attributes_AttributeValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Specifies one attribute value with different type.
type Attributes_AttributeValue struct {
	// The attribute value.
	//
	// Types that are valid to be assigned to Value:
	//	*Attributes_AttributeValue_StringValue
	//	*Attributes_AttributeValue_Int64Value
	//	*Attributes_AttributeValue_DoubleValue
	//	*Attributes_AttributeValue_BoolValue
	//	*Attributes_AttributeValue_BytesValue
	//	*Attributes_AttributeValue_TimestampValue
	//	*Attributes_AttributeValue_DurationValue
	//	*Attributes_AttributeValue_StringMapValue
	Value                isAttributes_AttributeValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Attributes_AttributeValue) Reset()         { *m = Attributes_AttributeValue{} }
func (m *Attributes_AttributeValue) String() string { return proto.CompactTextString(m) }
func (*Attributes_AttributeValue) ProtoMessage()    {}
func (*Attributes_AttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6504964367320bd3, []int{0, 1}
}

func (m *Attributes_AttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes_AttributeValue.Unmarshal(m, b)
}
func (m *Attributes_AttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes_AttributeValue.Marshal(b, m, deterministic)
}
func (m *Attributes_AttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes_AttributeValue.Merge(m, src)
}
func (m *Attributes_AttributeValue) XXX_Size() int {
	return xxx_messageInfo_Attributes_AttributeValue.Size(m)
}
func (m *Attributes_AttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes_AttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes_AttributeValue proto.InternalMessageInfo

type isAttributes_AttributeValue_Value interface {
	isAttributes_AttributeValue_Value()
}

type Attributes_AttributeValue_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Attributes_AttributeValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Attributes_AttributeValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Attributes_AttributeValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Attributes_AttributeValue_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,6,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Attributes_AttributeValue_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Attributes_AttributeValue_DurationValue struct {
	DurationValue *duration.Duration `protobuf:"bytes,8,opt,name=duration_value,json=durationValue,proto3,oneof"`
}

type Attributes_AttributeValue_StringMapValue struct {
	StringMapValue *Attributes_StringMap `protobuf:"bytes,9,opt,name=string_map_value,json=stringMapValue,proto3,oneof"`
}

func (*Attributes_AttributeValue_StringValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_Int64Value) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_DoubleValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_BoolValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_BytesValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_TimestampValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_DurationValue) isAttributes_AttributeValue_Value() {}

func (*Attributes_AttributeValue_StringMapValue) isAttributes_AttributeValue_Value() {}

func (m *Attributes_AttributeValue) GetValue() isAttributes_AttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Attributes_AttributeValue) GetStringValue() string {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Attributes_AttributeValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Attributes_AttributeValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Attributes_AttributeValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Attributes_AttributeValue) GetBytesValue() []byte {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Attributes_AttributeValue) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Attributes_AttributeValue) GetDurationValue() *duration.Duration {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_DurationValue); ok {
		return x.DurationValue
	}
	return nil
}

func (m *Attributes_AttributeValue) GetStringMapValue() *Attributes_StringMap {
	if x, ok := m.GetValue().(*Attributes_AttributeValue_StringMapValue); ok {
		return x.StringMapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Attributes_AttributeValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Attributes_AttributeValue_StringValue)(nil),
		(*Attributes_AttributeValue_Int64Value)(nil),
		(*Attributes_AttributeValue_DoubleValue)(nil),
		(*Attributes_AttributeValue_BoolValue)(nil),
		(*Attributes_AttributeValue_BytesValue)(nil),
		(*Attributes_AttributeValue_TimestampValue)(nil),
		(*Attributes_AttributeValue_DurationValue)(nil),
		(*Attributes_AttributeValue_StringMapValue)(nil),
	}
}

// Defines a string map.
type Attributes_StringMap struct {
	// Holds a set of name/value pairs.
	Entries              map[string]string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Attributes_StringMap) Reset()         { *m = Attributes_StringMap{} }
func (m *Attributes_StringMap) String() string { return proto.CompactTextString(m) }
func (*Attributes_StringMap) ProtoMessage()    {}
func (*Attributes_StringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_6504964367320bd3, []int{0, 2}
}

func (m *Attributes_StringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes_StringMap.Unmarshal(m, b)
}
func (m *Attributes_StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes_StringMap.Marshal(b, m, deterministic)
}
func (m *Attributes_StringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes_StringMap.Merge(m, src)
}
func (m *Attributes_StringMap) XXX_Size() int {
	return xxx_messageInfo_Attributes_StringMap.Size(m)
}
func (m *Attributes_StringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes_StringMap.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes_StringMap proto.InternalMessageInfo

func (m *Attributes_StringMap) GetEntries() map[string]string {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Defines a list of attributes in compressed format optimized for transport.
// Within this message, strings are referenced using integer indices into
// one of two string dictionaries. Positive integers index into the global
// deployment-wide dictionary, whereas negative integers index into the message-level
// dictionary instead. The message-level dictionary is carried by the
// `words` field of this message, the deployment-wide dictionary is determined via
// configuration.
type CompressedAttributes struct {
	// The message-level dictionary.
	Words []string `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	// Holds attributes of type STRING, DNS_NAME, EMAIL_ADDRESS, URI
	Strings map[int32]int32 `protobuf:"bytes,2,rep,name=strings,proto3" json:"strings,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	// Holds attributes of type INT64
	Int64S map[int32]int64 `protobuf:"bytes,3,rep,name=int64s,proto3" json:"int64s,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Holds attributes of type DOUBLE
	Doubles map[int32]float64 `protobuf:"bytes,4,rep,name=doubles,proto3" json:"doubles,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// Holds attributes of type BOOL
	Bools map[int32]bool `protobuf:"bytes,5,rep,name=bools,proto3" json:"bools,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Holds attributes of type TIMESTAMP
	Timestamps map[int32]*timestamp.Timestamp `protobuf:"bytes,6,rep,name=timestamps,proto3" json:"timestamps,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Holds attributes of type DURATION
	Durations map[int32]*duration.Duration `protobuf:"bytes,7,rep,name=durations,proto3" json:"durations,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Holds attributes of type BYTES
	Bytes map[int32][]byte `protobuf:"bytes,8,rep,name=bytes,proto3" json:"bytes,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Holds attributes of type STRING_MAP
	StringMaps           map[int32]*StringMap `protobuf:"bytes,9,rep,name=string_maps,json=stringMaps,proto3" json:"string_maps,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CompressedAttributes) Reset()         { *m = CompressedAttributes{} }
func (m *CompressedAttributes) String() string { return proto.CompactTextString(m) }
func (*CompressedAttributes) ProtoMessage()    {}
func (*CompressedAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6504964367320bd3, []int{1}
}

func (m *CompressedAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompressedAttributes.Unmarshal(m, b)
}
func (m *CompressedAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompressedAttributes.Marshal(b, m, deterministic)
}
func (m *CompressedAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedAttributes.Merge(m, src)
}
func (m *CompressedAttributes) XXX_Size() int {
	return xxx_messageInfo_CompressedAttributes.Size(m)
}
func (m *CompressedAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedAttributes proto.InternalMessageInfo

func (m *CompressedAttributes) GetWords() []string {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *CompressedAttributes) GetStrings() map[int32]int32 {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *CompressedAttributes) GetInt64S() map[int32]int64 {
	if m != nil {
		return m.Int64S
	}
	return nil
}

func (m *CompressedAttributes) GetDoubles() map[int32]float64 {
	if m != nil {
		return m.Doubles
	}
	return nil
}

func (m *CompressedAttributes) GetBools() map[int32]bool {
	if m != nil {
		return m.Bools
	}
	return nil
}

func (m *CompressedAttributes) GetTimestamps() map[int32]*timestamp.Timestamp {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *CompressedAttributes) GetDurations() map[int32]*duration.Duration {
	if m != nil {
		return m.Durations
	}
	return nil
}

func (m *CompressedAttributes) GetBytes() map[int32][]byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *CompressedAttributes) GetStringMaps() map[int32]*StringMap {
	if m != nil {
		return m.StringMaps
	}
	return nil
}

// A map of string to string. The keys and values in this map are dictionary
// indices (see the [Attributes][istio.mixer.v1.CompressedAttributes] message for an explanation)
type StringMap struct {
	// Holds a set of name/value pairs.
	Entries              map[int32]int32 `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StringMap) Reset()         { *m = StringMap{} }
func (m *StringMap) String() string { return proto.CompactTextString(m) }
func (*StringMap) ProtoMessage()    {}
func (*StringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_6504964367320bd3, []int{2}
}

func (m *StringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMap.Unmarshal(m, b)
}
func (m *StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMap.Marshal(b, m, deterministic)
}
func (m *StringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMap.Merge(m, src)
}
func (m *StringMap) XXX_Size() int {
	return xxx_messageInfo_StringMap.Size(m)
}
func (m *StringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMap.DiscardUnknown(m)
}

var xxx_messageInfo_StringMap proto.InternalMessageInfo

func (m *StringMap) GetEntries() map[int32]int32 {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*Attributes)(nil), "istio.mixer.v1.Attributes")
	proto.RegisterMapType((map[string]*Attributes_AttributeValue)(nil), "istio.mixer.v1.Attributes.AttributesEntry")
	proto.RegisterType((*Attributes_AttributeValue)(nil), "istio.mixer.v1.Attributes.AttributeValue")
	proto.RegisterType((*Attributes_StringMap)(nil), "istio.mixer.v1.Attributes.StringMap")
	proto.RegisterMapType((map[string]string)(nil), "istio.mixer.v1.Attributes.StringMap.EntriesEntry")
	proto.RegisterType((*CompressedAttributes)(nil), "istio.mixer.v1.CompressedAttributes")
	proto.RegisterMapType((map[int32]bool)(nil), "istio.mixer.v1.CompressedAttributes.BoolsEntry")
	proto.RegisterMapType((map[int32][]byte)(nil), "istio.mixer.v1.CompressedAttributes.BytesEntry")
	proto.RegisterMapType((map[int32]float64)(nil), "istio.mixer.v1.CompressedAttributes.DoublesEntry")
	proto.RegisterMapType((map[int32]*duration.Duration)(nil), "istio.mixer.v1.CompressedAttributes.DurationsEntry")
	proto.RegisterMapType((map[int32]int64)(nil), "istio.mixer.v1.CompressedAttributes.Int64sEntry")
	proto.RegisterMapType((map[int32]*StringMap)(nil), "istio.mixer.v1.CompressedAttributes.StringMapsEntry")
	proto.RegisterMapType((map[int32]int32)(nil), "istio.mixer.v1.CompressedAttributes.StringsEntry")
	proto.RegisterMapType((map[int32]*timestamp.Timestamp)(nil), "istio.mixer.v1.CompressedAttributes.TimestampsEntry")
	proto.RegisterType((*StringMap)(nil), "istio.mixer.v1.StringMap")
	proto.RegisterMapType((map[int32]int32)(nil), "istio.mixer.v1.StringMap.EntriesEntry")
}

func init() { proto.RegisterFile("mixer/v1/attributes.proto", fileDescriptor_6504964367320bd3) }

var fileDescriptor_6504964367320bd3 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xcd, 0x6e, 0xd3, 0x4a,
	0x18, 0x86, 0x3d, 0x4d, 0xf3, 0xe3, 0x2f, 0x39, 0x49, 0x6b, 0xe5, 0x48, 0x6e, 0x16, 0x89, 0x4f,
	0xcf, 0xd1, 0x91, 0xe9, 0xc2, 0xee, 0x9f, 0x50, 0xe9, 0x06, 0x08, 0x8d, 0x14, 0x40, 0x48, 0xc8,
	0x20, 0xfe, 0x2a, 0x81, 0x1c, 0xc5, 0x04, 0x8b, 0x24, 0x13, 0x79, 0x26, 0x85, 0xac, 0x59, 0xb1,
	0x63, 0xc1, 0x82, 0x3b, 0x80, 0x4b, 0xe9, 0x25, 0xb0, 0x02, 0xb5, 0x7b, 0x24, 0x96, 0x2c, 0x91,
	0x67, 0xc6, 0xce, 0xc4, 0xa4, 0x8d, 0xb3, 0xb3, 0x67, 0xde, 0xf7, 0x99, 0xcf, 0x33, 0xef, 0x7c,
	0x09, 0x6c, 0x0c, 0xfc, 0xb7, 0x5e, 0x60, 0x9f, 0xec, 0xd8, 0x2e, 0xa5, 0x81, 0xdf, 0x19, 0x53,
	0x8f, 0x58, 0xa3, 0x00, 0x53, 0xac, 0x95, 0x7d, 0x42, 0x7d, 0x6c, 0x31, 0x81, 0x75, 0xb2, 0x53,
	0xab, 0xf6, 0x70, 0x0f, 0xb3, 0x29, 0x3b, 0x7c, 0xe2, 0xaa, 0x5a, 0xbd, 0x87, 0x71, 0xaf, 0xef,
	0xd9, 0xec, 0xad, 0x33, 0x7e, 0x69, 0x77, 0xc7, 0x81, 0x4b, 0x7d, 0x3c, 0x14, 0xf3, 0x8d, 0xe4,
	0x3c, 0xf5, 0x07, 0x1e, 0xa1, 0xee, 0x60, 0xc4, 0x05, 0x9b, 0xef, 0x72, 0x00, 0x37, 0xe3, 0xb5,
	0xb5, 0x3b, 0x00, 0xd3, 0x4a, 0x74, 0x64, 0x64, 0xcc, 0xe2, 0xee, 0x96, 0x35, 0x5b, 0x8a, 0x35,
	0xd5, 0x4b, 0x8f, 0xad, 0x21, 0x0d, 0x26, 0x8e, 0xe4, 0xae, 0xbd, 0x82, 0x4a, 0x62, 0x5a, 0x5b,
	0x83, 0xcc, 0x6b, 0x6f, 0xa2, 0x23, 0x03, 0x99, 0xaa, 0x13, 0x3e, 0x6a, 0xd7, 0x21, 0x7b, 0xe2,
	0xf6, 0xc7, 0x9e, 0xbe, 0x62, 0x20, 0xb3, 0xb8, 0x7b, 0x25, 0xcd, 0x5a, 0x8f, 0x42, 0x83, 0xc3,
	0x7d, 0x87, 0x2b, 0x07, 0xa8, 0xf6, 0x39, 0x03, 0xe5, 0xd9, 0x59, 0xed, 0x5f, 0x28, 0x11, 0x1a,
	0xf8, 0xc3, 0xde, 0x8b, 0x29, 0x5e, 0x6d, 0x2b, 0x4e, 0x91, 0x8f, 0x72, 0xd1, 0x3f, 0x50, 0xf4,
	0x87, 0xf4, 0xea, 0xbe, 0xd0, 0x64, 0x0c, 0x64, 0x66, 0xda, 0x8a, 0x03, 0x6c, 0x30, 0xe6, 0x74,
	0xf1, 0xb8, 0xd3, 0xf7, 0x84, 0x66, 0xd5, 0x40, 0x26, 0x0a, 0x39, 0x7c, 0x94, 0x8b, 0x1a, 0x00,
	0x1d, 0x8c, 0xfb, 0x42, 0x92, 0x35, 0x90, 0x59, 0x68, 0x2b, 0x8e, 0x1a, 0x8e, 0xc5, 0x0b, 0x75,
	0x26, 0xd4, 0x23, 0x42, 0x91, 0x33, 0x90, 0x59, 0x0a, 0x17, 0x62, 0x83, 0x5c, 0xd2, 0x82, 0x4a,
	0x7c, 0x36, 0x42, 0x96, 0x67, 0x5b, 0x52, 0xb3, 0xf8, 0x19, 0x5a, 0xd1, 0x19, 0x5a, 0x0f, 0x23,
	0x5d, 0x5b, 0x71, 0xca, 0xb1, 0x89, 0x63, 0x9a, 0x50, 0x8e, 0x22, 0x20, 0x28, 0x05, 0x46, 0xd9,
	0xf8, 0x83, 0x72, 0x24, 0x64, 0x6d, 0xc5, 0xf9, 0x2b, 0xb2, 0x70, 0xc6, 0x7d, 0x58, 0x13, 0x7b,
	0x37, 0x70, 0xa3, 0x5a, 0x54, 0x46, 0xf9, 0xef, 0x92, 0xe3, 0x79, 0xc0, 0x2c, 0xf7, 0x5c, 0x56,
	0x15, 0x89, 0x5e, 0x18, 0xb1, 0x99, 0x17, 0xa7, 0x5c, 0xfb, 0x88, 0x40, 0x8d, 0x85, 0xda, 0x5d,
	0xc8, 0x7b, 0x43, 0x1a, 0xf8, 0x71, 0xd4, 0x76, 0xd2, 0xf0, 0xad, 0x16, 0xf7, 0xf0, 0xc4, 0x45,
	0x84, 0xda, 0x21, 0x94, 0xe4, 0x89, 0x39, 0x59, 0xab, 0xca, 0x59, 0x53, 0xa5, 0x00, 0x6d, 0xfe,
	0x50, 0xa1, 0x7a, 0x0b, 0x0f, 0x46, 0x81, 0x47, 0x88, 0xd7, 0x95, 0xee, 0x43, 0x15, 0xb2, 0x6f,
	0x70, 0xd0, 0xe5, 0xf5, 0xa9, 0x0e, 0x7f, 0x09, 0xeb, 0xe6, 0x1f, 0x48, 0xf4, 0x95, 0xf9, 0x75,
	0xcf, 0x83, 0x89, 0x2f, 0x88, 0xea, 0x16, 0x04, 0xad, 0x0d, 0x39, 0x96, 0x37, 0xa2, 0x67, 0x18,
	0x6b, 0x3b, 0x15, 0xeb, 0x36, 0xb3, 0x70, 0x94, 0xf0, 0x87, 0x65, 0xf1, 0x54, 0x12, 0x7d, 0x75,
	0x89, 0xb2, 0x8e, 0xb8, 0x47, 0x94, 0x25, 0x08, 0x5a, 0x0b, 0xb2, 0x61, 0x7e, 0x89, 0x9e, 0x65,
	0x28, 0x3b, 0x15, 0xaa, 0x19, 0x3a, 0x38, 0x88, 0xbb, 0xb5, 0xe7, 0x00, 0x71, 0x42, 0x89, 0x9e,
	0x63, 0xac, 0xfd, 0x54, 0xac, 0x38, 0xe5, 0x1c, 0xd8, 0x2c, 0x9c, 0x7e, 0x6b, 0x28, 0x1f, 0xbe,
	0x37, 0x90, 0x23, 0x11, 0xb5, 0x63, 0x50, 0xa3, 0xf0, 0x12, 0x3d, 0xcf, 0xf0, 0x7b, 0xe9, 0xbe,
	0x3a, 0x72, 0x49, 0xf4, 0x4f, 0x21, 0x7d, 0xca, 0x63, 0x7b, 0x10, 0xde, 0x50, 0xbd, 0xb0, 0xcc,
	0x1e, 0x4c, 0xe2, 0x6e, 0xc8, 0xdd, 0xda, 0x31, 0x14, 0xa7, 0xf7, 0x89, 0xe8, 0xea, 0x12, 0x9b,
	0x10, 0x87, 0x5e, 0x94, 0xb9, 0x1a, 0x96, 0xe9, 0x40, 0x7c, 0xbd, 0x58, 0xec, 0xe5, 0x5c, 0xc9,
	0xb1, 0x5f, 0x9f, 0x13, 0xfb, 0x75, 0xb9, 0x6f, 0x5e, 0x83, 0xa2, 0x94, 0xa3, 0x45, 0xd6, 0x8c,
	0x6c, 0x3d, 0x84, 0x92, 0x9c, 0x9b, 0x45, 0x5e, 0x24, 0x7b, 0x0f, 0x00, 0xa6, 0x41, 0x59, 0xe4,
	0x2c, 0xc8, 0xce, 0xa7, 0x50, 0x49, 0xc4, 0x62, 0x8e, 0x7d, 0x7b, 0xf6, 0x27, 0xe5, 0x92, 0xfe,
	0x29, 0xa3, 0x1f, 0x43, 0x79, 0x36, 0x12, 0x73, 0xc8, 0xf6, 0x2c, 0xf9, 0xe2, 0x9e, 0x9a, 0xfc,
	0xda, 0x09, 0x4d, 0xb9, 0x4f, 0x25, 0xd9, 0xf9, 0x04, 0x2a, 0x89, 0xf3, 0x4f, 0x53, 0x53, 0x22,
	0x56, 0x31, 0x41, 0xee, 0x77, 0xef, 0x67, 0xda, 0xf0, 0x8d, 0x64, 0x1b, 0xfe, 0xff, 0x42, 0xc8,
	0xf2, 0xbd, 0x77, 0x51, 0x08, 0x9b, 0x5b, 0xa7, 0x67, 0x75, 0xe5, 0xe7, 0x59, 0x5d, 0xf9, 0x72,
	0x5e, 0x57, 0xbe, 0x9e, 0xd7, 0xd1, 0xb3, 0xbf, 0xf9, 0xea, 0x3e, 0xb6, 0xdd, 0x91, 0x6f, 0x47,
	0x7f, 0x91, 0x7e, 0x21, 0xd4, 0xc9, 0xb1, 0x9d, 0xde, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x24, 0x3f, 0xc3, 0x38, 0x09, 0x00, 0x00,
}
