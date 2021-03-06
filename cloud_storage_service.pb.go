// Code generated by protoc-gen-go.
// source: cloud_storage_service.proto
// DO NOT EDIT!

/*
Package cloud_storage_service is a generated protocol buffer package.

It is generated from these files:
	cloud_storage_service.proto

It has these top-level messages:
	StoreRequest
	DeleteRequest
	DeleteResponse
	StorageObject
*/
package cloud_storage_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StoreRequest struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	// The raw data of the object to be stored
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StoreRequest) Reset()                    { *m = StoreRequest{} }
func (m *StoreRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()               {}
func (*StoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *StoreRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteRequest struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeleteRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type DeleteResponse struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteResponse) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type StorageObject struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *StorageObject) Reset()                    { *m = StorageObject{} }
func (m *StorageObject) String() string            { return proto.CompactTextString(m) }
func (*StorageObject) ProtoMessage()               {}
func (*StorageObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StorageObject) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *StorageObject) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "cloud_storage_service.StoreRequest")
	proto.RegisterType((*DeleteRequest)(nil), "cloud_storage_service.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "cloud_storage_service.DeleteResponse")
	proto.RegisterType((*StorageObject)(nil), "cloud_storage_service.StorageObject")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CloudStorageService service

type CloudStorageServiceClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StorageObject, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type cloudStorageServiceClient struct {
	cc *grpc.ClientConn
}

func NewCloudStorageServiceClient(cc *grpc.ClientConn) CloudStorageServiceClient {
	return &cloudStorageServiceClient{cc}
}

func (c *cloudStorageServiceClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StorageObject, error) {
	out := new(StorageObject)
	err := grpc.Invoke(ctx, "/cloud_storage_service.CloudStorageService/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudStorageServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/cloud_storage_service.CloudStorageService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudStorageService service

type CloudStorageServiceServer interface {
	Store(context.Context, *StoreRequest) (*StorageObject, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterCloudStorageServiceServer(s *grpc.Server, srv CloudStorageServiceServer) {
	s.RegisterService(&_CloudStorageService_serviceDesc, srv)
}

func _CloudStorageService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudStorageServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_storage_service.CloudStorageService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudStorageServiceServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudStorageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudStorageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_storage_service.CloudStorageService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudStorageServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudStorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud_storage_service.CloudStorageService",
	HandlerType: (*CloudStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _CloudStorageService_Store_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudStorageService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud_storage_service.proto",
}

func init() { proto.RegisterFile("cloud_storage_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0xc9, 0x2f,
	0x4d, 0x89, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc5, 0x2a, 0xa9, 0x64, 0xc7, 0xc5, 0x13,
	0x5c, 0x92, 0x5f, 0x94, 0x1a, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc5, 0xc5, 0x91,
	0x96, 0x99, 0x93, 0x9a, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7,
	0x0b, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x81,
	0xd9, 0x4a, 0xda, 0x5c, 0xbc, 0x2e, 0xa9, 0x39, 0xa9, 0x25, 0xc4, 0x18, 0xa0, 0xa4, 0xc3, 0xc5,
	0x07, 0x53, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x57, 0xb5, 0x2d, 0x17, 0x6f, 0x30, 0xc4,
	0xb5, 0xfe, 0x49, 0x59, 0xa9, 0xc9, 0xf8, 0xdd, 0x26, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x03, 0x76,
	0x1a, 0x67, 0x10, 0x88, 0x69, 0x74, 0x84, 0x91, 0x4b, 0xd8, 0x19, 0xe4, 0x67, 0xa8, 0x21, 0xc1,
	0x10, 0x1f, 0x0b, 0x85, 0x70, 0xb1, 0x82, 0x7d, 0x2c, 0xa4, 0xac, 0x87, 0x3d, 0xbc, 0x90, 0xc3,
	0x43, 0x4a, 0x05, 0x8f, 0x22, 0xb8, 0xcb, 0x94, 0x18, 0x84, 0xc2, 0xb9, 0xd8, 0x20, 0x5e, 0x13,
	0xc2, 0xa5, 0x03, 0x25, 0x98, 0xa4, 0x54, 0x09, 0xa8, 0x82, 0x84, 0x8f, 0x12, 0x83, 0x93, 0x66,
	0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4e, 0x66, 0x4e,
	0x6a, 0x66, 0xbe, 0x3e, 0x56, 0xbd, 0x49, 0x6c, 0xe0, 0x98, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xbe, 0x59, 0x9c, 0xe1, 0x08, 0x02, 0x00, 0x00,
}
