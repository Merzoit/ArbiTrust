package repositories

import (
	"arb/constants"
	"arb/db"
	"arb/structures"
	"context"
	"errors"
	"log"
	"time"
)

type UserRepository interface {
	AddUser(user *structures.User) error
	GetUserById(id uint) (*structures.User, error)
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
	log.Println("DB: " + constants.CallDBAddingUser)
	err := db.DatabasePool.QueryRow(
		context.Background(), query, user.TID, user.Name,
		user.RegDate, user.Rating,
	).Scan(&user.ID)

	if err != nil {
		log.Printf("DB: "+constants.LogErrorAddingUser, err)
		return err
	}

	log.Printf("DB: "+constants.LogUserCreateSuccessfully, user.ID)
	return nil
}

func (repo *PgUserRepository) GetUserById(id uint) (*structures.User, error) {
	user := &structures.User{}
	query := "SELECT id, tid, name, reg_date, rating FROM users WHERE id=$1"

	log.Println("DB: " + constants.CallGetUserByID)
	err := db.DatabasePool.QueryRow(context.Background(), query, id).Scan(
		&user.ID, &user.TID, &user.Name, &user.RegDate, &user.Rating,
	)
	if err != nil {
		log.Printf("DB: "+constants.LogErrorFetchingUser, err)
		return nil, errors.New(constants.ErrUserNotFound)
	}

	log.Printf("DB: "+constants.LogUserCreateSuccessfully, err)
	return user, nil
}

func (repo *PgUserRepository) DeleteUser(id uint) error {
	query := "DELETE FROM users WHERE id=$1"

	log.Println("DB: " + constants.CallDeleteUser)
	_, err := db.DatabasePool.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf(constants.LogErrorDeletingUser, err)
		return errors.New(constants.ErrUserNotFound)
	}

	log.Printf(constants.LogUserDeleteSuccessfully, id)
	return nil
}

func (repo *PgUserRepository) UpdateUser(user *structures.User) error {
	query := "UPDATE users SET tid=$1, name=$2, rating=$3, reg_date=$4 WHERE id=$5"

	log.Println("DB: " + constants.CallUpdateUser)
	_, err := db.DatabasePool.Exec(
		context.Background(), query, user.TID, user.Name,
		user.Rating, user.RegDate, user.ID,
	)

	if err != nil {
		log.Printf("DB: "+constants.LogErrorUpdatingUser, err)
		return err
	}

	log.Printf("DB: "+constants.LogUserUpdateSuccessfully, user.ID)
	return nil
}
