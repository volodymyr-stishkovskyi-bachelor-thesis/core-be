// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/scraper.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScrapeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScrapeRequest) Reset() {
	*x = ScrapeRequest{}
	mi := &file_proto_scraper_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScrapeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapeRequest) ProtoMessage() {}

func (x *ScrapeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scraper_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapeRequest.ProtoReflect.Descriptor instead.
func (*ScrapeRequest) Descriptor() ([]byte, []int) {
	return file_proto_scraper_proto_rawDescGZIP(), []int{0}
}

func (x *ScrapeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ScrapeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScrapeResponse) Reset() {
	*x = ScrapeResponse{}
	mi := &file_proto_scraper_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScrapeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapeResponse) ProtoMessage() {}

func (x *ScrapeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scraper_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapeResponse.ProtoReflect.Descriptor instead.
func (*ScrapeResponse) Descriptor() ([]byte, []int) {
	return file_proto_scraper_proto_rawDescGZIP(), []int{1}
}

func (x *ScrapeResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ScrapeResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_proto_scraper_proto protoreflect.FileDescriptor

var file_proto_scraper_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x22, 0x21,
	0x0a, 0x0d, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0x4b, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x12,
	0x16, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x6f, 0x6c, 0x6f, 0x64, 0x79, 0x6d, 0x79, 0x72, 0x2d, 0x73, 0x74, 0x69, 0x73, 0x68, 0x6b, 0x6f,
	0x76, 0x73, 0x6b, 0x79, 0x69, 0x2d, 0x62, 0x61, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x72, 0x2d, 0x74,
	0x68, 0x65, 0x73, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x62, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_proto_scraper_proto_rawDescOnce sync.Once
	file_proto_scraper_proto_rawDescData []byte
)

func file_proto_scraper_proto_rawDescGZIP() []byte {
	file_proto_scraper_proto_rawDescOnce.Do(func() {
		file_proto_scraper_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scraper_proto_rawDesc), len(file_proto_scraper_proto_rawDesc)))
	})
	return file_proto_scraper_proto_rawDescData
}

var file_proto_scraper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_scraper_proto_goTypes = []any{
	(*ScrapeRequest)(nil),  // 0: scraper.ScrapeRequest
	(*ScrapeResponse)(nil), // 1: scraper.ScrapeResponse
}
var file_proto_scraper_proto_depIdxs = []int32{
	0, // 0: scraper.ScraperService.Scrape:input_type -> scraper.ScrapeRequest
	1, // 1: scraper.ScraperService.Scrape:output_type -> scraper.ScrapeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scraper_proto_init() }
func file_proto_scraper_proto_init() {
	if File_proto_scraper_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scraper_proto_rawDesc), len(file_proto_scraper_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scraper_proto_goTypes,
		DependencyIndexes: file_proto_scraper_proto_depIdxs,
		MessageInfos:      file_proto_scraper_proto_msgTypes,
	}.Build()
	File_proto_scraper_proto = out.File
	file_proto_scraper_proto_goTypes = nil
	file_proto_scraper_proto_depIdxs = nil
}
