package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/api/youtube/v3"
)

type apiResolver struct {
	ctx     context.Context
	w       http.ResponseWriter
	r       map[string]interface{}
	service *youtube.Service
}

func (r *apiResolver) List() {
	// get full list of playlists from youtube api to decide which use later
	// r parameter will be used later
	ctxWithTimeout, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()
	youtubeList := r.service.Playlists.List([]string{"contentDetails", "id", "snippet"}).Mine(true)
	youtubeList.Context(ctxWithTimeout)
	response, err := youtubeList.Do()
	if err != nil {
		r.w.WriteHeader(http.StatusRequestTimeout)
		message, _ := json.MarshalIndent(map[string]string{"error": "no answer from Google"}, "", "    ")
		r.w.Write(message)
	}
	responseByte, err := response.MarshalJSON()
	if err != nil {
		panic(err)
	}
	r.w.Write(responseByte)
}

func (r *apiResolver) ListItems() {
	ctxWithTimeout, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()
	playlistId, ok := r.r["playlistId"].(string)
	if !ok {
		r.w.WriteHeader(http.StatusBadRequest)
		r.w.Write(func() []byte {
			out, _ := json.MarshalIndent(map[string]string{"error": "no playlistId specified"}, "", "    ")
			return out
		}())
	}
	call := r.service.PlaylistItems.List([]string{"contentDetails", "id", "snippet"}).PlaylistId(playlistId)
	call.Context(ctxWithTimeout)
	response, err := call.Do()
	if err != nil {
		r.w.WriteHeader(http.StatusGatewayTimeout)
		r.w.Write(func() []byte {
			out, _ := json.MarshalIndent(map[string]string{"error": "no answer from Google"}, "", "    ")
			return out
		}())
	}
	r.w.Write(func() []byte {
		out, _ := response.MarshalJSON()
		return out
	}())
}

func(r * apiResolver) Import() {
	query, ok := r.r["query"]
	if ok != true {
		r.w.WriteHeader(http.StatusBadRequest)
		r.w.Write(func() []byte {
			out, _ := json.MarshalIndent(map[string]string{"error": "no query specified"}, "", "    ")
			return out
		}())
	}
	fmt.Fprintf(r.w, "query is: %v", query)
}