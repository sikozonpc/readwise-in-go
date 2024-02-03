package main

import "time"

type RawExtractBook struct {
	ASIN       string                `json:"asin"`
	Title      string                `json:"title"`
	Authors    string                `json:"authors"`
	Highlights []RawExtractHighlight `json:"highlights"`
}

type RawExtractHighlight struct {
	Text     string `json:"text"`
	Location struct {
		Value int    `json:"value"`
		URL   string `json:"url"`
	} `json:"location"`
	IsNoteOnly bool   `json:"isNoteOnly"`
	Note       string `json:"note"`
}

type Book struct {
	ISBN      string    `json:"isbn"`
	Title     string    `json:"title"`
	Authors   string    `json:"authors"`
	CreatedAt time.Time `json:"created_at"`
}

type Highlight struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Location string `json:"location"`
	Note     string `json:"note"`
	UserID   int    `json:"userId"`
	BookID   string    `json:"bookId"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"created_at"`
}

type DailyInsight struct {
	Text        string
	Note        string
	BookAuthors string
	BookTitle   string
}
