package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"testausserveri/testausbulkkikalendar/gcal"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// IndexHandler site handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	// Callback "URL" from Google auth
	if state == "state-token" {
		code := r.URL.Query().Get("code")
		authToken, err := gcal.GetTokenFromCode(code)
		if err != nil {
			http.Error(w, "Error retrieving auth token", http.StatusInternalServerError)
			return
		}

		jsonAuthToken, err := json.Marshal(authToken)
		if err != nil {
			http.Error(w, "Error transforming auth token", http.StatusInternalServerError)
			return
		}

		encodedAuthToken := base64.StdEncoding.EncodeToString(jsonAuthToken)

		cookie := &http.Cookie{
			Name:     "auth-token",
			Value:    encodedAuthToken,
			Path:     "/",  // Scope to root and all subpaths
			HttpOnly: true, // Accessible only by the server
			MaxAge:   3600, // Expires in 1 hour
		}

		// Set the cookie
		http.SetCookie(w, cookie)
		// Redirect to root
		http.Redirect(w, r, "/", http.StatusFound) // 302 Found
		return
	}

	// Get authentication attributes
	var authURL string
	authToken := &oauth2.Token{}
	if authValues, ok := r.Context().Value("auth").(*AuthContext); ok {
		authToken = authValues.AuthToken
		authURL = authValues.AuthURL
	} else {
		w.Write([]byte("Attributes not found"))
		return
	}

	var calendars []*calendar.CalendarListEntry
	if authToken.Valid() {
		var err error
		calendars, err = gcal.ListCalendars(authToken)
		if err != nil {
			fmt.Println(err)
			calendars = []*calendar.CalendarListEntry{}
		}
		fmt.Println(calendars)
	}

	data := struct {
		Title     string
		IsAuth    bool
		AuthURL   string
		Calendars []*calendar.CalendarListEntry
	}{
		Title:     "Testausbulkkikalendar",
		IsAuth:    authToken.Valid(),
		AuthURL:   authURL,
		Calendars: calendars,
	}

	// Render the "index.html" template
	templates.ExecuteTemplate(w, "index.html", data)
}
