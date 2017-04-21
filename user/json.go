package user

import (
	"encoding/json"
	"regexp"
)

// UserJSON is an intermediate structure for JSON (un)marshaling
type UserJSON struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Regexps  json.RawMessage `json:"regexps"`
}

// MarshalJSON returns the JSON encoding of u
func (u *User) MarshalJSON() ([]byte, error) {
	regexps, err := dumpRegexps(u.Regexps)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&UserJSON{
		Username: u.Username,
		Password: u.Password,
		Regexps:  regexps,
	})
}

// UnmarshalJSON parses JSON data and stores it in u
func (u *User) UnmarshalJSON(b []byte) error {
	var j UserJSON
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	regexps, err := loadRegexps(j.Regexps)
	if err != nil {
		return err
	}

	u.Username = j.Username
	u.Password = j.Password
	u.Regexps = regexps

	return nil
}

// dumpRegexps dumps slice of regexps into JSON encoding
func dumpRegexps(rr []*regexp.Regexp) ([]byte, error) {
	exprs := make([]string, 0, len(rr))
	for _, r := range rr {
		exprs = append(exprs, r.String())
	}
	return json.Marshal(exprs)
}

// loadRegexps converts JSON into a slice of regexps
func loadRegexps(b []byte) ([]*regexp.Regexp, error) {
	var exprs []string
	if err := json.Unmarshal(b, &exprs); err != nil {
		return nil, err
	}

	regexps := make([]*regexp.Regexp, 0, len(exprs))
	for _, expr := range exprs {
		r, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		regexps = append(regexps, r)
	}

	return regexps, nil
}
