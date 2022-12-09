package callback

import (
	"context"
	"fmt"
	"time"

	"github.com/via04/playlist_importer/internal/handlers/auth"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func ExchangeUserDataGoogle(ctx context.Context, code string) (responseJSON string, err error) {
	// test function to get lists from YouTube API
	tokenSource := new(tokenSource)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()
	token, err := auth.GoogleOauthConfig.Exchange(ctxWithTimeout, code)
	if err != nil {
		return "", fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	tokenSource.setToken(*token)
	youtubeService, err := youtube.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return "", fmt.Errorf("service creation error: %s", err.Error())
	}
	youtubeList := youtubeService.Playlists.List([]string{"contentDetails", "id"},).Mine(true)
	response, err := youtubeList.Do()
	if err != nil {
		return "", fmt.Errorf("youtube api query error: %s", err.Error())
	}
	responseByte, err := response.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(responseByte), nil
}