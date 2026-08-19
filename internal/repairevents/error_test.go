package repairevents

import (
	"errors"
	"net/http"
	"testing"
)

func TestArchivedEventReturnsMissingWithoutRetry(t *testing.T) {
	err := (Repository{}).Get("archived")
	if !errors.Is(err, ErrMissingEvent) {
		t.Errorf("repository broke missing chain: %v", err)
	}
	if got := Classify(err); got != KindMissing {
		t.Errorf("classifier returned %q", got)
	}
	if got := HTTPStatus(KindMissing); got != http.StatusNotFound {
		t.Errorf("handler returned %d", got)
	}
	if ShouldRetry(KindMissing, 0) {
		t.Error("missing event was retried")
	}
}
