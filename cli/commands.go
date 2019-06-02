package cli

import (
	"fmt"
	"os"

	"github.com/aarthikrao/payLaterService/connections"

	"github.com/aarthikrao/payLaterService/services"

	"github.com/abiosoft/ishell"
)

// For ,more info on cli library refer : https://github.com/abiosoft/ishell

// #########################  Add all possible commands here ############################
//
// New User: # new-user name email credit-limit
// New Merchant: # new-merchant name interest
// Transaction : # transaction user-name merchant-name amount
// Update merchant interest rate : # update-merchant-interest merchant-name interest
// Merchant details: # report discount merchant-name
// User payback : # user-payback user-name amount
// User dues for a particular user # report dues user-name
// Find users at credit limit # report users-at-credit-limit
// All user dues # report total-dues
//
// ######################################################################################

// US is used for all user related operations
var US *services.UserService

// MS is used for all merchnat related operations
var MS *services.MerchantService

// TS is used for transaction related services
var TS *services.TransactionService

// INITServices is used to initiate services
func INITServices() {
	US = services.NewUserService()
	MS = services.NewMerchantService()
	TS = services.NewTransactionService()
}

// RunCli : Add all the commandline logic here
func RunCli() {
	shell := ishell.New()

	// to add new user
	shell.AddCmd(&ishell.Cmd{
		Name: "new-user",
		Help: "usage : new-user name email credit-limit",
		Func: func(c *ishell.Context) {
			US.CreateUser(c.Args)
		},
	})

	// to add new merchant
	shell.AddCmd(&ishell.Cmd{
		Name: "new-merchant",
		Help: "usage : new-merchant name interest",
		Func: func(c *ishell.Context) {
			MS.CreateMerchant(c.Args)
		},
	})

	// to run transaction
	shell.AddCmd(&ishell.Cmd{
		Name: "transaction",
		Help: "usage : transaction user-name merchant-name amount",
		Func: func(c *ishell.Context) {
			TS.TransferToMerchant(c.Args)
		},
	})

	// commands related to update merchant interest
	shell.AddCmd(&ishell.Cmd{
		Name: "update-merchant-interest",
		Help: "usage : update-merchant-interest merchant-name interest",
		Func: func(c *ishell.Context) {
			MS.ChangeMerchantInterest(c.Args)
		},
	})

	// commands related to user-payback
	shell.AddCmd(&ishell.Cmd{
		Name: "user-payback",
		Help: "usage : user-payback user-name amount",
		Func: func(c *ishell.Context) {
			US.Payback(c.Args)
		},
	})

	// commands related to reporting
	shell.AddCmd(&ishell.Cmd{
		Name: "report",
		Help: "usage : report discount merchant-name || report dues user-name || report users-at-credit-limit || report total-dues",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			switch c.Args[0] {
			case "discount":
				MS.GetMerchantDiscount(c.Args)
			case "dues":
				US.GetUserDues(c.Args)
			case "users-at-credit-limit":
				US.GetUsersAtCreditLimit(c.Args)
			case "total-dues":
				US.GetTotalUserDues(c.Args)
			default:
				fmt.Println("Invalid command. Please refer 'help'")
			}
		},
	})

	// add all program termination related logic here
	shell.Interrupt(func(c *ishell.Context, count int, str string) {
		fmt.Println("Closing db connections")
		connections.ShutDown()
		fmt.Println("Exiting Pay Later Service ...")
		os.Exit(0)
	})

	shell.Run()

}
