package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5" 
	"golang.org/x/crypto/argon2"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` 
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
