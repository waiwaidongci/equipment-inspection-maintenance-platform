package checklist

func AddItem(s *Settings, name string) {
	if s.Items == nil {
		s.Items = make(map[string]bool)
	}
	s.Items[name] = true
}
