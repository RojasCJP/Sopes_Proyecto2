// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: management/management.proto

package management

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

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age         int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	VaccineType string `protobuf:"bytes,3,opt,name=vaccine_type,json=vaccineType,proto3" json:"vaccine_type,omitempty"`
	Location    string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	NDose       int32  `protobuf:"varint,5,opt,name=n_dose,json=nDose,proto3" json:"n_dose,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_management_management_proto_msgTypes[0]
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
	return file_management_management_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetVaccineType() string {
	if x != nil {
		return x.VaccineType
	}
	return ""
}

func (x *User) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *User) GetNDose() int32 {
	if x != nil {
		return x.NDose
	}
	return 0
}

var File_management_management_proto protoreflect.FileDescriptor

var file_management_management_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x63, 0x63,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x76, 0x61, 0x63, 0x63, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x5f, 0x64, 0x6f, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x44, 0x6f, 0x73, 0x65, 0x32, 0x46,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x35, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_management_proto_rawDescOnce sync.Once
	file_management_management_proto_rawDescData = file_management_management_proto_rawDesc
)

func file_management_management_proto_rawDescGZIP() []byte {
	file_management_management_proto_rawDescOnce.Do(func() {
		file_management_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_management_proto_rawDescData)
	})
	return file_management_management_proto_rawDescData
}

var file_management_management_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_management_management_proto_goTypes = []interface{}{
	(*User)(nil), // 0: management.User
}
var file_management_management_proto_depIdxs = []int32{
	0, // 0: management.UserManagment.CreateNewUser:input_type -> management.User
	0, // 1: management.UserManagment.CreateNewUser:output_type -> management.User
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_management_proto_init() }
func file_management_management_proto_init() {
	if File_management_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_management_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_management_proto_goTypes,
		DependencyIndexes: file_management_management_proto_depIdxs,
		MessageInfos:      file_management_management_proto_msgTypes,
	}.Build()
	File_management_management_proto = out.File
	file_management_management_proto_rawDesc = nil
	file_management_management_proto_goTypes = nil
	file_management_management_proto_depIdxs = nil
}
