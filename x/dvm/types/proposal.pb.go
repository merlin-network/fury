// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/dvm/proposal.proto

package types

import (
	fmt "fmt"
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

// ProposalResult is the enum type for the proposal result.
type ProposalResult int32

const (
	// unchosen value
	ProposalResult_PROPOSAL_RESULT_UNSPECIFIED ProposalResult = 0
	// approved
	ProposalResult_PROPOSAL_RESULT_APPROVED ProposalResult = 1
	// rejected
	ProposalResult_PROPOSAL_RESULT_REJECTED ProposalResult = 2
	// expired
	ProposalResult_PROPOSAL_RESULT_EXPIRED ProposalResult = 3
)

var ProposalResult_name = map[int32]string{
	0: "PROPOSAL_RESULT_UNSPECIFIED",
	1: "PROPOSAL_RESULT_APPROVED",
	2: "PROPOSAL_RESULT_REJECTED",
	3: "PROPOSAL_RESULT_EXPIRED",
}

var ProposalResult_value = map[string]int32{
	"PROPOSAL_RESULT_UNSPECIFIED": 0,
	"PROPOSAL_RESULT_APPROVED":    1,
	"PROPOSAL_RESULT_REJECTED":    2,
	"PROPOSAL_RESULT_EXPIRED":     3,
}

func (x ProposalResult) String() string {
	return proto.EnumName(ProposalResult_name, int32(x))
}

func (ProposalResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_29058a11b4a243d6, []int{0}
}

// PublicKeysChangeProposal is the type for the proposal for additions and
// removals of pub keys.
type PublicKeysChangeProposal struct {
	// id is the sequential id of the proposal generated by the blockchain.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// creator is the account address of the proposal creator.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// modifications contain the ticket payload of the proposal.
	Modifications PubkeysChangeProposalPayload `protobuf:"bytes,3,opt,name=modifications,proto3" json:"modifications"`
	// votes contains the votes of the proposal.
	Votes []*Vote `protobuf:"bytes,4,rep,name=votes,proto3" json:"votes,omitempty"`
	// start_ts is the block time that the proposal is set.
	StartTS int64 `protobuf:"varint,5,opt,name=start_ts,proto3" json:"start_ts"`
}

func (m *PublicKeysChangeProposal) Reset()         { *m = PublicKeysChangeProposal{} }
func (m *PublicKeysChangeProposal) String() string { return proto.CompactTextString(m) }
func (*PublicKeysChangeProposal) ProtoMessage()    {}
func (*PublicKeysChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_29058a11b4a243d6, []int{0}
}
func (m *PublicKeysChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKeysChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKeysChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKeysChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeysChangeProposal.Merge(m, src)
}
func (m *PublicKeysChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *PublicKeysChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeysChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeysChangeProposal proto.InternalMessageInfo

func (m *PublicKeysChangeProposal) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PublicKeysChangeProposal) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PublicKeysChangeProposal) GetModifications() PubkeysChangeProposalPayload {
	if m != nil {
		return m.Modifications
	}
	return PubkeysChangeProposalPayload{}
}

func (m *PublicKeysChangeProposal) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *PublicKeysChangeProposal) GetStartTS() int64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

// PublicKeysChangeFinishedProposal is the type used for the finished proposal
// for additions and removals of pub keys.
type PublicKeysChangeFinishedProposal struct {
	// proposal is the proposal for additions and removals of pub keys.
	Proposal PublicKeysChangeProposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
	// result is the result of the finished proposal.
	Result ProposalResult `protobuf:"varint,2,opt,name=result,proto3,enum=merlinnetwork.fury.dvm.ProposalResult" json:"result,omitempty"`
	// result_meta is the metadata related to the result of the finished proposal.
	ResultMeta string `protobuf:"bytes,3,opt,name=result_meta,json=resultMeta,proto3" json:"result_meta,omitempty"`
	// finish_ts is the block time that the proposal is set as finished.
	FinishTS int64 `protobuf:"varint,4,opt,name=finish_ts,proto3" json:"finish_ts"`
}

func (m *PublicKeysChangeFinishedProposal) Reset()         { *m = PublicKeysChangeFinishedProposal{} }
func (m *PublicKeysChangeFinishedProposal) String() string { return proto.CompactTextString(m) }
func (*PublicKeysChangeFinishedProposal) ProtoMessage()    {}
func (*PublicKeysChangeFinishedProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_29058a11b4a243d6, []int{1}
}
func (m *PublicKeysChangeFinishedProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKeysChangeFinishedProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKeysChangeFinishedProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKeysChangeFinishedProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeysChangeFinishedProposal.Merge(m, src)
}
func (m *PublicKeysChangeFinishedProposal) XXX_Size() int {
	return m.Size()
}
func (m *PublicKeysChangeFinishedProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeysChangeFinishedProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeysChangeFinishedProposal proto.InternalMessageInfo

func (m *PublicKeysChangeFinishedProposal) GetProposal() PublicKeysChangeProposal {
	if m != nil {
		return m.Proposal
	}
	return PublicKeysChangeProposal{}
}

func (m *PublicKeysChangeFinishedProposal) GetResult() ProposalResult {
	if m != nil {
		return m.Result
	}
	return ProposalResult_PROPOSAL_RESULT_UNSPECIFIED
}

func (m *PublicKeysChangeFinishedProposal) GetResultMeta() string {
	if m != nil {
		return m.ResultMeta
	}
	return ""
}

func (m *PublicKeysChangeFinishedProposal) GetFinishTS() int64 {
	if m != nil {
		return m.FinishTS
	}
	return 0
}

func init() {
	proto.RegisterEnum("merlinnetwork.fury.dvm.ProposalResult", ProposalResult_name, ProposalResult_value)
	proto.RegisterType((*PublicKeysChangeProposal)(nil), "merlinnetwork.fury.dvm.PublicKeysChangeProposal")
	proto.RegisterType((*PublicKeysChangeFinishedProposal)(nil), "merlinnetwork.fury.dvm.PublicKeysChangeFinishedProposal")
}

func init() { proto.RegisterFile("fury/dvm/proposal.proto", fileDescriptor_29058a11b4a243d6) }

var fileDescriptor_29058a11b4a243d6 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xda, 0x4c,
	0x14, 0xc5, 0x19, 0x43, 0x12, 0x18, 0xf4, 0x21, 0x34, 0x8a, 0xbe, 0xb8, 0x49, 0x65, 0x5b, 0x2c,
	0x2a, 0x54, 0xb5, 0xa6, 0x22, 0x8b, 0x4a, 0x5d, 0x35, 0x80, 0x23, 0xd1, 0xa6, 0x61, 0x34, 0x90,
	0xa8, 0xaa, 0x2a, 0x21, 0x63, 0x4f, 0xcc, 0x08, 0xcc, 0x20, 0xcf, 0x40, 0xcb, 0x13, 0xb4, 0xea,
	0xaa, 0x8f, 0x95, 0x65, 0x96, 0x5d, 0xa1, 0xca, 0xec, 0xfa, 0x14, 0x95, 0xff, 0x40, 0x14, 0x42,
	0x77, 0xd7, 0xe7, 0x1e, 0x9f, 0x3b, 0xf7, 0x37, 0x03, 0xff, 0x17, 0x1e, 0xad, 0xb9, 0x73, 0xbf,
	0x36, 0x0d, 0xf8, 0x94, 0x0b, 0x7b, 0x6c, 0x4e, 0x03, 0x2e, 0x39, 0x42, 0xc2, 0xa3, 0x13, 0x2a,
	0xbf, 0xf0, 0x60, 0x64, 0x0a, 0x8f, 0x9a, 0xee, 0xdc, 0x3f, 0x3e, 0xf4, 0xb8, 0xc7, 0xe3, 0x76,
	0x2d, 0xaa, 0x12, 0xe7, 0xf1, 0xe1, 0x3a, 0x41, 0x32, 0x67, 0x44, 0x65, 0xaa, 0xa2, 0xb5, 0x3a,
	0xe7, 0x92, 0x26, 0x5a, 0xe5, 0x9b, 0x02, 0x55, 0x3c, 0x1b, 0x8c, 0x99, 0xf3, 0x9e, 0x2e, 0x44,
	0x73, 0x68, 0x4f, 0x3c, 0x8a, 0xd3, 0xb1, 0xa8, 0x04, 0x15, 0xe6, 0xaa, 0xc0, 0x00, 0xd5, 0x1c,
	0x51, 0x98, 0x8b, 0x54, 0x78, 0xe0, 0x04, 0xd4, 0x96, 0x3c, 0x50, 0x15, 0x03, 0x54, 0x0b, 0x64,
	0xfd, 0x89, 0x3e, 0xc3, 0xff, 0x7c, 0xee, 0xb2, 0x1b, 0xe6, 0xd8, 0x92, 0xf1, 0x89, 0x50, 0xb3,
	0x06, 0xa8, 0x16, 0xeb, 0xaf, 0xcc, 0xc7, 0x47, 0x36, 0xf1, 0x6c, 0x30, 0x7a, 0x34, 0x0b, 0xdb,
	0x8b, 0x31, 0xb7, 0xdd, 0x46, 0xee, 0x76, 0xa9, 0x67, 0xc8, 0xc3, 0x30, 0x64, 0xc2, 0xbd, 0xe8,
	0xc8, 0x42, 0xcd, 0x19, 0xd9, 0x6a, 0xb1, 0xae, 0xee, 0x4a, 0xbd, 0xe6, 0x92, 0x92, 0xc4, 0x86,
	0x4e, 0x61, 0x5e, 0x48, 0x3b, 0x90, 0x7d, 0x29, 0xd4, 0x3d, 0x03, 0x54, 0xb3, 0x8d, 0xa3, 0x70,
	0xa9, 0x1f, 0x74, 0x23, 0xad, 0xd7, 0xfd, 0xb3, 0xd4, 0x37, 0x6d, 0xb2, 0xa9, 0x2a, 0xdf, 0x15,
	0x68, 0x6c, 0x93, 0x38, 0x67, 0x13, 0x26, 0x86, 0xd4, 0xdd, 0x10, 0xb9, 0x84, 0xf9, 0xf5, 0xa5,
	0xc4, 0x5c, 0x8a, 0xf5, 0x17, 0xff, 0x58, 0x71, 0x27, 0xd1, 0x74, 0xbd, 0x4d, 0x06, 0x7a, 0x03,
	0xf7, 0x03, 0x2a, 0x66, 0x63, 0x19, 0x03, 0x2d, 0xd5, 0x2b, 0x3b, 0xd3, 0x52, 0x37, 0x89, 0x9d,
	0x24, 0xfd, 0x03, 0xe9, 0xb0, 0x98, 0x54, 0x7d, 0x9f, 0x4a, 0x3b, 0x26, 0x5e, 0x20, 0x30, 0x91,
	0x3e, 0x50, 0x69, 0xa3, 0xd7, 0xb0, 0x70, 0x13, 0x2f, 0x10, 0x71, 0xc8, 0xc5, 0x1c, 0x9e, 0x84,
	0x4b, 0x3d, 0x9f, 0x6c, 0x15, 0x83, 0xb8, 0x37, 0x90, 0xfb, 0xf2, 0xf9, 0x0f, 0x00, 0x4b, 0x0f,
	0x87, 0x22, 0x1d, 0x9e, 0x60, 0xd2, 0xc1, 0x9d, 0xee, 0xd9, 0x45, 0x9f, 0x58, 0xdd, 0xab, 0x8b,
	0x5e, 0xff, 0xea, 0xb2, 0x8b, 0xad, 0x66, 0xfb, 0xbc, 0x6d, 0xb5, 0xca, 0x19, 0xf4, 0x14, 0xaa,
	0xdb, 0x86, 0x33, 0x8c, 0x49, 0xe7, 0xda, 0x6a, 0x95, 0xc1, 0xae, 0x2e, 0xb1, 0xde, 0x59, 0xcd,
	0x9e, 0xd5, 0x2a, 0x2b, 0xe8, 0x04, 0x1e, 0x6d, 0x77, 0xad, 0x8f, 0xb8, 0x4d, 0xac, 0x56, 0x39,
	0xdb, 0x78, 0x7b, 0x1b, 0x6a, 0xe0, 0x2e, 0xd4, 0xc0, 0xef, 0x50, 0x03, 0x3f, 0x57, 0x5a, 0xe6,
	0x6e, 0xa5, 0x65, 0x7e, 0xad, 0xb4, 0xcc, 0xa7, 0x67, 0x1e, 0x93, 0xc3, 0xd9, 0xc0, 0x74, 0xb8,
	0x5f, 0x13, 0x1e, 0x7d, 0x99, 0x72, 0x8b, 0xea, 0xda, 0xd7, 0xe4, 0xf9, 0x2f, 0xa6, 0x54, 0x0c,
	0xf6, 0xe3, 0xa7, 0x7e, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xd8, 0xe1, 0x28, 0x58, 0x03,
	0x00, 0x00,
}

func (m *PublicKeysChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKeysChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKeysChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartTS != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.Modifications.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PublicKeysChangeFinishedProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKeysChangeFinishedProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKeysChangeFinishedProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinishTS != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.FinishTS))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ResultMeta) > 0 {
		i -= len(m.ResultMeta)
		copy(dAtA[i:], m.ResultMeta)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ResultMeta)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Proposal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublicKeysChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProposal(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.Modifications.Size()
	n += 1 + l + sovProposal(uint64(l))
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if m.StartTS != 0 {
		n += 1 + sovProposal(uint64(m.StartTS))
	}
	return n
}

func (m *PublicKeysChangeFinishedProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Proposal.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.Result != 0 {
		n += 1 + sovProposal(uint64(m.Result))
	}
	l = len(m.ResultMeta)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.FinishTS != 0 {
		n += 1 + sovProposal(uint64(m.FinishTS))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicKeysChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: PublicKeysChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKeysChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modifications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Modifications.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *PublicKeysChangeFinishedProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: PublicKeysChangeFinishedProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKeysChangeFinishedProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proposal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= ProposalResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultMeta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultMeta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishTS", wireType)
			}
			m.FinishTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinishTS |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
