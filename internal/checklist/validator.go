package checklist

type Validator interface{ Check(string) error }
type required struct{}

func (*required) Check(string) error { return nil }
func OptionalValidator() Validator   { var v *required; return v }
