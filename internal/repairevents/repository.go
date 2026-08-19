package repairevents

import (
	"errors"
	"fmt"
)

var ErrMissingEvent = errors.New("repair event missing")

type Repository struct{}

func (Repository) Get(id string) error {
	if id == "archived" {
		return fmt.Errorf("get repair event %s: %w", id, ErrMissingEvent)
	}
	return nil
}
