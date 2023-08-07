package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	auth "github.com/trent-howard/go-rss/internal"
	"github.com/trent-howard/go-rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				respondWithError(w, http.StatusForbidden, "User not found")
			default:
				respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			}
			return
		}
		handler(w, r, user)
	}
}
