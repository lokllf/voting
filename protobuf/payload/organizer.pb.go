// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voting/protobuf/payload/organizer.proto

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

type OrganizerPayload_Action int32

const (
	OrganizerPayload_CREATE_VOTE  OrganizerPayload_Action = 0
	OrganizerPayload_UPDATE_VOTE  OrganizerPayload_Action = 1
	OrganizerPayload_DELETE_VOTE  OrganizerPayload_Action = 2
	OrganizerPayload_ADD_BALLOT   OrganizerPayload_Action = 3
	OrganizerPayload_COUNT_BALLOT OrganizerPayload_Action = 4
)

var OrganizerPayload_Action_name = map[int32]string{
	0: "CREATE_VOTE",
	1: "UPDATE_VOTE",
	2: "DELETE_VOTE",
	3: "ADD_BALLOT",
	4: "COUNT_BALLOT",
}

var OrganizerPayload_Action_value = map[string]int32{
	"CREATE_VOTE":  0,
	"UPDATE_VOTE":  1,
	"DELETE_VOTE":  2,
	"ADD_BALLOT":   3,
	"COUNT_BALLOT": 4,
}

func (x OrganizerPayload_Action) String() string {
	return proto.EnumName(OrganizerPayload_Action_name, int32(x))
}

func (OrganizerPayload_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 0}
}

// OrganizerPayload represents the format of a payload for organizer
type OrganizerPayload struct {
	Action               OrganizerPayload_Action           `protobuf:"varint,1,opt,name=action,proto3,enum=payload.OrganizerPayload_Action" json:"action,omitempty"`
	SubmittedAt          int64                             `protobuf:"varint,2,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	CreateVote           *OrganizerPayload_CreateVoteData  `protobuf:"bytes,3,opt,name=create_vote,json=createVote,proto3" json:"create_vote,omitempty"`
	UpdateVote           *OrganizerPayload_UpdateVoteData  `protobuf:"bytes,4,opt,name=update_vote,json=updateVote,proto3" json:"update_vote,omitempty"`
	DeleteVote           *OrganizerPayload_DeleteVoteData  `protobuf:"bytes,5,opt,name=delete_vote,json=deleteVote,proto3" json:"delete_vote,omitempty"`
	AddBallot            *OrganizerPayload_AddBallotData   `protobuf:"bytes,6,opt,name=add_ballot,json=addBallot,proto3" json:"add_ballot,omitempty"`
	CountBallot          *OrganizerPayload_CountBallotData `protobuf:"bytes,7,opt,name=count_ballot,json=countBallot,proto3" json:"count_ballot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *OrganizerPayload) Reset()         { *m = OrganizerPayload{} }
func (m *OrganizerPayload) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload) ProtoMessage()    {}
func (*OrganizerPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0}
}

func (m *OrganizerPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload.Unmarshal(m, b)
}
func (m *OrganizerPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload.Merge(m, src)
}
func (m *OrganizerPayload) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload.Size(m)
}
func (m *OrganizerPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload proto.InternalMessageInfo

func (m *OrganizerPayload) GetAction() OrganizerPayload_Action {
	if m != nil {
		return m.Action
	}
	return OrganizerPayload_CREATE_VOTE
}

func (m *OrganizerPayload) GetSubmittedAt() int64 {
	if m != nil {
		return m.SubmittedAt
	}
	return 0
}

func (m *OrganizerPayload) GetCreateVote() *OrganizerPayload_CreateVoteData {
	if m != nil {
		return m.CreateVote
	}
	return nil
}

func (m *OrganizerPayload) GetUpdateVote() *OrganizerPayload_UpdateVoteData {
	if m != nil {
		return m.UpdateVote
	}
	return nil
}

func (m *OrganizerPayload) GetDeleteVote() *OrganizerPayload_DeleteVoteData {
	if m != nil {
		return m.DeleteVote
	}
	return nil
}

func (m *OrganizerPayload) GetAddBallot() *OrganizerPayload_AddBallotData {
	if m != nil {
		return m.AddBallot
	}
	return nil
}

func (m *OrganizerPayload) GetCountBallot() *OrganizerPayload_CountBallotData {
	if m != nil {
		return m.CountBallot
	}
	return nil
}

type OrganizerPayload_CreateVoteData struct {
	Vote                 *voting.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrganizerPayload_CreateVoteData) Reset()         { *m = OrganizerPayload_CreateVoteData{} }
func (m *OrganizerPayload_CreateVoteData) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload_CreateVoteData) ProtoMessage()    {}
func (*OrganizerPayload_CreateVoteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 0}
}

func (m *OrganizerPayload_CreateVoteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload_CreateVoteData.Unmarshal(m, b)
}
func (m *OrganizerPayload_CreateVoteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload_CreateVoteData.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload_CreateVoteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload_CreateVoteData.Merge(m, src)
}
func (m *OrganizerPayload_CreateVoteData) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload_CreateVoteData.Size(m)
}
func (m *OrganizerPayload_CreateVoteData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload_CreateVoteData.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload_CreateVoteData proto.InternalMessageInfo

func (m *OrganizerPayload_CreateVoteData) GetVote() *voting.Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

type OrganizerPayload_UpdateVoteData struct {
	Vote                 *voting.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrganizerPayload_UpdateVoteData) Reset()         { *m = OrganizerPayload_UpdateVoteData{} }
func (m *OrganizerPayload_UpdateVoteData) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload_UpdateVoteData) ProtoMessage()    {}
func (*OrganizerPayload_UpdateVoteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 1}
}

func (m *OrganizerPayload_UpdateVoteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload_UpdateVoteData.Unmarshal(m, b)
}
func (m *OrganizerPayload_UpdateVoteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload_UpdateVoteData.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload_UpdateVoteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload_UpdateVoteData.Merge(m, src)
}
func (m *OrganizerPayload_UpdateVoteData) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload_UpdateVoteData.Size(m)
}
func (m *OrganizerPayload_UpdateVoteData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload_UpdateVoteData.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload_UpdateVoteData proto.InternalMessageInfo

func (m *OrganizerPayload_UpdateVoteData) GetVote() *voting.Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

type OrganizerPayload_DeleteVoteData struct {
	VoteId               string   `protobuf:"bytes,1,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizerPayload_DeleteVoteData) Reset()         { *m = OrganizerPayload_DeleteVoteData{} }
func (m *OrganizerPayload_DeleteVoteData) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload_DeleteVoteData) ProtoMessage()    {}
func (*OrganizerPayload_DeleteVoteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 2}
}

func (m *OrganizerPayload_DeleteVoteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload_DeleteVoteData.Unmarshal(m, b)
}
func (m *OrganizerPayload_DeleteVoteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload_DeleteVoteData.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload_DeleteVoteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload_DeleteVoteData.Merge(m, src)
}
func (m *OrganizerPayload_DeleteVoteData) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload_DeleteVoteData.Size(m)
}
func (m *OrganizerPayload_DeleteVoteData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload_DeleteVoteData.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload_DeleteVoteData proto.InternalMessageInfo

func (m *OrganizerPayload_DeleteVoteData) GetVoteId() string {
	if m != nil {
		return m.VoteId
	}
	return ""
}

type OrganizerPayload_AddBallotData struct {
	VoteId               string   `protobuf:"bytes,1,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
	HashedCode           string   `protobuf:"bytes,2,opt,name=hashed_code,json=hashedCode,proto3" json:"hashed_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizerPayload_AddBallotData) Reset()         { *m = OrganizerPayload_AddBallotData{} }
func (m *OrganizerPayload_AddBallotData) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload_AddBallotData) ProtoMessage()    {}
func (*OrganizerPayload_AddBallotData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 3}
}

func (m *OrganizerPayload_AddBallotData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload_AddBallotData.Unmarshal(m, b)
}
func (m *OrganizerPayload_AddBallotData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload_AddBallotData.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload_AddBallotData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload_AddBallotData.Merge(m, src)
}
func (m *OrganizerPayload_AddBallotData) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload_AddBallotData.Size(m)
}
func (m *OrganizerPayload_AddBallotData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload_AddBallotData.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload_AddBallotData proto.InternalMessageInfo

func (m *OrganizerPayload_AddBallotData) GetVoteId() string {
	if m != nil {
		return m.VoteId
	}
	return ""
}

func (m *OrganizerPayload_AddBallotData) GetHashedCode() string {
	if m != nil {
		return m.HashedCode
	}
	return ""
}

type OrganizerPayload_CountBallotData struct {
	VoteId               string   `protobuf:"bytes,1,opt,name=vote_id,json=voteId,proto3" json:"vote_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizerPayload_CountBallotData) Reset()         { *m = OrganizerPayload_CountBallotData{} }
func (m *OrganizerPayload_CountBallotData) String() string { return proto.CompactTextString(m) }
func (*OrganizerPayload_CountBallotData) ProtoMessage()    {}
func (*OrganizerPayload_CountBallotData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2dd429f0686dc6, []int{0, 4}
}

func (m *OrganizerPayload_CountBallotData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerPayload_CountBallotData.Unmarshal(m, b)
}
func (m *OrganizerPayload_CountBallotData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerPayload_CountBallotData.Marshal(b, m, deterministic)
}
func (m *OrganizerPayload_CountBallotData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerPayload_CountBallotData.Merge(m, src)
}
func (m *OrganizerPayload_CountBallotData) XXX_Size() int {
	return xxx_messageInfo_OrganizerPayload_CountBallotData.Size(m)
}
func (m *OrganizerPayload_CountBallotData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerPayload_CountBallotData.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerPayload_CountBallotData proto.InternalMessageInfo

func (m *OrganizerPayload_CountBallotData) GetVoteId() string {
	if m != nil {
		return m.VoteId
	}
	return ""
}

func init() {
	proto.RegisterEnum("payload.OrganizerPayload_Action", OrganizerPayload_Action_name, OrganizerPayload_Action_value)
	proto.RegisterType((*OrganizerPayload)(nil), "payload.OrganizerPayload")
	proto.RegisterType((*OrganizerPayload_CreateVoteData)(nil), "payload.OrganizerPayload.CreateVoteData")
	proto.RegisterType((*OrganizerPayload_UpdateVoteData)(nil), "payload.OrganizerPayload.UpdateVoteData")
	proto.RegisterType((*OrganizerPayload_DeleteVoteData)(nil), "payload.OrganizerPayload.DeleteVoteData")
	proto.RegisterType((*OrganizerPayload_AddBallotData)(nil), "payload.OrganizerPayload.AddBallotData")
	proto.RegisterType((*OrganizerPayload_CountBallotData)(nil), "payload.OrganizerPayload.CountBallotData")
}

func init() {
	proto.RegisterFile("voting/protobuf/payload/organizer.proto", fileDescriptor_db2dd429f0686dc6)
}

var fileDescriptor_db2dd429f0686dc6 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0x9b, 0x40,
	0x10, 0xc5, 0x4b, 0xec, 0x62, 0x79, 0x70, 0x09, 0xda, 0x4b, 0x11, 0x97, 0x92, 0x5c, 0x42, 0x7a,
	0xc0, 0x92, 0x7b, 0xe9, 0x95, 0x00, 0x95, 0x2c, 0xa1, 0x12, 0x21, 0x9c, 0x5b, 0x85, 0x16, 0x76,
	0x9b, 0x20, 0x51, 0xd6, 0x22, 0x4b, 0xa4, 0xf6, 0x4b, 0xf6, 0x2b, 0x55, 0xbb, 0xfc, 0x71, 0xb1,
	0xe4, 0x38, 0xc7, 0xf7, 0xe6, 0xbd, 0x9f, 0xbd, 0xa3, 0x01, 0x6e, 0x5e, 0x18, 0x2f, 0xeb, 0xc7,
	0xf5, 0xbe, 0x61, 0x9c, 0xe5, 0xed, 0xcf, 0xf5, 0x1e, 0xff, 0xae, 0x18, 0x26, 0x6b, 0xd6, 0x3c,
	0xe2, 0xba, 0xfc, 0x43, 0x1b, 0x57, 0x8e, 0xd0, 0xa2, 0x1f, 0x58, 0x57, 0xc7, 0x8d, 0x5e, 0xbf,
	0x30, 0x4e, 0xbb, 0xec, 0xf5, 0x5f, 0x15, 0x8c, 0x78, 0xe8, 0xdf, 0x77, 0x3d, 0xf4, 0x15, 0x54,
	0x5c, 0xf0, 0x92, 0xd5, 0xa6, 0x62, 0x2b, 0x8e, 0xbe, 0xb1, 0xdd, 0x9e, 0xe8, 0x1e, 0x47, 0x5d,
	0x4f, 0xe6, 0x92, 0x3e, 0x8f, 0xae, 0x60, 0xf5, 0xdc, 0xe6, 0xbf, 0x4a, 0xce, 0x29, 0xc9, 0x30,
	0x37, 0x2f, 0x6c, 0xc5, 0x99, 0x25, 0xda, 0xe8, 0x79, 0x1c, 0x6d, 0x41, 0x2b, 0x1a, 0x8a, 0x39,
	0xcd, 0xc4, 0xdf, 0x30, 0x67, 0xb6, 0xe2, 0x68, 0x1b, 0xe7, 0xf4, 0x2f, 0xf8, 0x32, 0xfc, 0xc0,
	0x38, 0x0d, 0x30, 0xc7, 0x09, 0x14, 0xa3, 0x16, 0xa8, 0x76, 0x4f, 0x46, 0xd4, 0xfc, 0x1c, 0x6a,
	0x27, 0xc3, 0x07, 0x54, 0x3b, 0x6a, 0x81, 0x22, 0xb4, 0xa2, 0x03, 0xea, 0xfd, 0x39, 0x54, 0x20,
	0xc3, 0x07, 0x14, 0x19, 0x35, 0xfa, 0x06, 0x80, 0x09, 0xc9, 0x72, 0x5c, 0x55, 0x8c, 0x9b, 0xaa,
	0x24, 0xdd, 0xbc, 0xb2, 0x41, 0x42, 0xee, 0x64, 0x54, 0x82, 0x96, 0x78, 0x90, 0x28, 0x82, 0x55,
	0xc1, 0xda, 0x9a, 0x0f, 0xa4, 0x85, 0x24, 0xdd, 0xbe, 0xb2, 0x29, 0x91, 0xfe, 0x8f, 0xa5, 0x15,
	0x07, 0xc3, 0xda, 0x80, 0x3e, 0xdd, 0x24, 0xb2, 0x61, 0x2e, 0xdf, 0xaa, 0x48, 0xee, 0xca, 0xed,
	0x8e, 0xc3, 0x15, 0xf3, 0x44, 0x4e, 0x44, 0x67, 0xba, 0xb2, 0x37, 0x74, 0x6e, 0x41, 0x9f, 0xee,
	0x06, 0x7d, 0x84, 0x85, 0x98, 0x64, 0x25, 0x91, 0xb5, 0x65, 0xa2, 0x0a, 0xb9, 0x25, 0xd6, 0x16,
	0x3e, 0x4c, 0x1e, 0x7f, 0x32, 0x89, 0x3e, 0x81, 0xf6, 0x84, 0x9f, 0x9f, 0x28, 0xc9, 0x0a, 0x46,
	0xa8, 0xbc, 0xaa, 0x65, 0x02, 0x9d, 0xe5, 0x33, 0x42, 0xad, 0xcf, 0x70, 0x79, 0xf4, 0xfa, 0x93,
	0xb0, 0xeb, 0x1f, 0xa0, 0x76, 0x57, 0x8b, 0x2e, 0x41, 0xf3, 0x93, 0xd0, 0x4b, 0xc3, 0xec, 0x21,
	0x4e, 0x43, 0xe3, 0x9d, 0x30, 0x76, 0xf7, 0xc1, 0x68, 0x28, 0xc2, 0x08, 0xc2, 0x28, 0x1c, 0x8c,
	0x0b, 0xa4, 0x03, 0x78, 0x41, 0x90, 0xdd, 0x79, 0x51, 0x14, 0xa7, 0xc6, 0x0c, 0x19, 0xb0, 0xf2,
	0xe3, 0xdd, 0xf7, 0x74, 0x70, 0xe6, 0xb9, 0x2a, 0x3f, 0xac, 0x2f, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0xd7, 0x5e, 0x8c, 0xaf, 0x03, 0x00, 0x00,
}