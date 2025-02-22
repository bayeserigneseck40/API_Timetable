package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	_ "modernc.org/sqlite"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Start       string    `json:"dtstart"`
	End         string    `json:"dtend"`
	UID         string    `json:"uid"`
	ResourceId  []string  `json:"resource_id"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "collections.db")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base SQLite : %v", err)
	}
}

func storeOrUpdateEvent(event Event) error {
	event.ID = uuid.New()
	resourceIdStr := strings.Join(event.ResourceId, ",")
	var existingID string
	err := db.QueryRow("SELECT uid FROM events WHERE uid = ?", event.UID).Scan(&existingID)
	if err == sql.ErrNoRows {
		_, err = db.Exec(`INSERT INTO events (id,summary,description, location, start, end, uid,resources_id) VALUES (?, ?, ?, ?, ?, ?, ?,?)`, event.ID, event.Summary, event.Description, event.Location, event.Start, event.End, event.UID, resourceIdStr)
		if err != nil {
			return err
		}
		fmt.Println("🆕 Nouvel événement inséré !")
	} else if err == nil {
		_, err = db.Exec(`UPDATE events SET summary=?, description=?, location=?, start=?, end=?, uid=?,resources_id=? WHERE id=?`,
			event.Summary, event.Description, event.Location, event.Start, event.End, event.UID, resourceIdStr, event.ID)
		if err != nil {
			return err
		}
		fmt.Println("✏️  Événement mis à jour !")
	} else {
		return err
	}
	return nil
}

func processMessage(msg jetstream.Msg) {
	log.Printf("Message reçu: %s", string(msg.Data())) // ✅ Ajout du log

	var event Event
	if err := json.Unmarshal(msg.Data(), &event); err != nil {
		log.Printf("Erreur de parsing JSON : %v", err)
		return
	}

	if err := storeOrUpdateEvent(event); err != nil {
		log.Printf("Erreur lors de l'insertion/mise à jour : %v", err)
	} else {
		msg.Ack()
	}
}

func main() {
	initDB()
	defer db.Close()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Erreur de connexion à NATS : %v", err)
	}
	defer nc.Close()

	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// ✅ Création du stream s'il n'existe pas
	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "USERS",
		Subjects: []string{"USERS.*"},
	})
	if err != nil && !errors.Is(err, jetstream.ErrStreamNameAlreadyInUse) {
		log.Fatalf("Erreur lors de la création du stream : %v", err)
	}

	stream, err := js.Stream(ctx, "USERS")
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du stream : %v", err)
	}

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable: "timetable_consumer",
		Name:    "timetable_consumer",
	})
	if err != nil {
		log.Fatalf("Erreur lors de la création du consumer : %v", err)
	}

	sub, err := consumer.Consume(processMessage)
	if err != nil {
		log.Fatalf("Erreur lors de la consommation des messages : %v", err)
	}

	<-sub.Closed()
}
