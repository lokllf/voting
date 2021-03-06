// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voting/protobuf/payload/voter.proto

package payload

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	voting "voting/protobuf/voting"
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

type VoterPayload_Action int32

const (
	VoterPayload_CAST_BALLOT VoterPayload_Action = 0
)

var VoterPayload_Action_name = map[int32]string{
	0: "CAST_BALLOT",
}

var VoterPayload_Action_value = map[string]int32{
	"CAST_BALLOT": 0,
}

func (x VoterPayload_Action) String() string {
	return proto.EnumName(VoterPayload_Action_name, int32(x))
}

func (VoterPayload_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dce8917277f4f4c0, []int{0, 0}
}

// VoterPayload represents the format of a payload for voter
type VoterPayload struct {
	Action               VoterPayload_Action          `protobuf:"varint,1,opt,name=action,proto3,enum=payload.VoterPayload_Action" json:"action,omitempty"`
	SubmittedAt          int64                        `protobuf:"varint,2,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	CastBallot           *VoterPayload_CastBallotData `protobuf:"bytes,3,opt,name=cast_ballot,json=castBallot,proto3" json:"cast_ballot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *VoterPayload) Reset()         { *m = VoterPayload{} }
func (m *VoterPayload) String() string { return proto.CompactTextString(m) }
func (*VoterPayload) ProtoMessage()    {}
func (*VoterPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce8917277f4f4c0, []int{0}
}

func (m *VoterPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoterPayload.Unmarshal(m, b)
}
func (m *VoterPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoterPayload.Marshal(b, m, deterministic)
}
func (m *VoterPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoterPayload.Merge(m, src)
}
func (m *VoterPayload) XXX_Size() int {
	return xxx_messageInfo_VoterPayload.Size(m)
}
func (m *VoterPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_VoterPayload.DiscardUnknown(m)
}

var xxx_messageInfo_VoterPayload proto.InternalMessageInfo

func (m *VoterPayload) GetAction() VoterPayload_Action {
	if m != nil {
		return m.Action
	}
	return VoterPayload_CAST_BALLOT
}

func (m *VoterPayload) GetSubmittedAt() int64 {
	if m != nil {
		return m.SubmittedAt
	}
	return 0
}

func (m *VoterPayload) GetCastBallot() *VoterPayload_CastBallotData {
	if m != nil {
		return m.CastBallot
	}
	return nil
}

type VoterPayload_CastBallotData struct {
	Ballot               *voting.Ballot `protobuf:"bytes,1,opt,name=ballot,proto3" json:"ballot,omitempty"`
	Code                 string         `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VoterPayload_CastBallotData) Reset()         { *m = VoterPayload_CastBallotData{} }
func (m *VoterPayload_CastBallotData) String() string { return proto.CompactTextString(m) }
func (*VoterPayload_CastBallotData) ProtoMessage()    {}
func (*VoterPayload_CastBallotData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce8917277f4f4c0, []int{0, 0}
}

func (m *VoterPayload_CastBallotData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoterPayload_CastBallotData.Unmarshal(m, b)
}
func (m *VoterPayload_CastBallotData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoterPayload_CastBallotData.Marshal(b, m, deterministic)
}
func (m *VoterPayload_CastBallotData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoterPayload_CastBallotData.Merge(m, src)
}
func (m *VoterPayload_CastBallotData) XXX_Size() int {
	return xxx_messageInfo_VoterPayload_CastBallotData.Size(m)
}
func (m *VoterPayload_CastBallotData) XXX_DiscardUnknown() {
	xxx_messageInfo_VoterPayload_CastBallotData.DiscardUnknown(m)
}

var xxx_messageInfo_VoterPayload_CastBallotData proto.InternalMessageInfo

func (m *VoterPayload_CastBallotData) GetBallot() *voting.Ballot {
	if m != nil {
		return m.Ballot
	}
	return nil
}

func (m *VoterPayload_CastBallotData) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterEnum("payload.VoterPayload_Action", VoterPayload_Action_name, VoterPayload_Action_value)
	proto.RegisterType((*VoterPayload)(nil), "payload.VoterPayload")
	proto.RegisterType((*VoterPayload_CastBallotData)(nil), "payload.VoterPayload.CastBallotData")
}

func init() {
	proto.RegisterFile("voting/protobuf/payload/voter.proto", fileDescriptor_dce8917277f4f4c0)
}

var fileDescriptor_dce8917277f4f4c0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xcd, 0x26, 0x15, 0xaf, 0xa3, 0x4a, 0x9e, 0xea, 0xf0, 0xa1, 0x4e, 0x91, 0x3e, 0xa5,
	0x30, 0xfd, 0x02, 0xdd, 0xf4, 0xad, 0xa0, 0xc4, 0xe1, 0x6b, 0xb9, 0xa6, 0x55, 0x0a, 0x75, 0x37,
	0xd6, 0xdb, 0xc0, 0x4f, 0xe2, 0xd7, 0x15, 0x93, 0x50, 0x14, 0xf7, 0x96, 0xfb, 0xdf, 0xef, 0x7e,
	0xc9, 0x05, 0xae, 0xf7, 0xc4, 0xed, 0xfa, 0x3d, 0xdb, 0x6c, 0x89, 0xa9, 0xda, 0xbd, 0x65, 0x1b,
	0xfc, 0xec, 0x08, 0xeb, 0x6c, 0x4f, 0xdc, 0x6c, 0x95, 0x8d, 0xe5, 0x89, 0x0f, 0xa7, 0xff, 0x68,
	0x5f, 0x57, 0xd8, 0x75, 0xc4, 0x8e, 0x9e, 0x7d, 0x8d, 0x60, 0xf2, 0xfa, 0x33, 0xfd, 0xec, 0xa6,
	0xe4, 0x3d, 0x04, 0x68, 0xb8, 0xa5, 0x75, 0x2c, 0x12, 0x91, 0x46, 0xf3, 0x4b, 0xe5, 0x7d, 0xea,
	0x37, 0xa6, 0x72, 0xcb, 0x68, 0xcf, 0xca, 0x2b, 0x98, 0xf4, 0xbb, 0xea, 0xa3, 0x65, 0x6e, 0xea,
	0x12, 0x39, 0x1e, 0x25, 0x22, 0x1d, 0xeb, 0x70, 0xc8, 0x72, 0x96, 0x8f, 0x10, 0x1a, 0xec, 0xb9,
	0x74, 0xd7, 0xc7, 0xe3, 0x44, 0xa4, 0xe1, 0xfc, 0xe6, 0xb0, 0x7d, 0x89, 0x3d, 0x2f, 0x2c, 0xf7,
	0x80, 0x8c, 0x1a, 0xcc, 0x50, 0x4f, 0x0b, 0x88, 0xfe, 0x76, 0xe5, 0x2d, 0x04, 0xde, 0x29, 0xac,
	0x33, 0x52, 0x6e, 0x51, 0xe5, 0x18, 0xed, 0xbb, 0x52, 0xc2, 0xb1, 0xa1, 0xba, 0xb1, 0x6f, 0x3b,
	0xd5, 0xf6, 0x3c, 0xbb, 0x80, 0xc0, 0x6d, 0x22, 0xcf, 0x20, 0x5c, 0xe6, 0x2f, 0xab, 0x72, 0x91,
	0x17, 0xc5, 0xd3, 0xea, 0xfc, 0xa8, 0x0a, 0xec, 0x07, 0xdd, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x7a, 0xd8, 0xf7, 0x75, 0x01, 0x00, 0x00,
}
