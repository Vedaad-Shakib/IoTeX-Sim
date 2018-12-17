// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trie.proto

package trie

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BranchNodePb struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BranchNodePb) Reset()         { *m = BranchNodePb{} }
func (m *BranchNodePb) String() string { return proto.CompactTextString(m) }
func (*BranchNodePb) ProtoMessage()    {}
func (*BranchNodePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_trie_4cef26911a66d6bd, []int{0}
}
func (m *BranchNodePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchNodePb.Unmarshal(m, b)
}
func (m *BranchNodePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchNodePb.Marshal(b, m, deterministic)
}
func (dst *BranchNodePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchNodePb.Merge(dst, src)
}
func (m *BranchNodePb) XXX_Size() int {
	return xxx_messageInfo_BranchNodePb.Size(m)
}
func (m *BranchNodePb) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchNodePb.DiscardUnknown(m)
}

var xxx_messageInfo_BranchNodePb proto.InternalMessageInfo

func (m *BranchNodePb) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BranchNodePb) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type BranchPb struct {
	Branches             []*BranchNodePb `protobuf:"bytes,1,rep,name=branches,proto3" json:"branches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BranchPb) Reset()         { *m = BranchPb{} }
func (m *BranchPb) String() string { return proto.CompactTextString(m) }
func (*BranchPb) ProtoMessage()    {}
func (*BranchPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_trie_4cef26911a66d6bd, []int{1}
}
func (m *BranchPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchPb.Unmarshal(m, b)
}
func (m *BranchPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchPb.Marshal(b, m, deterministic)
}
func (dst *BranchPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchPb.Merge(dst, src)
}
func (m *BranchPb) XXX_Size() int {
	return xxx_messageInfo_BranchPb.Size(m)
}
func (m *BranchPb) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchPb.DiscardUnknown(m)
}

var xxx_messageInfo_BranchPb proto.InternalMessageInfo

func (m *BranchPb) GetBranches() []*BranchNodePb {
	if m != nil {
		return m.Branches
	}
	return nil
}

type LeafPb struct {
	Ext                  uint32   `protobuf:"varint,1,opt,name=ext,proto3" json:"ext,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeafPb) Reset()         { *m = LeafPb{} }
func (m *LeafPb) String() string { return proto.CompactTextString(m) }
func (*LeafPb) ProtoMessage()    {}
func (*LeafPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_trie_4cef26911a66d6bd, []int{2}
}
func (m *LeafPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeafPb.Unmarshal(m, b)
}
func (m *LeafPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeafPb.Marshal(b, m, deterministic)
}
func (dst *LeafPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafPb.Merge(dst, src)
}
func (m *LeafPb) XXX_Size() int {
	return xxx_messageInfo_LeafPb.Size(m)
}
func (m *LeafPb) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafPb.DiscardUnknown(m)
}

var xxx_messageInfo_LeafPb proto.InternalMessageInfo

func (m *LeafPb) GetExt() uint32 {
	if m != nil {
		return m.Ext
	}
	return 0
}

func (m *LeafPb) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *LeafPb) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ExtendPb struct {
	Path                 []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtendPb) Reset()         { *m = ExtendPb{} }
func (m *ExtendPb) String() string { return proto.CompactTextString(m) }
func (*ExtendPb) ProtoMessage()    {}
func (*ExtendPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_trie_4cef26911a66d6bd, []int{3}
}
func (m *ExtendPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendPb.Unmarshal(m, b)
}
func (m *ExtendPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendPb.Marshal(b, m, deterministic)
}
func (dst *ExtendPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendPb.Merge(dst, src)
}
func (m *ExtendPb) XXX_Size() int {
	return xxx_messageInfo_ExtendPb.Size(m)
}
func (m *ExtendPb) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendPb.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendPb proto.InternalMessageInfo

func (m *ExtendPb) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ExtendPb) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NodePb struct {
	// Types that are valid to be assigned to Node:
	//	*NodePb_Branch
	//	*NodePb_Leaf
	//	*NodePb_Extend
	Node                 isNodePb_Node `protobuf_oneof:"node"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodePb) Reset()         { *m = NodePb{} }
func (m *NodePb) String() string { return proto.CompactTextString(m) }
func (*NodePb) ProtoMessage()    {}
func (*NodePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_trie_4cef26911a66d6bd, []int{4}
}
func (m *NodePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePb.Unmarshal(m, b)
}
func (m *NodePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePb.Marshal(b, m, deterministic)
}
func (dst *NodePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePb.Merge(dst, src)
}
func (m *NodePb) XXX_Size() int {
	return xxx_messageInfo_NodePb.Size(m)
}
func (m *NodePb) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePb.DiscardUnknown(m)
}

var xxx_messageInfo_NodePb proto.InternalMessageInfo

type isNodePb_Node interface {
	isNodePb_Node()
}

type NodePb_Branch struct {
	Branch *BranchPb `protobuf:"bytes,2,opt,name=branch,proto3,oneof"`
}

type NodePb_Leaf struct {
	Leaf *LeafPb `protobuf:"bytes,3,opt,name=leaf,proto3,oneof"`
}

type NodePb_Extend struct {
	Extend *ExtendPb `protobuf:"bytes,4,opt,name=extend,proto3,oneof"`
}

func (*NodePb_Branch) isNodePb_Node() {}

func (*NodePb_Leaf) isNodePb_Node() {}

func (*NodePb_Extend) isNodePb_Node() {}

func (m *NodePb) GetNode() isNodePb_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *NodePb) GetBranch() *BranchPb {
	if x, ok := m.GetNode().(*NodePb_Branch); ok {
		return x.Branch
	}
	return nil
}

func (m *NodePb) GetLeaf() *LeafPb {
	if x, ok := m.GetNode().(*NodePb_Leaf); ok {
		return x.Leaf
	}
	return nil
}

func (m *NodePb) GetExtend() *ExtendPb {
	if x, ok := m.GetNode().(*NodePb_Extend); ok {
		return x.Extend
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NodePb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NodePb_OneofMarshaler, _NodePb_OneofUnmarshaler, _NodePb_OneofSizer, []interface{}{
		(*NodePb_Branch)(nil),
		(*NodePb_Leaf)(nil),
		(*NodePb_Extend)(nil),
	}
}

func _NodePb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NodePb)
	// node
	switch x := m.Node.(type) {
	case *NodePb_Branch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Branch); err != nil {
			return err
		}
	case *NodePb_Leaf:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leaf); err != nil {
			return err
		}
	case *NodePb_Extend:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Extend); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodePb.Node has unexpected type %T", x)
	}
	return nil
}

func _NodePb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NodePb)
	switch tag {
	case 2: // node.branch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BranchPb)
		err := b.DecodeMessage(msg)
		m.Node = &NodePb_Branch{msg}
		return true, err
	case 3: // node.leaf
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeafPb)
		err := b.DecodeMessage(msg)
		m.Node = &NodePb_Leaf{msg}
		return true, err
	case 4: // node.extend
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExtendPb)
		err := b.DecodeMessage(msg)
		m.Node = &NodePb_Extend{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NodePb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NodePb)
	// node
	switch x := m.Node.(type) {
	case *NodePb_Branch:
		s := proto.Size(x.Branch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NodePb_Leaf:
		s := proto.Size(x.Leaf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NodePb_Extend:
		s := proto.Size(x.Extend)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BranchNodePb)(nil), "trie.branchNodePb")
	proto.RegisterType((*BranchPb)(nil), "trie.branchPb")
	proto.RegisterType((*LeafPb)(nil), "trie.leafPb")
	proto.RegisterType((*ExtendPb)(nil), "trie.extendPb")
	proto.RegisterType((*NodePb)(nil), "trie.nodePb")
}

func init() { proto.RegisterFile("trie.proto", fileDescriptor_trie_4cef26911a66d6bd) }

var fileDescriptor_trie_4cef26911a66d6bd = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0xa3, 0x58, 0x15, 0xe1, 0xec, 0x96, 0x72, 0x74, 0xd0, 0x68, 0x34, 0x79, 0xf2, 0xe0,
	0x76, 0x28, 0x1d, 0x4b, 0x87, 0x4c, 0x45, 0xe8, 0x1f, 0x58, 0xf5, 0x95, 0x04, 0x82, 0x1d, 0x5c,
	0xb5, 0xf8, 0x27, 0xf4, 0x67, 0x07, 0xe9, 0x92, 0xa0, 0x21, 0xdb, 0x93, 0xef, 0xde, 0x7b, 0xdf,
	0x61, 0x80, 0x30, 0xef, 0xa9, 0x3d, 0xce, 0x53, 0x98, 0x50, 0x46, 0x6d, 0x5e, 0xa1, 0xf2, 0x73,
	0x3f, 0x7e, 0xed, 0x3e, 0xa7, 0x81, 0xac, 0xc7, 0x27, 0xb8, 0xdb, 0x8f, 0x03, 0x2d, 0x5a, 0xd4,
	0xa2, 0xb9, 0x77, 0xfc, 0x40, 0x04, 0x79, 0xec, 0xc3, 0x4e, 0xaf, 0x6b, 0xd1, 0x54, 0x2e, 0x69,
	0xf3, 0x06, 0x1b, 0x76, 0x5a, 0x8f, 0xed, 0x45, 0xd3, 0x8f, 0x16, 0x75, 0xd1, 0x94, 0x1d, 0xb6,
	0xa9, 0x2a, 0xcf, 0x76, 0xd7, 0x1d, 0xf3, 0x01, 0xea, 0x40, 0xfd, 0xb7, 0xf5, 0xf8, 0x08, 0x05,
	0x2d, 0xe1, 0xdc, 0x16, 0xe5, 0xad, 0xae, 0x48, 0xf5, 0xd7, 0x1f, 0x7e, 0x49, 0x17, 0xe9, 0x23,
	0x3f, 0xcc, 0x0b, 0x6c, 0x68, 0x09, 0x34, 0x0e, 0xd6, 0x5f, 0x5d, 0xe2, 0x96, 0x6b, 0x9d, 0xbb,
	0xfe, 0x05, 0xa8, 0x91, 0x8f, 0x6d, 0x40, 0x31, 0x52, 0xda, 0x28, 0xbb, 0x87, 0x1c, 0xda, 0xfa,
	0xed, 0xca, 0x9d, 0xe7, 0x68, 0x40, 0x46, 0xe0, 0xd4, 0x5f, 0x76, 0x15, 0xef, 0xf1, 0x09, 0xdb,
	0x95, 0x4b, 0xb3, 0x98, 0xc6, 0x38, 0x5a, 0xe6, 0x69, 0x17, 0xc4, 0x98, 0xc6, 0xfa, 0x5d, 0x81,
	0x8c, 0x04, 0x5e, 0xa5, 0x3f, 0xf1, 0x7c, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x5a, 0x49, 0xc1,
	0x97, 0x01, 0x00, 0x00,
}
