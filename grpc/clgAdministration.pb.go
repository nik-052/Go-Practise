// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: clgAdministration.proto

package grpc

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

type CourseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	CourseName string `protobuf:"bytes,2,opt,name=CourseName,proto3" json:"CourseName"`
}

func (x *CourseInfo) Reset() {
	*x = CourseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clgAdministration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseInfo) ProtoMessage() {}

func (x *CourseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_clgAdministration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseInfo.ProtoReflect.Descriptor instead.
func (*CourseInfo) Descriptor() ([]byte, []int) {
	return file_clgAdministration_proto_rawDescGZIP(), []int{0}
}

func (x *CourseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CourseInfo) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clgAdministration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_clgAdministration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_clgAdministration_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type StudentMarks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	StudentId  string `protobuf:"bytes,2,opt,name=StudentId,proto3" json:"StudentId"`
	CourseId   string `protobuf:"bytes,3,opt,name=CourseId,proto3" json:"CourseId"`
	CourseName string `protobuf:"bytes,4,opt,name=CourseName,proto3" json:"CourseName"`
	Marks      int64  `protobuf:"varint,5,opt,name=Marks,proto3" json:"Marks"`
	Grade      string `protobuf:"bytes,6,opt,name=Grade,proto3" json:"Grade"`
}

func (x *StudentMarks) Reset() {
	*x = StudentMarks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clgAdministration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentMarks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentMarks) ProtoMessage() {}

func (x *StudentMarks) ProtoReflect() protoreflect.Message {
	mi := &file_clgAdministration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentMarks.ProtoReflect.Descriptor instead.
func (*StudentMarks) Descriptor() ([]byte, []int) {
	return file_clgAdministration_proto_rawDescGZIP(), []int{2}
}

func (x *StudentMarks) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StudentMarks) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *StudentMarks) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *StudentMarks) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *StudentMarks) GetMarks() int64 {
	if x != nil {
		return x.Marks
	}
	return 0
}

func (x *StudentMarks) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

type StudentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string        `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	Name            string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name"`
	RollNumber      string        `protobuf:"bytes,3,opt,name=RollNumber,proto3" json:"RollNumber"`
	Age             int64         `protobuf:"varint,4,opt,name=Age,proto3" json:"Age"`
	CourseId        string        `protobuf:"bytes,5,opt,name=CourseId,proto3" json:"CourseId"`
	MarksId         string        `protobuf:"bytes,6,opt,name=MarksId,proto3" json:"MarksId"`
	ClassesEnrolled *CourseInfo   `protobuf:"bytes,7,opt,name=ClassesEnrolled,proto3" json:"ClassesEnrolled"`
	StudentMarks    *StudentMarks `protobuf:"bytes,8,opt,name=StudentMarks,proto3" json:"StudentMarks"`
}

func (x *StudentInfo) Reset() {
	*x = StudentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clgAdministration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentInfo) ProtoMessage() {}

func (x *StudentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_clgAdministration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentInfo.ProtoReflect.Descriptor instead.
func (*StudentInfo) Descriptor() ([]byte, []int) {
	return file_clgAdministration_proto_rawDescGZIP(), []int{3}
}

func (x *StudentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StudentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentInfo) GetRollNumber() string {
	if x != nil {
		return x.RollNumber
	}
	return ""
}

func (x *StudentInfo) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *StudentInfo) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *StudentInfo) GetMarksId() string {
	if x != nil {
		return x.MarksId
	}
	return ""
}

func (x *StudentInfo) GetClassesEnrolled() *CourseInfo {
	if x != nil {
		return x.ClassesEnrolled
	}
	return nil
}

func (x *StudentInfo) GetStudentMarks() *StudentMarks {
	if x != nil {
		return x.StudentMarks
	}
	return nil
}

type InstructorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	InstructorCode  string      `protobuf:"bytes,2,opt,name=InstructorCode,proto3" json:"InstructorCode"`
	InstructorName  string      `protobuf:"bytes,3,opt,name=InstructorName,proto3" json:"InstructorName"`
	Department      string      `protobuf:"bytes,4,opt,name=Department,proto3" json:"Department"`
	CourseId        string      `protobuf:"bytes,5,opt,name=CourseId,proto3" json:"CourseId"`
	CourseName      string      `protobuf:"bytes,6,opt,name=CourseName,proto3" json:"CourseName"`
	ClassesEnrolled *CourseInfo `protobuf:"bytes,7,opt,name=ClassesEnrolled,proto3" json:"ClassesEnrolled"`
}

func (x *InstructorDetails) Reset() {
	*x = InstructorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clgAdministration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstructorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstructorDetails) ProtoMessage() {}

func (x *InstructorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_clgAdministration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstructorDetails.ProtoReflect.Descriptor instead.
func (*InstructorDetails) Descriptor() ([]byte, []int) {
	return file_clgAdministration_proto_rawDescGZIP(), []int{4}
}

func (x *InstructorDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstructorDetails) GetInstructorCode() string {
	if x != nil {
		return x.InstructorCode
	}
	return ""
}

func (x *InstructorDetails) GetInstructorName() string {
	if x != nil {
		return x.InstructorName
	}
	return ""
}

func (x *InstructorDetails) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *InstructorDetails) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *InstructorDetails) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *InstructorDetails) GetClassesEnrolled() *CourseInfo {
	if x != nil {
		return x.ClassesEnrolled
	}
	return nil
}

var File_clgAdministration_proto protoreflect.FileDescriptor

var file_clgAdministration_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6c, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x47, 0x6f, 0x50, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x22, 0x99, 0x02,
	0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x6f, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x41, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0f, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x47, 0x6f, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x47, 0x6f, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x0c, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x11, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x47, 0x6f, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x32, 0x4b, 0x0a,
	0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x16, 0x2e, 0x47, 0x6f, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x47, 0x6f, 0x50, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x47, 0x6f,
	0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clgAdministration_proto_rawDescOnce sync.Once
	file_clgAdministration_proto_rawDescData = file_clgAdministration_proto_rawDesc
)

func file_clgAdministration_proto_rawDescGZIP() []byte {
	file_clgAdministration_proto_rawDescOnce.Do(func() {
		file_clgAdministration_proto_rawDescData = protoimpl.X.CompressGZIP(file_clgAdministration_proto_rawDescData)
	})
	return file_clgAdministration_proto_rawDescData
}

var file_clgAdministration_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_clgAdministration_proto_goTypes = []interface{}{
	(*CourseInfo)(nil),        // 0: GoPractise.CourseInfo
	(*Res)(nil),               // 1: GoPractise.Res
	(*StudentMarks)(nil),      // 2: GoPractise.StudentMarks
	(*StudentInfo)(nil),       // 3: GoPractise.StudentInfo
	(*InstructorDetails)(nil), // 4: GoPractise.InstructorDetails
}
var file_clgAdministration_proto_depIdxs = []int32{
	0, // 0: GoPractise.StudentInfo.ClassesEnrolled:type_name -> GoPractise.CourseInfo
	2, // 1: GoPractise.StudentInfo.StudentMarks:type_name -> GoPractise.StudentMarks
	0, // 2: GoPractise.InstructorDetails.ClassesEnrolled:type_name -> GoPractise.CourseInfo
	0, // 3: GoPractise.Administration.CreateCourse:input_type -> GoPractise.CourseInfo
	1, // 4: GoPractise.Administration.CreateCourse:output_type -> GoPractise.Res
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_clgAdministration_proto_init() }
func file_clgAdministration_proto_init() {
	if File_clgAdministration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clgAdministration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseInfo); i {
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
		file_clgAdministration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
		file_clgAdministration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentMarks); i {
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
		file_clgAdministration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentInfo); i {
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
		file_clgAdministration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstructorDetails); i {
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
			RawDescriptor: file_clgAdministration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clgAdministration_proto_goTypes,
		DependencyIndexes: file_clgAdministration_proto_depIdxs,
		MessageInfos:      file_clgAdministration_proto_msgTypes,
	}.Build()
	File_clgAdministration_proto = out.File
	file_clgAdministration_proto_rawDesc = nil
	file_clgAdministration_proto_goTypes = nil
	file_clgAdministration_proto_depIdxs = nil
}