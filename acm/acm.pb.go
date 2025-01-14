// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acm.proto

package acm

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	permission "github.com/hyperledger/burrow/permission"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address"`
	PublicKey            crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey"`
	Sequence             uint64                                       `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Balance              uint64                                       `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
	EVMCode              Bytecode                                     `protobuf:"bytes,5,opt,name=EVMCode,proto3,customtype=Bytecode" json:"EVMCode"`
	Permissions          permission.AccountPermissions                `protobuf:"bytes,6,opt,name=Permissions,proto3" json:"Permissions"`
	WASMCode             Bytecode                                     `protobuf:"bytes,7,opt,name=WASMCode,proto3,customtype=Bytecode" json:",omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Account) Reset()      { *m = Account{} }
func (*Account) ProtoMessage() {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ed775bc0a6adf6, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func (m *Account) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetPermissions() permission.AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return permission.AccountPermissions{}
}

func (*Account) XXX_MessageName() string {
	return "acm.Account"
}
func init() {
	proto.RegisterType((*Account)(nil), "acm.Account")
	golang_proto.RegisterType((*Account)(nil), "acm.Account")
}

func init() { proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }
func init() { golang_proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }

var fileDescriptor_49ed775bc0a6adf6 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x3f, 0x4f, 0xc2, 0x40,
	0x1c, 0xe5, 0xa0, 0x52, 0x38, 0x18, 0xf0, 0xa6, 0x86, 0xe1, 0x8a, 0x4e, 0xc4, 0x60, 0x9b, 0xf8,
	0x67, 0xc1, 0x89, 0x1a, 0x5d, 0x8c, 0x86, 0x94, 0x44, 0x13, 0xb7, 0xf6, 0x7a, 0x96, 0x26, 0x94,
	0xab, 0xd7, 0x36, 0xa6, 0xdf, 0xc4, 0xd1, 0x8f, 0xe2, 0xc8, 0xe8, 0x68, 0x1c, 0x88, 0x29, 0x9b,
	0x9f, 0xc1, 0xc1, 0x70, 0x1c, 0xb5, 0x71, 0x70, 0xeb, 0xeb, 0x7b, 0xef, 0xf7, 0x5e, 0xde, 0xc1,
	0xa6, 0x43, 0x42, 0x23, 0xe2, 0x2c, 0x61, 0xa8, 0xe6, 0x90, 0xb0, 0x7b, 0xe8, 0x07, 0xc9, 0x34,
	0x75, 0x0d, 0xc2, 0x42, 0xd3, 0x67, 0x3e, 0x33, 0x05, 0xe7, 0xa6, 0x0f, 0x02, 0x09, 0x20, 0xbe,
	0x36, 0x9e, 0x6e, 0x27, 0xa2, 0x3c, 0x0c, 0xe2, 0x38, 0x60, 0x73, 0xf9, 0xa7, 0x4d, 0x78, 0x16,
	0x25, 0x92, 0xdf, 0xff, 0xae, 0x42, 0x75, 0x44, 0x08, 0x4b, 0xe7, 0x09, 0xba, 0x81, 0xea, 0xc8,
	0xf3, 0x38, 0x8d, 0x63, 0x0d, 0xf4, 0x40, 0xbf, 0x6d, 0x9d, 0x2c, 0x96, 0x7a, 0xe5, 0x63, 0xa9,
	0x0f, 0x4a, 0x99, 0xd3, 0x2c, 0xa2, 0x7c, 0x46, 0x3d, 0x9f, 0x72, 0xd3, 0x4d, 0x39, 0x67, 0x4f,
	0xa6, 0x3c, 0x28, 0xbd, 0xf6, 0xf6, 0x08, 0x3a, 0x85, 0xcd, 0x71, 0xea, 0xce, 0x02, 0x72, 0x45,
	0x33, 0xad, 0xda, 0x03, 0xfd, 0xd6, 0xd1, 0xae, 0x21, 0xc5, 0x05, 0x61, 0x29, 0xeb, 0x10, 0xfb,
	0x57, 0x89, 0xba, 0xb0, 0x31, 0xa1, 0x8f, 0x29, 0x9d, 0x13, 0xaa, 0xd5, 0x7a, 0xa0, 0xaf, 0xd8,
	0x05, 0x46, 0x1a, 0x54, 0x2d, 0x67, 0xe6, 0xac, 0x29, 0x45, 0x50, 0x5b, 0x88, 0x0e, 0xa0, 0x7a,
	0x71, 0x7b, 0x7d, 0xce, 0x3c, 0xaa, 0xed, 0x88, 0xf2, 0x1d, 0x59, 0xbe, 0x61, 0x65, 0x09, 0x25,
	0xcc, 0xa3, 0xf6, 0x56, 0x80, 0x2e, 0x61, 0x6b, 0x5c, 0xcc, 0x12, 0x6b, 0x75, 0x51, 0x0d, 0x1b,
	0xa5, 0xa9, 0xe4, 0x24, 0x25, 0x95, 0xec, 0x59, 0x36, 0xa2, 0x21, 0x6c, 0xdc, 0x8d, 0x26, 0x9b,
	0x50, 0x55, 0x84, 0xe2, 0xbf, 0xa1, 0x5f, 0x4b, 0x1d, 0x0e, 0x58, 0x18, 0x24, 0x34, 0x8c, 0x92,
	0xcc, 0x2e, 0xf4, 0x43, 0xe5, 0xf9, 0x45, 0xaf, 0x58, 0x67, 0x8b, 0x1c, 0x83, 0xb7, 0x1c, 0x83,
	0xf7, 0x1c, 0x83, 0xcf, 0x1c, 0x83, 0xd7, 0x15, 0x06, 0x8b, 0x15, 0x06, 0xf7, 0x7b, 0xff, 0x6f,
	0xee, 0x90, 0xd0, 0xad, 0x8b, 0x27, 0x3c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xba, 0x95, 0x25,
	0xdd, 0x23, 0x02, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Address.Size()))
	n1, err1 := m.Address.MarshalTo(dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.PublicKey.Size()))
	n2, err2 := m.PublicKey.MarshalTo(dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if m.Sequence != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Sequence))
	}
	if m.Balance != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Balance))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.EVMCode.Size()))
	n3, err3 := m.EVMCode.MarshalTo(dAtA[i:])
	if err3 != nil {
		return 0, err3
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Permissions.Size()))
	n4, err4 := m.Permissions.MarshalTo(dAtA[i:])
	if err4 != nil {
		return 0, err4
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.WASMCode.Size()))
	n5, err5 := m.WASMCode.MarshalTo(dAtA[i:])
	if err5 != nil {
		return 0, err5
	}
	i += n5
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAcm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.PublicKey.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.Sequence != 0 {
		n += 1 + sovAcm(uint64(m.Sequence))
	}
	if m.Balance != 0 {
		n += 1 + sovAcm(uint64(m.Balance))
	}
	l = m.EVMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.Permissions.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.WASMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAcm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAcm(x uint64) (n int) {
	return sovAcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EVMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WASMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAcm
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
				return 0, ErrInvalidLengthAcm
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAcm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAcm
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAcm(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAcm
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAcm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcm   = fmt.Errorf("proto: integer overflow")
)
