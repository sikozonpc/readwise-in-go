package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	// initialize the tables
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createBooksTable(); err != nil {
		return nil, err
	}

	if err := s.createHighlightsTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLStorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err

}

func (s *MySQLStorage) createBooksTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			isbn VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			authors VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (isbn),
			UNIQUE KEY (isbn)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}

func (s *MySQLStorage) createHighlightsTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS highlights (
			id INT NOT NULL AUTO_INCREMENT,
			text TEXT,
			location VARCHAR(255) NOT NULL,
			note TEXT,
			userId INT UNSIGNED NOT NULL,
			bookId VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (userId) REFERENCES users(id),
			FOREIGN KEY (bookId) REFERENCES books(isbn)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}