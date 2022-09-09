// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: internal/transport/grpc/service/block.proto

package proto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Block string `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_service_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_service_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_service_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Block) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Block string `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Date  string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_service_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_service_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_service_block_proto_rawDescGZIP(), []int{1}
}

func (x *BlockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlockRequest) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *BlockRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Blocks  []*Block `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_service_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_service_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_service_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BlockResponse) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type EstimateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Count   int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *EstimateResponse) Reset() {
	*x = EstimateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_service_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateResponse) ProtoMessage() {}

func (x *EstimateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_service_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateResponse.ProtoReflect.Descriptor instead.
func (*EstimateResponse) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_service_block_proto_rawDescGZIP(), []int{3}
}

func (x *EstimateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EstimateResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_internal_transport_grpc_service_block_proto protoreflect.FileDescriptor

var file_internal_transport_grpc_service_block_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x48, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x4f, 0x0a,
	0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x42,
	0x0a, 0x10, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xcb, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_transport_grpc_service_block_proto_rawDescOnce sync.Once
	file_internal_transport_grpc_service_block_proto_rawDescData = file_internal_transport_grpc_service_block_proto_rawDesc
)

func file_internal_transport_grpc_service_block_proto_rawDescGZIP() []byte {
	file_internal_transport_grpc_service_block_proto_rawDescOnce.Do(func() {
		file_internal_transport_grpc_service_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_transport_grpc_service_block_proto_rawDescData)
	})
	return file_internal_transport_grpc_service_block_proto_rawDescData
}

var file_internal_transport_grpc_service_block_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_transport_grpc_service_block_proto_goTypes = []interface{}{
	(*Block)(nil),            // 0: proto.Block
	(*BlockRequest)(nil),     // 1: proto.BlockRequest
	(*BlockResponse)(nil),    // 2: proto.BlockResponse
	(*EstimateResponse)(nil), // 3: proto.EstimateResponse
}
var file_internal_transport_grpc_service_block_proto_depIdxs = []int32{
	0, // 0: proto.BlockResponse.blocks:type_name -> proto.Block
	1, // 1: proto.BlockService.StoreBlockation:input_type -> proto.BlockRequest
	1, // 2: proto.BlockService.ShowBlockInBlock:input_type -> proto.BlockRequest
	1, // 3: proto.BlockService.Estimate:input_type -> proto.BlockRequest
	2, // 4: proto.BlockService.StoreBlockation:output_type -> proto.BlockResponse
	2, // 5: proto.BlockService.ShowBlockInBlock:output_type -> proto.BlockResponse
	3, // 6: proto.BlockService.Estimate:output_type -> proto.EstimateResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_transport_grpc_service_block_proto_init() }
func file_internal_transport_grpc_service_block_proto_init() {
	if File_internal_transport_grpc_service_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_transport_grpc_service_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_internal_transport_grpc_service_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_internal_transport_grpc_service_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_internal_transport_grpc_service_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateResponse); i {
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
			RawDescriptor: file_internal_transport_grpc_service_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_transport_grpc_service_block_proto_goTypes,
		DependencyIndexes: file_internal_transport_grpc_service_block_proto_depIdxs,
		MessageInfos:      file_internal_transport_grpc_service_block_proto_msgTypes,
	}.Build()
	File_internal_transport_grpc_service_block_proto = out.File
	file_internal_transport_grpc_service_block_proto_rawDesc = nil
	file_internal_transport_grpc_service_block_proto_goTypes = nil
	file_internal_transport_grpc_service_block_proto_depIdxs = nil
}
