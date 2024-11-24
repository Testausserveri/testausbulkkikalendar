package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"testausserveri/testausbulkkikalendar/oauth"
)

// Also init templates for use in handlers package
var templates *template.Template

func Init(templateGlob string) {
	templates = template.Must(template.ParseGlob(templateGlob))
}

// Index site handler
func Index(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	// Callback "URL" from Google auth
	if state == "state-token" {
		code := r.URL.Query().Get("code")
		fmt.Println(code)
		authToken, err := oauth.GetTokenFromCode(code)
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

	cookie, err := r.Cookie("auth-token")
	if err != nil {
		if err != http.ErrNoCookie {
			// Other errors
			http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
			return
		}
	}

	var isAuthenticated bool
	var authUrl string
	if cookie == nil {
		isAuthenticated = false
		authUrl = oauth.GetAuthURL()
	} else {
		isAuthenticated = true
	}

	data := struct {
		Title   string
		IsAuth  bool
		AuthUrl string
	}{
		Title:   "Testausbulkkikalendar",
		IsAuth:  isAuthenticated,
		AuthUrl: authUrl,
	}

	// Render the "index.html" template
	templates.ExecuteTemplate(w, "index.html", data)
}
