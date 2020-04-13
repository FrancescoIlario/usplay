// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activitycomm/activity.proto

package activitycomm

import (
	context "context"
	fmt "fmt"
	activitytypecomm "github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the activity details
type CreateActivityRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ActTypeID            string   `protobuf:"bytes,4,opt,name=ActTypeID,proto3" json:"ActTypeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActivityRequest) Reset()         { *m = CreateActivityRequest{} }
func (m *CreateActivityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateActivityRequest) ProtoMessage()    {}
func (*CreateActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{0}
}

func (m *CreateActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityRequest.Unmarshal(m, b)
}
func (m *CreateActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityRequest.Marshal(b, m, deterministic)
}
func (m *CreateActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityRequest.Merge(m, src)
}
func (m *CreateActivityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateActivityRequest.Size(m)
}
func (m *CreateActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityRequest proto.InternalMessageInfo

func (m *CreateActivityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateActivityRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateActivityRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateActivityRequest) GetActTypeID() string {
	if m != nil {
		return m.ActTypeID
	}
	return ""
}

// The response message containing the id of the activity
type CreateActivityReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActivityReply) Reset()         { *m = CreateActivityReply{} }
func (m *CreateActivityReply) String() string { return proto.CompactTextString(m) }
func (*CreateActivityReply) ProtoMessage()    {}
func (*CreateActivityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{1}
}

func (m *CreateActivityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityReply.Unmarshal(m, b)
}
func (m *CreateActivityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityReply.Marshal(b, m, deterministic)
}
func (m *CreateActivityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityReply.Merge(m, src)
}
func (m *CreateActivityReply) XXX_Size() int {
	return xxx_messageInfo_CreateActivityReply.Size(m)
}
func (m *CreateActivityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityReply proto.InternalMessageInfo

func (m *CreateActivityReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The request message for the read request
type ReadActivityRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadActivityRequest) Reset()         { *m = ReadActivityRequest{} }
func (m *ReadActivityRequest) String() string { return proto.CompactTextString(m) }
func (*ReadActivityRequest) ProtoMessage()    {}
func (*ReadActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{2}
}

func (m *ReadActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadActivityRequest.Unmarshal(m, b)
}
func (m *ReadActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadActivityRequest.Marshal(b, m, deterministic)
}
func (m *ReadActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadActivityRequest.Merge(m, src)
}
func (m *ReadActivityRequest) XXX_Size() int {
	return xxx_messageInfo_ReadActivityRequest.Size(m)
}
func (m *ReadActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadActivityRequest proto.InternalMessageInfo

func (m *ReadActivityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the read request
type ReadActivityReply struct {
	Activity             *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadActivityReply) Reset()         { *m = ReadActivityReply{} }
func (m *ReadActivityReply) String() string { return proto.CompactTextString(m) }
func (*ReadActivityReply) ProtoMessage()    {}
func (*ReadActivityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{3}
}

func (m *ReadActivityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadActivityReply.Unmarshal(m, b)
}
func (m *ReadActivityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadActivityReply.Marshal(b, m, deterministic)
}
func (m *ReadActivityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadActivityReply.Merge(m, src)
}
func (m *ReadActivityReply) XXX_Size() int {
	return xxx_messageInfo_ReadActivityReply.Size(m)
}
func (m *ReadActivityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadActivityReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadActivityReply proto.InternalMessageInfo

func (m *ReadActivityReply) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

// The request message for the delete request
type DeleteActivityRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteActivityRequest) Reset()         { *m = DeleteActivityRequest{} }
func (m *DeleteActivityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteActivityRequest) ProtoMessage()    {}
func (*DeleteActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{4}
}

func (m *DeleteActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteActivityRequest.Unmarshal(m, b)
}
func (m *DeleteActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteActivityRequest.Marshal(b, m, deterministic)
}
func (m *DeleteActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteActivityRequest.Merge(m, src)
}
func (m *DeleteActivityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteActivityRequest.Size(m)
}
func (m *DeleteActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteActivityRequest proto.InternalMessageInfo

func (m *DeleteActivityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the delete request
type DeleteActivityReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteActivityReply) Reset()         { *m = DeleteActivityReply{} }
func (m *DeleteActivityReply) String() string { return proto.CompactTextString(m) }
func (*DeleteActivityReply) ProtoMessage()    {}
func (*DeleteActivityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{5}
}

func (m *DeleteActivityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteActivityReply.Unmarshal(m, b)
}
func (m *DeleteActivityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteActivityReply.Marshal(b, m, deterministic)
}
func (m *DeleteActivityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteActivityReply.Merge(m, src)
}
func (m *DeleteActivityReply) XXX_Size() int {
	return xxx_messageInfo_DeleteActivityReply.Size(m)
}
func (m *DeleteActivityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteActivityReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteActivityReply proto.InternalMessageInfo

// The request message for the update request
type UpdateActivityRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	ActTypeID            string   `protobuf:"bytes,5,opt,name=ActTypeID,proto3" json:"ActTypeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActivityRequest) Reset()         { *m = UpdateActivityRequest{} }
func (m *UpdateActivityRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityRequest) ProtoMessage()    {}
func (*UpdateActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{6}
}

func (m *UpdateActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityRequest.Unmarshal(m, b)
}
func (m *UpdateActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityRequest.Marshal(b, m, deterministic)
}
func (m *UpdateActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityRequest.Merge(m, src)
}
func (m *UpdateActivityRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityRequest.Size(m)
}
func (m *UpdateActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityRequest proto.InternalMessageInfo

func (m *UpdateActivityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateActivityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateActivityRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UpdateActivityRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateActivityRequest) GetActTypeID() string {
	if m != nil {
		return m.ActTypeID
	}
	return ""
}

// The response message to the update request.
// Returns the old values for the Activity
type UpdateActivityReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActivityReply) Reset()         { *m = UpdateActivityReply{} }
func (m *UpdateActivityReply) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityReply) ProtoMessage()    {}
func (*UpdateActivityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{7}
}

func (m *UpdateActivityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityReply.Unmarshal(m, b)
}
func (m *UpdateActivityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityReply.Marshal(b, m, deterministic)
}
func (m *UpdateActivityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityReply.Merge(m, src)
}
func (m *UpdateActivityReply) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityReply.Size(m)
}
func (m *UpdateActivityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityReply proto.InternalMessageInfo

// The request message for the list activities request
type ListActivitiesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListActivitiesRequest) Reset()         { *m = ListActivitiesRequest{} }
func (m *ListActivitiesRequest) String() string { return proto.CompactTextString(m) }
func (*ListActivitiesRequest) ProtoMessage()    {}
func (*ListActivitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{8}
}

func (m *ListActivitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActivitiesRequest.Unmarshal(m, b)
}
func (m *ListActivitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActivitiesRequest.Marshal(b, m, deterministic)
}
func (m *ListActivitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActivitiesRequest.Merge(m, src)
}
func (m *ListActivitiesRequest) XXX_Size() int {
	return xxx_messageInfo_ListActivitiesRequest.Size(m)
}
func (m *ListActivitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActivitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListActivitiesRequest proto.InternalMessageInfo

// The response message to the list activities request.
type ListActivitiesReply struct {
	Activities           []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListActivitiesReply) Reset()         { *m = ListActivitiesReply{} }
func (m *ListActivitiesReply) String() string { return proto.CompactTextString(m) }
func (*ListActivitiesReply) ProtoMessage()    {}
func (*ListActivitiesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{9}
}

func (m *ListActivitiesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActivitiesReply.Unmarshal(m, b)
}
func (m *ListActivitiesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActivitiesReply.Marshal(b, m, deterministic)
}
func (m *ListActivitiesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActivitiesReply.Merge(m, src)
}
func (m *ListActivitiesReply) XXX_Size() int {
	return xxx_messageInfo_ListActivitiesReply.Size(m)
}
func (m *ListActivitiesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActivitiesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListActivitiesReply proto.InternalMessageInfo

func (m *ListActivitiesReply) GetActivities() []*Activity {
	if m != nil {
		return m.Activities
	}
	return nil
}

type Activity struct {
	Id                   string                         `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string                         `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Description          string                         `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	ActType              *activitytypecomm.ActivityType `protobuf:"bytes,5,opt,name=ActType,proto3" json:"ActType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7fe47ad1627e1f, []int{10}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Activity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Activity) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Activity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Activity) GetActType() *activitytypecomm.ActivityType {
	if m != nil {
		return m.ActType
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateActivityRequest)(nil), "activitycomm.CreateActivityRequest")
	proto.RegisterType((*CreateActivityReply)(nil), "activitycomm.CreateActivityReply")
	proto.RegisterType((*ReadActivityRequest)(nil), "activitycomm.ReadActivityRequest")
	proto.RegisterType((*ReadActivityReply)(nil), "activitycomm.ReadActivityReply")
	proto.RegisterType((*DeleteActivityRequest)(nil), "activitycomm.DeleteActivityRequest")
	proto.RegisterType((*DeleteActivityReply)(nil), "activitycomm.DeleteActivityReply")
	proto.RegisterType((*UpdateActivityRequest)(nil), "activitycomm.UpdateActivityRequest")
	proto.RegisterType((*UpdateActivityReply)(nil), "activitycomm.UpdateActivityReply")
	proto.RegisterType((*ListActivitiesRequest)(nil), "activitycomm.ListActivitiesRequest")
	proto.RegisterType((*ListActivitiesReply)(nil), "activitycomm.ListActivitiesReply")
	proto.RegisterType((*Activity)(nil), "activitycomm.Activity")
}

func init() {
	proto.RegisterFile("activitycomm/activity.proto", fileDescriptor_9a7fe47ad1627e1f)
}

var fileDescriptor_9a7fe47ad1627e1f = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x8b, 0xd4, 0x30,
	0x1c, 0x9d, 0x4e, 0xeb, 0xb8, 0xfb, 0x1b, 0x11, 0xcc, 0x50, 0x2d, 0xa3, 0xe8, 0xd8, 0x45, 0xf4,
	0xd4, 0xc2, 0x08, 0x22, 0x78, 0x90, 0x75, 0x07, 0x75, 0x60, 0x15, 0xa9, 0x7a, 0xf1, 0x96, 0x4d,
	0xc3, 0x1a, 0x6c, 0xb7, 0xb1, 0xc9, 0x0c, 0xf4, 0xe4, 0x47, 0xf0, 0x1b, 0x78, 0xf7, 0x5b, 0x4a,
	0xfa, 0xcf, 0xce, 0xcf, 0x38, 0x7b, 0xda, 0x5b, 0xf2, 0x7b, 0xaf, 0x2f, 0x2f, 0x8f, 0x97, 0xc2,
	0x5d, 0xca, 0xb4, 0xd8, 0x0a, 0x5d, 0xb1, 0x22, 0xcf, 0xe3, 0x6e, 0x13, 0xc9, 0xb2, 0xd0, 0x05,
	0xb9, 0x31, 0x04, 0xe7, 0x47, 0xdd, 0x4e, 0x57, 0x92, 0xef, 0xd0, 0xcd, 0xa0, 0xf9, 0x24, 0xfc,
	0x01, 0xfe, 0x49, 0xc9, 0xa9, 0xe6, 0xc7, 0x2d, 0x96, 0xf0, 0xef, 0x1b, 0xae, 0x34, 0x21, 0xe0,
	0xbd, 0xa7, 0x39, 0x0f, 0x9c, 0x85, 0xf3, 0xe4, 0x30, 0xa9, 0xd7, 0x66, 0x76, 0x52, 0xa4, 0x3c,
	0x18, 0x37, 0x33, 0xb3, 0x26, 0x0b, 0x98, 0xae, 0xb8, 0x62, 0xa5, 0x90, 0x5a, 0x14, 0x17, 0x81,
	0x5b, 0x43, 0xc3, 0x11, 0xb9, 0x07, 0x87, 0xc7, 0x4c, 0x7f, 0xaa, 0x24, 0x5f, 0xaf, 0x02, 0xaf,
	0xc6, 0xff, 0x0e, 0xc2, 0x47, 0x30, 0xc3, 0x06, 0x64, 0x56, 0x91, 0x9b, 0x30, 0x5e, 0xa7, 0xed,
	0xe1, 0xe3, 0x75, 0x6a, 0x68, 0x09, 0xa7, 0x29, 0x76, 0x89, 0x69, 0x6f, 0xe0, 0xd6, 0x2e, 0xcd,
	0x68, 0x2d, 0xe1, 0xa0, 0xbb, 0x79, 0x4d, 0x9d, 0x2e, 0x6f, 0x47, 0xc3, 0xa4, 0xa2, 0x9e, 0xde,
	0xf3, 0xc2, 0xc7, 0xe0, 0xaf, 0x78, 0xc6, 0xff, 0xcd, 0x05, 0x9f, 0xe8, 0xc3, 0x0c, 0x13, 0x65,
	0x56, 0x85, 0x3f, 0x1d, 0xf0, 0x3f, 0xcb, 0x94, 0x5e, 0x2a, 0xd0, 0x07, 0x3d, 0xb6, 0x04, 0xed,
	0xfe, 0x3f, 0x68, 0xef, 0x92, 0xa0, 0xaf, 0xe1, 0xa0, 0x7d, 0x98, 0x61, 0x43, 0xc6, 0xe8, 0x1d,
	0xf0, 0x4f, 0x85, 0xd2, 0xed, 0x50, 0x70, 0xd5, 0xfa, 0x0c, 0xdf, 0xc1, 0x0c, 0x03, 0x26, 0xcc,
	0x67, 0x00, 0xb4, 0x1f, 0x05, 0xce, 0xc2, 0xdd, 0x13, 0xe7, 0x80, 0x19, 0xfe, 0x72, 0xe0, 0xa0,
	0x03, 0xae, 0x30, 0x83, 0xe7, 0x70, 0xbd, 0xbd, 0x72, 0x9d, 0xc0, 0x74, 0x79, 0x3f, 0xc2, 0xcf,
	0xa0, 0xf7, 0x67, 0x58, 0x49, 0x47, 0x5f, 0xfe, 0x76, 0x61, 0xda, 0x21, 0x1f, 0xb7, 0x8c, 0x24,
	0x30, 0x69, 0x8a, 0x49, 0x8e, 0x76, 0xaf, 0x67, 0x7d, 0x2f, 0xf3, 0x87, 0xfb, 0x49, 0x26, 0xea,
	0x11, 0x39, 0x05, 0xcf, 0xd4, 0x93, 0x20, 0xb2, 0xa5, 0xd9, 0xf3, 0x07, 0xfb, 0x28, 0x8d, 0x5a,
	0x02, 0x93, 0xa6, 0x7a, 0xd8, 0xa1, 0xb5, 0xb9, 0xd8, 0xa1, 0xad, 0xb5, 0xb5, 0x66, 0xd3, 0x12,
	0xac, 0x69, 0x2d, 0x33, 0xd6, 0xb4, 0x15, 0x6c, 0x44, 0x3e, 0x80, 0x67, 0x9a, 0x84, 0x15, 0xad,
	0xb5, 0xc3, 0x8a, 0x96, 0x0a, 0x86, 0xa3, 0x57, 0x2f, 0xdf, 0xba, 0x5f, 0x5e, 0x9c, 0x0b, 0xfd,
	0x75, 0x73, 0x16, 0xb1, 0x22, 0x8f, 0x5f, 0x97, 0xf4, 0x82, 0x71, 0xc5, 0x8a, 0x75, 0x46, 0x4b,
	0x51, 0xc4, 0x1b, 0x25, 0x33, 0x5a, 0xc5, 0xf2, 0xdb, 0x79, 0xac, 0x78, 0xb9, 0x15, 0x8c, 0xab,
	0x78, 0x28, 0x7b, 0x36, 0xa9, 0xff, 0x7e, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0x09,
	0xe6, 0xe9, 0x4f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActivitySvcClient is the client API for ActivitySvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivitySvcClient interface {
	// Creates a new activity
	Create(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error)
	// Reads an activity
	Read(ctx context.Context, in *ReadActivityRequest, opts ...grpc.CallOption) (*ReadActivityReply, error)
	// Delete an activity
	Delete(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error)
	// Update an activity
	Update(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error)
	// List an activity
	List(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesReply, error)
}

type activitySvcClient struct {
	cc grpc.ClientConnInterface
}

func NewActivitySvcClient(cc grpc.ClientConnInterface) ActivitySvcClient {
	return &activitySvcClient{cc}
}

func (c *activitySvcClient) Create(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error) {
	out := new(CreateActivityReply)
	err := c.cc.Invoke(ctx, "/activitycomm.ActivitySvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activitySvcClient) Read(ctx context.Context, in *ReadActivityRequest, opts ...grpc.CallOption) (*ReadActivityReply, error) {
	out := new(ReadActivityReply)
	err := c.cc.Invoke(ctx, "/activitycomm.ActivitySvc/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activitySvcClient) Delete(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error) {
	out := new(DeleteActivityReply)
	err := c.cc.Invoke(ctx, "/activitycomm.ActivitySvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activitySvcClient) Update(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error) {
	out := new(UpdateActivityReply)
	err := c.cc.Invoke(ctx, "/activitycomm.ActivitySvc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activitySvcClient) List(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesReply, error) {
	out := new(ListActivitiesReply)
	err := c.cc.Invoke(ctx, "/activitycomm.ActivitySvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivitySvcServer is the server API for ActivitySvc service.
type ActivitySvcServer interface {
	// Creates a new activity
	Create(context.Context, *CreateActivityRequest) (*CreateActivityReply, error)
	// Reads an activity
	Read(context.Context, *ReadActivityRequest) (*ReadActivityReply, error)
	// Delete an activity
	Delete(context.Context, *DeleteActivityRequest) (*DeleteActivityReply, error)
	// Update an activity
	Update(context.Context, *UpdateActivityRequest) (*UpdateActivityReply, error)
	// List an activity
	List(context.Context, *ListActivitiesRequest) (*ListActivitiesReply, error)
}

// UnimplementedActivitySvcServer can be embedded to have forward compatible implementations.
type UnimplementedActivitySvcServer struct {
}

func (*UnimplementedActivitySvcServer) Create(ctx context.Context, req *CreateActivityRequest) (*CreateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedActivitySvcServer) Read(ctx context.Context, req *ReadActivityRequest) (*ReadActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedActivitySvcServer) Delete(ctx context.Context, req *DeleteActivityRequest) (*DeleteActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedActivitySvcServer) Update(ctx context.Context, req *UpdateActivityRequest) (*UpdateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedActivitySvcServer) List(ctx context.Context, req *ListActivitiesRequest) (*ListActivitiesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterActivitySvcServer(s *grpc.Server, srv ActivitySvcServer) {
	s.RegisterService(&_ActivitySvc_serviceDesc, srv)
}

func _ActivitySvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivitySvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomm.ActivitySvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivitySvcServer).Create(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivitySvc_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivitySvcServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomm.ActivitySvc/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivitySvcServer).Read(ctx, req.(*ReadActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivitySvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivitySvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomm.ActivitySvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivitySvcServer).Delete(ctx, req.(*DeleteActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivitySvc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivitySvcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomm.ActivitySvc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivitySvcServer).Update(ctx, req.(*UpdateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivitySvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivitySvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitycomm.ActivitySvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivitySvcServer).List(ctx, req.(*ListActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivitySvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activitycomm.ActivitySvc",
	HandlerType: (*ActivitySvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ActivitySvc_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ActivitySvc_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ActivitySvc_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ActivitySvc_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ActivitySvc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activitycomm/activity.proto",
}
