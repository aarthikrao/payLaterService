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
// Args name interest
func (ms *MerchantService) CreateMerchant(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	interestRate, err := utils.StrToFloat32(args[1])
	if err != nil {
		fmt.Println("Please provide valid rate of interest")
		return false
	}
	merchant := models.Merchant{
		Name:          args[0],
		InterestRate:  interestRate,
		TotalAmount:   0,
		TotalDiscount: 0,
	}
	if !merchant.ValidateAndClean() {
		fmt.Println("Invalid merchant details")
		return false
	}
	err = ms.merchantData.AddNewMerchant(merchant)
	if err != nil {
		fmt.Println("Merchant creation failed ..!", err)
		return false
	}
	fmt.Printf("Created new merchant: \n %+v\n", merchant)
	return true

}

// ChangeMerchantInterest is used to add new merchants to the system
// Args : merchant-name interest
func (ms *MerchantService) ChangeMerchantInterest(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	interestRate, err := utils.StrToFloat32(args[1])
	if err != nil {
		fmt.Println("Please provide valid rate of interest")
		return false
	}
	merchant := models.Merchant{
		InterestRate: interestRate,
		Name:         args[0],
	}
	if !merchant.ValidateAndClean() {
		fmt.Println("Invalid Merchant details")
		return false
	}
	err = ms.merchantData.UpdateMerchantByName(merchant.Name, merchant)
	if err != nil {
		fmt.Println("Invalid Merchant details")
		return false
	}
	return true
}

// GetMerchantDiscount is used to add new merchants to the system
// Args : discount merchant-name
func (ms *MerchantService) GetMerchantDiscount(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	merchant, err := ms.merchantData.GetMerchantByName(args[1])
	if err != nil {
		fmt.Println("Error in fetching details for merchant:", args[1])
		return false
	}
	fmt.Println("Merchant name:", merchant.Name, " Total amount recieved:", merchant.TotalAmount, " Discount:", merchant.TotalDiscount, " Current interest rate :", merchant.InterestRate)
	return true
}
