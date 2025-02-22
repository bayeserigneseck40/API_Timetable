package helpers

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// ConnectDB initialise la base de données SQLite
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "collections.db")
	if err != nil {
		return nil, err
	}

	// Création des tables si elles n'existent pas
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// createTables crée les tables nécessaires
func createTables(db *sql.DB) error {
	query := `
CREATE TABLE IF NOT EXISTS events (
    id TEXT PRIMARY KEY,
    summary TEXT,
	Description,
	Location TEXT,
	Start  DATETIME,
	End    DATETIME,
	UID    TEXT,
	Resources_id TEXT
);

	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Erreur lors de la création des tables : %v", err)
		return err
	}

	fmt.Println("✅ Base de données et tables créées avec succès !")
	return nil
}
