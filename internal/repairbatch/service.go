package repairbatch

func Validate(l *Ledger, invalid bool) error {
	x := l.Acquire()
	if invalid {
		return ErrInvalid
	}
	x.Release()
	return nil
}
