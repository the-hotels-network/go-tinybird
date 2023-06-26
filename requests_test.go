package tinybird_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/the-hotels-network/go-tinybird"
)

func newRequest(workspaceName, pipeName, alias, responseData string) tinybird.Request {
	return tinybird.Request{
		Pipe: tinybird.Pipe{
			Name:  pipeName,
			Alias: alias,
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
	rs.Add(newRequest("workspace1", "pipe1", "alias1", `{"id":1}`))
	rs.Add(newRequest("workspace2", "pipe2", "alias2", `{"id":2}`))
	rs.Add(newRequest("workspace2", "pipe2", "alias3", `{"id":2}`))

	rs.Execute()

	assert.Len(t, rs, 3)
	assert.Equal(t, rs[0].Pipe.Name, rs.Get("workspace1", "pipe1").Pipe.Name)
	assert.Equal(t, rs[2].Pipe.Name, rs.GetByAlias("workspace2", "pipe2", "alias3").Pipe.Name)
	assert.Empty(t, rs.Get("random-workspace", "random-pipe").Pipe.Name)
	assert.Empty(t, rs.GetByAlias("random-workspace", "random-pipe", "random-alias").Pipe.Name)
	assert.Equal(t, `{"id":1}`, rs[0].Response.Body)
	assert.Equal(t, `{"id":2}`, rs[1].Response.Body)
	assert.GreaterOrEqual(t, rs.Duration(), time.Duration(0).Nanoseconds())
}
