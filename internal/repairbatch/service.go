package repairbatch

func Validate(l *Ledger, invalid bool) error {
	x := l.Acquire()
	defer x.Release()
	if invalid {
		return ErrInvalid
	}
	return nil
}
