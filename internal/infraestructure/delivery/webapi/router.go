package webapi

import (
	"github.com/gin-gonic/gin"
)

func MapURL() {
	router := gin.Default()
	router.GET("test", buildDependencies().testHandler.GetFactory)
	router.Run("localhost:8080")
}
