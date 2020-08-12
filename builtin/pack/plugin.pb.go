// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: waypoint/builtin/pack/plugin.proto

package pack

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// DockerImage is the artifact type for the builder.
type DockerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image       string            `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Tag         string            `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	BuildLabels map[string]string `protobuf:"bytes,3,rep,name=build_labels,json=buildLabels,proto3" json:"build_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DockerImage) Reset() {
	*x = DockerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_pack_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerImage) ProtoMessage() {}

func (x *DockerImage) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_pack_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerImage.ProtoReflect.Descriptor instead.
func (*DockerImage) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_pack_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *DockerImage) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DockerImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *DockerImage) GetBuildLabels() map[string]string {
	if x != nil {
		return x.BuildLabels
	}
	return nil
}

var File_waypoint_builtin_pack_plugin_proto protoreflect.FileDescriptor

var file_waypoint_builtin_pack_plugin_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x69, 0x6e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x22, 0xbc, 0x01, 0x0a, 0x0b, 0x44,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x45, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x2e,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x77, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x2f, 0x70, 0x61,
	0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waypoint_builtin_pack_plugin_proto_rawDescOnce sync.Once
	file_waypoint_builtin_pack_plugin_proto_rawDescData = file_waypoint_builtin_pack_plugin_proto_rawDesc
)

func file_waypoint_builtin_pack_plugin_proto_rawDescGZIP() []byte {
	file_waypoint_builtin_pack_plugin_proto_rawDescOnce.Do(func() {
		file_waypoint_builtin_pack_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_waypoint_builtin_pack_plugin_proto_rawDescData)
	})
	return file_waypoint_builtin_pack_plugin_proto_rawDescData
}

var file_waypoint_builtin_pack_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waypoint_builtin_pack_plugin_proto_goTypes = []interface{}{
	(*DockerImage)(nil), // 0: pack.DockerImage
	nil,                 // 1: pack.DockerImage.BuildLabelsEntry
}
var file_waypoint_builtin_pack_plugin_proto_depIdxs = []int32{
	1, // 0: pack.DockerImage.build_labels:type_name -> pack.DockerImage.BuildLabelsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_waypoint_builtin_pack_plugin_proto_init() }
func file_waypoint_builtin_pack_plugin_proto_init() {
	if File_waypoint_builtin_pack_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waypoint_builtin_pack_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerImage); i {
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
			RawDescriptor: file_waypoint_builtin_pack_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waypoint_builtin_pack_plugin_proto_goTypes,
		DependencyIndexes: file_waypoint_builtin_pack_plugin_proto_depIdxs,
		MessageInfos:      file_waypoint_builtin_pack_plugin_proto_msgTypes,
	}.Build()
	File_waypoint_builtin_pack_plugin_proto = out.File
	file_waypoint_builtin_pack_plugin_proto_rawDesc = nil
	file_waypoint_builtin_pack_plugin_proto_goTypes = nil
	file_waypoint_builtin_pack_plugin_proto_depIdxs = nil
}
