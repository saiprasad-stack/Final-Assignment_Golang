package models

import "time"

type BillingData struct {
	ID             uint
	MeterID        uint
	BillDate       time.Time
	ImportKwhTotal float64
}
