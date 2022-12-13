package callback

import (
	"context"
	"fmt"
	"time"

	"github.com/via04/playlist_importer/internal/handlers/auth"
	"golang.org/x/oauth2"
)

// type savedSessions struct{
// 	sessions map[]
// }

func HandshakeYoutube(ctx context.Context, code string) (*oauth2.Token, error) {
	// test function to get lists from YouTube API
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	token, err := auth.GoogleOauthConfig.Exchange(ctxWithTimeout, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	return token, err
	// fmt.Fprintf(os.Stdout, "Expiration time: %v", token.Expiry)
	// tokenSource := auth.GoogleOauthConfig.TokenSource(ctxWithTimeout, token) // create autorefreshing tokenSource from config
	// youtubeService, err := youtube.NewService(ctxWithTimeout, option.WithTokenSource(tokenSource))
	// if err != nil {
	// 	return "", fmt.Errorf("service creation error: %s", err.Error())
	// }
	// youtubeList := youtubeService.Playlists.List([]string{"contentDetails", "id"}).Mine(true)
	// response, err := youtubeList.Do()
	// if err != nil {
	// 	return "", fmt.Errorf("youtube api query error: %s", err.Error())
	// }
	// responseByte, err := response.MarshalJSON()
	// if err != nil {
	// 	panic(err)
	// }
	// return string(responseByte), nil
}
