package repairevents

import "net/http"

func HTTPStatus(kind Kind) int {
	if kind == KindMissing {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
