package repository

import (
	"github.com/jmoiron/sqlx"
	"microServices/bet_maker"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) CreateEvent(event bet_maker.Event) (int, error) {
	return 0, nil
}
