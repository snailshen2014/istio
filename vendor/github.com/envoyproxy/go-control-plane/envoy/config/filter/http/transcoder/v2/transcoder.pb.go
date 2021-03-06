// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_07004c00c48edf58, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(dst, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GrpcJsonTranscoder_OneofMarshaler, _GrpcJsonTranscoder_OneofUnmarshaler, _GrpcJsonTranscoder_OneofSizer, []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

func _GrpcJsonTranscoder_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.ProtoDescriptorBin)
	case nil:
	default:
		return fmt.Errorf("GrpcJsonTranscoder.DescriptorSet has unexpected type %T", x)
	}
	return nil
}

func _GrpcJsonTranscoder_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GrpcJsonTranscoder)
	switch tag {
	case 1: // descriptor_set.proto_descriptor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{x}
		return true, err
	case 4: // descriptor_set.proto_descriptor_bin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{x}
		return true, err
	default:
		return false, nil
	}
}

func _GrpcJsonTranscoder_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptor)))
		n += len(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptorBin)))
		n += len(x.ProtoDescriptorBin)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_07004c00c48edf58, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(dst, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}
func (m *GrpcJsonTranscoder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DescriptorSet != nil {
		nn1, err := m.DescriptorSet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
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
	if m.PrintOptions != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTranscoder(dAtA, i, uint64(m.PrintOptions.Size()))
		n2, err := m.PrintOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.MatchIncomingRequestRoute {
		dAtA[i] = 0x28
		i++
		if m.MatchIncomingRequestRoute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for _, s := range m.IgnoredQueryParameters {
			dAtA[i] = 0x32
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
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptor)))
	i += copy(dAtA[i:], m.ProtoDescriptor)
	return i, nil
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ProtoDescriptorBin != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptorBin)))
		i += copy(dAtA[i:], m.ProtoDescriptorBin)
	}
	return i, nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AddWhitespace {
		dAtA[i] = 0x8
		i++
		if m.AddWhitespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AlwaysPrintPrimitiveFields {
		dAtA[i] = 0x10
		i++
		if m.AlwaysPrintPrimitiveFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AlwaysPrintEnumsAsInts {
		dAtA[i] = 0x18
		i++
		if m.AlwaysPrintEnumsAsInts {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.PreserveProtoFieldNames {
		dAtA[i] = 0x20
		i++
		if m.PreserveProtoFieldNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTranscoder(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GrpcJsonTranscoder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DescriptorSet != nil {
		n += m.DescriptorSet.Size()
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.PrintOptions != nil {
		l = m.PrintOptions.Size()
		n += 1 + l + sovTranscoder(uint64(l))
	}
	if m.MatchIncomingRequestRoute {
		n += 2
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for _, s := range m.IgnoredQueryParameters {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptor)
	n += 1 + l + sovTranscoder(uint64(l))
	return n
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProtoDescriptorBin != nil {
		l = len(m.ProtoDescriptorBin)
		n += 1 + l + sovTranscoder(uint64(l))
	}
	return n
}
func (m *GrpcJsonTranscoder_PrintOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddWhitespace {
		n += 2
	}
	if m.AlwaysPrintPrimitiveFields {
		n += 2
	}
	if m.AlwaysPrintEnumsAsInts {
		n += 2
	}
	if m.PreserveProtoFieldNames {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTranscoder(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTranscoder(x uint64) (n int) {
	return sovTranscoder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcJsonTranscoder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: GrpcJsonTranscoder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcJsonTranscoder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrintOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrintOptions == nil {
				m.PrintOptions = &GrpcJsonTranscoder_PrintOptions{}
			}
			if err := m.PrintOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptorBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{v}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchIncomingRequestRoute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MatchIncomingRequestRoute = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredQueryParameters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredQueryParameters = append(m.IgnoredQueryParameters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
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
func (m *GrpcJsonTranscoder_PrintOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: PrintOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrintOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddWhitespace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddWhitespace = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintPrimitiveFields", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintPrimitiveFields = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintEnumsAsInts", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintEnumsAsInts = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreserveProtoFieldNames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreserveProtoFieldNames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
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
func skipTranscoder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
				return 0, ErrInvalidLengthTranscoder
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTranscoder
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
				next, err := skipTranscoder(dAtA[start:])
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
	ErrInvalidLengthTranscoder = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTranscoder   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_transcoder_07004c00c48edf58)
}

var fileDescriptor_transcoder_07004c00c48edf58 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x3b, 0x8d, 0x29, 0xdb, 0xb1, 0xad, 0x32, 0x48, 0xb3, 0x06, 0x0d, 0x41, 0xb0, 0x04,
	0x84, 0x5d, 0x88, 0x82, 0xa2, 0x07, 0x69, 0x50, 0x9b, 0x7a, 0xd0, 0x38, 0x08, 0x82, 0x97, 0x61,
	0xba, 0xfb, 0x92, 0x0c, 0x64, 0x67, 0xa6, 0x33, 0x93, 0xad, 0xf9, 0x1a, 0xfa, 0x65, 0xc4, 0x53,
	0x8f, 0x1e, 0xfd, 0x08, 0x92, 0x5b, 0x8f, 0x7e, 0x03, 0x99, 0xd9, 0x26, 0xac, 0xf5, 0xd2, 0xdb,
	0xce, 0xfe, 0x7f, 0xff, 0xf7, 0xe6, 0xbd, 0xff, 0xe0, 0xa7, 0x20, 0x4b, 0xb5, 0x48, 0x33, 0x25,
	0xc7, 0x62, 0x92, 0x8e, 0xc5, 0xcc, 0x81, 0x49, 0xa7, 0xce, 0xe9, 0xd4, 0x19, 0x2e, 0x6d, 0xa6,
	0x72, 0x30, 0x69, 0xd9, 0xaf, 0x9d, 0x12, 0x6d, 0x94, 0x53, 0xe4, 0x20, 0x18, 0x93, 0xca, 0x98,
	0x54, 0xc6, 0xc4, 0x1b, 0x93, 0x1a, 0x5a, 0xf6, 0xdb, 0xad, 0x92, 0xcf, 0x44, 0xce, 0x1d, 0xa4,
	0xab, 0x8f, 0xaa, 0xc0, 0x83, 0x6f, 0x4d, 0x4c, 0x8e, 0x8c, 0xce, 0xde, 0x5a, 0x25, 0x3f, 0xae,
	0x2d, 0xe4, 0x11, 0xbe, 0x1d, 0x74, 0x96, 0x83, 0xcd, 0x8c, 0xd0, 0x4e, 0x99, 0x18, 0x75, 0x51,
	0x6f, 0x7b, 0xb8, 0x41, 0x6f, 0x05, 0xe5, 0xd5, 0x5a, 0x20, 0x07, 0x38, 0xb2, 0x60, 0x4a, 0x91,
	0x81, 0x8d, 0x37, 0xbb, 0x8d, 0xde, 0xf6, 0x00, 0xff, 0xb8, 0x38, 0x6f, 0x34, 0xbf, 0xa2, 0xcd,
	0x08, 0xd1, 0xb5, 0x46, 0x66, 0x78, 0x57, 0x1b, 0x21, 0x1d, 0x53, 0xda, 0x09, 0x25, 0x6d, 0xdc,
	0xe8, 0xa2, 0xde, 0xcd, 0xfe, 0x51, 0x72, 0xbd, 0x21, 0x92, 0xff, 0xef, 0x99, 0x8c, 0x7c, 0xbd,
	0xf7, 0x55, 0x39, 0xba, 0xa3, 0x6b, 0x27, 0xd2, 0xc7, 0x77, 0xae, 0x8e, 0xc0, 0x4e, 0x84, 0x8c,
	0x6f, 0x74, 0x51, 0x6f, 0x67, 0xb8, 0x41, 0xc9, 0x95, 0x31, 0x06, 0x42, 0x92, 0x97, 0xf8, 0x5e,
	0xc1, 0x5d, 0x36, 0x65, 0x42, 0x66, 0xaa, 0x10, 0x72, 0xc2, 0x0c, 0x9c, 0xce, 0xc1, 0x3a, 0x66,
	0xd4, 0xdc, 0x41, 0xdc, 0xec, 0xa2, 0x5e, 0x44, 0xef, 0x06, 0xe6, 0xf8, 0x12, 0xa1, 0x15, 0x41,
	0x3d, 0x40, 0x9e, 0xe1, 0x58, 0x4c, 0xa4, 0x32, 0x90, 0xb3, 0xd3, 0x39, 0x98, 0x05, 0xd3, 0xdc,
	0xf0, 0x02, 0x1c, 0x18, 0x1b, 0x6f, 0xf9, 0xd5, 0xd0, 0xfd, 0x4b, 0xfd, 0x83, 0x97, 0x47, 0x6b,
	0xb5, 0xfd, 0x07, 0xe1, 0x9d, 0xfa, 0x34, 0xe4, 0x21, 0xde, 0xe3, 0x79, 0xce, 0xce, 0xa6, 0xc2,
	0x81, 0xd5, 0x3c, 0x83, 0x10, 0x40, 0x44, 0x77, 0x79, 0x9e, 0x7f, 0x5a, 0xff, 0x24, 0x87, 0xf8,
	0x3e, 0x9f, 0x9d, 0xf1, 0x85, 0x65, 0xd5, 0x6e, 0xb5, 0x11, 0x85, 0x70, 0xa2, 0x04, 0x36, 0x16,
	0x30, 0xcb, 0x7d, 0x22, 0xde, 0xd5, 0xae, 0xa0, 0xd0, 0x61, 0xb4, 0x42, 0xde, 0x04, 0x82, 0x3c,
	0xc7, 0xed, 0x7f, 0x4a, 0x80, 0x9c, 0x17, 0x96, 0x71, 0xcb, 0x84, 0x74, 0x55, 0x48, 0x11, 0xdd,
	0xaf, 0xf9, 0x5f, 0x7b, 0xfd, 0xd0, 0x1e, 0x4b, 0x67, 0xc9, 0x0b, 0xdc, 0xd6, 0x06, 0x7c, 0xc4,
	0xc0, 0xaa, 0x75, 0x87, 0xb6, 0x4c, 0xf2, 0x02, 0x6c, 0xd8, 0x75, 0x44, 0x5b, 0x2b, 0x62, 0xe4,
	0x81, 0xd0, 0xf4, 0x9d, 0x97, 0x07, 0x2d, 0xbc, 0x57, 0x0b, 0xc7, 0x82, 0x23, 0xcd, 0xef, 0x17,
	0xe7, 0x0d, 0x34, 0x18, 0xfe, 0x5c, 0x76, 0xd0, 0xaf, 0x65, 0x07, 0xfd, 0x5e, 0x76, 0x10, 0x7e,
	0x22, 0x54, 0xf5, 0x44, 0xb4, 0x51, 0x5f, 0x16, 0xd7, 0x7c, 0x2d, 0x23, 0xf4, 0x79, 0xb3, 0xec,
	0x9f, 0x6c, 0x85, 0x4b, 0x3d, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x22, 0xb7, 0xee, 0x24, 0x62,
	0x03, 0x00, 0x00,
}
