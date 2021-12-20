// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/service/status/v2/csds.proto

package envoy_service_status_v2

import (
	context "context"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/admin/v2alpha"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Status of a config.
type ConfigStatus int32

const (
	// Status info is not available/unknown.
	ConfigStatus_UNKNOWN ConfigStatus = 0
	// Management server has sent the config to client and received ACK.
	ConfigStatus_SYNCED ConfigStatus = 1
	// Config is not sent.
	ConfigStatus_NOT_SENT ConfigStatus = 2
	// Management server has sent the config to client but hasn’t received
	// ACK/NACK.
	ConfigStatus_STALE ConfigStatus = 3
	// Management server has sent the config to client but received NACK.
	ConfigStatus_ERROR ConfigStatus = 4
)

// Enum value maps for ConfigStatus.
var (
	ConfigStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SYNCED",
		2: "NOT_SENT",
		3: "STALE",
		4: "ERROR",
	}
	ConfigStatus_value = map[string]int32{
		"UNKNOWN":  0,
		"SYNCED":   1,
		"NOT_SENT": 2,
		"STALE":    3,
		"ERROR":    4,
	}
)

func (x ConfigStatus) Enum() *ConfigStatus {
	p := new(ConfigStatus)
	*p = x
	return p
}

func (x ConfigStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_service_status_v2_csds_proto_enumTypes[0].Descriptor()
}

func (ConfigStatus) Type() protoreflect.EnumType {
	return &file_envoy_service_status_v2_csds_proto_enumTypes[0]
}

func (x ConfigStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigStatus.Descriptor instead.
func (ConfigStatus) EnumDescriptor() ([]byte, []int) {
	return file_envoy_service_status_v2_csds_proto_rawDescGZIP(), []int{0}
}

// Request for client status of clients identified by a list of NodeMatchers.
type ClientStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Management server can use these match criteria to identify clients.
	// The match follows OR semantics.
	NodeMatchers []*matcher.NodeMatcher `protobuf:"bytes,1,rep,name=node_matchers,json=nodeMatchers,proto3" json:"node_matchers,omitempty"`
}

func (x *ClientStatusRequest) Reset() {
	*x = ClientStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_status_v2_csds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStatusRequest) ProtoMessage() {}

func (x *ClientStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_status_v2_csds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStatusRequest.ProtoReflect.Descriptor instead.
func (*ClientStatusRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_status_v2_csds_proto_rawDescGZIP(), []int{0}
}

func (x *ClientStatusRequest) GetNodeMatchers() []*matcher.NodeMatcher {
	if x != nil {
		return x.NodeMatchers
	}
	return nil
}

// Detailed config (per xDS) with status.
// [#next-free-field: 6]
type PerXdsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ConfigStatus `protobuf:"varint,1,opt,name=status,proto3,enum=envoy.service.status.v2.ConfigStatus" json:"status,omitempty"`
	// Types that are assignable to PerXdsConfig:
	//	*PerXdsConfig_ListenerConfig
	//	*PerXdsConfig_ClusterConfig
	//	*PerXdsConfig_RouteConfig
	//	*PerXdsConfig_ScopedRouteConfig
	PerXdsConfig isPerXdsConfig_PerXdsConfig `protobuf_oneof:"per_xds_config"`
}

func (x *PerXdsConfig) Reset() {
	*x = PerXdsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_status_v2_csds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerXdsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerXdsConfig) ProtoMessage() {}

func (x *PerXdsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_status_v2_csds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerXdsConfig.ProtoReflect.Descriptor instead.
func (*PerXdsConfig) Descriptor() ([]byte, []int) {
	return file_envoy_service_status_v2_csds_proto_rawDescGZIP(), []int{1}
}

func (x *PerXdsConfig) GetStatus() ConfigStatus {
	if x != nil {
		return x.Status
	}
	return ConfigStatus_UNKNOWN
}

func (m *PerXdsConfig) GetPerXdsConfig() isPerXdsConfig_PerXdsConfig {
	if m != nil {
		return m.PerXdsConfig
	}
	return nil
}

func (x *PerXdsConfig) GetListenerConfig() *v2alpha.ListenersConfigDump {
	if x, ok := x.GetPerXdsConfig().(*PerXdsConfig_ListenerConfig); ok {
		return x.ListenerConfig
	}
	return nil
}

func (x *PerXdsConfig) GetClusterConfig() *v2alpha.ClustersConfigDump {
	if x, ok := x.GetPerXdsConfig().(*PerXdsConfig_ClusterConfig); ok {
		return x.ClusterConfig
	}
	return nil
}

func (x *PerXdsConfig) GetRouteConfig() *v2alpha.RoutesConfigDump {
	if x, ok := x.GetPerXdsConfig().(*PerXdsConfig_RouteConfig); ok {
		return x.RouteConfig
	}
	return nil
}

func (x *PerXdsConfig) GetScopedRouteConfig() *v2alpha.ScopedRoutesConfigDump {
	if x, ok := x.GetPerXdsConfig().(*PerXdsConfig_ScopedRouteConfig); ok {
		return x.ScopedRouteConfig
	}
	return nil
}

type isPerXdsConfig_PerXdsConfig interface {
	isPerXdsConfig_PerXdsConfig()
}

type PerXdsConfig_ListenerConfig struct {
	ListenerConfig *v2alpha.ListenersConfigDump `protobuf:"bytes,2,opt,name=listener_config,json=listenerConfig,proto3,oneof"`
}

type PerXdsConfig_ClusterConfig struct {
	ClusterConfig *v2alpha.ClustersConfigDump `protobuf:"bytes,3,opt,name=cluster_config,json=clusterConfig,proto3,oneof"`
}

type PerXdsConfig_RouteConfig struct {
	RouteConfig *v2alpha.RoutesConfigDump `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3,oneof"`
}

type PerXdsConfig_ScopedRouteConfig struct {
	ScopedRouteConfig *v2alpha.ScopedRoutesConfigDump `protobuf:"bytes,5,opt,name=scoped_route_config,json=scopedRouteConfig,proto3,oneof"`
}

func (*PerXdsConfig_ListenerConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_ClusterConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_RouteConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_ScopedRouteConfig) isPerXdsConfig_PerXdsConfig() {}

// All xds configs for a particular client.
type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node for a particular client.
	Node      *core.Node      `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XdsConfig []*PerXdsConfig `protobuf:"bytes,2,rep,name=xds_config,json=xdsConfig,proto3" json:"xds_config,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_status_v2_csds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_status_v2_csds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_envoy_service_status_v2_csds_proto_rawDescGZIP(), []int{2}
}

func (x *ClientConfig) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *ClientConfig) GetXdsConfig() []*PerXdsConfig {
	if x != nil {
		return x.XdsConfig
	}
	return nil
}

type ClientStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Client configs for the clients specified in the ClientStatusRequest.
	Config []*ClientConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *ClientStatusResponse) Reset() {
	*x = ClientStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_status_v2_csds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStatusResponse) ProtoMessage() {}

func (x *ClientStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_status_v2_csds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStatusResponse.ProtoReflect.Descriptor instead.
func (*ClientStatusResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_status_v2_csds_proto_rawDescGZIP(), []int{3}
}

func (x *ClientStatusResponse) GetConfig() []*ClientConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_service_status_v2_csds_proto protoreflect.FileDescriptor

var file_envoy_service_status_v2_csds_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x73, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x25, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x6e,
	0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0xb1, 0x03, 0x0a, 0x0c,
	0x50, 0x65, 0x72, 0x58, 0x64, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a, 0x0f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x50, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5d,
	0x0a, 0x13, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x11, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x10, 0x0a,
	0x0e, 0x70, 0x65, 0x72, 0x5f, 0x78, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x81, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a,
	0x0a, 0x78, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x58,
	0x64, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x78, 0x64, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x55, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x4b, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x4e, 0x43, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32, 0xb8, 0x02, 0x0a, 0x1c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x9e, 0x01, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x1b, 0x2f, 0x76,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a,
	0x01, 0x2a, 0x42, 0x3f, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x43, 0x73, 0x64,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_status_v2_csds_proto_rawDescOnce sync.Once
	file_envoy_service_status_v2_csds_proto_rawDescData = file_envoy_service_status_v2_csds_proto_rawDesc
)

func file_envoy_service_status_v2_csds_proto_rawDescGZIP() []byte {
	file_envoy_service_status_v2_csds_proto_rawDescOnce.Do(func() {
		file_envoy_service_status_v2_csds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_status_v2_csds_proto_rawDescData)
	})
	return file_envoy_service_status_v2_csds_proto_rawDescData
}

var file_envoy_service_status_v2_csds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_service_status_v2_csds_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_service_status_v2_csds_proto_goTypes = []interface{}{
	(ConfigStatus)(0),                      // 0: envoy.service.status.v2.ConfigStatus
	(*ClientStatusRequest)(nil),            // 1: envoy.service.status.v2.ClientStatusRequest
	(*PerXdsConfig)(nil),                   // 2: envoy.service.status.v2.PerXdsConfig
	(*ClientConfig)(nil),                   // 3: envoy.service.status.v2.ClientConfig
	(*ClientStatusResponse)(nil),           // 4: envoy.service.status.v2.ClientStatusResponse
	(*matcher.NodeMatcher)(nil),            // 5: envoy.type.matcher.NodeMatcher
	(*v2alpha.ListenersConfigDump)(nil),    // 6: envoy.admin.v2alpha.ListenersConfigDump
	(*v2alpha.ClustersConfigDump)(nil),     // 7: envoy.admin.v2alpha.ClustersConfigDump
	(*v2alpha.RoutesConfigDump)(nil),       // 8: envoy.admin.v2alpha.RoutesConfigDump
	(*v2alpha.ScopedRoutesConfigDump)(nil), // 9: envoy.admin.v2alpha.ScopedRoutesConfigDump
	(*core.Node)(nil),                      // 10: envoy.api.v2.core.Node
}
var file_envoy_service_status_v2_csds_proto_depIdxs = []int32{
	5,  // 0: envoy.service.status.v2.ClientStatusRequest.node_matchers:type_name -> envoy.type.matcher.NodeMatcher
	0,  // 1: envoy.service.status.v2.PerXdsConfig.status:type_name -> envoy.service.status.v2.ConfigStatus
	6,  // 2: envoy.service.status.v2.PerXdsConfig.listener_config:type_name -> envoy.admin.v2alpha.ListenersConfigDump
	7,  // 3: envoy.service.status.v2.PerXdsConfig.cluster_config:type_name -> envoy.admin.v2alpha.ClustersConfigDump
	8,  // 4: envoy.service.status.v2.PerXdsConfig.route_config:type_name -> envoy.admin.v2alpha.RoutesConfigDump
	9,  // 5: envoy.service.status.v2.PerXdsConfig.scoped_route_config:type_name -> envoy.admin.v2alpha.ScopedRoutesConfigDump
	10, // 6: envoy.service.status.v2.ClientConfig.node:type_name -> envoy.api.v2.core.Node
	2,  // 7: envoy.service.status.v2.ClientConfig.xds_config:type_name -> envoy.service.status.v2.PerXdsConfig
	3,  // 8: envoy.service.status.v2.ClientStatusResponse.config:type_name -> envoy.service.status.v2.ClientConfig
	1,  // 9: envoy.service.status.v2.ClientStatusDiscoveryService.StreamClientStatus:input_type -> envoy.service.status.v2.ClientStatusRequest
	1,  // 10: envoy.service.status.v2.ClientStatusDiscoveryService.FetchClientStatus:input_type -> envoy.service.status.v2.ClientStatusRequest
	4,  // 11: envoy.service.status.v2.ClientStatusDiscoveryService.StreamClientStatus:output_type -> envoy.service.status.v2.ClientStatusResponse
	4,  // 12: envoy.service.status.v2.ClientStatusDiscoveryService.FetchClientStatus:output_type -> envoy.service.status.v2.ClientStatusResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_service_status_v2_csds_proto_init() }
func file_envoy_service_status_v2_csds_proto_init() {
	if File_envoy_service_status_v2_csds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_status_v2_csds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStatusRequest); i {
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
		file_envoy_service_status_v2_csds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerXdsConfig); i {
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
		file_envoy_service_status_v2_csds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_envoy_service_status_v2_csds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStatusResponse); i {
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
	file_envoy_service_status_v2_csds_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PerXdsConfig_ListenerConfig)(nil),
		(*PerXdsConfig_ClusterConfig)(nil),
		(*PerXdsConfig_RouteConfig)(nil),
		(*PerXdsConfig_ScopedRouteConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_status_v2_csds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_status_v2_csds_proto_goTypes,
		DependencyIndexes: file_envoy_service_status_v2_csds_proto_depIdxs,
		EnumInfos:         file_envoy_service_status_v2_csds_proto_enumTypes,
		MessageInfos:      file_envoy_service_status_v2_csds_proto_msgTypes,
	}.Build()
	File_envoy_service_status_v2_csds_proto = out.File
	file_envoy_service_status_v2_csds_proto_rawDesc = nil
	file_envoy_service_status_v2_csds_proto_goTypes = nil
	file_envoy_service_status_v2_csds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientStatusDiscoveryServiceClient is the client API for ClientStatusDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientStatusDiscoveryServiceClient interface {
	StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error)
	FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error)
}

type clientStatusDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStatusDiscoveryServiceClient(cc grpc.ClientConnInterface) ClientStatusDiscoveryServiceClient {
	return &clientStatusDiscoveryServiceClient{cc}
}

func (c *clientStatusDiscoveryServiceClient) StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientStatusDiscoveryService_serviceDesc.Streams[0], "/envoy.service.status.v2.ClientStatusDiscoveryService/StreamClientStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStatusDiscoveryServiceStreamClientStatusClient{stream}
	return x, nil
}

type ClientStatusDiscoveryService_StreamClientStatusClient interface {
	Send(*ClientStatusRequest) error
	Recv() (*ClientStatusResponse, error)
	grpc.ClientStream
}

type clientStatusDiscoveryServiceStreamClientStatusClient struct {
	grpc.ClientStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Send(m *ClientStatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Recv() (*ClientStatusResponse, error) {
	m := new(ClientStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientStatusDiscoveryServiceClient) FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error) {
	out := new(ClientStatusResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.status.v2.ClientStatusDiscoveryService/FetchClientStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientStatusDiscoveryServiceServer is the server API for ClientStatusDiscoveryService service.
type ClientStatusDiscoveryServiceServer interface {
	StreamClientStatus(ClientStatusDiscoveryService_StreamClientStatusServer) error
	FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error)
}

// UnimplementedClientStatusDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientStatusDiscoveryServiceServer struct {
}

func (*UnimplementedClientStatusDiscoveryServiceServer) StreamClientStatus(ClientStatusDiscoveryService_StreamClientStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientStatus not implemented")
}
func (*UnimplementedClientStatusDiscoveryServiceServer) FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClientStatus not implemented")
}

func RegisterClientStatusDiscoveryServiceServer(s *grpc.Server, srv ClientStatusDiscoveryServiceServer) {
	s.RegisterService(&_ClientStatusDiscoveryService_serviceDesc, srv)
}

func _ClientStatusDiscoveryService_StreamClientStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStatusDiscoveryServiceServer).StreamClientStatus(&clientStatusDiscoveryServiceStreamClientStatusServer{stream})
}

type ClientStatusDiscoveryService_StreamClientStatusServer interface {
	Send(*ClientStatusResponse) error
	Recv() (*ClientStatusRequest, error)
	grpc.ServerStream
}

type clientStatusDiscoveryServiceStreamClientStatusServer struct {
	grpc.ServerStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Send(m *ClientStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Recv() (*ClientStatusRequest, error) {
	m := new(ClientStatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientStatusDiscoveryService_FetchClientStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.status.v2.ClientStatusDiscoveryService/FetchClientStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, req.(*ClientStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientStatusDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.status.v2.ClientStatusDiscoveryService",
	HandlerType: (*ClientStatusDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClientStatus",
			Handler:    _ClientStatusDiscoveryService_FetchClientStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientStatus",
			Handler:       _ClientStatusDiscoveryService_StreamClientStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/status/v2/csds.proto",
}
