package user

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cznic/sqlite"
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
	case "sqlite3":
	default:
		return nil, fmt.Errorf("%q scheme is not supported", chunks[0])
	}

	switch chunks[0] {
	case "postgres":
		chunks[1] = url
	case "sqlite3":
		chunks[0] = "sqlite"
	}

	db, err := sql.Open(chunks[0], chunks[1])
	if err != nil {
		return nil, err
	}

	// create tables
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		username VARCHAR(32) NOT NULL PRIMARY KEY,
		password VARCHAR(32) NOT NULL,
		globs    TEXT
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
func (db *DB) exec(tx *sql.Tx, q string, v ...interface{}) (sql.Result, error) {
	if tx == nil {
		return db.Exec(db.prep(q), v...)
	}
	return tx.Exec(db.prep(q), v...)
}

// query executes a query that returns rows
func (db *DB) query(tx *sql.Tx, q string, v ...interface{}) (*sql.Rows, error) {
	if tx == nil {
		return db.Query(db.prep(q), v...)
	}
	return tx.Query(db.prep(q), v...)
}

// prep is needed for mysql compatibility
// it replaces postgres $ placeholders with ?
func (db *DB) prep(q string) string {
	switch db.Driver().(type) {
	case *mysql.MySQLDriver, *sqlite.Driver:
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

// UserSave creates or updates existing user
func (db *DB) Save(u *User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()

	rows, err := db.query(tx, `SELECT 1 FROM users WHERE username = $1`, u.Username)
	if err != nil {
		return err
	}
	defer rows.Close()

	// hash password
	u.Password = hash(u.Password, db.salt)

	globs, err := marshalStringSlice(u.Globs)
	if err != nil {
		return err
	}

	// update existing user if it exists
	if rows.Next() {
		_, err = db.exec(tx, `UPDATE users SET password = $1, globs = $2 WHERE username = $3`,
			u.Password, globs, u.Username)
		return err
	}

	// create a new record
	_, err = db.exec(tx, `INSERT INTO users (username, password, globs) VALUES ($1, $2, $3)`,
		u.Username, u.Password, globs)
	return err
}

// ErrNotFound returned when user cannot be found
var ErrNotFound = errors.New("user not found")

func (db *DB) FindByUsername(username string) (*User, error) {
	return nil, errors.New("not implemented")
}

// FindByUsernameAndPassword
func (db *DB) FindByUsernameAndPassword(username, password string) (*User, error) {
	rows, err := db.query(nil, `
		SELECT username, password, globs
		FROM users
		WHERE username = $1
			AND password = $2
		LIMIT 1
	`, username, hash(password, db.salt))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u User
	var b []byte
	for rows.Next() {
		if err = rows.Scan(&u.Username, &u.Password, &b); err != nil {
			return nil, err
		}
	}

	// user is not found
	if u.Username == "" {
		return nil, ErrNotFound
	}

	globs, err := unmarshalStringSlice(b)
	if err != nil {
		return nil, err
	}
	u.Globs = globs

	return &u, nil
}

// UserDelete deletes the named user
func (db *DB) Delete(username string) error {
	_, err := db.exec(nil, `DELETE FROM users WHERE username = $1`, username)
	return err
}

// hash returns base64 encoded hash of the provided string plus salt
func hash(s, salt string) string {
	h := sha256.New()
	h.Write([]byte(s))
	h.Write([]byte(salt))

	// base16 is 64 bytes long we need only 32
	return fmt.Sprintf("%x", h.Sum(nil))[:32]
}

// marshalStringSlice dumps a slice of strings into byte encoding
func marshalStringSlice(ss []string) ([]byte, error) {
	return json.Marshal(&ss)
}

// unmarshalStringSlice converts byte encoding into a slice of strings
func unmarshalStringSlice(b []byte) ([]string, error) {
	var ss []string
	if err := json.Unmarshal(b, &ss); err != nil {
		return nil, err
	}
	return ss, nil
}
