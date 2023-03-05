package main

import (
	"example/patrones/internal/infraestructure/delivery/webapi"
)

func main() {
	// router := gin.Default()
	// router.GET("test", getTest)

	// router.Run("localhost:8080")
	webapi.MapURL()
}

// func getTest(c *gin.Context) {
// 	// c.IndentedJSON(http.StatusOK, nil)
// }
