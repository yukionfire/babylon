// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/bls_key.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_crypto_bls12381 "github.com/babylonchain/babylon/crypto/bls12381"
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

// BlsKey wraps BLS public key with PoP
type BlsKey struct {
	// pubkey is the BLS public key of a validator
	Pubkey *github_com_babylonchain_babylon_crypto_bls12381.PublicKey `protobuf:"bytes,1,opt,name=pubkey,proto3,customtype=github.com/babylonchain/babylon/crypto/bls12381.PublicKey" json:"pubkey,omitempty"`
	// pop is the proof-of-possession of the BLS key
	Pop *ProofOfPossession `protobuf:"bytes,2,opt,name=pop,proto3" json:"pop,omitempty"`
}

func (m *BlsKey) Reset()         { *m = BlsKey{} }
func (m *BlsKey) String() string { return proto.CompactTextString(m) }
func (*BlsKey) ProtoMessage()    {}
func (*BlsKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7e926461cc70111, []int{0}
}
func (m *BlsKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlsKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlsKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlsKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlsKey.Merge(m, src)
}
func (m *BlsKey) XXX_Size() int {
	return m.Size()
}
func (m *BlsKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BlsKey.DiscardUnknown(m)
}

var xxx_messageInfo_BlsKey proto.InternalMessageInfo

func (m *BlsKey) GetPop() *ProofOfPossession {
	if m != nil {
		return m.Pop
	}
	return nil
}

// ProofOfPossession defines proof for the ownership of Ed25519 and BLS private keys
type ProofOfPossession struct {
	// bls_sig is calculated by encrypt(key = BLS_sk, data = encrypt(key = Ed25519_sk, data = Secp256k1_pk))
	BlsSig *github_com_babylonchain_babylon_crypto_bls12381.Signature `protobuf:"bytes,1,opt,name=bls_sig,json=blsSig,proto3,customtype=github.com/babylonchain/babylon/crypto/bls12381.Signature" json:"bls_sig,omitempty"`
}

func (m *ProofOfPossession) Reset()         { *m = ProofOfPossession{} }
func (m *ProofOfPossession) String() string { return proto.CompactTextString(m) }
func (*ProofOfPossession) ProtoMessage()    {}
func (*ProofOfPossession) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7e926461cc70111, []int{1}
}
func (m *ProofOfPossession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOfPossession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOfPossession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOfPossession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOfPossession.Merge(m, src)
}
func (m *ProofOfPossession) XXX_Size() int {
	return m.Size()
}
func (m *ProofOfPossession) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOfPossession.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOfPossession proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BlsKey)(nil), "babylon.checkpointing.v1.BlsKey")
	proto.RegisterType((*ProofOfPossession)(nil), "babylon.checkpointing.v1.ProofOfPossession")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/bls_key.proto", fileDescriptor_a7e926461cc70111)
}

var fileDescriptor_a7e926461cc70111 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0xc9, 0xcc,
	0x4b, 0xd7, 0x4f, 0xca, 0x29, 0x8e, 0xcf, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x80, 0x2a, 0xd2, 0x43, 0x51, 0xa4, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xa4, 0x0f, 0x62, 0x41, 0xd4, 0x2b, 0xcd, 0x63, 0xe4, 0x62, 0x73, 0xca, 0x29, 0xf6, 0x4e,
	0xad, 0x14, 0x0a, 0xe5, 0x62, 0x2b, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x71, 0xb2, 0xbd, 0x75, 0x4f, 0xde, 0x32, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0x6a, 0x72, 0x72, 0x46, 0x62, 0x66, 0x9e, 0x3e, 0xdc, 0x2d, 0x45, 0x95, 0x05,
	0x25, 0xf9, 0x20, 0x47, 0x18, 0x1a, 0x19, 0x5b, 0x18, 0xea, 0x05, 0x94, 0x26, 0xe5, 0x64, 0x26,
	0x7b, 0xa7, 0x56, 0x06, 0x41, 0x0d, 0x13, 0xb2, 0xe5, 0x62, 0x2e, 0xc8, 0x2f, 0x90, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x36, 0xd2, 0xd6, 0xc3, 0xe5, 0x3e, 0xbd, 0x80, 0xa2, 0xfc, 0xfc, 0x34, 0xff,
	0xb4, 0x80, 0xfc, 0xe2, 0xe2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0xbc, 0x20, 0x90, 0x3e, 0xa5, 0x6c,
	0x2e, 0x41, 0x0c, 0x19, 0xa1, 0x30, 0x2e, 0x76, 0x90, 0xb7, 0x8b, 0x33, 0xd3, 0x29, 0x71, 0x6b,
	0x70, 0x66, 0x7a, 0x5e, 0x62, 0x49, 0x69, 0x51, 0x6a, 0x10, 0x5b, 0x52, 0x4e, 0x71, 0x70, 0x66,
	0xba, 0x93, 0xff, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0x12, 0x32,
	0xbc, 0x02, 0x2d, 0x5a, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xa1, 0x6c, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x85, 0x5d, 0xf5, 0xbc, 0x01, 0x00, 0x00,
}

func (m *BlsKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlsKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlsKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pop != nil {
		{
			size, err := m.Pop.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Pubkey != nil {
		{
			size := m.Pubkey.Size()
			i -= size
			if _, err := m.Pubkey.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofOfPossession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOfPossession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOfPossession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlsSig != nil {
		{
			size := m.BlsSig.Size()
			i -= size
			if _, err := m.BlsSig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlsKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlsKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlsKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	if m.Pop != nil {
		l = m.Pop.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	return n
}

func (m *ProofOfPossession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlsSig != nil {
		l = m.BlsSig.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	return n
}

func sovBlsKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlsKey(x uint64) (n int) {
	return sovBlsKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlsKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: BlsKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlsKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_crypto_bls12381.PublicKey
			m.Pubkey = &v
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pop == nil {
				m.Pop = &ProofOfPossession{}
			}
			if err := m.Pop.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func (m *ProofOfPossession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: ProofOfPossession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOfPossession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_crypto_bls12381.Signature
			m.BlsSig = &v
			if err := m.BlsSig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func skipBlsKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlsKey
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
					return 0, ErrIntOverflowBlsKey
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
					return 0, ErrIntOverflowBlsKey
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
				return 0, ErrInvalidLengthBlsKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlsKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlsKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlsKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlsKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlsKey = fmt.Errorf("proto: unexpected end of group")
)