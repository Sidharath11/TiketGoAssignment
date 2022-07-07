package app

import (
	"university/domain"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnMongoDB()
	domain.ConnPostgress()
	domain.RedisCli()
	router.Run(":8080")
}
