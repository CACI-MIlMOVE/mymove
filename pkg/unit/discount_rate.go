package unit

// DiscountRate represents a percentage, usually represented as a value between 0%
// and 100%.
type DiscountRate float64

// Float64 returns the Rate's value as a float64 with 1 representing one whole unit.
func (r DiscountRate) Float64() float64 {
	return float64(r)
}

// Apply returns the remaining charge after applying a DiscountRate.
//
// Note: The value returned is calculated by first determining the Percent of Tariff by
// subtracting the discount rate from 1 and then multiplying the result by the
// Cents parameter.
func (r DiscountRate) Apply(c Cents) Cents {
	percentOfTariff := 1 - r.Float64()
	return c.MultiplyFloat64(percentOfTariff)
}

// NewDiscountRateFromPercent createDecimals a new DiscountRate using a float64 with
// 100 representing one whole unit.
func NewDiscountRateFromPercent(input float64) DiscountRate {
	return DiscountRate(input / 100.0)
}
