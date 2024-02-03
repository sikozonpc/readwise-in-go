package main

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

type Storage interface {
	CreateBook(Book) error
	CreateHighlights([]Highlight) error
	GetBookByISBN(string) (*Book, error)
	GetRandomHighlights(limit, userId int) ([]*Highlight, error)
	GetUsers() ([]*User, error)
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateBook(b Book) error {
	_, err := s.db.Exec(`
		INSERT INTO books (isbn, title, authors) 
		VALUES (?, ?, ?)
	`, b.ISBN, b.Title, b.Authors)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateHighlights(hs []Highlight) error {
	values := []interface{}{}

	query := "INSERT INTO highlights (text, location, note, userId, bookId) VALUES "
	for _, h := range hs {
		query += "(?, ?, ?, ?, ?),"
		values = append(values, h.Text, h.Location, h.Note, h.UserID, h.BookID)
	}

	query = query[:len(query)-1]

	_, err := s.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetBookByISBN(isbn string) (*Book, error) {
	rows, err := s.db.Query(`
		SELECT * FROM books WHERE isbn = ?
	`, isbn)
	if err != nil {
		return nil, err
	}

	book := new(Book)
	for rows.Next() {
		if err := rows.Scan(&book.ISBN, &book.Title, &book.Authors, &book.CreatedAt); err != nil {
			return nil, err
		}
	}

	if book.ISBN == "" {
		return nil, fmt.Errorf("book not found")
	}

	return book, nil
}

func (s *Store) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	for rows.Next() {
		u := new(User)

		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.FirstName,
			&u.LastName,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (s *Store) GetRandomHighlights(n, userID int) ([]*Highlight, error) {
	rows, err := s.db.Query("SELECT * FROM highlights WHERE userId = ? ORDER BY RAND() LIMIT ?", userID, n)
	if err != nil {
		return nil, err
	}

	var highlights []*Highlight
	for rows.Next() {
		h := new(Highlight)

		if err := rows.Scan(
			&h.ID,
			&h.Text,
			&h.Location,
			&h.Note,
			&h.UserID,
			&h.BookID,
			&h.CreatedAt,
		); err != nil {
			return nil, err
		}

		highlights = append(highlights, h)
	}

	return highlights, nil
}
