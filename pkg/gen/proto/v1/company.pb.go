// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: company.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CompanyINN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *CompanyINN) Reset() {
	*x = CompanyINN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyINN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyINN) ProtoMessage() {}

func (x *CompanyINN) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyINN.ProtoReflect.Descriptor instead.
func (*CompanyINN) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyINN) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

type CompanyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn           string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp           string `protobuf:"bytes,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	CompanyName   string `protobuf:"bytes,3,opt,name=companyName,proto3" json:"companyName,omitempty"`
	OwnerFullName string `protobuf:"bytes,4,opt,name=ownerFullName,proto3" json:"ownerFullName,omitempty"`
}

func (x *CompanyInfo) Reset() {
	*x = CompanyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyInfo) ProtoMessage() {}

func (x *CompanyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyInfo.ProtoReflect.Descriptor instead.
func (*CompanyInfo) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyInfo) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *CompanyInfo) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *CompanyInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CompanyInfo) GetOwnerFullName() string {
	if x != nil {
		return x.OwnerFullName
	}
	return ""
}

var File_company_proto protoreflect.FileDescriptor

var file_company_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x22, 0x79, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0x66, 0x0a, 0x0a, 0x52, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x4e, 0x4e, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x69,
	0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x69, 0x6e, 0x6e, 0x7d, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x6e, 0x69, 0x75, 0x73, 0x6b, 0x61,
	0x61, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_company_proto_rawDescOnce sync.Once
	file_company_proto_rawDescData = file_company_proto_rawDesc
)

func file_company_proto_rawDescGZIP() []byte {
	file_company_proto_rawDescOnce.Do(func() {
		file_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_proto_rawDescData)
	})
	return file_company_proto_rawDescData
}

var file_company_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_company_proto_goTypes = []interface{}{
	(*CompanyINN)(nil),  // 0: company.CompanyINN
	(*CompanyInfo)(nil), // 1: company.CompanyInfo
}
var file_company_proto_depIdxs = []int32{
	0, // 0: company.RusProfile.GetCompanyInfo:input_type -> company.CompanyINN
	1, // 1: company.RusProfile.GetCompanyInfo:output_type -> company.CompanyInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_company_proto_init() }
func file_company_proto_init() {
	if File_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyINN); i {
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
		file_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyInfo); i {
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
			RawDescriptor: file_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_proto_goTypes,
		DependencyIndexes: file_company_proto_depIdxs,
		MessageInfos:      file_company_proto_msgTypes,
	}.Build()
	File_company_proto = out.File
	file_company_proto_rawDesc = nil
	file_company_proto_goTypes = nil
	file_company_proto_depIdxs = nil
}
