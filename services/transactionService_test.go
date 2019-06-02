package services

import (
	"fmt"
	"os"
	"testing"

	conn "github.com/aarthikrao/payLaterService/connections"

	"github.com/aarthikrao/payLaterService/config"
)

// US is used for all user related operations
var US *UserService

// MS is used for all merchnat related operations
var MS *MerchantService

// TS is used for transaction related services
var TS *TransactionService

// INITServices is used to initiate services
func INITServices() {
	US = NewUserService()
	MS = NewMerchantService()
	TS = NewTransactionService()
}

// TestMain : Add all setup and teardown here
func TestMain(m *testing.M) {
	// Initialise configs
	config.InitConfigs()

	// Initialise the postgres db connection
	conn.InitDatabase()
	// Close db connection at shutdown
	defer conn.ShutDown()

	INITServices()

	// Run all the tests and fetch the exit code
	runResult := m.Run()

	// Clearing all DB data related to test
	conn.PGDB.Exec("delete from merchants where name= 'merchantTest'")
	conn.PGDB.Exec("delete from users where name = 'userTest'")
	fmt.Println("Deleting merchantTest and userTest after test ")

	os.Exit(runResult)
}

// TestCreateUser : Create a user
func TestCreateUser(t *testing.T) {
	actual := US.CreateUser([]string{"userTest", "test@test.com", "100"})
	if !actual {
		t.Errorf("User was not created")
	}
}

// #################################### MerchantService #############################

func TestCreateMerchant(t *testing.T) {
	actual := MS.CreateMerchant([]string{"merchantTest", "10"})
	if !actual {
		t.Errorf("Merchant not created")
	}
}

func TestChangeMerchantInterest(t *testing.T) {
	actual := MS.ChangeMerchantInterest([]string{"merchantTest", "10"})
	if !actual {
		t.Errorf("ChangeMerchantInterest failed")
	}
}

func TestGetMerchantDiscount(t *testing.T) {
	actual := MS.GetMerchantDiscount([]string{"discount", "merchantTest"})
	if !actual {
		t.Errorf("GetMerchantDiscount failed")
	}
}

// ############################################## Transaction ###################################
// We are running a transaction assuming the user and merchant are already created in the tests above
func TestRunTransaction(t *testing.T) {
	actual := TS.TransferToMerchant([]string{"userTest", "merchantTest", "10"})
	if !actual {
		t.Errorf("Transaction was not successful")
	}
}

func TestRunTransactionInvalidUser(t *testing.T) {
	actual := TS.TransferToMerchant([]string{"userTestInvalid", "merchantTest", "10"})
	if actual {
		t.Errorf("Transaction was successful despite invalid user")
	}
}

func TestRunTransactionInvalidMerchant(t *testing.T) {
	actual := TS.TransferToMerchant([]string{"userTest", "merchantTestInvalid", "10"})
	if actual {
		t.Errorf("Transaction was successful despite invalid merchant")
	}
}

func TestRunTransactionInvalidAmount(t *testing.T) {
	actual := TS.TransferToMerchant([]string{"userTest", "merchantTest", "10000"})
	if actual {
		t.Errorf("Transaction was successful despite invalid amount")
	}
}

func TestRunTransactionStringAmount(t *testing.T) {
	actual := TS.TransferToMerchant([]string{"userTest", "merchantTest", "abcd"})
	if actual {
		t.Errorf("Transaction was successful despite invalid merchant")
	}
}
