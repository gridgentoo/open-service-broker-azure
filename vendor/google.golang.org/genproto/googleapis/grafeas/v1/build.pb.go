// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/build.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Public key formats.
type BuildSignature_KeyType int32

const (
	// `KeyType` is not set.
	BuildSignature_KEY_TYPE_UNSPECIFIED BuildSignature_KeyType = 0
	// `PGP ASCII Armored` public key.
	BuildSignature_PGP_ASCII_ARMORED BuildSignature_KeyType = 1
	// `PKIX PEM` public key.
	BuildSignature_PKIX_PEM BuildSignature_KeyType = 2
)

var BuildSignature_KeyType_name = map[int32]string{
	0: "KEY_TYPE_UNSPECIFIED",
	1: "PGP_ASCII_ARMORED",
	2: "PKIX_PEM",
}

var BuildSignature_KeyType_value = map[string]int32{
	"KEY_TYPE_UNSPECIFIED": 0,
	"PGP_ASCII_ARMORED":    1,
	"PKIX_PEM":             2,
}

func (x BuildSignature_KeyType) String() string {
	return proto.EnumName(BuildSignature_KeyType_name, int32(x))
}

func (BuildSignature_KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_056b00916e290af1, []int{1, 0}
}

// Note holding the version of the provider's builder and the signature of the
// provenance message in the build details occurrence.
type BuildNote struct {
	// Required. Immutable. Version of the builder which produced this build.
	BuilderVersion string `protobuf:"bytes,1,opt,name=builder_version,json=builderVersion,proto3" json:"builder_version,omitempty"`
	// Signature of the build in occurrences pointing to this build note
	// containing build details.
	Signature            *BuildSignature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BuildNote) Reset()         { *m = BuildNote{} }
func (m *BuildNote) String() string { return proto.CompactTextString(m) }
func (*BuildNote) ProtoMessage()    {}
func (*BuildNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b00916e290af1, []int{0}
}

func (m *BuildNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildNote.Unmarshal(m, b)
}
func (m *BuildNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildNote.Marshal(b, m, deterministic)
}
func (m *BuildNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildNote.Merge(m, src)
}
func (m *BuildNote) XXX_Size() int {
	return xxx_messageInfo_BuildNote.Size(m)
}
func (m *BuildNote) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildNote.DiscardUnknown(m)
}

var xxx_messageInfo_BuildNote proto.InternalMessageInfo

func (m *BuildNote) GetBuilderVersion() string {
	if m != nil {
		return m.BuilderVersion
	}
	return ""
}

func (m *BuildNote) GetSignature() *BuildSignature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Message encapsulating the signature of the verified build.
type BuildSignature struct {
	// Public key of the builder which can be used to verify that the related
	// findings are valid and unchanged. If `key_type` is empty, this defaults
	// to PEM encoded public keys.
	//
	// This field may be empty if `key_id` references an external key.
	//
	// For Cloud Build based signatures, this is a PEM encoded public
	// key. To verify the Cloud Build signature, place the contents of
	// this field into a file (public.pem). The signature field is base64-decoded
	// into its binary representation in signature.bin, and the provenance bytes
	// from `BuildDetails` are base64-decoded into a binary representation in
	// signed.bin. OpenSSL can then verify the signature:
	// `openssl sha256 -verify public.pem -signature signature.bin signed.bin`
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Required. Signature of the related `BuildProvenance`. In JSON, this is
	// base-64 encoded.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// An ID for the key used to sign. This could be either an ID for the key
	// stored in `public_key` (such as the ID or fingerprint for a PGP key, or the
	// CN for a cert), or a reference to an external key (such as a reference to a
	// key in Cloud Key Management Service).
	KeyId string `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// The type of the key, either stored in `public_key` or referenced in
	// `key_id`.
	KeyType              BuildSignature_KeyType `protobuf:"varint,4,opt,name=key_type,json=keyType,proto3,enum=grafeas.v1.BuildSignature_KeyType" json:"key_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BuildSignature) Reset()         { *m = BuildSignature{} }
func (m *BuildSignature) String() string { return proto.CompactTextString(m) }
func (*BuildSignature) ProtoMessage()    {}
func (*BuildSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b00916e290af1, []int{1}
}

func (m *BuildSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildSignature.Unmarshal(m, b)
}
func (m *BuildSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildSignature.Marshal(b, m, deterministic)
}
func (m *BuildSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildSignature.Merge(m, src)
}
func (m *BuildSignature) XXX_Size() int {
	return xxx_messageInfo_BuildSignature.Size(m)
}
func (m *BuildSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildSignature.DiscardUnknown(m)
}

var xxx_messageInfo_BuildSignature proto.InternalMessageInfo

func (m *BuildSignature) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *BuildSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *BuildSignature) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *BuildSignature) GetKeyType() BuildSignature_KeyType {
	if m != nil {
		return m.KeyType
	}
	return BuildSignature_KEY_TYPE_UNSPECIFIED
}

// Details of a build occurrence.
type BuildOccurrence struct {
	// Required. The actual provenance for the build.
	Provenance *BuildProvenance `protobuf:"bytes,1,opt,name=provenance,proto3" json:"provenance,omitempty"`
	// Serialized JSON representation of the provenance, used in generating the
	// build signature in the corresponding build note. After verifying the
	// signature, `provenance_bytes` can be unmarshalled and compared to the
	// provenance to confirm that it is unchanged. A base64-encoded string
	// representation of the provenance bytes is used for the signature in order
	// to interoperate with openssl which expects this format for signature
	// verification.
	//
	// The serialized form is captured both to avoid ambiguity in how the
	// provenance is marshalled to json as well to prevent incompatibilities with
	// future changes.
	ProvenanceBytes      string   `protobuf:"bytes,2,opt,name=provenance_bytes,json=provenanceBytes,proto3" json:"provenance_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildOccurrence) Reset()         { *m = BuildOccurrence{} }
func (m *BuildOccurrence) String() string { return proto.CompactTextString(m) }
func (*BuildOccurrence) ProtoMessage()    {}
func (*BuildOccurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b00916e290af1, []int{2}
}

func (m *BuildOccurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildOccurrence.Unmarshal(m, b)
}
func (m *BuildOccurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildOccurrence.Marshal(b, m, deterministic)
}
func (m *BuildOccurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildOccurrence.Merge(m, src)
}
func (m *BuildOccurrence) XXX_Size() int {
	return xxx_messageInfo_BuildOccurrence.Size(m)
}
func (m *BuildOccurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildOccurrence.DiscardUnknown(m)
}

var xxx_messageInfo_BuildOccurrence proto.InternalMessageInfo

func (m *BuildOccurrence) GetProvenance() *BuildProvenance {
	if m != nil {
		return m.Provenance
	}
	return nil
}

func (m *BuildOccurrence) GetProvenanceBytes() string {
	if m != nil {
		return m.ProvenanceBytes
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1.BuildSignature_KeyType", BuildSignature_KeyType_name, BuildSignature_KeyType_value)
	proto.RegisterType((*BuildNote)(nil), "grafeas.v1.BuildNote")
	proto.RegisterType((*BuildSignature)(nil), "grafeas.v1.BuildSignature")
	proto.RegisterType((*BuildOccurrence)(nil), "grafeas.v1.BuildOccurrence")
}

func init() { proto.RegisterFile("grafeas/v1/build.proto", fileDescriptor_056b00916e290af1) }

var fileDescriptor_056b00916e290af1 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x0a, 0x6d, 0x3d, 0x2d, 0x49, 0x58, 0x51, 0x64, 0xb5, 0x20, 0x45, 0xbe, 0x10,
	0x2e, 0xb6, 0x1a, 0x2e, 0x95, 0x2a, 0x0e, 0x49, 0x6b, 0x8a, 0x15, 0xb5, 0x5d, 0x36, 0x05, 0x11,
	0x2e, 0x2b, 0xc7, 0x19, 0x56, 0x56, 0xa2, 0x5d, 0x6b, 0xed, 0x44, 0xda, 0xd7, 0xe1, 0xfd, 0x78,
	0x07, 0x94, 0xb5, 0xc1, 0x06, 0xc4, 0x6d, 0xfc, 0xcd, 0x3f, 0xff, 0x3f, 0xb6, 0x07, 0x5e, 0x08,
	0x9d, 0x7c, 0xc3, 0xa4, 0x08, 0xb7, 0xe7, 0xe1, 0x62, 0x93, 0xad, 0x97, 0x41, 0xae, 0x55, 0xa9,
	0x08, 0xd4, 0x3c, 0xd8, 0x9e, 0x9f, 0x9e, 0xb5, 0x34, 0xb9, 0x56, 0x5b, 0x94, 0x89, 0x4c, 0xb1,
	0x12, 0xfa, 0x12, 0xdc, 0xc9, 0x6e, 0xee, 0x4e, 0x95, 0x48, 0x5e, 0x43, 0xcf, 0x9a, 0xa0, 0xe6,
	0x5b, 0xd4, 0x45, 0xa6, 0xa4, 0xe7, 0x0c, 0x9c, 0xa1, 0xcb, 0xba, 0x35, 0xfe, 0x5c, 0x51, 0x72,
	0x01, 0x6e, 0x91, 0x09, 0x99, 0x94, 0x1b, 0x8d, 0x5e, 0x67, 0xe0, 0x0c, 0x8f, 0x46, 0xa7, 0x41,
	0x13, 0x19, 0x58, 0xcb, 0xd9, 0x2f, 0x05, 0x6b, 0xc4, 0xfe, 0x0f, 0x07, 0xba, 0x7f, 0x76, 0xc9,
	0x2b, 0x80, 0x7c, 0xb3, 0x58, 0x67, 0x29, 0x5f, 0xa1, 0xa9, 0x03, 0xdd, 0x8a, 0x4c, 0xd1, 0x90,
	0x97, 0x7f, 0x67, 0x1d, 0xb7, 0xfc, 0xc8, 0x09, 0xec, 0xaf, 0xd0, 0xf0, 0x6c, 0xe9, 0xed, 0xd9,
	0xc1, 0x27, 0x2b, 0x34, 0xf1, 0x92, 0xbc, 0x83, 0xc3, 0x1d, 0x2e, 0x4d, 0x8e, 0xde, 0xe3, 0x81,
	0x33, 0xec, 0x8e, 0xfc, 0xff, 0xef, 0x17, 0x4c, 0xd1, 0x3c, 0x98, 0x1c, 0xd9, 0xc1, 0xaa, 0x2a,
	0xfc, 0x0f, 0x70, 0x50, 0x33, 0xe2, 0xc1, 0xf3, 0x69, 0x34, 0xe7, 0x0f, 0x73, 0x1a, 0xf1, 0x4f,
	0x77, 0x33, 0x1a, 0x5d, 0xc5, 0xef, 0xe3, 0xe8, 0xba, 0xff, 0x88, 0x9c, 0xc0, 0x33, 0x7a, 0x43,
	0xf9, 0x78, 0x76, 0x15, 0xc7, 0x7c, 0xcc, 0x6e, 0xef, 0x59, 0x74, 0xdd, 0x77, 0xc8, 0x31, 0x1c,
	0xd2, 0x69, 0xfc, 0x85, 0xd3, 0xe8, 0xb6, 0xdf, 0xf1, 0x0d, 0xf4, 0x6c, 0xd8, 0x7d, 0x9a, 0x6e,
	0xb4, 0x46, 0x99, 0x22, 0xb9, 0x04, 0x68, 0x7e, 0x83, 0x7d, 0xdf, 0xa3, 0xd1, 0xd9, 0x3f, 0xdb,
	0xd1, 0xdf, 0x12, 0xd6, 0x92, 0x93, 0x37, 0xd0, 0x6f, 0x9e, 0xf8, 0xc2, 0x94, 0x58, 0xd8, 0x8f,
	0xe2, 0xb2, 0x5e, 0xc3, 0x27, 0x3b, 0x3c, 0xf9, 0x08, 0x4f, 0x33, 0xd5, 0xf2, 0xa5, 0xce, 0xd7,
	0x0b, 0xa1, 0x94, 0x58, 0x63, 0x20, 0xd4, 0x3a, 0x91, 0x22, 0x50, 0x5a, 0x84, 0x02, 0xa5, 0xbd,
	0x84, 0xb0, 0x6a, 0x25, 0x79, 0x56, 0x84, 0xcd, 0xc1, 0x5c, 0xd6, 0xe5, 0xf7, 0xce, 0xde, 0x0d,
	0x1b, 0x2f, 0xf6, 0xad, 0xf4, 0xed, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x6f, 0x32, 0x3d,
	0x77, 0x02, 0x00, 0x00,
}
