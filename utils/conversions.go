package utils

import (
	"strconv"
)

// StrToFloat32 : converts string to float
func StrToFloat32(f string) (floatString float32, err error) {
	s, err := strconv.ParseFloat(f, 32)
	if err == nil {
		floatString = float32(s)
	}
	return
}

// GetAmountAfterInterest is used to get the split amount after merchant disount
func GetAmountAfterInterest(totalAmount float32, rateOfInterest float32) (merchantAmount float32, ourDiscount float32) {
	ourDiscount = totalAmount * rateOfInterest / 100
	merchantAmount = totalAmount - ourDiscount
	return
}
