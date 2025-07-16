package db

import (
	"fmt"
	"log"
	"strings"
)

type QueryBuilder struct {
	cols []string
	q    string
}

type ClosedQuery interface {
	GetQuery() string
}

type Closed struct {
	str string
}

func (c Closed) GetQuery() string {
	return c.str
}

func NewQueryBuilder(c []string) *QueryBuilder {
	log.Printf("[DEBUG] QueryBuilder - new query builder %v", c)
	return &QueryBuilder{cols: c}
}

func (qb *QueryBuilder) InsertInto(table string) *QueryBuilder {
	columns := strings.Join(qb.cols, ",")

	ph := make([]string, len(qb.cols))
	for i := range qb.cols {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	values := strings.Join(ph, ",")

	qb.q = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, values)

	return qb
}

func (qb *QueryBuilder) Select(cols []string) ClosedQuery {
	columns := strings.Join(cols, ",")

	if qb.q == "" {
		log.Printf("[ERROR] QueryBuilder - no query to select")
		return &Closed{str: ""}
	}
	q := fmt.Sprintf("%s RETURNING %s", qb.q, columns)
	return &Closed{str: q}
}

func (qb *QueryBuilder) Close() ClosedQuery {
	return &Closed{str: qb.q}
}
