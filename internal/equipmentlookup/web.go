package equipmentlookup

func Status(r Result) int {
	if r == Missing {
		return 404
	}
	return 500
}
