// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeLimitDiscountDao is internal type for wrapping internal DAO implements.
type internalEventeLimitDiscountDao = *internal.EventeLimitDiscountDao

// eventeLimitDiscountDao is the data access object for table evente_limit_discount.
// You can define custom methods on it to extend its functionality as you wish.
type eventeLimitDiscountDao struct {
	internalEventeLimitDiscountDao
}

var (
	// EventeLimitDiscount is globally public accessible object for table evente_limit_discount operations.
	EventeLimitDiscount = eventeLimitDiscountDao{
		internal.NewEventeLimitDiscountDao(),
	}
)

// Fill with you ideas below.
