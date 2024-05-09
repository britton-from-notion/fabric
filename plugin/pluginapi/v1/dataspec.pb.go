// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: pluginapi/v1/dataspec.proto

package pluginapiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*Spec_Attr
	//	*Spec_Block
	//	*Spec_ObjSpec
	//	*Spec_ObjDump
	//	*Spec_Opaque
	Data isSpec_Data `protobuf_oneof:"data"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{0}
}

func (m *Spec) GetData() isSpec_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Spec) GetAttr() *AttrSpec {
	if x, ok := x.GetData().(*Spec_Attr); ok {
		return x.Attr
	}
	return nil
}

func (x *Spec) GetBlock() *BlockSpec {
	if x, ok := x.GetData().(*Spec_Block); ok {
		return x.Block
	}
	return nil
}

func (x *Spec) GetObjSpec() *ObjectSpec {
	if x, ok := x.GetData().(*Spec_ObjSpec); ok {
		return x.ObjSpec
	}
	return nil
}

func (x *Spec) GetObjDump() *ObjDumpSpec {
	if x, ok := x.GetData().(*Spec_ObjDump); ok {
		return x.ObjDump
	}
	return nil
}

func (x *Spec) GetOpaque() *OpaqueSpec {
	if x, ok := x.GetData().(*Spec_Opaque); ok {
		return x.Opaque
	}
	return nil
}

type isSpec_Data interface {
	isSpec_Data()
}

type Spec_Attr struct {
	Attr *AttrSpec `protobuf:"bytes,1,opt,name=attr,proto3,oneof"`
}

type Spec_Block struct {
	Block *BlockSpec `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

type Spec_ObjSpec struct {
	ObjSpec *ObjectSpec `protobuf:"bytes,3,opt,name=obj_spec,json=objSpec,proto3,oneof"`
}

type Spec_ObjDump struct {
	ObjDump *ObjDumpSpec `protobuf:"bytes,4,opt,name=obj_dump,json=objDump,proto3,oneof"`
}

type Spec_Opaque struct {
	Opaque *OpaqueSpec `protobuf:"bytes,5,opt,name=opaque,proto3,oneof"`
}

func (*Spec_Attr) isSpec_Data() {}

func (*Spec_Block) isSpec_Data() {}

func (*Spec_ObjSpec) isSpec_Data() {}

func (*Spec_ObjDump) isSpec_Data() {}

func (*Spec_Opaque) isSpec_Data() {}

type AttrSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       *CtyType  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	DefaultVal *CtyValue `protobuf:"bytes,3,opt,name=default_val,json=defaultVal,proto3" json:"default_val,omitempty"`
	ExampleVal *CtyValue `protobuf:"bytes,4,opt,name=example_val,json=exampleVal,proto3" json:"example_val,omitempty"`
	Doc        string    `protobuf:"bytes,5,opt,name=doc,proto3" json:"doc,omitempty"`
	// bool     required    = 6; // made obsolete by constraints system
	Constraints  uint32      `protobuf:"varint,7,opt,name=constraints,proto3" json:"constraints,omitempty"`
	OneOf        []*CtyValue `protobuf:"bytes,8,rep,name=one_of,json=oneOf,proto3" json:"one_of,omitempty"`
	MinInclusive *CtyValue   `protobuf:"bytes,9,opt,name=min_inclusive,json=minInclusive,proto3" json:"min_inclusive,omitempty"`
	MaxInclusive *CtyValue   `protobuf:"bytes,10,opt,name=max_inclusive,json=maxInclusive,proto3" json:"max_inclusive,omitempty"`
	Deprecated   string      `protobuf:"bytes,11,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *AttrSpec) Reset() {
	*x = AttrSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrSpec) ProtoMessage() {}

func (x *AttrSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrSpec.ProtoReflect.Descriptor instead.
func (*AttrSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{1}
}

func (x *AttrSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrSpec) GetType() *CtyType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *AttrSpec) GetDefaultVal() *CtyValue {
	if x != nil {
		return x.DefaultVal
	}
	return nil
}

func (x *AttrSpec) GetExampleVal() *CtyValue {
	if x != nil {
		return x.ExampleVal
	}
	return nil
}

func (x *AttrSpec) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

func (x *AttrSpec) GetConstraints() uint32 {
	if x != nil {
		return x.Constraints
	}
	return 0
}

func (x *AttrSpec) GetOneOf() []*CtyValue {
	if x != nil {
		return x.OneOf
	}
	return nil
}

func (x *AttrSpec) GetMinInclusive() *CtyValue {
	if x != nil {
		return x.MinInclusive
	}
	return nil
}

func (x *AttrSpec) GetMaxInclusive() *CtyValue {
	if x != nil {
		return x.MaxInclusive
	}
	return nil
}

func (x *AttrSpec) GetDeprecated() string {
	if x != nil {
		return x.Deprecated
	}
	return ""
}

type BlockSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Nested   *Spec  `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
	Doc      string `protobuf:"bytes,3,opt,name=doc,proto3" json:"doc,omitempty"`
	Required bool   `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *BlockSpec) Reset() {
	*x = BlockSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSpec) ProtoMessage() {}

func (x *BlockSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSpec.ProtoReflect.Descriptor instead.
func (*BlockSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{2}
}

func (x *BlockSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockSpec) GetNested() *Spec {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *BlockSpec) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

func (x *BlockSpec) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

type ObjectSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Specs []*ObjectSpec_ObjectSpecChild `protobuf:"bytes,1,rep,name=specs,proto3" json:"specs,omitempty"`
}

func (x *ObjectSpec) Reset() {
	*x = ObjectSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSpec) ProtoMessage() {}

func (x *ObjectSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSpec.ProtoReflect.Descriptor instead.
func (*ObjectSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectSpec) GetSpecs() []*ObjectSpec_ObjectSpecChild {
	if x != nil {
		return x.Specs
	}
	return nil
}

type ObjDumpSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc string `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *ObjDumpSpec) Reset() {
	*x = ObjDumpSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjDumpSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjDumpSpec) ProtoMessage() {}

func (x *ObjDumpSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjDumpSpec.ProtoReflect.Descriptor instead.
func (*ObjDumpSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{4}
}

func (x *ObjDumpSpec) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

type OpaqueSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *HclSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Doc  string   `protobuf:"bytes,2,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *OpaqueSpec) Reset() {
	*x = OpaqueSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpaqueSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpaqueSpec) ProtoMessage() {}

func (x *OpaqueSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpaqueSpec.ProtoReflect.Descriptor instead.
func (*OpaqueSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{5}
}

func (x *OpaqueSpec) GetSpec() *HclSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *OpaqueSpec) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

type KeyForObjectSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *Spec  `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyForObjectSpec) Reset() {
	*x = KeyForObjectSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyForObjectSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyForObjectSpec) ProtoMessage() {}

func (x *KeyForObjectSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyForObjectSpec.ProtoReflect.Descriptor instead.
func (*KeyForObjectSpec) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{6}
}

func (x *KeyForObjectSpec) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *KeyForObjectSpec) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ObjectSpec_ObjectSpecChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*ObjectSpec_ObjectSpecChild_Attr
	//	*ObjectSpec_ObjectSpecChild_Block
	//	*ObjectSpec_ObjectSpecChild_Named
	Data isObjectSpec_ObjectSpecChild_Data `protobuf_oneof:"data"`
}

func (x *ObjectSpec_ObjectSpecChild) Reset() {
	*x = ObjectSpec_ObjectSpecChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginapi_v1_dataspec_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectSpec_ObjectSpecChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSpec_ObjectSpecChild) ProtoMessage() {}

func (x *ObjectSpec_ObjectSpecChild) ProtoReflect() protoreflect.Message {
	mi := &file_pluginapi_v1_dataspec_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSpec_ObjectSpecChild.ProtoReflect.Descriptor instead.
func (*ObjectSpec_ObjectSpecChild) Descriptor() ([]byte, []int) {
	return file_pluginapi_v1_dataspec_proto_rawDescGZIP(), []int{3, 0}
}

func (m *ObjectSpec_ObjectSpecChild) GetData() isObjectSpec_ObjectSpecChild_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ObjectSpec_ObjectSpecChild) GetAttr() *AttrSpec {
	if x, ok := x.GetData().(*ObjectSpec_ObjectSpecChild_Attr); ok {
		return x.Attr
	}
	return nil
}

func (x *ObjectSpec_ObjectSpecChild) GetBlock() *BlockSpec {
	if x, ok := x.GetData().(*ObjectSpec_ObjectSpecChild_Block); ok {
		return x.Block
	}
	return nil
}

func (x *ObjectSpec_ObjectSpecChild) GetNamed() *KeyForObjectSpec {
	if x, ok := x.GetData().(*ObjectSpec_ObjectSpecChild_Named); ok {
		return x.Named
	}
	return nil
}

type isObjectSpec_ObjectSpecChild_Data interface {
	isObjectSpec_ObjectSpecChild_Data()
}

type ObjectSpec_ObjectSpecChild_Attr struct {
	Attr *AttrSpec `protobuf:"bytes,1,opt,name=attr,proto3,oneof"`
}

type ObjectSpec_ObjectSpecChild_Block struct {
	Block *BlockSpec `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

type ObjectSpec_ObjectSpecChild_Named struct {
	Named *KeyForObjectSpec `protobuf:"bytes,3,opt,name=named,proto3,oneof"`
}

func (*ObjectSpec_ObjectSpecChild_Attr) isObjectSpec_ObjectSpecChild_Data() {}

func (*ObjectSpec_ObjectSpecChild_Block) isObjectSpec_ObjectSpecChild_Data() {}

func (*ObjectSpec_ObjectSpecChild_Named) isObjectSpec_ObjectSpecChild_Data() {}

var File_pluginapi_v1_dataspec_proto protoreflect.FileDescriptor

var file_pluginapi_v1_dataspec_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x63, 0x6c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00,
	0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x53, 0x70, 0x65, 0x63, 0x12, 0x36,
	0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x44, 0x75, 0x6d, 0x70, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x07, 0x6f,
	0x62, 0x6a, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xb8, 0x03, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x74, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x6f, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6f, 0x6e,
	0x65, 0x4f, 0x66, 0x12, 0x3b, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x74, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65,
	0x12, 0x3b, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0c, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x79, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0xff, 0x01, 0x0a, 0x0a, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3e, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x52, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a, 0xb0, 0x01, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x48, 0x00, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x65,
	0x63, 0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x0a, 0x0b, 0x4f, 0x62,
	0x6a, 0x44, 0x75, 0x6d, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x22, 0x49, 0x0a, 0x0a, 0x4f,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x63, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x22, 0x4c, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x42, 0xb3, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x6b, 0x2d, 0x69, 0x6f, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x18, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pluginapi_v1_dataspec_proto_rawDescOnce sync.Once
	file_pluginapi_v1_dataspec_proto_rawDescData = file_pluginapi_v1_dataspec_proto_rawDesc
)

func file_pluginapi_v1_dataspec_proto_rawDescGZIP() []byte {
	file_pluginapi_v1_dataspec_proto_rawDescOnce.Do(func() {
		file_pluginapi_v1_dataspec_proto_rawDescData = protoimpl.X.CompressGZIP(file_pluginapi_v1_dataspec_proto_rawDescData)
	})
	return file_pluginapi_v1_dataspec_proto_rawDescData
}

var file_pluginapi_v1_dataspec_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pluginapi_v1_dataspec_proto_goTypes = []interface{}{
	(*Spec)(nil),                       // 0: pluginapi.v1.Spec
	(*AttrSpec)(nil),                   // 1: pluginapi.v1.AttrSpec
	(*BlockSpec)(nil),                  // 2: pluginapi.v1.BlockSpec
	(*ObjectSpec)(nil),                 // 3: pluginapi.v1.ObjectSpec
	(*ObjDumpSpec)(nil),                // 4: pluginapi.v1.ObjDumpSpec
	(*OpaqueSpec)(nil),                 // 5: pluginapi.v1.OpaqueSpec
	(*KeyForObjectSpec)(nil),           // 6: pluginapi.v1.KeyForObjectSpec
	(*ObjectSpec_ObjectSpecChild)(nil), // 7: pluginapi.v1.ObjectSpec.ObjectSpecChild
	(*CtyType)(nil),                    // 8: pluginapi.v1.CtyType
	(*CtyValue)(nil),                   // 9: pluginapi.v1.CtyValue
	(*HclSpec)(nil),                    // 10: pluginapi.v1.HclSpec
}
var file_pluginapi_v1_dataspec_proto_depIdxs = []int32{
	1,  // 0: pluginapi.v1.Spec.attr:type_name -> pluginapi.v1.AttrSpec
	2,  // 1: pluginapi.v1.Spec.block:type_name -> pluginapi.v1.BlockSpec
	3,  // 2: pluginapi.v1.Spec.obj_spec:type_name -> pluginapi.v1.ObjectSpec
	4,  // 3: pluginapi.v1.Spec.obj_dump:type_name -> pluginapi.v1.ObjDumpSpec
	5,  // 4: pluginapi.v1.Spec.opaque:type_name -> pluginapi.v1.OpaqueSpec
	8,  // 5: pluginapi.v1.AttrSpec.type:type_name -> pluginapi.v1.CtyType
	9,  // 6: pluginapi.v1.AttrSpec.default_val:type_name -> pluginapi.v1.CtyValue
	9,  // 7: pluginapi.v1.AttrSpec.example_val:type_name -> pluginapi.v1.CtyValue
	9,  // 8: pluginapi.v1.AttrSpec.one_of:type_name -> pluginapi.v1.CtyValue
	9,  // 9: pluginapi.v1.AttrSpec.min_inclusive:type_name -> pluginapi.v1.CtyValue
	9,  // 10: pluginapi.v1.AttrSpec.max_inclusive:type_name -> pluginapi.v1.CtyValue
	0,  // 11: pluginapi.v1.BlockSpec.nested:type_name -> pluginapi.v1.Spec
	7,  // 12: pluginapi.v1.ObjectSpec.specs:type_name -> pluginapi.v1.ObjectSpec.ObjectSpecChild
	10, // 13: pluginapi.v1.OpaqueSpec.spec:type_name -> pluginapi.v1.HclSpec
	0,  // 14: pluginapi.v1.KeyForObjectSpec.spec:type_name -> pluginapi.v1.Spec
	1,  // 15: pluginapi.v1.ObjectSpec.ObjectSpecChild.attr:type_name -> pluginapi.v1.AttrSpec
	2,  // 16: pluginapi.v1.ObjectSpec.ObjectSpecChild.block:type_name -> pluginapi.v1.BlockSpec
	6,  // 17: pluginapi.v1.ObjectSpec.ObjectSpecChild.named:type_name -> pluginapi.v1.KeyForObjectSpec
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_pluginapi_v1_dataspec_proto_init() }
func file_pluginapi_v1_dataspec_proto_init() {
	if File_pluginapi_v1_dataspec_proto != nil {
		return
	}
	file_pluginapi_v1_cty_proto_init()
	file_pluginapi_v1_hclspec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pluginapi_v1_dataspec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjDumpSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpaqueSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyForObjectSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pluginapi_v1_dataspec_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectSpec_ObjectSpecChild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pluginapi_v1_dataspec_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Spec_Attr)(nil),
		(*Spec_Block)(nil),
		(*Spec_ObjSpec)(nil),
		(*Spec_ObjDump)(nil),
		(*Spec_Opaque)(nil),
	}
	file_pluginapi_v1_dataspec_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ObjectSpec_ObjectSpecChild_Attr)(nil),
		(*ObjectSpec_ObjectSpecChild_Block)(nil),
		(*ObjectSpec_ObjectSpecChild_Named)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pluginapi_v1_dataspec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pluginapi_v1_dataspec_proto_goTypes,
		DependencyIndexes: file_pluginapi_v1_dataspec_proto_depIdxs,
		MessageInfos:      file_pluginapi_v1_dataspec_proto_msgTypes,
	}.Build()
	File_pluginapi_v1_dataspec_proto = out.File
	file_pluginapi_v1_dataspec_proto_rawDesc = nil
	file_pluginapi_v1_dataspec_proto_goTypes = nil
	file_pluginapi_v1_dataspec_proto_depIdxs = nil
}
