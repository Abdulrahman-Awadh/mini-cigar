// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: analytics.proto

package analytics

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

type BringTotalSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BringTotalSalesRequest) Reset() {
	*x = BringTotalSalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BringTotalSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BringTotalSalesRequest) ProtoMessage() {}

func (x *BringTotalSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BringTotalSalesRequest.ProtoReflect.Descriptor instead.
func (*BringTotalSalesRequest) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

type BringTotalSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice float32 `protobuf:"fixed32,1,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *BringTotalSalesResponse) Reset() {
	*x = BringTotalSalesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BringTotalSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BringTotalSalesResponse) ProtoMessage() {}

func (x *BringTotalSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BringTotalSalesResponse.ProtoReflect.Descriptor instead.
func (*BringTotalSalesResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *BringTotalSalesResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type BringSalesByProductIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *BringSalesByProductIdRequest) Reset() {
	*x = BringSalesByProductIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BringSalesByProductIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BringSalesByProductIdRequest) ProtoMessage() {}

func (x *BringSalesByProductIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BringSalesByProductIdRequest.ProtoReflect.Descriptor instead.
func (*BringSalesByProductIdRequest) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *BringSalesByProductIdRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type BringSalesByProductIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice int64 `protobuf:"varint,1,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *BringSalesByProductIdResponse) Reset() {
	*x = BringSalesByProductIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BringSalesByProductIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BringSalesByProductIdResponse) ProtoMessage() {}

func (x *BringSalesByProductIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BringSalesByProductIdResponse.ProtoReflect.Descriptor instead.
func (*BringSalesByProductIdResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *BringSalesByProductIdResponse) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type ListTopFiveCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTopFiveCustomersRequest) Reset() {
	*x = ListTopFiveCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopFiveCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopFiveCustomersRequest) ProtoMessage() {}

func (x *ListTopFiveCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopFiveCustomersRequest.ProtoReflect.Descriptor instead.
func (*ListTopFiveCustomersRequest) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{4}
}

type ListTopFiveCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId []string `protobuf:"bytes,1,rep,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ListTopFiveCustomersResponse) Reset() {
	*x = ListTopFiveCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopFiveCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopFiveCustomersResponse) ProtoMessage() {}

func (x *ListTopFiveCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopFiveCustomersResponse.ProtoReflect.Descriptor instead.
func (*ListTopFiveCustomersResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{5}
}

func (x *ListTopFiveCustomersResponse) GetCustomerId() []string {
	if x != nil {
		return x.CustomerId
	}
	return nil
}

var File_analytics_proto protoreflect.FileDescriptor

var file_analytics_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x18, 0x0a, 0x16, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x17, 0x42,
	0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x1c, 0x42, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x1d, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x70, 0x46, 0x69, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x70, 0x46, 0x69, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x32, 0x85, 0x02, 0x0a, 0x10, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x0f, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x12, 0x17, 0x2e, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x42, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x42,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x42, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x69, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x69, 0x76,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x69, 0x76, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_proto_rawDescOnce sync.Once
	file_analytics_proto_rawDescData = file_analytics_proto_rawDesc
)

func file_analytics_proto_rawDescGZIP() []byte {
	file_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_proto_rawDescData)
	})
	return file_analytics_proto_rawDescData
}

var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_analytics_proto_goTypes = []interface{}{
	(*BringTotalSalesRequest)(nil),        // 0: BringTotalSalesRequest
	(*BringTotalSalesResponse)(nil),       // 1: BringTotalSalesResponse
	(*BringSalesByProductIdRequest)(nil),  // 2: BringSalesByProductIdRequest
	(*BringSalesByProductIdResponse)(nil), // 3: BringSalesByProductIdResponse
	(*ListTopFiveCustomersRequest)(nil),   // 4: ListTopFiveCustomersRequest
	(*ListTopFiveCustomersResponse)(nil),  // 5: ListTopFiveCustomersResponse
}
var file_analytics_proto_depIdxs = []int32{
	0, // 0: AnalyticsService.BringTotalSales:input_type -> BringTotalSalesRequest
	2, // 1: AnalyticsService.BringSalesByProductId:input_type -> BringSalesByProductIdRequest
	4, // 2: AnalyticsService.ListTopFiveCustomers:input_type -> ListTopFiveCustomersRequest
	1, // 3: AnalyticsService.BringTotalSales:output_type -> BringTotalSalesResponse
	3, // 4: AnalyticsService.BringSalesByProductId:output_type -> BringSalesByProductIdResponse
	5, // 5: AnalyticsService.ListTopFiveCustomers:output_type -> ListTopFiveCustomersResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_analytics_proto_init() }
func file_analytics_proto_init() {
	if File_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BringTotalSalesRequest); i {
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
		file_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BringTotalSalesResponse); i {
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
		file_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BringSalesByProductIdRequest); i {
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
		file_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BringSalesByProductIdResponse); i {
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
		file_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopFiveCustomersRequest); i {
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
		file_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopFiveCustomersResponse); i {
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
			RawDescriptor: file_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_proto_depIdxs,
		MessageInfos:      file_analytics_proto_msgTypes,
	}.Build()
	File_analytics_proto = out.File
	file_analytics_proto_rawDesc = nil
	file_analytics_proto_goTypes = nil
	file_analytics_proto_depIdxs = nil
}
