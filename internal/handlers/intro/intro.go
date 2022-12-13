package intro
// Welcome page code

import (
	"net/http"
)

var Intro = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
		w.Write([]byte("intro page"))
}