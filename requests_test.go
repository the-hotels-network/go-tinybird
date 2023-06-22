package tinybird_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/the-hotels-network/go-tinybird"
)

func newRequest(workspaceName, pipeName string) tinybird.Request {
	return tinybird.Request{
		Pipe: tinybird.Pipe{
			Name: pipeName,
			Workspace: tinybird.Workspace{
				Name: workspaceName,
			},
		},
		Before: func(r *tinybird.Request) bool {
			return false
		},
	}
}

func TestMultipleRequestsSuccess(t *testing.T) {
	rs := tinybird.Requests{}
	rs.Add(newRequest("workspace1", "pipe1"))
	rs.Add(newRequest("workspace2", "pipe2"))

	rs.Execute()

	assert.Len(t, rs, 2)
	assert.ObjectsAreEqual(rs[0], rs.Get("workspace1", "pipe1"))
	assert.ObjectsAreEqual(tinybird.Request{}, rs.Get("random-workspace", "random-pipe"))
	assert.GreaterOrEqual(t, rs.Duration(), time.Duration(0).Nanoseconds())
}
