package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"testausserveri/testausbulkkikalendar/gcal"
	"testausserveri/testausbulkkikalendar/structs"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// Search calendar events
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Map form data to struct
	var q structs.Query
	q.Calendar = r.FormValue("calendar")
	q.Query = r.FormValue("query")
	q.StartDate = r.FormValue("dateStart")
	q.EndDate = r.FormValue("dateEnd")
	maxResults, err := strconv.ParseInt(r.FormValue("maxResults"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid maxResults value", http.StatusBadRequest)
		return
	}
	q.MaxResults = maxResults

	if q.StartDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", q.StartDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		// Format the date into RFC3339
		q.StartDate = parsedStartDate.Format(time.RFC3339)
	}

	if q.EndDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", q.EndDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		// Format the date into RFC3339
		q.EndDate = parsedEndDate.Format(time.RFC3339)
	}

	authToken := &oauth2.Token{}
	if authValues, ok := r.Context().Value("auth").(*AuthContext); ok {
		authToken = authValues.AuthToken
	} else {
		w.Write([]byte("Attributes not found"))
		return
	}

	eventList, err := gcal.QueryEvents(authToken, q)
	if err != nil {
		fmt.Println(err)
		return
	}

	events := eventList.Items
	templates.ExecuteTemplate(w, "query_results", struct{ Events []*calendar.Event }{events})
}
