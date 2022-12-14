package config

// Unmarshals file secret.json into exported variable Credentials

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Web struct {
	Web Content `json:"Pjc"`
}

type Content struct {
	ClientId     string   `json:"client_id"`
	ProjectId    string   `json:"project_id"`
	TokenUri     string   `json:"token_uri"`
	AuthProvider string   `json:"auth_provider_x509_cert_url"`
	ClientSecret string   `json:"client_secret"`
	RedirectUris []string `json:"redirect_uris"`
}

func credentials() Web {
	var credentials *Web
	var createCredentials sync.Once
	return func() Web {
		createCredentials.Do(func() {
			secret, err := os.ReadFile("./data/secret.json")
			if err != nil {
			fmt.Fprintln(os.Stderr, "IOError: no secret.json")
			}
			credentials = new(Web)
			json.Unmarshal(secret, credentials)
			fmt.Fprintf(os.Stdout, "new_addr:%p\n", credentials)
		})
		fmt.Fprintf(os.Stdout, "saved_addr:%p\n", credentials)
		return *credentials
	}()
}

var Credentials = credentials()