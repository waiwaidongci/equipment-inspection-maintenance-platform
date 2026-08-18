package equipmentlookup

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("equipment missing")

func StorageError(err error) error { return fmt.Errorf("read equipment: %v", err) }
