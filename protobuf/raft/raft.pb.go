// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/raft/raft.proto

package raft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BindAddr             string   `protobuf:"bytes,2,opt,name=bind_addr,json=bindAddr,proto3" json:"bind_addr,omitempty"`
	GrpcAddr             string   `protobuf:"bytes,3,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	HttpAddr             string   `protobuf:"bytes,4,opt,name=http_addr,json=httpAddr,proto3" json:"http_addr,omitempty"`
	Leader               bool     `protobuf:"varint,5,opt,name=leader,proto3" json:"leader,omitempty"`
	DataDir              string   `protobuf:"bytes,6,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_028aa12295c796d4, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *Node) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *Node) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

func (m *Node) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

func (m *Node) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

type Cluster struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_028aa12295c796d4, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "raft.Node")
	proto.RegisterType((*Cluster)(nil), "raft.Cluster")
}

func init() { proto.RegisterFile("protobuf/raft/raft.proto", fileDescriptor_028aa12295c796d4) }

var fileDescriptor_028aa12295c796d4 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x6a, 0xc3, 0x30,
	0x0c, 0x87, 0x71, 0x9a, 0xa6, 0xa9, 0x06, 0x3b, 0xf8, 0x30, 0x3c, 0x76, 0x09, 0x85, 0xb1, 0xc0,
	0x20, 0x81, 0xed, 0x09, 0xf6, 0xe7, 0xbc, 0x43, 0x8e, 0xbb, 0x14, 0xbb, 0x72, 0x5b, 0xb3, 0x34,
	0x0e, 0x8a, 0xfc, 0x3e, 0x7b, 0xd4, 0x61, 0x7b, 0x3b, 0xec, 0x22, 0xf4, 0xfb, 0x3e, 0x04, 0x92,
	0x40, 0xcd, 0xe4, 0xd9, 0x9b, 0x70, 0xec, 0x49, 0x1f, 0x39, 0x95, 0x2e, 0x21, 0x59, 0xc6, 0x7e,
	0xf7, 0x2d, 0xa0, 0xfc, 0xf0, 0x68, 0xe5, 0x35, 0x14, 0x0e, 0x95, 0x68, 0x44, 0xbb, 0x1d, 0x0a,
	0x87, 0xf2, 0x0e, 0xb6, 0xc6, 0x4d, 0xb8, 0xd7, 0x88, 0xa4, 0x8a, 0x84, 0xeb, 0x08, 0x5e, 0x10,
	0x29, 0xca, 0x13, 0xcd, 0x87, 0x2c, 0x57, 0x59, 0x46, 0xf0, 0x27, 0xcf, 0xcc, 0x73, 0x96, 0x65,
	0x96, 0x11, 0x24, 0x79, 0x03, 0xd5, 0x68, 0x35, 0x5a, 0x52, 0xeb, 0x46, 0xb4, 0xf5, 0xf0, 0x9b,
	0xe4, 0x2d, 0xd4, 0xa8, 0x59, 0xef, 0xd1, 0x91, 0xaa, 0xd2, 0xcc, 0x26, 0xe6, 0x77, 0x47, 0xbb,
	0x47, 0xd8, 0xbc, 0x8d, 0x61, 0x61, 0x4b, 0xb2, 0x81, 0xf5, 0xe4, 0xd1, 0x2e, 0x4a, 0x34, 0xab,
	0xf6, 0xea, 0x09, 0xba, 0x74, 0x4f, 0xdc, 0x7f, 0xc8, 0xe2, 0xf5, 0xe1, 0xf3, 0xfe, 0xe4, 0xf8,
	0x1c, 0x4c, 0x77, 0xf0, 0x97, 0xfe, 0xe2, 0x97, 0xf0, 0xa5, 0x7b, 0x33, 0xea, 0x85, 0xfb, 0x7f,
	0x9f, 0x30, 0x55, 0x8a, 0xcf, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x41, 0x78, 0x3c, 0x21,
	0x01, 0x00, 0x00,
}
