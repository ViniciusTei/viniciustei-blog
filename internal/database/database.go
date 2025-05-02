package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseImpl struct {
	conn *pgxpool.Pool
}

func (d *DatabaseImpl) Conn() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()
	d = &DatabaseImpl{
		conn: dbpool,
	}
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

	return result, nil
}

func (d *DatabaseImpl) SelectOne(query string, args ...interface{}) error {
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
	if d.conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	_, err := d.conn.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
