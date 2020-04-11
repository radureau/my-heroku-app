package middleware

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

const headerNthRequest = "Nth-Request"
const headerNthFrom = "Nth-Request-Metadata"

type counter struct {
	v   int
	mux sync.Mutex
}

func (k *counter) inc() {
	k.mux.Lock()
	k.v++
	k.mux.Unlock()
}

var k = new(counter)

// NthRequest return a gin middleware wich set a response a header
// with the number of times it was called
func NthRequest(metadata ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		k.inc()
		c.Header(headerNthRequest, fmt.Sprintf("%d", k.v))
		if len(metadata) > 0 {
			c.Header(headerNthFrom, fmt.Sprint(metadata...))
		}
		c.Next()
	}
}
