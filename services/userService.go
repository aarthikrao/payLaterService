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
	// convert amount to float value
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
		return
	}
	err = us.userData.AddNewUser(user)
	if err != nil {
		fmt.Println("User creation failed. User already exsists")
		return
	}
	fmt.Printf("Created user : %+v\n", user)
	return

}

// Payback is used to payback the due
func (us *UserService) Payback(args []string) {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	// get amount in float value
	amt, err := utils.StrToFloat32(args[1])
	if err != nil || amt <= 0 {
		fmt.Println("Please enter valid amount")
		return
	}
	// Fetch user from db
	user, err := us.userData.GetUserByName(args[0])
	if err != nil {
		fmt.Println("Invalid Username")
		return
	}
	result, err := us.userData.PaybackUserDues(amt, user.Name)
	if err != nil {
		fmt.Println("Error in processing request")
		return
	}
	if result.RowsAffected() != 1 {
		fmt.Println("Request declined : amount greater than dues")
		return
	}
	fmt.Println("User:", user.Name, " remaining amount:", (user.Spent - amt))
	return
}

// GetUserDues is used to fetch the user dues
func (us *UserService) GetUserDues(args []string) {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	user, err := us.userData.GetUserByName(args[1])
	if err != nil {
		fmt.Println("Error in fetching user. User does not exist.")
		return
	}
	fmt.Println("User:", user.Name, " amount due :", user.Spent)
}

// GetUsersAtCreditLimit is used to fetch the users with credit == limit
func (us *UserService) GetUsersAtCreditLimit(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	users, err := us.userData.GetAllUsersWithQuery("spent = credit_limit")
	if err != nil {
		fmt.Println("Error in fetching users")
		return
	}
	if len(users) < 1 {
		fmt.Println("No users found at credit limit")
		return
	}
	fmt.Println("Users at credit limit:")
	for _, user := range users {
		fmt.Println("User:", user.Name, " Due:", user.Spent)
	}
}

// GetTotalUserDues is used to fetch the users with credit == limit
func (us *UserService) GetTotalUserDues(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect input, try 'help'")
		return
	}
	users, err := us.userData.GetAllUsersWithQuery("spent > 0")
	if err != nil {
		fmt.Println("Error in fetching users")
		return
	}
	if len(users) < 1 {
		fmt.Println("No users found with dues")
		return
	}
	fmt.Println("Users with dues:")
	for _, user := range users {
		fmt.Println("User:", user.Name, " Due:", user.Spent)
	}
}
