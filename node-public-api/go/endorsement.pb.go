// Copyright (c) 2023 MASSA LABS <info@massa.net>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: endorsement.proto

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

// An endorsement, as sent in the network
type Endorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Slot in which the endorsement can be included
	Slot *Slot `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
	// Endorsement index inside the including block
	Index uint32 `protobuf:"fixed32,2,opt,name=index,proto3" json:"index,omitempty"`
	// Hash of endorsed block
	// This is the parent in thread `self.slot.thread` of the block in which the endorsement is included
	EndorsedBlock string `protobuf:"bytes,3,opt,name=endorsed_block,json=endorsedBlock,proto3" json:"endorsed_block,omitempty"`
}

func (x *Endorsement) Reset() {
	*x = Endorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endorsement) ProtoMessage() {}

func (x *Endorsement) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endorsement.ProtoReflect.Descriptor instead.
func (*Endorsement) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{0}
}

func (x *Endorsement) GetSlot() *Slot {
	if x != nil {
		return x.Slot
	}
	return nil
}

func (x *Endorsement) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Endorsement) GetEndorsedBlock() string {
	if x != nil {
		return x.EndorsedBlock
	}
	return ""
}

// Signed endorsement
type SignedEndorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endorsement
	Content *Endorsement `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// A cryptographically generated value using `serialized_data` and a public key.
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The public-key component used in the generation of the signature
	ContentCreatorPubKey string `protobuf:"bytes,3,opt,name=content_creator_pub_key,json=contentCreatorPubKey,proto3" json:"content_creator_pub_key,omitempty"`
	// Derived from the same public key used to generate the signature
	ContentCreatorAddress string `protobuf:"bytes,4,opt,name=content_creator_address,json=contentCreatorAddress,proto3" json:"content_creator_address,omitempty"`
	// A secure hash of the data. See also [massa_hash::Hash]
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SignedEndorsement) Reset() {
	*x = SignedEndorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedEndorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedEndorsement) ProtoMessage() {}

func (x *SignedEndorsement) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedEndorsement.ProtoReflect.Descriptor instead.
func (*SignedEndorsement) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{1}
}

func (x *SignedEndorsement) GetContent() *Endorsement {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SignedEndorsement) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SignedEndorsement) GetContentCreatorPubKey() string {
	if x != nil {
		return x.ContentCreatorPubKey
	}
	return ""
}

func (x *SignedEndorsement) GetContentCreatorAddress() string {
	if x != nil {
		return x.ContentCreatorAddress
	}
	return ""
}

func (x *SignedEndorsement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_endorsement_proto protoreflect.FileDescriptor

var file_endorsement_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a,
	0x0b, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x73,
	0x73, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x22, 0xe5, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x6f,
	0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x75,
	0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0xa3, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x10,
	0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x73, 0x73, 0x61, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02,
	0x0c, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xba, 0x02, 0x04,
	0x47, 0x52, 0x50, 0x43, 0xca, 0x02, 0x0c, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_endorsement_proto_rawDescOnce sync.Once
	file_endorsement_proto_rawDescData = file_endorsement_proto_rawDesc
)

func file_endorsement_proto_rawDescGZIP() []byte {
	file_endorsement_proto_rawDescOnce.Do(func() {
		file_endorsement_proto_rawDescData = protoimpl.X.CompressGZIP(file_endorsement_proto_rawDescData)
	})
	return file_endorsement_proto_rawDescData
}

var file_endorsement_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_endorsement_proto_goTypes = []interface{}{
	(*Endorsement)(nil),       // 0: massa.api.v1.Endorsement
	(*SignedEndorsement)(nil), // 1: massa.api.v1.SignedEndorsement
	(*Slot)(nil),              // 2: massa.api.v1.Slot
}
var file_endorsement_proto_depIdxs = []int32{
	2, // 0: massa.api.v1.Endorsement.slot:type_name -> massa.api.v1.Slot
	0, // 1: massa.api.v1.SignedEndorsement.content:type_name -> massa.api.v1.Endorsement
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_endorsement_proto_init() }
func file_endorsement_proto_init() {
	if File_endorsement_proto != nil {
		return
	}
	file_slot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_endorsement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endorsement); i {
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
		file_endorsement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedEndorsement); i {
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
			RawDescriptor: file_endorsement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_endorsement_proto_goTypes,
		DependencyIndexes: file_endorsement_proto_depIdxs,
		MessageInfos:      file_endorsement_proto_msgTypes,
	}.Build()
	File_endorsement_proto = out.File
	file_endorsement_proto_rawDesc = nil
	file_endorsement_proto_goTypes = nil
	file_endorsement_proto_depIdxs = nil
}
