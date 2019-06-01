package services

import (
	"fmt"

	"github.com/aarthikrao/payLaterService/models"

	"github.com/aarthikrao/payLaterService/utils"

	"github.com/aarthikrao/payLaterService/datalayer"
)

// MerchantService : merchant operations will happen here
type MerchantService struct {
	merchantData *datalayer.MerchantData
}

// NewMerchantService is used to create new instance of merchant service
func NewMerchantService() *MerchantService {
	return &MerchantService{
		merchantData: datalayer.NewMerchantData(),
	}
}

// CreateMerchant is used to add new merchants to the system
func (ms *MerchantService) CreateMerchant(args []string) {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	interestRate, err := utils.StrToFloat32(args[1])
	if err != nil {
		fmt.Println("Please provide a valid interest rate")
		return
	}
	merchant := models.Merchant{
		Name:          args[0],
		InterestRate:  interestRate,
		TotalAmount:   0,
		TotalDiscount: 0,
	}
	if !merchant.ValidateAndClean() {
		fmt.Println("Invalid merchant details")
	} else {
		err := ms.merchantData.AddNewMerchant(merchant)
		if err != nil {
			fmt.Println("Merchant creation failed ..!", err)
			return
		}
		fmt.Printf("Created new merchant: %+v\n", merchant)
		return
	}
	fmt.Println("Merchant creation failed ..!")
}

// ChangeMerchantInterest is used to add new merchants to the system
func (ms *MerchantService) ChangeMerchantInterest(args []string) {

}

// GetAllMerchantDiscount is used to add new merchants to the system
func (ms *MerchantService) GetAllMerchantDiscount(args []string) {

}
