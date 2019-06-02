package services

import (
	"fmt"

	"github.com/aarthikrao/payLaterService/models"

	"github.com/aarthikrao/payLaterService/utils"

	"github.com/aarthikrao/payLaterService/datalayer"
)

// TransactionService : All transaction happen here
type TransactionService struct {
	ud *datalayer.UserData
	md *datalayer.MerchantData
	td *datalayer.TransactionData
}

// NewTransactionService is used to create new insatnces of transactions
func NewTransactionService() *TransactionService {
	return &TransactionService{
		ud: datalayer.NewUserData(),
		md: datalayer.NewMerchantData(),
		td: datalayer.NewTransactionData(),
	}
}

// TransferToMerchant is used to transfer the amount from user to merchant
// Args : user-name merchant-name amount
func (ts *TransactionService) TransferToMerchant(args []string) bool {
	if len(args) != 3 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	amount, err := utils.StrToFloat32(args[2])
	if err != nil {
		fmt.Println("Please enter a valid amount:", amount)
		return false
	}
	transaction := models.Transaction{
		UserName:     args[0],
		MerchantName: args[1],
		TotalAmount:  amount,
	}
	// Validating user
	user, err := ts.ud.GetUserByName(transaction.UserName)
	if err != nil || user.Name == "" {
		fmt.Println("Invalid username")
		return false
	}

	// Validate merchant
	merchant, err := ts.md.GetMerchantByName(transaction.MerchantName)
	if err != nil || merchant.Name == "" {
		fmt.Println("Invalid Merchant")
		return false
	}
	// Calculate the amount after interest and populate the transaction struct
	merchantAmount, ourDiscount := utils.GetAmountAfterInterest(transaction.TotalAmount, merchant.InterestRate)
	transaction.MerchantAmount = merchantAmount
	transaction.OurDiscount = ourDiscount
	transaction.InterestRate = merchant.InterestRate

	if !transaction.ValidateAndClean() {
		fmt.Println("Invalid transaction")
		return false
	}
	// Begin transaction
	returnResult, err := ts.td.RunTransaction(transaction)
	if err != nil {
		fmt.Println("Transaction unsuccessful")
		return false
	}
	return returnResult
}
