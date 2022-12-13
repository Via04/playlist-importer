package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	auth "github.com/via04/playlist_importer/internal/handlers/youtube"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Decide(w http.ResponseWriter, r *http.Request) {
	jsonBody := make(map[string]interface{})
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message, _ := json.MarshalIndent(map[string]string{"error": "cannot unmarshal body"}, "", "    ")
		w.Write(message)
		return
	}
	session, ok := jsonBody["token"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message, _ := json.MarshalIndent(map[string]string{"error": "no token passed"}, "", "    ")
		w.Write(message)
		return
	}
	token, ok := auth.Sessions.Read(session)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		message, _ := json.MarshalIndent(map[string]string{"error": "wrong token"}, "", "    ")
		w.Write(message)
		return
	} else {
		fmt.Printf("got token from body: %v\n", token)
	}
	method, ok := jsonBody["method"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message, _ := json.MarshalIndent(map[string]string{"error": "no method passed"}, "", "    ")
		w.Write(message)
		return
	}
	switch method {
	case "list":
		List(ctx, w, r, token)
	}
}

func List(ctx context.Context, w http.ResponseWriter, r *http.Request, token oauth2.Token) {
	// get full playlist from youtube api
	// r parameter will be used later
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()
	tokenSource := auth.GoogleOauthConfig.TokenSource(ctxWithTimeout, &token)
	youtubeService, err := youtube.NewService(ctxWithTimeout, option.WithTokenSource(tokenSource))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		message, _ := json.MarshalIndent(map[string]string{"error": "cannot connect to Google servers"}, "", "    ")
		w.Write(message)
	}
	youtubeList := youtubeService.Playlists.List([]string{"contentDetails", "id"}).Mine(true)
	response, err := youtubeList.Do()
	if err != nil {
		w.WriteHeader(http.StatusRequestTimeout)
		message, _ := json.MarshalIndent(map[string]string{"error": "no answer from Google"}, "", "    ")
		w.Write(message)
	}
	responseByte, err := response.MarshalJSON()
	if err != nil {
		panic(err)
	}
	w.Write(responseByte)
}