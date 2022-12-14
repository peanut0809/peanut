// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeLimitDiscountExtendDao is the data access object for table evente_limit_discount_extend.
type EventeLimitDiscountExtendDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns EventeLimitDiscountExtendColumns // columns contains all the column names of Table for convenient usage.
}

// EventeLimitDiscountExtendColumns defines and stores column names for table evente_limit_discount_extend.
type EventeLimitDiscountExtendColumns struct {
	Id              string // 主键ID
	OrgId           string // 主办ID
	DiscountId      string // 限时折扣ID
	ProductType     string // 产品类型 （event活动 goods商品)
	ProductId       string // 产品id
	ProductName     string // 产品名称
	ScreeningsId    string // 场次ID
	ScreeningsName  string // 场次名称
	PriceList       string // 票品id集合 jason串
	DiscountPercent string // 折扣比例
	GoodsPrice      string // 商品价格
	Status          string // 状态 1正常 2删除 // 预留
	CreateDate      string // 创建时间
	UpdateDate      string // 修改时间
}

// eventeLimitDiscountExtendColumns holds the columns for table evente_limit_discount_extend.
var eventeLimitDiscountExtendColumns = EventeLimitDiscountExtendColumns{
	Id:              "id",
	OrgId:           "org_id",
	DiscountId:      "discount_id",
	ProductType:     "product_type",
	ProductId:       "product_id",
	ProductName:     "product_name",
	ScreeningsId:    "screenings_id",
	ScreeningsName:  "screenings_name",
	PriceList:       "price_list",
	DiscountPercent: "discount_percent",
	GoodsPrice:      "goods_price",
	Status:          "status",
	CreateDate:      "create_date",
	UpdateDate:      "update_date",
}

// NewEventeLimitDiscountExtendDao creates and returns a new DAO object for table data access.
func NewEventeLimitDiscountExtendDao() *EventeLimitDiscountExtendDao {
	return &EventeLimitDiscountExtendDao{
		group:   "default",
		table:   "evente_limit_discount_extend",
		columns: eventeLimitDiscountExtendColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeLimitDiscountExtendDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeLimitDiscountExtendDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeLimitDiscountExtendDao) Columns() EventeLimitDiscountExtendColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeLimitDiscountExtendDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeLimitDiscountExtendDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeLimitDiscountExtendDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
