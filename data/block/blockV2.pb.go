// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockV2.proto

package block

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ThotaGopichandThota_gn_core2_data "github.com/ThotaGopichandThota/gn-core2/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// HeaderV2 extends the Header structure with extra fields for version 2
type HeaderV2 struct {
	Header                   *Header       `protobuf:"bytes,1,opt,name=Header,proto3" json:"header,omitempty"`
	ScheduledRootHash        []byte        `protobuf:"bytes,2,opt,name=ScheduledRootHash,proto3" json:"scheduledRootHash,omitempty"`
	ScheduledAccumulatedFees *math_big.Int `protobuf:"bytes,3,opt,name=ScheduledAccumulatedFees,proto3,casttypewith=math/big.Int;github.com/ThotaGopichandThota/gn-core2/data.BigIntCaster" json:"scheduledAccumulatedFees,omitempty"`
	ScheduledDeveloperFees   *math_big.Int `protobuf:"bytes,4,opt,name=ScheduledDeveloperFees,proto3,casttypewith=math/big.Int;github.com/ThotaGopichandThota/gn-core2/data.BigIntCaster" json:"scheduledDeveloperFees,omitempty"`
	ScheduledGasProvided     uint64        `protobuf:"varint,5,opt,name=ScheduledGasProvided,proto3" json:"scheduledGasProvided"`
	ScheduledGasPenalized    uint64        `protobuf:"varint,6,opt,name=ScheduledGasPenalized,proto3" json:"scheduledGasPenalized"`
	ScheduledGasRefunded     uint64        `protobuf:"varint,7,opt,name=ScheduledGasRefunded,proto3" json:"scheduledGasRefunded"`
}

func (m *HeaderV2) Reset()      { *m = HeaderV2{} }
func (*HeaderV2) ProtoMessage() {}
func (*HeaderV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{0}
}
func (m *HeaderV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeaderV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderV2.Merge(m, src)
}
func (m *HeaderV2) XXX_Size() int {
	return m.Size()
}
func (m *HeaderV2) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderV2.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderV2 proto.InternalMessageInfo

func (m *HeaderV2) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderV2) GetScheduledRootHash() []byte {
	if m != nil {
		return m.ScheduledRootHash
	}
	return nil
}

func (m *HeaderV2) GetScheduledAccumulatedFees() *math_big.Int {
	if m != nil {
		return m.ScheduledAccumulatedFees
	}
	return nil
}

func (m *HeaderV2) GetScheduledDeveloperFees() *math_big.Int {
	if m != nil {
		return m.ScheduledDeveloperFees
	}
	return nil
}

func (m *HeaderV2) GetScheduledGasProvided() uint64 {
	if m != nil {
		return m.ScheduledGasProvided
	}
	return 0
}

func (m *HeaderV2) GetScheduledGasPenalized() uint64 {
	if m != nil {
		return m.ScheduledGasPenalized
	}
	return 0
}

func (m *HeaderV2) GetScheduledGasRefunded() uint64 {
	if m != nil {
		return m.ScheduledGasRefunded
	}
	return 0
}

type MiniBlockReserved struct {
	ExecutionType    ProcessingType `protobuf:"varint,1,opt,name=ExecutionType,proto3,enum=proto.ProcessingType" json:"executionType"`
	TransactionsType []byte         `protobuf:"bytes,2,opt,name=TransactionsType,proto3" json:"transactionsType"`
}

func (m *MiniBlockReserved) Reset()      { *m = MiniBlockReserved{} }
func (*MiniBlockReserved) ProtoMessage() {}
func (*MiniBlockReserved) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{1}
}
func (m *MiniBlockReserved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniBlockReserved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniBlockReserved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniBlockReserved.Merge(m, src)
}
func (m *MiniBlockReserved) XXX_Size() int {
	return m.Size()
}
func (m *MiniBlockReserved) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniBlockReserved.DiscardUnknown(m)
}

var xxx_messageInfo_MiniBlockReserved proto.InternalMessageInfo

func (m *MiniBlockReserved) GetExecutionType() ProcessingType {
	if m != nil {
		return m.ExecutionType
	}
	return Normal
}

func (m *MiniBlockReserved) GetTransactionsType() []byte {
	if m != nil {
		return m.TransactionsType
	}
	return nil
}

type MiniBlockHeaderReserved struct {
	ExecutionType           ProcessingType `protobuf:"varint,1,opt,name=ExecutionType,proto3,enum=proto.ProcessingType" json:"executionType"`
	State                   MiniBlockState `protobuf:"varint,2,opt,name=State,proto3,enum=proto.MiniBlockState" json:"state"`
	IndexOfFirstTxProcessed int32          `protobuf:"varint,3,opt,name=IndexOfFirstTxProcessed,proto3" json:"indexOfFirstTxProcessed"`
	IndexOfLastTxProcessed  int32          `protobuf:"varint,4,opt,name=IndexOfLastTxProcessed,proto3" json:"indexOfLastTxProcessed"`
}

func (m *MiniBlockHeaderReserved) Reset()      { *m = MiniBlockHeaderReserved{} }
func (*MiniBlockHeaderReserved) ProtoMessage() {}
func (*MiniBlockHeaderReserved) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{2}
}
func (m *MiniBlockHeaderReserved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniBlockHeaderReserved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniBlockHeaderReserved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniBlockHeaderReserved.Merge(m, src)
}
func (m *MiniBlockHeaderReserved) XXX_Size() int {
	return m.Size()
}
func (m *MiniBlockHeaderReserved) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniBlockHeaderReserved.DiscardUnknown(m)
}

var xxx_messageInfo_MiniBlockHeaderReserved proto.InternalMessageInfo

func (m *MiniBlockHeaderReserved) GetExecutionType() ProcessingType {
	if m != nil {
		return m.ExecutionType
	}
	return Normal
}

func (m *MiniBlockHeaderReserved) GetState() MiniBlockState {
	if m != nil {
		return m.State
	}
	return Final
}

func (m *MiniBlockHeaderReserved) GetIndexOfFirstTxProcessed() int32 {
	if m != nil {
		return m.IndexOfFirstTxProcessed
	}
	return 0
}

func (m *MiniBlockHeaderReserved) GetIndexOfLastTxProcessed() int32 {
	if m != nil {
		return m.IndexOfLastTxProcessed
	}
	return 0
}

func init() {
	proto.RegisterType((*HeaderV2)(nil), "proto.HeaderV2")
	proto.RegisterType((*MiniBlockReserved)(nil), "proto.MiniBlockReserved")
	proto.RegisterType((*MiniBlockHeaderReserved)(nil), "proto.MiniBlockHeaderReserved")
}

func init() { proto.RegisterFile("blockV2.proto", fileDescriptor_17a3844aa051366e) }

var fileDescriptor_17a3844aa051366e = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x4f, 0xd4, 0x40,
	0x18, 0xde, 0x41, 0x8a, 0x3a, 0xb8, 0x06, 0x1a, 0x3e, 0x2a, 0x24, 0xd3, 0xcd, 0x9e, 0x38, 0xc8,
	0x6e, 0xb2, 0x26, 0x5e, 0x38, 0xa8, 0x55, 0xf9, 0x48, 0x40, 0xc8, 0xb0, 0x72, 0xf0, 0x36, 0xdb,
	0xbe, 0xb4, 0x13, 0x77, 0x3b, 0x4d, 0x67, 0x4a, 0xc0, 0x93, 0x3f, 0xc1, 0x3f, 0xa1, 0x31, 0xc6,
	0x1f, 0xe2, 0xc1, 0x03, 0x47, 0x12, 0x93, 0x2a, 0xe5, 0x62, 0x7a, 0xe2, 0x27, 0x98, 0x9d, 0xae,
	0xcb, 0x57, 0xf7, 0xc8, 0xa9, 0xef, 0xfb, 0x3e, 0x1f, 0xf3, 0x64, 0xf2, 0x4e, 0x71, 0xb5, 0xd3,
	0x15, 0xee, 0xfb, 0xbd, 0x56, 0x23, 0x8a, 0x85, 0x12, 0xa6, 0xa1, 0x3f, 0x0b, 0xcb, 0x3e, 0x57,
	0x41, 0xd2, 0x69, 0xb8, 0xa2, 0xd7, 0xf4, 0x85, 0x2f, 0x9a, 0x7a, 0xdc, 0x49, 0xf6, 0x75, 0xa7,
	0x1b, 0x5d, 0x15, 0xaa, 0x85, 0x49, 0x6d, 0x52, 0x34, 0xf5, 0x5f, 0x06, 0xbe, 0xb7, 0x0e, 0xcc,
	0x83, 0x78, 0xaf, 0x65, 0xae, 0xe0, 0x89, 0xa2, 0xb6, 0x50, 0x0d, 0x2d, 0x4d, 0xb6, 0xaa, 0x05,
	0xa9, 0x51, 0x0c, 0x9d, 0x99, 0x3c, 0xb5, 0xa7, 0x02, 0x5d, 0x3f, 0x16, 0x3d, 0xae, 0xa0, 0x17,
	0xa9, 0x23, 0x3a, 0x90, 0x98, 0x5b, 0x78, 0x7a, 0xd7, 0x0d, 0xc0, 0x4b, 0xba, 0xe0, 0x51, 0x21,
	0xd4, 0x3a, 0x93, 0x81, 0x35, 0x56, 0x43, 0x4b, 0x0f, 0x1c, 0x3b, 0x4f, 0xed, 0x45, 0x79, 0x1d,
	0xbc, 0xe4, 0x71, 0x53, 0x69, 0x7e, 0x47, 0xd8, 0x1a, 0x4e, 0x5f, 0xb8, 0x6e, 0xd2, 0x4b, 0xba,
	0x4c, 0x81, 0xb7, 0x0a, 0x20, 0xad, 0x3b, 0xda, 0x36, 0xca, 0x53, 0xbb, 0x2e, 0x47, 0x70, 0x2e,
	0xdc, 0xbf, 0xfd, 0xb6, 0x57, 0x7b, 0x4c, 0x05, 0xcd, 0x0e, 0xf7, 0x1b, 0x1b, 0xa1, 0x5a, 0xb9,
	0x74, 0x5d, 0xed, 0x40, 0x28, 0xb6, 0x26, 0x22, 0xee, 0x06, 0x2c, 0xf4, 0x74, 0xd7, 0xf4, 0xc3,
	0x65, 0x57, 0xc4, 0xd0, 0x6a, 0x7a, 0x4c, 0xb1, 0x86, 0xc3, 0xfd, 0x8d, 0x50, 0xbd, 0x64, 0x52,
	0x41, 0x4c, 0x47, 0x26, 0x32, 0xbf, 0x20, 0x3c, 0x37, 0x04, 0x5f, 0xc1, 0x01, 0x74, 0x45, 0x04,
	0xb1, 0x0e, 0x3b, 0xae, 0xc3, 0x86, 0x79, 0x6a, 0xd7, 0x64, 0x29, 0xe3, 0x56, 0xa2, 0x8e, 0x48,
	0x63, 0x6e, 0xe2, 0x99, 0x21, 0xb2, 0xc6, 0xe4, 0x4e, 0x2c, 0x0e, 0xb8, 0x07, 0x9e, 0x65, 0xd4,
	0xd0, 0xd2, 0xb8, 0x63, 0xe5, 0xa9, 0x3d, 0x23, 0x4b, 0x70, 0x5a, 0xaa, 0x32, 0xb7, 0xf1, 0xec,
	0x95, 0x39, 0x84, 0xac, 0xcb, 0x3f, 0x80, 0x67, 0x4d, 0x68, 0xbb, 0x47, 0x79, 0x6a, 0xcf, 0xca,
	0x32, 0x02, 0x2d, 0xd7, 0x5d, 0x8f, 0x47, 0x61, 0x3f, 0x09, 0xfb, 0xf1, 0xee, 0x96, 0xc7, 0xfb,
	0x8f, 0xd3, 0x52, 0x55, 0xfd, 0x33, 0xc2, 0xd3, 0x5b, 0x3c, 0xe4, 0x4e, 0x7f, 0xe3, 0x29, 0x48,
	0x88, 0x0f, 0xc0, 0x33, 0xdf, 0xe0, 0xea, 0xeb, 0x43, 0x70, 0x13, 0xc5, 0x45, 0xd8, 0x3e, 0x8a,
	0x40, 0x6f, 0xfb, 0xc3, 0xd6, 0xec, 0x60, 0xdb, 0x77, 0x62, 0xe1, 0x82, 0x94, 0x3c, 0xf4, 0xfb,
	0xa0, 0x33, 0x9d, 0xa7, 0x76, 0x15, 0x2e, 0xf3, 0xe9, 0x55, 0xb9, 0xf9, 0x1c, 0x4f, 0xb5, 0x63,
	0x16, 0x4a, 0xe6, 0xf6, 0x47, 0x52, 0x5b, 0x16, 0x8b, 0xaf, 0x5f, 0x8c, 0xba, 0x86, 0xd1, 0x1b,
	0xec, 0xfa, 0xcf, 0x31, 0x3c, 0x3f, 0xcc, 0x59, 0xbc, 0xa7, 0x5b, 0x4b, 0xfb, 0x14, 0x1b, 0xbb,
	0x8a, 0xa9, 0x22, 0xe2, 0x85, 0xcf, 0xf0, 0x78, 0x0d, 0x3a, 0xf7, 0xf3, 0xd4, 0x36, 0x64, 0xbf,
	0xa4, 0x05, 0xdd, 0x7c, 0x8b, 0xe7, 0x37, 0x42, 0x0f, 0x0e, 0xb7, 0xf7, 0x57, 0x79, 0x2c, 0x55,
	0xfb, 0x70, 0x70, 0x32, 0x78, 0xfa, 0x39, 0x1a, 0xce, 0x62, 0x9e, 0xda, 0xf3, 0xbc, 0x9c, 0x42,
	0x47, 0x69, 0x4d, 0x8a, 0xe7, 0x06, 0xd0, 0x26, 0xbb, 0xea, 0x3a, 0xae, 0x5d, 0x17, 0xf2, 0xd4,
	0x9e, 0xe3, 0xa5, 0x0c, 0x3a, 0x42, 0xe9, 0x3c, 0x3b, 0x3e, 0x25, 0x95, 0x93, 0x53, 0x52, 0x39,
	0x3f, 0x25, 0xe8, 0x63, 0x46, 0xd0, 0xd7, 0x8c, 0xa0, 0x1f, 0x19, 0x41, 0xc7, 0x19, 0x41, 0x27,
	0x19, 0x41, 0x7f, 0x32, 0x82, 0xfe, 0x66, 0xa4, 0x72, 0x9e, 0x11, 0xf4, 0xe9, 0x8c, 0x54, 0x8e,
	0xcf, 0x48, 0xe5, 0xe4, 0x8c, 0x54, 0xde, 0x19, 0xfa, 0xdf, 0xd8, 0x99, 0xd0, 0x77, 0xf2, 0xe4,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xfc, 0xe1, 0xf9, 0x70, 0x05, 0x00, 0x00,
}

func (this *HeaderV2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderV2)
	if !ok {
		that2, ok := that.(HeaderV2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !bytes.Equal(this.ScheduledRootHash, that1.ScheduledRootHash) {
		return false
	}
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		if !__caster.Equal(this.ScheduledAccumulatedFees, that1.ScheduledAccumulatedFees) {
			return false
		}
	}
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		if !__caster.Equal(this.ScheduledDeveloperFees, that1.ScheduledDeveloperFees) {
			return false
		}
	}
	if this.ScheduledGasProvided != that1.ScheduledGasProvided {
		return false
	}
	if this.ScheduledGasPenalized != that1.ScheduledGasPenalized {
		return false
	}
	if this.ScheduledGasRefunded != that1.ScheduledGasRefunded {
		return false
	}
	return true
}
func (this *MiniBlockReserved) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniBlockReserved)
	if !ok {
		that2, ok := that.(MiniBlockReserved)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ExecutionType != that1.ExecutionType {
		return false
	}
	if !bytes.Equal(this.TransactionsType, that1.TransactionsType) {
		return false
	}
	return true
}
func (this *MiniBlockHeaderReserved) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniBlockHeaderReserved)
	if !ok {
		that2, ok := that.(MiniBlockHeaderReserved)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ExecutionType != that1.ExecutionType {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if this.IndexOfFirstTxProcessed != that1.IndexOfFirstTxProcessed {
		return false
	}
	if this.IndexOfLastTxProcessed != that1.IndexOfLastTxProcessed {
		return false
	}
	return true
}
func (this *HeaderV2) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&block.HeaderV2{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
	s = append(s, "ScheduledRootHash: "+fmt.Sprintf("%#v", this.ScheduledRootHash)+",\n")
	s = append(s, "ScheduledAccumulatedFees: "+fmt.Sprintf("%#v", this.ScheduledAccumulatedFees)+",\n")
	s = append(s, "ScheduledDeveloperFees: "+fmt.Sprintf("%#v", this.ScheduledDeveloperFees)+",\n")
	s = append(s, "ScheduledGasProvided: "+fmt.Sprintf("%#v", this.ScheduledGasProvided)+",\n")
	s = append(s, "ScheduledGasPenalized: "+fmt.Sprintf("%#v", this.ScheduledGasPenalized)+",\n")
	s = append(s, "ScheduledGasRefunded: "+fmt.Sprintf("%#v", this.ScheduledGasRefunded)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MiniBlockReserved) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&block.MiniBlockReserved{")
	s = append(s, "ExecutionType: "+fmt.Sprintf("%#v", this.ExecutionType)+",\n")
	s = append(s, "TransactionsType: "+fmt.Sprintf("%#v", this.TransactionsType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MiniBlockHeaderReserved) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&block.MiniBlockHeaderReserved{")
	s = append(s, "ExecutionType: "+fmt.Sprintf("%#v", this.ExecutionType)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "IndexOfFirstTxProcessed: "+fmt.Sprintf("%#v", this.IndexOfFirstTxProcessed)+",\n")
	s = append(s, "IndexOfLastTxProcessed: "+fmt.Sprintf("%#v", this.IndexOfLastTxProcessed)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBlockV2(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HeaderV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScheduledGasRefunded != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasRefunded))
		i--
		dAtA[i] = 0x38
	}
	if m.ScheduledGasPenalized != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasPenalized))
		i--
		dAtA[i] = 0x30
	}
	if m.ScheduledGasProvided != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasProvided))
		i--
		dAtA[i] = 0x28
	}
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		size := __caster.Size(m.ScheduledDeveloperFees)
		i -= size
		if _, err := __caster.MarshalTo(m.ScheduledDeveloperFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlockV2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		size := __caster.Size(m.ScheduledAccumulatedFees)
		i -= size
		if _, err := __caster.MarshalTo(m.ScheduledAccumulatedFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlockV2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ScheduledRootHash) > 0 {
		i -= len(m.ScheduledRootHash)
		copy(dAtA[i:], m.ScheduledRootHash)
		i = encodeVarintBlockV2(dAtA, i, uint64(len(m.ScheduledRootHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlockV2(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MiniBlockReserved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniBlockReserved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniBlockReserved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransactionsType) > 0 {
		i -= len(m.TransactionsType)
		copy(dAtA[i:], m.TransactionsType)
		i = encodeVarintBlockV2(dAtA, i, uint64(len(m.TransactionsType)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExecutionType != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ExecutionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MiniBlockHeaderReserved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniBlockHeaderReserved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniBlockHeaderReserved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IndexOfLastTxProcessed != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.IndexOfLastTxProcessed))
		i--
		dAtA[i] = 0x20
	}
	if m.IndexOfFirstTxProcessed != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.IndexOfFirstTxProcessed))
		i--
		dAtA[i] = 0x18
	}
	if m.State != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.ExecutionType != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ExecutionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockV2(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockV2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBlockV2(uint64(l))
	}
	l = len(m.ScheduledRootHash)
	if l > 0 {
		n += 1 + l + sovBlockV2(uint64(l))
	}
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		l = __caster.Size(m.ScheduledAccumulatedFees)
		n += 1 + l + sovBlockV2(uint64(l))
	}
	{
		__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
		l = __caster.Size(m.ScheduledDeveloperFees)
		n += 1 + l + sovBlockV2(uint64(l))
	}
	if m.ScheduledGasProvided != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasProvided))
	}
	if m.ScheduledGasPenalized != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasPenalized))
	}
	if m.ScheduledGasRefunded != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasRefunded))
	}
	return n
}

func (m *MiniBlockReserved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecutionType != 0 {
		n += 1 + sovBlockV2(uint64(m.ExecutionType))
	}
	l = len(m.TransactionsType)
	if l > 0 {
		n += 1 + l + sovBlockV2(uint64(l))
	}
	return n
}

func (m *MiniBlockHeaderReserved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecutionType != 0 {
		n += 1 + sovBlockV2(uint64(m.ExecutionType))
	}
	if m.State != 0 {
		n += 1 + sovBlockV2(uint64(m.State))
	}
	if m.IndexOfFirstTxProcessed != 0 {
		n += 1 + sovBlockV2(uint64(m.IndexOfFirstTxProcessed))
	}
	if m.IndexOfLastTxProcessed != 0 {
		n += 1 + sovBlockV2(uint64(m.IndexOfLastTxProcessed))
	}
	return n
}

func sovBlockV2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockV2(x uint64) (n int) {
	return sovBlockV2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HeaderV2) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeaderV2{`,
		`Header:` + strings.Replace(fmt.Sprintf("%v", this.Header), "Header", "Header", 1) + `,`,
		`ScheduledRootHash:` + fmt.Sprintf("%v", this.ScheduledRootHash) + `,`,
		`ScheduledAccumulatedFees:` + fmt.Sprintf("%v", this.ScheduledAccumulatedFees) + `,`,
		`ScheduledDeveloperFees:` + fmt.Sprintf("%v", this.ScheduledDeveloperFees) + `,`,
		`ScheduledGasProvided:` + fmt.Sprintf("%v", this.ScheduledGasProvided) + `,`,
		`ScheduledGasPenalized:` + fmt.Sprintf("%v", this.ScheduledGasPenalized) + `,`,
		`ScheduledGasRefunded:` + fmt.Sprintf("%v", this.ScheduledGasRefunded) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MiniBlockReserved) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniBlockReserved{`,
		`ExecutionType:` + fmt.Sprintf("%v", this.ExecutionType) + `,`,
		`TransactionsType:` + fmt.Sprintf("%v", this.TransactionsType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MiniBlockHeaderReserved) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniBlockHeaderReserved{`,
		`ExecutionType:` + fmt.Sprintf("%v", this.ExecutionType) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`IndexOfFirstTxProcessed:` + fmt.Sprintf("%v", this.IndexOfFirstTxProcessed) + `,`,
		`IndexOfLastTxProcessed:` + fmt.Sprintf("%v", this.IndexOfLastTxProcessed) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlockV2(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HeaderV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: HeaderV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledRootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScheduledRootHash = append(m.ScheduledRootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ScheduledRootHash == nil {
				m.ScheduledRootHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledAccumulatedFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.ScheduledAccumulatedFees = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledDeveloperFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ThotaGopichandThota_gn_core2_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.ScheduledDeveloperFees = tmp
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasProvided", wireType)
			}
			m.ScheduledGasProvided = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasProvided |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasPenalized", wireType)
			}
			m.ScheduledGasPenalized = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasPenalized |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasRefunded", wireType)
			}
			m.ScheduledGasRefunded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasRefunded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func (m *MiniBlockReserved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: MiniBlockReserved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniBlockReserved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionType", wireType)
			}
			m.ExecutionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionType |= ProcessingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsType", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionsType = append(m.TransactionsType[:0], dAtA[iNdEx:postIndex]...)
			if m.TransactionsType == nil {
				m.TransactionsType = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func (m *MiniBlockHeaderReserved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: MiniBlockHeaderReserved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniBlockHeaderReserved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionType", wireType)
			}
			m.ExecutionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionType |= ProcessingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= MiniBlockState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOfFirstTxProcessed", wireType)
			}
			m.IndexOfFirstTxProcessed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOfFirstTxProcessed |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOfLastTxProcessed", wireType)
			}
			m.IndexOfLastTxProcessed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOfLastTxProcessed |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func skipBlockV2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockV2
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
					return 0, ErrIntOverflowBlockV2
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
					return 0, ErrIntOverflowBlockV2
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
				return 0, ErrInvalidLengthBlockV2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockV2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockV2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockV2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockV2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockV2 = fmt.Errorf("proto: unexpected end of group")
)