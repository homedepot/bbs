// Code generated by protoc-gen-gogo.
// source: error.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import io "io"
import fmt "fmt"

import strings "strings"
import reflect "reflect"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Error_Type int32

const (
	Error_UnknownError               Error_Type = 0
	Error_InvalidDomain              Error_Type = 1
	Error_UnkownVersion              Error_Type = 2
	Error_InvalidRecord              Error_Type = 3
	Error_InvalidRequest             Error_Type = 4
	Error_InvalidResponse            Error_Type = 5
	Error_InvalidProtobufMessage     Error_Type = 6
	Error_InvalidJson                Error_Type = 7
	Error_FailedToOpenEnvelope       Error_Type = 8
	Error_InvalidStateTransition     Error_Type = 9
	Error_Unauthorized               Error_Type = 10
	Error_ResourceConflict           Error_Type = 11
	Error_ResourceExists             Error_Type = 12
	Error_ResourceNotFound           Error_Type = 13
	Error_RouterError                Error_Type = 14
	Error_ActualLRPCannotBeClaimed   Error_Type = 15
	Error_ActualLRPCannotBeStarted   Error_Type = 16
	Error_ActualLRPCannotBeCrashed   Error_Type = 17
	Error_ActualLRPCannotBeFailed    Error_Type = 18
	Error_ActualLRPCannotBeRemoved   Error_Type = 19
	Error_ActualLRPCannotBeStopped   Error_Type = 20
	Error_ActualLRPCannotBeUnclaimed Error_Type = 21
	Error_ActualLRPCannotBeEvacuated Error_Type = 22
	Error_DesiredLRPCannotBeUpdated  Error_Type = 23
	Error_RunningOnDifferentCell     Error_Type = 24
	Error_NetworkError               Error_Type = 25
)

var Error_Type_name = map[int32]string{
	0:  "UnknownError",
	1:  "InvalidDomain",
	2:  "UnkownVersion",
	3:  "InvalidRecord",
	4:  "InvalidRequest",
	5:  "InvalidResponse",
	6:  "InvalidProtobufMessage",
	7:  "InvalidJson",
	8:  "FailedToOpenEnvelope",
	9:  "InvalidStateTransition",
	10: "Unauthorized",
	11: "ResourceConflict",
	12: "ResourceExists",
	13: "ResourceNotFound",
	14: "RouterError",
	15: "ActualLRPCannotBeClaimed",
	16: "ActualLRPCannotBeStarted",
	17: "ActualLRPCannotBeCrashed",
	18: "ActualLRPCannotBeFailed",
	19: "ActualLRPCannotBeRemoved",
	20: "ActualLRPCannotBeStopped",
	21: "ActualLRPCannotBeUnclaimed",
	22: "ActualLRPCannotBeEvacuated",
	23: "DesiredLRPCannotBeUpdated",
	24: "RunningOnDifferentCell",
	25: "NetworkError",
}
var Error_Type_value = map[string]int32{
	"UnknownError":               0,
	"InvalidDomain":              1,
	"UnkownVersion":              2,
	"InvalidRecord":              3,
	"InvalidRequest":             4,
	"InvalidResponse":            5,
	"InvalidProtobufMessage":     6,
	"InvalidJson":                7,
	"FailedToOpenEnvelope":       8,
	"InvalidStateTransition":     9,
	"Unauthorized":               10,
	"ResourceConflict":           11,
	"ResourceExists":             12,
	"ResourceNotFound":           13,
	"RouterError":                14,
	"ActualLRPCannotBeClaimed":   15,
	"ActualLRPCannotBeStarted":   16,
	"ActualLRPCannotBeCrashed":   17,
	"ActualLRPCannotBeFailed":    18,
	"ActualLRPCannotBeRemoved":   19,
	"ActualLRPCannotBeStopped":   20,
	"ActualLRPCannotBeUnclaimed": 21,
	"ActualLRPCannotBeEvacuated": 22,
	"DesiredLRPCannotBeUpdated":  23,
	"RunningOnDifferentCell":     24,
	"NetworkError":               25,
}

func (x Error_Type) Enum() *Error_Type {
	p := new(Error_Type)
	*p = x
	return p
}
func (x Error_Type) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(Error_Type_name, int32(x))
}
func (x *Error_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Type_value, data, "Error_Type")
	if err != nil {
		return err
	}
	*x = Error_Type(value)
	return nil
}

type Error struct {
	Type    Error_Type `protobuf:"varint,1,opt,name=type,enum=models.Error_Type" json:"type"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_UnknownError
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("models.Error_Type", Error_Type_name, Error_Type_value)
}
func (m *Error) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Error_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipError(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipError(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipError(data[start:])
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
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
)

func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringError(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovError(uint64(m.Type))
	l = len(m.Message)
	n += 1 + l + sovError(uint64(l))
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Error) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintError(data, i, uint64(m.Type))
	data[i] = 0x12
	i++
	i = encodeVarintError(data, i, uint64(len(m.Message)))
	i += copy(data[i:], m.Message)
	return i, nil
}

func encodeFixed64Error(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Error(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintError(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.Error{` +
		`Type:` + fmt.Sprintf("%#v", this.Type),
		`Message:` + fmt.Sprintf("%#v", this.Message) + `}`}, ", ")
	return s
}
func valueToGoStringError(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringError(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (x Error_Type) String() string {
	s, ok := Error_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
