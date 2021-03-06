// Code generated by protoc-gen-go.
// source: github.com/micro/kv-srv/proto/store/store.proto
// DO NOT EDIT!

/*
Package go_micro_srv_kv_store is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/kv-srv/proto/store/store.proto

It has these top-level messages:
	Item
	GetRequest
	GetResponse
	PutRequest
	PutResponse
	DelRequest
	DelResponse
*/
package go_micro_srv_kv_store

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Item struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration int64  `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type PutRequest struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DelRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DelResponse struct {
}

func (m *DelResponse) Reset()                    { *m = DelResponse{} }
func (m *DelResponse) String() string            { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()               {}
func (*DelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Item)(nil), "go.micro.srv.kv.store.Item")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.kv.store.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.kv.store.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "go.micro.srv.kv.store.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "go.micro.srv.kv.store.PutResponse")
	proto.RegisterType((*DelRequest)(nil), "go.micro.srv.kv.store.DelRequest")
	proto.RegisterType((*DelResponse)(nil), "go.micro.srv.kv.store.DelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Store service

type StoreClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...client.CallOption) (*PutResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error)
}

type storeClient struct {
	c           client.Client
	serviceName string
}

func NewStoreClient(serviceName string, c client.Client) StoreClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.kv.store"
	}
	return &storeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *storeClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Store.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Put(ctx context.Context, in *PutRequest, opts ...client.CallOption) (*PutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Store.Put", in)
	out := new(PutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Del(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Store.Del", in)
	out := new(DelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	Put(context.Context, *PutRequest, *PutResponse) error
	Del(context.Context, *DelRequest, *DelResponse) error
}

func RegisterStoreHandler(s server.Server, hdlr StoreHandler) {
	s.Handle(s.NewHandler(&Store{hdlr}))
}

type Store struct {
	StoreHandler
}

func (h *Store) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.StoreHandler.Get(ctx, in, out)
}

func (h *Store) Put(ctx context.Context, in *PutRequest, out *PutResponse) error {
	return h.StoreHandler.Put(ctx, in, out)
}

func (h *Store) Del(ctx context.Context, in *DelRequest, out *DelResponse) error {
	return h.StoreHandler.Del(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xb7, 0xff, 0xee, 0xfe, 0xc1, 0xa9, 0x82, 0x04, 0x85, 0xa2, 0x20, 0x9a, 0x93, 0x17,
	0x13, 0x58, 0xcf, 0x7a, 0x12, 0xc4, 0xcb, 0x22, 0xf1, 0x13, 0xec, 0x2e, 0xc3, 0x1a, 0xba, 0xdd,
	0xd4, 0x24, 0x2d, 0xfa, 0xc9, 0xbd, 0x9a, 0x4c, 0xc5, 0x16, 0x34, 0x82, 0x97, 0xd2, 0xce, 0x7b,
	0xf9, 0xcd, 0x7b, 0x25, 0x20, 0x37, 0xda, 0x3f, 0xb7, 0x2b, 0xb1, 0x36, 0xb5, 0xac, 0xf5, 0xda,
	0x1a, 0x59, 0x75, 0x57, 0xce, 0x76, 0xb2, 0xb1, 0xc6, 0x1b, 0xe9, 0xbc, 0xb1, 0xd8, 0x3f, 0x05,
	0x4d, 0xd8, 0xf1, 0xc6, 0x08, 0x32, 0x8a, 0xe0, 0x12, 0x55, 0x27, 0x48, 0xe4, 0x0b, 0x98, 0x3e,
	0x78, 0xac, 0xd9, 0x21, 0xe4, 0x15, 0xbe, 0x95, 0xd9, 0x79, 0x76, 0xb9, 0xa7, 0xe2, 0x2b, 0x3b,
	0x82, 0x59, 0xb7, 0xdc, 0xb6, 0x58, 0xfe, 0x0b, 0xb3, 0x7d, 0xd5, 0x7f, 0xb0, 0x33, 0x00, 0x7c,
	0x6d, 0xb4, 0x5d, 0x7a, 0x6d, 0x76, 0x65, 0x1e, 0xa4, 0x5c, 0x8d, 0x26, 0x3c, 0xe8, 0xf7, 0xe8,
	0x15, 0xbe, 0xb4, 0xe8, 0xfc, 0x77, 0x2a, 0xbf, 0x85, 0x82, 0x74, 0xd7, 0x98, 0x9d, 0x43, 0x26,
	0x61, 0xaa, 0xc3, 0x7a, 0x72, 0x14, 0xf3, 0x53, 0xf1, 0x63, 0x48, 0x11, 0x13, 0x2a, 0x32, 0xf2,
	0x1b, 0x80, 0xc7, 0xf6, 0x8b, 0xff, 0xe7, 0xe3, 0x07, 0x50, 0xd0, 0xf1, 0x7e, 0x7d, 0x4c, 0x7b,
	0x87, 0xdb, 0x74, 0xda, 0x60, 0x27, 0xbd, 0xb7, 0xcf, 0xdf, 0x33, 0x98, 0x3d, 0x45, 0x24, 0x5b,
	0x40, 0x1e, 0x6a, 0xb0, 0x8b, 0xc4, 0xc6, 0xe1, 0x17, 0x9c, 0xf0, 0xdf, 0x2c, 0x9f, 0x31, 0x26,
	0x91, 0x17, 0x72, 0x25, 0x79, 0x43, 0xe5, 0x24, 0x6f, 0x5c, 0x8b, 0x78, 0x21, 0x78, 0x92, 0x37,
	0x94, 0x4e, 0xf2, 0x46, 0xbd, 0xf9, 0x64, 0xf5, 0x9f, 0x2e, 0xd1, 0xf5, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb2, 0x7e, 0x24, 0xd2, 0x77, 0x02, 0x00, 0x00,
}
