package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

type Database struct {
    Conn *sql.DB
}

var DB *Database

func Connect() error {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    conn, err := sql.Open("postgres", dsn)
    if err != nil {
        return err
    }

    if err = conn.Ping(); err != nil {
        return err
    }

    DB = &Database{Conn: conn}
    log.Println("Database connected successfully")
    return nil
}

func (db *Database) Close() error {
    return db.Conn.Close()
}