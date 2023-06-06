// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/sportevent/ticket.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SportEventAddTicketPayload indicates data of add sport-event ticket
type SportEventAddTicketPayload struct {
	// uid is the universal unique identifier of the sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the sport-event.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// status is the current status of the sport-event.
	Status SportEventStatus `protobuf:"varint,5,opt,name=status,proto3,enum=fanfury.fury.sportevent.SportEventStatus" json:"status,omitempty"`
	// creator is the address of the creator of sport-event.
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// meta contains human-readable metadata of the sport-event.
	Meta string `protobuf:"bytes,9,opt,name=meta,proto3" json:"meta,omitempty"`
	// sr_contribution_for_house is the portion of payout that should be paid by
	// sr
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
}

func (m *SportEventAddTicketPayload) Reset()         { *m = SportEventAddTicketPayload{} }
func (m *SportEventAddTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventAddTicketPayload) ProtoMessage()    {}
func (*SportEventAddTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{0}
}
func (m *SportEventAddTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventAddTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventAddTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventAddTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventAddTicketPayload.Merge(m, src)
}
func (m *SportEventAddTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventAddTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventAddTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventAddTicketPayload proto.InternalMessageInfo

func (m *SportEventAddTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *SportEventAddTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func (m *SportEventAddTicketPayload) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// SportEventUpdateTicketPayload indicates data of update sport-event ticket
type SportEventUpdateTicketPayload struct {
	// uid is the uuid of the sport-event
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// status is the status of the resolution.
	Status SportEventStatus `protobuf:"varint,6,opt,name=status,proto3,enum=fanfury.fury.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *SportEventUpdateTicketPayload) Reset()         { *m = SportEventUpdateTicketPayload{} }
func (m *SportEventUpdateTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventUpdateTicketPayload) ProtoMessage()    {}
func (*SportEventUpdateTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{1}
}
func (m *SportEventUpdateTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventUpdateTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventUpdateTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventUpdateTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventUpdateTicketPayload.Merge(m, src)
}
func (m *SportEventUpdateTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventUpdateTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventUpdateTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventUpdateTicketPayload proto.InternalMessageInfo

func (m *SportEventUpdateTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventUpdateTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

// SportEventResolutionTicketPayload indicates data of the
// resolution of the sport-event ticket.
type SportEventResolutionTicketPayload struct {
	// uid is the universal unique identifier of sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// resolution_ts is the resolution timestamp of the sport-event.
	ResolutionTS uint64 `protobuf:"varint,2,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// winner_odds_uids is the universal unique identifier list of the winner
	// odds.
	WinnerOddsUIDs []string `protobuf:"bytes,3,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the status of the resolution.
	Status SportEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=fanfury.fury.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *SportEventResolutionTicketPayload) Reset()         { *m = SportEventResolutionTicketPayload{} }
func (m *SportEventResolutionTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventResolutionTicketPayload) ProtoMessage()    {}
func (*SportEventResolutionTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{2}
}
func (m *SportEventResolutionTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventResolutionTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventResolutionTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventResolutionTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventResolutionTicketPayload.Merge(m, src)
}
func (m *SportEventResolutionTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventResolutionTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventResolutionTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventResolutionTicketPayload proto.InternalMessageInfo

func (m *SportEventResolutionTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventResolutionTicketPayload) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *SportEventResolutionTicketPayload) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *SportEventResolutionTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SportEventAddTicketPayload)(nil), "fanfury.fury.sportevent.SportEventAddTicketPayload")
	proto.RegisterType((*SportEventUpdateTicketPayload)(nil), "fanfury.fury.sportevent.SportEventUpdateTicketPayload")
	proto.RegisterType((*SportEventResolutionTicketPayload)(nil), "fanfury.fury.sportevent.SportEventResolutionTicketPayload")
}

func init() { proto.RegisterFile("fury/sportevent/ticket.proto", fileDescriptor_b7e970e0fce71fcc) }

var fileDescriptor_b7e970e0fce71fcc = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcb, 0x4e, 0x1b, 0x31,
	0x14, 0xcd, 0x90, 0x17, 0xb8, 0x34, 0xaa, 0xdc, 0xaa, 0x1d, 0xa8, 0x1a, 0x07, 0x16, 0x28, 0x12,
	0x62, 0x22, 0xc1, 0x17, 0x10, 0x5e, 0x65, 0xd5, 0xca, 0x01, 0x55, 0xea, 0x66, 0x34, 0x89, 0x2f,
	0xc3, 0x08, 0x62, 0x47, 0xb6, 0xa7, 0x94, 0x2f, 0xe8, 0xaa, 0x52, 0xff, 0xa1, 0x8b, 0xfe, 0x0a,
	0x4b, 0x96, 0x55, 0x17, 0x56, 0x35, 0xec, 0xf2, 0x15, 0x95, 0xcd, 0x94, 0x09, 0x20, 0x2a, 0x41,
	0xbb, 0xe8, 0x26, 0xbe, 0xbe, 0xf7, 0x9c, 0x63, 0xfb, 0x1e, 0xc7, 0x83, 0x5e, 0xaa, 0x18, 0x3a,
	0x6a, 0x24, 0xa4, 0x86, 0x0f, 0xc0, 0x75, 0x47, 0x27, 0x83, 0x23, 0xd0, 0xc1, 0x48, 0x0a, 0x2d,
	0xf0, 0x9c, 0x8a, 0x81, 0x83, 0x3e, 0x11, 0xf2, 0x28, 0x50, 0x31, 0x04, 0x05, 0x6e, 0xbe, 0x75,
	0x83, 0xe7, 0xc2, 0xd0, 0xc5, 0x97, 0xe4, 0xf9, 0xb9, 0x1b, 0x08, 0xc1, 0x98, 0xca, 0x4b, 0xcf,
	0x62, 0x11, 0x0b, 0x17, 0x76, 0x6c, 0x74, 0x99, 0x5d, 0xfc, 0x5c, 0x45, 0xf3, 0x3d, 0x8b, 0xdf,
	0xb2, 0xf8, 0x75, 0xc6, 0xf6, 0xdc, 0x5e, 0xde, 0x46, 0xa7, 0xc7, 0x22, 0x62, 0xb8, 0x85, 0xca,
	0x69, 0xc2, 0x7c, 0xaf, 0xe5, 0xb5, 0x67, 0xba, 0x8d, 0xcc, 0x90, 0xf2, 0xfe, 0xee, 0xe6, 0xd8,
	0x10, 0x9b, 0xa5, 0xf6, 0x07, 0xaf, 0xa1, 0x69, 0xa5, 0x23, 0xa9, 0x43, 0xad, 0xfc, 0xa9, 0x96,
	0xd7, 0xae, 0x74, 0x5f, 0x64, 0x86, 0xd4, 0x7b, 0x36, 0xb7, 0xd7, 0x1b, 0x1b, 0x72, 0x55, 0xa6,
	0x57, 0x11, 0x5e, 0x46, 0x35, 0xe0, 0xcc, 0x52, 0xca, 0x8e, 0xf2, 0x34, 0x33, 0xa4, 0xba, 0xc5,
	0x99, 0x23, 0xe4, 0x25, 0x9a, 0x8f, 0x78, 0x0d, 0x55, 0xec, 0x31, 0xfc, 0x4a, 0xab, 0xdc, 0x7e,
	0xb4, 0x4a, 0x82, 0x3b, 0xfb, 0x13, 0xbc, 0x61, 0x4c, 0x51, 0x07, 0xc6, 0x1b, 0xa8, 0xa6, 0x74,
	0xa4, 0x53, 0xe5, 0x57, 0x5b, 0x5e, 0xbb, 0xb1, 0xba, 0xfc, 0x07, 0x5a, 0x71, 0xfe, 0x9e, 0xa3,
	0xd0, 0x9c, 0x8a, 0x7d, 0x54, 0x1f, 0x48, 0x88, 0xb4, 0x90, 0x7e, 0xcd, 0x76, 0x80, 0xfe, 0x9e,
	0xe2, 0x3d, 0xd4, 0x18, 0x26, 0x3c, 0xec, 0x83, 0x0e, 0xa3, 0xa1, 0x48, 0xb9, 0xf6, 0xeb, 0xae,
	0x45, 0xc1, 0x99, 0x21, 0xa5, 0x1f, 0x86, 0x2c, 0xc5, 0x89, 0x3e, 0x4c, 0xfb, 0xc1, 0x40, 0x0c,
	0x3b, 0x03, 0xa1, 0x86, 0x42, 0xe5, 0xc3, 0x8a, 0x62, 0x47, 0x1d, 0x7d, 0x3a, 0x02, 0x15, 0xec,
	0x72, 0x4d, 0x67, 0x87, 0x09, 0xef, 0x82, 0x5e, 0x77, 0x1a, 0x78, 0x07, 0xd5, 0xad, 0xe2, 0x01,
	0x80, 0x3f, 0xfd, 0x20, 0xb9, 0x5a, 0x1f, 0xf4, 0x36, 0x00, 0xc6, 0xa8, 0x32, 0x04, 0x1d, 0xf9,
	0x33, 0x6e, 0xd7, 0x2e, 0xc6, 0x5f, 0x3d, 0x34, 0xa7, 0x64, 0x38, 0x10, 0x5c, 0xcb, 0xa4, 0x9f,
	0xea, 0x44, 0xf0, 0xf0, 0x40, 0xc8, 0xf0, 0x50, 0xa4, 0x0a, 0x7c, 0xe4, 0xd6, 0x83, 0xfb, 0xad,
	0x97, 0x19, 0xf2, 0xbc, 0x27, 0x37, 0x26, 0x14, 0xb7, 0x85, 0x7c, 0x6d, 0xf5, 0xc6, 0x86, 0xdc,
	0xbd, 0x18, 0xbd, 0xbb, 0xb4, 0xf8, 0xa9, 0x8c, 0x5e, 0x15, 0x7e, 0xec, 0x8f, 0x58, 0xa4, 0xe1,
	0xff, 0xbb, 0x92, 0xb7, 0xed, 0xaf, 0xfc, 0x5b, 0xfb, 0xab, 0x7f, 0x65, 0x7f, 0x71, 0xf9, 0x6b,
	0x0f, 0xbe, 0xfc, 0x8b, 0xdf, 0xa6, 0xd0, 0x42, 0x51, 0xa4, 0xa0, 0xc4, 0xb1, 0x33, 0xeb, 0xbe,
	0x6e, 0xec, 0xa0, 0xc7, 0xf2, 0x8a, 0x5c, 0x58, 0xb2, 0x90, 0x19, 0x32, 0x3b, 0xa1, 0x6a, 0xdb,
	0x7c, 0x1d, 0x48, 0xaf, 0x4f, 0x31, 0x45, 0x4f, 0x4e, 0x12, 0xce, 0x41, 0x86, 0xf6, 0x1f, 0x1e,
	0xa6, 0x09, 0xb3, 0x5e, 0x95, 0xdb, 0x33, 0xdd, 0xa5, 0xcc, 0x90, 0xc6, 0x3b, 0x57, 0xb3, 0x4f,
	0xc0, 0xfe, 0xee, 0xa6, 0x1a, 0x1b, 0x72, 0x0b, 0x4d, 0x6f, 0x65, 0x26, 0x3a, 0x55, 0x79, 0x70,
	0xa7, 0xba, 0x3b, 0x67, 0x59, 0xd3, 0x3b, 0xcf, 0x9a, 0xde, 0xcf, 0xac, 0xe9, 0x7d, 0xb9, 0x68,
	0x96, 0xce, 0x2f, 0x9a, 0xa5, 0xef, 0x17, 0xcd, 0xd2, 0xfb, 0x95, 0x09, 0xe3, 0x54, 0x0c, 0x2b,
	0xb9, 0xb2, 0x8d, 0x3b, 0x1f, 0xaf, 0x7d, 0x01, 0xac, 0x87, 0xfd, 0x9a, 0x7b, 0x93, 0xd7, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x7b, 0x43, 0xf8, 0x20, 0x06, 0x00, 0x00,
}

func (m *SportEventAddTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventAddTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventAddTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x4a
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventUpdateTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventUpdateTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventUpdateTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventResolutionTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventResolutionTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventResolutionTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SportEventAddTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func (m *SportEventUpdateTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func (m *SportEventResolutionTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovTicket(uint64(m.ResolutionTS))
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SportEventAddTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SportEventAddTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventAddTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Odds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Odds = append(m.Odds, &Odds{})
			if err := m.Odds[len(m.Odds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SportEventUpdateTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SportEventResolutionTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolutionTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
