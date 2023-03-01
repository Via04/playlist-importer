package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/via04/playlist_importer/internal/handlers/youtube"
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
	tokenSource := auth.GoogleOauthConfig.TokenSource(ctx, &token)
	youtubeService, err := youtube.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write(func() []byte {
			out, _ := json.MarshalIndent(map[string]string{"err": "no answer from Google Services"}, "", "    ")
			return out
		}())
	}
	resolver := apiResolver{
		ctx:     ctx,
		w:       w,
		r:       jsonBody,
		service: youtubeService,
	}
	switch method {
	case "list":
		resolver.List()
	case "listItems":
		resolver.ListItems()
	}
}
