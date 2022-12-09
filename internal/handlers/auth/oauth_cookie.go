package auth

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"time"
)



func generateStateOauthCookie(w http.ResponseWriter) string {
	// creates cookie with random key for each session to prevent CSRF attack
	expiration := time.Now().Add(time.Hour * 24)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration, SameSite: http.SameSiteNoneMode, Secure: true, Path: "/"}
	http.SetCookie(w, &cookie)
	return state
}