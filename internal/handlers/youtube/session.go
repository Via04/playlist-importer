package youtube

import (
	"encoding/base64"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

type session struct {
	// Contains active user session with session_token as key and Google oauth2.token as value
	rwmux   sync.RWMutex
	session map[string]oauth2.Token
	
}
var initOnce sync.Once

func (s *session) Append(token oauth2.Token)(session string) {
	// Multithread-safe append to sessions
	initOnce.Do(func() {s.session = make(map[string]oauth2.Token)})
	s.rwmux.Lock()
	defer s.rwmux.Unlock()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	key := make([]byte, 32)
	r.Read(key)
	session = base64.URLEncoding.EncodeToString(key)
	s.session[session] = token
	return session
}

func (s *session) Read(session string)(token oauth2.Token, ok bool) {
	// Multithread-safe read for sessions
	s.rwmux.RLock()
	defer s.rwmux.RUnlock()
	token, ok = s.session[session]
	return token, ok
}

var Sessions = new(session)