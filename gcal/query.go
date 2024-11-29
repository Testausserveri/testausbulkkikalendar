package gcal

import (
	"context"
	"testausserveri/testausbulkkikalendar/structs"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func QueryEvents(
	token *oauth2.Token,
	q structs.Query,
) (*calendar.Events, error) {
	ctx := context.Background()
	client := getClient(token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	eventQuery := srv.Events.List(q.Calendar).ShowDeleted(false).SingleEvents(true)
	if q.StartDate != "" {
		eventQuery = eventQuery.TimeMin(q.StartDate)
	}
	if q.EndDate != "" {
		eventQuery = eventQuery.TimeMax(q.EndDate)
	}
	if q.MaxResults != -1 {
		eventQuery = eventQuery.MaxResults(q.MaxResults)
	}
	if q.Query != "" {
		eventQuery = eventQuery.Q(q.Query)
	}
	events, err := eventQuery.Do()
	if err != nil {
		return nil, err
	}
	return events, nil
}
