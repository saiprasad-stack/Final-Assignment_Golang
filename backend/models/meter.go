package models

import "time" // ✅ REQUIRED

type Meter struct {
	ID         uint
	ConsumerID uint
	Reading    float64
	Date       time.Time
}
