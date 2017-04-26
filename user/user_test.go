package user

import "testing"

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
