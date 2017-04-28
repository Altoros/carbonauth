package user

import "testing"

func TestUser_CanQuery(t *testing.T) {
	t.Parallel()

	u := User{Globs: []string{"foo.*", "baz.*", "bam"}}
	for q, want := range map[string]bool{
		"*":       false,
		"bar.*":   false,
		"foo.*":   true,
		"foo.bar": true,
		"foo":     false,
		"baz.a":   true,
		"bam":     true,
		"bam1":    false,
		"ba":      false,
	} {
		if u.Can(q) != want {
			t.Errorf("Can(%q) = %t, want %t", q, u.Can(q), want)
		}
	}
}
