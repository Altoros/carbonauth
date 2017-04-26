package user

// User is user entity representation
type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Globs    []string `json:"globs"`
}

// CanQuery returns true when q matches at least one of user's globs
func (u *User) CanQuery(q string) bool {
	for _, g := range u.Globs {
		for i := 0; i < len(q) && i < len(g); i++ {
			if g[i] == '*' {
				return true
			}

			if g[i] != q[i] {
				break
			}
		}
	}
	return false
}
