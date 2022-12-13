package auth

import (
	"fmt"
	"net/http"

	"github.com/via04/playlist_importer/internal/config"
)

var OauthGoogleLogin = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	oauthState := generateStateOauthCookie(w)
	u := GoogleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "%s\naddr:%p\n", config.Credentials.Web.ClientId, &config.Credentials)
	fmt.Fprintln(w, "Google Login")
}