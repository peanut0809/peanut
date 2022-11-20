// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeIntegralOrderListDao is internal type for wrapping internal DAO implements.
type internalEventeIntegralOrderListDao = *internal.EventeIntegralOrderListDao

// eventeIntegralOrderListDao is the data access object for table evente_integral_order_list.
// You can define custom methods on it to extend its functionality as you wish.
type eventeIntegralOrderListDao struct {
	internalEventeIntegralOrderListDao
}

var (
	// EventeIntegralOrderList is globally public accessible object for table evente_integral_order_list operations.
	EventeIntegralOrderList = eventeIntegralOrderListDao{
		internal.NewEventeIntegralOrderListDao(),
	}
)

// Fill with you ideas below.
