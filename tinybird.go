package tinybird

import (
	"github.com/the-hotels-network/go-tinybird/internal/env"
)

var URL_BASE string
var NDJSON bool

// Initialize module.
func init() {
	URL_BASE = env.Get("TB_URL_BASE", "https://api.tinybird.co/v0/pipes")
	NDJSON   = env.GetBool("TB_NDJSON", false)
}

// Return the URL base.
func URL() string {
	return URL_BASE
}

// Return the JSON format response.
func Format() string {
	if NDJSON {
		return "ndjson"
	}

	return "json"
}
