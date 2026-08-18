package equipmentlookup

func Retryable(status int) bool { return status >= 500 }
