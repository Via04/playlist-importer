package auth

import (
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