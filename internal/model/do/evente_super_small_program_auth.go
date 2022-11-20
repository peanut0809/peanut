// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeSuperSmallProgramAuth is the golang structure of table evente_super_small_program_auth for DAO operations like Where/Data.
type EventeSuperSmallProgramAuth struct {
	g.Meta                 `orm:"table:evente_super_small_program_auth, do:true"`
	Id                     interface{} //
	NickName               interface{} // 授权方昵称
	ServiceTypeInfo        interface{} // 类型，0订阅号，1升级后订阅号，2服务号
	VerifyTypeInfo         interface{} // 认证类型，-1未认证，0微信认证，1新浪微博认证，2腾讯微博认证，3:0+未通过名称认证，4:1+未通过名称认证，5:2+未通过名称认证
	UserName               interface{} // 授权方小程序的原始ID
	PrincipalName          interface{} // 小程序的主体名称
	Alias                  interface{} // 小程序的主体名称
	HeadImg                interface{} // 授权方头像
	QrcodeUrl              interface{} // 二维码
	BusinessInfo           interface{} // 商户功能 序列化
	AuthorizerAppid        interface{} // app_id
	AuthorizerRefreshToken interface{} //
	FuncInfo               interface{} // 小程序授权给开发者的权限集列表 序列化
	MiniProgramInfo        interface{} // 小程序信息 序列化
	AuthorizerInfo         interface{} // 微信返回的序列化
	CreateDate             *gtime.Time // 新增时间
	UpdateDate             *gtime.Time // 更新时间
	DeleteDate             *gtime.Time // 删除时间
	Status                 interface{} // 状态 1:正常 2:取消授权
}
