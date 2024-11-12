package main

import (
	"fmt"
	"net/http"

	"github.com/FilledEther/rss_reader/internal/auth"
	"github.com/FilledEther/rss_reader/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		// log.Fatalf("%v", apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user%v", err))
		}

		handler(w, r, user)

	}
}