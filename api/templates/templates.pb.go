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

	FeatureId  uint64 `protobuf:"varint,1,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	TemplateId uint64 `protobuf:"varint,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
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

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PriorityId uint64 `protobuf:"varint,3,opt,name=priority_id,json=priorityId,proto3" json:"priority_id,omitempty"`
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

func (x *FeatureStruct) GetPriorityId() uint64 {
	if x != nil {
		return x.PriorityId
	}
	return 0
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

var File_templates_proto protoreflect.FileDescriptor

var file_templates_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x19,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x47, 0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x54, 0x0a, 0x0d, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x45, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xf8, 0x03, 0x0a, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12,
	0x4e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12,
	0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x17, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x17, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x1e,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0f,
	0x5a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_templates_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_templates_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: TemplatesService.Empty
	(*FeatureLinkTemplateStruct)(nil), // 1: TemplatesService.FeatureLinkTemplateStruct
	(*TemplateStruct)(nil),            // 2: TemplatesService.TemplateStruct
	(*TemplatesList)(nil),             // 3: TemplatesService.TemplatesList
	(*FeatureStruct)(nil),             // 4: TemplatesService.FeatureStruct
	(*FeaturesList)(nil),              // 5: TemplatesService.FeaturesList
	(*IdStruct)(nil),                  // 6: TemplatesService.IdStruct
}
var file_templates_proto_depIdxs = []int32{
	2, // 0: TemplatesService.TemplatesList.items:type_name -> TemplatesService.TemplateStruct
	4, // 1: TemplatesService.FeaturesList.items:type_name -> TemplatesService.FeatureStruct
	1, // 2: TemplatesService.TemplatesStorage.CreateLink:input_type -> TemplatesService.FeatureLinkTemplateStruct
	2, // 3: TemplatesService.TemplatesStorage.CreateTemplate:input_type -> TemplatesService.TemplateStruct
	6, // 4: TemplatesService.TemplatesStorage.DeleteTemplate:input_type -> TemplatesService.IdStruct
	1, // 5: TemplatesService.TemplatesStorage.DeleteLink:input_type -> TemplatesService.FeatureLinkTemplateStruct
	0, // 6: TemplatesService.TemplatesStorage.GetAllTemplates:input_type -> TemplatesService.Empty
	6, // 7: TemplatesService.TemplatesStorage.GetFeaturesByTemplateId:input_type -> TemplatesService.IdStruct
	6, // 8: TemplatesService.TemplatesStorage.CreateLink:output_type -> TemplatesService.IdStruct
	6, // 9: TemplatesService.TemplatesStorage.CreateTemplate:output_type -> TemplatesService.IdStruct
	0, // 10: TemplatesService.TemplatesStorage.DeleteTemplate:output_type -> TemplatesService.Empty
	0, // 11: TemplatesService.TemplatesStorage.DeleteLink:output_type -> TemplatesService.Empty
	3, // 12: TemplatesService.TemplatesStorage.GetAllTemplates:output_type -> TemplatesService.TemplatesList
	5, // 13: TemplatesService.TemplatesStorage.GetFeaturesByTemplateId:output_type -> TemplatesService.FeaturesList
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_templates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_templates_proto_goTypes,
		DependencyIndexes: file_templates_proto_depIdxs,
		MessageInfos:      file_templates_proto_msgTypes,
	}.Build()
	File_templates_proto = out.File
	file_templates_proto_rawDesc = nil
	file_templates_proto_goTypes = nil
	file_templates_proto_depIdxs = nil
}