// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeCouponDao is internal type for wrapping internal DAO implements.
type internalEventeCouponDao = *internal.EventeCouponDao

// eventeCouponDao is the data access object for table evente_coupon.
// You can define custom methods on it to extend its functionality as you wish.
type eventeCouponDao struct {
	internalEventeCouponDao
}

var (
	// EventeCoupon is globally public accessible object for table evente_coupon operations.
	EventeCoupon = eventeCouponDao{
		internal.NewEventeCouponDao(),
	}
)

// Fill with you ideas below.
