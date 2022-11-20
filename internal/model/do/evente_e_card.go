// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECard is the golang structure of table evente_e_card for DAO operations like Where/Data.
type EventeECard struct {
	g.Meta             `orm:"table:evente_e_card, do:true"`
	Id                 interface{} //
	OrgId              interface{} // 主办id
	CardName           interface{} // E通卡名称
	SaleStartDate      *gtime.Time // 售卖开始时间
	SaleStopDate       *gtime.Time // 售卖结束时间
	Colour             interface{} // 年卡背景颜色
	ImgId              interface{} // 年卡图片id
	IssueNum           interface{} // 发行量
	Price              interface{} // 价格
	ValidStatus        interface{} // 有效期状态（1.购买日自然年有效 2.购买日多少天有效 3.自定义）
	ValidUntil         interface{} // 开始结束时间用，分割/多少天有效
	Usage              interface{} // 使用方式 1、预约 2、免预约
	BeforehandNum      interface{} // 总预约次数
	UserInformation    interface{} // 1、姓名 2、手机号 3、身份证 4、一寸照 多个，分割
	CardExplain        interface{} // E通卡说明
	BuyNum             interface{} // 直接购买量
	ActivationCodeNum  interface{} // 激活码生成总量
	ActivationUsageNum interface{} // 激活码使用量（激活量）
	Stock              interface{} // 库存
	SaleStatus         interface{} // 售卖状态 1、开启售卖 2、关闭售卖
	CreateDate         *gtime.Time // 创建时间
	UpdateDate         *gtime.Time // 修改时间
	CheckNum           interface{} // 每天可核销数量
	TextColour         interface{} // 年卡文字颜色
	Cancellation       interface{} // 是否可取消预约
	CancellationTime   interface{} // 年卡可取消时间序列化
	GiveMember         interface{} // 赠送会员等级
	DeliveryAddress    interface{} // 商品取货地址
	ContainGoods       interface{} // 是否包含商品
}
