package callback

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/via04/playlist_importer/internal/handlers/auth"
)

var OauthGoogleCallback = func(w http.ResponseWriter, r *http.Request) {

	if !auth.IsAuthenticated(w, r) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	ctx := context.Background()
	fmt.Fprintf(os.Stdout, "Request code value: %v\n", r.FormValue("code"))
	response, err := ExchangeUserDataGoogle(ctx, r.FormValue("code"))
	if err != nil {
		fmt.Fprintln(w, http.StatusBadRequest)
	}
	fmt.Fprintln(w, response)
}