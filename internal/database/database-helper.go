package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitializeDatabase() {
	db, err := sql.Open("sqlite", "./scaler_data.db")
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}

	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS scaler_cluster (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        ip TEXT,
        status TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		last_changed DATETIME
    );`

	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("Error %s when creating table\n", err)
	}
	log.Printf("Database initialized successfully")

}
