// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: user.proto

package userproto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Udid          string `protobuf:"bytes,2,opt,name=udid,proto3" json:"udid,omitempty"`
	UserId        string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserShortname string `protobuf:"bytes,4,opt,name=user_shortname,json=userShortname,proto3" json:"user_shortname,omitempty"`
	UserLongname  string `protobuf:"bytes,5,opt,name=user_longname,json=userLongname,proto3" json:"user_longname,omitempty"`
	AuthToken     string `protobuf:"bytes,6,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	PasswordHash  []byte `protobuf:"bytes,7,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Hidden        bool   `protobuf:"varint,8,opt,name=hidden,proto3" json:"hidden,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUserShortname() string {
	if x != nil {
		return x.UserShortname
	}
	return ""
}

func (x *User) GetUserLongname() string {
	if x != nil {
		return x.UserLongname
	}
	return ""
}

func (x *User) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *User) GetPasswordHash() []byte {
	if x != nil {
		return x.PasswordHash
	}
	return nil
}

func (x *User) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x64, 0x6d,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x64, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil), // 0: userproto.User
}
var file_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
