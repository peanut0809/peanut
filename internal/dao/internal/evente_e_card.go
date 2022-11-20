// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardDao is the data access object for table evente_e_card.
type EventeECardDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EventeECardColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardColumns defines and stores column names for table evente_e_card.
type EventeECardColumns struct {
	Id                 string //
	OrgId              string // 主办id
	CardName           string // E通卡名称
	SaleStartDate      string // 售卖开始时间
	SaleStopDate       string // 售卖结束时间
	Colour             string // 年卡背景颜色
	ImgId              string // 年卡图片id
	IssueNum           string // 发行量
	Price              string // 价格
	ValidStatus        string // 有效期状态（1.购买日自然年有效 2.购买日多少天有效 3.自定义）
	ValidUntil         string // 开始结束时间用，分割/多少天有效
	Usage              string // 使用方式 1、预约 2、免预约
	BeforehandNum      string // 总预约次数
	UserInformation    string // 1、姓名 2、手机号 3、身份证 4、一寸照 多个，分割
	CardExplain        string // E通卡说明
	BuyNum             string // 直接购买量
	ActivationCodeNum  string // 激活码生成总量
	ActivationUsageNum string // 激活码使用量（激活量）
	Stock              string // 库存
	SaleStatus         string // 售卖状态 1、开启售卖 2、关闭售卖
	CreateDate         string // 创建时间
	UpdateDate         string // 修改时间
	CheckNum           string // 每天可核销数量
	TextColour         string // 年卡文字颜色
	Cancellation       string // 是否可取消预约
	CancellationTime   string // 年卡可取消时间序列化
	GiveMember         string // 赠送会员等级
	DeliveryAddress    string // 商品取货地址
	ContainGoods       string // 是否包含商品
}

// eventeECardColumns holds the columns for table evente_e_card.
var eventeECardColumns = EventeECardColumns{
	Id:                 "id",
	OrgId:              "org_id",
	CardName:           "card_name",
	SaleStartDate:      "sale_start_date",
	SaleStopDate:       "sale_stop_date",
	Colour:             "colour",
	ImgId:              "img_id",
	IssueNum:           "issue_num",
	Price:              "price",
	ValidStatus:        "valid_status",
	ValidUntil:         "valid_until",
	Usage:              "usage",
	BeforehandNum:      "beforehand_num",
	UserInformation:    "user_information",
	CardExplain:        "card_explain",
	BuyNum:             "buy_num",
	ActivationCodeNum:  "activation_code_num",
	ActivationUsageNum: "activation_usage_num",
	Stock:              "stock",
	SaleStatus:         "sale_status",
	CreateDate:         "create_date",
	UpdateDate:         "update_date",
	CheckNum:           "check_num",
	TextColour:         "text_colour",
	Cancellation:       "cancellation",
	CancellationTime:   "cancellation_time",
	GiveMember:         "give_member",
	DeliveryAddress:    "delivery_address",
	ContainGoods:       "contain_goods",
}

// NewEventeECardDao creates and returns a new DAO object for table data access.
func NewEventeECardDao() *EventeECardDao {
	return &EventeECardDao{
		group:   "default",
		table:   "evente_e_card",
		columns: eventeECardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardDao) Columns() EventeECardColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}