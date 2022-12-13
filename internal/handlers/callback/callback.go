package callback

import (
	"context"
	"encoding/json"
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
	token, err := HandshakeYoutube(ctx, r.FormValue("code"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// now create session token for user
	fmt.Fprintf(os.Stdout, "token: %v", token)
	session := Sessions.Append(*token)
	fmt.Fprintf(os.Stdout, "Session: %v, key %v", Sessions.Read(session), session)
	jsonOut, err := json.MarshalIndent(session, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonOut)
}
