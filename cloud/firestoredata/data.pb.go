// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.6
// source: cloud/firestore/v1/data.proto

package firestoredata

import (
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The data within all Firestore document events.
type DocumentEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A Document object containing a post-operation document snapshot.
	// This is not populated for delete events.
	Value *Document `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// A Document object containing a pre-operation document snapshot.
	// This is only populated for update and delete events.
	OldValue *Document `protobuf:"bytes,2,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	// A DocumentMask object that lists changed fields.
	// This is only populated for update events.
	UpdateMask    *DocumentMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DocumentEventData) Reset() {
	*x = DocumentEventData{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentEventData) ProtoMessage() {}

func (x *DocumentEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentEventData.ProtoReflect.Descriptor instead.
func (*DocumentEventData) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentEventData) GetValue() *Document {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *DocumentEventData) GetOldValue() *Document {
	if x != nil {
		return x.OldValue
	}
	return nil
}

func (x *DocumentEventData) GetUpdateMask() *DocumentMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// A set of field paths on a document.
type DocumentMask struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of field paths in the mask.
	// See [Document.fields][google.cloud.firestore.v1.events.Document.fields]
	// for a field path syntax reference.
	FieldPaths    []string `protobuf:"bytes,1,rep,name=field_paths,json=fieldPaths,proto3" json:"field_paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DocumentMask) Reset() {
	*x = DocumentMask{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentMask) ProtoMessage() {}

func (x *DocumentMask) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentMask.ProtoReflect.Descriptor instead.
func (*DocumentMask) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentMask) GetFieldPaths() []string {
	if x != nil {
		return x.FieldPaths
	}
	return nil
}

// A Firestore document.
type Document struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the document. For example:
	// `projects/{project_id}/databases/{database_id}/documents/{document_path}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The document's fields.
	//
	// The map keys represent field names.
	//
	// A simple field name contains only characters `a` to `z`, `A` to `Z`,
	// `0` to `9`, or `_`, and must not start with `0` to `9`. For example,
	// `foo_bar_17`.
	//
	// Field names matching the regular expression `__.*__` are reserved. Reserved
	// field names are forbidden except in certain documented contexts. The map
	// keys, represented as UTF-8, must not exceed 1,500 bytes and cannot be
	// empty.
	//
	// Field paths may be used in other contexts to refer to structured fields
	// defined here. For `map_value`, the field path is represented by the simple
	// or quoted field names of the containing fields, delimited by `.`. For
	// example, the structured field
	// `"foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would be
	// represented by the field path `foo.x&y`.
	//
	// Within a field path, a quoted field name starts and ends with “ ` “ and
	// may contain any character. Some characters, including “ ` “, must be
	// escaped using a `\`. For example, “ `x&y` “ represents `x&y` and
	// “ `bak\`tik` “ represents “ bak`tik “.
	Fields map[string]*Value `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The time at which the document was created.
	//
	// This value increases monotonically when a document is deleted then
	// recreated. It can also be compared to values from other documents and
	// the `read_time` of a query.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the document was last changed.
	//
	// This value is initially set to the `create_time` then increases
	// monotonically with each change to the document. It can also be
	// compared to values from other documents and the `read_time` of a query.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Document) Reset() {
	*x = Document{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *Document) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Document) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Document) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Document) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// A message that can hold any of the supported value types.
type Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_StringValue
	//	*Value_BytesValue
	//	*Value_ReferenceValue
	//	*Value_GeoPointValue
	//	*Value_ArrayValue
	//	*Value_MapValue
	ValueType     isValue_ValueType `protobuf_oneof:"value_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{3}
}

func (x *Value) GetValueType() isValue_ValueType {
	if x != nil {
		return x.ValueType
	}
	return nil
}

func (x *Value) GetNullValue() structpb.NullValue {
	if x != nil {
		if x, ok := x.ValueType.(*Value_NullValue); ok {
			return x.NullValue
		}
	}
	return structpb.NullValue(0)
}

func (x *Value) GetBooleanValue() bool {
	if x != nil {
		if x, ok := x.ValueType.(*Value_BooleanValue); ok {
			return x.BooleanValue
		}
	}
	return false
}

func (x *Value) GetIntegerValue() int64 {
	if x != nil {
		if x, ok := x.ValueType.(*Value_IntegerValue); ok {
			return x.IntegerValue
		}
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x != nil {
		if x, ok := x.ValueType.(*Value_DoubleValue); ok {
			return x.DoubleValue
		}
	}
	return 0
}

func (x *Value) GetTimestampValue() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.ValueType.(*Value_TimestampValue); ok {
			return x.TimestampValue
		}
	}
	return nil
}

func (x *Value) GetStringValue() string {
	if x != nil {
		if x, ok := x.ValueType.(*Value_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *Value) GetBytesValue() []byte {
	if x != nil {
		if x, ok := x.ValueType.(*Value_BytesValue); ok {
			return x.BytesValue
		}
	}
	return nil
}

func (x *Value) GetReferenceValue() string {
	if x != nil {
		if x, ok := x.ValueType.(*Value_ReferenceValue); ok {
			return x.ReferenceValue
		}
	}
	return ""
}

func (x *Value) GetGeoPointValue() *latlng.LatLng {
	if x != nil {
		if x, ok := x.ValueType.(*Value_GeoPointValue); ok {
			return x.GeoPointValue
		}
	}
	return nil
}

func (x *Value) GetArrayValue() *ArrayValue {
	if x != nil {
		if x, ok := x.ValueType.(*Value_ArrayValue); ok {
			return x.ArrayValue
		}
	}
	return nil
}

func (x *Value) GetMapValue() *MapValue {
	if x != nil {
		if x, ok := x.ValueType.(*Value_MapValue); ok {
			return x.MapValue
		}
	}
	return nil
}

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	// A null value.
	NullValue structpb.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	// A boolean value.
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	// An integer value.
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	// A double value.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	// A timestamp value.
	//
	// Precise only to microseconds. When stored, any additional precision is
	// rounded down.
	TimestampValue *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_StringValue struct {
	// A string value.
	//
	// The string, represented as UTF-8, must not exceed 1 MiB - 89 bytes.
	// Only the first 1,500 bytes of the UTF-8 representation are considered by
	// queries.
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BytesValue struct {
	// A bytes value.
	//
	// Must not exceed 1 MiB - 89 bytes.
	// Only the first 1,500 bytes are considered by queries.
	BytesValue []byte `protobuf:"bytes,18,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_ReferenceValue struct {
	// A reference to a document. For example:
	// `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	ReferenceValue string `protobuf:"bytes,5,opt,name=reference_value,json=referenceValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	// A geo point value representing a point on the surface of Earth.
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	// An array value.
	//
	// Cannot directly contain another array value, though can contain an
	// map which contains another array.
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type Value_MapValue struct {
	// A map value.
	MapValue *MapValue `protobuf:"bytes,6,opt,name=map_value,json=mapValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BytesValue) isValue_ValueType() {}

func (*Value_ReferenceValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

func (*Value_MapValue) isValue_ValueType() {}

// An array value.
type ArrayValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Values in the array.
	Values        []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArrayValue) Reset() {
	*x = ArrayValue{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArrayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayValue) ProtoMessage() {}

func (x *ArrayValue) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayValue.ProtoReflect.Descriptor instead.
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{4}
}

func (x *ArrayValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

// A map value.
type MapValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The map's fields.
	//
	// The map keys represent field names. Field names matching the regular
	// expression `__.*__` are reserved. Reserved field names are forbidden except
	// in certain documented contexts. The map keys, represented as UTF-8, must
	// not exceed 1,500 bytes and cannot be empty.
	Fields        map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapValue) Reset() {
	*x = MapValue{}
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValue) ProtoMessage() {}

func (x *MapValue) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_firestore_v1_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValue.ProtoReflect.Descriptor instead.
func (*MapValue) Descriptor() ([]byte, []int) {
	return file_cloud_firestore_v1_data_proto_rawDescGZIP(), []int{5}
}

func (x *MapValue) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_cloud_firestore_v1_data_proto protoreflect.FileDescriptor

const file_cloud_firestore_v1_data_proto_rawDesc = "" +
	"\n" +
	"\x1dcloud/firestore/v1/data.proto\x12 google.events.cloud.firestore.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18google/type/latlng.proto\"\xef\x01\n" +
	"\x11DocumentEventData\x12@\n" +
	"\x05value\x18\x01 \x01(\v2*.google.events.cloud.firestore.v1.DocumentR\x05value\x12G\n" +
	"\told_value\x18\x02 \x01(\v2*.google.events.cloud.firestore.v1.DocumentR\boldValue\x12O\n" +
	"\vupdate_mask\x18\x03 \x01(\v2..google.events.cloud.firestore.v1.DocumentMaskR\n" +
	"updateMask\"/\n" +
	"\fDocumentMask\x12\x1f\n" +
	"\vfield_paths\x18\x01 \x03(\tR\n" +
	"fieldPaths\"\xcc\x02\n" +
	"\bDocument\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12N\n" +
	"\x06fields\x18\x02 \x03(\v26.google.events.cloud.firestore.v1.Document.FieldsEntryR\x06fields\x12;\n" +
	"\vcreate_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x1ab\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12=\n" +
	"\x05value\x18\x02 \x01(\v2'.google.events.cloud.firestore.v1.ValueR\x05value:\x028\x01\"\xda\x04\n" +
	"\x05Value\x12;\n" +
	"\n" +
	"null_value\x18\v \x01(\x0e2\x1a.google.protobuf.NullValueH\x00R\tnullValue\x12%\n" +
	"\rboolean_value\x18\x01 \x01(\bH\x00R\fbooleanValue\x12%\n" +
	"\rinteger_value\x18\x02 \x01(\x03H\x00R\fintegerValue\x12#\n" +
	"\fdouble_value\x18\x03 \x01(\x01H\x00R\vdoubleValue\x12E\n" +
	"\x0ftimestamp_value\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12#\n" +
	"\fstring_value\x18\x11 \x01(\tH\x00R\vstringValue\x12!\n" +
	"\vbytes_value\x18\x12 \x01(\fH\x00R\n" +
	"bytesValue\x12)\n" +
	"\x0freference_value\x18\x05 \x01(\tH\x00R\x0ereferenceValue\x12=\n" +
	"\x0fgeo_point_value\x18\b \x01(\v2\x13.google.type.LatLngH\x00R\rgeoPointValue\x12O\n" +
	"\varray_value\x18\t \x01(\v2,.google.events.cloud.firestore.v1.ArrayValueH\x00R\n" +
	"arrayValue\x12I\n" +
	"\tmap_value\x18\x06 \x01(\v2*.google.events.cloud.firestore.v1.MapValueH\x00R\bmapValueB\f\n" +
	"\n" +
	"value_type\"M\n" +
	"\n" +
	"ArrayValue\x12?\n" +
	"\x06values\x18\x01 \x03(\v2'.google.events.cloud.firestore.v1.ValueR\x06values\"\xbe\x01\n" +
	"\bMapValue\x12N\n" +
	"\x06fields\x18\x01 \x03(\v26.google.events.cloud.firestore.v1.MapValue.FieldsEntryR\x06fields\x1ab\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12=\n" +
	"\x05value\x18\x02 \x01(\v2'.google.events.cloud.firestore.v1.ValueR\x05value:\x028\x01Bv\xaa\x02)Google.Events.Protobuf.Cloud.Firestore.V1\xca\x02 Google\\Events\\Cloud\\Firestore\\V1\xea\x02$Google::Events::Cloud::Firestore::V1b\x06proto3"

var (
	file_cloud_firestore_v1_data_proto_rawDescOnce sync.Once
	file_cloud_firestore_v1_data_proto_rawDescData []byte
)

func file_cloud_firestore_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_firestore_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_firestore_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloud_firestore_v1_data_proto_rawDesc), len(file_cloud_firestore_v1_data_proto_rawDesc)))
	})
	return file_cloud_firestore_v1_data_proto_rawDescData
}

var file_cloud_firestore_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cloud_firestore_v1_data_proto_goTypes = []any{
	(*DocumentEventData)(nil),     // 0: google.events.cloud.firestore.v1.DocumentEventData
	(*DocumentMask)(nil),          // 1: google.events.cloud.firestore.v1.DocumentMask
	(*Document)(nil),              // 2: google.events.cloud.firestore.v1.Document
	(*Value)(nil),                 // 3: google.events.cloud.firestore.v1.Value
	(*ArrayValue)(nil),            // 4: google.events.cloud.firestore.v1.ArrayValue
	(*MapValue)(nil),              // 5: google.events.cloud.firestore.v1.MapValue
	nil,                           // 6: google.events.cloud.firestore.v1.Document.FieldsEntry
	nil,                           // 7: google.events.cloud.firestore.v1.MapValue.FieldsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(structpb.NullValue)(0),       // 9: google.protobuf.NullValue
	(*latlng.LatLng)(nil),         // 10: google.type.LatLng
}
var file_cloud_firestore_v1_data_proto_depIdxs = []int32{
	2,  // 0: google.events.cloud.firestore.v1.DocumentEventData.value:type_name -> google.events.cloud.firestore.v1.Document
	2,  // 1: google.events.cloud.firestore.v1.DocumentEventData.old_value:type_name -> google.events.cloud.firestore.v1.Document
	1,  // 2: google.events.cloud.firestore.v1.DocumentEventData.update_mask:type_name -> google.events.cloud.firestore.v1.DocumentMask
	6,  // 3: google.events.cloud.firestore.v1.Document.fields:type_name -> google.events.cloud.firestore.v1.Document.FieldsEntry
	8,  // 4: google.events.cloud.firestore.v1.Document.create_time:type_name -> google.protobuf.Timestamp
	8,  // 5: google.events.cloud.firestore.v1.Document.update_time:type_name -> google.protobuf.Timestamp
	9,  // 6: google.events.cloud.firestore.v1.Value.null_value:type_name -> google.protobuf.NullValue
	8,  // 7: google.events.cloud.firestore.v1.Value.timestamp_value:type_name -> google.protobuf.Timestamp
	10, // 8: google.events.cloud.firestore.v1.Value.geo_point_value:type_name -> google.type.LatLng
	4,  // 9: google.events.cloud.firestore.v1.Value.array_value:type_name -> google.events.cloud.firestore.v1.ArrayValue
	5,  // 10: google.events.cloud.firestore.v1.Value.map_value:type_name -> google.events.cloud.firestore.v1.MapValue
	3,  // 11: google.events.cloud.firestore.v1.ArrayValue.values:type_name -> google.events.cloud.firestore.v1.Value
	7,  // 12: google.events.cloud.firestore.v1.MapValue.fields:type_name -> google.events.cloud.firestore.v1.MapValue.FieldsEntry
	3,  // 13: google.events.cloud.firestore.v1.Document.FieldsEntry.value:type_name -> google.events.cloud.firestore.v1.Value
	3,  // 14: google.events.cloud.firestore.v1.MapValue.FieldsEntry.value:type_name -> google.events.cloud.firestore.v1.Value
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_cloud_firestore_v1_data_proto_init() }
func file_cloud_firestore_v1_data_proto_init() {
	if File_cloud_firestore_v1_data_proto != nil {
		return
	}
	file_cloud_firestore_v1_data_proto_msgTypes[3].OneofWrappers = []any{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_ReferenceValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_ArrayValue)(nil),
		(*Value_MapValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloud_firestore_v1_data_proto_rawDesc), len(file_cloud_firestore_v1_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_firestore_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_firestore_v1_data_proto_depIdxs,
		MessageInfos:      file_cloud_firestore_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_firestore_v1_data_proto = out.File
	file_cloud_firestore_v1_data_proto_goTypes = nil
	file_cloud_firestore_v1_data_proto_depIdxs = nil
}
