// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform/api/proto/api.proto

package go_micro_api_platform

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

type CreateServiceRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{0}
}

func (m *CreateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceRequest.Unmarshal(m, b)
}
func (m *CreateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceRequest.Merge(m, src)
}
func (m *CreateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceRequest.Size(m)
}
func (m *CreateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceRequest proto.InternalMessageInfo

func (m *CreateServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type CreateServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceResponse) Reset()         { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()    {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{1}
}

func (m *CreateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceResponse.Unmarshal(m, b)
}
func (m *CreateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceResponse.Merge(m, src)
}
func (m *CreateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceResponse.Size(m)
}
func (m *CreateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceResponse proto.InternalMessageInfo

type ReadServiceRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadServiceRequest) Reset()         { *m = ReadServiceRequest{} }
func (m *ReadServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ReadServiceRequest) ProtoMessage()    {}
func (*ReadServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{2}
}

func (m *ReadServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadServiceRequest.Unmarshal(m, b)
}
func (m *ReadServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadServiceRequest.Marshal(b, m, deterministic)
}
func (m *ReadServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadServiceRequest.Merge(m, src)
}
func (m *ReadServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ReadServiceRequest.Size(m)
}
func (m *ReadServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadServiceRequest proto.InternalMessageInfo

func (m *ReadServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type ReadServiceResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadServiceResponse) Reset()         { *m = ReadServiceResponse{} }
func (m *ReadServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ReadServiceResponse) ProtoMessage()    {}
func (*ReadServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{3}
}

func (m *ReadServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadServiceResponse.Unmarshal(m, b)
}
func (m *ReadServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadServiceResponse.Marshal(b, m, deterministic)
}
func (m *ReadServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadServiceResponse.Merge(m, src)
}
func (m *ReadServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ReadServiceResponse.Size(m)
}
func (m *ReadServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadServiceResponse proto.InternalMessageInfo

func (m *ReadServiceResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type UpdateServiceRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateServiceRequest) Reset()         { *m = UpdateServiceRequest{} }
func (m *UpdateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceRequest) ProtoMessage()    {}
func (*UpdateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{4}
}

func (m *UpdateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceRequest.Unmarshal(m, b)
}
func (m *UpdateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceRequest.Merge(m, src)
}
func (m *UpdateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceRequest.Size(m)
}
func (m *UpdateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceRequest proto.InternalMessageInfo

func (m *UpdateServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type UpdateServiceResponse struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateServiceResponse) Reset()         { *m = UpdateServiceResponse{} }
func (m *UpdateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceResponse) ProtoMessage()    {}
func (*UpdateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{5}
}

func (m *UpdateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceResponse.Unmarshal(m, b)
}
func (m *UpdateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceResponse.Marshal(b, m, deterministic)
}
func (m *UpdateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceResponse.Merge(m, src)
}
func (m *UpdateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceResponse.Size(m)
}
func (m *UpdateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceResponse proto.InternalMessageInfo

func (m *UpdateServiceResponse) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteServiceRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceRequest) Reset()         { *m = DeleteServiceRequest{} }
func (m *DeleteServiceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceRequest) ProtoMessage()    {}
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{6}
}

func (m *DeleteServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceRequest.Unmarshal(m, b)
}
func (m *DeleteServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceRequest.Merge(m, src)
}
func (m *DeleteServiceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceRequest.Size(m)
}
func (m *DeleteServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceRequest proto.InternalMessageInfo

func (m *DeleteServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceResponse) Reset()         { *m = DeleteServiceResponse{} }
func (m *DeleteServiceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceResponse) ProtoMessage()    {}
func (*DeleteServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{7}
}

func (m *DeleteServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceResponse.Unmarshal(m, b)
}
func (m *DeleteServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceResponse.Marshal(b, m, deterministic)
}
func (m *DeleteServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceResponse.Merge(m, src)
}
func (m *DeleteServiceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceResponse.Size(m)
}
func (m *DeleteServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceResponse proto.InternalMessageInfo

type ListServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServicesRequest) Reset()         { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()    {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{8}
}

func (m *ListServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesRequest.Unmarshal(m, b)
}
func (m *ListServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesRequest.Marshal(b, m, deterministic)
}
func (m *ListServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesRequest.Merge(m, src)
}
func (m *ListServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListServicesRequest.Size(m)
}
func (m *ListServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesRequest proto.InternalMessageInfo

type ListServicesResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListServicesResponse) Reset()         { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()    {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{9}
}

func (m *ListServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesResponse.Unmarshal(m, b)
}
func (m *ListServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesResponse.Marshal(b, m, deterministic)
}
func (m *ListServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesResponse.Merge(m, src)
}
func (m *ListServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListServicesResponse.Size(m)
}
func (m *ListServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesResponse proto.InternalMessageInfo

func (m *ListServicesResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Service struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Source               string            `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Type                 string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_17269eb7b0f7e9ad, []int{10}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceRequest)(nil), "go.micro.api.platform.CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "go.micro.api.platform.CreateServiceResponse")
	proto.RegisterType((*ReadServiceRequest)(nil), "go.micro.api.platform.ReadServiceRequest")
	proto.RegisterType((*ReadServiceResponse)(nil), "go.micro.api.platform.ReadServiceResponse")
	proto.RegisterType((*UpdateServiceRequest)(nil), "go.micro.api.platform.UpdateServiceRequest")
	proto.RegisterType((*UpdateServiceResponse)(nil), "go.micro.api.platform.UpdateServiceResponse")
	proto.RegisterType((*DeleteServiceRequest)(nil), "go.micro.api.platform.DeleteServiceRequest")
	proto.RegisterType((*DeleteServiceResponse)(nil), "go.micro.api.platform.DeleteServiceResponse")
	proto.RegisterType((*ListServicesRequest)(nil), "go.micro.api.platform.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "go.micro.api.platform.ListServicesResponse")
	proto.RegisterType((*Service)(nil), "go.micro.api.platform.Service")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.api.platform.Service.MetadataEntry")
}

func init() { proto.RegisterFile("platform/api/proto/api.proto", fileDescriptor_17269eb7b0f7e9ad) }

var fileDescriptor_17269eb7b0f7e9ad = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0xca, 0xd3, 0x40,
	0x10, 0xfd, 0xf2, 0xf5, 0xd7, 0xa9, 0x05, 0x99, 0x26, 0x1a, 0x82, 0x48, 0xc9, 0x95, 0xda, 0x92,
	0x42, 0xbd, 0x29, 0xf5, 0x52, 0x05, 0x2f, 0x54, 0xda, 0x88, 0x0f, 0xb0, 0xb6, 0x53, 0x09, 0x26,
	0xd9, 0xb8, 0xbb, 0x2d, 0xf4, 0x89, 0x7c, 0x37, 0x9f, 0x42, 0x92, 0x6c, 0x4a, 0x53, 0x13, 0x09,
	0x1f, 0xb9, 0x9b, 0x9d, 0x9c, 0x3d, 0xe7, 0xec, 0xec, 0x9e, 0xc0, 0xf3, 0x24, 0x64, 0xea, 0xc0,
	0x45, 0xb4, 0x60, 0x49, 0xb0, 0x48, 0x04, 0x57, 0x3c, 0xad, 0xbc, 0xac, 0x42, 0xeb, 0x07, 0xf7,
	0xa2, 0x60, 0x27, 0xb8, 0x97, 0xf5, 0x34, 0xd4, 0xdd, 0x80, 0xf9, 0x4e, 0x10, 0x53, 0xf4, 0x95,
	0xc4, 0x29, 0xd8, 0x91, 0x4f, 0xbf, 0x8e, 0x24, 0x15, 0xae, 0x60, 0x20, 0xf3, 0x8e, 0x6d, 0x4c,
	0x8d, 0x97, 0xa3, 0xe5, 0x0b, 0xaf, 0x92, 0xc0, 0x2b, 0xf6, 0x15, 0x70, 0xf7, 0x19, 0x58, 0x37,
	0x8c, 0x32, 0xe1, 0xb1, 0x24, 0xf7, 0x0b, 0xa0, 0x4f, 0x6c, 0xdf, 0x9a, 0xd0, 0x16, 0x26, 0x25,
	0xbe, 0x5c, 0x06, 0xd7, 0x30, 0xd4, 0x08, 0x69, 0x1b, 0xd3, 0x4e, 0x03, 0xc6, 0x0b, 0x3e, 0x9d,
	0xc6, 0xb7, 0x64, 0xdf, 0xe6, 0x34, 0xb6, 0x60, 0xdd, 0x30, 0x6a, 0x9b, 0x0f, 0xa7, 0xdc, 0x80,
	0xf9, 0x9e, 0x42, 0x6a, 0xf7, 0xca, 0x6e, 0x18, 0xf5, 0x95, 0x59, 0x30, 0xf9, 0x14, 0x48, 0xa5,
	0xdb, 0x52, 0x2b, 0xb9, 0x3e, 0x98, 0xe5, 0x76, 0x0b, 0xa3, 0xff, 0x63, 0xc0, 0x40, 0x77, 0x11,
	0xa1, 0x1b, 0xb3, 0x28, 0x3f, 0xc6, 0x23, 0x3f, 0xab, 0xd1, 0x86, 0xc1, 0x89, 0x84, 0x0c, 0x78,
	0x6c, 0xdf, 0x67, 0xed, 0x62, 0x89, 0x4f, 0xa1, 0x2f, 0xf9, 0x51, 0xec, 0xc8, 0xee, 0x64, 0x1f,
	0xf4, 0x2a, 0x65, 0x51, 0xe7, 0x84, 0xec, 0x6e, 0xce, 0x92, 0xd6, 0xf8, 0x11, 0x86, 0x11, 0x29,
	0xb6, 0x67, 0x8a, 0xd9, 0xbd, 0xcc, 0xe1, 0xfc, 0xff, 0x0e, 0xbd, 0xcf, 0x1a, 0xfe, 0x21, 0x56,
	0xe2, 0xec, 0x5f, 0x76, 0x3b, 0x6f, 0x61, 0x5c, 0xfa, 0x84, 0x4f, 0xa0, 0xf3, 0x93, 0xce, 0xda,
	0x73, 0x5a, 0xa2, 0x09, 0xbd, 0x13, 0x0b, 0x8f, 0xa4, 0x0d, 0xe7, 0x8b, 0xf5, 0xfd, 0xca, 0x58,
	0xfe, 0xee, 0xc2, 0x70, 0xa3, 0x95, 0x30, 0x84, 0x71, 0x29, 0x30, 0x38, 0xab, 0xb1, 0x54, 0x15,
	0x54, 0x67, 0xde, 0x0c, 0xac, 0x2f, 0xf4, 0x0e, 0x0f, 0x30, 0xba, 0x4a, 0x0d, 0xbe, 0xaa, 0xd9,
	0xfe, 0x6f, 0x52, 0x9d, 0xd7, 0x4d, 0xa0, 0x17, 0x9d, 0x10, 0xc6, 0xa5, 0x87, 0x5f, 0x7b, 0xaa,
	0xaa, 0xc0, 0xd5, 0x9e, 0xaa, 0x32, 0x4b, 0xb9, 0x5a, 0xe9, 0x05, 0xd7, 0xaa, 0x55, 0x25, 0xa7,
	0x56, 0xad, 0x3a, 0x14, 0x77, 0x18, 0xc0, 0xe3, 0xeb, 0xf7, 0x8f, 0x75, 0x93, 0xa9, 0xc8, 0x8e,
	0x33, 0x6b, 0x84, 0x2d, 0xa4, 0xbe, 0xf7, 0xb3, 0xbf, 0xf7, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x8f, 0xa9, 0x38, 0xdd, 0x05, 0x00, 0x00,
}