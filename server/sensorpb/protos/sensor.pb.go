// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/sensor.proto

package sensorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SensorRequest) Reset() {
	*x = SensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorRequest) ProtoMessage() {}

func (x *SensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorRequest.ProtoReflect.Descriptor instead.
func (*SensorRequest) Descriptor() ([]byte, []int) {
	return file_protos_sensor_proto_rawDescGZIP(), []int{0}
}

type SensorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SensorResponse) Reset() {
	*x = SensorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorResponse) ProtoMessage() {}

func (x *SensorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorResponse.ProtoReflect.Descriptor instead.
func (*SensorResponse) Descriptor() ([]byte, []int) {
	return file_protos_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *SensorResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ThresholdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ThresholdRequest) Reset() {
	*x = ThresholdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdRequest) ProtoMessage() {}

func (x *ThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdRequest.ProtoReflect.Descriptor instead.
func (*ThresholdRequest) Descriptor() ([]byte, []int) {
	return file_protos_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *ThresholdRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ThresholdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ThresholdResponse) Reset() {
	*x = ThresholdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdResponse) ProtoMessage() {}

func (x *ThresholdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdResponse.ProtoReflect.Descriptor instead.
func (*ThresholdResponse) Descriptor() ([]byte, []int) {
	return file_protos_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *ThresholdResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protos_sensor_proto protoreflect.FileDescriptor

var file_protos_sensor_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x22, 0x0f,
	0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x26, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x29, 0x0a, 0x11, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xaa, 0x02, 0x0a,
	0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0e, 0x48, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x45, 0x0a, 0x0e, 0x54, 0x6f, 0x78, 0x69, 0x63, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x54,
	0x6f, 0x78, 0x69, 0x63, 0x69, 0x74, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_sensor_proto_rawDescOnce sync.Once
	file_protos_sensor_proto_rawDescData = file_protos_sensor_proto_rawDesc
)

func file_protos_sensor_proto_rawDescGZIP() []byte {
	file_protos_sensor_proto_rawDescOnce.Do(func() {
		file_protos_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_sensor_proto_rawDescData)
	})
	return file_protos_sensor_proto_rawDescData
}

var file_protos_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_sensor_proto_goTypes = []interface{}{
	(*SensorRequest)(nil),     // 0: sensors.SensorRequest
	(*SensorResponse)(nil),    // 1: sensors.SensorResponse
	(*ThresholdRequest)(nil),  // 2: sensors.ThresholdRequest
	(*ThresholdResponse)(nil), // 3: sensors.ThresholdResponse
}
var file_protos_sensor_proto_depIdxs = []int32{
	0, // 0: sensors.Sensor.TempSensor:input_type -> sensors.SensorRequest
	0, // 1: sensors.Sensor.HumiditySensor:input_type -> sensors.SensorRequest
	0, // 2: sensors.Sensor.ToxicitySensor:input_type -> sensors.SensorRequest
	2, // 3: sensors.Sensor.SetToxicityThreshold:input_type -> sensors.ThresholdRequest
	1, // 4: sensors.Sensor.TempSensor:output_type -> sensors.SensorResponse
	1, // 5: sensors.Sensor.HumiditySensor:output_type -> sensors.SensorResponse
	1, // 6: sensors.Sensor.ToxicitySensor:output_type -> sensors.SensorResponse
	3, // 7: sensors.Sensor.SetToxicityThreshold:output_type -> sensors.ThresholdResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_sensor_proto_init() }
func file_protos_sensor_proto_init() {
	if File_protos_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorRequest); i {
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
		file_protos_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorResponse); i {
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
		file_protos_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdRequest); i {
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
		file_protos_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdResponse); i {
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
			RawDescriptor: file_protos_sensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_sensor_proto_goTypes,
		DependencyIndexes: file_protos_sensor_proto_depIdxs,
		MessageInfos:      file_protos_sensor_proto_msgTypes,
	}.Build()
	File_protos_sensor_proto = out.File
	file_protos_sensor_proto_rawDesc = nil
	file_protos_sensor_proto_goTypes = nil
	file_protos_sensor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SensorClient is the client API for Sensor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SensorClient interface {
	TempSensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_TempSensorClient, error)
	HumiditySensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_HumiditySensorClient, error)
	ToxicitySensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_ToxicitySensorClient, error)
	SetToxicityThreshold(ctx context.Context, in *ThresholdRequest, opts ...grpc.CallOption) (*ThresholdResponse, error)
}

type sensorClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorClient(cc grpc.ClientConnInterface) SensorClient {
	return &sensorClient{cc}
}

func (c *sensorClient) TempSensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_TempSensorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sensor_serviceDesc.Streams[0], "/sensors.Sensor/TempSensor", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorTempSensorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sensor_TempSensorClient interface {
	Recv() (*SensorResponse, error)
	grpc.ClientStream
}

type sensorTempSensorClient struct {
	grpc.ClientStream
}

func (x *sensorTempSensorClient) Recv() (*SensorResponse, error) {
	m := new(SensorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sensorClient) HumiditySensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_HumiditySensorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sensor_serviceDesc.Streams[1], "/sensors.Sensor/HumiditySensor", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorHumiditySensorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sensor_HumiditySensorClient interface {
	Recv() (*SensorResponse, error)
	grpc.ClientStream
}

type sensorHumiditySensorClient struct {
	grpc.ClientStream
}

func (x *sensorHumiditySensorClient) Recv() (*SensorResponse, error) {
	m := new(SensorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sensorClient) ToxicitySensor(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (Sensor_ToxicitySensorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sensor_serviceDesc.Streams[2], "/sensors.Sensor/ToxicitySensor", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorToxicitySensorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sensor_ToxicitySensorClient interface {
	Recv() (*SensorResponse, error)
	grpc.ClientStream
}

type sensorToxicitySensorClient struct {
	grpc.ClientStream
}

func (x *sensorToxicitySensorClient) Recv() (*SensorResponse, error) {
	m := new(SensorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sensorClient) SetToxicityThreshold(ctx context.Context, in *ThresholdRequest, opts ...grpc.CallOption) (*ThresholdResponse, error) {
	out := new(ThresholdResponse)
	err := c.cc.Invoke(ctx, "/sensors.Sensor/SetToxicityThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorServer is the server API for Sensor service.
type SensorServer interface {
	TempSensor(*SensorRequest, Sensor_TempSensorServer) error
	HumiditySensor(*SensorRequest, Sensor_HumiditySensorServer) error
	ToxicitySensor(*SensorRequest, Sensor_ToxicitySensorServer) error
	SetToxicityThreshold(context.Context, *ThresholdRequest) (*ThresholdResponse, error)
}

// UnimplementedSensorServer can be embedded to have forward compatible implementations.
type UnimplementedSensorServer struct {
}

func (*UnimplementedSensorServer) TempSensor(*SensorRequest, Sensor_TempSensorServer) error {
	return status.Errorf(codes.Unimplemented, "method TempSensor not implemented")
}
func (*UnimplementedSensorServer) HumiditySensor(*SensorRequest, Sensor_HumiditySensorServer) error {
	return status.Errorf(codes.Unimplemented, "method HumiditySensor not implemented")
}
func (*UnimplementedSensorServer) ToxicitySensor(*SensorRequest, Sensor_ToxicitySensorServer) error {
	return status.Errorf(codes.Unimplemented, "method ToxicitySensor not implemented")
}
func (*UnimplementedSensorServer) SetToxicityThreshold(context.Context, *ThresholdRequest) (*ThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToxicityThreshold not implemented")
}

func RegisterSensorServer(s *grpc.Server, srv SensorServer) {
	s.RegisterService(&_Sensor_serviceDesc, srv)
}

func _Sensor_TempSensor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SensorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorServer).TempSensor(m, &sensorTempSensorServer{stream})
}

type Sensor_TempSensorServer interface {
	Send(*SensorResponse) error
	grpc.ServerStream
}

type sensorTempSensorServer struct {
	grpc.ServerStream
}

func (x *sensorTempSensorServer) Send(m *SensorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Sensor_HumiditySensor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SensorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorServer).HumiditySensor(m, &sensorHumiditySensorServer{stream})
}

type Sensor_HumiditySensorServer interface {
	Send(*SensorResponse) error
	grpc.ServerStream
}

type sensorHumiditySensorServer struct {
	grpc.ServerStream
}

func (x *sensorHumiditySensorServer) Send(m *SensorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Sensor_ToxicitySensor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SensorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorServer).ToxicitySensor(m, &sensorToxicitySensorServer{stream})
}

type Sensor_ToxicitySensorServer interface {
	Send(*SensorResponse) error
	grpc.ServerStream
}

type sensorToxicitySensorServer struct {
	grpc.ServerStream
}

func (x *sensorToxicitySensorServer) Send(m *SensorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Sensor_SetToxicityThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServer).SetToxicityThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensors.Sensor/SetToxicityThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServer).SetToxicityThreshold(ctx, req.(*ThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sensor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensors.Sensor",
	HandlerType: (*SensorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetToxicityThreshold",
			Handler:    _Sensor_SetToxicityThreshold_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TempSensor",
			Handler:       _Sensor_TempSensor_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HumiditySensor",
			Handler:       _Sensor_HumiditySensor_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ToxicitySensor",
			Handler:       _Sensor_ToxicitySensor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/sensor.proto",
}
