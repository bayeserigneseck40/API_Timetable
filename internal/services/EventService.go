package services

import (
	"API_timetable/internal/models"
	"API_timetable/internal/repositories"
	"github.com/google/uuid"
)

// EventService gère la logique métier des événements.
type EventService struct {
	Repo *repositories.EventRepository
}

// GetAllEvents récupère tous les événements.
func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.Repo.GetAllEvents()
}

// GetEventsByResource récupère tous les événements d'une resource spécifique.
func (s *EventService) GetEventsByResource(resourceID uuid.UUID) ([]models.Event, error) {
	return s.Repo.GetEventsByResource(resourceID)
}

// GetEventsByResource récupère tous les événements d'une resource spécifique.
func (s *EventService) GetEventsById(id uuid.UUID) (*models.Event, error) {
	return s.Repo.GetByID(id)
}
