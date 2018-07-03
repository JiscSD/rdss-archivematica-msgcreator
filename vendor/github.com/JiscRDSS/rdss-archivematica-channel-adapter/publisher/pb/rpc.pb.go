// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publisher/pb/rpc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	publisher/pb/rpc.proto

It has these top-level messages:
	MetadataReadRequest
	MetadataReadResponse
	MetadataCreateRequest
	MetadataCreateResponse
	MetadataUpdateRequest
	MetadataUpdateResponse
	MetadataDeleteRequest
	MetadataDeleteResponse
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MetadataReadRequest struct {
}

func (m *MetadataReadRequest) Reset()                    { *m = MetadataReadRequest{} }
func (m *MetadataReadRequest) String() string            { return proto.CompactTextString(m) }
func (*MetadataReadRequest) ProtoMessage()               {}
func (*MetadataReadRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

type MetadataReadResponse struct {
}

func (m *MetadataReadResponse) Reset()                    { *m = MetadataReadResponse{} }
func (m *MetadataReadResponse) String() string            { return proto.CompactTextString(m) }
func (*MetadataReadResponse) ProtoMessage()               {}
func (*MetadataReadResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

type MetadataCreateRequest struct {
}

func (m *MetadataCreateRequest) Reset()                    { *m = MetadataCreateRequest{} }
func (m *MetadataCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*MetadataCreateRequest) ProtoMessage()               {}
func (*MetadataCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{2} }

type MetadataCreateResponse struct {
}

func (m *MetadataCreateResponse) Reset()                    { *m = MetadataCreateResponse{} }
func (m *MetadataCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*MetadataCreateResponse) ProtoMessage()               {}
func (*MetadataCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{3} }

type MetadataUpdateRequest struct {
}

func (m *MetadataUpdateRequest) Reset()                    { *m = MetadataUpdateRequest{} }
func (m *MetadataUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*MetadataUpdateRequest) ProtoMessage()               {}
func (*MetadataUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{4} }

type MetadataUpdateResponse struct {
}

func (m *MetadataUpdateResponse) Reset()                    { *m = MetadataUpdateResponse{} }
func (m *MetadataUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*MetadataUpdateResponse) ProtoMessage()               {}
func (*MetadataUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{5} }

type MetadataDeleteRequest struct {
}

func (m *MetadataDeleteRequest) Reset()                    { *m = MetadataDeleteRequest{} }
func (m *MetadataDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*MetadataDeleteRequest) ProtoMessage()               {}
func (*MetadataDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{6} }

type MetadataDeleteResponse struct {
}

func (m *MetadataDeleteResponse) Reset()                    { *m = MetadataDeleteResponse{} }
func (m *MetadataDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*MetadataDeleteResponse) ProtoMessage()               {}
func (*MetadataDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{7} }

func init() {
	proto.RegisterType((*MetadataReadRequest)(nil), "pb.MetadataReadRequest")
	proto.RegisterType((*MetadataReadResponse)(nil), "pb.MetadataReadResponse")
	proto.RegisterType((*MetadataCreateRequest)(nil), "pb.MetadataCreateRequest")
	proto.RegisterType((*MetadataCreateResponse)(nil), "pb.MetadataCreateResponse")
	proto.RegisterType((*MetadataUpdateRequest)(nil), "pb.MetadataUpdateRequest")
	proto.RegisterType((*MetadataUpdateResponse)(nil), "pb.MetadataUpdateResponse")
	proto.RegisterType((*MetadataDeleteRequest)(nil), "pb.MetadataDeleteRequest")
	proto.RegisterType((*MetadataDeleteResponse)(nil), "pb.MetadataDeleteResponse")
}
func (this *MetadataReadRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataReadRequest)
	if !ok {
		that2, ok := that.(MetadataReadRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataReadResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataReadResponse)
	if !ok {
		that2, ok := that.(MetadataReadResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataCreateRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataCreateRequest)
	if !ok {
		that2, ok := that.(MetadataCreateRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataCreateResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataCreateResponse)
	if !ok {
		that2, ok := that.(MetadataCreateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataUpdateRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataUpdateRequest)
	if !ok {
		that2, ok := that.(MetadataUpdateRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataUpdateResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataUpdateResponse)
	if !ok {
		that2, ok := that.(MetadataUpdateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataDeleteRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataDeleteRequest)
	if !ok {
		that2, ok := that.(MetadataDeleteRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *MetadataDeleteResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MetadataDeleteResponse)
	if !ok {
		that2, ok := that.(MetadataDeleteResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rdss service

type RdssClient interface {
	MetadataRead(ctx context.Context, in *MetadataReadRequest, opts ...grpc.CallOption) (*MetadataReadResponse, error)
	MetadataCreate(ctx context.Context, in *MetadataCreateRequest, opts ...grpc.CallOption) (*MetadataCreateResponse, error)
	MetadataUpdate(ctx context.Context, in *MetadataUpdateRequest, opts ...grpc.CallOption) (*MetadataUpdateResponse, error)
	MetadataDelete(ctx context.Context, in *MetadataDeleteRequest, opts ...grpc.CallOption) (*MetadataDeleteResponse, error)
}

type rdssClient struct {
	cc *grpc.ClientConn
}

func NewRdssClient(cc *grpc.ClientConn) RdssClient {
	return &rdssClient{cc}
}

func (c *rdssClient) MetadataRead(ctx context.Context, in *MetadataReadRequest, opts ...grpc.CallOption) (*MetadataReadResponse, error) {
	out := new(MetadataReadResponse)
	err := grpc.Invoke(ctx, "/pb.Rdss/MetadataRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdssClient) MetadataCreate(ctx context.Context, in *MetadataCreateRequest, opts ...grpc.CallOption) (*MetadataCreateResponse, error) {
	out := new(MetadataCreateResponse)
	err := grpc.Invoke(ctx, "/pb.Rdss/MetadataCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdssClient) MetadataUpdate(ctx context.Context, in *MetadataUpdateRequest, opts ...grpc.CallOption) (*MetadataUpdateResponse, error) {
	out := new(MetadataUpdateResponse)
	err := grpc.Invoke(ctx, "/pb.Rdss/MetadataUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdssClient) MetadataDelete(ctx context.Context, in *MetadataDeleteRequest, opts ...grpc.CallOption) (*MetadataDeleteResponse, error) {
	out := new(MetadataDeleteResponse)
	err := grpc.Invoke(ctx, "/pb.Rdss/MetadataDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rdss service

type RdssServer interface {
	MetadataRead(context.Context, *MetadataReadRequest) (*MetadataReadResponse, error)
	MetadataCreate(context.Context, *MetadataCreateRequest) (*MetadataCreateResponse, error)
	MetadataUpdate(context.Context, *MetadataUpdateRequest) (*MetadataUpdateResponse, error)
	MetadataDelete(context.Context, *MetadataDeleteRequest) (*MetadataDeleteResponse, error)
}

func RegisterRdssServer(s *grpc.Server, srv RdssServer) {
	s.RegisterService(&_Rdss_serviceDesc, srv)
}

func _Rdss_MetadataRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdssServer).MetadataRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Rdss/MetadataRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdssServer).MetadataRead(ctx, req.(*MetadataReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdss_MetadataCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdssServer).MetadataCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Rdss/MetadataCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdssServer).MetadataCreate(ctx, req.(*MetadataCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdss_MetadataUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdssServer).MetadataUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Rdss/MetadataUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdssServer).MetadataUpdate(ctx, req.(*MetadataUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdss_MetadataDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdssServer).MetadataDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Rdss/MetadataDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdssServer).MetadataDelete(ctx, req.(*MetadataDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rdss_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Rdss",
	HandlerType: (*RdssServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MetadataRead",
			Handler:    _Rdss_MetadataRead_Handler,
		},
		{
			MethodName: "MetadataCreate",
			Handler:    _Rdss_MetadataCreate_Handler,
		},
		{
			MethodName: "MetadataUpdate",
			Handler:    _Rdss_MetadataUpdate_Handler,
		},
		{
			MethodName: "MetadataDelete",
			Handler:    _Rdss_MetadataDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publisher/pb/rpc.proto",
}

func NewPopulatedMetadataReadRequest(r randyRpc, easy bool) *MetadataReadRequest {
	this := &MetadataReadRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataReadResponse(r randyRpc, easy bool) *MetadataReadResponse {
	this := &MetadataReadResponse{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataCreateRequest(r randyRpc, easy bool) *MetadataCreateRequest {
	this := &MetadataCreateRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataCreateResponse(r randyRpc, easy bool) *MetadataCreateResponse {
	this := &MetadataCreateResponse{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataUpdateRequest(r randyRpc, easy bool) *MetadataUpdateRequest {
	this := &MetadataUpdateRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataUpdateResponse(r randyRpc, easy bool) *MetadataUpdateResponse {
	this := &MetadataUpdateResponse{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataDeleteRequest(r randyRpc, easy bool) *MetadataDeleteRequest {
	this := &MetadataDeleteRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetadataDeleteResponse(r randyRpc, easy bool) *MetadataDeleteResponse {
	this := &MetadataDeleteResponse{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyRpc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRpc(r randyRpc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringRpc(r randyRpc) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneRpc(r)
	}
	return string(tmps)
}
func randUnrecognizedRpc(r randyRpc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldRpc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldRpc(dAtA []byte, r randyRpc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateRpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateRpc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("publisher/pb/rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0x4d, 0xca,
	0xc9, 0x2c, 0xce, 0x48, 0x2d, 0xd2, 0x2f, 0x48, 0xd2, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x25, 0x95, 0xa6, 0x81, 0x79, 0x60,
	0x0e, 0x98, 0x05, 0xd1, 0xa2, 0x24, 0xca, 0x25, 0xec, 0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58, 0x92,
	0x18, 0x94, 0x9a, 0x98, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x24, 0xc6, 0x25, 0x82,
	0x2a, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0x24, 0xce, 0x25, 0x0a, 0x13, 0x77, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x85, 0x69, 0x90, 0xe0, 0x12, 0x43, 0x97, 0xc0, 0xd4, 0x12, 0x5a, 0x90, 0x82,
	0x5d, 0x0b, 0x4c, 0x02, 0x53, 0x8b, 0x4b, 0x6a, 0x4e, 0x2a, 0x56, 0x2d, 0x30, 0x09, 0x88, 0x16,
	0xa3, 0x25, 0x4c, 0x5c, 0x2c, 0x41, 0x29, 0xc5, 0xc5, 0x42, 0x8e, 0x5c, 0x3c, 0xc8, 0x2e, 0x17,
	0x12, 0xd7, 0x2b, 0x48, 0xd2, 0xc3, 0xe2, 0x45, 0x29, 0x09, 0x4c, 0x09, 0x88, 0x59, 0x42, 0xee,
	0x5c, 0x7c, 0xa8, 0x7e, 0x11, 0x92, 0x44, 0x56, 0x8b, 0xe2, 0x71, 0x29, 0x29, 0x6c, 0x52, 0x98,
	0x06, 0x41, 0x7c, 0x88, 0x6a, 0x10, 0x4a, 0x70, 0xa0, 0x1a, 0x84, 0x1a, 0x20, 0xc8, 0x06, 0x41,
	0xfc, 0x8d, 0x6a, 0x10, 0x4a, 0x20, 0xa1, 0x1a, 0x84, 0x1a, 0x4c, 0x4e, 0x3c, 0x3f, 0x1e, 0xca,
	0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0x63, 0x12, 0x1b, 0x38, 0x0d, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0x10, 0xf5, 0x7b, 0x50, 0x02, 0x00, 0x00,
}