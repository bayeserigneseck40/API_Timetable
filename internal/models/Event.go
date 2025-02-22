package models

import (
	"github.com/google/uuid"
)

// Modèle d'un événement dans l'emploi du temps
type Event struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Summary     string    `json:"summary" db:"summary"`
	Description string    `json:"description" db:"description"`
	Location    string    `json:"location" db:"location"`
	Start       string    `json:"dtstart" db:"start"`
	End         string    `json:"dtend" db:"end"`
	UID         string    `json:"uid" db:"uid"`
	ResourceId  string    `json:"resource_id" db:"resources_id"`
}
