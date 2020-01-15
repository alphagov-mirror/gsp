// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads           []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor  *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter      *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	DisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                     `json:"-"`
	XXX_unrecognized        []byte                                       `json:"-"`
	XXX_sizecache           int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0, 1}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *_type.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0, 1, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_3c8760f38742c17f) }

var fileDescriptor_3c8760f38742c17f = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x6b, 0x27, 0x4d, 0xd3, 0x6d, 0x49, 0xab, 0x05, 0x1a, 0x63, 0x85, 0x36, 0x4a, 0x7b,
	0x88, 0x82, 0xe4, 0xa0, 0x54, 0xa8, 0xd0, 0x5b, 0x43, 0x1a, 0x01, 0xaa, 0x68, 0xe4, 0x8a, 0x8f,
	0x53, 0xc3, 0xc6, 0xde, 0x86, 0x15, 0xce, 0xee, 0xb2, 0xbb, 0x31, 0x58, 0xdc, 0x38, 0x71, 0xe7,
	0x2d, 0x78, 0x18, 0x2e, 0x3c, 0x00, 0x17, 0x5e, 0x82, 0x9e, 0x90, 0x3f, 0x53, 0xa7, 0x2d, 0xe2,
	0xc0, 0x6d, 0xed, 0xff, 0xcc, 0x6f, 0xc6, 0x33, 0x7f, 0x2f, 0xd8, 0xc0, 0xd4, 0x67, 0x41, 0x1b,
	0x71, 0xd2, 0xf6, 0x3b, 0x6d, 0xec, 0x4a, 0x8b, 0x0b, 0xa6, 0x18, 0x5c, 0x8d, 0xde, 0x5b, 0x88,
	0x13, 0xcb, 0xef, 0x98, 0xb5, 0x5c, 0x94, 0x4b, 0xa4, 0xc3, 0x7c, 0x2c, 0x82, 0x38, 0xd6, 0xdc,
	0xc9, 0x33, 0xa8, 0xcb, 0x19, 0xa1, 0x2a, 0x3b, 0x24, 0x51, 0x46, 0x1c, 0xa5, 0x02, 0x8e, 0xdb,
	0x1c, 0x0b, 0x07, 0x67, 0x4a, 0x6d, 0xcc, 0xd8, 0xd8, 0xc3, 0x11, 0x00, 0x51, 0xca, 0x14, 0x52,
	0x84, 0xd1, 0xa4, 0x13, 0xb3, 0xea, 0x23, 0x8f, 0xb8, 0x48, 0xe1, 0x76, 0x7a, 0x48, 0x84, 0xcd,
	0x24, 0x2d, 0x7a, 0x1a, 0x4d, 0xcf, 0xda, 0x1f, 0x04, 0xe2, 0x1c, 0x0b, 0x79, 0x9d, 0xee, 0x4e,
	0x45, 0x44, 0x8e, 0xf5, 0xc6, 0xef, 0x12, 0xb8, 0xfd, 0xd8, 0x9b, 0x4a, 0x85, 0xc5, 0x11, 0x43,
	0xee, 0x81, 0x94, 0x64, 0x4c, 0x27, 0x98, 0x2a, 0xd8, 0x02, 0xab, 0x4e, 0x2c, 0x0c, 0x29, 0x9a,
	0x60, 0x43, 0xab, 0x6b, 0xcd, 0xe5, 0xee, 0xd2, 0x79, 0xb7, 0x28, 0xf4, 0xba, 0x66, 0xaf, 0x24,
	0xe2, 0x73, 0x34, 0xc1, 0xf0, 0x09, 0x58, 0x4e, 0x3f, 0x54, 0x1a, 0x7a, 0xbd, 0xd0, 0x5c, 0xe9,
	0xb4, 0xac, 0x8b, 0xc3, 0xb3, 0xb2, 0x39, 0x1c, 0x31, 0x07, 0x79, 0x44, 0x05, 0x47, 0xa3, 0xc3,
	0x34, 0xc3, 0x9e, 0x25, 0xc3, 0x37, 0x60, 0x2d, 0xac, 0xe6, 0x0e, 0x67, 0xbc, 0xc5, 0x88, 0xb7,
	0x97, 0xe7, 0x5d, 0xd9, 0xb3, 0x15, 0x36, 0xe3, 0x66, 0xdc, 0x43, 0xaa, 0x44, 0x60, 0x57, 0x68,
	0xee, 0x25, 0xec, 0x82, 0x12, 0x67, 0x1e, 0x71, 0x02, 0xa3, 0x58, 0xd7, 0x2e, 0x37, 0x7a, 0x35,
	0x78, 0x10, 0x65, 0xd8, 0x49, 0xa6, 0x39, 0x02, 0x37, 0xaf, 0x28, 0x05, 0xd7, 0x41, 0xe1, 0x1d,
	0x0e, 0xe2, 0x49, 0xd9, 0xe1, 0x11, 0x3e, 0x00, 0x8b, 0x3e, 0xf2, 0xa6, 0xd8, 0xd0, 0xa3, 0x5a,
	0x5b, 0xd7, 0x0c, 0x25, 0xe5, 0xd8, 0x71, 0xf4, 0xbe, 0xfe, 0x50, 0x33, 0x7f, 0x16, 0x40, 0x29,
	0x2e, 0x0b, 0x4f, 0x41, 0xc5, 0x15, 0x8c, 0x0f, 0x43, 0xbf, 0x79, 0x0c, 0xb9, 0xe9, 0x8c, 0xf7,
	0xfe, 0xbd, 0x75, 0xab, 0x27, 0x18, 0x3f, 0x4e, 0xf2, 0xed, 0x1b, 0xee, 0x85, 0x27, 0x09, 0x4f,
	0x41, 0x35, 0x44, 0x73, 0xc1, 0x7c, 0x22, 0x09, 0xa3, 0x84, 0x8e, 0x87, 0x67, 0xc8, 0x51, 0x4c,
	0x18, 0x85, 0xa8, 0xef, 0x9a, 0x15, 0xdb, 0xc8, 0x4a, 0x6d, 0x64, 0xbd, 0x78, 0x4a, 0xd5, 0x6e,
	0xe7, 0x65, 0xd8, 0x6d, 0xe4, 0x89, 0x96, 0x5e, 0x5f, 0xb0, 0x37, 0xe6, 0x29, 0xfd, 0x08, 0x02,
	0x5f, 0x81, 0x5b, 0xe9, 0xa7, 0x0e, 0xa5, 0x42, 0x1e, 0x1e, 0xa2, 0x33, 0x85, 0x45, 0xb2, 0x80,
	0x3b, 0x97, 0xe0, 0xbd, 0xc4, 0xa3, 0x5d, 0x70, 0xde, 0x5d, 0xfa, 0xa6, 0x15, 0x5b, 0x7a, 0x79,
	0xc1, 0x86, 0x29, 0xe2, 0x24, 0x24, 0x1c, 0x84, 0x00, 0xf8, 0x08, 0x18, 0x2e, 0x91, 0x68, 0xe4,
	0xe1, 0xe1, 0x7c, 0x69, 0x63, 0xb1, 0xae, 0x35, 0xcb, 0x76, 0x35, 0xd1, 0x8f, 0xe7, 0x64, 0xf3,
	0x13, 0x58, 0xbd, 0x38, 0x12, 0xb8, 0x0d, 0xca, 0x0e, 0x52, 0x78, 0xcc, 0x44, 0x30, 0x6f, 0xf5,
	0x4c, 0x80, 0x7d, 0xb0, 0x16, 0x2d, 0x22, 0xf9, 0x75, 0xd1, 0x38, 0x5d, 0xec, 0xdd, 0x64, 0x13,
	0xe1, 0x8f, 0x6d, 0xf5, 0x05, 0x72, 0xc2, 0xf6, 0x91, 0x37, 0x88, 0xe3, 0xec, 0x68, 0x7d, 0x83,
	0x2c, 0xe9, 0x59, 0xb1, 0xac, 0xad, 0xeb, 0x9d, 0xef, 0x3a, 0x30, 0xd2, 0xcd, 0xf7, 0xd2, 0xeb,
	0xe4, 0x04, 0x0b, 0x9f, 0x38, 0x18, 0xbe, 0x06, 0x6b, 0x27, 0x4a, 0x60, 0x34, 0x99, 0x39, 0x77,
	0x33, 0xbf, 0xee, 0x2c, 0xc5, 0xc6, 0xef, 0xa7, 0x58, 0x2a, 0x73, 0xeb, 0x5a, 0x5d, 0x72, 0x46,
	0x25, 0x6e, 0x2c, 0x34, 0xb5, 0xfb, 0x1a, 0x44, 0xa0, 0xd2, 0xc3, 0x9e, 0x42, 0x33, 0xf0, 0xf6,
	0x5c, 0x62, 0xa8, 0x5e, 0xa2, 0xef, 0xfc, 0x3d, 0x28, 0x57, 0x62, 0x0a, 0x2a, 0x7d, 0xac, 0x9c,
	0xb7, 0xff, 0xb1, 0xf7, 0xc6, 0xe7, 0x1f, 0xbf, 0xbe, 0xea, 0xb5, 0x46, 0x35, 0x77, 0xf9, 0xee,
	0x67, 0xd7, 0xc4, 0xbe, 0xd6, 0xea, 0xde, 0x03, 0x26, 0x61, 0x31, 0x88, 0x0b, 0xf6, 0x31, 0xc8,
	0x31, 0xbb, 0xe5, 0x43, 0x57, 0x0e, 0x42, 0x8b, 0x0d, 0xb4, 0x2f, 0x9a, 0x36, 0x2a, 0x45, 0x76,
	0xdb, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xf6, 0xeb, 0xf4, 0xfd, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}
