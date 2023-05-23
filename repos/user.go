package repos

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/models"
)

const (
	getAllUsers = `SELECT id, first_name, last_name, email, phone, address, city FROM users`
	addUser     = `INSERT INTO users (first_name, last_name, email, phone, address, city) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	getUserByID = `SELECT id, first_name, last_name, email, phone, address, city FROM users WHERE id = $1`
	updateUser  = `UPDATE users SET first_name = $1, last_name = $2, email = $3, phone = $4, address = $5, city = $6 WHERE id = $7`
	deleteUser  = `DELETE FROM users WHERE id = $1`
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) GetUsers() ([]*models.User, error) {
	rows, err := ur.db.Query(getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Address, &user.City)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) CreateUser(user *models.User) error {
	var userID int
	err := ur.db.QueryRow(addUser, user.FirstName, user.LastName, user.Email, user.Phone, user.Address,
		user.City).Scan(&userID)
	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

func (ur *UserRepo) GetUserByID(userID int) (*models.User, error) {
	row := ur.db.QueryRow(getUserByID, userID)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Address, &user.City)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) UpdateUser(userID int, updatedValues *models.User) error {
	_, err := ur.db.Exec(updateUser, updatedValues.FirstName, updatedValues.LastName, updatedValues.Email,
		updatedValues.Phone, updatedValues.Address, updatedValues.City, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) DeleteUser(userID int) error {
	_, err := ur.db.Exec(deleteUser, userID)
	if err != nil {
		return err
	}
	return nil
}
