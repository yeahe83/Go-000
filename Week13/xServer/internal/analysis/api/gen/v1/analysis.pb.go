// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: analysis.proto

package analysispb

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

type GetYoYRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetYoYRequest) Reset() {
	*x = GetYoYRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetYoYRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetYoYRequest) ProtoMessage() {}

func (x *GetYoYRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetYoYRequest.ProtoReflect.Descriptor instead.
func (*GetYoYRequest) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{0}
}

type GetYoYResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetYoYResponse) Reset() {
	*x = GetYoYResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetYoYResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetYoYResponse) ProtoMessage() {}

func (x *GetYoYResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetYoYResponse.ProtoReflect.Descriptor instead.
func (*GetYoYResponse) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{1}
}

type GetMoMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMoMRequest) Reset() {
	*x = GetMoMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoMRequest) ProtoMessage() {}

func (x *GetMoMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoMRequest.ProtoReflect.Descriptor instead.
func (*GetMoMRequest) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{2}
}

type GetMoMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMoMResponse) Reset() {
	*x = GetMoMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoMResponse) ProtoMessage() {}

func (x *GetMoMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoMResponse.ProtoReflect.Descriptor instead.
func (*GetMoMResponse) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{3}
}

var File_analysis_proto protoreflect.FileDescriptor

var file_analysis_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x0f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x59, 0x6f, 0x59, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x59, 0x6f, 0x59, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x59, 0x6f,
	0x59, 0x12, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x59, 0x6f, 0x59, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x59,
	0x6f, 0x59, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x4d, 0x12, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a,
	0x26, 0x78, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analysis_proto_rawDescOnce sync.Once
	file_analysis_proto_rawDescData = file_analysis_proto_rawDesc
)

func file_analysis_proto_rawDescGZIP() []byte {
	file_analysis_proto_rawDescOnce.Do(func() {
		file_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_analysis_proto_rawDescData)
	})
	return file_analysis_proto_rawDescData
}

var file_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_analysis_proto_goTypes = []interface{}{
	(*GetYoYRequest)(nil),  // 0: analysis.v1.GetYoYRequest
	(*GetYoYResponse)(nil), // 1: analysis.v1.GetYoYResponse
	(*GetMoMRequest)(nil),  // 2: analysis.v1.GetMoMRequest
	(*GetMoMResponse)(nil), // 3: analysis.v1.GetMoMResponse
}
var file_analysis_proto_depIdxs = []int32{
	0, // 0: analysis.v1.AnalysisService.GetYoY:input_type -> analysis.v1.GetYoYRequest
	2, // 1: analysis.v1.AnalysisService.GetMoM:input_type -> analysis.v1.GetMoMRequest
	1, // 2: analysis.v1.AnalysisService.GetYoY:output_type -> analysis.v1.GetYoYResponse
	3, // 3: analysis.v1.AnalysisService.GetMoM:output_type -> analysis.v1.GetMoMResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_analysis_proto_init() }
func file_analysis_proto_init() {
	if File_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetYoYRequest); i {
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
		file_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetYoYResponse); i {
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
		file_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoMRequest); i {
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
		file_analysis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoMResponse); i {
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
			RawDescriptor: file_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analysis_proto_goTypes,
		DependencyIndexes: file_analysis_proto_depIdxs,
		MessageInfos:      file_analysis_proto_msgTypes,
	}.Build()
	File_analysis_proto = out.File
	file_analysis_proto_rawDesc = nil
	file_analysis_proto_goTypes = nil
	file_analysis_proto_depIdxs = nil
}
