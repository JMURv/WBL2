package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type SuccessResponse struct {
	Result any `json:"result"`
}

type ErrorResponse struct {
	Error any `json:"error"`
}

func ErrResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&ErrorResponse{Error: data})
}

func OkResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&SuccessResponse{Result: data})
}

func parseEventFromBody(r *http.Request) (*Event, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var e Event
	if err := json.Unmarshal(body, &e); err != nil {
		return nil, err
	}

	return &e, nil
}

func validateEvent(e *Event) error {
	if e.ID == 0 {
		return fmt.Errorf("ID is required")
	}
	if e.UserID == 0 {
		return fmt.Errorf("UserID is required")
	}
	if e.Title == "" {
		return fmt.Errorf("UserID is required")
	}
	if e.Date == "" {
		return fmt.Errorf("date is required")
	}
	if _, err := time.Parse("2006-01-02", e.Date); err != nil {
		return fmt.Errorf("неверный формат даты: %v", err)
	}
	return nil
}

func parseAndValidateRequest(r *http.Request) (*Event, error) {
	idStr := r.FormValue("id")
	userIDStr := r.FormValue("user_id")
	title := r.FormValue("title")
	dateStr := r.FormValue("date")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("неверный ID события: %v", err)
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("неверный ID пользователя: %v", err)
	}

	_, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("неверный формат даты: %v", err)
	}

	return &Event{
		ID:     id,
		UserID: userID,
		Title:  title,
		Date:   dateStr,
	}, nil
}
