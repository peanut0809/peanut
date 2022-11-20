// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeECardGoodsDao is internal type for wrapping internal DAO implements.
type internalEventeECardGoodsDao = *internal.EventeECardGoodsDao

// eventeECardGoodsDao is the data access object for table evente_e_card_goods.
// You can define custom methods on it to extend its functionality as you wish.
type eventeECardGoodsDao struct {
	internalEventeECardGoodsDao
}

var (
	// EventeECardGoods is globally public accessible object for table evente_e_card_goods operations.
	EventeECardGoods = eventeECardGoodsDao{
		internal.NewEventeECardGoodsDao(),
	}
)

// Fill with you ideas below.