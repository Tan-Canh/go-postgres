package i_user_repo

import (
	"database/sql"
	"fmt"
	"lesson_3/models"
	"lesson_3/repositories"
)

type IUserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repositories.UserRepo {
	return &IUserRepo{Db: db}
}

func (iUser *IUserRepo) Select() ([]models.User, error) {
	users := make([]models.User, 0)

	sqlSelect := "select * from users"
	rows, err := iUser.Db.Query(sqlSelect)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (iUser *IUserRepo) Insert(user models.User) error {
	sqlInsert := `insert into users (id, name, gender, email)
				values ($1,$2,$3,$4)`

	_, err := iUser.Db.Exec(sqlInsert, user.ID, user.Name, user.Gender, user.Email)
	if err != nil {
		return err
	}

	fmt.Println("Recorded Add", user)

	return nil
}
