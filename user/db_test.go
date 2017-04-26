package user

import (
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	u := &User{
		Username: "user",
		Password: "secret",
		Globs:    []string{"foo.*"},
	}

	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	db, err := Open(databaseURL, "seasalt")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Clean()

	if err = db.Save(u); err != nil {
		t.Fatal(err)
	}

	u2, err := db.FindByUsernameAndPassword("user", "secret")
	if err != nil {
		t.Fatal(err)
	}

	if u.Username != u2.Username &&
		u.Password != u2.Password &&
		len(u.Globs) != len(u2.Globs) {
		t.Error("users are not equal")
	}

	if err = db.Delete("user"); err != nil {
		t.Fatal(err)
	}

	_, err = db.FindByUsernameAndPassword("user", "secret")
	if err != ErrInvalidCredentials {
		t.Errorf("err = %v, want %v", err, ErrInvalidCredentials)
	}
}
