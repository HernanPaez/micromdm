// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: device_command.proto

package devicecommandproto

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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid           string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Payload        []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	CreatedAt      int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastSentAt     int64  `protobuf:"varint,4,opt,name=last_sent_at,json=lastSentAt,proto3" json:"last_sent_at,omitempty"`
	Acknowledged   int64  `protobuf:"varint,5,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	TimesSent      int64  `protobuf:"varint,6,opt,name=times_sent,json=timesSent,proto3" json:"times_sent,omitempty"`
	LastStatus     string `protobuf:"bytes,7,opt,name=last_status,json=lastStatus,proto3" json:"last_status,omitempty"`
	FailureMessage []byte `protobuf:"bytes,8,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_device_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_device_command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Command) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Command) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Command) GetLastSentAt() int64 {
	if x != nil {
		return x.LastSentAt
	}
	return 0
}

func (x *Command) GetAcknowledged() int64 {
	if x != nil {
		return x.Acknowledged
	}
	return 0
}

func (x *Command) GetTimesSent() int64 {
	if x != nil {
		return x.TimesSent
	}
	return 0
}

func (x *Command) GetLastStatus() string {
	if x != nil {
		return x.LastStatus
	}
	return ""
}

func (x *Command) GetFailureMessage() []byte {
	if x != nil {
		return x.FailureMessage
	}
	return nil
}

type DeviceCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUdid string     `protobuf:"bytes,1,opt,name=device_udid,json=deviceUdid,proto3" json:"device_udid,omitempty"`
	Commands   []*Command `protobuf:"bytes,2,rep,name=commands,proto3" json:"commands,omitempty"`
	Completed  []*Command `protobuf:"bytes,3,rep,name=completed,proto3" json:"completed,omitempty"`
	Failed     []*Command `protobuf:"bytes,4,rep,name=failed,proto3" json:"failed,omitempty"`
	NotNow     []*Command `protobuf:"bytes,5,rep,name=not_now,json=notNow,proto3" json:"not_now,omitempty"`
}

func (x *DeviceCommand) Reset() {
	*x = DeviceCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCommand) ProtoMessage() {}

func (x *DeviceCommand) ProtoReflect() protoreflect.Message {
	mi := &file_device_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCommand.ProtoReflect.Descriptor instead.
func (*DeviceCommand) Descriptor() ([]byte, []int) {
	return file_device_command_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCommand) GetDeviceUdid() string {
	if x != nil {
		return x.DeviceUdid
	}
	return ""
}

func (x *DeviceCommand) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *DeviceCommand) GetCompleted() []*Command {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *DeviceCommand) GetFailed() []*Command {
	if x != nil {
		return x.Failed
	}
	return nil
}

func (x *DeviceCommand) GetNotNow() []*Command {
	if x != nil {
		return x.NotNow
	}
	return nil
}

var File_device_command_proto protoreflect.FileDescriptor

var file_device_command_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75,
	0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x55, 0x64, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x39,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x6e, 0x6f,
	0x74, 0x4e, 0x6f, 0x77, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x64, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x6d, 0x64, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_command_proto_rawDescOnce sync.Once
	file_device_command_proto_rawDescData = file_device_command_proto_rawDesc
)

func file_device_command_proto_rawDescGZIP() []byte {
	file_device_command_proto_rawDescOnce.Do(func() {
		file_device_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_command_proto_rawDescData)
	})
	return file_device_command_proto_rawDescData
}

var file_device_command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_device_command_proto_goTypes = []interface{}{
	(*Command)(nil),       // 0: devicecommandproto.Command
	(*DeviceCommand)(nil), // 1: devicecommandproto.DeviceCommand
}
var file_device_command_proto_depIdxs = []int32{
	0, // 0: devicecommandproto.DeviceCommand.commands:type_name -> devicecommandproto.Command
	0, // 1: devicecommandproto.DeviceCommand.completed:type_name -> devicecommandproto.Command
	0, // 2: devicecommandproto.DeviceCommand.failed:type_name -> devicecommandproto.Command
	0, // 3: devicecommandproto.DeviceCommand.not_now:type_name -> devicecommandproto.Command
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_device_command_proto_init() }
func file_device_command_proto_init() {
	if File_device_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_device_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCommand); i {
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
			RawDescriptor: file_device_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_device_command_proto_goTypes,
		DependencyIndexes: file_device_command_proto_depIdxs,
		MessageInfos:      file_device_command_proto_msgTypes,
	}.Build()
	File_device_command_proto = out.File
	file_device_command_proto_rawDesc = nil
	file_device_command_proto_goTypes = nil
	file_device_command_proto_depIdxs = nil
}
