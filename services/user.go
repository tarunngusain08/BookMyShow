package services

import (
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type User struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) GetUsers() ([]*models.User, error) {
	return u.repo.GetUsers()
}

func (u *User) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

func (u *User) GetUser(userID int) (*models.User, error) {
	return u.repo.GetUserByID(userID)
}

func (u *User) UpdateUser(userID int, updatedUser *models.User) error {
	return u.repo.UpdateUser(userID, updatedUser)
}

func (u *User) DeleteUser(userID int) error {
	return u.repo.DeleteUser(userID)
}
