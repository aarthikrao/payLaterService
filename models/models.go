package models

import "time"

// More about sql tags at https://github.com/go-pg/pg/wiki/Model-Definition

// User is used to store user information
type User struct {
	ID          int
	Name        string `sql:",unique,notnull"`
	Email       string
	Spent       float32 `sql:"default:0"`
	CreditLimit float32
}

// Merchant is used to store merchant information
type Merchant struct {
	ID            int
	Name          string `sql:",unique,notnull"`
	InterestRate  float32
	TotalAmount   float32 `sql:"default:0"`
	TotalDiscount float32 `sql:"default:0"`
}

// Transaction is used to store transaction information
type Transaction struct {
	ID             int
	UserName       string
	MerchantName   string
	TotalAmount    float32
	MerchantAmount float32
	OurDiscount    float32
	InterestRate   float32
	Time           time.Time `sql:"default:now()"`
}

// ValidateAndClean is used to valudate all the user related information
func (u *User) ValidateAndClean() bool {

	return true
}

// ValidateAndClean is used to valudate all the merchant related information
func (m *Merchant) ValidateAndClean() bool {

	return true
}

// ValidateAndClean is used to valudate all the Transaction related information
func (t *Transaction) ValidateAndClean() bool {

	return true
}
