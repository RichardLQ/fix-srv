// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: service.proto

package stub

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x02, 0x0a, 0x0a, 0x46, 0x69, 0x78, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x69, 0x63, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x63, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x63,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x52, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x70, 0x69, 0x63, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x69, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x31, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x31, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x74,
	0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x31, 0x52, 0x73,
	0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x70, 0x69, 0x63, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x74, 0x75, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*GetPicCategory1Req)(nil), // 0: stub.GetPicCategory1Req
	(*GetPicList1Req)(nil),     // 1: stub.GetPicList1Req
	(*GetBannerReq)(nil),       // 2: stub.GetBannerReq
	(*GetPicCategory1Rsp)(nil), // 3: stub.GetPicCategory1Rsp
	(*GetPicList1Rsp)(nil),     // 4: stub.GetPicList1Rsp
	(*GetBannerRsp)(nil),       // 5: stub.GetBannerRsp
}
var file_service_proto_depIdxs = []int32{
	0, // 0: stub.FixService.GetPicCategory1:input_type -> stub.GetPicCategory1Req
	1, // 1: stub.FixService.GetPicList1:input_type -> stub.GetPicList1Req
	2, // 2: stub.FixService.GetBanner:input_type -> stub.GetBannerReq
	3, // 3: stub.FixService.GetPicCategory1:output_type -> stub.GetPicCategory1Rsp
	4, // 4: stub.FixService.GetPicList1:output_type -> stub.GetPicList1Rsp
	5, // 5: stub.FixService.GetBanner:output_type -> stub.GetBannerRsp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_pics_struct_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
