package repositories

import (
	"API_timetable/internal/models"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

// EventRepository gère l'accès aux événements dans la base de données.
type EventRepository struct {
	DB *sql.DB
}

// GetAllEvents récupère tous les événements.
func (repo *EventRepository) GetAllEvents() ([]models.Event, error) {
	rows, err := repo.DB.Query("SELECT id, summary,description, location, start, end, uid,resources_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Summary, &event.Description, &event.Location, &event.Start, &event.End, &event.UID, &event.ResourceId)
		if err != nil {
			log.Println("Erreur lors du scan des événements :", err)
			continue
		}
		events = append(events, event)
	}
	return events, nil
}

// GetEventsByResource récupère tous les événements liés à une resource donnée.
func (repo *EventRepository) GetEventsByResource(resourceID uuid.UUID) ([]models.Event, error) {
	rows, err := repo.DB.Query("SELECT id, summary,description, location, start,resources_id end FROM events WHERE resource_id = ?", resourceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.ID, &event.Summary, &event.Description, &event.Location, &event.Start, &event.End, &event.UID, &event.ResourceId)
		if err != nil {
			log.Println("Erreur lors du scan des événements :", err)
			continue
		}
		events = append(events, event)
	}
	return events, nil
}

// Récupère une ressource par son ID
func (r *EventRepository) GetByID(id uuid.UUID) (*models.Event, error) {
	query := "SELECT id, summary,description, location, start,resources_id end FROM events WHERE resource_id = ?"
	row := r.DB.QueryRow(query, id)

	event := &models.Event{}
	err := row.Scan(&event.ID, &event.Summary, &event.Description, &event.Location, &event.Start, &event.End, &event.UID, &event.ResourceId)
	if err != nil {
		return nil, err
	}
	return event, nil
}
