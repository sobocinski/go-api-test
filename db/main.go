package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// NewDatabase create database connection
func NewDatabase(ctx context.Context) (*pg.DB, error) {
	opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/db_name?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if err := db.Ping(ctx); err != nil {
		return db, err
	}

	//add logger
	db.AddQueryHook(dbLogger{})
	return db, nil
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	log.Println(string(query))
	return nil
}
