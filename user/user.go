package user

// User is user entity representation
type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Globs    []string `json:"globs"`
}

// Can returns true when s matches at least one of user's globs
func (u *User) Can(s string) bool {
	for _, g := range u.Globs {
		for i := 0; i < len(s) && i < len(g); i++ {
			if g[i] == '*' {
				return true
			}

			if g[i] != s[i] {
				break
			}

			if i == len(s) - 1 && i == len(g) - 1 {
				return true
			}
		}
	}
	return false
}
