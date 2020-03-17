// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oracle_artifact_provider.proto

package proto

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

// Configuration for the Oracle artifact provider.
type OracleArtifactProvider struct {
	// Whether the Oracle artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Oracle artifact accounts.
	Accounts             []*OracleArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *OracleArtifactProvider) Reset()         { *m = OracleArtifactProvider{} }
func (m *OracleArtifactProvider) String() string { return proto.CompactTextString(m) }
func (*OracleArtifactProvider) ProtoMessage()    {}
func (*OracleArtifactProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30a104913dd69e6, []int{0}
}

func (m *OracleArtifactProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleArtifactProvider.Unmarshal(m, b)
}
func (m *OracleArtifactProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleArtifactProvider.Marshal(b, m, deterministic)
}
func (m *OracleArtifactProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleArtifactProvider.Merge(m, src)
}
func (m *OracleArtifactProvider) XXX_Size() int {
	return xxx_messageInfo_OracleArtifactProvider.Size(m)
}
func (m *OracleArtifactProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleArtifactProvider.DiscardUnknown(m)
}

var xxx_messageInfo_OracleArtifactProvider proto.InternalMessageInfo

func (m *OracleArtifactProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *OracleArtifactProvider) GetAccounts() []*OracleArtifactAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Configuration for an Oracle artifact account.
type OracleArtifactAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// The namespace in which the bucket and objects should be created.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,4,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// An Oracle region (e.g., `us-phoenix-1`).
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,6,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,7,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
	// The OCID of the Oracle User with which to authenticate.
	UserId               string   `protobuf:"bytes,8,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleArtifactAccount) Reset()         { *m = OracleArtifactAccount{} }
func (m *OracleArtifactAccount) String() string { return proto.CompactTextString(m) }
func (*OracleArtifactAccount) ProtoMessage()    {}
func (*OracleArtifactAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30a104913dd69e6, []int{1}
}

func (m *OracleArtifactAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleArtifactAccount.Unmarshal(m, b)
}
func (m *OracleArtifactAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleArtifactAccount.Marshal(b, m, deterministic)
}
func (m *OracleArtifactAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleArtifactAccount.Merge(m, src)
}
func (m *OracleArtifactAccount) XXX_Size() int {
	return xxx_messageInfo_OracleArtifactAccount.Size(m)
}
func (m *OracleArtifactAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleArtifactAccount.DiscardUnknown(m)
}

var xxx_messageInfo_OracleArtifactAccount proto.InternalMessageInfo

func (m *OracleArtifactAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OracleArtifactAccount) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *OracleArtifactAccount) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *OracleArtifactAccount) GetPrivateKeyPassphrase() string {
	if m != nil {
		return m.PrivateKeyPassphrase
	}
	return ""
}

func (m *OracleArtifactAccount) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *OracleArtifactAccount) GetSshPrivateKeyFilePath() string {
	if m != nil {
		return m.SshPrivateKeyFilePath
	}
	return ""
}

func (m *OracleArtifactAccount) GetTenancyId() string {
	if m != nil {
		return m.TenancyId
	}
	return ""
}

func (m *OracleArtifactAccount) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*OracleArtifactProvider)(nil), "proto.OracleArtifactProvider")
	proto.RegisterType((*OracleArtifactAccount)(nil), "proto.OracleArtifactAccount")
}

func init() { proto.RegisterFile("oracle_artifact_provider.proto", fileDescriptor_c30a104913dd69e6) }

var fileDescriptor_c30a104913dd69e6 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x6a, 0xf3, 0x30,
	0x14, 0xc5, 0xc9, 0x7f, 0xe7, 0x66, 0x13, 0x5f, 0x82, 0x86, 0xf0, 0x61, 0x32, 0x65, 0xca, 0x90,
	0x76, 0xe8, 0x9a, 0xa5, 0x10, 0x3a, 0xd4, 0xf8, 0x05, 0xc2, 0x8d, 0x7c, 0x13, 0x0b, 0x5c, 0x49,
	0x5c, 0x29, 0x81, 0x3c, 0x4f, 0x5f, 0xb4, 0x58, 0x76, 0xdd, 0x16, 0x3c, 0x49, 0xe7, 0x9c, 0x9f,
	0xee, 0xb9, 0x08, 0xfe, 0x5b, 0x46, 0x55, 0xd1, 0x09, 0x39, 0xe8, 0x0b, 0xaa, 0x70, 0x72, 0x6c,
	0xef, 0xba, 0x20, 0xde, 0x39, 0xb6, 0xc1, 0x8a, 0x49, 0x3c, 0x36, 0x15, 0xac, 0xde, 0x23, 0x78,
	0x68, 0xb9, 0xac, 0xc5, 0x84, 0x84, 0x19, 0x19, 0x3c, 0x57, 0x54, 0xc8, 0x41, 0x3a, 0xd8, 0x26,
	0xf9, 0xb7, 0x14, 0x2f, 0x90, 0xa0, 0x52, 0xf6, 0x66, 0x82, 0x97, 0xc3, 0x74, 0xb4, 0x5d, 0xec,
	0xd7, 0xcd, 0xd0, 0xdd, 0xdf, 0x51, 0x87, 0x06, 0xca, 0x3b, 0x7a, 0xf3, 0x39, 0x84, 0x65, 0x2f,
	0x23, 0x04, 0x8c, 0x0d, 0x7e, 0x50, 0xac, 0x9a, 0xe7, 0xf1, 0x2e, 0x52, 0x58, 0x5c, 0xb4, 0xb9,
	0x12, 0x3b, 0xd6, 0x26, 0xc8, 0x61, 0x8c, 0x7e, 0x5b, 0x62, 0x0d, 0xf3, 0x9a, 0xf4, 0x0e, 0x15,
	0xc9, 0x51, 0xcc, 0x7f, 0x0c, 0xb1, 0x87, 0x7f, 0x8e, 0xf5, 0x1d, 0x03, 0xbd, 0xd1, 0x23, 0x43,
	0xef, 0x5d, 0xc9, 0xe8, 0x49, 0x8e, 0x23, 0xd8, 0x9b, 0x89, 0x15, 0x4c, 0x99, 0xae, 0xda, 0x1a,
	0x39, 0x89, 0x54, 0xab, 0xc4, 0x33, 0x2c, 0xbd, 0x2f, 0xb3, 0xee, 0xc9, 0xab, 0xae, 0x28, 0xc3,
	0x50, 0xca, 0x69, 0xc4, 0xfa, 0xc3, 0x7a, 0xbf, 0x40, 0x06, 0x8d, 0x7a, 0x1c, 0x0b, 0x39, 0x6b,
	0xf6, 0xeb, 0x8c, 0xba, 0xeb, 0xe6, 0x89, 0x8f, 0x85, 0x4c, 0x9a, 0xae, 0x46, 0x9d, 0xa7, 0xf1,
	0x33, 0x9f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x01, 0x91, 0x79, 0xc3, 0x01, 0x00, 0x00,
}