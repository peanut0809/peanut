// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeVoucherActivationDao is internal type for wrapping internal DAO implements.
type internalEventeVoucherActivationDao = *internal.EventeVoucherActivationDao

// eventeVoucherActivationDao is the data access object for table evente_voucher_activation.
// You can define custom methods on it to extend its functionality as you wish.
type eventeVoucherActivationDao struct {
	internalEventeVoucherActivationDao
}

var (
	// EventeVoucherActivation is globally public accessible object for table evente_voucher_activation operations.
	EventeVoucherActivation = eventeVoucherActivationDao{
		internal.NewEventeVoucherActivationDao(),
	}
)

// Fill with you ideas below.