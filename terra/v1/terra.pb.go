// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stellarproject/nebula/terra/v1/terra.proto

package v1 // import "github.com/stellarproject/nebula/terra/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import time "time"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NodeStatus_Status int32

const (
	NodeStatus_UNKNOWN  NodeStatus_Status = 0
	NodeStatus_OK       NodeStatus_Status = 1
	NodeStatus_UPDATING NodeStatus_Status = 2
	NodeStatus_FAILURE  NodeStatus_Status = 3
)

var NodeStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "UPDATING",
	3: "FAILURE",
}
var NodeStatus_Status_value = map[string]int32{
	"UNKNOWN":  0,
	"OK":       1,
	"UPDATING": 2,
	"FAILURE":  3,
}

func (x NodeStatus_Status) String() string {
	return proto.EnumName(NodeStatus_Status_name, int32(x))
}
func (NodeStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{10, 0}
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	ManifestList         *ManifestList `protobuf:"bytes,1,opt,name=manifest_list,json=manifestList" json:"manifest_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetManifestList() *ManifestList {
	if m != nil {
		return m.ManifestList
	}
	return nil
}

type Assembly struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assembly) Reset()         { *m = Assembly{} }
func (m *Assembly) String() string { return proto.CompactTextString(m) }
func (*Assembly) ProtoMessage()    {}
func (*Assembly) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{2}
}
func (m *Assembly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assembly.Unmarshal(m, b)
}
func (m *Assembly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assembly.Marshal(b, m, deterministic)
}
func (dst *Assembly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assembly.Merge(dst, src)
}
func (m *Assembly) XXX_Size() int {
	return xxx_messageInfo_Assembly.Size(m)
}
func (m *Assembly) XXX_DiscardUnknown() {
	xxx_messageInfo_Assembly.DiscardUnknown(m)
}

var xxx_messageInfo_Assembly proto.InternalMessageInfo

func (m *Assembly) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type Manifest struct {
	NodeID               string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Assemblies           []*Assembly       `protobuf:"bytes,3,rep,name=assemblies" json:"assemblies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{3}
}
func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest.Unmarshal(m, b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
}
func (dst *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(dst, src)
}
func (m *Manifest) XXX_Size() int {
	return xxx_messageInfo_Manifest.Size(m)
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *Manifest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Manifest) GetAssemblies() []*Assembly {
	if m != nil {
		return m.Assemblies
	}
	return nil
}

type ManifestList struct {
	Manifests            []*Manifest `protobuf:"bytes,1,rep,name=manifests" json:"manifests,omitempty"`
	Updated              time.Time   `protobuf:"bytes,2,opt,name=updated,stdtime" json:"updated"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ManifestList) Reset()         { *m = ManifestList{} }
func (m *ManifestList) String() string { return proto.CompactTextString(m) }
func (*ManifestList) ProtoMessage()    {}
func (*ManifestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{4}
}
func (m *ManifestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestList.Unmarshal(m, b)
}
func (m *ManifestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestList.Marshal(b, m, deterministic)
}
func (dst *ManifestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestList.Merge(dst, src)
}
func (m *ManifestList) XXX_Size() int {
	return xxx_messageInfo_ManifestList.Size(m)
}
func (m *ManifestList) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestList.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestList proto.InternalMessageInfo

func (m *ManifestList) GetManifests() []*Manifest {
	if m != nil {
		return m.Manifests
	}
	return nil
}

func (m *ManifestList) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

type ApplyRequest struct {
	ManifestList         *ManifestList `protobuf:"bytes,1,opt,name=manifest_list,json=manifestList" json:"manifest_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ApplyRequest) Reset()         { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()    {}
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{5}
}
func (m *ApplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRequest.Unmarshal(m, b)
}
func (m *ApplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRequest.Merge(dst, src)
}
func (m *ApplyRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRequest.Size(m)
}
func (m *ApplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRequest proto.InternalMessageInfo

func (m *ApplyRequest) GetManifestList() *ManifestList {
	if m != nil {
		return m.ManifestList
	}
	return nil
}

type NodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{6}
}
func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (dst *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(dst, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

type Node struct {
	ID                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{7}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type NodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{8}
}
func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (dst *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(dst, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{9}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type NodeStatus struct {
	NodeID               string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Status               NodeStatus_Status `protobuf:"varint,2,opt,name=status,proto3,enum=io.stellarproject.terra.v1.NodeStatus_Status" json:"status,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeStatus) Reset()         { *m = NodeStatus{} }
func (m *NodeStatus) String() string { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()    {}
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{10}
}
func (m *NodeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatus.Unmarshal(m, b)
}
func (m *NodeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatus.Marshal(b, m, deterministic)
}
func (dst *NodeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatus.Merge(dst, src)
}
func (m *NodeStatus) XXX_Size() int {
	return xxx_messageInfo_NodeStatus.Size(m)
}
func (m *NodeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatus proto.InternalMessageInfo

func (m *NodeStatus) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *NodeStatus) GetStatus() NodeStatus_Status {
	if m != nil {
		return m.Status
	}
	return NodeStatus_UNKNOWN
}

func (m *NodeStatus) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type StatusResponse struct {
	Nodes                []*NodeStatus `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_terra_42a76df99df76677, []int{11}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetNodes() []*NodeStatus {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "io.stellarproject.terra.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "io.stellarproject.terra.v1.ListResponse")
	proto.RegisterType((*Assembly)(nil), "io.stellarproject.terra.v1.Assembly")
	proto.RegisterType((*Manifest)(nil), "io.stellarproject.terra.v1.Manifest")
	proto.RegisterMapType((map[string]string)(nil), "io.stellarproject.terra.v1.Manifest.LabelsEntry")
	proto.RegisterType((*ManifestList)(nil), "io.stellarproject.terra.v1.ManifestList")
	proto.RegisterType((*ApplyRequest)(nil), "io.stellarproject.terra.v1.ApplyRequest")
	proto.RegisterType((*NodesRequest)(nil), "io.stellarproject.terra.v1.NodesRequest")
	proto.RegisterType((*Node)(nil), "io.stellarproject.terra.v1.Node")
	proto.RegisterMapType((map[string]string)(nil), "io.stellarproject.terra.v1.Node.LabelsEntry")
	proto.RegisterType((*NodesResponse)(nil), "io.stellarproject.terra.v1.NodesResponse")
	proto.RegisterType((*StatusRequest)(nil), "io.stellarproject.terra.v1.StatusRequest")
	proto.RegisterType((*NodeStatus)(nil), "io.stellarproject.terra.v1.NodeStatus")
	proto.RegisterType((*StatusResponse)(nil), "io.stellarproject.terra.v1.StatusResponse")
	proto.RegisterEnum("io.stellarproject.terra.v1.NodeStatus_Status", NodeStatus_Status_name, NodeStatus_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TerraClient is the client API for Terra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TerraClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type terraClient struct {
	cc *grpc.ClientConn
}

func NewTerraClient(cc *grpc.ClientConn) TerraClient {
	return &terraClient{cc}
}

func (c *terraClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.terra.v1.Terra/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.terra.v1.Terra/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraClient) Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error) {
	out := new(NodesResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.terra.v1.Terra/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.terra.v1.Terra/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TerraServer is the server API for Terra service.
type TerraServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Apply(context.Context, *ApplyRequest) (*types.Empty, error)
	Nodes(context.Context, *NodesRequest) (*NodesResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

func RegisterTerraServer(s *grpc.Server, srv TerraServer) {
	s.RegisterService(&_Terra_serviceDesc, srv)
}

func _Terra_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.terra.v1.Terra/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terra_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.terra.v1.Terra/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terra_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.terra.v1.Terra/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraServer).Nodes(ctx, req.(*NodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terra_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.terra.v1.Terra/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Terra_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.stellarproject.terra.v1.Terra",
	HandlerType: (*TerraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Terra_List_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Terra_Apply_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Terra_Nodes_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Terra_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stellarproject/nebula/terra/v1/terra.proto",
}

func init() {
	proto.RegisterFile("github.com/stellarproject/nebula/terra/v1/terra.proto", fileDescriptor_terra_42a76df99df76677)
}

var fileDescriptor_terra_42a76df99df76677 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xc6, 0x49, 0x27, 0x49, 0x89, 0x56, 0x55, 0x65, 0x99, 0x87, 0x44, 0x06, 0x41,
	0x40, 0xc5, 0xa6, 0x41, 0xa0, 0x72, 0x11, 0x52, 0xa2, 0x84, 0x12, 0xb5, 0x4d, 0x91, 0x69, 0xc5,
	0x45, 0xa0, 0xca, 0xa9, 0xb7, 0xc1, 0x60, 0xc7, 0xc6, 0xbb, 0x8e, 0x94, 0xbf, 0x40, 0xfc, 0x0d,
	0x7f, 0xc0, 0x07, 0xf0, 0x5c, 0xa4, 0x3e, 0xf3, 0x05, 0x3c, 0x21, 0xef, 0xae, 0x83, 0xcb, 0x25,
	0x89, 0x04, 0x6f, 0x9e, 0xdd, 0x73, 0x66, 0xcf, 0xcc, 0x9c, 0x49, 0xe0, 0xf6, 0xd0, 0xa5, 0x6f,
	0xe2, 0x81, 0x71, 0x1c, 0xf8, 0x26, 0xa1, 0xd8, 0xf3, 0xec, 0x28, 0x8c, 0x82, 0xb7, 0xf8, 0x98,
	0x9a, 0x23, 0x3c, 0x88, 0x3d, 0xdb, 0xa4, 0x38, 0x8a, 0x6c, 0x73, 0xbc, 0xc9, 0x3f, 0x8c, 0x30,
	0x0a, 0x68, 0x80, 0x34, 0x37, 0x30, 0xce, 0xc3, 0x0d, 0x7e, 0x3d, 0xde, 0xd4, 0xd6, 0x86, 0xc1,
	0x30, 0x60, 0x30, 0x33, 0xf9, 0xe2, 0x0c, 0xad, 0x36, 0x0c, 0x82, 0xa1, 0x87, 0x4d, 0x16, 0x0d,
	0xe2, 0x13, 0x93, 0xba, 0x3e, 0x26, 0xd4, 0xf6, 0x43, 0x01, 0xb8, 0xf8, 0x2b, 0x00, 0xfb, 0x21,
	0x9d, 0xf0, 0x4b, 0xbd, 0x02, 0xa5, 0x5d, 0x97, 0x50, 0x0b, 0xbf, 0x8f, 0x31, 0xa1, 0xfa, 0x6b,
	0x28, 0xf3, 0x90, 0x84, 0xc1, 0x88, 0x60, 0xb4, 0x07, 0x15, 0xdf, 0x1e, 0xb9, 0x27, 0x98, 0xd0,
	0x23, 0xcf, 0x25, 0x54, 0x95, 0xea, 0x52, 0xa3, 0xd4, 0x6c, 0x18, 0x7f, 0x97, 0x69, 0xec, 0x09,
	0x02, 0x4b, 0x54, 0xf6, 0x33, 0x91, 0x5e, 0x87, 0x62, 0x8b, 0x10, 0xec, 0x0f, 0xbc, 0x09, 0x5a,
	0x83, 0xbc, 0xeb, 0xdb, 0x43, 0xcc, 0x52, 0xae, 0x58, 0x3c, 0xd0, 0xbf, 0x4b, 0x50, 0x4c, 0x13,
	0xa0, 0x4b, 0x50, 0x18, 0x05, 0x0e, 0x3e, 0x72, 0x1d, 0x0e, 0x6a, 0xc3, 0xd9, 0x69, 0x4d, 0xe9,
	0x07, 0x0e, 0xee, 0x75, 0x2c, 0x25, 0xb9, 0xea, 0x39, 0xe8, 0x31, 0x28, 0x9e, 0x3d, 0xc0, 0x1e,
	0x51, 0xe5, 0x7a, 0xae, 0x51, 0x6a, 0xde, 0x5c, 0x44, 0x9b, 0xb1, 0xcb, 0x28, 0xdd, 0x11, 0x8d,
	0x26, 0x96, 0xe0, 0xa3, 0x0e, 0x80, 0xcd, 0xd5, 0xb9, 0x98, 0xa8, 0x39, 0x96, 0xed, 0xf2, 0xac,
	0x6c, 0x69, 0x2d, 0x56, 0x86, 0xa7, 0xdd, 0x85, 0x52, 0x26, 0x39, 0xaa, 0x42, 0xee, 0x1d, 0x9e,
	0x88, 0x22, 0x93, 0xcf, 0xa4, 0xf0, 0xb1, 0xed, 0xc5, 0x58, 0x95, 0x79, 0xe1, 0x2c, 0xb8, 0x27,
	0x6f, 0x49, 0xfa, 0x47, 0x09, 0xca, 0xd9, 0xee, 0xa1, 0x36, 0xac, 0xa4, 0xfd, 0x23, 0xaa, 0x34,
	0x5f, 0x50, 0x4a, 0xb6, 0x7e, 0xd2, 0xd0, 0x43, 0x28, 0xc4, 0xa1, 0x63, 0x53, 0xec, 0xb0, 0x07,
	0x4b, 0x4d, 0xcd, 0xe0, 0x86, 0x30, 0x52, 0x43, 0x18, 0x07, 0xa9, 0x63, 0xda, 0xc5, 0xcf, 0xa7,
	0xb5, 0xa5, 0x0f, 0x5f, 0x6b, 0x92, 0x95, 0x92, 0x12, 0x4b, 0xb4, 0xc2, 0xd0, 0x9b, 0x08, 0x8b,
	0xfc, 0x6f, 0x4b, 0xac, 0x42, 0x39, 0x19, 0x28, 0x49, 0x1d, 0xf8, 0x49, 0x82, 0xe5, 0xe4, 0x00,
	0xad, 0x83, 0x3c, 0x9d, 0xbb, 0x72, 0x76, 0x5a, 0x93, 0x7b, 0x1d, 0x4b, 0x76, 0x1d, 0xa4, 0x42,
	0xc1, 0x76, 0x9c, 0x08, 0x13, 0x22, 0x1a, 0x98, 0x86, 0xa8, 0x33, 0x75, 0x02, 0x9f, 0xdd, 0xc6,
	0x2c, 0x49, 0xc9, 0x1b, 0x7f, 0x72, 0xc1, 0xbf, 0xcc, 0x6f, 0x1b, 0x2a, 0xa2, 0x16, 0xb1, 0x3e,
	0x77, 0x20, 0x9f, 0xb8, 0x34, 0x9d, 0x5d, 0x7d, 0x9e, 0x20, 0x8b, 0xc3, 0xf5, 0x0b, 0x50, 0x79,
	0x4a, 0x6d, 0x1a, 0x4f, 0xbb, 0xf2, 0x45, 0x02, 0x48, 0x00, 0xfc, 0x74, 0xb1, 0xc5, 0xe8, 0x82,
	0x42, 0x18, 0x9c, 0x09, 0x5d, 0x6d, 0xde, 0x98, 0xf7, 0x3a, 0x4f, 0x6e, 0x88, 0x97, 0x05, 0x19,
	0xd5, 0xa1, 0xe4, 0x60, 0x72, 0x1c, 0xb9, 0x21, 0x75, 0x83, 0x91, 0x9a, 0x63, 0x45, 0x67, 0x8f,
	0xf4, 0x2d, 0x50, 0x84, 0xae, 0x12, 0x14, 0x0e, 0xfb, 0x3b, 0xfd, 0xfd, 0x67, 0xfd, 0xea, 0x12,
	0x52, 0x40, 0xde, 0xdf, 0xa9, 0x4a, 0xa8, 0x0c, 0xc5, 0xc3, 0x27, 0x9d, 0xd6, 0x41, 0xaf, 0xbf,
	0x5d, 0x95, 0x13, 0xc8, 0xa3, 0x56, 0x6f, 0xf7, 0xd0, 0xea, 0x56, 0x73, 0x7a, 0x1f, 0x56, 0xd3,
	0x3a, 0x45, 0xc7, 0x1e, 0x9c, 0xef, 0xd8, 0x95, 0xc5, 0x34, 0x8b, 0xbe, 0x35, 0xbf, 0xc9, 0x90,
	0x3f, 0x48, 0xae, 0xd1, 0x0b, 0x58, 0x66, 0x1b, 0x74, 0x75, 0x56, 0x82, 0xcc, 0x2f, 0x9f, 0xd6,
	0x98, 0x0f, 0x14, 0x12, 0x7b, 0x90, 0x67, 0x0b, 0x81, 0x66, 0x52, 0xb2, 0x3b, 0xa3, 0xad, 0xff,
	0xb6, 0x72, 0xdd, 0xe4, 0x37, 0x18, 0xbd, 0x82, 0x3c, 0x33, 0xcc, 0xec, 0x54, 0xd9, 0xfd, 0xd0,
	0xae, 0x2d, 0x80, 0x14, 0x42, 0x8f, 0xa6, 0x73, 0x99, 0x49, 0x3a, 0xe7, 0x34, 0xed, 0xfa, 0x22,
	0x50, 0xfe, 0x40, 0xdb, 0x78, 0xb9, 0xb1, 0xf0, 0xbf, 0xdc, 0xfd, 0xf1, 0xe6, 0xf3, 0xa5, 0x81,
	0xc2, 0x1a, 0x70, 0xeb, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb4, 0xa4, 0x41, 0x1e, 0x07,
	0x00, 0x00,
}
