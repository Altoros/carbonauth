package user

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUser_CanQuery(t *testing.T) {
	t.Parallel()

	u := User{Globs: []string{"foo.*", "baz.*"}}
	for q, want := range map[string]bool{
		"*":       false,
		"bar.*":   false,
		"foo.*":   true,
		"foo.bar": true,
		"foo":     false,
		"baz.a":   true,
	} {
		if u.CanQuery(q) != want {
			t.Errorf("CanQuery(%q) = %t, want %t", q, u.CanQuery(q), want)
		}
	}
}

func TestUser_MarshalUnmarshalJSON(t *testing.T) {
	t.Parallel()

	u1 := &User{
		Username: "admin",
		Password: "secret",
		Globs:    []string{"foo.*", "bar.*"},
	}

	b, err := json.Marshal(u1)
	if err != nil {
		t.Fatal(err)
	}

	u2 := &User{}
	if err := json.Unmarshal(b, u2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(u1, u2) {
		t.Error("u != Unmarshal(Marshal*u))")
	}
}
