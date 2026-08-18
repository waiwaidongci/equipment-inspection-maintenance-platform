package taskpipeline

import "errors"

var ErrInvalid = errors.New("invalid task")

func ErrorChannel() chan error { return make(chan error, 1) }
