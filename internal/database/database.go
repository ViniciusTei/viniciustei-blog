package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseImpl struct {
	conn *pgxpool.Pool
}

func Conn() (*DatabaseImpl, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, fmt.Errorf("Missing DATABASE_URL env")
	}

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	databaseConfig := dbpool.Config()
	log.Printf("User %s Connected Database: <%s>/<%s>\n",
		databaseConfig.ConnConfig.User,
		databaseConfig.ConnConfig.Host,
		databaseConfig.ConnConfig.Database,
	)
	return &DatabaseImpl{
		conn: dbpool,
	}, nil
}

func (d *DatabaseImpl) SelectAll(query string, args ...interface{}) ([][]interface{}, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	rows, err := d.conn.Query(context.Background(), query, args...)
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

func (d *DatabaseImpl) SelectOne(query string, args ...interface{}) error {
	log.Printf("Query: %s, Args: %v\n", query, args)
	if d.conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	row := d.conn.QueryRow(context.Background(), query, args...)
	var columns []interface{}
	err := row.Scan(&columns)
	if err != nil {
		return err
	}

	return nil
}

func (d *DatabaseImpl) Insert(query string, args ...interface{}) error {
	log.Printf("Query: %s, Args: %v\n", query, args)
	if d.conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	_, err := d.conn.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
