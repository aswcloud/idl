// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kubernetes.proto

package v1

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

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid       *Uuid         `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OwnerUuid  []*Uuid       `protobuf:"bytes,3,rep,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	Deployment []*Deployment `protobuf:"bytes,4,rep,name=deployment,proto3" json:"deployment,omitempty"`
	Service    []*Service    `protobuf:"bytes,5,rep,name=service,proto3" json:"service,omitempty"`
	Role       []*Role       `protobuf:"bytes,6,rep,name=role,proto3" json:"role,omitempty"`
	Storage    []*Storage    `protobuf:"bytes,7,rep,name=storage,proto3" json:"storage,omitempty"`
	Pod        []*Pod        `protobuf:"bytes,8,rep,name=pod,proto3" json:"pod,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Namespace) GetUuid() *Uuid {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *Namespace) GetOwnerUuid() []*Uuid {
	if x != nil {
		return x.OwnerUuid
	}
	return nil
}

func (x *Namespace) GetDeployment() []*Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *Namespace) GetService() []*Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Namespace) GetRole() []*Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *Namespace) GetStorage() []*Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *Namespace) GetPod() []*Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod []*Pod `protobuf:"bytes,8,rep,name=pod,proto3" json:"pod,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{1}
}

func (x *Deployment) GetPod() []*Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{2}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{3}
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{4}
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_kubernetes_proto_rawDescGZIP(), []int{5}
}

var File_kubernetes_proto protoreflect.FileDescriptor

var file_kubernetes_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x75, 0x69, 0x64,
	0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x25, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03, 0x70,
	0x6f, 0x64, 0x22, 0x27, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x09,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x05, 0x0a, 0x03, 0x50, 0x6f, 0x64,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubernetes_proto_rawDescOnce sync.Once
	file_kubernetes_proto_rawDescData = file_kubernetes_proto_rawDesc
)

func file_kubernetes_proto_rawDescGZIP() []byte {
	file_kubernetes_proto_rawDescOnce.Do(func() {
		file_kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubernetes_proto_rawDescData)
	})
	return file_kubernetes_proto_rawDescData
}

var file_kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kubernetes_proto_goTypes = []interface{}{
	(*Namespace)(nil),  // 0: v1.Namespace
	(*Deployment)(nil), // 1: v1.Deployment
	(*Service)(nil),    // 2: v1.Service
	(*Role)(nil),       // 3: v1.Role
	(*Storage)(nil),    // 4: v1.Storage
	(*Pod)(nil),        // 5: v1.Pod
	(*Uuid)(nil),       // 6: v1.Uuid
}
var file_kubernetes_proto_depIdxs = []int32{
	6, // 0: v1.Namespace.uuid:type_name -> v1.Uuid
	6, // 1: v1.Namespace.owner_uuid:type_name -> v1.Uuid
	1, // 2: v1.Namespace.deployment:type_name -> v1.Deployment
	2, // 3: v1.Namespace.service:type_name -> v1.Service
	3, // 4: v1.Namespace.role:type_name -> v1.Role
	4, // 5: v1.Namespace.storage:type_name -> v1.Storage
	5, // 6: v1.Namespace.pod:type_name -> v1.Pod
	5, // 7: v1.Deployment.pod:type_name -> v1.Pod
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_kubernetes_proto_init() }
func file_kubernetes_proto_init() {
	if File_kubernetes_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_kubernetes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_kubernetes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_kubernetes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
		file_kubernetes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
			RawDescriptor: file_kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kubernetes_proto_goTypes,
		DependencyIndexes: file_kubernetes_proto_depIdxs,
		MessageInfos:      file_kubernetes_proto_msgTypes,
	}.Build()
	File_kubernetes_proto = out.File
	file_kubernetes_proto_rawDesc = nil
	file_kubernetes_proto_goTypes = nil
	file_kubernetes_proto_depIdxs = nil
}
