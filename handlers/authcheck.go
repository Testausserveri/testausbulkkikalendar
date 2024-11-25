package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"testausserveri/testausbulkkikalendar/oauth"

	"golang.org/x/oauth2"
)

type AuthContext struct {
	AuthToken *oauth2.Token
	AuthURL   string
}

func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth-token")
		if err != nil {
			if err != http.ErrNoCookie {
				// Other errors
				http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
				return
			}
		}

		authToken := &oauth2.Token{}
		var authURL string
		if cookie == nil {
			authURL = oauth.GetAuthURL()
		} else {
			// Decode the string
			decodedToken, err := base64.StdEncoding.DecodeString(cookie.Value)
			if err != nil {
				http.Error(w, "Error decoding cookie", http.StatusInternalServerError)
				return
			}
			json.Unmarshal(decodedToken, authToken)
		}

		// Add an attribute to the context
		ctx := context.WithValue(r.Context(), "auth", &AuthContext{authToken, authURL})
		// Create a new request with the updated context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
