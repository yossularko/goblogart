package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WithParamsTest(role string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Your handler code goes in here - e.g.
		fmt.Println(role)
	}

	return gin.HandlerFunc(fn)
}
