package repositories

import (
	"arb/db"
	"arb/structures"
	"context"
	"errors"
	"time"
)

type UserRepository interface {
	AddUser(user *structures.User) error
	GetUserById(i uint) (*structures.User, error)
	DeleteUser(id uint) error
	UpdateUser(user *structures.User) error
}

type PgUserRepository struct{}

func NewPgUserRepository() UserRepository {
	return &PgUserRepository{}
}

func (repo *PgUserRepository) AddUser(user *structures.User) error {
	user.RegDate = time.Now()
	user.Rating = 0.0

	query := `
        INSERT INTO users (tid, name, reg_date, rating)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	err := db.DatabasePool.QueryRow(context.Background(), query, user.TID, user.Name, user.RegDate, user.Rating).Scan(&user.ID)
	return err
}

func (repo *PgUserRepository) GetUserById(id uint) (*structures.User, error) {
	user := &structures.User{}
	query := "SELECT id, tid, name, reg_date, rating FROM users WHERE id=$1"
	err := db.DatabasePool.QueryRow(context.Background(), query, id).Scan(
		&user.ID, &user.TID, &user.Name, &user.RegDate, &user.Rating,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repo *PgUserRepository) DeleteUser(id uint) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := db.DatabasePool.Exec(context.Background(), query, id)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (repo *PgUserRepository) UpdateUser(user *structures.User) error {
	query := "UPDATE users SET tid=$1, name=$2, rating=$3, reg_date=$4 WHERE id=$5"
	_, err := db.DatabasePool.Exec(context.Background(), query, user.TID, user.Name, user.Rating, user.RegDate, user.ID)
	return err
}
