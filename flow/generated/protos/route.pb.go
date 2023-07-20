// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: route.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePeerFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfigs *FlowConnectionConfigs `protobuf:"bytes,1,opt,name=connection_configs,json=connectionConfigs,proto3" json:"connection_configs,omitempty"`
}

func (x *CreatePeerFlowRequest) Reset() {
	*x = CreatePeerFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePeerFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeerFlowRequest) ProtoMessage() {}

func (x *CreatePeerFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeerFlowRequest.ProtoReflect.Descriptor instead.
func (*CreatePeerFlowRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePeerFlowRequest) GetConnectionConfigs() *FlowConnectionConfigs {
	if x != nil {
		return x.ConnectionConfigs
	}
	return nil
}

type CreatePeerFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorflowId string `protobuf:"bytes,1,opt,name=worflow_id,json=worflowId,proto3" json:"worflow_id,omitempty"`
}

func (x *CreatePeerFlowResponse) Reset() {
	*x = CreatePeerFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePeerFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeerFlowResponse) ProtoMessage() {}

func (x *CreatePeerFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeerFlowResponse.ProtoReflect.Descriptor instead.
func (*CreatePeerFlowResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePeerFlowResponse) GetWorflowId() string {
	if x != nil {
		return x.WorflowId
	}
	return ""
}

type CreateQRepFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrepConfig *QRepConfig `protobuf:"bytes,1,opt,name=qrep_config,json=qrepConfig,proto3" json:"qrep_config,omitempty"`
}

func (x *CreateQRepFlowRequest) Reset() {
	*x = CreateQRepFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQRepFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQRepFlowRequest) ProtoMessage() {}

func (x *CreateQRepFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQRepFlowRequest.ProtoReflect.Descriptor instead.
func (*CreateQRepFlowRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQRepFlowRequest) GetQrepConfig() *QRepConfig {
	if x != nil {
		return x.QrepConfig
	}
	return nil
}

type CreateQRepFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorflowId string `protobuf:"bytes,1,opt,name=worflow_id,json=worflowId,proto3" json:"worflow_id,omitempty"`
}

func (x *CreateQRepFlowResponse) Reset() {
	*x = CreateQRepFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQRepFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQRepFlowResponse) ProtoMessage() {}

func (x *CreateQRepFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQRepFlowResponse.ProtoReflect.Descriptor instead.
func (*CreateQRepFlowResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{3}
}

func (x *CreateQRepFlowResponse) GetWorflowId() string {
	if x != nil {
		return x.WorflowId
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{4}
}

func (x *HealthCheckRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheckResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId      string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	FlowJobName     string `protobuf:"bytes,2,opt,name=flow_job_name,json=flowJobName,proto3" json:"flow_job_name,omitempty"`
	SourcePeer      *Peer  `protobuf:"bytes,3,opt,name=source_peer,json=sourcePeer,proto3" json:"source_peer,omitempty"`
	DestinationPeer *Peer  `protobuf:"bytes,4,opt,name=destination_peer,json=destinationPeer,proto3" json:"destination_peer,omitempty"`
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{6}
}

func (x *ShutdownRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *ShutdownRequest) GetFlowJobName() string {
	if x != nil {
		return x.FlowJobName
	}
	return ""
}

func (x *ShutdownRequest) GetSourcePeer() *Peer {
	if x != nil {
		return x.SourcePeer
	}
	return nil
}

func (x *ShutdownRequest) GetDestinationPeer() *Peer {
	if x != nil {
		return x.DestinationPeer
	}
	return nil
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{7}
}

func (x *ShutdownResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ShutdownResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_route_proto protoreflect.FileDescriptor

var file_route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70,
	0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65, 0x65,
	0x72, 0x64, 0x62, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x11,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x22, 0x37, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x71, 0x72, 0x65, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64,
	0x62, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x52, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0a, 0x71, 0x72, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x37, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xca, 0x01,
	0x0a, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x77, 0x4a,
	0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65,
	0x65, 0x72, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x10, 0x53, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xf2, 0x02, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65,
	0x72, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x65, 0x65,
	0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70,
	0x46, 0x6c, 0x6f, 0x77, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x65, 0x70, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x52, 0x65, 0x70, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x20, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x53, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7c, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x64, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x0a, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0xca, 0x02, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52, 0x6f, 0x75, 0x74, 0x65, 0xe2, 0x02,
	0x17, 0x50, 0x65, 0x65, 0x72, 0x64, 0x62, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x64,
	0x62, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_proto_rawDescOnce sync.Once
	file_route_proto_rawDescData = file_route_proto_rawDesc
)

func file_route_proto_rawDescGZIP() []byte {
	file_route_proto_rawDescOnce.Do(func() {
		file_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_proto_rawDescData)
	})
	return file_route_proto_rawDescData
}

var file_route_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_route_proto_goTypes = []interface{}{
	(*CreatePeerFlowRequest)(nil),  // 0: peerdb_route.CreatePeerFlowRequest
	(*CreatePeerFlowResponse)(nil), // 1: peerdb_route.CreatePeerFlowResponse
	(*CreateQRepFlowRequest)(nil),  // 2: peerdb_route.CreateQRepFlowRequest
	(*CreateQRepFlowResponse)(nil), // 3: peerdb_route.CreateQRepFlowResponse
	(*HealthCheckRequest)(nil),     // 4: peerdb_route.HealthCheckRequest
	(*HealthCheckResponse)(nil),    // 5: peerdb_route.HealthCheckResponse
	(*ShutdownRequest)(nil),        // 6: peerdb_route.ShutdownRequest
	(*ShutdownResponse)(nil),       // 7: peerdb_route.ShutdownResponse
	(*FlowConnectionConfigs)(nil),  // 8: peerdb_flow.FlowConnectionConfigs
	(*QRepConfig)(nil),             // 9: peerdb_flow.QRepConfig
	(*Peer)(nil),                   // 10: peerdb_peers.Peer
}
var file_route_proto_depIdxs = []int32{
	8,  // 0: peerdb_route.CreatePeerFlowRequest.connection_configs:type_name -> peerdb_flow.FlowConnectionConfigs
	9,  // 1: peerdb_route.CreateQRepFlowRequest.qrep_config:type_name -> peerdb_flow.QRepConfig
	10, // 2: peerdb_route.ShutdownRequest.source_peer:type_name -> peerdb_peers.Peer
	10, // 3: peerdb_route.ShutdownRequest.destination_peer:type_name -> peerdb_peers.Peer
	0,  // 4: peerdb_route.FlowService.CreatePeerFlow:input_type -> peerdb_route.CreatePeerFlowRequest
	2,  // 5: peerdb_route.FlowService.CreateQRepFlow:input_type -> peerdb_route.CreateQRepFlowRequest
	4,  // 6: peerdb_route.FlowService.HealthCheck:input_type -> peerdb_route.HealthCheckRequest
	6,  // 7: peerdb_route.FlowService.ShutdownFlow:input_type -> peerdb_route.ShutdownRequest
	1,  // 8: peerdb_route.FlowService.CreatePeerFlow:output_type -> peerdb_route.CreatePeerFlowResponse
	3,  // 9: peerdb_route.FlowService.CreateQRepFlow:output_type -> peerdb_route.CreateQRepFlowResponse
	5,  // 10: peerdb_route.FlowService.HealthCheck:output_type -> peerdb_route.HealthCheckResponse
	7,  // 11: peerdb_route.FlowService.ShutdownFlow:output_type -> peerdb_route.ShutdownResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_route_proto_init() }
func file_route_proto_init() {
	if File_route_proto != nil {
		return
	}
	file_peers_proto_init()
	file_flow_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePeerFlowRequest); i {
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
		file_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePeerFlowResponse); i {
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
		file_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQRepFlowRequest); i {
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
		file_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQRepFlowResponse); i {
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
		file_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
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
			RawDescriptor: file_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_proto_goTypes,
		DependencyIndexes: file_route_proto_depIdxs,
		MessageInfos:      file_route_proto_msgTypes,
	}.Build()
	File_route_proto = out.File
	file_route_proto_rawDesc = nil
	file_route_proto_goTypes = nil
	file_route_proto_depIdxs = nil
}