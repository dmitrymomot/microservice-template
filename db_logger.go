package main

import (
	ctx "context"
	"log"

	"github.com/go-pg/pg/v9"
)

type dbLogger struct {
}

func newDBLogger() dbLogger {
	return dbLogger{}
}

func (d dbLogger) BeforeQuery(c ctx.Context, q *pg.QueryEvent) (ctx.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c ctx.Context, q *pg.QueryEvent) error {
	qs, err := q.FormattedQuery()
	if err != nil {
		log.Printf("query string error: %+v", err)
	}
	log.Printf("query string: %s", qs)
	return nil
}
