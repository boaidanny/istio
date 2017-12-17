// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bazel-out/k8-fastbuild/genfiles/mixer/tools/codegen/pkg/interfacegen/testdata/quota_tmpl.proto

/*
	Package istio_mixer_adapter_quota is a generated protocol buffer package.

	It is generated from these files:
		bazel-out/k8-fastbuild/genfiles/mixer/tools/codegen/pkg/interfacegen/testdata/quota_tmpl.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package istio_mixer_adapter_quota

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"
import istio_mixer_v1_config_descriptor "istio.io/api/mixer/v1/config/descriptor"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

// template ...
type Type struct {
	// dimensions are ...
	Dimensions       map[string]istio_mixer_v1_config_descriptor.ValueType `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType"`
	AnotherValueType istio_mixer_v1_config_descriptor.ValueType            `protobuf:"varint,7,opt,name=anotherValueType,proto3,enum=istio.mixer.v1.config.descriptor.ValueType" json:"anotherValueType,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorQuotaTmpl, []int{0} }

func (m *Type) GetDimensions() map[string]istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Type) GetAnotherValueType() istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.AnotherValueType
	}
	return istio_mixer_v1_config_descriptor.VALUE_TYPE_UNSPECIFIED
}

type InstanceParam struct {
	Dimensions                     map[string]string `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int64Primitive                 string            `protobuf:"bytes,3,opt,name=int64Primitive,proto3" json:"int64Primitive,omitempty"`
	BoolPrimitive                  string            `protobuf:"bytes,4,opt,name=boolPrimitive,proto3" json:"boolPrimitive,omitempty"`
	DoublePrimitive                string            `protobuf:"bytes,5,opt,name=doublePrimitive,proto3" json:"doublePrimitive,omitempty"`
	StringPrimitive                string            `protobuf:"bytes,6,opt,name=stringPrimitive,proto3" json:"stringPrimitive,omitempty"`
	AnotherValueType               string            `protobuf:"bytes,7,opt,name=anotherValueType,proto3" json:"anotherValueType,omitempty"`
	DimensionsFixedInt64ValueDType map[string]string `protobuf:"bytes,8,rep,name=dimensionsFixedInt64ValueDType" json:"dimensionsFixedInt64ValueDType,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TimeStamp                      string            `protobuf:"bytes,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Duration                       string            `protobuf:"bytes,10,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *InstanceParam) Reset()                    { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage()               {}
func (*InstanceParam) Descriptor() ([]byte, []int) { return fileDescriptorQuotaTmpl, []int{1} }

func (m *InstanceParam) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *InstanceParam) GetInt64Primitive() string {
	if m != nil {
		return m.Int64Primitive
	}
	return ""
}

func (m *InstanceParam) GetBoolPrimitive() string {
	if m != nil {
		return m.BoolPrimitive
	}
	return ""
}

func (m *InstanceParam) GetDoublePrimitive() string {
	if m != nil {
		return m.DoublePrimitive
	}
	return ""
}

func (m *InstanceParam) GetStringPrimitive() string {
	if m != nil {
		return m.StringPrimitive
	}
	return ""
}

func (m *InstanceParam) GetAnotherValueType() string {
	if m != nil {
		return m.AnotherValueType
	}
	return ""
}

func (m *InstanceParam) GetDimensionsFixedInt64ValueDType() map[string]string {
	if m != nil {
		return m.DimensionsFixedInt64ValueDType
	}
	return nil
}

func (m *InstanceParam) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func (m *InstanceParam) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func init() {
	proto.RegisterType((*Type)(nil), "istio.mixer.adapter.quota.Type")
	proto.RegisterType((*InstanceParam)(nil), "istio.mixer.adapter.quota.InstanceParam")
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
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
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if this.AnotherValueType != that1.AnotherValueType {
		return false
	}
	return true
}
func (this *InstanceParam) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceParam)
	if !ok {
		that2, ok := that.(InstanceParam)
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
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if this.Int64Primitive != that1.Int64Primitive {
		return false
	}
	if this.BoolPrimitive != that1.BoolPrimitive {
		return false
	}
	if this.DoublePrimitive != that1.DoublePrimitive {
		return false
	}
	if this.StringPrimitive != that1.StringPrimitive {
		return false
	}
	if this.AnotherValueType != that1.AnotherValueType {
		return false
	}
	if len(this.DimensionsFixedInt64ValueDType) != len(that1.DimensionsFixedInt64ValueDType) {
		return false
	}
	for i := range this.DimensionsFixedInt64ValueDType {
		if this.DimensionsFixedInt64ValueDType[i] != that1.DimensionsFixedInt64ValueDType[i] {
			return false
		}
	}
	if this.TimeStamp != that1.TimeStamp {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&istio_mixer_adapter_quota.Type{")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "AnotherValueType: "+fmt.Sprintf("%#v", this.AnotherValueType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&istio_mixer_adapter_quota.InstanceParam{")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "Int64Primitive: "+fmt.Sprintf("%#v", this.Int64Primitive)+",\n")
	s = append(s, "BoolPrimitive: "+fmt.Sprintf("%#v", this.BoolPrimitive)+",\n")
	s = append(s, "DoublePrimitive: "+fmt.Sprintf("%#v", this.DoublePrimitive)+",\n")
	s = append(s, "StringPrimitive: "+fmt.Sprintf("%#v", this.StringPrimitive)+",\n")
	s = append(s, "AnotherValueType: "+fmt.Sprintf("%#v", this.AnotherValueType)+",\n")
	keysForDimensionsFixedInt64ValueDType := make([]string, 0, len(this.DimensionsFixedInt64ValueDType))
	for k, _ := range this.DimensionsFixedInt64ValueDType {
		keysForDimensionsFixedInt64ValueDType = append(keysForDimensionsFixedInt64ValueDType, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensionsFixedInt64ValueDType)
	mapStringForDimensionsFixedInt64ValueDType := "map[string]string{"
	for _, k := range keysForDimensionsFixedInt64ValueDType {
		mapStringForDimensionsFixedInt64ValueDType += fmt.Sprintf("%#v: %#v,", k, this.DimensionsFixedInt64ValueDType[k])
	}
	mapStringForDimensionsFixedInt64ValueDType += "}"
	if this.DimensionsFixedInt64ValueDType != nil {
		s = append(s, "DimensionsFixedInt64ValueDType: "+mapStringForDimensionsFixedInt64ValueDType+",\n")
	}
	s = append(s, "TimeStamp: "+fmt.Sprintf("%#v", this.TimeStamp)+",\n")
	s = append(s, "Duration: "+fmt.Sprintf("%#v", this.Duration)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQuotaTmpl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + sovQuotaTmpl(uint64(v))
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(v))
		}
	}
	if m.AnotherValueType != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(m.AnotherValueType))
	}
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + len(v) + sovQuotaTmpl(uint64(len(v)))
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Int64Primitive) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.Int64Primitive)))
		i += copy(dAtA[i:], m.Int64Primitive)
	}
	if len(m.BoolPrimitive) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.BoolPrimitive)))
		i += copy(dAtA[i:], m.BoolPrimitive)
	}
	if len(m.DoublePrimitive) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.DoublePrimitive)))
		i += copy(dAtA[i:], m.DoublePrimitive)
	}
	if len(m.StringPrimitive) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.StringPrimitive)))
		i += copy(dAtA[i:], m.StringPrimitive)
	}
	if len(m.AnotherValueType) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.AnotherValueType)))
		i += copy(dAtA[i:], m.AnotherValueType)
	}
	if len(m.DimensionsFixedInt64ValueDType) > 0 {
		for k, _ := range m.DimensionsFixedInt64ValueDType {
			dAtA[i] = 0x42
			i++
			v := m.DimensionsFixedInt64ValueDType[k]
			mapSize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + len(v) + sovQuotaTmpl(uint64(len(v)))
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.TimeStamp) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.TimeStamp)))
		i += copy(dAtA[i:], m.TimeStamp)
	}
	if len(m.Duration) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintQuotaTmpl(dAtA, i, uint64(len(m.Duration)))
		i += copy(dAtA[i:], m.Duration)
	}
	return i, nil
}

func encodeVarintQuotaTmpl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + sovQuotaTmpl(uint64(v))
			n += mapEntrySize + 1 + sovQuotaTmpl(uint64(mapEntrySize))
		}
	}
	if m.AnotherValueType != 0 {
		n += 1 + sovQuotaTmpl(uint64(m.AnotherValueType))
	}
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + len(v) + sovQuotaTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovQuotaTmpl(uint64(mapEntrySize))
		}
	}
	l = len(m.Int64Primitive)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	l = len(m.BoolPrimitive)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	l = len(m.DoublePrimitive)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	l = len(m.StringPrimitive)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	l = len(m.AnotherValueType)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	if len(m.DimensionsFixedInt64ValueDType) > 0 {
		for k, v := range m.DimensionsFixedInt64ValueDType {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuotaTmpl(uint64(len(k))) + 1 + len(v) + sovQuotaTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovQuotaTmpl(uint64(mapEntrySize))
		}
	}
	l = len(m.TimeStamp)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovQuotaTmpl(uint64(l))
	}
	return n
}

func sovQuotaTmpl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuotaTmpl(x uint64) (n int) {
	return sovQuotaTmpl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Type{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`AnotherValueType:` + fmt.Sprintf("%v", this.AnotherValueType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	keysForDimensionsFixedInt64ValueDType := make([]string, 0, len(this.DimensionsFixedInt64ValueDType))
	for k, _ := range this.DimensionsFixedInt64ValueDType {
		keysForDimensionsFixedInt64ValueDType = append(keysForDimensionsFixedInt64ValueDType, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensionsFixedInt64ValueDType)
	mapStringForDimensionsFixedInt64ValueDType := "map[string]string{"
	for _, k := range keysForDimensionsFixedInt64ValueDType {
		mapStringForDimensionsFixedInt64ValueDType += fmt.Sprintf("%v: %v,", k, this.DimensionsFixedInt64ValueDType[k])
	}
	mapStringForDimensionsFixedInt64ValueDType += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`Int64Primitive:` + fmt.Sprintf("%v", this.Int64Primitive) + `,`,
		`BoolPrimitive:` + fmt.Sprintf("%v", this.BoolPrimitive) + `,`,
		`DoublePrimitive:` + fmt.Sprintf("%v", this.DoublePrimitive) + `,`,
		`StringPrimitive:` + fmt.Sprintf("%v", this.StringPrimitive) + `,`,
		`AnotherValueType:` + fmt.Sprintf("%v", this.AnotherValueType) + `,`,
		`DimensionsFixedInt64ValueDType:` + mapStringForDimensionsFixedInt64ValueDType + `,`,
		`TimeStamp:` + fmt.Sprintf("%v", this.TimeStamp) + `,`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuotaTmpl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuotaTmpl
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
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType)
			}
			var mapkey string
			var mapvalue istio_mixer_v1_config_descriptor.ValueType
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuotaTmpl
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (istio_mixer_v1_config_descriptor.ValueType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuotaTmpl(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnotherValueType", wireType)
			}
			m.AnotherValueType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnotherValueType |= (istio_mixer_v1_config_descriptor.ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuotaTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuotaTmpl
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
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuotaTmpl
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
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuotaTmpl
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuotaTmpl(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Primitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Int64Primitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BoolPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoublePrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DoublePrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnotherValueType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AnotherValueType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DimensionsFixedInt64ValueDType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DimensionsFixedInt64ValueDType == nil {
				m.DimensionsFixedInt64ValueDType = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuotaTmpl
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuotaTmpl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuotaTmpl(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuotaTmpl
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DimensionsFixedInt64ValueDType[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeStamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaTmpl
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
				return ErrInvalidLengthQuotaTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuotaTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuotaTmpl
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
func skipQuotaTmpl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuotaTmpl
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
					return 0, ErrIntOverflowQuotaTmpl
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
					return 0, ErrIntOverflowQuotaTmpl
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
				return 0, ErrInvalidLengthQuotaTmpl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuotaTmpl
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
				next, err := skipQuotaTmpl(dAtA[start:])
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
	ErrInvalidLengthQuotaTmpl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuotaTmpl   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("bazel-out/k8-fastbuild/genfiles/mixer/tools/codegen/pkg/interfacegen/testdata/quota_tmpl.proto", fileDescriptorQuotaTmpl)
}

var fileDescriptorQuotaTmpl = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0xb6, 0x3f, 0xec, 0x8e, 0xf4, 0x07, 0x83, 0x87, 0x18, 0x64, 0x28, 0x55, 0x64,
	0xb1, 0x34, 0x43, 0xab, 0x48, 0x11, 0x3c, 0x28, 0x55, 0x28, 0x08, 0xd6, 0x2a, 0xea, 0xc9, 0x32,
	0xbb, 0x79, 0xbb, 0x8e, 0x4d, 0x66, 0xe2, 0xe4, 0x65, 0xe9, 0x7a, 0xf2, 0x4f, 0x10, 0xf4, 0x24,
	0xfe, 0x01, 0xfe, 0x29, 0x1e, 0x8b, 0x27, 0x8f, 0x36, 0x7a, 0xf0, 0xd8, 0xa3, 0x47, 0xc9, 0x44,
	0x36, 0xdd, 0xed, 0x0f, 0x5b, 0x6f, 0xc9, 0xcb, 0xe7, 0xfb, 0x7d, 0xef, 0x7d, 0x67, 0x42, 0x5f,
	0xb4, 0xe4, 0x1b, 0x88, 0x96, 0x4c, 0x86, 0x62, 0x7b, 0x75, 0xa9, 0x23, 0x53, 0x6c, 0x65, 0x2a,
	0x0a, 0x45, 0x17, 0x74, 0x47, 0x45, 0x90, 0x8a, 0x58, 0xed, 0x80, 0x15, 0x68, 0x4c, 0x94, 0x8a,
	0xb6, 0x09, 0xa1, 0x0b, 0x5a, 0x24, 0xdb, 0x5d, 0xa1, 0x34, 0x82, 0xed, 0xc8, 0xb6, 0x2b, 0x20,
	0xa4, 0x18, 0x4a, 0x94, 0xe2, 0x75, 0x66, 0x50, 0x6e, 0x61, 0x9c, 0x44, 0x41, 0x62, 0x0d, 0x1a,
	0x76, 0x51, 0xa5, 0xa8, 0x4c, 0xe0, 0x5c, 0x02, 0x19, 0xca, 0x04, 0xc1, 0x06, 0x8e, 0xf2, 0x17,
	0x4a, 0xeb, 0xde, 0xb2, 0x40, 0x88, 0x93, 0x48, 0x22, 0x08, 0xd8, 0x41, 0xd0, 0xa9, 0x32, 0x3a,
	0x2d, 0xe5, 0xfe, 0xe2, 0x80, 0x69, 0x1b, 0xdd, 0x51, 0x5d, 0x11, 0x42, 0xda, 0xb6, 0x2a, 0x41,
	0x63, 0x45, 0x4f, 0x46, 0x19, 0x6c, 0x61, 0x3f, 0x81, 0x12, 0x5e, 0xf8, 0x54, 0xa7, 0xe3, 0x4f,
	0xfa, 0x09, 0xb0, 0x87, 0x94, 0x86, 0x2a, 0xfe, 0xeb, 0xe4, 0x91, 0xf9, 0xb1, 0xe6, 0xf9, 0x15,
	0x11, 0x1c, 0x3b, 0x49, 0x50, 0x88, 0x82, 0xb5, 0x81, 0xe2, 0x9e, 0x46, 0xdb, 0xdf, 0x3c, 0x60,
	0xc1, 0x9e, 0xd1, 0x39, 0xa9, 0x0d, 0xbe, 0x04, 0xfb, 0xb4, 0x68, 0x5a, 0xf0, 0xde, 0xb9, 0x79,
	0xd2, 0x9c, 0x59, 0x59, 0x1c, 0xb2, 0xed, 0x2d, 0x07, 0xe5, 0x9c, 0x41, 0x35, 0x67, 0x30, 0x90,
	0x6c, 0x1e, 0x32, 0xf1, 0x5f, 0xd1, 0xd9, 0x91, 0xbe, 0x6c, 0x8e, 0x8e, 0x6d, 0x43, 0xdf, 0x23,
	0xf3, 0xa4, 0xd9, 0xd8, 0x2c, 0x1e, 0xd9, 0x1d, 0x3a, 0xe1, 0x76, 0xf5, 0xea, 0x67, 0x6f, 0x59,
	0x2a, 0x6f, 0xd5, 0x57, 0xc9, 0xc2, 0xc7, 0x09, 0x3a, 0xbd, 0xae, 0x53, 0x94, 0xba, 0x0d, 0x1b,
	0xd2, 0xca, 0x98, 0x3d, 0x3f, 0x22, 0xa7, 0xd5, 0x13, 0x72, 0x1a, 0x52, 0x9f, 0x18, 0xd8, 0x55,
	0x3a, 0xa3, 0x34, 0xde, 0xbc, 0xb1, 0x61, 0x55, 0xac, 0x50, 0xf5, 0xc0, 0x1b, 0x73, 0xfb, 0x8c,
	0x54, 0xd9, 0x15, 0x3a, 0xdd, 0x32, 0x26, 0xaa, 0xb0, 0x71, 0x87, 0x0d, 0x17, 0x59, 0x93, 0xce,
	0x86, 0x26, 0x6b, 0x45, 0x50, 0x71, 0x13, 0x8e, 0x1b, 0x2d, 0x17, 0x64, 0x8a, 0x56, 0xe9, 0x6e,
	0x45, 0x4e, 0x96, 0xe4, 0x48, 0x99, 0x5d, 0x3b, 0xe6, 0x48, 0x1b, 0x87, 0x4f, 0x89, 0x7d, 0x20,
	0x94, 0x57, 0xcb, 0xdd, 0x57, 0x3b, 0x10, 0xae, 0x17, 0x7b, 0x38, 0x62, 0xcd, 0x49, 0xa7, 0x5c,
	0x78, 0x0f, 0xfe, 0x23, 0xbc, 0xa3, 0xec, 0xca, 0x40, 0xff, 0xd1, 0x93, 0x5d, 0xa2, 0x0d, 0x54,
	0x31, 0x3c, 0x46, 0x19, 0x27, 0x5e, 0xc3, 0xcd, 0x5e, 0x15, 0x98, 0x4f, 0xa7, 0xc2, 0xcc, 0x4a,
	0x54, 0x46, 0x7b, 0xd4, 0x7d, 0x1c, 0xbc, 0xfb, 0xb7, 0x4f, 0x73, 0xed, 0x2e, 0x1c, 0xbc, 0x76,
	0x8d, 0x03, 0x37, 0xc9, 0x7f, 0x44, 0x2f, 0x9f, 0x62, 0xfe, 0xb3, 0x58, 0xde, 0x5d, 0xd9, 0xdd,
	0xe3, 0xb5, 0x6f, 0x7b, 0xbc, 0xb6, 0xbf, 0xc7, 0xc9, 0xdb, 0x9c, 0x93, 0xcf, 0x39, 0x27, 0x5f,
	0x72, 0x4e, 0x76, 0x73, 0x4e, 0xbe, 0xe7, 0x9c, 0xfc, 0xca, 0x79, 0x6d, 0x3f, 0xe7, 0xe4, 0xdd,
	0x0f, 0x5e, 0xfb, 0xfd, 0xf5, 0xe7, 0xfb, 0x7a, 0xbd, 0x35, 0xe9, 0x7e, 0xfb, 0xeb, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x7c, 0xa5, 0xdd, 0xc4, 0x04, 0x00, 0x00,
}
