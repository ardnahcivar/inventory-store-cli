package models

import (
	"time"
)

type InventoryItem struct {
	Name             string
	Price            float64
	ManufacturedDate time.Time
	Brand            string
	Id               int64
	//Quantity         string
}

// func (i InventoryItem) String() string {
// 	return fmt.Sprintf("Item is: %s,Price:%v,ManufacturedDate is %v, Brand is %s, Id is %v", i.Name, i.Price, i.ManufacturedDate, i.Brand, i.Id)
// }
