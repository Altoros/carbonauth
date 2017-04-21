package user

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"
)

// User is user entity representation
type User struct {
	Username string
	Password string
	Regexps  []*regexp.Regexp
}

// CanRead checks whether user has access to the named key
func (u *User) CanRead(s string) bool {
	for _, m := range u.Regexps {
		if m.MatchString(s) {
			return true
		}
	}
	return false
}

// UserSave creates or updates existing user
func (db *DB) Save(u *User) error {
	rows, err := db.query(`SELECT 1 FROM users WHERE username = $1`, u.Username)
	if err != nil {
		return err
	}

	// hash password
	u.Password = hash(u.Password, db.salt)

	regexps, err := dumpRegexps(u.Regexps)
	if err != nil {
		return err
	}

	// update existing user if it exists
	if rows.Next() {
		_, err = db.exec(`UPDATE users SET password = $1, regexps = $2 WHERE username = $3`,
			u.Password, string(regexps), u.Username)
		return err
	}

	// create a new record
	_, err = db.exec(`INSERT INTO users (username, password, regexps) VALUES ($1, $2, $3)`,
		u.Username, u.Password, string(regexps))
	return err
}

// ErrInvalidCredentials returned when user cannot be found
var ErrInvalidCredentials = errors.New("invalid username or password")

// UserFind
func (db *DB) Find(username, password string) (*User, error) {
	rows, err := db.query(`
		SELECT username, password, regexps
		FROM users
		WHERE Username = $1
			AND Password = $2
		LIMIT 1
	`, username, hash(password, db.salt))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := &User{}
	b := []byte{}
	for rows.Next() {
		if err = rows.Scan(&u.Username, &u.Password, &b); err != nil {
			return nil, err
		}
	}

	// user is not found
	if u.Username == "" {
		return nil, ErrInvalidCredentials
	}

	// convert bytes to regexps
	regexps, err := loadRegexps(b)
	if err != nil {
		return nil, err
	}
	u.Regexps = regexps

	return u, nil
}

// UserDelete deletes the named user
func (db *DB) Delete(username string) error {
	_, err := db.exec(`DELETE FROM users WHERE username = $1`, username)
	return err
}

// hash returns base64 encoded hash of the provided string plus salt
func hash(s, salt string) string {
	h := sha256.New()
	h.Write([]byte(s))
	h.Write([]byte(salt))

	return base64.StdEncoding.EncodeToString(h.Sum([]byte{}))[:32]
}
