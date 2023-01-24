package service

import (
	"microServices/bet_maker"
	"microServices/bet_maker/pkg/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event bet_maker.Event) (int, error) {
	return s.repo.CreateEvent(event)
}
