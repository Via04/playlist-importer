package callback

import (
	"fmt"

	"golang.org/x/oauth2"
)

type tokenSource struct {
	// have to write this structure implementing TokenSource interface to satisfy Google API service
	token oauth2.Token
}
func (t *tokenSource) setToken(token oauth2.Token) {
	t.token = token
}
func (t *tokenSource) Token() (*oauth2.Token, error) {
	if t.token == (oauth2.Token{}) {
		return nil, fmt.Errorf("tokenSource: token is not set")
	}
	return &t.token, nil
}