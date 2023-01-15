package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(URI string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", URI)
	// db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(db *sqlx.DB) error {
	query := fmt.Sprintf(`CREATE TABLE %s
	(
		id serial not null unique,
		title varchar(200) not null unique,
		description varchar(1000),
		url1 varchar(255),
		url2 varchar(255),
		url3 varchar(255),
		price money,
		date_create timestamp default CURRENT_TIMESTAMP
	);`, tableProduct)
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
