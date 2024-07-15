package calendar

import (
	"context"
)

type Calendar struct { // TODO
}

type Logger interface { // TODO
}

type Storage interface { // TODO
}

func New(storage Storage) *Calendar {
	return &Calendar{}
}

func (a *Calendar) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
