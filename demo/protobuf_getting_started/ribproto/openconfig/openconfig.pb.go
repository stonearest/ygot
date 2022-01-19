// openconfig is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - yang/rib/openconfig-rib-bgp.yang
// Include paths:
//   - yang/...

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: ribproto/openconfig/openconfig.proto

package openconfig

import (
	openconfig_rib_bgp "github.com/openconfig/ygot/demo/protobuf_getting_started/ribproto/openconfig/openconfig_rib_bgp"
	_ "github.com/openconfig/ygot/proto/yext"
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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BgpRib *openconfig_rib_bgp.BgpRib `protobuf:"bytes,314942212,opt,name=bgp_rib,json=bgpRib,proto3" json:"bgp_rib,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ribproto_openconfig_openconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_ribproto_openconfig_openconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_ribproto_openconfig_openconfig_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetBgpRib() *openconfig_rib_bgp.BgpRib {
	if x != nil {
		return x.BgpRib
	}
	return nil
}

var File_ribproto_openconfig_openconfig_proto protoreflect.FileDescriptor

var file_ribproto_openconfig_openconfig_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x79, 0x67, 0x6f, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x65, 0x78, 0x74, 0x2f, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x78, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x79, 0x67, 0x6f, 0x74,
	0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x67,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x72,
	0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x69,
	0x62, 0x5f, 0x62, 0x67, 0x70, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x72, 0x69, 0x62, 0x5f, 0x62, 0x67, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59,
	0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x62, 0x67, 0x70, 0x5f,
	0x72, 0x69, 0x62, 0x18, 0x84, 0xc6, 0x96, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x62, 0x67, 0x70, 0x2e, 0x42, 0x67,
	0x70, 0x52, 0x69, 0x62, 0x42, 0x0b, 0x82, 0x41, 0x08, 0x2f, 0x62, 0x67, 0x70, 0x2d, 0x72, 0x69,
	0x62, 0x52, 0x06, 0x62, 0x67, 0x70, 0x52, 0x69, 0x62, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x79, 0x67, 0x6f, 0x74, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x67, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x72, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ribproto_openconfig_openconfig_proto_rawDescOnce sync.Once
	file_ribproto_openconfig_openconfig_proto_rawDescData = file_ribproto_openconfig_openconfig_proto_rawDesc
)

func file_ribproto_openconfig_openconfig_proto_rawDescGZIP() []byte {
	file_ribproto_openconfig_openconfig_proto_rawDescOnce.Do(func() {
		file_ribproto_openconfig_openconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_ribproto_openconfig_openconfig_proto_rawDescData)
	})
	return file_ribproto_openconfig_openconfig_proto_rawDescData
}

var file_ribproto_openconfig_openconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ribproto_openconfig_openconfig_proto_goTypes = []interface{}{
	(*Device)(nil),                    // 0: openconfig.Device
	(*openconfig_rib_bgp.BgpRib)(nil), // 1: openconfig.openconfig_rib_bgp.BgpRib
}
var file_ribproto_openconfig_openconfig_proto_depIdxs = []int32{
	1, // 0: openconfig.Device.bgp_rib:type_name -> openconfig.openconfig_rib_bgp.BgpRib
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ribproto_openconfig_openconfig_proto_init() }
func file_ribproto_openconfig_openconfig_proto_init() {
	if File_ribproto_openconfig_openconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ribproto_openconfig_openconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
			RawDescriptor: file_ribproto_openconfig_openconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ribproto_openconfig_openconfig_proto_goTypes,
		DependencyIndexes: file_ribproto_openconfig_openconfig_proto_depIdxs,
		MessageInfos:      file_ribproto_openconfig_openconfig_proto_msgTypes,
	}.Build()
	File_ribproto_openconfig_openconfig_proto = out.File
	file_ribproto_openconfig_openconfig_proto_rawDesc = nil
	file_ribproto_openconfig_openconfig_proto_goTypes = nil
	file_ribproto_openconfig_openconfig_proto_depIdxs = nil
}
