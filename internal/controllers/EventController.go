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

// GetAllEventsHandler retourne tous les événements
// @Summary Récupère tous les événements
// @Description Récupère la liste complète des événements disponibles
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {array} models.Event
// @Failure 500 {string} string "Erreur lors de la récupération des événements"
// @Router /events [get]
func (c *EventController) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := c.Service.GetAllEvents()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GetEventsByResourceHandler retourne les événements d'une ressource donnée
// @Summary Récupère les événements d'une ressource
// @Description Récupère les événements associés à une ressource spécifique
// @Tags Events
// @Accept json
// @Produce json
// @Param resource_id path string true "ID de la ressource"
// @Success 200 {array} models.Event
// @Failure 400 {string} string "ID invalide"
// @Failure 500 {string} string "Erreur lors de la récupération des événements"
// @Router /events/resource/{resource_id} [get]
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
