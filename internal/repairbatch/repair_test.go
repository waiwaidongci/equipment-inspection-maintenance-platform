package repairbatch

import (
	"errors"
	"testing"
)

func TestRepairBatchReleasesLeasesAndKeepsErrors(t *testing.T) {
	t.Run("batch", func(t *testing.T) {
		l := &Ledger{}
		RunBatch(l, 7)
		if l.Peak > 1 || l.Open != 0 {
			t.Fatalf("peak=%d open=%d", l.Peak, l.Open)
		}
	})
	t.Run("error", func(t *testing.T) {
		if !errors.Is(Save(&Tx{}, ErrInvalid), ErrInvalid) {
			t.Fatal("validation error lost")
		}
	})
	t.Run("validation", func(t *testing.T) {
		l := &Ledger{}
		if !errors.Is(Validate(l, true), ErrInvalid) {
			t.Fatal("invalid repair accepted")
		}
		if l.Open != 0 {
			t.Fatalf("open=%d", l.Open)
		}
	})
	t.Run("release twice", func(t *testing.T) {
		l := &Ledger{}
		x := l.Acquire()
		x.Release()
		x.Release()
		if l.Open != 0 {
			t.Fatalf("open=%d", l.Open)
		}
	})
}
