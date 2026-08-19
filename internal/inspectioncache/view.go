package inspectioncache

type Snapshot map[string]int

func (snapshot Snapshot) Copy() Snapshot {
	out := make(Snapshot, len(snapshot))
	for key, value := range snapshot {
		out[key] = value
	}
	return out
}
