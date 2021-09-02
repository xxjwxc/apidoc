// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.11.4
// source: shares/shares.proto

package shares

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	common "rpc/common"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SharesInfo 股票信息
type SharesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`             // 股票代码
	SimpleCode string  `protobuf:"bytes,2,opt,name=simpleCode,proto3" json:"simpleCode"` // 股票代码简写
	Ext        string  `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext"`               // 后缀
	Name       string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`             // 股票名字
	Price      float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price"`         // 当前价格
	Percent    float64 `protobuf:"fixed64,6,opt,name=percent,proto3" json:"percent"`     // 百分比
	Color      string  `protobuf:"bytes,7,opt,name=color,proto3" json:"color"`           // 颜色
}

func (x *SharesInfo) Reset() {
	*x = SharesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_shares_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharesInfo) ProtoMessage() {}

func (x *SharesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shares_shares_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharesInfo.ProtoReflect.Descriptor instead.
func (*SharesInfo) Descriptor() ([]byte, []int) {
	return file_shares_shares_proto_rawDescGZIP(), []int{0}
}

func (x *SharesInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SharesInfo) GetSimpleCode() string {
	if x != nil {
		return x.SimpleCode
	}
	return ""
}

func (x *SharesInfo) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

func (x *SharesInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SharesInfo) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SharesInfo) GetPercent() float64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *SharesInfo) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name"` // 分组名
	List []*SharesInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list"` // 列表
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_shares_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_shares_shares_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_shares_shares_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetList() []*SharesInfo {
	if x != nil {
		return x.List
	}
	return nil
}

// GetGroupResp 请求结构
type GetGroupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Group `protobuf:"bytes,1,rep,name=list,proto3" json:"list"` // 列表
}

func (x *GetGroupResp) Reset() {
	*x = GetGroupResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_shares_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResp) ProtoMessage() {}

func (x *GetGroupResp) ProtoReflect() protoreflect.Message {
	mi := &file_shares_shares_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResp.ProtoReflect.Descriptor instead.
func (*GetGroupResp) Descriptor() ([]byte, []int) {
	return file_shares_shares_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupResp) GetList() []*Group {
	if x != nil {
		return x.List
	}
	return nil
}

// SearchReq 搜索请求
type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"` // 股票代码
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_shares_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_shares_shares_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_shares_shares_proto_rawDescGZIP(), []int{3}
}

func (x *SearchReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// SearchReq 搜索请求
type SearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SharesInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info"` // 返回信息
}

func (x *SearchResp) Reset() {
	*x = SearchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shares_shares_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResp) ProtoMessage() {}

func (x *SearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_shares_shares_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResp.ProtoReflect.Descriptor instead.
func (*SearchResp) Descriptor() ([]byte, []int) {
	return file_shares_shares_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResp) GetInfo() *SharesInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_shares_shares_proto protoreflect.FileDescriptor

var file_shares_shares_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x1a, 0x1a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x31, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x1f, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x6e, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shares_shares_proto_rawDescOnce sync.Once
	file_shares_shares_proto_rawDescData = file_shares_shares_proto_rawDesc
)

func file_shares_shares_proto_rawDescGZIP() []byte {
	file_shares_shares_proto_rawDescOnce.Do(func() {
		file_shares_shares_proto_rawDescData = protoimpl.X.CompressGZIP(file_shares_shares_proto_rawDescData)
	})
	return file_shares_shares_proto_rawDescData
}

var file_shares_shares_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shares_shares_proto_goTypes = []interface{}{
	(*SharesInfo)(nil),   // 0: shares.SharesInfo
	(*Group)(nil),        // 1: shares.Group
	(*GetGroupResp)(nil), // 2: shares.GetGroupResp
	(*SearchReq)(nil),    // 3: shares.SearchReq
	(*SearchResp)(nil),   // 4: shares.SearchResp
	(*common.Empty)(nil), // 5: common.Empty
}
var file_shares_shares_proto_depIdxs = []int32{
	0, // 0: shares.Group.list:type_name -> shares.SharesInfo
	1, // 1: shares.GetGroupResp.list:type_name -> shares.Group
	0, // 2: shares.SearchResp.info:type_name -> shares.SharesInfo
	5, // 3: shares.shares.GetGroup:input_type -> common.Empty
	3, // 4: shares.shares.Search:input_type -> shares.SearchReq
	2, // 5: shares.shares.GetGroup:output_type -> shares.GetGroupResp
	4, // 6: shares.shares.Search:output_type -> shares.SearchResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shares_shares_proto_init() }
func file_shares_shares_proto_init() {
	if File_shares_shares_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shares_shares_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharesInfo); i {
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
		file_shares_shares_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_shares_shares_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResp); i {
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
		file_shares_shares_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_shares_shares_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResp); i {
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
			RawDescriptor: file_shares_shares_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shares_shares_proto_goTypes,
		DependencyIndexes: file_shares_shares_proto_depIdxs,
		MessageInfos:      file_shares_shares_proto_msgTypes,
	}.Build()
	File_shares_shares_proto = out.File
	file_shares_shares_proto_rawDesc = nil
	file_shares_shares_proto_goTypes = nil
	file_shares_shares_proto_depIdxs = nil
}
