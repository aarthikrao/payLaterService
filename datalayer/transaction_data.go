package datalayer

import (
	"fmt"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"

	conn "github.com/aarthikrao/payLaterService/connections"
	"github.com/aarthikrao/payLaterService/models"
)

// TransactionData is struct used for all transactions
type TransactionData struct{}

// NewTransactionData is used for creating new instance of transaction data
func NewTransactionData() *TransactionData {
	return &TransactionData{}
}

// RunTransaction is used to transfer money from user account to merchant account
func (ts *TransactionData) RunTransaction(txn models.Transaction) (err error) {
	tx, err := conn.PGDB.Begin()
	if err != nil {
		fmt.Println("Error in begin transaction")
		return
	}

	// Deduct amount from user
	userRes, err := deductAmountFromUser(tx, txn.TotalAmount, txn.UserName)
	if err != nil {
		fmt.Println("Error in decrementing value from user")
		tx.Rollback()
		return
	}
	if userRes.RowsAffected() != 1 {
		fmt.Println("Transaction rejected due to insufficient balance")
		tx.Rollback()
		return
	}
	// Adding amount to merchant
	merchantRes, err := addAmountToMerchant(tx, txn.MerchantName, txn.MerchantAmount, txn.OurDiscount)
	if err != nil {
		fmt.Println("Error in decrementing value from user")
		tx.Rollback()
		return
	}
	if merchantRes.RowsAffected() != 1 {
		fmt.Println("Transaction rejected: Amount not credited to merchant")
		tx.Rollback()
		return
	}
	tx.Insert(&txn)
	tx.Commit()
	fmt.Printf("Transaction successful : %+v", txn)
	return
}

// DeductAmountFromUser is used for transactions
func deductAmountFromUser(tx *pg.Tx, amount float32, name string) (result orm.Result, err error) {
	amountString := fmt.Sprintf("%f", amount)
	query := "update users set spent = spent + " + amountString + " where spent + " + amountString + " <= credit_limit and name = '" + name + "'"
	result, err = tx.Exec(query)
	return
}

// AddAmountToMerchant is used to add money in merchants account during transaction
func addAmountToMerchant(tx *pg.Tx, merchantName string, merchantAmount float32, ourInterest float32) (result orm.Result, err error) {
	merchantAmountString := fmt.Sprintf("%f", merchantAmount)
	ourInterestString := fmt.Sprintf("%f", ourInterest)
	query := "update merchants set total_amount = total_amount + " + merchantAmountString + " , total_discount = total_discount + " + ourInterestString + " where name = '" + merchantName + "'"
	result, err = tx.Exec(query)
	return
}
