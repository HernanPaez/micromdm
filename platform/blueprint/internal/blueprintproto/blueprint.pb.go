// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: blueprint.proto

package blueprintproto

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

type Blueprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid                                string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                                string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ManifestUrls                        []string `protobuf:"bytes,3,rep,name=manifest_urls,json=manifestUrls,proto3" json:"manifest_urls,omitempty"`
	ProfileIds                          []string `protobuf:"bytes,5,rep,name=profile_ids,json=profileIds,proto3" json:"profile_ids,omitempty"`
	ApplyAt                             []string `protobuf:"bytes,6,rep,name=apply_at,json=applyAt,proto3" json:"apply_at,omitempty"`
	UserUuid                            []string `protobuf:"bytes,7,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	SkipPrimarySetupAccountCreation     bool     `protobuf:"varint,8,opt,name=skip_primary_setup_account_creation,json=skipPrimarySetupAccountCreation,proto3" json:"skip_primary_setup_account_creation,omitempty"`
	SetPrimarySetupAccountAsRegularUser bool     `protobuf:"varint,9,opt,name=set_primary_setup_account_as_regular_user,json=setPrimarySetupAccountAsRegularUser,proto3" json:"set_primary_setup_account_as_regular_user,omitempty"`
}

func (x *Blueprint) Reset() {
	*x = Blueprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blueprint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blueprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blueprint) ProtoMessage() {}

func (x *Blueprint) ProtoReflect() protoreflect.Message {
	mi := &file_blueprint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blueprint.ProtoReflect.Descriptor instead.
func (*Blueprint) Descriptor() ([]byte, []int) {
	return file_blueprint_proto_rawDescGZIP(), []int{0}
}

func (x *Blueprint) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Blueprint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Blueprint) GetManifestUrls() []string {
	if x != nil {
		return x.ManifestUrls
	}
	return nil
}

func (x *Blueprint) GetProfileIds() []string {
	if x != nil {
		return x.ProfileIds
	}
	return nil
}

func (x *Blueprint) GetApplyAt() []string {
	if x != nil {
		return x.ApplyAt
	}
	return nil
}

func (x *Blueprint) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *Blueprint) GetSkipPrimarySetupAccountCreation() bool {
	if x != nil {
		return x.SkipPrimarySetupAccountCreation
	}
	return false
}

func (x *Blueprint) GetSetPrimarySetupAccountAsRegularUser() bool {
	if x != nil {
		return x.SetPrimarySetupAccountAsRegularUser
	}
	return false
}

var File_blueprint_proto protoreflect.FileDescriptor

var file_blueprint_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xec, 0x02, 0x0a, 0x09, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x23, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1f, 0x73, 0x6b, 0x69, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x29, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x23, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x73,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x52, 0x0d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x64, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x64, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_blueprint_proto_rawDescOnce sync.Once
	file_blueprint_proto_rawDescData = file_blueprint_proto_rawDesc
)

func file_blueprint_proto_rawDescGZIP() []byte {
	file_blueprint_proto_rawDescOnce.Do(func() {
		file_blueprint_proto_rawDescData = protoimpl.X.CompressGZIP(file_blueprint_proto_rawDescData)
	})
	return file_blueprint_proto_rawDescData
}

var file_blueprint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blueprint_proto_goTypes = []interface{}{
	(*Blueprint)(nil), // 0: blueprintproto.Blueprint
}
var file_blueprint_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blueprint_proto_init() }
func file_blueprint_proto_init() {
	if File_blueprint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blueprint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blueprint); i {
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
			RawDescriptor: file_blueprint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blueprint_proto_goTypes,
		DependencyIndexes: file_blueprint_proto_depIdxs,
		MessageInfos:      file_blueprint_proto_msgTypes,
	}.Build()
	File_blueprint_proto = out.File
	file_blueprint_proto_rawDesc = nil
	file_blueprint_proto_goTypes = nil
	file_blueprint_proto_depIdxs = nil
}
