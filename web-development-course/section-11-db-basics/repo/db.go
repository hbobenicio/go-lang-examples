package repo

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"app/config"
)

// DB is the Database pool
var DB *sql.DB

// InitDB inicializa o pool de conexão com o banco de dados
func InitDB(cfg *config.Config) {
	DB = newDB(cfg)
	setupDB(DB)
}

// newDB cria uma nova instancia do pool de conexão com o banco
func newDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf(
		"dbname='%s' user='%s' password='%s' host='%s' sslmode='%s'",
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		"disable",
	)

	// initialize a new connection pool
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("connection error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db
}

// setupDB configura o pool de conexão com o banco
func setupDB(db *sql.DB) {
	// Sets the maximum limit of 10 concurrently open connections.
	// If 10 connections are already open and another new connection is needed,
	// then the application will be forced to wait
	// until one of the 10 open connections is freed up and becomes idle.
	// (Default value: unlimited)
	db.SetMaxOpenConns(10)

	// MaxIdleConns should always be less than or equal to MaxOpenConns.
	// Go enforces this and will automatically reduce MaxIdleConns if necessary.
	//
	// "There is no point in ever having any more idle connections than the maximum allowed open connections,
	// because if you could instantaneously grab all the allowed open connections,
	// the remain idle connections would always remain idle.
	// It's like having a bridge with four lanes,
	// but only ever allowing three vehicles to drive across it at once."
	// (Default value: 2)
	db.SetMaxIdleConns(10)

	db.SetConnMaxLifetime(1 * time.Hour)
}
