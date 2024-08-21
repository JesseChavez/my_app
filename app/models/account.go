package models

import "time"

type Account struct {
	ID        int
	Email     string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
