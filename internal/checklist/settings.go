package checklist

type Settings struct{ Items map[string]bool }

func DefaultSettings() Settings { return Settings{} }
