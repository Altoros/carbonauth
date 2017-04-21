package main
//
//import (
//	"bytes"
//	"encoding/json"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"regexp"
//	"testing"
//)
//
//func TestUserAuthorization(t *testing.T) {
//	databaseURL := os.Getenv("TEST_DATABASE_URL")
//	if databaseURL == "" {
//		t.Skip("TEST_DATABASE_URL is not provided")
//	}
//
//	db, err := connectDB(databaseURL)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer cleanDB(db)
//
//	// saveUser handler
//	a := basicAuth(saveUser(db, "salt"), "admin", "secret")
//
//	// handler that requires user authorization
//	h := authorizeUser(db, "salt", func(w http.ResponseWriter, _ *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	})
//
//	testSaveUser(t, a, "create_user", "admin", "secret", &User{
//		Username: "user",
//		Password: "secret",
//		Regexps:  []*regexp.Regexp{regexp.MustCompile("^.*$")},
//	})
//
//	testAccess(t, h, "authorized_access", "user", "secret", http.StatusOK)
//	testAccess(t, h, "unauthorized_access", "user", "secret132", http.StatusUnauthorized)
//
//	testSaveUser(t, a, "change_user_creds", "admin", "secret", &User{
//		Username: "user",
//		Password: "secret132",
//		Regexps:  []*regexp.Regexp{regexp.MustCompile("^.*$")},
//	})
//
//	testAccess(t, h, "new_credentials", "user", "secret132", http.StatusOK)
//
//	// delete user
//	w := httptest.NewRecorder()
//	r := httptest.NewRequest(http.MethodDelete, "/users?username=user", nil)
//	r.SetBasicAuth("admin", "secret")
//	a.ServeHTTP(w, r)
//	if w.Code != http.StatusOK {
//		t.Fatalf("delete user status = %d, want %d", w.Code, http.StatusOK)
//	}
//
//	testAccess(t, h, "deleted_user", "user", "secret132", http.StatusUnauthorized)
//}
//
//func testAccess(t *testing.T, h http.HandlerFunc, name, username, password string, want int) {
//	t.Run(name, func(t *testing.T) {
//		w := httptest.NewRecorder()
//		r := httptest.NewRequest(http.MethodGet, "/", nil)
//		r.SetBasicAuth(username, password)
//		h.ServeHTTP(w, r)
//		if w.Code != want {
//			t.Errorf("status = %d, want %d", w.Code, want)
//		}
//	})
//}
//
//func testSaveUser(t *testing.T, h http.Handler, name, username, password string, user *User) {
//	t.Run(name, func(t *testing.T) {
//		b, err := json.Marshal(user)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		w := httptest.NewRecorder()
//		r := httptest.NewRequest(http.MethodPatch, "/users", bytes.NewReader(b))
//		r.SetBasicAuth(username, password)
//		h.ServeHTTP(w, r)
//
//		if w.Code != http.StatusOK {
//			t.Fatalf("create user status = %d, want %d", w.Code, http.StatusOK)
//		}
//	})
//}
