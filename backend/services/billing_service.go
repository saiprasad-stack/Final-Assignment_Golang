package services

import (
	"billing-system/backend/repositories"
)

func CalculateBill(consumerID uint, date string) (float64, []map[string]interface{}, float64, float64) {

	// 🔹 Step 1: Get units from repository (CORRECT LOGIC)
	units, firstReading, lastReading := repositories.GetMonthlyConsumption(consumerID, date)

	// 🔹 Step 2: Get consumer + tariffs
	consumer := repositories.GetConsumer(consumerID)
	slabs := repositories.GetTariffs(consumer.LocationID)

	total := 0.0
	remaining := int(units)

	breakdown := []map[string]interface{}{}

	// 🔹 Step 3: Progressive billing
	for _, slab := range slabs {

		if remaining <= 0 {
			break
		}

		slabSize := slab.MaxUnits - slab.MinUnits + 1
		if slab.MinUnits == 0 {
			slabSize = slab.MaxUnits
		}

		used := remaining
		if remaining > slabSize {
			used = slabSize
		}

		amount := float64(used) * slab.Rate

		breakdown = append(breakdown, map[string]interface{}{
			"units":  used,
			"rate":   slab.Rate,
			"amount": amount,
		})

		total += amount
		remaining -= used
	}

	return total, breakdown, firstReading, lastReading
}
