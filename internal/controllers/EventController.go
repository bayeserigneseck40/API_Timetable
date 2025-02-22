package controllers

import (
	"API_timetable/internal/services"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// EventController gère les endpoints pour les événements.
type EventController struct {
	Service *services.EventService
}

// GetAllEventsHandler retourne tous les événements.
// GetEvents
// @Tags         events
// @Summary      Get a event.
// @Description  Get a event.
// @Param        id           	path      string  true  "Event UUID formatted ID"
// @Success      200            {object}  models.Event
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /events/ [get]
func (c *EventController) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := c.Service.GetAllEvents()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GetEventsByResourceHandler retourne les événements d'une resource donnée.
func (c *EventController) GetEventsByResourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resourceID, err := uuid.Parse(vars["resource_id"])
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	events, err := c.Service.GetEventsByResource(resourceID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
