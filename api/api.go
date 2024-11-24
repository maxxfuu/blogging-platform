// A set of rules and standard for requests and responses between the client and server
package api

import (
	"github.com/gin-gonic/gin"
)

//Write Parameters struct which represents the parameters our API endpoint will take

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define router and their handles. Function definitions are defined in /api/api.go
	// router.<HTTP_METHOD>(<ROUTE_PATH>, <HANDLER_FUNCTION>)
	router.GET("/articles", getArticle)
	router.POST("/article", postArticle)
	router.PUT("/article", putArticle)
	router.DELETE("/article", deleteArticle)

	return router
}
