package repairevents

import "net/http"

func HTTPStatus(kind Kind) int {
	statuses := map[Kind]int{
		KindMissing: http.StatusNotFound,
		KindSystem:  http.StatusInternalServerError,
	}
	if kind == KindMissing {
		return statuses[KindMissing]
	}
	return statuses[KindSystem]
}
