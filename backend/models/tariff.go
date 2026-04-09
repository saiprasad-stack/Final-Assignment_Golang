package models

type Tariff struct {
	ID         uint
	LocationID uint
	MinUnits   int
	MaxUnits   int
	Rate       float64
}
