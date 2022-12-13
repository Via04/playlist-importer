package list

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/via04/playlist_importer/internal/handlers/auth"
	"github.com/via04/playlist_importer/internal/handlers/callback"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YoutubeList(w http.ResponseWriter, r *http.Request) {
	// get full playlist from youtube api
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	session := r.URL.Query().Get("token")
	fmt.Fprintf(os.Stdout, "\nsession token: %v\n", session)
	token := callback.Sessions.Read(session)
	fmt.Fprintf(os.Stdout, "token: %v\n", token)
	tokenSource := auth.GoogleOauthConfig.TokenSource(context.TODO(), &token)
	youtubeService, err := youtube.NewService(context.TODO(), option.WithTokenSource(tokenSource))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	youtubeList := youtubeService.Playlists.List([]string{"contentDetails", "id"}).Mine(true)
	response, err := youtubeList.Do()
	if err != nil {
		w.WriteHeader(http.StatusRequestTimeout)
	}
	responseByte, err := response.MarshalJSON()
	if err != nil {
		panic(err)
	}
	w.Write(responseByte)
}