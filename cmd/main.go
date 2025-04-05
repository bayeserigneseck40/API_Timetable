package main

import (
	"API_timetable/internal/controllers"
	"API_timetable/internal/helpers"
	"API_timetable/internal/repositories"
	"API_timetable/internal/services"
	"github.com/gorilla/mux"
	"log"
	"net/http" // Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
)

// @title API Timetable & Config
// @version 1.0
// @description Documentation de l'API Timetable et Config.
// @host localhost:8080
// @BasePath /
func main() {
	db, _ := helpers.InitDB()
	defer db.Close()

	// Initialisation des composants
	eventRepo := &repositories.EventRepository{DB: db}
	eventService := &services.EventService{Repo: eventRepo}
	eventController := &controllers.EventController{Service: eventService}

	// Définition du routeur
	router := mux.NewRouter()

	// Routes pour les événements
	router.HandleFunc("/events", eventController.GetAllEventsHandler).Methods("GET")
	router.HandleFunc("/events/{resource_id}", eventController.GetEventsByResourceHandler).Methods("GET")

	// Lancement du serveur
	log.Println("🚀 Serveur Timetable en écoute sur http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
