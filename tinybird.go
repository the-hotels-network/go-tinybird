package tinybird

import (
	"github.com/the-hotels-network/go-tinybird/internal/env"
)

var URL_BASE string
var MAX_IDLE_CONNS int
var MAX_CONNS_PER_HOST int
var MAX_IDLE_CONNS_PER_HOST int
var CONNS_TIMEOUT int

// Initialize module.
func init() {
	URL_BASE = env.Get("TB_URL_BASE", "https://api.tinybird.co/v0/pipes")
	MAX_IDLE_CONNS = env.GetInt("TB_MAX_IDLE_CONNS", 100)
	MAX_CONNS_PER_HOST = env.GetInt("TB_MAX_CONNS_PER_HOST", 100)
	MAX_IDLE_CONNS_PER_HOST = env.GetInt("TB_MAX_IDLE_CONNS_PER_HOST", 100)
	CONNS_TIMEOUT = env.GetInt("TB_CONNS_TIMEOUT", 30)
}
