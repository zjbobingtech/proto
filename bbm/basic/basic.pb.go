// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.3
// source: bbm/basic/basic.proto

package basic

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

type BasicBehaviorTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BehaviorType    string `protobuf:"bytes,1,opt,name=behavior_type,json=behaviorType,proto3" json:"behavior_type,omitempty"`
	BehaviorContent string `protobuf:"bytes,2,opt,name=behavior_content,json=behaviorContent,proto3" json:"behavior_content,omitempty"`
}

func (x *BasicBehaviorTraceRequest) Reset() {
	*x = BasicBehaviorTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicBehaviorTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicBehaviorTraceRequest) ProtoMessage() {}

func (x *BasicBehaviorTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicBehaviorTraceRequest.ProtoReflect.Descriptor instead.
func (*BasicBehaviorTraceRequest) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{0}
}

func (x *BasicBehaviorTraceRequest) GetBehaviorType() string {
	if x != nil {
		return x.BehaviorType
	}
	return ""
}

func (x *BasicBehaviorTraceRequest) GetBehaviorContent() string {
	if x != nil {
		return x.BehaviorContent
	}
	return ""
}

type BasicGlobalParametersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64                        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Code       string                       `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Data       []*BasicGlobalParametersData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BasicGlobalParametersResponse) Reset() {
	*x = BasicGlobalParametersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicGlobalParametersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicGlobalParametersResponse) ProtoMessage() {}

func (x *BasicGlobalParametersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicGlobalParametersResponse.ProtoReflect.Descriptor instead.
func (*BasicGlobalParametersResponse) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{1}
}

func (x *BasicGlobalParametersResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BasicGlobalParametersResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BasicGlobalParametersResponse) GetData() []*BasicGlobalParametersData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BasicGlobalParametersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value         string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	DefaultValue  string `protobuf:"bytes,4,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Type          string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	DbType        string `protobuf:"bytes,7,opt,name=db_type,json=dbType,proto3" json:"db_type,omitempty"`
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	Group         string `protobuf:"bytes,9,opt,name=group,proto3" json:"group,omitempty"`
	GroupSort     int64  `protobuf:"varint,10,opt,name=group_sort,json=groupSort,proto3" json:"group_sort,omitempty"`
	Sort          int64  `protobuf:"varint,11,opt,name=sort,proto3" json:"sort,omitempty"`
	ValueType     string `protobuf:"bytes,12,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	ValueTemplate string `protobuf:"bytes,13,opt,name=value_template,json=valueTemplate,proto3" json:"value_template,omitempty"`
	ValueStyle    string `protobuf:"bytes,14,opt,name=value_style,json=valueStyle,proto3" json:"value_style,omitempty"`
	DisableBy     string `protobuf:"bytes,15,opt,name=disable_by,json=disableBy,proto3" json:"disable_by,omitempty"`
	HasPrivate    int64  `protobuf:"varint,16,opt,name=has_private,json=hasPrivate,proto3" json:"has_private,omitempty"`
}

func (x *BasicGlobalParametersData) Reset() {
	*x = BasicGlobalParametersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicGlobalParametersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicGlobalParametersData) ProtoMessage() {}

func (x *BasicGlobalParametersData) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicGlobalParametersData.ProtoReflect.Descriptor instead.
func (*BasicGlobalParametersData) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{2}
}

func (x *BasicGlobalParametersData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BasicGlobalParametersData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BasicGlobalParametersData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *BasicGlobalParametersData) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *BasicGlobalParametersData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BasicGlobalParametersData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BasicGlobalParametersData) GetDbType() string {
	if x != nil {
		return x.DbType
	}
	return ""
}

func (x *BasicGlobalParametersData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BasicGlobalParametersData) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *BasicGlobalParametersData) GetGroupSort() int64 {
	if x != nil {
		return x.GroupSort
	}
	return 0
}

func (x *BasicGlobalParametersData) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *BasicGlobalParametersData) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *BasicGlobalParametersData) GetValueTemplate() string {
	if x != nil {
		return x.ValueTemplate
	}
	return ""
}

func (x *BasicGlobalParametersData) GetValueStyle() string {
	if x != nil {
		return x.ValueStyle
	}
	return ""
}

func (x *BasicGlobalParametersData) GetDisableBy() string {
	if x != nil {
		return x.DisableBy
	}
	return ""
}

func (x *BasicGlobalParametersData) GetHasPrivate() int64 {
	if x != nil {
		return x.HasPrivate
	}
	return 0
}

type BasicGlobalParametersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *BasicGlobalParametersRequest) Reset() {
	*x = BasicGlobalParametersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicGlobalParametersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicGlobalParametersRequest) ProtoMessage() {}

func (x *BasicGlobalParametersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicGlobalParametersRequest.ProtoReflect.Descriptor instead.
func (*BasicGlobalParametersRequest) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{3}
}

func (x *BasicGlobalParametersRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type BasicVerifyRbacRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method   string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`                     //请求方法
	FullPath string `protobuf:"bytes,2,opt,name=full_path,json=fullPath,proto3" json:"full_path,omitempty"` //请求路由
	RoleId   int64  `protobuf:"varint,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`      //角色id
}

func (x *BasicVerifyRbacRequest) Reset() {
	*x = BasicVerifyRbacRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicVerifyRbacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicVerifyRbacRequest) ProtoMessage() {}

func (x *BasicVerifyRbacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicVerifyRbacRequest.ProtoReflect.Descriptor instead.
func (*BasicVerifyRbacRequest) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{4}
}

func (x *BasicVerifyRbacRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *BasicVerifyRbacRequest) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

func (x *BasicVerifyRbacRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type BasicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BasicResponse) Reset() {
	*x = BasicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbm_basic_basic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicResponse) ProtoMessage() {}

func (x *BasicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bbm_basic_basic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicResponse.ProtoReflect.Descriptor instead.
func (*BasicResponse) Descriptor() ([]byte, []int) {
	return file_bbm_basic_basic_proto_rawDescGZIP(), []int{5}
}

func (x *BasicResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BasicResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_bbm_basic_basic_proto protoreflect.FileDescriptor

var file_bbm_basic_basic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x62, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0x6b,
	0x0a, 0x19, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x1d,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcd, 0x03, 0x0a, 0x19, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x61,
	0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x32, 0x0a, 0x1c, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x66, 0x0a, 0x16,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x62, 0x61, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x81, 0x02, 0x0a, 0x05, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x12, 0x46, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x52, 0x62, 0x61, 0x63, 0x12, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x62, 0x61, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x62, 0x62, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x3b, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bbm_basic_basic_proto_rawDescOnce sync.Once
	file_bbm_basic_basic_proto_rawDescData = file_bbm_basic_basic_proto_rawDesc
)

func file_bbm_basic_basic_proto_rawDescGZIP() []byte {
	file_bbm_basic_basic_proto_rawDescOnce.Do(func() {
		file_bbm_basic_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_bbm_basic_basic_proto_rawDescData)
	})
	return file_bbm_basic_basic_proto_rawDescData
}

var file_bbm_basic_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bbm_basic_basic_proto_goTypes = []interface{}{
	(*BasicBehaviorTraceRequest)(nil),     // 0: basic.BasicBehaviorTraceRequest
	(*BasicGlobalParametersResponse)(nil), // 1: basic.BasicGlobalParametersResponse
	(*BasicGlobalParametersData)(nil),     // 2: basic.BasicGlobalParametersData
	(*BasicGlobalParametersRequest)(nil),  // 3: basic.BasicGlobalParametersRequest
	(*BasicVerifyRbacRequest)(nil),        // 4: basic.BasicVerifyRbacRequest
	(*BasicResponse)(nil),                 // 5: basic.BasicResponse
}
var file_bbm_basic_basic_proto_depIdxs = []int32{
	2, // 0: basic.BasicGlobalParametersResponse.data:type_name -> basic.BasicGlobalParametersData
	4, // 1: basic.Basic.BasicVerifyRbac:input_type -> basic.BasicVerifyRbacRequest
	3, // 2: basic.Basic.BasicGlobalParameters:input_type -> basic.BasicGlobalParametersRequest
	0, // 3: basic.Basic.BasicBehaviorTrace:input_type -> basic.BasicBehaviorTraceRequest
	5, // 4: basic.Basic.BasicVerifyRbac:output_type -> basic.BasicResponse
	1, // 5: basic.Basic.BasicGlobalParameters:output_type -> basic.BasicGlobalParametersResponse
	5, // 6: basic.Basic.BasicBehaviorTrace:output_type -> basic.BasicResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bbm_basic_basic_proto_init() }
func file_bbm_basic_basic_proto_init() {
	if File_bbm_basic_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bbm_basic_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicBehaviorTraceRequest); i {
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
		file_bbm_basic_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicGlobalParametersResponse); i {
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
		file_bbm_basic_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicGlobalParametersData); i {
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
		file_bbm_basic_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicGlobalParametersRequest); i {
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
		file_bbm_basic_basic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicVerifyRbacRequest); i {
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
		file_bbm_basic_basic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicResponse); i {
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
			RawDescriptor: file_bbm_basic_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bbm_basic_basic_proto_goTypes,
		DependencyIndexes: file_bbm_basic_basic_proto_depIdxs,
		MessageInfos:      file_bbm_basic_basic_proto_msgTypes,
	}.Build()
	File_bbm_basic_basic_proto = out.File
	file_bbm_basic_basic_proto_rawDesc = nil
	file_bbm_basic_basic_proto_goTypes = nil
	file_bbm_basic_basic_proto_depIdxs = nil
}
