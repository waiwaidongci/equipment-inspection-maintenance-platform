package repairevents

import (
	"errors"
	"fmt"
)

var ErrMissingEvent = errors.New("repair event missing")

type Repository struct{}

func (Repository) Get(id string) error {
	if id == "archived" {
		message := fmt.Sprintf("get repair event %s", id)
		return fmt.Errorf("%s: %w", message, ErrMissingEvent)
	}
	return nil
}
