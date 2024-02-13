package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}

type Handler struct {
	repo Repo
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e, err := parseEventFromBody(r)
		if err = validateEvent(e); err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Ошибка валидации данных: %v", err))
			return
		}

		e, _ = h.repo.CreateEvent(e)
		OkResponse(w, http.StatusCreated, e)
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e, err := parseEventFromBody(r)
		if err = validateEvent(e); err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Ошибка валидации данных: %v", err))
			return
		}

		e, err = h.repo.UpdateEvent(e)
		if err != nil {
			ErrResponse(w, http.StatusNotFound, err.Error())
			return
		}

		OkResponse(w, http.StatusOK, e)
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e, err := parseEventFromBody(r)
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Неверный ID события: %v", err))
			return
		}

		err = h.repo.DeleteEvent(e.ID)
		if err != nil {
			ErrResponse(w, http.StatusNotFound, err.Error())
			return
		}

		OkResponse(w, http.StatusNoContent, "Событие успешно удалено")
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Неверный формат даты: %v", err))
			return
		}

		events := h.repo.EventsForDay(date)
		OkResponse(w, http.StatusOK, events)
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Неверный формат даты: %v", err))
			return
		}

		events := h.repo.EventsForWeek(date)
		OkResponse(w, http.StatusOK, events)
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Неверный формат даты: %v", err))
			return
		}

		events := h.repo.EventsForMonth(date)
		OkResponse(w, http.StatusOK, events)
	default:
		ErrResponse(w, http.StatusMethodNotAllowed, MethodNotAllowed)
		return
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func main() {
	var conf Config
	confData, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Ошибка чтения конф. файла: %v", err)
	}

	if err = yaml.Unmarshal(confData, &conf); err != nil {
		log.Fatalf("Ошибка декодирования конф. файла: %v", err)
	}

	h := Handler{repo: NewMemoryRepo()}
	http.HandleFunc("/create_event", LoggingMiddleware(h.CreateEvent))
	http.HandleFunc("/update_event", LoggingMiddleware(h.UpdateEvent))
	http.HandleFunc("/delete_event", LoggingMiddleware(h.DeleteEvent))
	http.HandleFunc("/events_for_day", LoggingMiddleware(h.EventsForDay))
	http.HandleFunc("/events_for_week", LoggingMiddleware(h.EventsForWeek))
	http.HandleFunc("/events_for_month", LoggingMiddleware(h.EventsForMonth))

	fmt.Printf("Сервер запущен на порту %s\n", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, nil))
}
