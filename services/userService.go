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
// Args : name email credit-limit
func (us *UserService) CreateUser(args []string) bool {
	if len(args) != 3 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	// convert amount to float value
	creditLimit, err := utils.StrToFloat32(args[2])
	if err != nil {
		fmt.Println("Please provide a number for credit limit")
		return false
	}
	user := models.User{
		Name:        args[0],
		Email:       args[1],
		CreditLimit: creditLimit,
		Spent:       0,
	}
	if !user.ValidateAndClean() {
		fmt.Println("Invalid user details")
		return false
	}
	err = us.userData.AddNewUser(user)
	if err != nil {
		fmt.Println("User creation failed. User already exsists")
		return false
	}
	fmt.Printf("Created user : %+v\n", user)
	return true

}

// Payback is used to payback the due
// Args : user-name amount
func (us *UserService) Payback(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	// get amount in float value
	amt, err := utils.StrToFloat32(args[1])
	if err != nil || amt <= 0 {
		fmt.Println("Please enter valid amount")
		return false
	}
	// Fetch user from db
	user, err := us.userData.GetUserByName(args[0])
	if err != nil {
		fmt.Println("Invalid Username")
		return false
	}
	result, err := us.userData.PaybackUserDues(amt, user.Name)
	if err != nil {
		fmt.Println("Error in processing request")
		return false
	}
	if result.RowsAffected() != 1 {
		fmt.Println("Request declined : amount greater than dues")
		return false
	}
	fmt.Println("User:", user.Name, " remaining amount:", (user.Spent - amt))
	return true
}

// GetUserDues is used to fetch the user dues
// Args : dues user-name
func (us *UserService) GetUserDues(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	user, err := us.userData.GetUserByName(args[1])
	if err != nil {
		fmt.Println("Error in fetching user. User does not exist.")
		return false
	}
	fmt.Println("User:", user.Name, " amount due :", user.Spent)
	return true
}

// GetUsersAtCreditLimit is used to fetch the users with credit == limit
// Args : users-at-credit-limit
func (us *UserService) GetUsersAtCreditLimit(args []string) bool {
	if len(args) != 1 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	users, err := us.userData.GetAllUsersWithQuery("spent = credit_limit")
	if err != nil {
		fmt.Println("Error in fetching users")
		return false
	}
	if len(users) < 1 {
		fmt.Println("No users found at credit limit")
		return false
	}
	fmt.Println("Users at credit limit:")
	for _, user := range users {
		fmt.Println("User:", user.Name, " Due:", user.Spent)
	}
	return true
}

// GetTotalUserDues is used to fetch the users with credit == limit
// Args : total-duess
func (us *UserService) GetTotalUserDues(args []string) bool {
	if len(args) != 1 {
		fmt.Println("Incorrect input, try 'help'")
		return false
	}
	users, err := us.userData.GetAllUsersWithQuery("spent > 0")
	if err != nil {
		fmt.Println("Error in fetching users")
		return false
	}
	if len(users) < 1 {
		fmt.Println("No users found with dues")
		return false
	}
	dues := float32(0.0)
	for _, user := range users {
		fmt.Println("User:", user.Name, " Due:", user.Spent)
		dues = dues + user.Spent
	}
	fmt.Println("Total:", dues)
	return true
}
