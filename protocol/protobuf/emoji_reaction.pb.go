// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emoji_reaction.proto

package protobuf

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

type EmojiReaction_Type int32

const (
	EmojiReaction_UNKNOWN_EMOJI_REACTION_TYPE EmojiReaction_Type = 0
	EmojiReaction_LOVE                        EmojiReaction_Type = 1
	EmojiReaction_THUMBS_UP                   EmojiReaction_Type = 2
	EmojiReaction_THUMBS_DOWN                 EmojiReaction_Type = 3
	EmojiReaction_LAUGH                       EmojiReaction_Type = 4
	EmojiReaction_SAD                         EmojiReaction_Type = 5
	EmojiReaction_ANGRY                       EmojiReaction_Type = 6
)

var EmojiReaction_Type_name = map[int32]string{
	0: "UNKNOWN_EMOJI_REACTION_TYPE",
	1: "LOVE",
	2: "THUMBS_UP",
	3: "THUMBS_DOWN",
	4: "LAUGH",
	5: "SAD",
	6: "ANGRY",
}

var EmojiReaction_Type_value = map[string]int32{
	"UNKNOWN_EMOJI_REACTION_TYPE": 0,
	"LOVE":                        1,
	"THUMBS_UP":                   2,
	"THUMBS_DOWN":                 3,
	"LAUGH":                       4,
	"SAD":                         5,
	"ANGRY":                       6,
}

func (x EmojiReaction_Type) String() string {
	return proto.EnumName(EmojiReaction_Type_name, int32(x))
}

func (EmojiReaction_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0a088c907bbc7ed6, []int{0, 0}
}

type EmojiReaction struct {
	// clock Lamport timestamp of the chat message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// chat_id the ID of the chat the message belongs to, for query efficiency the chat_id is stored in the db even though the
	// target message also stores the chat_id
	ChatId string `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// message_id the ID of the target message that the user wishes to react to
	MessageId string `protobuf:"bytes,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// message_type is (somewhat confusingly) the ID of the type of chat the message belongs to
	MessageType MessageType `protobuf:"varint,4,opt,name=message_type,json=messageType,proto3,enum=protobuf.MessageType" json:"message_type,omitempty"`
	// type the ID of the emoji the user wishes to react with
	Type EmojiReaction_Type `protobuf:"varint,5,opt,name=type,proto3,enum=protobuf.EmojiReaction_Type" json:"type,omitempty"`
	// whether this is a rectraction of a previously sent emoji
	Retracted bool `protobuf:"varint,6,opt,name=retracted,proto3" json:"retracted,omitempty"`
	// Grant for organisation chat messages
	Grant                []byte   `protobuf:"bytes,7,opt,name=grant,proto3" json:"grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojiReaction) Reset()         { *m = EmojiReaction{} }
func (m *EmojiReaction) String() string { return proto.CompactTextString(m) }
func (*EmojiReaction) ProtoMessage()    {}
func (*EmojiReaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a088c907bbc7ed6, []int{0}
}

func (m *EmojiReaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojiReaction.Unmarshal(m, b)
}
func (m *EmojiReaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojiReaction.Marshal(b, m, deterministic)
}
func (m *EmojiReaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojiReaction.Merge(m, src)
}
func (m *EmojiReaction) XXX_Size() int {
	return xxx_messageInfo_EmojiReaction.Size(m)
}
func (m *EmojiReaction) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojiReaction.DiscardUnknown(m)
}

var xxx_messageInfo_EmojiReaction proto.InternalMessageInfo

func (m *EmojiReaction) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *EmojiReaction) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *EmojiReaction) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *EmojiReaction) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *EmojiReaction) GetType() EmojiReaction_Type {
	if m != nil {
		return m.Type
	}
	return EmojiReaction_UNKNOWN_EMOJI_REACTION_TYPE
}

func (m *EmojiReaction) GetRetracted() bool {
	if m != nil {
		return m.Retracted
	}
	return false
}

func (m *EmojiReaction) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.EmojiReaction_Type", EmojiReaction_Type_name, EmojiReaction_Type_value)
	proto.RegisterType((*EmojiReaction)(nil), "protobuf.EmojiReaction")
}

func init() {
	proto.RegisterFile("emoji_reaction.proto", fileDescriptor_0a088c907bbc7ed6)
}

var fileDescriptor_0a088c907bbc7ed6 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xcd, 0x6e, 0xaa, 0x40,
	0x14, 0xc7, 0x2f, 0x0a, 0x28, 0x07, 0xbd, 0x77, 0x32, 0xf1, 0xa6, 0xa4, 0xb5, 0x29, 0x71, 0xc5,
	0x8a, 0x36, 0xed, 0xa6, 0x49, 0x57, 0x58, 0x89, 0xd2, 0x2a, 0x98, 0x11, 0x6a, 0xec, 0x86, 0x20,
	0x4c, 0xad, 0x6d, 0x11, 0x82, 0xe3, 0xc2, 0xa7, 0xee, 0x2b, 0x34, 0x0c, 0x1a, 0xd3, 0xd5, 0xcc,
	0xef, 0xff, 0x91, 0x73, 0x0e, 0x74, 0x68, 0x9a, 0x7d, 0xac, 0xc3, 0x82, 0x46, 0x31, 0x5b, 0x67,
	0x1b, 0x33, 0x2f, 0x32, 0x96, 0xe1, 0x26, 0x7f, 0x96, 0xbb, 0xb7, 0x73, 0x95, 0x6e, 0x76, 0xe9,
	0xb6, 0x92, 0x7b, 0xdf, 0x35, 0x68, 0xdb, 0x65, 0x9e, 0x1c, 0xe2, 0xb8, 0x03, 0x52, 0xfc, 0x95,
	0xc5, 0x9f, 0x9a, 0xa0, 0x0b, 0x86, 0x48, 0x2a, 0xc0, 0x67, 0xd0, 0x88, 0xdf, 0x23, 0x16, 0xae,
	0x13, 0xad, 0xa6, 0x0b, 0x86, 0x42, 0xe4, 0x12, 0x9d, 0x04, 0x5f, 0x02, 0xa4, 0x74, 0xbb, 0x8d,
	0x56, 0xb4, 0xf4, 0xea, 0xdc, 0x53, 0x0e, 0x8a, 0x93, 0xe0, 0x7b, 0x68, 0x1d, 0x6d, 0xb6, 0xcf,
	0xa9, 0x26, 0xea, 0x82, 0xf1, 0xf7, 0xf6, 0xbf, 0x79, 0xdc, 0xc6, 0x9c, 0x54, 0xae, 0xbf, 0xcf,
	0x29, 0x51, 0xd3, 0x13, 0xe0, 0x1b, 0x10, 0x79, 0x43, 0xe2, 0x8d, 0xee, 0xa9, 0xf1, 0x6b, 0x5d,
	0x93, 0x17, 0x79, 0x12, 0x77, 0x41, 0x29, 0x28, 0x2b, 0xa2, 0x98, 0xd1, 0x44, 0x93, 0x75, 0xc1,
	0x68, 0x92, 0x93, 0x50, 0xde, 0xb5, 0x2a, 0xa2, 0x0d, 0xd3, 0x1a, 0xba, 0x60, 0xb4, 0x48, 0x05,
	0xbd, 0x1c, 0x44, 0x3e, 0xed, 0x0a, 0x2e, 0x02, 0xf7, 0xd9, 0xf5, 0xe6, 0x6e, 0x68, 0x4f, 0xbc,
	0x27, 0x27, 0x24, 0xb6, 0xf5, 0xe8, 0x3b, 0x9e, 0x1b, 0xfa, 0x8b, 0xa9, 0x8d, 0xfe, 0xe0, 0x26,
	0x88, 0x63, 0xef, 0xc5, 0x46, 0x02, 0x6e, 0x83, 0xe2, 0x8f, 0x82, 0x49, 0x7f, 0x16, 0x06, 0x53,
	0x54, 0xc3, 0xff, 0x40, 0x3d, 0xe0, 0xc0, 0x9b, 0xbb, 0xa8, 0x8e, 0x15, 0x90, 0xc6, 0x56, 0x30,
	0x1c, 0x21, 0x11, 0x37, 0xa0, 0x3e, 0xb3, 0x06, 0x48, 0x2a, 0x35, 0xcb, 0x1d, 0x92, 0x05, 0x92,
	0xfb, 0xed, 0x57, 0xd5, 0xbc, 0x7e, 0x38, 0x5e, 0xb3, 0x94, 0xf9, 0xef, 0xee, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x57, 0x12, 0xd9, 0xb6, 0x01, 0x00, 0x00,
}