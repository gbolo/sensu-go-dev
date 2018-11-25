// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: adhoc.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdhocRequest struct {
	// Name is the name of the requested adhoc check.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"check"`
	// Subscriptions is the list of entity subscriptions.
	Subscriptions []string `protobuf:"bytes,2,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// Creator is the author of the adhoc request.
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// Reason is used to provide context to the request.
	Reason               string   `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdhocRequest) Reset()         { *m = AdhocRequest{} }
func (m *AdhocRequest) String() string { return proto.CompactTextString(m) }
func (*AdhocRequest) ProtoMessage()    {}
func (*AdhocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adhoc_be0c18be64c02f59, []int{0}
}
func (m *AdhocRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdhocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdhocRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AdhocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdhocRequest.Merge(dst, src)
}
func (m *AdhocRequest) XXX_Size() int {
	return m.Size()
}
func (m *AdhocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdhocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdhocRequest proto.InternalMessageInfo

func (m *AdhocRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdhocRequest) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *AdhocRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AdhocRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AdhocRequest)(nil), "sensu.api.core.v2.AdhocRequest")
}
func (this *AdhocRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AdhocRequest)
	if !ok {
		that2, ok := that.(AdhocRequest)
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
	if this.Name != that1.Name {
		return false
	}
	if len(this.Subscriptions) != len(that1.Subscriptions) {
		return false
	}
	for i := range this.Subscriptions {
		if this.Subscriptions[i] != that1.Subscriptions[i] {
			return false
		}
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *AdhocRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdhocRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAdhoc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Subscriptions) > 0 {
		for _, s := range m.Subscriptions {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Creator) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAdhoc(dAtA, i, uint64(len(m.Creator)))
		i += copy(dAtA[i:], m.Creator)
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAdhoc(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAdhoc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAdhocRequest(r randyAdhoc, easy bool) *AdhocRequest {
	this := &AdhocRequest{}
	this.Name = string(randStringAdhoc(r))
	v1 := r.Intn(10)
	this.Subscriptions = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Subscriptions[i] = string(randStringAdhoc(r))
	}
	this.Creator = string(randStringAdhoc(r))
	this.Reason = string(randStringAdhoc(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAdhoc(r, 5)
	}
	return this
}

type randyAdhoc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAdhoc(r randyAdhoc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAdhoc(r randyAdhoc) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneAdhoc(r)
	}
	return string(tmps)
}
func randUnrecognizedAdhoc(r randyAdhoc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAdhoc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAdhoc(dAtA []byte, r randyAdhoc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAdhoc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAdhoc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AdhocRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAdhoc(uint64(l))
	}
	if len(m.Subscriptions) > 0 {
		for _, s := range m.Subscriptions {
			l = len(s)
			n += 1 + l + sovAdhoc(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdhoc(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovAdhoc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdhoc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAdhoc(x uint64) (n int) {
	return sovAdhoc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AdhocRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdhoc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdhocRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdhocRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdhoc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdhoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdhoc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdhoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdhoc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdhoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdhoc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdhoc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdhoc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdhoc
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
func skipAdhoc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdhoc
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
					return 0, ErrIntOverflowAdhoc
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
					return 0, ErrIntOverflowAdhoc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAdhoc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAdhoc
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
				next, err := skipAdhoc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthAdhoc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdhoc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("adhoc.proto", fileDescriptor_adhoc_be0c18be64c02f59) }

var fileDescriptor_adhoc_be0c18be64c02f59 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0xc9, 0xc8,
	0x4f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x4b,
	0x2c, 0xc8, 0xd4, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x33, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4c, 0x2a,
	0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62, 0x82, 0xd2, 0x74, 0x46, 0x2e, 0x1e, 0x47, 0x90,
	0x89, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xb2, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x9c, 0xaf, 0xee, 0xc9, 0xb3, 0x26, 0x67, 0xa4, 0x26,
	0x67, 0x07, 0x81, 0x85, 0x85, 0xb4, 0xb8, 0x78, 0x8b, 0x4b, 0x93, 0x8a, 0x93, 0x8b, 0x32, 0x0b,
	0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0x38, 0x9d, 0x58, 0x4e, 0xdc, 0x93,
	0x67, 0x0c, 0x42, 0x95, 0x12, 0x92, 0xe3, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92,
	0x60, 0x06, 0x9b, 0x06, 0x51, 0x05, 0x13, 0x14, 0x92, 0xe1, 0x62, 0x2b, 0x4a, 0x4d, 0x2c, 0xce,
	0xcf, 0x93, 0x60, 0x41, 0x92, 0x86, 0x8a, 0x39, 0x29, 0xfc, 0x78, 0x28, 0xc7, 0xb8, 0xe2, 0x91,
	0x1c, 0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x4c, 0x65, 0x46, 0x49, 0x6c, 0x60, 0x2f, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0x44, 0x34, 0xe5, 0x13, 0x01, 0x00, 0x00,
}
