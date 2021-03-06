package github

import (
	"net/http"
	"time"

	"github.com/sniperkit/xvcs/plugin/provider/github/go-github/pkg"
)

// isValidToken check if token is valid.
func (g *Github) isValidToken(httpClient *http.Client) bool {
	defer funcTrack(time.Now())

	resp, err := g.makeRequest(httpClient)
	if err != nil {
		return false
	}
	err = github.CheckResponse(resp)
	if _, ok := err.(*github.TwoFactorAuthError); ok {
		return false
	}

	return true
}
