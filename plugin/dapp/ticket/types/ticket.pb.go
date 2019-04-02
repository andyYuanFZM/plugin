// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ticket.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	ticket.proto

It has these top-level messages:
	Ticket
	TicketAction
	TicketMiner
	TicketMinerOld
	MinerFlag
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ReqBindMiner
	ReplyBindMiner
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types1 "github.com/33cn/chain33/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Ticket struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	// 0 -> 未成熟 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,3,opt,name=isGenesis" json:"isGenesis,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime" json:"createTime,omitempty"`
	// 挖矿时间
	MinerTime int64 `protobuf:"varint,5,opt,name=minerTime" json:"minerTime,omitempty"`
	// 挖到的币的数目
	MinerValue   int64  `protobuf:"varint,8,opt,name=minerValue" json:"minerValue,omitempty"`
	MinerAddress string `protobuf:"bytes,6,opt,name=minerAddress" json:"minerAddress,omitempty"`
	// return wallet
	ReturnAddress string `protobuf:"bytes,7,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *Ticket) Reset()                    { *m = Ticket{} }
func (m *Ticket) String() string            { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()               {}
func (*Ticket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ticket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *Ticket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Ticket) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *Ticket) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Ticket) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *Ticket) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *Ticket) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *Ticket) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

// message for execs.ticket
type TicketAction struct {
	// Types that are valid to be assigned to Value:
	//	*TicketAction_Tbind
	//	*TicketAction_Topen
	//	*TicketAction_Genesis
	//	*TicketAction_Tclose
	//	*TicketAction_Miner
	Value isTicketAction_Value `protobuf_oneof:"value"`
	Ty    int32                `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *TicketAction) Reset()                    { *m = TicketAction{} }
func (m *TicketAction) String() string            { return proto.CompactTextString(m) }
func (*TicketAction) ProtoMessage()               {}
func (*TicketAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTicketAction_Value interface {
	isTicketAction_Value()
}

type TicketAction_Tbind struct {
	Tbind *TicketBind `protobuf:"bytes,5,opt,name=tbind,oneof"`
}
type TicketAction_Topen struct {
	Topen *TicketOpen `protobuf:"bytes,1,opt,name=topen,oneof"`
}
type TicketAction_Genesis struct {
	Genesis *TicketGenesis `protobuf:"bytes,2,opt,name=genesis,oneof"`
}
type TicketAction_Tclose struct {
	Tclose *TicketClose `protobuf:"bytes,3,opt,name=tclose,oneof"`
}
type TicketAction_Miner struct {
	Miner *TicketMiner `protobuf:"bytes,4,opt,name=miner,oneof"`
}

func (*TicketAction_Tbind) isTicketAction_Value()   {}
func (*TicketAction_Topen) isTicketAction_Value()   {}
func (*TicketAction_Genesis) isTicketAction_Value() {}
func (*TicketAction_Tclose) isTicketAction_Value()  {}
func (*TicketAction_Miner) isTicketAction_Value()   {}

func (m *TicketAction) GetValue() isTicketAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TicketAction) GetTbind() *TicketBind {
	if x, ok := m.GetValue().(*TicketAction_Tbind); ok {
		return x.Tbind
	}
	return nil
}

func (m *TicketAction) GetTopen() *TicketOpen {
	if x, ok := m.GetValue().(*TicketAction_Topen); ok {
		return x.Topen
	}
	return nil
}

func (m *TicketAction) GetGenesis() *TicketGenesis {
	if x, ok := m.GetValue().(*TicketAction_Genesis); ok {
		return x.Genesis
	}
	return nil
}

func (m *TicketAction) GetTclose() *TicketClose {
	if x, ok := m.GetValue().(*TicketAction_Tclose); ok {
		return x.Tclose
	}
	return nil
}

func (m *TicketAction) GetMiner() *TicketMiner {
	if x, ok := m.GetValue().(*TicketAction_Miner); ok {
		return x.Miner
	}
	return nil
}

func (m *TicketAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TicketAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TicketAction_OneofMarshaler, _TicketAction_OneofUnmarshaler, _TicketAction_OneofSizer, []interface{}{
		(*TicketAction_Tbind)(nil),
		(*TicketAction_Topen)(nil),
		(*TicketAction_Genesis)(nil),
		(*TicketAction_Tclose)(nil),
		(*TicketAction_Miner)(nil),
	}
}

func _TicketAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TicketAction)
	// value
	switch x := m.Value.(type) {
	case *TicketAction_Tbind:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tbind); err != nil {
			return err
		}
	case *TicketAction_Topen:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Topen); err != nil {
			return err
		}
	case *TicketAction_Genesis:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Genesis); err != nil {
			return err
		}
	case *TicketAction_Tclose:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tclose); err != nil {
			return err
		}
	case *TicketAction_Miner:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Miner); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TicketAction.Value has unexpected type %T", x)
	}
	return nil
}

func _TicketAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TicketAction)
	switch tag {
	case 5: // value.tbind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketBind)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Tbind{msg}
		return true, err
	case 1: // value.topen
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketOpen)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Topen{msg}
		return true, err
	case 2: // value.genesis
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketGenesis)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Genesis{msg}
		return true, err
	case 3: // value.tclose
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketClose)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Tclose{msg}
		return true, err
	case 4: // value.miner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketMiner)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Miner{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TicketAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TicketAction)
	// value
	switch x := m.Value.(type) {
	case *TicketAction_Tbind:
		s := proto.Size(x.Tbind)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Topen:
		s := proto.Size(x.Topen)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Genesis:
		s := proto.Size(x.Genesis)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Tclose:
		s := proto.Size(x.Tclose)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Miner:
		s := proto.Size(x.Miner)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TicketMiner struct {
	Bits     uint32 `protobuf:"varint,1,opt,name=bits" json:"bits,omitempty"`
	Reward   int64  `protobuf:"varint,2,opt,name=reward" json:"reward,omitempty"`
	TicketId string `protobuf:"bytes,3,opt,name=ticketId" json:"ticketId,omitempty"`
	Modify   []byte `protobuf:"bytes,4,opt,name=modify,proto3" json:"modify,omitempty"`
	// 挖到区块时公开
	PrivHash []byte `protobuf:"bytes,5,opt,name=privHash,proto3" json:"privHash,omitempty"`
}

func (m *TicketMiner) Reset()                    { *m = TicketMiner{} }
func (m *TicketMiner) String() string            { return proto.CompactTextString(m) }
func (*TicketMiner) ProtoMessage()               {}
func (*TicketMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TicketMiner) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *TicketMiner) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *TicketMiner) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *TicketMiner) GetModify() []byte {
	if m != nil {
		return m.Modify
	}
	return nil
}

func (m *TicketMiner) GetPrivHash() []byte {
	if m != nil {
		return m.PrivHash
	}
	return nil
}

type TicketMinerOld struct {
	Bits     uint32 `protobuf:"varint,1,opt,name=bits" json:"bits,omitempty"`
	Reward   int64  `protobuf:"varint,2,opt,name=reward" json:"reward,omitempty"`
	TicketId string `protobuf:"bytes,3,opt,name=ticketId" json:"ticketId,omitempty"`
	Modify   []byte `protobuf:"bytes,4,opt,name=modify,proto3" json:"modify,omitempty"`
}

func (m *TicketMinerOld) Reset()                    { *m = TicketMinerOld{} }
func (m *TicketMinerOld) String() string            { return proto.CompactTextString(m) }
func (*TicketMinerOld) ProtoMessage()               {}
func (*TicketMinerOld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TicketMinerOld) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *TicketMinerOld) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *TicketMinerOld) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *TicketMinerOld) GetModify() []byte {
	if m != nil {
		return m.Modify
	}
	return nil
}

type MinerFlag struct {
	Flag    int32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
	Reserve int64 `protobuf:"varint,2,opt,name=reserve" json:"reserve,omitempty"`
}

func (m *MinerFlag) Reset()                    { *m = MinerFlag{} }
func (m *MinerFlag) String() string            { return proto.CompactTextString(m) }
func (*MinerFlag) ProtoMessage()               {}
func (*MinerFlag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MinerFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *MinerFlag) GetReserve() int64 {
	if m != nil {
		return m.Reserve
	}
	return 0
}

type TicketBind struct {
	MinerAddress  string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	ReturnAddress string `protobuf:"bytes,2,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *TicketBind) Reset()                    { *m = TicketBind{} }
func (m *TicketBind) String() string            { return proto.CompactTextString(m) }
func (*TicketBind) ProtoMessage()               {}
func (*TicketBind) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TicketBind) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketBind) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type TicketOpen struct {
	// 用户挖矿的ticket 地址
	MinerAddress string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	// 购买ticket的数目
	Count int32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	// 币实际存储的地址
	ReturnAddress string `protobuf:"bytes,3,opt,name=returnAddress" json:"returnAddress,omitempty"`
	// 随机种子
	RandSeed int64 `protobuf:"varint,4,opt,name=randSeed" json:"randSeed,omitempty"`
	// 购买ticket时公开
	PubHashes [][]byte `protobuf:"bytes,5,rep,name=pubHashes,proto3" json:"pubHashes,omitempty"`
}

func (m *TicketOpen) Reset()                    { *m = TicketOpen{} }
func (m *TicketOpen) String() string            { return proto.CompactTextString(m) }
func (*TicketOpen) ProtoMessage()               {}
func (*TicketOpen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TicketOpen) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketOpen) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TicketOpen) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func (m *TicketOpen) GetRandSeed() int64 {
	if m != nil {
		return m.RandSeed
	}
	return 0
}

func (m *TicketOpen) GetPubHashes() [][]byte {
	if m != nil {
		return m.PubHashes
	}
	return nil
}

type TicketGenesis struct {
	MinerAddress  string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	ReturnAddress string `protobuf:"bytes,2,opt,name=returnAddress" json:"returnAddress,omitempty"`
	Count         int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *TicketGenesis) Reset()                    { *m = TicketGenesis{} }
func (m *TicketGenesis) String() string            { return proto.CompactTextString(m) }
func (*TicketGenesis) ProtoMessage()               {}
func (*TicketGenesis) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TicketGenesis) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketGenesis) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func (m *TicketGenesis) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TicketClose struct {
	TicketId     []string `protobuf:"bytes,1,rep,name=ticketId" json:"ticketId,omitempty"`
	MinerAddress string   `protobuf:"bytes,2,opt,name=minerAddress" json:"minerAddress,omitempty"`
}

func (m *TicketClose) Reset()                    { *m = TicketClose{} }
func (m *TicketClose) String() string            { return proto.CompactTextString(m) }
func (*TicketClose) ProtoMessage()               {}
func (*TicketClose) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TicketClose) GetTicketId() []string {
	if m != nil {
		return m.TicketId
	}
	return nil
}

func (m *TicketClose) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

type TicketList struct {
	Addr   string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *TicketList) Reset()                    { *m = TicketList{} }
func (m *TicketList) String() string            { return proto.CompactTextString(m) }
func (*TicketList) ProtoMessage()               {}
func (*TicketList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TicketList) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *TicketList) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type TicketInfos struct {
	TicketIds []string `protobuf:"bytes,1,rep,name=ticketIds" json:"ticketIds,omitempty"`
}

func (m *TicketInfos) Reset()                    { *m = TicketInfos{} }
func (m *TicketInfos) String() string            { return proto.CompactTextString(m) }
func (*TicketInfos) ProtoMessage()               {}
func (*TicketInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TicketInfos) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

type ReplyTicketList struct {
	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *ReplyTicketList) Reset()                    { *m = ReplyTicketList{} }
func (m *ReplyTicketList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTicketList) ProtoMessage()               {}
func (*ReplyTicketList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplyTicketList) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type ReplyWalletTickets struct {
	Tickets  []*Ticket `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
	Privkeys [][]byte  `protobuf:"bytes,2,rep,name=privkeys,proto3" json:"privkeys,omitempty"`
}

func (m *ReplyWalletTickets) Reset()                    { *m = ReplyWalletTickets{} }
func (m *ReplyWalletTickets) String() string            { return proto.CompactTextString(m) }
func (*ReplyWalletTickets) ProtoMessage()               {}
func (*ReplyWalletTickets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReplyWalletTickets) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *ReplyWalletTickets) GetPrivkeys() [][]byte {
	if m != nil {
		return m.Privkeys
	}
	return nil
}

type ReceiptTicket struct {
	TicketId   string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	Status     int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	PrevStatus int32  `protobuf:"varint,3,opt,name=prevStatus" json:"prevStatus,omitempty"`
	Addr       string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *ReceiptTicket) Reset()                    { *m = ReceiptTicket{} }
func (m *ReceiptTicket) String() string            { return proto.CompactTextString(m) }
func (*ReceiptTicket) ProtoMessage()               {}
func (*ReceiptTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ReceiptTicket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *ReceiptTicket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiptTicket) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *ReceiptTicket) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ReceiptTicketBind struct {
	OldMinerAddress string `protobuf:"bytes,1,opt,name=oldMinerAddress" json:"oldMinerAddress,omitempty"`
	NewMinerAddress string `protobuf:"bytes,2,opt,name=newMinerAddress" json:"newMinerAddress,omitempty"`
	ReturnAddress   string `protobuf:"bytes,3,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *ReceiptTicketBind) Reset()                    { *m = ReceiptTicketBind{} }
func (m *ReceiptTicketBind) String() string            { return proto.CompactTextString(m) }
func (*ReceiptTicketBind) ProtoMessage()               {}
func (*ReceiptTicketBind) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ReceiptTicketBind) GetOldMinerAddress() string {
	if m != nil {
		return m.OldMinerAddress
	}
	return ""
}

func (m *ReceiptTicketBind) GetNewMinerAddress() string {
	if m != nil {
		return m.NewMinerAddress
	}
	return ""
}

func (m *ReceiptTicketBind) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type ReqBindMiner struct {
	BindAddr     string `protobuf:"bytes,1,opt,name=bindAddr" json:"bindAddr,omitempty"`
	OriginAddr   string `protobuf:"bytes,2,opt,name=originAddr" json:"originAddr,omitempty"`
	Amount       int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	CheckBalance bool   `protobuf:"varint,4,opt,name=checkBalance" json:"checkBalance,omitempty"`
}

func (m *ReqBindMiner) Reset()                    { *m = ReqBindMiner{} }
func (m *ReqBindMiner) String() string            { return proto.CompactTextString(m) }
func (*ReqBindMiner) ProtoMessage()               {}
func (*ReqBindMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ReqBindMiner) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *ReqBindMiner) GetOriginAddr() string {
	if m != nil {
		return m.OriginAddr
	}
	return ""
}

func (m *ReqBindMiner) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqBindMiner) GetCheckBalance() bool {
	if m != nil {
		return m.CheckBalance
	}
	return false
}

type ReplyBindMiner struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReplyBindMiner) Reset()                    { *m = ReplyBindMiner{} }
func (m *ReplyBindMiner) String() string            { return proto.CompactTextString(m) }
func (*ReplyBindMiner) ProtoMessage()               {}
func (*ReplyBindMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ReplyBindMiner) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func init() {
	proto.RegisterType((*Ticket)(nil), "types.Ticket")
	proto.RegisterType((*TicketAction)(nil), "types.TicketAction")
	proto.RegisterType((*TicketMiner)(nil), "types.TicketMiner")
	proto.RegisterType((*TicketMinerOld)(nil), "types.TicketMinerOld")
	proto.RegisterType((*MinerFlag)(nil), "types.MinerFlag")
	proto.RegisterType((*TicketBind)(nil), "types.TicketBind")
	proto.RegisterType((*TicketOpen)(nil), "types.TicketOpen")
	proto.RegisterType((*TicketGenesis)(nil), "types.TicketGenesis")
	proto.RegisterType((*TicketClose)(nil), "types.TicketClose")
	proto.RegisterType((*TicketList)(nil), "types.TicketList")
	proto.RegisterType((*TicketInfos)(nil), "types.TicketInfos")
	proto.RegisterType((*ReplyTicketList)(nil), "types.ReplyTicketList")
	proto.RegisterType((*ReplyWalletTickets)(nil), "types.ReplyWalletTickets")
	proto.RegisterType((*ReceiptTicket)(nil), "types.ReceiptTicket")
	proto.RegisterType((*ReceiptTicketBind)(nil), "types.ReceiptTicketBind")
	proto.RegisterType((*ReqBindMiner)(nil), "types.ReqBindMiner")
	proto.RegisterType((*ReplyBindMiner)(nil), "types.ReplyBindMiner")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ticket service

type TicketClient interface {
	// 创建绑定挖矿
	CreateBindMiner(ctx context.Context, in *ReqBindMiner, opts ...grpc.CallOption) (*ReplyBindMiner, error)
	// 查询钱包票数
	GetTicketCount(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*types1.Int64, error)
	// Miner
	// 设置自动挖矿
	SetAutoMining(ctx context.Context, in *MinerFlag, opts ...grpc.CallOption) (*types1.Reply, error)
}

type ticketClient struct {
	cc *grpc.ClientConn
}

func NewTicketClient(cc *grpc.ClientConn) TicketClient {
	return &ticketClient{cc}
}

func (c *ticketClient) CreateBindMiner(ctx context.Context, in *ReqBindMiner, opts ...grpc.CallOption) (*ReplyBindMiner, error) {
	out := new(ReplyBindMiner)
	err := grpc.Invoke(ctx, "/types.ticket/CreateBindMiner", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketClient) GetTicketCount(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*types1.Int64, error) {
	out := new(types1.Int64)
	err := grpc.Invoke(ctx, "/types.ticket/GetTicketCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketClient) SetAutoMining(ctx context.Context, in *MinerFlag, opts ...grpc.CallOption) (*types1.Reply, error) {
	out := new(types1.Reply)
	err := grpc.Invoke(ctx, "/types.ticket/SetAutoMining", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ticket service

type TicketServer interface {
	// 创建绑定挖矿
	CreateBindMiner(context.Context, *ReqBindMiner) (*ReplyBindMiner, error)
	// 查询钱包票数
	GetTicketCount(context.Context, *types1.ReqNil) (*types1.Int64, error)
	// Miner
	// 设置自动挖矿
	SetAutoMining(context.Context, *MinerFlag) (*types1.Reply, error)
}

func RegisterTicketServer(s *grpc.Server, srv TicketServer) {
	s.RegisterService(&_Ticket_serviceDesc, srv)
}

func _Ticket_CreateBindMiner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBindMiner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).CreateBindMiner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/CreateBindMiner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).CreateBindMiner(ctx, req.(*ReqBindMiner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticket_GetTicketCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).GetTicketCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/GetTicketCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).GetTicketCount(ctx, req.(*types1.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticket_SetAutoMining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerFlag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).SetAutoMining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/SetAutoMining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).SetAutoMining(ctx, req.(*MinerFlag))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ticket_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.ticket",
	HandlerType: (*TicketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBindMiner",
			Handler:    _Ticket_CreateBindMiner_Handler,
		},
		{
			MethodName: "GetTicketCount",
			Handler:    _Ticket_GetTicketCount_Handler,
		},
		{
			MethodName: "SetAutoMining",
			Handler:    _Ticket_SetAutoMining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}

func init() { proto.RegisterFile("ticket.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0xe3, 0x3a, 0x49, 0x5f, 0x9d, 0x96, 0x1d, 0x0a, 0xb2, 0x22, 0xb4, 0x8a, 0x46, 0x08,
	0xc2, 0x0f, 0x15, 0x28, 0x08, 0x01, 0x17, 0xd4, 0x56, 0x62, 0x53, 0x89, 0xb0, 0xd2, 0x74, 0xb5,
	0x88, 0xa3, 0x6b, 0xbf, 0x66, 0x47, 0x75, 0xc6, 0xc6, 0x9e, 0xb4, 0x9b, 0x2b, 0x07, 0xa4, 0x3d,
	0xf0, 0x6f, 0x70, 0xe3, 0x7f, 0x44, 0xf3, 0x3c, 0xfe, 0x95, 0xf4, 0x50, 0x09, 0xf6, 0xe6, 0xf7,
	0xe6, 0x7b, 0x7e, 0xdf, 0x7c, 0xf3, 0xbd, 0xb1, 0xc1, 0xd7, 0x32, 0xba, 0x45, 0x7d, 0x92, 0xe5,
	0xa9, 0x4e, 0x99, 0xa7, 0x37, 0x19, 0x16, 0x13, 0x3f, 0x4a, 0x57, 0xab, 0x54, 0x95, 0x49, 0xfe,
	0x47, 0x1f, 0x06, 0x2f, 0x08, 0xc5, 0x26, 0x30, 0x2a, 0xf1, 0x97, 0x71, 0xe0, 0x4c, 0x9d, 0xd9,
	0xbe, 0xa8, 0x63, 0xf6, 0x3e, 0x0c, 0x0a, 0x1d, 0xea, 0x75, 0x11, 0xf4, 0xa7, 0xce, 0xcc, 0x13,
	0x36, 0x62, 0x1f, 0xc0, 0xbe, 0x2c, 0x9e, 0xa1, 0xc2, 0x42, 0x16, 0x81, 0x3b, 0x75, 0x66, 0x23,
	0xd1, 0x24, 0xd8, 0x53, 0x80, 0x28, 0xc7, 0x50, 0xe3, 0x0b, 0xb9, 0xc2, 0x60, 0x6f, 0xea, 0xcc,
	0x5c, 0xd1, 0xca, 0x98, 0xea, 0x95, 0x54, 0x98, 0xd3, 0xb2, 0x47, 0xcb, 0x4d, 0xc2, 0x54, 0x53,
	0xf0, 0x32, 0x4c, 0xd6, 0x18, 0x8c, 0xca, 0xea, 0x26, 0xc3, 0x38, 0xf8, 0x14, 0x9d, 0xc5, 0x71,
	0x8e, 0x45, 0x11, 0x0c, 0x88, 0x73, 0x27, 0xc7, 0x3e, 0x84, 0x71, 0x8e, 0x7a, 0x9d, 0xab, 0x0a,
	0x34, 0x24, 0x50, 0x37, 0xc9, 0xdf, 0xf4, 0xc1, 0x2f, 0x45, 0x38, 0x8b, 0xb4, 0x4c, 0x15, 0xfb,
	0x04, 0x3c, 0x7d, 0x2d, 0x55, 0x4c, 0xa4, 0x0e, 0x4e, 0x9f, 0x9c, 0x90, 0x74, 0x27, 0x25, 0xe6,
	0x5c, 0xaa, 0x78, 0xde, 0x13, 0x25, 0x82, 0xa0, 0x69, 0x86, 0x8a, 0x24, 0xdb, 0x86, 0x3e, 0xcf,
	0x50, 0x11, 0xd4, 0x20, 0xd8, 0x97, 0x30, 0x5c, 0x5a, 0xa9, 0xfa, 0x04, 0x3e, 0xee, 0x80, 0xad,
	0x6a, 0xf3, 0x9e, 0xa8, 0x60, 0xec, 0x73, 0x18, 0xe8, 0x28, 0x49, 0x0b, 0x24, 0x6d, 0x0f, 0x4e,
	0x59, 0xa7, 0xe0, 0xc2, 0xac, 0xcc, 0x7b, 0xc2, 0x62, 0xd8, 0xa7, 0xe0, 0xd1, 0xe6, 0x49, 0xe9,
	0x6d, 0xf0, 0xc2, 0xac, 0x18, 0x2e, 0x04, 0x61, 0x87, 0xd0, 0xd7, 0x9b, 0x00, 0xe8, 0x30, 0xfb,
	0x7a, 0x73, 0x3e, 0x04, 0xef, 0xce, 0xa8, 0xca, 0xdf, 0x38, 0x70, 0xd0, 0xaa, 0x60, 0x0c, 0xf6,
	0xae, 0xa5, 0x2e, 0x68, 0x7b, 0x63, 0x41, 0xcf, 0xc6, 0x0d, 0x39, 0xde, 0x87, 0x79, 0x4c, 0xfb,
	0x70, 0x85, 0x8d, 0x3a, 0x0e, 0x72, 0x77, 0x1d, 0xb4, 0x4a, 0x63, 0x79, 0xb3, 0x21, 0x76, 0xbe,
	0xb0, 0x91, 0xa9, 0xc9, 0x72, 0x79, 0x37, 0x0f, 0x8b, 0x57, 0xa4, 0xb6, 0x2f, 0xea, 0x98, 0x67,
	0x70, 0xd8, 0xa2, 0xf2, 0x3c, 0x89, 0xdf, 0x36, 0x1b, 0xfe, 0x3d, 0xec, 0x53, 0xaf, 0x9f, 0x92,
	0x70, 0x69, 0x9a, 0xdd, 0x24, 0xe1, 0x92, 0x9a, 0x79, 0x82, 0x9e, 0x59, 0x00, 0xc3, 0x1c, 0x0b,
	0xcc, 0xef, 0xd0, 0x76, 0xab, 0x42, 0xfe, 0x12, 0xa0, 0xf1, 0xc7, 0x8e, 0x39, 0x9d, 0xc7, 0x98,
	0xb3, 0xff, 0x90, 0x39, 0xff, 0x76, 0xaa, 0x17, 0x1b, 0x37, 0x3d, 0xea, 0xc5, 0xc7, 0xe0, 0x45,
	0xe9, 0x5a, 0x69, 0x3b, 0xac, 0x65, 0xb0, 0xdb, 0xce, 0x7d, 0xa0, 0x9d, 0x51, 0x2d, 0x0f, 0x55,
	0x7c, 0x85, 0x18, 0xdb, 0x89, 0xad, 0x63, 0x33, 0xaf, 0xd9, 0xfa, 0xda, 0x1c, 0x0d, 0x16, 0x81,
	0x37, 0x75, 0x67, 0xbe, 0x68, 0x12, 0x3c, 0x85, 0x71, 0xc7, 0xc8, 0xff, 0x9f, 0x06, 0xcd, 0x86,
	0xdc, 0xd6, 0x86, 0xf8, 0xa2, 0x72, 0x2a, 0x0d, 0xc2, 0xd6, 0xfd, 0xe5, 0x76, 0xce, 0x7b, 0x9b,
	0x4a, 0x7f, 0x97, 0x0a, 0xff, 0xae, 0xd2, 0xf9, 0x67, 0x59, 0x68, 0x73, 0xf8, 0x61, 0x1c, 0xe7,
	0x96, 0x34, 0x3d, 0xb7, 0x6e, 0x41, 0xb7, 0x7d, 0x0b, 0xf2, 0xcf, 0x2a, 0x22, 0x97, 0xea, 0x26,
	0xa5, 0x4b, 0xb1, 0x6a, 0x5c, 0x58, 0x26, 0x4d, 0x82, 0xff, 0x00, 0x47, 0x02, 0xb3, 0x64, 0xd3,
	0xea, 0xf5, 0x31, 0x0c, 0xcb, 0xf5, 0x12, 0x7e, 0x70, 0x3a, 0xee, 0x8c, 0xae, 0xa8, 0x56, 0xf9,
	0x6f, 0xc0, 0xa8, 0xf6, 0xd7, 0x30, 0x49, 0x50, 0x97, 0xab, 0xc5, 0xa3, 0xcb, 0xab, 0x59, 0xbb,
	0xc5, 0x8d, 0x51, 0xc0, 0xad, 0x66, 0xcd, 0xc4, 0xfc, 0x1e, 0xc6, 0x02, 0x23, 0x94, 0x99, 0xfe,
	0x0f, 0x9f, 0x83, 0xa7, 0x00, 0x59, 0x8e, 0x77, 0x57, 0x6d, 0x91, 0x5a, 0x99, 0x5a, 0xd4, 0xbd,
	0x46, 0x54, 0xfe, 0x97, 0x03, 0x4f, 0x3a, 0x9d, 0x69, 0x7e, 0x66, 0x70, 0x94, 0x26, 0xf1, 0x62,
	0xd7, 0x3e, 0xdb, 0x69, 0x83, 0x54, 0x78, 0xbf, 0xd8, 0x3d, 0xdd, 0xed, 0xf4, 0xe3, 0x06, 0x80,
	0xff, 0xe9, 0x80, 0x2f, 0xf0, 0x77, 0xc3, 0xa2, 0xbc, 0x01, 0x27, 0x30, 0x32, 0x37, 0xfd, 0x59,
	0xe3, 0x86, 0x3a, 0x36, 0x1b, 0x4e, 0x73, 0xb9, 0x94, 0x54, 0x6d, 0xfb, 0xb6, 0x32, 0x46, 0xa8,
	0x70, 0x55, 0x3b, 0xd7, 0x15, 0x36, 0x32, 0x7e, 0x8c, 0x5e, 0x61, 0x74, 0x7b, 0x1e, 0x26, 0xa1,
	0x8a, 0xca, 0x6f, 0xe3, 0x48, 0x74, 0x72, 0xfc, 0x23, 0x38, 0xa4, 0xc3, 0x6e, 0x98, 0x1c, 0x83,
	0xa7, 0x5f, 0xcf, 0xf1, 0xb5, 0xa5, 0x51, 0x06, 0xa7, 0xff, 0x38, 0x30, 0x28, 0x4f, 0x86, 0xfd,
	0x08, 0x47, 0x17, 0xf4, 0x79, 0x6d, 0x6a, 0xde, 0xb5, 0x5e, 0x68, 0x6f, 0x69, 0xf2, 0x5e, 0x9d,
	0x6c, 0xbf, 0x9f, 0xf7, 0xd8, 0x17, 0x70, 0xf8, 0xac, 0x32, 0xd6, 0x05, 0x31, 0x1d, 0x37, 0xf5,
	0xbf, 0xc8, 0x64, 0xe2, 0xdb, 0xf0, 0x52, 0xe9, 0x6f, 0xbf, 0xe1, 0x3d, 0xf6, 0x15, 0x8c, 0xaf,
	0x50, 0x9f, 0xad, 0x75, 0xba, 0x90, 0x4a, 0xaa, 0x25, 0x7b, 0xc7, 0x02, 0xea, 0x6b, 0xb4, 0x2e,
	0xa1, 0x66, 0xbc, 0x77, 0x3d, 0xa0, 0x3f, 0x8f, 0xaf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd7,
	0x1e, 0x82, 0xa2, 0x9e, 0x08, 0x00, 0x00,
}
