// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: monitoring_backup.proto

package monitoring_backup

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

type CreatedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kebun     string `protobuf:"bytes,1,opt,name=kebun,proto3" json:"kebun,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreatedNotify) Reset() {
	*x = CreatedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedNotify) ProtoMessage() {}

func (x *CreatedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedNotify.ProtoReflect.Descriptor instead.
func (*CreatedNotify) Descriptor() ([]byte, []int) {
	return file_monitoring_backup_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedNotify) GetKebun() string {
	if x != nil {
		return x.Kebun
	}
	return ""
}

func (x *CreatedNotify) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CreatedNotify) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ResponseNotif struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseNotif) Reset() {
	*x = ResponseNotif{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_backup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseNotif) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseNotif) ProtoMessage() {}

func (x *ResponseNotif) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_backup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseNotif.ProtoReflect.Descriptor instead.
func (*ResponseNotif) Descriptor() ([]byte, []int) {
	return file_monitoring_backup_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseNotif) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ResponseNotif) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MonitoringLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kebun     string `protobuf:"bytes,1,opt,name=kebun,proto3" json:"kebun,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MonitoringLog) Reset() {
	*x = MonitoringLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_backup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitoringLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitoringLog) ProtoMessage() {}

func (x *MonitoringLog) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_backup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitoringLog.ProtoReflect.Descriptor instead.
func (*MonitoringLog) Descriptor() ([]byte, []int) {
	return file_monitoring_backup_proto_rawDescGZIP(), []int{2}
}

func (x *MonitoringLog) GetKebun() string {
	if x != nil {
		return x.Kebun
	}
	return ""
}

func (x *MonitoringLog) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *MonitoringLog) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MonitoringLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonitoringLogs []*MonitoringLog `protobuf:"bytes,1,rep,name=monitoringLogs,proto3" json:"monitoringLogs,omitempty"`
}

func (x *MonitoringLogs) Reset() {
	*x = MonitoringLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_backup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitoringLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitoringLogs) ProtoMessage() {}

func (x *MonitoringLogs) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_backup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitoringLogs.ProtoReflect.Descriptor instead.
func (*MonitoringLogs) Descriptor() ([]byte, []int) {
	return file_monitoring_backup_proto_rawDescGZIP(), []int{3}
}

func (x *MonitoringLogs) GetMonitoringLogs() []*MonitoringLog {
	if x != nil {
		return x.MonitoringLogs
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_backup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_backup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_monitoring_backup_proto_rawDescGZIP(), []int{4}
}

var File_monitoring_backup_proto protoreflect.FileDescriptor

var file_monitoring_backup_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x5b, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6b, 0x65, 0x62, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65,
	0x62, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5b, 0x0a, 0x0d,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x6b, 0x65, 0x62, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65,
	0x62, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x48, 0x0a, 0x0e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xb9,
	0x01, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x12, 0x51, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x1a, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_monitoring_backup_proto_rawDescOnce sync.Once
	file_monitoring_backup_proto_rawDescData = file_monitoring_backup_proto_rawDesc
)

func file_monitoring_backup_proto_rawDescGZIP() []byte {
	file_monitoring_backup_proto_rawDescOnce.Do(func() {
		file_monitoring_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitoring_backup_proto_rawDescData)
	})
	return file_monitoring_backup_proto_rawDescData
}

var file_monitoring_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_monitoring_backup_proto_goTypes = []interface{}{
	(*CreatedNotify)(nil),  // 0: monitoring_backup.CreatedNotify
	(*ResponseNotif)(nil),  // 1: monitoring_backup.ResponseNotif
	(*MonitoringLog)(nil),  // 2: monitoring_backup.MonitoringLog
	(*MonitoringLogs)(nil), // 3: monitoring_backup.MonitoringLogs
	(*Empty)(nil),          // 4: monitoring_backup.Empty
}
var file_monitoring_backup_proto_depIdxs = []int32{
	2, // 0: monitoring_backup.MonitoringLogs.monitoringLogs:type_name -> monitoring_backup.MonitoringLog
	0, // 1: monitoring_backup.MonitoringBackup.SendNotif:input_type -> monitoring_backup.CreatedNotify
	4, // 2: monitoring_backup.MonitoringBackup.GetMonitoringLogs:input_type -> monitoring_backup.Empty
	1, // 3: monitoring_backup.MonitoringBackup.SendNotif:output_type -> monitoring_backup.ResponseNotif
	3, // 4: monitoring_backup.MonitoringBackup.GetMonitoringLogs:output_type -> monitoring_backup.MonitoringLogs
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_monitoring_backup_proto_init() }
func file_monitoring_backup_proto_init() {
	if File_monitoring_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitoring_backup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedNotify); i {
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
		file_monitoring_backup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseNotif); i {
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
		file_monitoring_backup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitoringLog); i {
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
		file_monitoring_backup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitoringLogs); i {
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
		file_monitoring_backup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_monitoring_backup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitoring_backup_proto_goTypes,
		DependencyIndexes: file_monitoring_backup_proto_depIdxs,
		MessageInfos:      file_monitoring_backup_proto_msgTypes,
	}.Build()
	File_monitoring_backup_proto = out.File
	file_monitoring_backup_proto_rawDesc = nil
	file_monitoring_backup_proto_goTypes = nil
	file_monitoring_backup_proto_depIdxs = nil
}
