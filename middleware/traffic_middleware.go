package middleware

import (
	"backend-noted/domain"
	"time"

	"github.com/gin-gonic/gin"
)

func TrafficLogger(repo domain.TrafficRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/stats" {
			c.Next()
			return
		}

		now := time.Now()
		roundedTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
		timestampMilli := roundedTime.UnixMilli()

		method := c.Request.Method
		incGet, incPost, incPut, incDel := 0, 0, 0, 0

		switch method {
		case "GET":
			incGet = 1
		case "POST":
			incPost = 1
		case "PATCH", "PUT":
			incPut = 1
		case "DELETE":
			incDel = 1
		}

		stats := domain.TrafficStat{
			Timestamp: timestampMilli,
			GET:       incGet,
			POST:      incPost,
			PUT:       incPut,
			DELETEReq: incDel,
		}

		_ = repo.UpsertTraffic(&stats, incGet, incPost, incPut, incDel)

		c.Next()
	}
}