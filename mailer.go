package main

type Mailer interface {
	SendInsights(ins []*DailyInsight, u *User) error
}

