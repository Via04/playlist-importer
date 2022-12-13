package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
)

var OauthGoogleCallback = func(w http.ResponseWriter, r *http.Request) {
	// callback from Google oauth url checks value of csrf token from google and compares to user saved value
	if !isAuthenticated(w, r) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	ctx := context.Background()
	fmt.Fprintf(os.Stdout, "Request code value: %v\n", r.FormValue("code"))
	token, err := handshakeYoutube(ctx, r.FormValue("code"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// now create session token for user
	fmt.Fprintf(os.Stdout, "token: %v", token)
	session := Sessions.Append(*token)
	fmt.Fprintf(os.Stdout, "Session: %v, key %v",
				func()oauth2.Token{
					token, _ := Sessions.Read(session) 
					return token
				}(), session)
	jsonOut, err := json.MarshalIndent(session, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonOut)
}

func handshakeYoutube(ctx context.Context, code string) (*oauth2.Token, error) {
	// test function to get lists from YouTube API
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	token, err := GoogleOauthConfig.Exchange(ctxWithTimeout, code)
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
