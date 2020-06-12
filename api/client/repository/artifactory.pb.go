// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: repository/artifactory.proto

package repository

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

// Artifactory repository integration.
type Artifactory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the Artifactory integration is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Artifactory service search configuration.
	Searches []*Artifactory_Search `protobuf:"bytes,2,rep,name=searches,proto3" json:"searches,omitempty"`
}

func (x *Artifactory) Reset() {
	*x = Artifactory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_artifactory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifactory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifactory) ProtoMessage() {}

func (x *Artifactory) ProtoReflect() protoreflect.Message {
	mi := &file_repository_artifactory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifactory.ProtoReflect.Descriptor instead.
func (*Artifactory) Descriptor() ([]byte, []int) {
	return file_repository_artifactory_proto_rawDescGZIP(), []int{0}
}

func (x *Artifactory) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Artifactory) GetSearches() []*Artifactory_Search {
	if x != nil {
		return x.Searches
	}
	return nil
}

// Artifactory service search configuration.
type Artifactory_Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the search.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The base URL at which your Artifactory search is reachable.
	BaseUrl string `protobuf:"bytes,2,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	// The repo in your Artifactory to be searched.
	Repo string `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	// The group ID in your Artifactory to be searched.
	GroupId string `protobuf:"bytes,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	// The package type of repo in your Artifactory to be searched. Defaults to
	// `MAVEN`.
	RepoType string `protobuf:"bytes,5,opt,name=repoType,proto3" json:"repoType,omitempty"`
	// The username of the Artifactory user to authenticate as.
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	// The password of the Artifactory user to authenticate as.
	Password string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Artifactory_Search) Reset() {
	*x = Artifactory_Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_artifactory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifactory_Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifactory_Search) ProtoMessage() {}

func (x *Artifactory_Search) ProtoReflect() protoreflect.Message {
	mi := &file_repository_artifactory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifactory_Search.ProtoReflect.Descriptor instead.
func (*Artifactory_Search) Descriptor() ([]byte, []int) {
	return file_repository_artifactory_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Artifactory_Search) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artifactory_Search) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *Artifactory_Search) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Artifactory_Search) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Artifactory_Search) GetRepoType() string {
	if x != nil {
		return x.RepoType
	}
	return ""
}

func (x *Artifactory_Search) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Artifactory_Search) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_repository_artifactory_proto protoreflect.FileDescriptor

var file_repository_artifactory_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0xa4, 0x02, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x1a, 0xb8, 0x01, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_repository_artifactory_proto_rawDescOnce sync.Once
	file_repository_artifactory_proto_rawDescData = file_repository_artifactory_proto_rawDesc
)

func file_repository_artifactory_proto_rawDescGZIP() []byte {
	file_repository_artifactory_proto_rawDescOnce.Do(func() {
		file_repository_artifactory_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_artifactory_proto_rawDescData)
	})
	return file_repository_artifactory_proto_rawDescData
}

var file_repository_artifactory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_repository_artifactory_proto_goTypes = []interface{}{
	(*Artifactory)(nil),        // 0: proto.repository.Artifactory
	(*Artifactory_Search)(nil), // 1: proto.repository.Artifactory.Search
}
var file_repository_artifactory_proto_depIdxs = []int32{
	1, // 0: proto.repository.Artifactory.searches:type_name -> proto.repository.Artifactory.Search
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_repository_artifactory_proto_init() }
func file_repository_artifactory_proto_init() {
	if File_repository_artifactory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repository_artifactory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifactory); i {
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
		file_repository_artifactory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifactory_Search); i {
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
			RawDescriptor: file_repository_artifactory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_repository_artifactory_proto_goTypes,
		DependencyIndexes: file_repository_artifactory_proto_depIdxs,
		MessageInfos:      file_repository_artifactory_proto_msgTypes,
	}.Build()
	File_repository_artifactory_proto = out.File
	file_repository_artifactory_proto_rawDesc = nil
	file_repository_artifactory_proto_goTypes = nil
	file_repository_artifactory_proto_depIdxs = nil
}