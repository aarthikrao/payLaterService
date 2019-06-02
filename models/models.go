package models

import (
	"fmt"
	"time"
)

// More about sql tags at https://github.com/go-pg/pg/wiki/Model-Definition

// User is used to store user information
type User struct {
	// Auto Incrementing ID
	ID int

	// Name of user
	Name string `sql:",unique,notnull"`

	// Email of user
	Email string

	// Total amount spent
	Spent float32 `sql:"default:0"`

	// Max credit limit for user
	CreditLimit float32
}

// Merchant is used to store merchant information
type Merchant struct {
	// Auto incremented ID
	ID int

	// Name of the merchant
	Name string `sql:",unique,notnull"`

	// Interest rate for this merchant. Note: It might change
	InterestRate float32

	// Total amount credited into users account. defaulted to zero at creation
	TotalAmount float32 `sql:"default:0"`

	// Total Discount from the merchant
	TotalDiscount float32 `sql:"default:0"`
}

// Transaction is used to store transaction information
type Transaction struct {
	// Auto incremented ID
	ID int

	// Name of the user
	UserName string

	// Name of the merchant
	MerchantName string

	// Total transaction amount
	TotalAmount float32

	// Amount to be credited to merchant
	MerchantAmount float32

	// Our discount for this particular transaction
	OurDiscount float32

	// The interest rate for this transaction
	InterestRate float32

	// Time of the transaction
	Time time.Time `sql:"default:now()"`
}

// ValidateAndClean is used to valudate all the user related information
func (u *User) ValidateAndClean() bool {
	if u.Name == "" {
		fmt.Println("Invalid username")
		return false
	}
	if u.CreditLimit < 0 {
		fmt.Println("Invalid negative credit limit")
		return false
	}
	return true
}

// ValidateAndClean is used to valudate all the merchant related information
func (m *Merchant) ValidateAndClean() bool {
	if m.InterestRate <= 0 || m.InterestRate > 100 {
		fmt.Println("Interest rate must be within 0 to 100")
		return false
	}
	if m.Name == "" {
		fmt.Println("Invalid merchant name")
		return false
	}
	return true
}

// ValidateAndClean is used to valudate all the Transaction related information
func (t *Transaction) ValidateAndClean() bool {
	if t.InterestRate <= 0 || t.InterestRate > 100 {
		fmt.Println("Interest rate must be within 0 to 100")
		return false
	}
	if t.TotalAmount < 0 {
		fmt.Println("Total amount should not be negative")
		return false
	}
	return true
}
