// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: serverk8s.proto

package kubernetes

import (
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

var File_serverk8s_proto protoreflect.FileDescriptor

var file_serverk8s_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6b, 0x38, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x70, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea,
	0x04, 0x0a, 0x0a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a,
	0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0d,
	0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x0a, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0d,
	0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x12, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x76, 0x63, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x70, 0x76, 0x63, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x38, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x0d,
	0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x0c, 0x2e,
	0x76, 0x31, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x76, 0x63, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x77, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_serverk8s_proto_goTypes = []interface{}{
	(*Namespace)(nil),        // 0: v1.namespace
	(*Empty)(nil),            // 1: v1.empty
	(*Service)(nil),          // 2: v1.service
	(*DeleteService)(nil),    // 3: v1.delete_service
	(*Deployment)(nil),       // 4: v1.deployment
	(*DeleteDeployment)(nil), // 5: v1.delete_deployment
	(*Pvc)(nil),              // 6: v1.pvc
	(*DeletePvc)(nil),        // 7: v1.delete_pvc
	(*Result)(nil),           // 8: v1.result
	(*ListNamespace)(nil),    // 9: v1.list_namespace
	(*ListService)(nil),      // 10: v1.list_service
	(*ListDeployment)(nil),   // 11: v1.list_deployment
	(*ListPvc)(nil),          // 12: v1.list_pvc
}
var file_serverk8s_proto_depIdxs = []int32{
	0,  // 0: v1.Kubernetes.CreateNamespace:input_type -> v1.namespace
	0,  // 1: v1.Kubernetes.DeleteNamespace:input_type -> v1.namespace
	1,  // 2: v1.Kubernetes.ListNamespace:input_type -> v1.empty
	2,  // 3: v1.Kubernetes.CreateService:input_type -> v1.service
	3,  // 4: v1.Kubernetes.DeleteService:input_type -> v1.delete_service
	0,  // 5: v1.Kubernetes.ListService:input_type -> v1.namespace
	4,  // 6: v1.Kubernetes.CreateDeployment:input_type -> v1.deployment
	5,  // 7: v1.Kubernetes.DeleteDeployment:input_type -> v1.delete_deployment
	0,  // 8: v1.Kubernetes.ListDeployment:input_type -> v1.namespace
	6,  // 9: v1.Kubernetes.CreatePersistentVolumeClaim:input_type -> v1.pvc
	7,  // 10: v1.Kubernetes.DeletePersistentVolumeClaim:input_type -> v1.delete_pvc
	0,  // 11: v1.Kubernetes.ListPersistentVolumeClaim:input_type -> v1.namespace
	8,  // 12: v1.Kubernetes.CreateNamespace:output_type -> v1.result
	8,  // 13: v1.Kubernetes.DeleteNamespace:output_type -> v1.result
	9,  // 14: v1.Kubernetes.ListNamespace:output_type -> v1.list_namespace
	2,  // 15: v1.Kubernetes.CreateService:output_type -> v1.service
	8,  // 16: v1.Kubernetes.DeleteService:output_type -> v1.result
	10, // 17: v1.Kubernetes.ListService:output_type -> v1.list_service
	8,  // 18: v1.Kubernetes.CreateDeployment:output_type -> v1.result
	8,  // 19: v1.Kubernetes.DeleteDeployment:output_type -> v1.result
	11, // 20: v1.Kubernetes.ListDeployment:output_type -> v1.list_deployment
	8,  // 21: v1.Kubernetes.CreatePersistentVolumeClaim:output_type -> v1.result
	8,  // 22: v1.Kubernetes.DeletePersistentVolumeClaim:output_type -> v1.result
	12, // 23: v1.Kubernetes.ListPersistentVolumeClaim:output_type -> v1.list_pvc
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_serverk8s_proto_init() }
func file_serverk8s_proto_init() {
	if File_serverk8s_proto != nil {
		return
	}
	file_deployment_proto_init()
	file_pvc_proto_init()
	file_service_proto_init()
	file_namespace_proto_init()
	file_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_serverk8s_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serverk8s_proto_goTypes,
		DependencyIndexes: file_serverk8s_proto_depIdxs,
	}.Build()
	File_serverk8s_proto = out.File
	file_serverk8s_proto_rawDesc = nil
	file_serverk8s_proto_goTypes = nil
	file_serverk8s_proto_depIdxs = nil
}
