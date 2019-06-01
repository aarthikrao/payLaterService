package main

import (
	"github.com/aarthikrao/payLaterService/cli"
	"github.com/aarthikrao/payLaterService/config"
	conn "github.com/aarthikrao/payLaterService/connections"
)

func main() {
	// Initialise configs
	config.InitConfigs()

	// Initialse the postgres db connection
	conn.InitDatabase()
	defer conn.ShutDown()

	cli.INITServices()

	// user := &models.User{
	// 	Name:  "user1",
	// 	Email: "name@domain.com",
	// }
	// err := conn.PGDB.Insert(user)
	// if err != nil {
	// 	panic(err)
	// }
	// user2 := &models.User{
	// 	Email: "aarthik@yay.com",
	// 	ID:    1,
	// }
	// conn.PGDB.Update(user)
	// conn.PGDB.Model(user2).UpdateNotNull(user2)
	// conn.PGDB.Model(user2).Where("name = ?", "user1").UpdateNotNull(user2)
	// userData := datalayer.NewUserData()
	// userD, err := userData.GetUserByName("Aarthik")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Email:", userD.Email, " ID:", userD.ID)
	// x := "20"
	// name := "Aarthik"
	// res, err := conn.PGDB.Exec("update users set spent = spent + "+x+" where spent + "+x+" < credit_limit and name = ?", name)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Rows effected :", res.RowsAffected())
	// tx := services.NewTransactionService()
	// tx.TransferToMerchent([]string{"u1", "m1", "10"})
	cli.RunCli()
}

// https://github.com/abiosoft/ishell
