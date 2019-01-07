// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voting/protobuf/voting/result.proto

package voting

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Result represents a voting result
type Result struct {
	VoteId               string            `protobuf:"bytes,1,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
	Total                uint32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Casted               uint32            `protobuf:"varint,3,opt,name=casted,proto3" json:"casted,omitempty"`
	Counts               map[string]uint32 `protobuf:"bytes,4,rep,name=counts,proto3" json:"counts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	CreatedAt            int64             `protobuf:"varint,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_23cf5db4f419449e, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetVoteId() string {
	if m != nil {
		return m.VoteId
	}
	return ""
}

func (m *Result) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Result) GetCasted() uint32 {
	if m != nil {
		return m.Casted
	}
	return 0
}

func (m *Result) GetCounts() map[string]uint32 {
	if m != nil {
		return m.Counts
	}
	return nil
}

func (m *Result) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Result)(nil), "voting.Result")
	proto.RegisterMapType((map[string]uint32)(nil), "voting.Result.CountsEntry")
}

func init() {
	proto.RegisterFile("voting/protobuf/voting/result.proto", fileDescriptor_23cf5db4f419449e)
}

var fileDescriptor_23cf5db4f419449e = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x49, 0xa3, 0x91, 0x4e, 0x11, 0x25, 0x88, 0x86, 0x82, 0x10, 0xf4, 0x92, 0x53, 0x0a,
	0xf5, 0xa2, 0xde, 0x44, 0x3c, 0x78, 0xcd, 0x0b, 0x94, 0x74, 0x33, 0x4a, 0x71, 0xd9, 0x48, 0x76,
	0xb2, 0xd0, 0x87, 0xf5, 0x5d, 0x64, 0x93, 0x08, 0xde, 0xf2, 0x7d, 0x33, 0xe1, 0xff, 0x07, 0xee,
	0xa7, 0x48, 0x87, 0xe1, 0x73, 0xf3, 0x9d, 0x22, 0xc5, 0x7d, 0xfe, 0xd8, 0x34, 0x4e, 0x38, 0xe6,
	0x9e, 0x6c, 0xd1, 0x52, 0x54, 0x79, 0xf7, 0xc3, 0x40, 0xb8, 0x32, 0x90, 0x37, 0x70, 0x36, 0x45,
	0xc2, 0xdd, 0x21, 0x28, 0xa6, 0x99, 0x59, 0xba, 0x79, 0x07, 0xdf, 0x83, 0xbc, 0x82, 0x53, 0x8a,
	0xe4, 0x7b, 0xb5, 0xd0, 0xcc, 0x9c, 0xbb, 0x0a, 0xf2, 0x1a, 0x44, 0xe7, 0x47, 0xc2, 0xa0, 0x78,
	0xd1, 0x8d, 0xe4, 0x16, 0x44, 0x17, 0xf3, 0x40, 0xa3, 0x3a, 0xd1, 0xdc, 0xac, 0xb6, 0x6b, 0x5b,
	0xa3, 0x6c, 0x8d, 0xb1, 0xaf, 0x65, 0xf8, 0x36, 0x50, 0x3a, 0xba, 0xb6, 0x29, 0x6f, 0x01, 0xba,
	0x84, 0x9e, 0x30, 0xec, 0x3c, 0xa9, 0x0b, 0xcd, 0x0c, 0x77, 0xcb, 0x66, 0x5e, 0x68, 0xfd, 0x04,
	0xab, 0x7f, 0xbf, 0xe4, 0x25, 0xf0, 0x2f, 0x3c, 0xb6, 0x92, 0xf3, 0x73, 0x6e, 0x38, 0xf9, 0x3e,
	0xe3, 0x5f, 0xc3, 0x02, 0xcf, 0x8b, 0x47, 0xb6, 0x17, 0xe5, 0xdc, 0x87, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0x2a, 0xe1, 0x40, 0x15, 0x01, 0x00, 0x00,
}
