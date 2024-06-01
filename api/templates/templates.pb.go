// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: templates.proto

package templates

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

type FeatureStruct_FeatureType int32

const (
	FeatureStruct_RANGE FeatureStruct_FeatureType = 0
	FeatureStruct_LIST  FeatureStruct_FeatureType = 1
)

// Enum value maps for FeatureStruct_FeatureType.
var (
	FeatureStruct_FeatureType_name = map[int32]string{
		0: "RANGE",
		1: "LIST",
	}
	FeatureStruct_FeatureType_value = map[string]int32{
		"RANGE": 0,
		"LIST":  1,
	}
)

func (x FeatureStruct_FeatureType) Enum() *FeatureStruct_FeatureType {
	p := new(FeatureStruct_FeatureType)
	*p = x
	return p
}

func (x FeatureStruct_FeatureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureStruct_FeatureType) Descriptor() protoreflect.EnumDescriptor {
	return file_templates_proto_enumTypes[0].Descriptor()
}

func (FeatureStruct_FeatureType) Type() protoreflect.EnumType {
	return &file_templates_proto_enumTypes[0]
}

func (x FeatureStruct_FeatureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureStruct_FeatureType.Descriptor instead.
func (FeatureStruct_FeatureType) EnumDescriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{4, 0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{0}
}

type FeatureLinkTemplateStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeatureId  uint64 `protobuf:"varint,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	TemplateId uint64 `protobuf:"varint,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FeatureLinkTemplateStruct) Reset() {
	*x = FeatureLinkTemplateStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureLinkTemplateStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureLinkTemplateStruct) ProtoMessage() {}

func (x *FeatureLinkTemplateStruct) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureLinkTemplateStruct.ProtoReflect.Descriptor instead.
func (*FeatureLinkTemplateStruct) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureLinkTemplateStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeatureLinkTemplateStruct) GetFeatureId() uint64 {
	if x != nil {
		return x.FeatureId
	}
	return 0
}

func (x *FeatureLinkTemplateStruct) GetTemplateId() uint64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *FeatureLinkTemplateStruct) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TemplateStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TemplateStruct) Reset() {
	*x = TemplateStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateStruct) ProtoMessage() {}

func (x *TemplateStruct) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateStruct.ProtoReflect.Descriptor instead.
func (*TemplateStruct) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateStruct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type TemplatesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TemplateStruct `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TemplatesList) Reset() {
	*x = TemplatesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplatesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplatesList) ProtoMessage() {}

func (x *TemplatesList) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplatesList.ProtoReflect.Descriptor instead.
func (*TemplatesList) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{3}
}

func (x *TemplatesList) GetItems() []*TemplateStruct {
	if x != nil {
		return x.Items
	}
	return nil
}

type FeatureStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FeatureType FeatureStruct_FeatureType `protobuf:"varint,3,opt,name=feature_type,json=featureType,proto3,enum=TemplatesService.FeatureStruct_FeatureType" json:"feature_type,omitempty"`
}

func (x *FeatureStruct) Reset() {
	*x = FeatureStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureStruct) ProtoMessage() {}

func (x *FeatureStruct) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureStruct.ProtoReflect.Descriptor instead.
func (*FeatureStruct) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{4}
}

func (x *FeatureStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeatureStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureStruct) GetFeatureType() FeatureStruct_FeatureType {
	if x != nil {
		return x.FeatureType
	}
	return FeatureStruct_RANGE
}

type FeaturesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FeatureStruct `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *FeaturesList) Reset() {
	*x = FeaturesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturesList) ProtoMessage() {}

func (x *FeaturesList) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturesList.ProtoReflect.Descriptor instead.
func (*FeaturesList) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{5}
}

func (x *FeaturesList) GetItems() []*FeatureStruct {
	if x != nil {
		return x.Items
	}
	return nil
}

type IdStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdStruct) Reset() {
	*x = IdStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdStruct) ProtoMessage() {}

func (x *IdStruct) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdStruct.ProtoReflect.Descriptor instead.
func (*IdStruct) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{6}
}

func (x *IdStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// фича с ссылкой
type FeatureLinkTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link    *FeatureLinkTemplateStruct `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Feature *FeatureStruct             `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *FeatureLinkTemplate) Reset() {
	*x = FeatureLinkTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureLinkTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureLinkTemplate) ProtoMessage() {}

func (x *FeatureLinkTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureLinkTemplate.ProtoReflect.Descriptor instead.
func (*FeatureLinkTemplate) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{7}
}

func (x *FeatureLinkTemplate) GetLink() *FeatureLinkTemplateStruct {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *FeatureLinkTemplate) GetFeature() *FeatureStruct {
	if x != nil {
		return x.Feature
	}
	return nil
}

// хранит на 1 шаблон кучу фич с описанием данных из ссылок
type HibridFeatureLinkTemplateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FeatureLinkTemplate `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *HibridFeatureLinkTemplateList) Reset() {
	*x = HibridFeatureLinkTemplateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_templates_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HibridFeatureLinkTemplateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HibridFeatureLinkTemplateList) ProtoMessage() {}

func (x *HibridFeatureLinkTemplateList) ProtoReflect() protoreflect.Message {
	mi := &file_templates_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HibridFeatureLinkTemplateList.ProtoReflect.Descriptor instead.
func (*HibridFeatureLinkTemplateList) Descriptor() ([]byte, []int) {
	return file_templates_proto_rawDescGZIP(), []int{8}
}

func (x *HibridFeatureLinkTemplateList) GetItems() []*FeatureLinkTemplate {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_templates_proto protoreflect.FileDescriptor

var file_templates_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x81, 0x01, 0x0a,
	0x19, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x56, 0x0a, 0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x22, 0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x22, 0x45, 0x0a, 0x0c, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x91,
	0x01, 0x0a, 0x13, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x39, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x5c, 0x0a, 0x1d, 0x48, 0x69, 0x62, 0x72, 0x69, 0x64, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0x82, 0x07, 0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x55,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x52, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x4b, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12,
	0x49, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x66, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x79, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x1a, 0x2f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x62, 0x72, 0x69, 0x64, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_templates_proto_rawDescOnce sync.Once
	file_templates_proto_rawDescData = file_templates_proto_rawDesc
)

func file_templates_proto_rawDescGZIP() []byte {
	file_templates_proto_rawDescOnce.Do(func() {
		file_templates_proto_rawDescData = protoimpl.X.CompressGZIP(file_templates_proto_rawDescData)
	})
	return file_templates_proto_rawDescData
}

var file_templates_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_templates_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_templates_proto_goTypes = []interface{}{
	(FeatureStruct_FeatureType)(0),        // 0: TemplatesService.FeatureStruct.FeatureType
	(*Empty)(nil),                         // 1: TemplatesService.Empty
	(*FeatureLinkTemplateStruct)(nil),     // 2: TemplatesService.FeatureLinkTemplateStruct
	(*TemplateStruct)(nil),                // 3: TemplatesService.TemplateStruct
	(*TemplatesList)(nil),                 // 4: TemplatesService.TemplatesList
	(*FeatureStruct)(nil),                 // 5: TemplatesService.FeatureStruct
	(*FeaturesList)(nil),                  // 6: TemplatesService.FeaturesList
	(*IdStruct)(nil),                      // 7: TemplatesService.IdStruct
	(*FeatureLinkTemplate)(nil),           // 8: TemplatesService.FeatureLinkTemplate
	(*HibridFeatureLinkTemplateList)(nil), // 9: TemplatesService.HibridFeatureLinkTemplateList
}
var file_templates_proto_depIdxs = []int32{
	3,  // 0: TemplatesService.TemplatesList.items:type_name -> TemplatesService.TemplateStruct
	0,  // 1: TemplatesService.FeatureStruct.feature_type:type_name -> TemplatesService.FeatureStruct.FeatureType
	5,  // 2: TemplatesService.FeaturesList.items:type_name -> TemplatesService.FeatureStruct
	2,  // 3: TemplatesService.FeatureLinkTemplate.link:type_name -> TemplatesService.FeatureLinkTemplateStruct
	5,  // 4: TemplatesService.FeatureLinkTemplate.feature:type_name -> TemplatesService.FeatureStruct
	8,  // 5: TemplatesService.HibridFeatureLinkTemplateList.items:type_name -> TemplatesService.FeatureLinkTemplate
	2,  // 6: TemplatesService.Templates.CreateLink:input_type -> TemplatesService.FeatureLinkTemplateStruct
	2,  // 7: TemplatesService.Templates.UpdateLink:input_type -> TemplatesService.FeatureLinkTemplateStruct
	2,  // 8: TemplatesService.Templates.DeleteLink:input_type -> TemplatesService.FeatureLinkTemplateStruct
	3,  // 9: TemplatesService.Templates.CreateTemplate:input_type -> TemplatesService.TemplateStruct
	3,  // 10: TemplatesService.Templates.UpdateTemplate:input_type -> TemplatesService.TemplateStruct
	7,  // 11: TemplatesService.Templates.DeleteTemplate:input_type -> TemplatesService.IdStruct
	5,  // 12: TemplatesService.Templates.CreateFeature:input_type -> TemplatesService.FeatureStruct
	5,  // 13: TemplatesService.Templates.UpdateFeature:input_type -> TemplatesService.FeatureStruct
	7,  // 14: TemplatesService.Templates.DeleteFeature:input_type -> TemplatesService.IdStruct
	1,  // 15: TemplatesService.Templates.GetAllTemplates:input_type -> TemplatesService.Empty
	7,  // 16: TemplatesService.Templates.GetFeaturesByTemplateId:input_type -> TemplatesService.IdStruct
	7,  // 17: TemplatesService.Templates.CreateLink:output_type -> TemplatesService.IdStruct
	1,  // 18: TemplatesService.Templates.UpdateLink:output_type -> TemplatesService.Empty
	1,  // 19: TemplatesService.Templates.DeleteLink:output_type -> TemplatesService.Empty
	7,  // 20: TemplatesService.Templates.CreateTemplate:output_type -> TemplatesService.IdStruct
	1,  // 21: TemplatesService.Templates.UpdateTemplate:output_type -> TemplatesService.Empty
	1,  // 22: TemplatesService.Templates.DeleteTemplate:output_type -> TemplatesService.Empty
	7,  // 23: TemplatesService.Templates.CreateFeature:output_type -> TemplatesService.IdStruct
	1,  // 24: TemplatesService.Templates.UpdateFeature:output_type -> TemplatesService.Empty
	1,  // 25: TemplatesService.Templates.DeleteFeature:output_type -> TemplatesService.Empty
	4,  // 26: TemplatesService.Templates.GetAllTemplates:output_type -> TemplatesService.TemplatesList
	9,  // 27: TemplatesService.Templates.GetFeaturesByTemplateId:output_type -> TemplatesService.HibridFeatureLinkTemplateList
	17, // [17:28] is the sub-list for method output_type
	6,  // [6:17] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_templates_proto_init() }
func file_templates_proto_init() {
	if File_templates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_templates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_templates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureLinkTemplateStruct); i {
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
		file_templates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateStruct); i {
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
		file_templates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplatesList); i {
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
		file_templates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureStruct); i {
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
		file_templates_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturesList); i {
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
		file_templates_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdStruct); i {
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
		file_templates_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureLinkTemplate); i {
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
		file_templates_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HibridFeatureLinkTemplateList); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_templates_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_templates_proto_goTypes,
		DependencyIndexes: file_templates_proto_depIdxs,
		EnumInfos:         file_templates_proto_enumTypes,
		MessageInfos:      file_templates_proto_msgTypes,
	}.Build()
	File_templates_proto = out.File
	file_templates_proto_rawDesc = nil
	file_templates_proto_goTypes = nil
	file_templates_proto_depIdxs = nil
}
