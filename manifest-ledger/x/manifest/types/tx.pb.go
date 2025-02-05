// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: liftedinit/manifest/v1/tx.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgPayout is the Msg/Payout request type.
type MsgPayout struct {
	// authority is the address of the controlling account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// payout_pairs are the pairs of addresses and coins to be paid out.
	PayoutPairs []PayoutPair `protobuf:"bytes,2,rep,name=payout_pairs,json=payoutPairs,proto3" json:"payout_pairs"`
}

func (m *MsgPayout) Reset()         { *m = MsgPayout{} }
func (m *MsgPayout) String() string { return proto.CompactTextString(m) }
func (*MsgPayout) ProtoMessage()    {}
func (*MsgPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c6d73bffbe53f, []int{0}
}
func (m *MsgPayout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPayout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayout.Merge(m, src)
}
func (m *MsgPayout) XXX_Size() int {
	return m.Size()
}
func (m *MsgPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayout.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayout proto.InternalMessageInfo

func (m *MsgPayout) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgPayout) GetPayoutPairs() []PayoutPair {
	if m != nil {
		return m.PayoutPairs
	}
	return nil
}

// PayoutPair is the object that pairs an address with a coin to be paid out.
type PayoutPair struct {
	Address string                                  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=coin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *PayoutPair) Reset()         { *m = PayoutPair{} }
func (m *PayoutPair) String() string { return proto.CompactTextString(m) }
func (*PayoutPair) ProtoMessage()    {}
func (*PayoutPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c6d73bffbe53f, []int{1}
}
func (m *PayoutPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayoutPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayoutPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PayoutPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayoutPair.Merge(m, src)
}
func (m *PayoutPair) XXX_Size() int {
	return m.Size()
}
func (m *PayoutPair) XXX_DiscardUnknown() {
	xxx_messageInfo_PayoutPair.DiscardUnknown(m)
}

var xxx_messageInfo_PayoutPair proto.InternalMessageInfo

func (m *PayoutPair) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PayoutPair) GetCoin() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Coin
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

// MsgPayout is the Msg/BurnHeldBalance request type.
type MsgBurnHeldBalance struct {
	// sender is the address of the tokenholder.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// burn_coins are the coins to be burned by the tokenholder.
	BurnCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=burn_coins,json=burnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burn_coins"`
}

func (m *MsgBurnHeldBalance) Reset()         { *m = MsgBurnHeldBalance{} }
func (m *MsgBurnHeldBalance) String() string { return proto.CompactTextString(m) }
func (*MsgBurnHeldBalance) ProtoMessage()    {}
func (*MsgBurnHeldBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c6d73bffbe53f, []int{2}
}
func (m *MsgBurnHeldBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnHeldBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnHeldBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnHeldBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnHeldBalance.Merge(m, src)
}
func (m *MsgBurnHeldBalance) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnHeldBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnHeldBalance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnHeldBalance proto.InternalMessageInfo

func (m *MsgBurnHeldBalance) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgBurnHeldBalance) GetBurnCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BurnCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgPayout)(nil), "liftedinit.manifest.v1.MsgPayout")
	proto.RegisterType((*PayoutPair)(nil), "liftedinit.manifest.v1.PayoutPair")
	proto.RegisterType((*MsgBurnHeldBalance)(nil), "liftedinit.manifest.v1.MsgBurnHeldBalance")
}

func init() { proto.RegisterFile("liftedinit/manifest/v1/tx.proto", fileDescriptor_b20c6d73bffbe53f) }

var fileDescriptor_b20c6d73bffbe53f = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x7d, 0x01, 0x81, 0x7c, 0xe9, 0x82, 0x55, 0x81, 0x53, 0x24, 0x3b, 0xf2, 0x42, 0x14,
	0x29, 0x3e, 0x52, 0xa4, 0x4a, 0x64, 0xc3, 0x20, 0xc4, 0x52, 0x29, 0x0a, 0x1b, 0x4b, 0x74, 0xb6,
	0xaf, 0xd7, 0x13, 0xb6, 0xcf, 0xf2, 0x9d, 0xa3, 0x66, 0x43, 0x8c, 0x9d, 0x98, 0x99, 0x18, 0x11,
	0x53, 0x06, 0x56, 0x66, 0x3a, 0x56, 0x4c, 0x4c, 0x05, 0x25, 0x43, 0xf8, 0x1b, 0x98, 0xd0, 0xf9,
	0x9c, 0xba, 0xaa, 0x22, 0x84, 0x58, 0xec, 0xbb, 0x7b, 0xef, 0xde, 0xfb, 0x7e, 0xdf, 0xb3, 0xa1,
	0x9b, 0xb0, 0x23, 0x49, 0x62, 0x96, 0x31, 0x89, 0x52, 0x9c, 0xb1, 0x23, 0x22, 0x24, 0x9a, 0x0d,
	0x91, 0x3c, 0xf1, 0xf3, 0x82, 0x4b, 0x6e, 0xdd, 0x6d, 0x12, 0xfc, 0x4d, 0x82, 0x3f, 0x1b, 0xee,
	0xdd, 0x8b, 0xb8, 0x48, 0xb9, 0x40, 0xa9, 0xa0, 0x2a, 0x3f, 0x15, 0x54, 0x5f, 0xd8, 0xdb, 0xa5,
	0x9c, 0xf2, 0x6a, 0x89, 0xd4, 0xaa, 0x3e, 0xed, 0xe8, 0xf4, 0xa9, 0x0e, 0xe8, 0x4d, 0x1d, 0x72,
	0xea, 0x4a, 0x21, 0x16, 0x04, 0xcd, 0x86, 0x21, 0x91, 0x78, 0x88, 0x22, 0xce, 0xb2, 0x3a, 0x7e,
	0x07, 0xa7, 0x2c, 0xe3, 0xa8, 0x7a, 0xea, 0x23, 0xef, 0x2b, 0x80, 0xe6, 0xa1, 0xa0, 0x63, 0x3c,
	0xe7, 0xa5, 0xb4, 0x0e, 0xa0, 0x89, 0x4b, 0x79, 0xcc, 0x0b, 0x26, 0xe7, 0x36, 0xe8, 0x82, 0x9e,
	0x19, 0xd8, 0xdf, 0x3e, 0x0f, 0x76, 0xeb, 0x2e, 0x4f, 0xe2, 0xb8, 0x20, 0x42, 0xbc, 0x94, 0x05,
	0xcb, 0xe8, 0xa4, 0x49, 0xb5, 0xc6, 0x70, 0x27, 0xaf, 0x2a, 0x4c, 0x73, 0xcc, 0x0a, 0x61, 0xb7,
	0xba, 0x37, 0x7a, 0xed, 0x7d, 0xcf, 0xdf, 0x4e, 0xec, 0xeb, 0x6e, 0x63, 0xcc, 0x8a, 0xc0, 0x3c,
	0xbb, 0x70, 0x8d, 0x8f, 0xeb, 0x45, 0x1f, 0x4c, 0xda, 0xf9, 0xe5, 0xb1, 0x18, 0x3d, 0xfc, 0xf5,
	0xc1, 0x35, 0xde, 0xae, 0x17, 0xfd, 0xa6, 0xcb, 0xe9, 0x7a, 0xd1, 0xef, 0xe8, 0x8a, 0x8d, 0xc1,
	0x97, 0xda, 0xbd, 0x2f, 0x00, 0xc2, 0xa6, 0xb0, 0x65, 0xc3, 0xdb, 0x58, 0xcb, 0xd5, 0x20, 0x93,
	0xcd, 0xd6, 0x12, 0xf0, 0xa6, 0xf2, 0xc4, 0x6e, 0x75, 0x41, 0xaf, 0xbd, 0xdf, 0xf1, 0x6b, 0x38,
	0x65, 0x9a, 0x5f, 0x9b, 0xe6, 0x3f, 0xe5, 0x2c, 0x0b, 0x9e, 0x29, 0x6d, 0xbf, 0x2f, 0xdc, 0x07,
	0x94, 0xc9, 0xe3, 0x32, 0xf4, 0x23, 0x9e, 0xd6, 0x7e, 0xd7, 0xaf, 0x81, 0x88, 0x5f, 0x23, 0x39,
	0xcf, 0x89, 0xa8, 0x2e, 0xbc, 0x5f, 0x2f, 0xfa, 0xed, 0x84, 0x50, 0x1c, 0xcd, 0xa7, 0xaa, 0x83,
	0xc6, 0xaa, 0x9a, 0x8d, 0xba, 0x4a, 0xfd, 0xfd, 0xeb, 0xea, 0x35, 0xf1, 0x40, 0x99, 0xe6, 0x9d,
	0xb6, 0xa0, 0x75, 0x28, 0x68, 0x50, 0x16, 0xd9, 0x0b, 0x92, 0xc4, 0x01, 0x4e, 0x70, 0x16, 0x91,
	0xff, 0x1e, 0xc9, 0x1b, 0x00, 0x61, 0x58, 0x16, 0x59, 0xa5, 0x64, 0x33, 0x91, 0xbf, 0xc0, 0x3e,
	0x57, 0xb0, 0x9f, 0x7e, 0xb8, 0xbd, 0x7f, 0x84, 0x15, 0x8a, 0x76, 0xe7, 0x0a, 0xad, 0xd0, 0xb8,
	0xa6, 0x6a, 0x5a, 0x25, 0x8c, 0x1e, 0x6f, 0x9f, 0xa1, 0xb7, 0x65, 0x86, 0xd7, 0xa8, 0x83, 0xf1,
	0xd9, 0xd2, 0x01, 0xe7, 0x4b, 0x07, 0xfc, 0x5c, 0x3a, 0xe0, 0xdd, 0xca, 0x31, 0xce, 0x57, 0x8e,
	0xf1, 0x7d, 0xe5, 0x18, 0xaf, 0x0e, 0xae, 0xe8, 0xdb, 0xf2, 0xc7, 0x0d, 0x12, 0x12, 0x53, 0x52,
	0xa0, 0x93, 0xa6, 0x7c, 0xa5, 0x39, 0xbc, 0x55, 0x7d, 0xef, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x93, 0x14, 0x38, 0x3e, 0xa7, 0x03, 0x00, 0x00,
}

func (m *MsgPayout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPayout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPayout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PayoutPairs) > 0 {
		for iNdEx := len(m.PayoutPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PayoutPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PayoutPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayoutPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PayoutPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnHeldBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnHeldBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnHeldBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnCoins) > 0 {
		for iNdEx := len(m.BurnCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPayout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.PayoutPairs) > 0 {
		for _, e := range m.PayoutPairs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *PayoutPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurnHeldBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.BurnCoins) > 0 {
		for _, e := range m.BurnCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPayout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPayout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPayout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayoutPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayoutPairs = append(m.PayoutPairs, PayoutPair{})
			if err := m.PayoutPairs[len(m.PayoutPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *PayoutPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: PayoutPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayoutPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBurnHeldBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnHeldBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnHeldBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnCoins = append(m.BurnCoins, types.Coin{})
			if err := m.BurnCoins[len(m.BurnCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
