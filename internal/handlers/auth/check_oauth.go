package auth

import (
	"fmt"
	"net/http"
	"os"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	oauthState, _ := r.Cookie("oauthstate")
	fmt.Fprintf(os.Stdout, "isAuth: cookie: %v\n", oauthState)
	fmt.Fprintf(os.Stdout, "isAuth: state: %v\n", r.FormValue("state"))
	if r.FormValue("state") != oauthState.Value {
		fmt.Fprintln(os.Stdout, "invalid oauth state")
		return false
	}
	return true
}