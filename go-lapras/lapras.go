// Package lapras provide method for LAPRAS inc. API.
package lapras

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Lapras provide method for LAPRAS inc. API
// API document: https://github.com/lapras-inc/public-api-schema
type Lapras struct {
	// baseURL is base url of LAPRAS inc. API
	baseURL url.URL
}

// NewLapras return Lapras struct pointer.
func NewLapras() *Lapras {
	return &Lapras{
		// https://lapras.com/public
		baseURL: url.URL{
			Scheme: "https",
			Host:   "lapras.com",
			Path:   "/public",
		},
	}
}

// GetPerson returns detailed information about the Lapras user.
// shareID is an identifier for users to publish their own information.
func (l *Lapras) GetPerson(shareID string) (*Person, error) {
	endpoint, err := url.JoinPath(l.baseURL.String(), shareID+".json")
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var person Person
	if err := json.Unmarshal(body, &person); err != nil {
		return nil, err
	}
	return &person, nil
}
