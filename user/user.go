package user

import "encoding/json"

// User is user entity representation
type User struct {
	Username string
	Password string
	Globs    []string
}

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

type userJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Globs    []byte `json:"globs"`
}

// MarshalJSON returns the JSON encoding of u
func (u *User) MarshalJSON() ([]byte, error) {
	globs, err := json.Marshal(u.Globs)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&userJSON{
		Username: u.Username,
		Password: u.Password,
		Globs:    globs,
	})
}

// UnmarshalJSON parses JSON data and stores it in u
func (u *User) UnmarshalJSON(b []byte) error {
	var j userJSON
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	globs, err := loadGlobs(j.Globs)
	if err != nil {
		return err
	}

	u.Username = j.Username
	u.Password = j.Password
	u.Globs = globs

	return nil
}

// dumpGlobs dumps slice of strings into JSON encoding
func dumpGlobs(globs []string) ([]byte, error) {
	return json.Marshal(&globs)
}

// loadGlobs converts JSON into a slice of strings
func loadGlobs(b []byte) ([]string, error) {
	var globs []string
	if err := json.Unmarshal(b, &globs); err != nil {
		return nil, err
	}
	return globs, nil
}
