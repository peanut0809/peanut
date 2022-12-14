// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeOrgPartyAuth is the golang structure for table evente_org_party_auth.
type EventeOrgPartyAuth struct {
	Id                     uint        `json:"id"                     description:""`
	OrgId                  int         `json:"orgId"                  description:"org_id"`
	AppType                int         `json:"appType"                description:"app 类型，1:小程序，2:小游戏"`
	AppState               int         `json:"appState"               description:"app 状态，0:未发布，1:已发布，2:已下线"`
	PartyType              string      `json:"partyType"              description:"第三方工具类型"`
	NickName               string      `json:"nickName"               description:"授权方昵称"`
	UserName               string      `json:"userName"               description:"授权方原始ID"`
	PrincipalName          string      `json:"principalName"          description:"第三方的主体名称"`
	HeadImg                string      `json:"headImg"                description:"授权方头像"`
	QrcodeUrl              string      `json:"qrcodeUrl"              description:"二维码"`
	AuthorizerAppid        string      `json:"authorizerAppid"        description:"app_id"`
	AuthorizerRefreshToken string      `json:"authorizerRefreshToken" description:" 有效期 1 个月，且只可使用一次，使用后失效。"`
	FuncInfo               string      `json:"funcInfo"               description:"第三方授权给开发者的权限集列表 序列化"`
	AuthorizerInfo         string      `json:"authorizerInfo"         description:"第三方返回的序列化数据集合"`
	CreateDate             *gtime.Time `json:"createDate"             description:"新增时间"`
	UpdateDate             *gtime.Time `json:"updateDate"             description:"更新时间"`
	DeleteDate             *gtime.Time `json:"deleteDate"             description:"删除时间"`
	Status                 int         `json:"status"                 description:"状态 1:正常 2:取消授权"`
}
