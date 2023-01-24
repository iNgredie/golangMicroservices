package repository

import (
	"github.com/jmoiron/sqlx"
	"microServices/bet_maker"
)

type Event interface {
	CreateEvent(event bet_maker.Event) (int, error)
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Event: NewEventPostgres(db),
	}
}
