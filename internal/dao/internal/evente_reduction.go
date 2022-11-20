// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeReductionDao is the data access object for table evente_reduction.
type EventeReductionDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns EventeReductionColumns // columns contains all the column names of Table for convenient usage.
}

// EventeReductionColumns defines and stores column names for table evente_reduction.
type EventeReductionColumns struct {
	Id            string //
	OrgId         string // org id
	Title         string // 标题
	ReductionMode string // reduction_mode 满减方式(1.满额减免  2.满件减免)
	ReductionType string // reduction_type 满减类型（1.仅按最高减免力度减免一次  2.每满足一次减免一次）
	StartDate     string // 开始时间
	EndDate       string // 结束时间
	CoverImg      string // 封面图ID
	Rule          string // 规则（json 格式数据）
	IsCurrency    string // 是否夸商品优惠（1是 2.否）
	ProductType   string // 适用产品类型（event 演出活动,goods 商品类型）
	ProductList   string // 产品ID集合
	Status        string // 状态（1正常 2.删除）
	CreateDate    string // 创建时间
	UpdateDate    string // 更新时间
}

// eventeReductionColumns holds the columns for table evente_reduction.
var eventeReductionColumns = EventeReductionColumns{
	Id:            "id",
	OrgId:         "org_id",
	Title:         "title",
	ReductionMode: "reduction_mode",
	ReductionType: "reduction_type",
	StartDate:     "start_date",
	EndDate:       "end_date",
	CoverImg:      "cover_img",
	Rule:          "rule",
	IsCurrency:    "is_currency",
	ProductType:   "product_type",
	ProductList:   "product_list",
	Status:        "status",
	CreateDate:    "create_date",
	UpdateDate:    "update_date",
}

// NewEventeReductionDao creates and returns a new DAO object for table data access.
func NewEventeReductionDao() *EventeReductionDao {
	return &EventeReductionDao{
		group:   "default",
		table:   "evente_reduction",
		columns: eventeReductionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeReductionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeReductionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeReductionDao) Columns() EventeReductionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeReductionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeReductionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeReductionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
