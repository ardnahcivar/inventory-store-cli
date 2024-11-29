package models

import "time"

type InventoryItem struct {
	Name             string
	Price            float64
	ManufacturedDate time.Time
	Brand            string
	Id               int64
}
