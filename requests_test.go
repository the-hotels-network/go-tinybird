package tinybird_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/the-hotels-network/go-tinybird"
)

func newRequest(workspaceName, pipeName, responseData string) tinybird.Request {
	return tinybird.Request{
		Pipe: tinybird.Pipe{
			Name: pipeName,
			Workspace: tinybird.Workspace{
				Name: workspaceName,
			},
		},
		Before: func(r *tinybird.Request) bool {
			r.Response = tinybird.Response{
				Body: responseData,
			}

			return false
		},
	}
}

func TestMultipleRequestsSuccess(t *testing.T) {
	rs := tinybird.Requests{}
	rs.Add(newRequest("workspace1", "pipe1", `{"id":1}`))
	rs.Add(newRequest("workspace2", "pipe2", `{"id":2}`))

	rs.Execute()

	assert.Len(t, rs, 2)
	assert.ObjectsAreEqual(rs[0], rs.Get("workspace1", "pipe1"))
	assert.ObjectsAreEqual(tinybird.Request{}, rs.Get("random-workspace", "random-pipe"))
	assert.Equal(t, `{"id":1}`, rs[0].Response.Body)
	assert.Equal(t, `{"id":2}`, rs[1].Response.Body)
	assert.GreaterOrEqual(t, rs.Duration(), time.Duration(0).Nanoseconds())
}
