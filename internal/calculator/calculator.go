// Package calculator contains the business logic of calculating the amount of jesmonite powder required for a certain liquid amount
package calculator

func CalcPowder(liquidQuantity float64) float64 {
	liquidMultiplier := 2.5
	return float64(liquidQuantity) * liquidMultiplier
}
