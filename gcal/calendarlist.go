package gcal

import (
	"context"
	"sort"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func ListCalendars(token *oauth2.Token) ([]*calendar.CalendarListEntry, error) {
	ctx := context.Background()
	client := getClient(token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	calendarList, err := srv.CalendarList.List().Do()
	if err != nil {
		return nil, err
	}
	calendars := calendarList.Items
	sort.Slice(calendars, func(i, j int) bool {
		return calendars[i].Primary && !calendars[j].Primary
	})

	return calendars, nil
}
