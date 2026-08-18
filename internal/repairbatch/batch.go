package repairbatch

func RunBatch(l *Ledger, n int) {
	for i := 0; i < n; i++ {
		runOne(l)
	}
}

// runOne is extracted so that the deferred Release executes at the end of
// each iteration, not when RunBatch returns.  Without this, every lease
// is held until the entire loop finishes, exhausting connections / file
// handles when n is large.
func runOne(l *Ledger) {
	x := l.Acquire()
	defer x.Release()
}
