// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeOrgPartyAuth is the golang structure of table evente_org_party_auth for DAO operations like Where/Data.
type EventeOrgPartyAuth struct {
	g.Meta                 `orm:"table:evente_org_party_auth, do:true"`
	Id                     interface{} //
	OrgId                  interface{} // org_id
	AppType                interface{} // app 类型，1:小程序，2:小游戏
	AppState               interface{} // app 状态，0:未发布，1:已发布，2:已下线
	PartyType              interface{} // 第三方工具类型
	NickName               interface{} // 授权方昵称
	UserName               interface{} // 授权方原始ID
	PrincipalName          interface{} // 第三方的主体名称
	HeadImg                interface{} // 授权方头像
	QrcodeUrl              interface{} // 二维码
	AuthorizerAppid        interface{} // app_id
	AuthorizerRefreshToken interface{} //  有效期 1 个月，且只可使用一次，使用后失效。
	FuncInfo               interface{} // 第三方授权给开发者的权限集列表 序列化
	AuthorizerInfo         interface{} // 第三方返回的序列化数据集合
	CreateDate             *gtime.Time // 新增时间
	UpdateDate             *gtime.Time // 更新时间
	DeleteDate             *gtime.Time // 删除时间
	Status                 interface{} // 状态 1:正常 2:取消授权
}
