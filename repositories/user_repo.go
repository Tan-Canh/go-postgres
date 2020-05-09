package repositories

import "lesson_3/models"

type UserRepo interface {
	Select() ([]models.User, error)
	Insert(models.User) error
}
