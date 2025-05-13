package database

import (
	"context"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseImpl struct {
	Conn *pgxpool.Pool
}

func Conn(url string) (DatabaseImpl, error) {
	if url == "" {
		log.Fatalln("DATABASE_URL is not set. Closing the application.")
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	databaseConfig := dbpool.Config()
	log.Printf("User %s Connected Database: <%s>/<%s>\n",
		strings.Split(databaseConfig.ConnConfig.User, "@")[0],
		databaseConfig.ConnConfig.Host,
		databaseConfig.ConnConfig.Database,
	)
	return DatabaseImpl{Conn: dbpool}, nil
}
