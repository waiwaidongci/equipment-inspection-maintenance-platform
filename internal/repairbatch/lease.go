package repairbatch

type Ledger struct{ Open, Peak int }
type Lease struct {
	l      *Ledger
	closed bool
}

func (l *Ledger) Acquire() *Lease {
	l.Open++
	if l.Open > l.Peak {
		l.Peak = l.Open
	}
	return &Lease{l: l}
}
func (x *Lease) Release() {
	if x.closed {
		return
	}
	x.l.Open--
	x.closed = true
}
