package handlers

// contains handlers for each path. For each new path new subpackage was created

import (
	"net/http"

	"github.com/via04/playlist_importer/internal/handlers/youtube"
	"github.com/via04/playlist_importer/internal/handlers/intro"
	"github.com/via04/playlist_importer/internal/handlers/youtube/api"
)

func New() http.Handler {
	smux := http.NewServeMux()
	smux.HandleFunc("/", intro.Intro)
	smux.HandleFunc("/youtube/login", youtube.OauthGoogleLogin)
	smux.HandleFunc("/youtube/callback", youtube.OauthGoogleCallback)
	smux.HandleFunc("/youtube/api", api.Decide)
	return smux
}