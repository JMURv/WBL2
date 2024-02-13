package main

import (
	"sync"
	"time"
)

type Repo interface {
	CreateEvent(e *Event) (*Event, error)
	UpdateEvent(e *Event) (*Event, error)
	DeleteEvent(eventID int) error
	EventsForDay(date time.Time) []*Event
	EventsForWeek(date time.Time) []*Event
	EventsForMonth(date time.Time) []*Event
}

type MemoryRepo struct {
	sync.RWMutex
	data map[int]*Event
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{data: make(map[int]*Event)}
}

func (r *MemoryRepo) CreateEvent(e *Event) (*Event, error) {
	r.Lock()
	r.data[e.ID] = e
	r.Unlock()
	return e, nil
}

func (r *MemoryRepo) UpdateEvent(e *Event) (*Event, error) {
	r.RLock()
	_, ok := r.data[e.ID]
	r.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}

	r.Lock()
	defer r.Unlock()
	r.data[e.ID] = e
	return r.data[e.ID], nil
}

func (r *MemoryRepo) DeleteEvent(eventID int) error {
	r.RLock()
	_, exists := r.data[eventID]
	r.RUnlock()

	if !exists {
		return ErrNotFound
	}

	r.Lock()
	delete(r.data, eventID)
	r.Unlock()
	return nil

}

func (r *MemoryRepo) EventsForDay(date time.Time) []*Event {
	r.RLock()
	events := make([]*Event, 0, len(r.data))
	for _, e := range r.data {
		evDate, _ := time.Parse("2006-01-02", e.Date)
		if evDate.Year() == date.Year() &&
			evDate.Month() == date.Month() &&
			evDate.Day() == date.Day() {
			events = append(events, e)
		}
	}
	r.RUnlock()
	return events
}

func (r *MemoryRepo) EventsForWeek(date time.Time) []*Event {
	r.RLock()
	events := make([]*Event, 0, len(r.data))
	nextWeek := date.AddDate(0, 0, 7) // Добавляем 7 дней к дате
	for _, e := range r.data {
		evDate, _ := time.Parse("2006-01-02", e.Date)
		if evDate.After(date) && evDate.Before(nextWeek) {
			events = append(events, e)
		}
	}
	r.RUnlock()
	return events
}

func (r *MemoryRepo) EventsForMonth(date time.Time) []*Event {
	r.RLock()
	events := make([]*Event, 0, len(r.data))
	for _, e := range r.data {
		evDate, _ := time.Parse("2006-01-02", e.Date)
		if evDate.Year() == date.Year() &&
			evDate.Month() == date.Month() {
			events = append(events, e)
		}
	}
	r.RUnlock()
	return events
}
