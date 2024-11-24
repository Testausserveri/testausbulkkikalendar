package oauth

import (
	"context"
	"log"
	"os"
	"testausserveri/testausbulkkikalendar/constants"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

var (
	ctx    context.Context
	config *oauth2.Config
)

func GetTokenFromCode(code string) (*oauth2.Token, error) {
	return config.Exchange(context.TODO(), code)
}

func GetAuthURL() string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func Init() {
	ctx = context.Background()
	b, err := os.ReadFile(constants.SECRETS_PATH + "/oauth.json")
	if err != nil {
		log.Fatalf("[OAUTH] Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err = google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("[OAUTH] Unable to parse client secret file to config: %v", err)
	}
}
