// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lolibrary/lolibrary/service.category/proto/category.proto

package categoryproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lolibrary/lolibrary/cmd/protoc-gen-router/proto"
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

type Category struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// timestamps
	CreatedAt            string   `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Category) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type GETReadCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadCategoryRequest) Reset()         { *m = GETReadCategoryRequest{} }
func (m *GETReadCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GETReadCategoryRequest) ProtoMessage()    {}
func (*GETReadCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{1}
}

func (m *GETReadCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadCategoryRequest.Unmarshal(m, b)
}
func (m *GETReadCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GETReadCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadCategoryRequest.Merge(m, src)
}
func (m *GETReadCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GETReadCategoryRequest.Size(m)
}
func (m *GETReadCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadCategoryRequest proto.InternalMessageInfo

func (m *GETReadCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GETReadCategoryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GETReadCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GETReadCategoryResponse) Reset()         { *m = GETReadCategoryResponse{} }
func (m *GETReadCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*GETReadCategoryResponse) ProtoMessage()    {}
func (*GETReadCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{2}
}

func (m *GETReadCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadCategoryResponse.Unmarshal(m, b)
}
func (m *GETReadCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadCategoryResponse.Marshal(b, m, deterministic)
}
func (m *GETReadCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadCategoryResponse.Merge(m, src)
}
func (m *GETReadCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_GETReadCategoryResponse.Size(m)
}
func (m *GETReadCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadCategoryResponse proto.InternalMessageInfo

func (m *GETReadCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type GETListCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETListCategoriesRequest) Reset()         { *m = GETListCategoriesRequest{} }
func (m *GETListCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GETListCategoriesRequest) ProtoMessage()    {}
func (*GETListCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{3}
}

func (m *GETListCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListCategoriesRequest.Unmarshal(m, b)
}
func (m *GETListCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *GETListCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListCategoriesRequest.Merge(m, src)
}
func (m *GETListCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GETListCategoriesRequest.Size(m)
}
func (m *GETListCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETListCategoriesRequest proto.InternalMessageInfo

type GETListCategoriesResponse struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GETListCategoriesResponse) Reset()         { *m = GETListCategoriesResponse{} }
func (m *GETListCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GETListCategoriesResponse) ProtoMessage()    {}
func (*GETListCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{4}
}

func (m *GETListCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListCategoriesResponse.Unmarshal(m, b)
}
func (m *GETListCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GETListCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListCategoriesResponse.Merge(m, src)
}
func (m *GETListCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GETListCategoriesResponse.Size(m)
}
func (m *GETListCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETListCategoriesResponse proto.InternalMessageInfo

func (m *GETListCategoriesResponse) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type POSTCreateCategoryRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateCategoryRequest) Reset()         { *m = POSTCreateCategoryRequest{} }
func (m *POSTCreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*POSTCreateCategoryRequest) ProtoMessage()    {}
func (*POSTCreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{5}
}

func (m *POSTCreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateCategoryRequest.Unmarshal(m, b)
}
func (m *POSTCreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *POSTCreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateCategoryRequest.Merge(m, src)
}
func (m *POSTCreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_POSTCreateCategoryRequest.Size(m)
}
func (m *POSTCreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateCategoryRequest proto.InternalMessageInfo

func (m *POSTCreateCategoryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *POSTCreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type POSTCreateCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *POSTCreateCategoryResponse) Reset()         { *m = POSTCreateCategoryResponse{} }
func (m *POSTCreateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*POSTCreateCategoryResponse) ProtoMessage()    {}
func (*POSTCreateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{6}
}

func (m *POSTCreateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateCategoryResponse.Unmarshal(m, b)
}
func (m *POSTCreateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *POSTCreateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateCategoryResponse.Merge(m, src)
}
func (m *POSTCreateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_POSTCreateCategoryResponse.Size(m)
}
func (m *POSTCreateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateCategoryResponse proto.InternalMessageInfo

func (m *POSTCreateCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type PUTUpdateCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PUTUpdateCategoryRequest) Reset()         { *m = PUTUpdateCategoryRequest{} }
func (m *PUTUpdateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateCategoryRequest) ProtoMessage()    {}
func (*PUTUpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{7}
}

func (m *PUTUpdateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Unmarshal(m, b)
}
func (m *PUTUpdateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *PUTUpdateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateCategoryRequest.Merge(m, src)
}
func (m *PUTUpdateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Size(m)
}
func (m *PUTUpdateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateCategoryRequest proto.InternalMessageInfo

func (m *PUTUpdateCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PUTUpdateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PUTUpdateCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PUTUpdateCategoryResponse) Reset()         { *m = PUTUpdateCategoryResponse{} }
func (m *PUTUpdateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateCategoryResponse) ProtoMessage()    {}
func (*PUTUpdateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{8}
}

func (m *PUTUpdateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Unmarshal(m, b)
}
func (m *PUTUpdateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *PUTUpdateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateCategoryResponse.Merge(m, src)
}
func (m *PUTUpdateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Size(m)
}
func (m *PUTUpdateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateCategoryResponse proto.InternalMessageInfo

func (m *PUTUpdateCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type DELETERemoveCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveCategoryRequest) Reset()         { *m = DELETERemoveCategoryRequest{} }
func (m *DELETERemoveCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveCategoryRequest) ProtoMessage()    {}
func (*DELETERemoveCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{9}
}

func (m *DELETERemoveCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Unmarshal(m, b)
}
func (m *DELETERemoveCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Marshal(b, m, deterministic)
}
func (m *DELETERemoveCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveCategoryRequest.Merge(m, src)
}
func (m *DELETERemoveCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Size(m)
}
func (m *DELETERemoveCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveCategoryRequest proto.InternalMessageInfo

func (m *DELETERemoveCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DELETERemoveCategoryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveCategoryResponse) Reset()         { *m = DELETERemoveCategoryResponse{} }
func (m *DELETERemoveCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveCategoryResponse) ProtoMessage()    {}
func (*DELETERemoveCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af727c152a2e3192, []int{10}
}

func (m *DELETERemoveCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Unmarshal(m, b)
}
func (m *DELETERemoveCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Marshal(b, m, deterministic)
}
func (m *DELETERemoveCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveCategoryResponse.Merge(m, src)
}
func (m *DELETERemoveCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Size(m)
}
func (m *DELETERemoveCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveCategoryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Category)(nil), "categoryproto.Category")
	proto.RegisterType((*GETReadCategoryRequest)(nil), "categoryproto.GETReadCategoryRequest")
	proto.RegisterType((*GETReadCategoryResponse)(nil), "categoryproto.GETReadCategoryResponse")
	proto.RegisterType((*GETListCategoriesRequest)(nil), "categoryproto.GETListCategoriesRequest")
	proto.RegisterType((*GETListCategoriesResponse)(nil), "categoryproto.GETListCategoriesResponse")
	proto.RegisterType((*POSTCreateCategoryRequest)(nil), "categoryproto.POSTCreateCategoryRequest")
	proto.RegisterType((*POSTCreateCategoryResponse)(nil), "categoryproto.POSTCreateCategoryResponse")
	proto.RegisterType((*PUTUpdateCategoryRequest)(nil), "categoryproto.PUTUpdateCategoryRequest")
	proto.RegisterType((*PUTUpdateCategoryResponse)(nil), "categoryproto.PUTUpdateCategoryResponse")
	proto.RegisterType((*DELETERemoveCategoryRequest)(nil), "categoryproto.DELETERemoveCategoryRequest")
	proto.RegisterType((*DELETERemoveCategoryResponse)(nil), "categoryproto.DELETERemoveCategoryResponse")
}

func init() {
	proto.RegisterFile("github.com/lolibrary/lolibrary/service.category/proto/category.proto", fileDescriptor_af727c152a2e3192)
}

var fileDescriptor_af727c152a2e3192 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xd3, 0x40, 0x9b, 0x69, 0x6b, 0xa1, 0x3d, 0xd0, 0x8d, 0x21, 0xa8, 0x5a, 0x09, 0x28,
	0xa0, 0xd8, 0x52, 0x7b, 0xe0, 0x82, 0x90, 0xaa, 0xc4, 0xca, 0xa5, 0x82, 0x10, 0xdc, 0x33, 0x72,
	0xbc, 0xa3, 0x60, 0xe4, 0xc4, 0x61, 0xbd, 0xae, 0xe8, 0x0d, 0xe5, 0x61, 0xfa, 0x06, 0x7d, 0x33,
	0x1e, 0x00, 0x65, 0xd7, 0x4e, 0x1d, 0xc7, 0x69, 0x22, 0xe5, 0xb6, 0xf3, 0xb3, 0xdf, 0xf7, 0xf9,
	0x9b, 0x59, 0x43, 0x77, 0x14, 0xca, 0x9f, 0xe9, 0xd0, 0x0e, 0xe2, 0xb1, 0x13, 0xc5, 0x51, 0x38,
	0x14, 0xbe, 0xb8, 0x2d, 0x9c, 0x12, 0x14, 0x37, 0x61, 0x80, 0x76, 0xe0, 0x4b, 0x1c, 0xc5, 0xe2,
	0xd6, 0x99, 0x8a, 0x58, 0xc6, 0x4e, 0x1e, 0xda, 0x2a, 0x24, 0xc7, 0x79, 0xac, 0x42, 0xab, 0xb7,
	0x01, 0x34, 0x18, 0x73, 0x8d, 0x13, 0xb4, 0x47, 0x38, 0x69, 0x8b, 0x38, 0x95, 0x28, 0x32, 0x64,
	0x1d, 0x68, 0x5c, 0xf6, 0xd7, 0x80, 0x83, 0x4e, 0x06, 0x4d, 0x4c, 0xa8, 0x85, 0x9c, 0x1a, 0xa7,
	0xc6, 0x59, 0x63, 0x50, 0x0b, 0x39, 0x21, 0x50, 0x4f, 0xa2, 0x74, 0x44, 0x6b, 0x2a, 0xa3, 0xce,
	0xf3, 0xdc, 0xc4, 0x1f, 0x23, 0xad, 0xeb, 0xdc, 0xfc, 0x4c, 0x5a, 0x00, 0x81, 0x40, 0x5f, 0x22,
	0xff, 0xe1, 0x4b, 0xca, 0x55, 0xa5, 0x91, 0x65, 0x2e, 0xe5, 0xbc, 0x9c, 0x4e, 0x79, 0x5e, 0x46,
	0x5d, 0xce, 0x32, 0x97, 0x92, 0x7d, 0x82, 0xe7, 0x3d, 0xd7, 0x1b, 0xa0, 0xcf, 0x73, 0x21, 0x03,
	0xfc, 0x9d, 0x62, 0x22, 0xb7, 0xd1, 0xc3, 0xbe, 0xc0, 0xc9, 0xca, 0xed, 0x64, 0x1a, 0x4f, 0x12,
	0x24, 0x17, 0x70, 0x90, 0xbb, 0xa6, 0x40, 0x0e, 0xcf, 0x4f, 0xec, 0x25, 0x1b, 0xed, 0xc5, 0x95,
	0x45, 0x23, 0xb3, 0x80, 0xf6, 0x5c, 0xef, 0x2a, 0x4c, 0x64, 0x56, 0x0c, 0x31, 0xc9, 0xf4, 0x30,
	0x0f, 0x9a, 0x15, 0xb5, 0x8c, 0xed, 0x23, 0x40, 0xb0, 0xc8, 0x52, 0xe3, 0x74, 0xef, 0x31, 0xbe,
	0x42, 0x2b, 0xeb, 0x40, 0xb3, 0xff, 0xf5, 0xbb, 0xd7, 0x51, 0x7e, 0x95, 0x2d, 0xc8, 0x3f, 0xd9,
	0xa8, 0x18, 0xc1, 0xde, 0xc3, 0x08, 0xd8, 0x37, 0xb0, 0xaa, 0x40, 0x76, 0x71, 0xe2, 0x33, 0xd0,
	0xfe, 0xb5, 0x77, 0xad, 0xe6, 0xb4, 0xc5, 0x64, 0x94, 0xa4, 0x5a, 0x41, 0x52, 0x1f, 0x9a, 0x15,
	0xf7, 0x77, 0x51, 0xd4, 0x86, 0x17, 0x5d, 0xf7, 0xca, 0xf5, 0xdc, 0x01, 0x8e, 0xe3, 0x9b, 0x4d,
	0xa2, 0xd8, 0x2b, 0x78, 0x59, 0xdd, 0xae, 0x35, 0x9c, 0xff, 0xab, 0x3f, 0x88, 0x20, 0xbf, 0xe0,
	0xa8, 0xb8, 0x44, 0xe4, 0x75, 0x49, 0x4e, 0xf5, 0x8a, 0x5a, 0x6f, 0x36, 0xb5, 0x69, 0x2e, 0x76,
	0x38, 0xbb, 0xa7, 0xfb, 0xf0, 0xc4, 0x11, 0xe8, 0x73, 0x22, 0xc0, 0x5c, 0xb6, 0x85, 0xbc, 0x2d,
	0xc1, 0xac, 0x33, 0xde, 0x3a, 0xdb, 0xdc, 0x98, 0x31, 0x1e, 0xcf, 0xee, 0x69, 0x03, 0xf6, 0x1d,
	0xfd, 0xd0, 0xc8, 0x1f, 0x30, 0xbb, 0x18, 0x61, 0x81, 0xf3, 0x7d, 0x09, 0xea, 0x11, 0x6b, 0xad,
	0x0f, 0x5b, 0xf5, 0x96, 0x98, 0xb9, 0xe2, 0x23, 0x12, 0xcc, 0xe5, 0xb5, 0x24, 0x2b, 0x1f, 0xb1,
	0x6e, 0xfd, 0xad, 0x77, 0x5b, 0x74, 0x96, 0x58, 0xf5, 0x7f, 0x87, 0x4c, 0xc1, 0x5c, 0x7e, 0xa8,
	0x2b, 0x1e, 0xaf, 0x7b, 0xe6, 0x2b, 0x1e, 0xaf, 0x7d, 0xf3, 0x8b, 0xa9, 0x46, 0xe1, 0x7c, 0x15,
	0x66, 0x77, 0x2d, 0x02, 0xcf, 0xca, 0xff, 0xf3, 0xd9, 0x5d, 0xeb, 0x88, 0x40, 0xd2, 0xce, 0xe3,
	0xe1, 0x53, 0x85, 0x7a, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xde, 0x4b, 0x52, 0x0a, 0x19, 0x06,
	0x00, 0x00,
}
