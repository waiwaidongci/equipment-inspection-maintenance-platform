package inspectioncache

type Snapshot map[string]int

func (snapshot Snapshot) Copy() Snapshot {
	return snapshot
}
