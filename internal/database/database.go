package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseImpl struct {
	Conn *pgxpool.Pool
}

func Conn() (DatabaseImpl, error) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
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
	return DatabaseImpl{
		Conn: dbpool,
	}, nil
}

func (d *DatabaseImpl) SelectAll(query string, args ...interface{}) ([][]interface{}, error) {
	if d.Conn == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	rows, err := d.Conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result [][]interface{}
	for rows.Next() {
		columns, err := rows.Values()
		if err != nil {
			return nil, err
		}
		result = append(result, columns)
	}

	log.Printf("Query: %s, Args: %v, Result: %v\n", query, args, result)
	return result, nil
}

func (d *DatabaseImpl) SelectFrom(query string, property string, from string, dest ...interface{}) error {
	log.Printf("Query: %s, Args: %s\n", query, property)
	if d.Conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	err := d.Conn.QueryRow(context.Background(), query, property, from).Scan(dest...)
	return err
}

func (d *DatabaseImpl) Insert(query string, args ...interface{}) error {
	log.Printf("Query: %s, Args: %v\n", query, args)
	if d.Conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	_, err := d.Conn.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
