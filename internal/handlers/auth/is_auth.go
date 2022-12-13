package auth

import (
	"fmt"
	"net/http"
	"os"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	// check state from cookie and value from Google callback value
	oauthState, _ := r.Cookie("oauthstate")
	fmt.Fprintf(os.Stdout, "isAuth: cookie: %v\n", oauthState.Value)
	fmt.Fprintf(os.Stdout, "isAuth: state: %v\n", r.FormValue("state"))
	if r.FormValue("state") != oauthState.Value {
		w.WriteHeader(http.StatusForbidden)
		return false
	}
	return true
}