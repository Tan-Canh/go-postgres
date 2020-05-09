package main

import (
	"fmt"
	"lesson_3/drivers"
	"lesson_3/models"
	"lesson_3/repositories/i_user_repo"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "ntc"
	password = "1"
	dbname   = "db_lession3"
)

func main() {
	db := drivers.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	userRepo := i_user_repo.NewUserRepo(db.SQL)

	user1 := models.User{
		ID:     1,
		Name:   "Canh",
		Gender: "20",
		Email:  "ng@gmail.com",
	}

	user2 := models.User{
		ID:     2,
		Name:   "Kiet",
		Gender: "20",
		Email:  "k@gmail.com",
	}

	err = userRepo.Insert(user1)
	fmt.Println(err)
	userRepo.Insert(user2)

	rows, _ := userRepo.Select()

	for i := range rows {
		fmt.Println(rows[i])
	}

}
