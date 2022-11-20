// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardAppointmentDao is the data access object for table evente_e_card_appointment.
type EventeECardAppointmentDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns EventeECardAppointmentColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardAppointmentColumns defines and stores column names for table evente_e_card_appointment.
type EventeECardAppointmentColumns struct {
	Id                  string // 自增id
	IncrementId         string //
	UserId              string // C端用户id
	OrgId               string // 主办id
	Mobile              string // 手机号
	Name                string // 预约联系人
	CardId              string // E通卡id
	CardSaleId          string //
	PicId               string // 活动图片id
	EventeId            string // 活动id
	EventeName          string // 活动名称
	ScreeningsId        string // 场次id
	ScreeningsName      string // 场次名称
	ScreeningsStartDate string // 场次开始时间
	ScreeningsEndDate   string // 场次结束时间
	Number              string // 预约数量
	Status              string // 状态 1预约 2取消预约 3已核销
	VerificationName    string // 核销人-主办名称
	VerificationDate    string // 核销时间
	CreateDate          string // 创建时间
	UpdateDate          string // 修改时间
	UseNumber           string // 本次消耗年卡次数
}

// eventeECardAppointmentColumns holds the columns for table evente_e_card_appointment.
var eventeECardAppointmentColumns = EventeECardAppointmentColumns{
	Id:                  "id",
	IncrementId:         "increment_id",
	UserId:              "user_id",
	OrgId:               "org_id",
	Mobile:              "mobile",
	Name:                "name",
	CardId:              "card_id",
	CardSaleId:          "card_sale_id",
	PicId:               "pic_id",
	EventeId:            "evente_id",
	EventeName:          "evente_name",
	ScreeningsId:        "screenings_id",
	ScreeningsName:      "screenings_name",
	ScreeningsStartDate: "screenings_start_date",
	ScreeningsEndDate:   "screenings_end_date",
	Number:              "number",
	Status:              "status",
	VerificationName:    "verification_name",
	VerificationDate:    "verification_date",
	CreateDate:          "create_date",
	UpdateDate:          "update_date",
	UseNumber:           "use_number",
}

// NewEventeECardAppointmentDao creates and returns a new DAO object for table data access.
func NewEventeECardAppointmentDao() *EventeECardAppointmentDao {
	return &EventeECardAppointmentDao{
		group:   "default",
		table:   "evente_e_card_appointment",
		columns: eventeECardAppointmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardAppointmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardAppointmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardAppointmentDao) Columns() EventeECardAppointmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardAppointmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardAppointmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardAppointmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}