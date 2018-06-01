// Code generated by protoc-gen-go. DO NOT EDIT.
// source: evmcontract.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	evmcontract.proto

It has these top-level messages:
	EVMContractObject
	EVMContractData
	EVMContractState
	EVMContractAction
	ReceiptEVMContract
	EVMContractDataCmd
	EVMContractStateCmd
	ReceiptEVMContractCmd
	CheckEVMAddrReq
	CheckEVMAddrResp
	EstimateEVMGasReq
	EstimateEVMGasResp
	EvmDebugReq
	EvmDebugResp
*/
package types

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

// 合约对象信息
type EVMContractObject struct {
	Addr  string            `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Data  *EVMContractData  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	State *EVMContractState `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *EVMContractObject) Reset()                    { *m = EVMContractObject{} }
func (m *EVMContractObject) String() string            { return proto.CompactTextString(m) }
func (*EVMContractObject) ProtoMessage()               {}
func (*EVMContractObject) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{0} }

func (m *EVMContractObject) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *EVMContractObject) GetData() *EVMContractData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EVMContractObject) GetState() *EVMContractState {
	if m != nil {
		return m.State
	}
	return nil
}

// 存放合约固定数据
type EVMContractData struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator" json:"creator,omitempty"`
	CreateTime int64  `protobuf:"varint,2,opt,name=createTime" json:"createTime,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Alias      string `protobuf:"bytes,4,opt,name=alias" json:"alias,omitempty"`
	Addr       string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty"`
	Code       []byte `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	CodeHash   []byte `protobuf:"bytes,7,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
}

func (m *EVMContractData) Reset()                    { *m = EVMContractData{} }
func (m *EVMContractData) String() string            { return proto.CompactTextString(m) }
func (*EVMContractData) ProtoMessage()               {}
func (*EVMContractData) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{1} }

func (m *EVMContractData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EVMContractData) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *EVMContractData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EVMContractData) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *EVMContractData) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *EVMContractData) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *EVMContractData) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

// 存放合约变化数据
type EVMContractState struct {
	Nonce       uint64            `protobuf:"varint,1,opt,name=nonce" json:"nonce,omitempty"`
	Suicided    bool              `protobuf:"varint,2,opt,name=suicided" json:"suicided,omitempty"`
	StorageHash []byte            `protobuf:"bytes,3,opt,name=storageHash,proto3" json:"storageHash,omitempty"`
	Storage     map[string][]byte `protobuf:"bytes,4,rep,name=storage" json:"storage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *EVMContractState) Reset()                    { *m = EVMContractState{} }
func (m *EVMContractState) String() string            { return proto.CompactTextString(m) }
func (*EVMContractState) ProtoMessage()               {}
func (*EVMContractState) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{2} }

func (m *EVMContractState) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EVMContractState) GetSuicided() bool {
	if m != nil {
		return m.Suicided
	}
	return false
}

func (m *EVMContractState) GetStorageHash() []byte {
	if m != nil {
		return m.StorageHash
	}
	return nil
}

func (m *EVMContractState) GetStorage() map[string][]byte {
	if m != nil {
		return m.Storage
	}
	return nil
}

// 创建/调用合约的请求结构
type EVMContractAction struct {
	// 转账金额
	Amount uint64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	// 消耗限制，默认为Transaction.Fee
	GasLimit uint64 `protobuf:"varint,2,opt,name=gasLimit" json:"gasLimit,omitempty"`
	// gas价格，默认为1
	GasPrice uint32 `protobuf:"varint,3,opt,name=gasPrice" json:"gasPrice,omitempty"`
	// 合约数据
	Code []byte `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	// 合约别名，方便识别
	Alias string `protobuf:"bytes,5,opt,name=alias" json:"alias,omitempty"`
	// 交易备注
	Note string `protobuf:"bytes,6,opt,name=note" json:"note,omitempty"`
}

func (m *EVMContractAction) Reset()                    { *m = EVMContractAction{} }
func (m *EVMContractAction) String() string            { return proto.CompactTextString(m) }
func (*EVMContractAction) ProtoMessage()               {}
func (*EVMContractAction) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{3} }

func (m *EVMContractAction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EVMContractAction) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *EVMContractAction) GetGasPrice() uint32 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *EVMContractAction) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *EVMContractAction) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *EVMContractAction) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

// 合约创建/调用日志
type ReceiptEVMContract struct {
	Caller       string `protobuf:"bytes,1,opt,name=caller" json:"caller,omitempty"`
	ContractName string `protobuf:"bytes,2,opt,name=contractName" json:"contractName,omitempty"`
	ContractAddr string `protobuf:"bytes,3,opt,name=contractAddr" json:"contractAddr,omitempty"`
	UsedGas      uint64 `protobuf:"varint,4,opt,name=usedGas" json:"usedGas,omitempty"`
	// 创建合约返回的代码
	Ret []byte `protobuf:"bytes,5,opt,name=ret,proto3" json:"ret,omitempty"`
}

func (m *ReceiptEVMContract) Reset()                    { *m = ReceiptEVMContract{} }
func (m *ReceiptEVMContract) String() string            { return proto.CompactTextString(m) }
func (*ReceiptEVMContract) ProtoMessage()               {}
func (*ReceiptEVMContract) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{4} }

func (m *ReceiptEVMContract) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *ReceiptEVMContract) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *ReceiptEVMContract) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *ReceiptEVMContract) GetUsedGas() uint64 {
	if m != nil {
		return m.UsedGas
	}
	return 0
}

func (m *ReceiptEVMContract) GetRet() []byte {
	if m != nil {
		return m.Ret
	}
	return nil
}

// 存放合约固定数据
type EVMContractDataCmd struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator" json:"creator,omitempty"`
	CreateTime string `protobuf:"bytes,2,opt,name=createTime" json:"createTime,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Alias      string `protobuf:"bytes,4,opt,name=alias" json:"alias,omitempty"`
	Addr       string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty"`
	Code       string `protobuf:"bytes,6,opt,name=code" json:"code,omitempty"`
	CodeHash   string `protobuf:"bytes,7,opt,name=codeHash" json:"codeHash,omitempty"`
}

func (m *EVMContractDataCmd) Reset()                    { *m = EVMContractDataCmd{} }
func (m *EVMContractDataCmd) String() string            { return proto.CompactTextString(m) }
func (*EVMContractDataCmd) ProtoMessage()               {}
func (*EVMContractDataCmd) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{5} }

func (m *EVMContractDataCmd) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EVMContractDataCmd) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *EVMContractDataCmd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EVMContractDataCmd) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *EVMContractDataCmd) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *EVMContractDataCmd) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *EVMContractDataCmd) GetCodeHash() string {
	if m != nil {
		return m.CodeHash
	}
	return ""
}

// 存放合约变化数据
type EVMContractStateCmd struct {
	Nonce       uint64            `protobuf:"varint,1,opt,name=nonce" json:"nonce,omitempty"`
	Suicided    bool              `protobuf:"varint,2,opt,name=suicided" json:"suicided,omitempty"`
	StorageHash string            `protobuf:"bytes,3,opt,name=storageHash" json:"storageHash,omitempty"`
	Storage     map[string]string `protobuf:"bytes,4,rep,name=storage" json:"storage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EVMContractStateCmd) Reset()                    { *m = EVMContractStateCmd{} }
func (m *EVMContractStateCmd) String() string            { return proto.CompactTextString(m) }
func (*EVMContractStateCmd) ProtoMessage()               {}
func (*EVMContractStateCmd) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{6} }

func (m *EVMContractStateCmd) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EVMContractStateCmd) GetSuicided() bool {
	if m != nil {
		return m.Suicided
	}
	return false
}

func (m *EVMContractStateCmd) GetStorageHash() string {
	if m != nil {
		return m.StorageHash
	}
	return ""
}

func (m *EVMContractStateCmd) GetStorage() map[string]string {
	if m != nil {
		return m.Storage
	}
	return nil
}

// 合约创建/调用日志
type ReceiptEVMContractCmd struct {
	Caller string `protobuf:"bytes,1,opt,name=caller" json:"caller,omitempty"`
	// 合约创建时才会返回此内容
	ContractName string `protobuf:"bytes,2,opt,name=contractName" json:"contractName,omitempty"`
	ContractAddr string `protobuf:"bytes,3,opt,name=contractAddr" json:"contractAddr,omitempty"`
	UsedGas      uint64 `protobuf:"varint,4,opt,name=usedGas" json:"usedGas,omitempty"`
	// 创建合约返回的代码
	Ret string `protobuf:"bytes,5,opt,name=ret" json:"ret,omitempty"`
}

func (m *ReceiptEVMContractCmd) Reset()                    { *m = ReceiptEVMContractCmd{} }
func (m *ReceiptEVMContractCmd) String() string            { return proto.CompactTextString(m) }
func (*ReceiptEVMContractCmd) ProtoMessage()               {}
func (*ReceiptEVMContractCmd) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{7} }

func (m *ReceiptEVMContractCmd) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *ReceiptEVMContractCmd) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *ReceiptEVMContractCmd) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *ReceiptEVMContractCmd) GetUsedGas() uint64 {
	if m != nil {
		return m.UsedGas
	}
	return 0
}

func (m *ReceiptEVMContractCmd) GetRet() string {
	if m != nil {
		return m.Ret
	}
	return ""
}

type CheckEVMAddrReq struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *CheckEVMAddrReq) Reset()                    { *m = CheckEVMAddrReq{} }
func (m *CheckEVMAddrReq) String() string            { return proto.CompactTextString(m) }
func (*CheckEVMAddrReq) ProtoMessage()               {}
func (*CheckEVMAddrReq) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{8} }

func (m *CheckEVMAddrReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type CheckEVMAddrResp struct {
	Contract     bool   `protobuf:"varint,1,opt,name=contract" json:"contract,omitempty"`
	CreateTime   string `protobuf:"bytes,2,opt,name=createTime" json:"createTime,omitempty"`
	ContractAddr string `protobuf:"bytes,3,opt,name=contractAddr" json:"contractAddr,omitempty"`
	ContractName string `protobuf:"bytes,4,opt,name=contractName" json:"contractName,omitempty"`
	AliasName    string `protobuf:"bytes,5,opt,name=aliasName" json:"aliasName,omitempty"`
}

func (m *CheckEVMAddrResp) Reset()                    { *m = CheckEVMAddrResp{} }
func (m *CheckEVMAddrResp) String() string            { return proto.CompactTextString(m) }
func (*CheckEVMAddrResp) ProtoMessage()               {}
func (*CheckEVMAddrResp) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{9} }

func (m *CheckEVMAddrResp) GetContract() bool {
	if m != nil {
		return m.Contract
	}
	return false
}

func (m *CheckEVMAddrResp) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *CheckEVMAddrResp) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *CheckEVMAddrResp) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *CheckEVMAddrResp) GetAliasName() string {
	if m != nil {
		return m.AliasName
	}
	return ""
}

type EstimateEVMGasReq struct {
	To   string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Code []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *EstimateEVMGasReq) Reset()                    { *m = EstimateEVMGasReq{} }
func (m *EstimateEVMGasReq) String() string            { return proto.CompactTextString(m) }
func (*EstimateEVMGasReq) ProtoMessage()               {}
func (*EstimateEVMGasReq) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{10} }

func (m *EstimateEVMGasReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EstimateEVMGasReq) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type EstimateEVMGasResp struct {
	Gas uint64 `protobuf:"varint,1,opt,name=gas" json:"gas,omitempty"`
}

func (m *EstimateEVMGasResp) Reset()                    { *m = EstimateEVMGasResp{} }
func (m *EstimateEVMGasResp) String() string            { return proto.CompactTextString(m) }
func (*EstimateEVMGasResp) ProtoMessage()               {}
func (*EstimateEVMGasResp) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{11} }

func (m *EstimateEVMGasResp) GetGas() uint64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

type EvmDebugReq struct {
	// 0 query, 1 set, -1 clear
	Optype int32 `protobuf:"varint,1,opt,name=optype" json:"optype,omitempty"`
}

func (m *EvmDebugReq) Reset()                    { *m = EvmDebugReq{} }
func (m *EvmDebugReq) String() string            { return proto.CompactTextString(m) }
func (*EvmDebugReq) ProtoMessage()               {}
func (*EvmDebugReq) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{12} }

func (m *EvmDebugReq) GetOptype() int32 {
	if m != nil {
		return m.Optype
	}
	return 0
}

type EvmDebugResp struct {
	DebugStatus string `protobuf:"bytes,1,opt,name=debugStatus" json:"debugStatus,omitempty"`
}

func (m *EvmDebugResp) Reset()                    { *m = EvmDebugResp{} }
func (m *EvmDebugResp) String() string            { return proto.CompactTextString(m) }
func (*EvmDebugResp) ProtoMessage()               {}
func (*EvmDebugResp) Descriptor() ([]byte, []int) { return fileDescriptor99, []int{13} }

func (m *EvmDebugResp) GetDebugStatus() string {
	if m != nil {
		return m.DebugStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*EVMContractObject)(nil), "types.EVMContractObject")
	proto.RegisterType((*EVMContractData)(nil), "types.EVMContractData")
	proto.RegisterType((*EVMContractState)(nil), "types.EVMContractState")
	proto.RegisterType((*EVMContractAction)(nil), "types.EVMContractAction")
	proto.RegisterType((*ReceiptEVMContract)(nil), "types.ReceiptEVMContract")
	proto.RegisterType((*EVMContractDataCmd)(nil), "types.EVMContractDataCmd")
	proto.RegisterType((*EVMContractStateCmd)(nil), "types.EVMContractStateCmd")
	proto.RegisterType((*ReceiptEVMContractCmd)(nil), "types.ReceiptEVMContractCmd")
	proto.RegisterType((*CheckEVMAddrReq)(nil), "types.CheckEVMAddrReq")
	proto.RegisterType((*CheckEVMAddrResp)(nil), "types.CheckEVMAddrResp")
	proto.RegisterType((*EstimateEVMGasReq)(nil), "types.EstimateEVMGasReq")
	proto.RegisterType((*EstimateEVMGasResp)(nil), "types.EstimateEVMGasResp")
	proto.RegisterType((*EvmDebugReq)(nil), "types.EvmDebugReq")
	proto.RegisterType((*EvmDebugResp)(nil), "types.EvmDebugResp")
}

func init() { proto.RegisterFile("evmcontract.proto", fileDescriptor99) }

var fileDescriptor99 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xe5, 0x36, 0xed, 0x96, 0xd3, 0x7e, 0xdf, 0x36, 0x03, 0x23, 0x9a, 0x10, 0xaa, 0x22,
	0x06, 0x15, 0x12, 0x15, 0x1a, 0x17, 0xa0, 0x5d, 0x20, 0x4d, 0x5d, 0x35, 0x2e, 0x18, 0x20, 0x0f,
	0xed, 0xde, 0x4b, 0xac, 0x2e, 0xac, 0x89, 0x43, 0xec, 0x4e, 0xda, 0x2d, 0x4f, 0x02, 0xe2, 0x15,
	0xb8, 0x41, 0xe2, 0x69, 0xb8, 0xe2, 0x31, 0xd0, 0xb1, 0x9d, 0x2c, 0xcd, 0x36, 0x01, 0xd2, 0x90,
	0xb8, 0xda, 0xf9, 0x1f, 0xff, 0x6d, 0x9f, 0x73, 0xfc, 0x5b, 0x0a, 0x6b, 0xe2, 0x34, 0x8d, 0x64,
	0xa6, 0x0b, 0x1e, 0xe9, 0x51, 0x5e, 0x48, 0x2d, 0x69, 0x47, 0x9f, 0xe5, 0x42, 0x85, 0x1f, 0x08,
	0xac, 0x4d, 0x0e, 0xf7, 0xc7, 0x6e, 0xf1, 0xf5, 0xd1, 0x3b, 0x11, 0x69, 0x4a, 0xc1, 0xe3, 0x71,
	0x5c, 0x04, 0x64, 0x40, 0x86, 0x3e, 0x33, 0x31, 0x7d, 0x08, 0x5e, 0xcc, 0x35, 0x0f, 0x5a, 0x03,
	0x32, 0xec, 0x6d, 0xad, 0x8f, 0xcc, 0xfe, 0x51, 0x6d, 0xef, 0x2e, 0xd7, 0x9c, 0x19, 0x0f, 0x7d,
	0x04, 0x1d, 0xa5, 0xb9, 0x16, 0x41, 0xdb, 0x98, 0x6f, 0x5f, 0x34, 0x1f, 0xe0, 0x32, 0xb3, 0xae,
	0xf0, 0x2b, 0x81, 0x95, 0xc6, 0x41, 0x34, 0x80, 0xa5, 0xa8, 0x10, 0x5c, 0xcb, 0xb2, 0x8a, 0x52,
	0xd2, 0xbb, 0x00, 0x26, 0x14, 0x6f, 0x93, 0x54, 0x98, 0x72, 0xda, 0xac, 0x96, 0xc1, 0xe2, 0x33,
	0x9e, 0xda, 0xbb, 0x7d, 0x66, 0x62, 0x7a, 0x13, 0x3a, 0x7c, 0x96, 0x70, 0x15, 0x78, 0x26, 0x69,
	0x45, 0xd5, 0x66, 0xa7, 0xd6, 0x26, 0x05, 0x2f, 0x92, 0xb1, 0x08, 0xba, 0x03, 0x32, 0xec, 0x33,
	0x13, 0xd3, 0x0d, 0x58, 0xc6, 0xbf, 0x2f, 0xb8, 0x3a, 0x0e, 0x96, 0x4c, 0xbe, 0xd2, 0xe1, 0x77,
	0x02, 0xab, 0xcd, 0xbe, 0xf0, 0xba, 0x4c, 0x66, 0x91, 0x30, 0xa5, 0x7b, 0xcc, 0x0a, 0x3c, 0x46,
	0xcd, 0x93, 0x28, 0x89, 0x45, 0x6c, 0xca, 0x5e, 0x66, 0x95, 0xa6, 0x03, 0xe8, 0x29, 0x2d, 0x0b,
	0x3e, 0xb5, 0xb7, 0xb4, 0xcd, 0x2d, 0xf5, 0x14, 0x7d, 0x0e, 0x4b, 0x4e, 0x06, 0xde, 0xa0, 0x3d,
	0xec, 0x6d, 0xdd, 0xbb, 0x62, 0xaa, 0xa3, 0x03, 0x6b, 0x9b, 0x64, 0xba, 0x38, 0x63, 0xe5, 0xa6,
	0x8d, 0x6d, 0xe8, 0xd7, 0x17, 0xe8, 0x2a, 0xb4, 0x4f, 0xc4, 0x99, 0x1b, 0x2e, 0x86, 0x58, 0xf5,
	0x29, 0x9f, 0xcd, 0xed, 0x4c, 0xfb, 0xcc, 0x8a, 0xed, 0xd6, 0x33, 0x12, 0x7e, 0x5a, 0xa4, 0x64,
	0x27, 0xd2, 0x89, 0xcc, 0xe8, 0x3a, 0x74, 0x79, 0x2a, 0xe7, 0x99, 0x76, 0x6d, 0x3a, 0x85, 0x7d,
	0x4e, 0xb9, 0x7a, 0x99, 0xa4, 0x89, 0x36, 0x47, 0x79, 0xac, 0xd2, 0x6e, 0xed, 0x4d, 0x91, 0x44,
	0xf6, 0x81, 0xfe, 0x63, 0x95, 0xae, 0x46, 0xef, 0xd5, 0x46, 0x5f, 0x3d, 0x5c, 0xa7, 0xf1, 0x70,
	0x99, 0xd4, 0xf6, 0x91, 0xf0, 0x89, 0xa5, 0x16, 0xe1, 0x47, 0x02, 0x94, 0x89, 0x48, 0x24, 0xb9,
	0xae, 0x95, 0x8a, 0x45, 0x46, 0x7c, 0x36, 0x13, 0x25, 0x46, 0x4e, 0xd1, 0x10, 0xfa, 0xe5, 0x7f,
	0xc4, 0x2b, 0xee, 0x38, 0xf2, 0xd9, 0x42, 0xae, 0xee, 0xd9, 0x41, 0x4e, 0xda, 0x8b, 0x1e, 0xcc,
	0x21, 0xa7, 0x73, 0x25, 0xe2, 0x3d, 0xc7, 0x96, 0xc7, 0x4a, 0x89, 0x03, 0x2e, 0x84, 0x36, 0x85,
	0xf7, 0x19, 0x86, 0xe1, 0x37, 0x02, 0xb4, 0xc1, 0xf9, 0x38, 0x8d, 0xff, 0x08, 0x75, 0xff, 0x2f,
	0xa1, 0xee, 0x5f, 0x81, 0xba, 0x5f, 0x43, 0xfd, 0x07, 0x81, 0x1b, 0x4d, 0xd8, 0xb0, 0xfe, 0x6b,
	0xa1, 0xdd, 0x5f, 0xa4, 0x7d, 0xa7, 0x49, 0xfb, 0x83, 0x2b, 0x68, 0x1f, 0xa7, 0xf1, 0xf5, 0x00,
	0xef, 0xd7, 0x81, 0xff, 0x4c, 0xe0, 0xd6, 0x45, 0x98, 0xb0, 0xd9, 0x7f, 0x82, 0x27, 0xdf, 0xf2,
	0xb4, 0x09, 0x2b, 0xe3, 0x63, 0x11, 0x9d, 0x4c, 0x0e, 0xf7, 0x71, 0x2f, 0x13, 0xef, 0x2f, 0xfb,
	0x72, 0x87, 0x5f, 0x08, 0xac, 0x2e, 0xfa, 0x54, 0x6e, 0x1f, 0xda, 0xde, 0x6b, 0xcc, 0xcb, 0xac,
	0xd2, 0xbf, 0xc4, 0xee, 0x77, 0xfa, 0x68, 0xce, 0xc3, 0xbb, 0x64, 0x1e, 0x77, 0xc0, 0x37, 0x74,
	0x1a, 0x83, 0xed, 0xeb, 0x3c, 0x11, 0x3e, 0x85, 0xb5, 0x89, 0xd2, 0x49, 0xca, 0xb5, 0x98, 0x1c,
	0xee, 0xef, 0x71, 0x85, 0xfd, 0xfd, 0x0f, 0x2d, 0x2d, 0x5d, 0x77, 0x2d, 0x2d, 0x2b, 0x86, 0x5b,
	0xe7, 0xdf, 0x8c, 0xf0, 0x3e, 0xd0, 0xe6, 0x46, 0x95, 0xe3, 0xf8, 0xa6, 0x5c, 0x39, 0x46, 0x31,
	0x0c, 0x37, 0xa1, 0x37, 0x39, 0x4d, 0x77, 0xc5, 0xd1, 0x7c, 0x8a, 0x47, 0xaf, 0x43, 0x57, 0xe6,
	0x08, 0x99, 0xf1, 0x74, 0x98, 0x53, 0xe1, 0x63, 0xe8, 0x9f, 0xdb, 0x54, 0x8e, 0xf0, 0xc6, 0x28,
	0x10, 0xbf, 0xb9, 0x72, 0xb5, 0xd4, 0x53, 0x47, 0x5d, 0xf3, 0x13, 0xfb, 0xe4, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0xc2, 0x36, 0x4b, 0x77, 0x07, 0x00, 0x00,
}