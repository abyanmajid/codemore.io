// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: course.proto

package course

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

// The Course message representing a course object
type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   string    `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Likes       int32     `protobuf:"varint,5,opt,name=likes,proto3" json:"likes,omitempty"`
	Students    []string  `protobuf:"bytes,6,rep,name=students,proto3" json:"students,omitempty"`
	Topics      []string  `protobuf:"bytes,7,rep,name=topics,proto3" json:"topics,omitempty"`
	Modules     []*Module `protobuf:"bytes,8,rep,name=modules,proto3" json:"modules,omitempty"`
	UpdatedAt   string    `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt   string    `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Course) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Course) GetStudents() []string {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *Course) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Course) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

func (x *Course) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Course) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// The Module message representing a module within a course
type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tasks []*Task `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{1}
}

func (x *Module) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Module) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Module) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// The Task message representing a task within a module
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Xp   int32  `protobuf:"varint,4,opt,name=xp,proto3" json:"xp,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Task) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Task) GetXp() int32 {
	if x != nil {
		return x.Xp
	}
	return 0
}

// Request message for creating a new course
type CreateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   string `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *CreateCourseRequest) Reset() {
	*x = CreateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseRequest) ProtoMessage() {}

func (x *CreateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseRequest.ProtoReflect.Descriptor instead.
func (*CreateCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCourseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCourseRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

// Response message for creating a new course
type CreateCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateCourseResponse) Reset() {
	*x = CreateCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseResponse) ProtoMessage() {}

func (x *CreateCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseResponse.ProtoReflect.Descriptor instead.
func (*CreateCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCourseResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request message for retrieving a course by id
type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{5}
}

func (x *GetCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Response message for retrieving a course by id
type GetCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *GetCourseResponse) Reset() {
	*x = GetCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseResponse) ProtoMessage() {}

func (x *GetCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseResponse.ProtoReflect.Descriptor instead.
func (*GetCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{6}
}

func (x *GetCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

// Request message for updating an existing course
type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   string    `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Likes       int32     `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
	Students    []string  `protobuf:"bytes,5,rep,name=students,proto3" json:"students,omitempty"`
	Topics      []string  `protobuf:"bytes,6,rep,name=topics,proto3" json:"topics,omitempty"`
	Modules     []*Module `protobuf:"bytes,7,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateCourseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCourseRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *UpdateCourseRequest) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *UpdateCourseRequest) GetStudents() []string {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *UpdateCourseRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *UpdateCourseRequest) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

// Response message for updating an existing course
type UpdateCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateCourseResponse) Reset() {
	*x = UpdateCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseResponse) ProtoMessage() {}

func (x *UpdateCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseResponse.ProtoReflect.Descriptor instead.
func (*UpdateCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCourseResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request message for deleting a course by id
type DeleteCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Response message for deleting a course by id
type DeleteCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCourseResponse) Reset() {
	*x = DeleteCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseResponse) ProtoMessage() {}

func (x *DeleteCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseResponse.ProtoReflect.Descriptor instead.
func (*DeleteCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCourseResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request message for retrieving all courses
type GetAllCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCoursesRequest) Reset() {
	*x = GetAllCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCoursesRequest) ProtoMessage() {}

func (x *GetAllCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCoursesRequest.ProtoReflect.Descriptor instead.
func (*GetAllCoursesRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{11}
}

// Response message for retrieving all courses
type GetAllCoursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *GetAllCoursesResponse) Reset() {
	*x = GetAllCoursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCoursesResponse) ProtoMessage() {}

func (x *GetAllCoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCoursesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCoursesResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{12}
}

func (x *GetAllCoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

var File_course_proto protoreflect.FileDescriptor

var file_course_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x52, 0x0a, 0x06, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x4e,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x78, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x70, 0x22, 0x6c,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3b,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x26,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x32, 0x80, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x18,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_proto_rawDescOnce sync.Once
	file_course_proto_rawDescData = file_course_proto_rawDesc
)

func file_course_proto_rawDescGZIP() []byte {
	file_course_proto_rawDescOnce.Do(func() {
		file_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_proto_rawDescData)
	})
	return file_course_proto_rawDescData
}

var file_course_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_course_proto_goTypes = []any{
	(*Course)(nil),                // 0: course.Course
	(*Module)(nil),                // 1: course.Module
	(*Task)(nil),                  // 2: course.Task
	(*CreateCourseRequest)(nil),   // 3: course.CreateCourseRequest
	(*CreateCourseResponse)(nil),  // 4: course.CreateCourseResponse
	(*GetCourseRequest)(nil),      // 5: course.GetCourseRequest
	(*GetCourseResponse)(nil),     // 6: course.GetCourseResponse
	(*UpdateCourseRequest)(nil),   // 7: course.UpdateCourseRequest
	(*UpdateCourseResponse)(nil),  // 8: course.UpdateCourseResponse
	(*DeleteCourseRequest)(nil),   // 9: course.DeleteCourseRequest
	(*DeleteCourseResponse)(nil),  // 10: course.DeleteCourseResponse
	(*GetAllCoursesRequest)(nil),  // 11: course.GetAllCoursesRequest
	(*GetAllCoursesResponse)(nil), // 12: course.GetAllCoursesResponse
}
var file_course_proto_depIdxs = []int32{
	1,  // 0: course.Course.modules:type_name -> course.Module
	2,  // 1: course.Module.tasks:type_name -> course.Task
	0,  // 2: course.GetCourseResponse.course:type_name -> course.Course
	1,  // 3: course.UpdateCourseRequest.modules:type_name -> course.Module
	0,  // 4: course.GetAllCoursesResponse.courses:type_name -> course.Course
	3,  // 5: course.CourseService.CreateCourse:input_type -> course.CreateCourseRequest
	11, // 6: course.CourseService.GetAllCourses:input_type -> course.GetAllCoursesRequest
	5,  // 7: course.CourseService.GetCourse:input_type -> course.GetCourseRequest
	7,  // 8: course.CourseService.UpdateCourse:input_type -> course.UpdateCourseRequest
	9,  // 9: course.CourseService.DeleteCourse:input_type -> course.DeleteCourseRequest
	4,  // 10: course.CourseService.CreateCourse:output_type -> course.CreateCourseResponse
	12, // 11: course.CourseService.GetAllCourses:output_type -> course.GetAllCoursesResponse
	6,  // 12: course.CourseService.GetCourse:output_type -> course.GetCourseResponse
	8,  // 13: course.CourseService.UpdateCourse:output_type -> course.UpdateCourseResponse
	10, // 14: course.CourseService.DeleteCourse:output_type -> course.DeleteCourseResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_course_proto_init() }
func file_course_proto_init() {
	if File_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Course); i {
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
		file_course_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Module); i {
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
		file_course_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Task); i {
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
		file_course_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCourseRequest); i {
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
		file_course_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCourseResponse); i {
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
		file_course_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCourseRequest); i {
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
		file_course_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCourseResponse); i {
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
		file_course_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCourseRequest); i {
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
		file_course_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCourseResponse); i {
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
		file_course_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCourseRequest); i {
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
		file_course_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCourseResponse); i {
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
		file_course_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCoursesRequest); i {
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
		file_course_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCoursesResponse); i {
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
			RawDescriptor: file_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_proto_goTypes,
		DependencyIndexes: file_course_proto_depIdxs,
		MessageInfos:      file_course_proto_msgTypes,
	}.Build()
	File_course_proto = out.File
	file_course_proto_rawDesc = nil
	file_course_proto_goTypes = nil
	file_course_proto_depIdxs = nil
}
