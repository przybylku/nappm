package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS processes (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        pid INTEGER NOT NULL
    );
    `
	if _, err = db.Exec(sqlStmt); err != nil {
		db.Close()
		log.Fatal(err)
	}
	log.Println("Table 'processes' is ready")
}

func CreateProcess(name string, pid int) (success bool, err error) {
	_, err = db.Exec(`INSERT INTO processes (name, pid) VALUES (?, ?)`, name, pid)
	if err != nil {
		return false, err
	}
	return true, nil

}
