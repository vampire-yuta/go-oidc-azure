package main

import (
	"context"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var once sync.Once

var provider *oidc.Provider
var oauth2Config *oauth2.Config

func getConfig() (*oauth2.Config, *oidc.Provider) {
	once.Do(func() {
		var err error
		// ここにissuer情報を設定
		provider, err = oidc.NewProvider(context.Background(), "https://login.microsoftonline.com/66674ed0-005d-4b5a-9138-2a7725b9df89/oauth2/v2.0/authorize")
		if err != nil {
			panic(err)
		}
		oauth2Config = &oauth2.Config{
			// ここにクライアントIDとクライアントシークレットを設定
			ClientID:     "95972570-f3ed-4c22-9f4c-c6a088a0c4bd",
			ClientSecret: "YfEsYuD_64etBd-r0VVH5s~0c9v8C3k~Zs",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://localhost:8080/callback",
		}
	})
	return oauth2Config, provider
}
