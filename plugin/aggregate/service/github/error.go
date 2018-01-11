package github

import (
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
)

var (
	errorNotStruct           = errors.New("error, not a struct")
	errorMappingJson         = errors.New("error while trying to map json response")
	errorRateLimitReached    = errors.New("error, rate limite for the current token is reached.")
	errorResponseIsNull      = errors.New("error while receiving the response of the http request.")
	errorMarshallingResponse = errors.New("error while trying to marshall the api response, entity object is nil")
)

func isTemporaryError(err error, wait bool) bool {
	defer funcTrack(time.Now())

	if err == nil {
		return false
	}
	// Get the underlying error, if this is a Wrapped error by the github.com/pkg/errors package.
	// If not, this will just return the error itself.
	underlyingErr := errors.Cause(err)
	switch underlyingErr.(type) {
	case *github.RateLimitError:
		return true
	case *github.AbuseRateLimitError:
		if wait {
			time.Sleep(2 * time.Second)
		}
		return true
	default:
		if strings.Contains(err.Error(), "abuse detection") {
			if wait {
				time.Sleep(2 * time.Second)
			}
			return true
		}
		if strings.Contains(err.Error(), "try again") {
			if wait {
				time.Sleep(2 * time.Second)
			}
			return true
		}
		return false
	}
}
