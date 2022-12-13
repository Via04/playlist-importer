package handlers

// contains handlers for each path. For each new path new subpackage was created

import (
	"net/http"

	"github.com/via04/playlist_importer/internal/handlers/auth"
	"github.com/via04/playlist_importer/internal/handlers/callback"
	"github.com/via04/playlist_importer/internal/handlers/intro"
	"github.com/via04/playlist_importer/internal/handlers/list"
)

func New() http.Handler {
	smux := http.NewServeMux()
	smux.HandleFunc("/", intro.Intro)
	smux.HandleFunc("/youtube/login", auth.OauthGoogleLogin)
	smux.HandleFunc("/youtube/callback", callback.OauthGoogleCallback)
	smux.HandleFunc("/youtube/api", list.YoutubeList)
	return smux
}