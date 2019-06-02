package services

import "testing"

func TestCreateMerchantInvalidRateN(t *testing.T) {
	actual := MS.CreateMerchant([]string{"merchantTest", "abc"})
	if actual {
		t.Errorf("Merchant Created despite invalid rate of interest")
	}
}

func TestCreateMerchantInvalidArgsN(t *testing.T) {
	actual := MS.CreateMerchant([]string{"merchantTest", "10", "some random param"})
	if actual {
		t.Errorf("Merchant created despite invalid args")
	}
}

func TestCreateMerchantInvalidInterestN(t *testing.T) {
	actual := MS.CreateMerchant([]string{"merchantTest", "-10", "some random param"})
	if actual {
		t.Errorf("Merchant created despite invalid interest rate")
	}
}

func TestGetMerchantDiscountInvalidArgsN(t *testing.T) {
	actual := MS.GetMerchantDiscount([]string{"merchantTest"})
	if actual {
		t.Errorf("GetMerchantDiscount passed despite InvalidArgs")
	}
}

func TestGetMerchantDiscountInvalidMerchantN(t *testing.T) {
	actual := MS.GetMerchantDiscount([]string{"discount", "merchantTest123"})
	if actual {
		t.Errorf("GetMerchantDiscount passed despite Invalid Merchant")
	}
}

func TestChangeMerchantInterest110N(t *testing.T) {
	actual := MS.ChangeMerchantInterest([]string{"merchantTest", "110"})
	if actual {
		t.Errorf("TestChangeMerchantInterest110N passed despite Invalid Rate of interest")
	}
}
