package user

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
	salt string
}

func Open(url, salt string) (*DB, error) {
	chunks := strings.Split(url, "://")
	if len(chunks) != 2 {
		return nil, errors.New("malformed database url")
	}

	switch chunks[0] {
	case "mysql":
	case "postgres":
	default:
		return nil, fmt.Errorf("%q scheme is not supported", chunks[0])
	}

	// postgres needs whole url address
	if chunks[0] == "postgres" {
		chunks[1] = url
	}

	db, err := sql.Open(chunks[0], chunks[1])
	if err != nil {
		return nil, err
	}

	// create tables
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		username VARCHAR(32) NOT NULL PRIMARY KEY,
		password VARCHAR(32) NOT NULL,
		regexps  TEXT
	)`)

	if err != nil {
		// close db since we are returning an error here
		db.Close()
		return nil, err
	}

	return &DB{DB: db, salt: salt}, nil
}

// placeholderRegexp looks for postgres $ placeholders
var placeholderRegexp = regexp.MustCompile("\\$\\d+")

// exec executes a query without returning rows
func (db *DB) exec(q string, v ...interface{}) (sql.Result, error) {
	return db.Exec(db.prep(q), v...)
}

// query executes a query that returns rows
func (db *DB) query(q string, v ...interface{}) (*sql.Rows, error) {
	return db.Query(db.prep(q), v...)
}

// prep is needed for mysql compatibility
// it replaces postgres $ placeholders with ?
func (db *DB) prep(q string) string {
	if _, ok := db.Driver().(*mysql.MySQLDriver); ok {
		q = placeholderRegexp.ReplaceAllString(q, "?")
	}
	return q
}

// Clean closes db connection and drops all tables
func (db *DB) Clean() error {
	if err := db.Close(); err != nil {
		return err
	}
	_, err := db.Exec("DROP TABLE users")
	return err
}
