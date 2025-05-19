package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func NewDb() (*sql.DB, func(), error) {
	err := godotenv.Load("/home/andhikadanger/cqrs/.env")
	if err != nil {
		NewErr("../logs/db", logrus.ErrorLevel, err)
		return nil, nil, err
	}
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, db_host, db_port, db_name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		NewErr("../logs/db", logrus.ErrorLevel, err)
		return nil, nil, err
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
