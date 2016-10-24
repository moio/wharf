// Code generated by protoc-gen-go.
// source: pwr/pwr.proto
// DO NOT EDIT!

/*
Package pwr is a generated protocol buffer package.

It is generated from these files:
	pwr/pwr.proto

It has these top-level messages:
	PatchHeader
	SyncHeader
	SyncOp
	SignatureHeader
	BlockHash
	CompressionSettings
	ManifestHeader
	ManifestBlockHash
	WoundsHeader
	Wound
	BsdiffControl
*/
package pwr

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

type CompressionAlgorithm int32

const (
	CompressionAlgorithm_NONE   CompressionAlgorithm = 0
	CompressionAlgorithm_BROTLI CompressionAlgorithm = 1
	CompressionAlgorithm_GZIP   CompressionAlgorithm = 2
)

var CompressionAlgorithm_name = map[int32]string{
	0: "NONE",
	1: "BROTLI",
	2: "GZIP",
}
var CompressionAlgorithm_value = map[string]int32{
	"NONE":   0,
	"BROTLI": 1,
	"GZIP":   2,
}

func (x CompressionAlgorithm) String() string {
	return proto.EnumName(CompressionAlgorithm_name, int32(x))
}
func (CompressionAlgorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HashAlgorithm int32

const (
	HashAlgorithm_SHAKE128_32 HashAlgorithm = 0
)

var HashAlgorithm_name = map[int32]string{
	0: "SHAKE128_32",
}
var HashAlgorithm_value = map[string]int32{
	"SHAKE128_32": 0,
}

func (x HashAlgorithm) String() string {
	return proto.EnumName(HashAlgorithm_name, int32(x))
}
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type WoundKind int32

const (
	WoundKind_FILE    WoundKind = 0
	WoundKind_SYMLINK WoundKind = 1
	WoundKind_DIR     WoundKind = 2
	// sent when a file portion has been verified as valid
	WoundKind_CLOSED_FILE WoundKind = 3
)

var WoundKind_name = map[int32]string{
	0: "FILE",
	1: "SYMLINK",
	2: "DIR",
	3: "CLOSED_FILE",
}
var WoundKind_value = map[string]int32{
	"FILE":        0,
	"SYMLINK":     1,
	"DIR":         2,
	"CLOSED_FILE": 3,
}

func (x WoundKind) String() string {
	return proto.EnumName(WoundKind_name, int32(x))
}
func (WoundKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SyncOp_Type int32

const (
	SyncOp_BLOCK_RANGE SyncOp_Type = 0
	SyncOp_DATA        SyncOp_Type = 1
	SyncOp_BSDIFF      SyncOp_Type = 2
	// REMOTE_DATA used to be 2 - shouldn't be in the wild, but better not re-use it?
	SyncOp_HEY_YOU_DID_IT SyncOp_Type = 2049
)

var SyncOp_Type_name = map[int32]string{
	0:    "BLOCK_RANGE",
	1:    "DATA",
	2:    "BSDIFF",
	2049: "HEY_YOU_DID_IT",
}
var SyncOp_Type_value = map[string]int32{
	"BLOCK_RANGE":    0,
	"DATA":           1,
	"BSDIFF":         2,
	"HEY_YOU_DID_IT": 2049,
}

func (x SyncOp_Type) String() string {
	return proto.EnumName(SyncOp_Type_name, int32(x))
}
func (SyncOp_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type PatchHeader struct {
	Compression *CompressionSettings `protobuf:"bytes,1,opt,name=compression" json:"compression,omitempty"`
}

func (m *PatchHeader) Reset()                    { *m = PatchHeader{} }
func (m *PatchHeader) String() string            { return proto.CompactTextString(m) }
func (*PatchHeader) ProtoMessage()               {}
func (*PatchHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PatchHeader) GetCompression() *CompressionSettings {
	if m != nil {
		return m.Compression
	}
	return nil
}

type SyncHeader struct {
	FileIndex int64 `protobuf:"varint,16,opt,name=fileIndex" json:"fileIndex,omitempty"`
}

func (m *SyncHeader) Reset()                    { *m = SyncHeader{} }
func (m *SyncHeader) String() string            { return proto.CompactTextString(m) }
func (*SyncHeader) ProtoMessage()               {}
func (*SyncHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SyncOp struct {
	Type       SyncOp_Type `protobuf:"varint,1,opt,name=type,enum=io.itch.wharf.pwr.SyncOp_Type" json:"type,omitempty"`
	FileIndex  int64       `protobuf:"varint,2,opt,name=fileIndex" json:"fileIndex,omitempty"`
	BlockIndex int64       `protobuf:"varint,3,opt,name=blockIndex" json:"blockIndex,omitempty"`
	BlockSpan  int64       `protobuf:"varint,4,opt,name=blockSpan" json:"blockSpan,omitempty"`
	Data       []byte      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SyncOp) Reset()                    { *m = SyncOp{} }
func (m *SyncOp) String() string            { return proto.CompactTextString(m) }
func (*SyncOp) ProtoMessage()               {}
func (*SyncOp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SignatureHeader struct {
	Compression *CompressionSettings `protobuf:"bytes,1,opt,name=compression" json:"compression,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SignatureHeader) GetCompression() *CompressionSettings {
	if m != nil {
		return m.Compression
	}
	return nil
}

type BlockHash struct {
	WeakHash   uint32 `protobuf:"varint,1,opt,name=weakHash" json:"weakHash,omitempty"`
	StrongHash []byte `protobuf:"bytes,2,opt,name=strongHash,proto3" json:"strongHash,omitempty"`
}

func (m *BlockHash) Reset()                    { *m = BlockHash{} }
func (m *BlockHash) String() string            { return proto.CompactTextString(m) }
func (*BlockHash) ProtoMessage()               {}
func (*BlockHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CompressionSettings struct {
	Algorithm CompressionAlgorithm `protobuf:"varint,1,opt,name=algorithm,enum=io.itch.wharf.pwr.CompressionAlgorithm" json:"algorithm,omitempty"`
	Quality   int32                `protobuf:"varint,2,opt,name=quality" json:"quality,omitempty"`
}

func (m *CompressionSettings) Reset()                    { *m = CompressionSettings{} }
func (m *CompressionSettings) String() string            { return proto.CompactTextString(m) }
func (*CompressionSettings) ProtoMessage()               {}
func (*CompressionSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ManifestHeader struct {
	Compression *CompressionSettings `protobuf:"bytes,1,opt,name=compression" json:"compression,omitempty"`
	Algorithm   HashAlgorithm        `protobuf:"varint,2,opt,name=algorithm,enum=io.itch.wharf.pwr.HashAlgorithm" json:"algorithm,omitempty"`
}

func (m *ManifestHeader) Reset()                    { *m = ManifestHeader{} }
func (m *ManifestHeader) String() string            { return proto.CompactTextString(m) }
func (*ManifestHeader) ProtoMessage()               {}
func (*ManifestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ManifestHeader) GetCompression() *CompressionSettings {
	if m != nil {
		return m.Compression
	}
	return nil
}

type ManifestBlockHash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *ManifestBlockHash) Reset()                    { *m = ManifestBlockHash{} }
func (m *ManifestBlockHash) String() string            { return proto.CompactTextString(m) }
func (*ManifestBlockHash) ProtoMessage()               {}
func (*ManifestBlockHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Wounds files format: header, container, then any
// number of Wounds
type WoundsHeader struct {
}

func (m *WoundsHeader) Reset()                    { *m = WoundsHeader{} }
func (m *WoundsHeader) String() string            { return proto.CompactTextString(m) }
func (*WoundsHeader) ProtoMessage()               {}
func (*WoundsHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Describe a corrupted portion of a file, in [start,end)
type Wound struct {
	Index int64     `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Start int64     `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	End   int64     `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
	Kind  WoundKind `protobuf:"varint,4,opt,name=kind,enum=io.itch.wharf.pwr.WoundKind" json:"kind,omitempty"`
}

func (m *Wound) Reset()                    { *m = Wound{} }
func (m *Wound) String() string            { return proto.CompactTextString(m) }
func (*Wound) ProtoMessage()               {}
func (*Wound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// bsdiff messages
type BsdiffControl struct {
	Add  []byte `protobuf:"bytes,1,opt,name=add,proto3" json:"add,omitempty"`
	Copy []byte `protobuf:"bytes,2,opt,name=copy,proto3" json:"copy,omitempty"`
	Seek int64  `protobuf:"varint,3,opt,name=seek" json:"seek,omitempty"`
}

func (m *BsdiffControl) Reset()                    { *m = BsdiffControl{} }
func (m *BsdiffControl) String() string            { return proto.CompactTextString(m) }
func (*BsdiffControl) ProtoMessage()               {}
func (*BsdiffControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*PatchHeader)(nil), "io.itch.wharf.pwr.PatchHeader")
	proto.RegisterType((*SyncHeader)(nil), "io.itch.wharf.pwr.SyncHeader")
	proto.RegisterType((*SyncOp)(nil), "io.itch.wharf.pwr.SyncOp")
	proto.RegisterType((*SignatureHeader)(nil), "io.itch.wharf.pwr.SignatureHeader")
	proto.RegisterType((*BlockHash)(nil), "io.itch.wharf.pwr.BlockHash")
	proto.RegisterType((*CompressionSettings)(nil), "io.itch.wharf.pwr.CompressionSettings")
	proto.RegisterType((*ManifestHeader)(nil), "io.itch.wharf.pwr.ManifestHeader")
	proto.RegisterType((*ManifestBlockHash)(nil), "io.itch.wharf.pwr.ManifestBlockHash")
	proto.RegisterType((*WoundsHeader)(nil), "io.itch.wharf.pwr.WoundsHeader")
	proto.RegisterType((*Wound)(nil), "io.itch.wharf.pwr.Wound")
	proto.RegisterType((*BsdiffControl)(nil), "io.itch.wharf.pwr.BsdiffControl")
	proto.RegisterEnum("io.itch.wharf.pwr.CompressionAlgorithm", CompressionAlgorithm_name, CompressionAlgorithm_value)
	proto.RegisterEnum("io.itch.wharf.pwr.HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
	proto.RegisterEnum("io.itch.wharf.pwr.WoundKind", WoundKind_name, WoundKind_value)
	proto.RegisterEnum("io.itch.wharf.pwr.SyncOp_Type", SyncOp_Type_name, SyncOp_Type_value)
}

func init() { proto.RegisterFile("pwr/pwr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0x8d, 0xc1, 0x24, 0x61, 0x08, 0xc4, 0xd9, 0xe4, 0x80, 0x7e, 0x8a, 0x22, 0xe4, 0xc3, 0x2f,
	0x11, 0x07, 0xb7, 0x25, 0x52, 0xd5, 0x43, 0x55, 0x89, 0x7f, 0x09, 0x16, 0x04, 0x22, 0x9b, 0x2a,
	0x4a, 0x7a, 0x40, 0x1b, 0xbc, 0xc0, 0x2a, 0x64, 0xd7, 0x5d, 0x6f, 0xea, 0x72, 0xec, 0xd7, 0xe8,
	0xa7, 0xec, 0x47, 0xa8, 0x76, 0x31, 0x81, 0xb4, 0x56, 0x4f, 0xb9, 0xcd, 0xcc, 0xce, 0xcc, 0x7b,
	0xfb, 0xfc, 0xd6, 0x50, 0x0c, 0x63, 0xf1, 0x26, 0x8c, 0x85, 0x13, 0x0a, 0x2e, 0x39, 0x3a, 0xa0,
	0xdc, 0xa1, 0x72, 0x3c, 0x73, 0xe2, 0x19, 0x16, 0x13, 0x27, 0x8c, 0x85, 0x7d, 0x03, 0x85, 0x6b,
	0x2c, 0xc7, 0xb3, 0x0e, 0xc1, 0x01, 0x11, 0xa8, 0x03, 0x85, 0x31, 0x7f, 0x0c, 0x05, 0x89, 0x22,
	0xca, 0x59, 0xd9, 0xa8, 0x18, 0x67, 0x85, 0xda, 0xff, 0xce, 0x5f, 0x73, 0x4e, 0x73, 0xdd, 0xe5,
	0x13, 0x29, 0x29, 0x9b, 0x46, 0xde, 0xe6, 0xa8, 0x5d, 0x05, 0xf0, 0x17, 0x6c, 0x9c, 0xec, 0x3d,
	0x86, 0xfc, 0x84, 0xce, 0x89, 0xcb, 0x02, 0xf2, 0xbd, 0x6c, 0x55, 0x8c, 0xb3, 0xac, 0xb7, 0x2e,
	0xd8, 0xbf, 0x0c, 0xd8, 0x56, 0xcd, 0x83, 0x10, 0xd5, 0xc0, 0x94, 0x8b, 0x90, 0x68, 0xe4, 0x52,
	0xed, 0x24, 0x05, 0x79, 0xd9, 0xe8, 0x0c, 0x17, 0x21, 0xf1, 0x74, 0xef, 0xcb, 0xe5, 0x99, 0x3f,
	0x96, 0xa3, 0x13, 0x80, 0xfb, 0x39, 0x1f, 0x3f, 0x2c, 0x8f, 0xb3, 0xfa, 0x78, 0xa3, 0xa2, 0xa6,
	0x75, 0xe6, 0x87, 0x98, 0x95, 0xcd, 0xe5, 0xf4, 0x73, 0x01, 0x21, 0x30, 0x03, 0x2c, 0x71, 0x39,
	0x57, 0x31, 0xce, 0xf6, 0x3c, 0x1d, 0xdb, 0x0d, 0x30, 0x15, 0x3a, 0xda, 0x87, 0x42, 0xa3, 0x37,
	0x68, 0x76, 0x47, 0x5e, 0xbd, 0x7f, 0xd9, 0xb6, 0xb6, 0xd0, 0x2e, 0x98, 0xad, 0xfa, 0xb0, 0x6e,
	0x19, 0x08, 0x60, 0xbb, 0xe1, 0xb7, 0xdc, 0x8b, 0x0b, 0x2b, 0x83, 0x0e, 0xa1, 0xd4, 0x69, 0xdf,
	0x8e, 0x6e, 0x07, 0x9f, 0x47, 0x2d, 0xb7, 0x35, 0x72, 0x87, 0xd6, 0x0f, 0xcb, 0xfe, 0x02, 0xfb,
	0x3e, 0x9d, 0x32, 0x2c, 0x9f, 0x04, 0x79, 0x75, 0xed, 0x2f, 0x21, 0xdf, 0x50, 0x37, 0xe8, 0xe0,
	0x68, 0x86, 0xfe, 0x83, 0xdd, 0x98, 0x60, 0x1d, 0xeb, 0x9d, 0x45, 0xef, 0x39, 0x57, 0xda, 0x44,
	0x52, 0x70, 0x36, 0xd5, 0xa7, 0x19, 0x7d, 0xc7, 0x8d, 0x8a, 0xfd, 0x0d, 0x0e, 0x53, 0xc0, 0x50,
	0x1b, 0xf2, 0x78, 0x3e, 0xe5, 0x82, 0xca, 0xd9, 0x63, 0xf2, 0xa5, 0x4e, 0xff, 0xcd, 0xb3, 0xbe,
	0x6a, 0xf7, 0xd6, 0x93, 0xa8, 0x0c, 0x3b, 0x5f, 0x9f, 0xf0, 0x9c, 0xca, 0x85, 0x86, 0xce, 0x79,
	0xab, 0xd4, 0xfe, 0x69, 0x40, 0xe9, 0x0a, 0x33, 0x3a, 0x21, 0x91, 0x7c, 0x6d, 0x75, 0xd0, 0xa7,
	0x4d, 0xf6, 0x19, 0xcd, 0xbe, 0x92, 0xb2, 0x47, 0x09, 0x90, 0x46, 0xdb, 0x3e, 0x85, 0x83, 0x15,
	0xb7, 0xb5, 0xca, 0x08, 0xcc, 0xd9, 0x4a, 0xe1, 0x3d, 0x4f, 0xc7, 0x76, 0x09, 0xf6, 0x6e, 0xf8,
	0x13, 0x0b, 0xa2, 0xe5, 0x15, 0xec, 0x18, 0x72, 0x3a, 0x47, 0x47, 0x90, 0xa3, 0xda, 0x8d, 0x86,
	0xb6, 0xdb, 0x32, 0x51, 0xd5, 0x48, 0x62, 0x21, 0x13, 0x0b, 0x2f, 0x13, 0x64, 0x41, 0x96, 0xb0,
	0x20, 0xf1, 0xad, 0x0a, 0xd1, 0x5b, 0x30, 0x1f, 0x28, 0x0b, 0xb4, 0x57, 0x4b, 0xb5, 0xe3, 0x14,
	0xea, 0x1a, 0xa5, 0x4b, 0x59, 0xe0, 0xe9, 0x4e, 0xdb, 0x85, 0x62, 0x23, 0x0a, 0xe8, 0x64, 0xd2,
	0xe4, 0x4c, 0x0a, 0x3e, 0x57, 0x4b, 0x71, 0x10, 0x24, 0x64, 0x55, 0xa8, 0xf8, 0x8f, 0x79, 0xb8,
	0x48, 0x3c, 0xa0, 0x63, 0x55, 0x8b, 0x08, 0x79, 0x48, 0xb0, 0x75, 0x5c, 0x7d, 0x0f, 0x47, 0x69,
	0x9f, 0x55, 0x59, 0xbf, 0x3f, 0xe8, 0xab, 0x47, 0xa0, 0xac, 0xef, 0x0d, 0x86, 0x3d, 0xd7, 0x32,
	0x54, 0xf5, 0xf2, 0xce, 0xbd, 0xb6, 0x32, 0xd5, 0x0a, 0x14, 0x5f, 0x08, 0xaa, 0x1e, 0x8f, 0xdf,
	0xa9, 0x77, 0xdb, 0xef, 0x6a, 0x1f, 0x46, 0xe7, 0x35, 0x6b, 0xab, 0xfa, 0x11, 0xf2, 0xcf, 0xbc,
	0xd5, 0xe0, 0x85, 0xdb, 0x53, 0xeb, 0x0a, 0xb0, 0xe3, 0xdf, 0x5e, 0xf5, 0xdc, 0x7e, 0xd7, 0x32,
	0xd0, 0x0e, 0x64, 0x5b, 0xae, 0x67, 0x65, 0xd4, 0x74, 0xb3, 0x37, 0xf0, 0xdb, 0xad, 0x91, 0x6e,
	0xcb, 0x36, 0x72, 0x77, 0xd9, 0x30, 0x16, 0xf7, 0xdb, 0xfa, 0x47, 0x77, 0xfe, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x44, 0xc7, 0x79, 0xad, 0xf9, 0x04, 0x00, 0x00,
}
