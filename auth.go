package main

//import (
//	"context"
//	"database/sql"
//	"encoding/json"
//	"net/http"
//)
//
//// userKey is a context key
//const userKey = 98
//
//// authorizeUser is authorization middleware for user-specific handlers
//func authorizeUser(db *sql.DB, salt string, h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		username, password, _ := r.BasicAuth()
//		u, err := UserFind(db, salt, username, password)
//		if err == ErrInvalidCredentials {
//			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//			return
//		}
//		if err != nil {
//			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//
//		// add User to current request context
//		c := context.WithValue(r.Context(), userKey, u)
//		h(w, r.WithContext(c))
//	}
//}
//
//// userFromRequest retrieves User object from request context
//func userFromRequest(r *http.Request) *User {
//	return r.Context().Value(userKey).(*User)
//}
//
//// ANY /users
//func saveUser(db *sql.DB, salt string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//
//		// DELETE /users
//		if r.Method == http.MethodDelete {
//			deleteUser(db, w, r)
//			return
//		}
//
//		u := &User{}
//		if err := json.NewDecoder(r.Body).Decode(u); err != nil {
//			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//			return
//		}
//
//		// validate fields
//		if len(u.Username) < 4 || len(u.Password) < 4 || len(u.Regexps) == 0 {
//			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//			return
//		}
//
//		if err := UserSave(u, db, salt); err != nil {
//			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//// DELETE /users
//func deleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//	username, ok := r.URL.Query()["username"]
//	if ok && username[0] != "" {
//		if err := UserDelete(db, username[0]); err != nil {
//			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
