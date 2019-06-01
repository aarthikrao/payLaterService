package services

import (
	"fmt"

	"github.com/aarthikrao/payLaterService/datalayer"
	"github.com/aarthikrao/payLaterService/utils"

	"github.com/aarthikrao/payLaterService/models"
)

// UserService : user operations and validations will happen here
type UserService struct {
	userData *datalayer.UserData
}

// NewUserService is used to create a new instance of UserService
func NewUserService() *UserService {
	return &UserService{
		userData: datalayer.NewUserData(),
	}
}

// CreateUser is used to add new users to the system
func (us *UserService) CreateUser(args []string) {
	if len(args) != 3 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	creditLimit, err := utils.StrToFloat32(args[2])
	if err != nil {
		fmt.Println("Please provide a number for credit limit")
		return
	}
	user := models.User{
		Name:        args[0],
		Email:       args[1],
		CreditLimit: creditLimit,
		Spent:       0,
	}
	if !user.ValidateAndClean() {
		fmt.Println("Invalid user details")
	} else {
		err := us.userData.AddNewUser(user)
		if err != nil {
			fmt.Println("User creation failed ..!")
		}
		fmt.Printf("Created user : %+v\n", user)
		return
	}
	fmt.Println("User creation failed ..!")
}

// ChangeUserLimit is used to change the user limit
func (us *UserService) ChangeUserLimit(args []string) {

}

// Payback is used to payback the due
func (us *UserService) Payback(args []string) {

}

// GetAllUserDues is used to fetch the user dues
func (us *UserService) GetAllUserDues(args []string) {

}

// GetUsersWithCreditLimit is used to fetch the users with credit == limit
func (us *UserService) GetUsersWithCreditLimit(args []string) {

}
