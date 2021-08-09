// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grpc/proto/address.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_INVALID_ADDRESS_FORMAT     ErrorCode = 0
	ErrorCode_UNABLE_TO_VALIDATE_ADDRESS ErrorCode = 1
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "INVALID_ADDRESS_FORMAT",
		1: "UNABLE_TO_VALIDATE_ADDRESS",
	}
	ErrorCode_value = map[string]int32{
		"INVALID_ADDRESS_FORMAT":     0,
		"UNABLE_TO_VALIDATE_ADDRESS": 1,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_proto_address_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_api_grpc_proto_address_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_proto_address_proto_rawDescGZIP(), []int{0}
}

type AddressValidationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressLines []string `protobuf:"bytes,1,rep,name=AddressLines,proto3" json:"AddressLines,omitempty"`
	Locality     string   `protobuf:"bytes,2,opt,name=Locality,proto3" json:"Locality,omitempty"`
	PostalCode   string   `protobuf:"bytes,3,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	Region       string   `protobuf:"bytes,4,opt,name=Region,proto3" json:"Region,omitempty"`
	CountryCode  string   `protobuf:"bytes,5,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
}

func (x *AddressValidationRequest) Reset() {
	*x = AddressValidationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressValidationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressValidationRequest) ProtoMessage() {}

func (x *AddressValidationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressValidationRequest.ProtoReflect.Descriptor instead.
func (*AddressValidationRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_address_proto_rawDescGZIP(), []int{0}
}

func (x *AddressValidationRequest) GetAddressLines() []string {
	if x != nil {
		return x.AddressLines
	}
	return nil
}

func (x *AddressValidationRequest) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *AddressValidationRequest) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *AddressValidationRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddressValidationRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type AddressValidationResponseSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string            `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	AddressLines []string          `protobuf:"bytes,2,rep,name=AddressLines,proto3" json:"AddressLines,omitempty"`
	Region       string            `protobuf:"bytes,3,opt,name=Region,proto3" json:"Region,omitempty"`
	Locality     string            `protobuf:"bytes,4,opt,name=Locality,proto3" json:"Locality,omitempty"`
	PostalCode   string            `protobuf:"bytes,5,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	CountryCode  string            `protobuf:"bytes,6,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
	Meta         map[string]string `protobuf:"bytes,7,rep,name=Meta,proto3" json:"Meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AddressValidationResponseSuccess) Reset() {
	*x = AddressValidationResponseSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressValidationResponseSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressValidationResponseSuccess) ProtoMessage() {}

func (x *AddressValidationResponseSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressValidationResponseSuccess.ProtoReflect.Descriptor instead.
func (*AddressValidationResponseSuccess) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_address_proto_rawDescGZIP(), []int{1}
}

func (x *AddressValidationResponseSuccess) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AddressValidationResponseSuccess) GetAddressLines() []string {
	if x != nil {
		return x.AddressLines
	}
	return nil
}

func (x *AddressValidationResponseSuccess) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddressValidationResponseSuccess) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *AddressValidationResponseSuccess) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *AddressValidationResponseSuccess) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *AddressValidationResponseSuccess) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AddressValidationResponseFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input     *AddressValidationRequest `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	ErrorCode ErrorCode                 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=ErrorCode" json:"error_code,omitempty"`
}

func (x *AddressValidationResponseFailure) Reset() {
	*x = AddressValidationResponseFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressValidationResponseFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressValidationResponseFailure) ProtoMessage() {}

func (x *AddressValidationResponseFailure) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressValidationResponseFailure.ProtoReflect.Descriptor instead.
func (*AddressValidationResponseFailure) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressValidationResponseFailure) GetInput() *AddressValidationRequest {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *AddressValidationResponseFailure) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_INVALID_ADDRESS_FORMAT
}

type AddressValidationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*AddressValidationResponse_SuccessResponse
	//	*AddressValidationResponse_ErrorResponse
	Response isAddressValidationResponse_Response `protobuf_oneof:"response"`
}

func (x *AddressValidationResponse) Reset() {
	*x = AddressValidationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressValidationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressValidationResponse) ProtoMessage() {}

func (x *AddressValidationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressValidationResponse.ProtoReflect.Descriptor instead.
func (*AddressValidationResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_address_proto_rawDescGZIP(), []int{3}
}

func (m *AddressValidationResponse) GetResponse() isAddressValidationResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AddressValidationResponse) GetSuccessResponse() *AddressValidationResponseSuccess {
	if x, ok := x.GetResponse().(*AddressValidationResponse_SuccessResponse); ok {
		return x.SuccessResponse
	}
	return nil
}

func (x *AddressValidationResponse) GetErrorResponse() *AddressValidationResponseFailure {
	if x, ok := x.GetResponse().(*AddressValidationResponse_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

type isAddressValidationResponse_Response interface {
	isAddressValidationResponse_Response()
}

type AddressValidationResponse_SuccessResponse struct {
	SuccessResponse *AddressValidationResponseSuccess `protobuf:"bytes,1,opt,name=success_response,json=successResponse,proto3,oneof"`
}

type AddressValidationResponse_ErrorResponse struct {
	ErrorResponse *AddressValidationResponseFailure `protobuf:"bytes,2,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*AddressValidationResponse_SuccessResponse) isAddressValidationResponse_Response() {}

func (*AddressValidationResponse_ErrorResponse) isAddressValidationResponse_Response() {}

var File_api_grpc_proto_address_proto protoreflect.FileDescriptor

var file_api_grpc_proto_address_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x01, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x20, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x7e, 0x0a, 0x20, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x10, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x0f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x47, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10,
	0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x32, 0x4e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x08,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_proto_address_proto_rawDescOnce sync.Once
	file_api_grpc_proto_address_proto_rawDescData = file_api_grpc_proto_address_proto_rawDesc
)

func file_api_grpc_proto_address_proto_rawDescGZIP() []byte {
	file_api_grpc_proto_address_proto_rawDescOnce.Do(func() {
		file_api_grpc_proto_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_proto_address_proto_rawDescData)
	})
	return file_api_grpc_proto_address_proto_rawDescData
}

var file_api_grpc_proto_address_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_grpc_proto_address_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_grpc_proto_address_proto_goTypes = []interface{}{
	(ErrorCode)(0),                           // 0: ErrorCode
	(*AddressValidationRequest)(nil),         // 1: AddressValidationRequest
	(*AddressValidationResponseSuccess)(nil), // 2: AddressValidationResponseSuccess
	(*AddressValidationResponseFailure)(nil), // 3: AddressValidationResponseFailure
	(*AddressValidationResponse)(nil),        // 4: AddressValidationResponse
	nil,                                      // 5: AddressValidationResponseSuccess.MetaEntry
}
var file_api_grpc_proto_address_proto_depIdxs = []int32{
	5, // 0: AddressValidationResponseSuccess.Meta:type_name -> AddressValidationResponseSuccess.MetaEntry
	1, // 1: AddressValidationResponseFailure.input:type_name -> AddressValidationRequest
	0, // 2: AddressValidationResponseFailure.error_code:type_name -> ErrorCode
	2, // 3: AddressValidationResponse.success_response:type_name -> AddressValidationResponseSuccess
	3, // 4: AddressValidationResponse.error_response:type_name -> AddressValidationResponseFailure
	1, // 5: Address.Validate:input_type -> AddressValidationRequest
	4, // 6: Address.Validate:output_type -> AddressValidationResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_grpc_proto_address_proto_init() }
func file_api_grpc_proto_address_proto_init() {
	if File_api_grpc_proto_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_proto_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressValidationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_grpc_proto_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressValidationResponseSuccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_grpc_proto_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressValidationResponseFailure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_grpc_proto_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressValidationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_grpc_proto_address_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AddressValidationResponse_SuccessResponse)(nil),
		(*AddressValidationResponse_ErrorResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpc_proto_address_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_proto_address_proto_goTypes,
		DependencyIndexes: file_api_grpc_proto_address_proto_depIdxs,
		EnumInfos:         file_api_grpc_proto_address_proto_enumTypes,
		MessageInfos:      file_api_grpc_proto_address_proto_msgTypes,
	}.Build()
	File_api_grpc_proto_address_proto = out.File
	file_api_grpc_proto_address_proto_rawDesc = nil
	file_api_grpc_proto_address_proto_goTypes = nil
	file_api_grpc_proto_address_proto_depIdxs = nil
}
