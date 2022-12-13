package callback

import (
	"encoding/base64"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

type session struct {
	rwmux   sync.RWMutex
	session map[string]oauth2.Token
	
}
var initOnce sync.Once

func (s *session) Append(token oauth2.Token)(session string) {
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

func (s *session) Read(session string)(token oauth2.Token) {
	s.rwmux.RLock()
	defer s.rwmux.RUnlock()
	return s.session[session]
}

var Sessions = new(session)