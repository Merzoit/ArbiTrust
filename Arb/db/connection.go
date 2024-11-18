package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DatabasePool *pgxpool.Pool

func InitDB() {
	var err error
	databaseUrl := "postgresql://gen_user:CVzh%5CBj%25%5C76o2Z@94.241.142.173:5432/default_db"

	DatabasePool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal("Отсутствует соединение", err)
	}

	fmt.Println("Соединение прошло успешно")
}
