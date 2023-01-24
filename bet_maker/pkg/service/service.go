package service

import (
	"microServices/bet_maker"
	"microServices/bet_maker/pkg/repository"
)

type Event interface {
	CreateEvent(event bet_maker.Event)
}

type Service struct {
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewEventService(repos.Event),
	}
}
