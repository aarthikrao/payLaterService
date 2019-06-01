package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/aarthikrao/payLaterService/connections"

	"github.com/aarthikrao/payLaterService/services"

	"github.com/abiosoft/ishell"
)

// For ,more info on cli library refer : https://github.com/abiosoft/ishell

// Add all possible commands here:
// New User: # new-user name email credit-limit
// New Merchant: # new-merchant name interest

// UDS is used for all user related operations
var UDS *services.UserService

// MDS is used for all merchnat related operations
var MDS *services.MerchantService

// TS is used for transaction related services
var TS *services.TransactionService

// INITServices is used to initiate services
func INITServices() {
	UDS = services.NewUserService()
	MDS = services.NewMerchantService()
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
			UDS.CreateUser(c.Args)
		},
	})

	// to add new merchant
	shell.AddCmd(&ishell.Cmd{
		Name: "new-merchant",
		Help: "usage : new-merchant name interest",
		Func: func(c *ishell.Context) {
			MDS.CreateMerchant(c.Args)
		},
	})

	// to run transaction
	shell.AddCmd(&ishell.Cmd{
		Name: "transaction",
		Help: "usage : transaction user-name merchant-name amount",
		Func: func(c *ishell.Context) {
			TS.TransferToMerchent(c.Args)
		},
	})

	// commands related to update merchant
	shell.AddCmd(&ishell.Cmd{
		Name: "update-merchant-interest",
		Help: "usage : update-merchant-interest merchant-name interest",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			fmt.Println("update-merchant-interest :", strings.Join(c.Args, " "))
		},
	})

	// commands related to update merchant
	shell.AddCmd(&ishell.Cmd{
		Name: "user-payback",
		Help: "usage : user-payback user-name amount",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			fmt.Println("user-payback :", strings.Join(c.Args, " "))
		},
	})

	// commands related to update merchant
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
				if len(c.Args) != 2 {
					fmt.Println("Incorrect input, try 'help'")
					return
				}
				fmt.Println("discount :", strings.Join(c.Args, " "))
			case "dues":
				if len(c.Args) != 2 {
					fmt.Println("Incorrect input, try 'help'")
					return
				}
				fmt.Println("dues :", strings.Join(c.Args, " "))
			case "users-at-credit-limit":
				if len(c.Args) != 1 {
					fmt.Println("Incorrect input, try 'help'")
					return
				}
				fmt.Println("users-at-credit-limit :", strings.Join(c.Args, " "))
			case "total-dues":
				if len(c.Args) != 1 {
					fmt.Println("Incorrect input, try 'help'")
					return
				}
				fmt.Println("total-dues :", strings.Join(c.Args, " "))
			default:
				fmt.Println("Invalid command. Please refer 'help'")
			}
		},
	})

	shell.Interrupt(func(c *ishell.Context, count int, str string) {
		fmt.Println("Closing db connections")
		connections.ShutDown()
		fmt.Println("Exiting Pay Later Service ...")
		os.Exit(0)
	})

	shell.Run()

}
