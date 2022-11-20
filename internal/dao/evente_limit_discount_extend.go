// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeLimitDiscountExtendDao is internal type for wrapping internal DAO implements.
type internalEventeLimitDiscountExtendDao = *internal.EventeLimitDiscountExtendDao

// eventeLimitDiscountExtendDao is the data access object for table evente_limit_discount_extend.
// You can define custom methods on it to extend its functionality as you wish.
type eventeLimitDiscountExtendDao struct {
	internalEventeLimitDiscountExtendDao
}

var (
	// EventeLimitDiscountExtend is globally public accessible object for table evente_limit_discount_extend operations.
	EventeLimitDiscountExtend = eventeLimitDiscountExtendDao{
		internal.NewEventeLimitDiscountExtendDao(),
	}
)

// Fill with you ideas below.
