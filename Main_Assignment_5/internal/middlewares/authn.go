package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5" // atau pustaka JWT yang disediakan starter code
	"golang.org/x/crypto/argon2"
)

// Contoh struct payload registrasi/login
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // student, judge, admin
}

package middlewares

import (
	"aaa_bekend/internal/lmodel"
	"net/http"
)

func AuthN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// TODO

		// TODO

		ctx := lmodel.ContextWithUser(r.Context(), userCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
