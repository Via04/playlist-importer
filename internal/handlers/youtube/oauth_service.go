package youtube

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/via04/playlist_importer/internal/config"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  config.Credentials.Web.RedirectUris[0],
	ClientID:     config.Credentials.Web.ClientId,
	ClientSecret: config.Credentials.Web.ClientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
	Endpoint:     google.Endpoint,
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
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

func generateStateOauthCookie(w http.ResponseWriter) string {
	// creates cookie with random key for each session to prevent CSRF attack
	expiration := time.Now().Add(time.Hour * 24)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration, SameSite: http.SameSiteNoneMode, Secure: true, Path: "/"}
	http.SetCookie(w, &cookie)
	return state
}
