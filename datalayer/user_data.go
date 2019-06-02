package datalayer

import (
	"fmt"

	conn "github.com/aarthikrao/payLaterService/connections"
	"github.com/aarthikrao/payLaterService/models"
	"github.com/go-pg/pg/orm"
)

// UserData will be used to fetched data
type UserData struct{}

// NewUserData is used to create new instance of NewUserData
func NewUserData() *UserData {
	return &UserData{}
}

// GetUserByName : Used to get user data by name
func (ud *UserData) GetUserByName(name string) (user models.User, err error) {
	// user := models.User{}
	err = conn.PGDB.Model(&user).Where("name = ?", name).Select()
	return
}

// AddNewUser is used to add new users to the system
func (ud *UserData) AddNewUser(user models.User) (err error) {
	err = conn.PGDB.Insert(&user)
	return
}

// UpdateUserByName is used to update user data
func (ud *UserData) UpdateUserByName(name string, user models.User) (err error) {
	_, err = conn.PGDB.Model(&user).Where("name = ?", name).UpdateNotNull(&user)
	return
}

// GetAllUsers fetches all user data, Note however that there is no row limit
func (ud *UserData) GetAllUsers() (users []models.User, err error) {
	err = conn.PGDB.Model(&users).Select()
	// TODO Add limit
	return
}

// GetAllUsersWithQuery fetches all user data, Note however that there is no row limit
func (ud *UserData) GetAllUsersWithQuery(query string) (users []models.User, err error) {
	err = conn.PGDB.Model(&users).Where(query).Select()
	// TODO Add limit
	return
}

// PaybackUserDues is used for user-payback
func (ud *UserData) PaybackUserDues(amount float32, name string) (result orm.Result, err error) {
	amountString := fmt.Sprintf("%f", amount)
	query := "update users set spent = spent - " + amountString + " where spent - " + amountString + " >= 0 and name = '" + name + "'"
	result, err = conn.PGDB.Exec(query)
	return
}
