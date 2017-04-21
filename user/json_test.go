package user

import (
	"encoding/json"
	"reflect"
	"regexp"
	"testing"
)

func TestUser_MarshalUnmarshalJSON(t *testing.T) {
	t.Parallel()

	u1 := &User{
		Username: "admin",
		Password: "secret",
		Regexps:  []*regexp.Regexp{regexp.MustCompile("^.*$")},
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

