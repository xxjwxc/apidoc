// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: ccmy/weixin.proto

package ccmy

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

// OauthReq 请求结构
type OauthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jscode       string `protobuf:"bytes,1,opt,name=jscode,proto3" json:"jscode"`              // 微信端jscode
	IsUpdateUser bool   `protobuf:"varint,2,opt,name=isUpdateUser,proto3" json:"isUpdateUser"` // 是否更新用户
}

func (x *OauthReq) Reset() {
	*x = OauthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthReq) ProtoMessage() {}

func (x *OauthReq) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthReq.ProtoReflect.Descriptor instead.
func (*OauthReq) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{0}
}

func (x *OauthReq) GetJscode() string {
	if x != nil {
		return x.Jscode
	}
	return ""
}

func (x *OauthReq) GetIsUpdateUser() bool {
	if x != nil {
		return x.IsUpdateUser
	}
	return false
}

// OauthResp 微信授权返回信息
type OauthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId"`      // 用户sessionid 用于当前交互使用
	OpenId      string `protobuf:"bytes,2,opt,name=openId,proto3" json:"openId"`            // openid
	OverdueTime int64  `protobuf:"varint,3,opt,name=overdueTime,proto3" json:"overdueTime"` // 逾期时间点(时间戳)
}

func (x *OauthResp) Reset() {
	*x = OauthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthResp) ProtoMessage() {}

func (x *OauthResp) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthResp.ProtoReflect.Descriptor instead.
func (*OauthResp) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{1}
}

func (x *OauthResp) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *OauthResp) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *OauthResp) GetOverdueTime() int64 {
	if x != nil {
		return x.OverdueTime
	}
	return 0
}

// WxUserinfo 用户信息
type WxUserinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID"` // 用户sessionid 用于当前交互使用
	NickName  string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName"`   //用户昵称
	AvatarURL string `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL"` //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender"`       //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      string `protobuf:"bytes,5,opt,name=city,proto3" json:"city"`           //用户所在城市
	Province  string `protobuf:"bytes,6,opt,name=province,proto3" json:"province"`   //用户所在省份
	Country   string `protobuf:"bytes,7,opt,name=country,proto3" json:"country"`     //用户所在国家
	Language  string `protobuf:"bytes,8,opt,name=language,proto3" json:"language"`   //用户的语言，简体中文为zh_CN
	Vip       bool   `protobuf:"varint,9,opt,name=vip,proto3" json:"vip"`            // 是否vip
}

func (x *WxUserinfo) Reset() {
	*x = WxUserinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxUserinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxUserinfo) ProtoMessage() {}

func (x *WxUserinfo) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxUserinfo.ProtoReflect.Descriptor instead.
func (*WxUserinfo) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{2}
}

func (x *WxUserinfo) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *WxUserinfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *WxUserinfo) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *WxUserinfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *WxUserinfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *WxUserinfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *WxUserinfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *WxUserinfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WxUserinfo) GetVip() bool {
	if x != nil {
		return x.Vip
	}
	return false
}

// WxUserinfoResp 用户信息返回
type WxUserinfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B bool `protobuf:"varint,1,opt,name=b,proto3" json:"b"`
}

func (x *WxUserinfoResp) Reset() {
	*x = WxUserinfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxUserinfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxUserinfoResp) ProtoMessage() {}

func (x *WxUserinfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxUserinfoResp.ProtoReflect.Descriptor instead.
func (*WxUserinfoResp) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{3}
}

func (x *WxUserinfoResp) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

// GetQrcodeResp 获取二维码
type GetQrcodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`   // 小程序页面头部
	Scene string `protobuf:"bytes,2,opt,name=scene,proto3" json:"scene"` // 附带参数
}

func (x *GetQrcodeReq) Reset() {
	*x = GetQrcodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQrcodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQrcodeReq) ProtoMessage() {}

func (x *GetQrcodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQrcodeReq.ProtoReflect.Descriptor instead.
func (*GetQrcodeReq) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{4}
}

func (x *GetQrcodeReq) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetQrcodeReq) GetScene() string {
	if x != nil {
		return x.Scene
	}
	return ""
}

// GetQrcodeResp 获取二维码
type GetQrcodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url"` // 二维码图片地址
}

func (x *GetQrcodeResp) Reset() {
	*x = GetQrcodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQrcodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQrcodeResp) ProtoMessage() {}

func (x *GetQrcodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQrcodeResp.ProtoReflect.Descriptor instead.
func (*GetQrcodeResp) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{5}
}

func (x *GetQrcodeResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`            // 用户openid
	IsChecked bool   `protobuf:"varint,2,opt,name=isChecked,proto3" json:"isChecked"` // 是否点击
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{6}
}

func (x *GroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupInfo) GetIsChecked() bool {
	if x != nil {
		return x.IsChecked
	}
	return false
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid    string       `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid"`       // 用户openid
	NickName  string       `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName"`   // 用户昵称
	AvatarURL string       `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL"` // 用户头像地址
	Group     []*GroupInfo `protobuf:"bytes,4,rep,name=group,proto3" json:"group"`         // 门店列表g
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfo) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *UserInfo) GetGroup() []*GroupInfo {
	if x != nil {
		return x.Group
	}
	return nil
}

// GetUserInfoResp 用户信息
type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid    string      `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid"`       // 用户openid
	NickName  string      `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName"`   // 用户昵称
	AvatarURL string      `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL"` // 用户头像地址
	Gender    string      `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender"`       // 用户的性别
	City      string      `protobuf:"bytes,5,opt,name=city,proto3" json:"city"`           // 用户所在城市
	Province  string      `protobuf:"bytes,6,opt,name=province,proto3" json:"province"`   // 用户所在省份
	Country   string      `protobuf:"bytes,7,opt,name=country,proto3" json:"country"`     // 用户所在国家
	Phone     string      `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone"`         // 用户手机号
	Group     string      `protobuf:"bytes,9,opt,name=group,proto3" json:"group"`         // 用户所在分组列表
	Rg        bool        `protobuf:"varint,10,opt,name=rg,proto3" json:"rg"`             // 是否涨绿跌红
	Only20    bool        `protobuf:"varint,11,opt,name=only20,proto3" json:"only20"`     // 是否20日线
	Capacity  int32       `protobuf:"varint,12,opt,name=capacity,proto3" json:"capacity"` // 可用容量
	UserList  []*UserInfo `protobuf:"bytes,13,rep,name=userList,proto3" json:"userList"`  // 用户列表
	Vip       bool        `protobuf:"varint,14,opt,name=vip,proto3" json:"vip"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserInfoResp) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *GetUserInfoResp) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GetUserInfoResp) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *GetUserInfoResp) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetUserInfoResp) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetUserInfoResp) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *GetUserInfoResp) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetUserInfoResp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetUserInfoResp) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetUserInfoResp) GetRg() bool {
	if x != nil {
		return x.Rg
	}
	return false
}

func (x *GetUserInfoResp) GetOnly20() bool {
	if x != nil {
		return x.Only20
	}
	return false
}

func (x *GetUserInfoResp) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *GetUserInfoResp) GetUserList() []*UserInfo {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *GetUserInfoResp) GetVip() bool {
	if x != nil {
		return x.Vip
	}
	return false
}

type UpsetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
}

func (x *UpsetUserInfoReq) Reset() {
	*x = UpsetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsetUserInfoReq) ProtoMessage() {}

func (x *UpsetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsetUserInfoReq.ProtoReflect.Descriptor instead.
func (*UpsetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{9}
}

func (x *UpsetUserInfoReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpsetUserInfoReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ReLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relogin     bool   `protobuf:"varint,1,opt,name=relogin,proto3" json:"relogin"`         // 是否重新登录:true 重新登录
	SessionId   string `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId"`      // 用户sessionid 用于当前交互使用
	OpenId      string `protobuf:"bytes,3,opt,name=openId,proto3" json:"openId"`            // openid
	OverdueTime int64  `protobuf:"varint,4,opt,name=overdueTime,proto3" json:"overdueTime"` // 逾期时间点(时间戳)
}

func (x *ReLoginResp) Reset() {
	*x = ReLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ccmy_weixin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReLoginResp) ProtoMessage() {}

func (x *ReLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_ccmy_weixin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReLoginResp.ProtoReflect.Descriptor instead.
func (*ReLoginResp) Descriptor() ([]byte, []int) {
	return file_ccmy_weixin_proto_rawDescGZIP(), []int{10}
}

func (x *ReLoginResp) GetRelogin() bool {
	if x != nil {
		return x.Relogin
	}
	return false
}

func (x *ReLoginResp) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReLoginResp) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *ReLoginResp) GetOverdueTime() int64 {
	if x != nil {
		return x.OverdueTime
	}
	return 0
}

var File_ccmy_weixin_proto protoreflect.FileDescriptor

var file_ccmy_weixin_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x63, 0x6d, 0x79, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x63, 0x6d, 0x79, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6a, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x63, 0x0a,
	0x09, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xf4, 0x01, 0x0a, 0x0a, 0x57, 0x78, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x76, 0x69, 0x70, 0x22, 0x1e, 0x0a, 0x0e, 0x57, 0x78, 0x55,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x62, 0x22, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3d, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xf3, 0x02, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52,
	0x4c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x72, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x72, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x79, 0x32, 0x30, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x6f, 0x6e, 0x6c, 0x79, 0x32, 0x30, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x69, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x76, 0x69,
	0x70, 0x22, 0x3a, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a,
	0x0b, 0x52, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72,
	0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xc1,
	0x02, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x0e, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x57,
	0x78, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x63,
	0x6d, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x63, 0x63, 0x6d, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0d, 0x55, 0x70, 0x73,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x63, 0x63, 0x6d,
	0x79, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x52, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e,
	0x63, 0x63, 0x6d, 0x79, 0x2e, 0x52, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x63, 0x6d, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ccmy_weixin_proto_rawDescOnce sync.Once
	file_ccmy_weixin_proto_rawDescData = file_ccmy_weixin_proto_rawDesc
)

func file_ccmy_weixin_proto_rawDescGZIP() []byte {
	file_ccmy_weixin_proto_rawDescOnce.Do(func() {
		file_ccmy_weixin_proto_rawDescData = protoimpl.X.CompressGZIP(file_ccmy_weixin_proto_rawDescData)
	})
	return file_ccmy_weixin_proto_rawDescData
}

var file_ccmy_weixin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ccmy_weixin_proto_goTypes = []interface{}{
	(*OauthReq)(nil),         // 0: ccmy.OauthReq
	(*OauthResp)(nil),        // 1: ccmy.OauthResp
	(*WxUserinfo)(nil),       // 2: ccmy.WxUserinfo
	(*WxUserinfoResp)(nil),   // 3: ccmy.WxUserinfoResp
	(*GetQrcodeReq)(nil),     // 4: ccmy.GetQrcodeReq
	(*GetQrcodeResp)(nil),    // 5: ccmy.GetQrcodeResp
	(*GroupInfo)(nil),        // 6: ccmy.GroupInfo
	(*UserInfo)(nil),         // 7: ccmy.userInfo
	(*GetUserInfoResp)(nil),  // 8: ccmy.GetUserInfoResp
	(*UpsetUserInfoReq)(nil), // 9: ccmy.UpsetUserInfoReq
	(*ReLoginResp)(nil),      // 10: ccmy.ReLoginResp
	(*common.Empty)(nil),     // 11: common.Empty
}
var file_ccmy_weixin_proto_depIdxs = []int32{
	6,  // 0: ccmy.userInfo.group:type_name -> ccmy.GroupInfo
	7,  // 1: ccmy.GetUserInfoResp.userList:type_name -> ccmy.userInfo
	0,  // 2: ccmy.Weixin.Oauth:input_type -> ccmy.OauthReq
	2,  // 3: ccmy.Weixin.UpdateUserInfo:input_type -> ccmy.WxUserinfo
	4,  // 4: ccmy.Weixin.GetQrcode:input_type -> ccmy.GetQrcodeReq
	11, // 5: ccmy.Weixin.GetUserInfo:input_type -> common.Empty
	9,  // 6: ccmy.Weixin.UpsetUserInfo:input_type -> ccmy.UpsetUserInfoReq
	11, // 7: ccmy.Weixin.ReLogin:input_type -> common.Empty
	1,  // 8: ccmy.Weixin.Oauth:output_type -> ccmy.OauthResp
	11, // 9: ccmy.Weixin.UpdateUserInfo:output_type -> common.Empty
	5,  // 10: ccmy.Weixin.GetQrcode:output_type -> ccmy.GetQrcodeResp
	8,  // 11: ccmy.Weixin.GetUserInfo:output_type -> ccmy.GetUserInfoResp
	11, // 12: ccmy.Weixin.UpsetUserInfo:output_type -> common.Empty
	10, // 13: ccmy.Weixin.ReLogin:output_type -> ccmy.ReLoginResp
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_ccmy_weixin_proto_init() }
func file_ccmy_weixin_proto_init() {
	if File_ccmy_weixin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ccmy_weixin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthReq); i {
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
		file_ccmy_weixin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthResp); i {
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
		file_ccmy_weixin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxUserinfo); i {
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
		file_ccmy_weixin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxUserinfoResp); i {
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
		file_ccmy_weixin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQrcodeReq); i {
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
		file_ccmy_weixin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQrcodeResp); i {
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
		file_ccmy_weixin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
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
		file_ccmy_weixin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_ccmy_weixin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_ccmy_weixin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsetUserInfoReq); i {
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
		file_ccmy_weixin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReLoginResp); i {
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
			RawDescriptor: file_ccmy_weixin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ccmy_weixin_proto_goTypes,
		DependencyIndexes: file_ccmy_weixin_proto_depIdxs,
		MessageInfos:      file_ccmy_weixin_proto_msgTypes,
	}.Build()
	File_ccmy_weixin_proto = out.File
	file_ccmy_weixin_proto_rawDesc = nil
	file_ccmy_weixin_proto_goTypes = nil
	file_ccmy_weixin_proto_depIdxs = nil
}
