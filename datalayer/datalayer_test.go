package datalayer

import (
	"fmt"
	"os"
	"testing"

	"github.com/aarthikrao/payLaterService/models"

	"github.com/aarthikrao/payLaterService/config"
	conn "github.com/aarthikrao/payLaterService/connections"
)

var UD *UserData
var MD *MerchantData
var TD *TransactionData

var merchantName = "merchantTest"

var interestRate float32 = 10.00

var userName = "userTest"

func INITDataLayer() {
	UD = NewUserData()
	MD = NewMerchantData()
	TD = NewTransactionData()
}

func TestMain(m *testing.M) {
	// Initialise configs
	config.InitConfigs()

	// Initialise the postgres db connection
	conn.InitDatabase()
	// Close db connection at shutdown
	defer conn.ShutDown()

	INITDataLayer()

	// Run all the tests and fetch the exit code
	runResult := m.Run()

	// Clearing all DB data related to test
	conn.PGDB.Exec("delete from merchants where name= 'merchantTest'")
	conn.PGDB.Exec("delete from users where name = 'userTest'")
	conn.PGDB.Exec("delete from transactions where user_name = 'userTest'")
	fmt.Println("Deleting merchantTest and userTest after test ")

	os.Exit(runResult)
}

// TestCreateUser : Create a user
func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:        userName,
		CreditLimit: float32(1000.00),
		Email:       "test@test.com",
	}
	err := UD.AddNewUser(user)
	if err != nil {
		t.Errorf("User was not created")
	}
}

func TestGetUserByName(t *testing.T) {
	user, err := UD.GetUserByName(userName)
	if err != nil {
		t.Errorf("User was not created")
		return
	}
	if user.CreditLimit != float32(1000.00) {
		t.Errorf("Amount is incorrect")
		return
	}
}

func TestCreateMerchant(t *testing.T) {
	merchant := models.Merchant{
		Name:         merchantName,
		InterestRate: interestRate,
	}
	if !merchant.ValidateAndClean() {
		t.Errorf("Merchant validity failed")
		return
	}
	err := MD.AddNewMerchant(merchant)
	if err != nil {
		t.Errorf("Error in creating merchant")
		return
	}
}

func TestGetMerchantByName(t *testing.T) {
	merchant, err := MD.GetMerchantByName(merchantName)
	if err != nil {
		t.Errorf("merchant was not found")
		return
	}
	if merchant.InterestRate != interestRate {
		t.Errorf("Error in interest rate")
	}
}

func TestTransaction(t *testing.T) {
	transaction := models.Transaction{
		InterestRate:   float32(10),
		MerchantAmount: float32(90),
		OurDiscount:    float32(10),
		MerchantName:   merchantName,
		UserName:       userName,
		TotalAmount:    float32(100),
	}
	result, err := TD.RunTransaction(transaction)
	if err != nil {
		t.Errorf("Error in running transaction")
		return
	}
	if !result {
		t.Errorf("Error in transaction")
		return
	}
}
