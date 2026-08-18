package repairbatch

func RunBatch(l *Ledger, n int) {
	for i := 0; i < n; i++ {
		func() { x := l.Acquire(); defer x.Release() }()
	}
}
