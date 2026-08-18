package equipmentlookup

import (
	"errors"
	"testing"
)

func TestUnknownEquipmentReturnsOneNotFoundWithoutRetry(t *testing.T) {
	err := StorageError(ErrMissing)
	if !errors.Is(err, ErrMissing) {
		t.Fatal("missing identity was discarded")
	}
	result := Decide(err)
	if result != Missing {
		t.Fatalf("result=%q", result)
	}
	code := Status(result)
	if code != 404 {
		t.Fatalf("code=%d", code)
	}
	if Retryable(code) {
		t.Fatal("missing equipment was retried")
	}
}
