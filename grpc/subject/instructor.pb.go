// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: instructor.proto

package subject

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

type PaginateInstructorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int64  `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Faculty    string `protobuf:"bytes,3,opt,name=faculty,proto3" json:"faculty,omitempty"`
}

func (x *PaginateInstructorRequest) Reset() {
	*x = PaginateInstructorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginateInstructorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginateInstructorRequest) ProtoMessage() {}

func (x *PaginateInstructorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginateInstructorRequest.ProtoReflect.Descriptor instead.
func (*PaginateInstructorRequest) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{0}
}

func (x *PaginateInstructorRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PaginateInstructorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaginateInstructorRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

type PaginateInstructorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber  int64                 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PerPage     int64                 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	PageCount   int64                 `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	TotalCount  int64                 `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Instructors []*InstructorMetadata `protobuf:"bytes,5,rep,name=instructors,proto3" json:"instructors,omitempty"`
}

func (x *PaginateInstructorResponse) Reset() {
	*x = PaginateInstructorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginateInstructorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginateInstructorResponse) ProtoMessage() {}

func (x *PaginateInstructorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginateInstructorResponse.ProtoReflect.Descriptor instead.
func (*PaginateInstructorResponse) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{1}
}

func (x *PaginateInstructorResponse) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PaginateInstructorResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PaginateInstructorResponse) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *PaginateInstructorResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *PaginateInstructorResponse) GetInstructors() []*InstructorMetadata {
	if x != nil {
		return x.Instructors
	}
	return nil
}

type GetInstructorbyIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInstructorbyIdRequest) Reset() {
	*x = GetInstructorbyIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstructorbyIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstructorbyIdRequest) ProtoMessage() {}

func (x *GetInstructorbyIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstructorbyIdRequest.ProtoReflect.Descriptor instead.
func (*GetInstructorbyIdRequest) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{2}
}

func (x *GetInstructorbyIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetInstructorbyIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructor *Instructor `protobuf:"bytes,1,opt,name=instructor,proto3" json:"instructor,omitempty"`
}

func (x *GetInstructorbyIdResponse) Reset() {
	*x = GetInstructorbyIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstructorbyIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstructorbyIdResponse) ProtoMessage() {}

func (x *GetInstructorbyIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstructorbyIdResponse.ProtoReflect.Descriptor instead.
func (*GetInstructorbyIdResponse) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{3}
}

func (x *GetInstructorbyIdResponse) GetInstructor() *Instructor {
	if x != nil {
		return x.Instructor
	}
	return nil
}

type CreateInstructorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName    string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Faculty     string `protobuf:"bytes,3,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Website     string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Degree      string `protobuf:"bytes,7,opt,name=degree,proto3" json:"degree,omitempty"`
	IsAdmin     bool   `protobuf:"varint,99,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *CreateInstructorRequest) Reset() {
	*x = CreateInstructorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstructorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstructorRequest) ProtoMessage() {}

func (x *CreateInstructorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstructorRequest.ProtoReflect.Descriptor instead.
func (*CreateInstructorRequest) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{4}
}

func (x *CreateInstructorRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateInstructorRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *CreateInstructorRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateInstructorRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateInstructorRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CreateInstructorRequest) GetDegree() string {
	if x != nil {
		return x.Degree
	}
	return ""
}

func (x *CreateInstructorRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type CreateInstructorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructor *Instructor `protobuf:"bytes,1,opt,name=instructor,proto3" json:"instructor,omitempty"`
}

func (x *CreateInstructorResponse) Reset() {
	*x = CreateInstructorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstructorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstructorResponse) ProtoMessage() {}

func (x *CreateInstructorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstructorResponse.ProtoReflect.Descriptor instead.
func (*CreateInstructorResponse) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{5}
}

func (x *CreateInstructorResponse) GetInstructor() *Instructor {
	if x != nil {
		return x.Instructor
	}
	return nil
}

type UpdateInstructorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName    string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Faculty     string `protobuf:"bytes,3,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Website     string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Degree      string `protobuf:"bytes,7,opt,name=degree,proto3" json:"degree,omitempty"`
	IsAdmin     bool   `protobuf:"varint,99,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *UpdateInstructorRequest) Reset() {
	*x = UpdateInstructorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInstructorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstructorRequest) ProtoMessage() {}

func (x *UpdateInstructorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstructorRequest.ProtoReflect.Descriptor instead.
func (*UpdateInstructorRequest) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateInstructorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateInstructorRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateInstructorRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *UpdateInstructorRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateInstructorRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateInstructorRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *UpdateInstructorRequest) GetDegree() string {
	if x != nil {
		return x.Degree
	}
	return ""
}

func (x *UpdateInstructorRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type UpdateInstructorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructor *Instructor `protobuf:"bytes,1,opt,name=instructor,proto3" json:"instructor,omitempty"`
}

func (x *UpdateInstructorResponse) Reset() {
	*x = UpdateInstructorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInstructorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstructorResponse) ProtoMessage() {}

func (x *UpdateInstructorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstructorResponse.ProtoReflect.Descriptor instead.
func (*UpdateInstructorResponse) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateInstructorResponse) GetInstructor() *Instructor {
	if x != nil {
		return x.Instructor
	}
	return nil
}

type DeleteInstructorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsAdmin bool  `protobuf:"varint,99,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *DeleteInstructorRequest) Reset() {
	*x = DeleteInstructorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstructorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstructorRequest) ProtoMessage() {}

func (x *DeleteInstructorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstructorRequest.ProtoReflect.Descriptor instead.
func (*DeleteInstructorRequest) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteInstructorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteInstructorRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type DeleteInstructorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructor *Instructor `protobuf:"bytes,1,opt,name=instructor,proto3" json:"instructor,omitempty"`
}

func (x *DeleteInstructorResponse) Reset() {
	*x = DeleteInstructorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instructor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstructorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstructorResponse) ProtoMessage() {}

func (x *DeleteInstructorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instructor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstructorResponse.ProtoReflect.Descriptor instead.
func (*DeleteInstructorResponse) Descriptor() ([]byte, []int) {
	return file_instructor_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteInstructorResponse) GetInstructor() *Instructor {
	if x != nil {
		return x.Instructor
	}
	return nil
}

var File_instructor_proto protoreflect.FileDescriptor

var file_instructor_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x0c, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x19, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x22, 0xd7, 0x01, 0x0a, 0x1a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3d, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22,
	0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x62, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xd6, 0x01,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x63, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x4f, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xe6, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x63, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x4f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x22, 0x44, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x63, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x4f, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x32, 0xd9, 0x03, 0x0a, 0x11, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d,
	0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x79,
	0x49, 0x64, 0x12, 0x21, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x2e,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x20, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instructor_proto_rawDescOnce sync.Once
	file_instructor_proto_rawDescData = file_instructor_proto_rawDesc
)

func file_instructor_proto_rawDescGZIP() []byte {
	file_instructor_proto_rawDescOnce.Do(func() {
		file_instructor_proto_rawDescData = protoimpl.X.CompressGZIP(file_instructor_proto_rawDescData)
	})
	return file_instructor_proto_rawDescData
}

var file_instructor_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_instructor_proto_goTypes = []interface{}{
	(*PaginateInstructorRequest)(nil),  // 0: subject.PaginateInstructorRequest
	(*PaginateInstructorResponse)(nil), // 1: subject.PaginateInstructorResponse
	(*GetInstructorbyIdRequest)(nil),   // 2: subject.GetInstructorbyIdRequest
	(*GetInstructorbyIdResponse)(nil),  // 3: subject.GetInstructorbyIdResponse
	(*CreateInstructorRequest)(nil),    // 4: subject.CreateInstructorRequest
	(*CreateInstructorResponse)(nil),   // 5: subject.CreateInstructorResponse
	(*UpdateInstructorRequest)(nil),    // 6: subject.UpdateInstructorRequest
	(*UpdateInstructorResponse)(nil),   // 7: subject.UpdateInstructorResponse
	(*DeleteInstructorRequest)(nil),    // 8: subject.DeleteInstructorRequest
	(*DeleteInstructorResponse)(nil),   // 9: subject.DeleteInstructorResponse
	(*InstructorMetadata)(nil),         // 10: subject.InstructorMetadata
	(*Instructor)(nil),                 // 11: subject.Instructor
}
var file_instructor_proto_depIdxs = []int32{
	10, // 0: subject.PaginateInstructorResponse.instructors:type_name -> subject.InstructorMetadata
	11, // 1: subject.GetInstructorbyIdResponse.instructor:type_name -> subject.Instructor
	11, // 2: subject.CreateInstructorResponse.instructor:type_name -> subject.Instructor
	11, // 3: subject.UpdateInstructorResponse.instructor:type_name -> subject.Instructor
	11, // 4: subject.DeleteInstructorResponse.instructor:type_name -> subject.Instructor
	0,  // 5: subject.InstructorService.PaginateInstructor:input_type -> subject.PaginateInstructorRequest
	2,  // 6: subject.InstructorService.GetInstructorbyId:input_type -> subject.GetInstructorbyIdRequest
	4,  // 7: subject.InstructorService.CreateInstructor:input_type -> subject.CreateInstructorRequest
	6,  // 8: subject.InstructorService.UpdateInstructor:input_type -> subject.UpdateInstructorRequest
	8,  // 9: subject.InstructorService.DeleteInstructor:input_type -> subject.DeleteInstructorRequest
	1,  // 10: subject.InstructorService.PaginateInstructor:output_type -> subject.PaginateInstructorResponse
	3,  // 11: subject.InstructorService.GetInstructorbyId:output_type -> subject.GetInstructorbyIdResponse
	5,  // 12: subject.InstructorService.CreateInstructor:output_type -> subject.CreateInstructorResponse
	7,  // 13: subject.InstructorService.UpdateInstructor:output_type -> subject.UpdateInstructorResponse
	9,  // 14: subject.InstructorService.DeleteInstructor:output_type -> subject.DeleteInstructorResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_instructor_proto_init() }
func file_instructor_proto_init() {
	if File_instructor_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_instructor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginateInstructorRequest); i {
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
		file_instructor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginateInstructorResponse); i {
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
		file_instructor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstructorbyIdRequest); i {
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
		file_instructor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstructorbyIdResponse); i {
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
		file_instructor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstructorRequest); i {
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
		file_instructor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstructorResponse); i {
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
		file_instructor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInstructorRequest); i {
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
		file_instructor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInstructorResponse); i {
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
		file_instructor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstructorRequest); i {
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
		file_instructor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstructorResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_instructor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_instructor_proto_goTypes,
		DependencyIndexes: file_instructor_proto_depIdxs,
		MessageInfos:      file_instructor_proto_msgTypes,
	}.Build()
	File_instructor_proto = out.File
	file_instructor_proto_rawDesc = nil
	file_instructor_proto_goTypes = nil
	file_instructor_proto_depIdxs = nil
}
