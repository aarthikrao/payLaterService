package utils

import (
	"fmt"
	"testing"
)

func TestStrToFloat32(t *testing.T) {
	val, err := StrToFloat32("100.123")
	if err != nil {
		t.Errorf("Error in conversion")
		return
	}
	if val != 100.123 {
		t.Errorf("Expected 100.123 but found " + fmt.Sprintf("%f", val))
	}
}

func TestGetAmountAfterInterest(t *testing.T) {
	mAmount, oAmount := GetAmountAfterInterest(100.00, 12)
	if mAmount != 88 || oAmount != 12 {
		t.Error("Expected 88 and 12 but found", mAmount, oAmount)
	}
}

func TestStrToFloat32abcN(t *testing.T) {
	_, err := StrToFloat32("abc")
	if err == nil {
		t.Errorf("Negative case for abc failing")
		return
	}
}
