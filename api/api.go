// A set of rules and standard for requests and responses between the client and server
package api

import (
	"github.com/gin-gonic/gin"
)

// Write Parameters struct which represents the parameters our API endpoint will take
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define router and their handles. Function definitions are defined in /api/api.go
	// router.<HTTP_METHOD>(<ROUTE_PATH>, <HANDLER_FUNCTION>)
	// URL PATH

	// Collection endpoint
	router.GET("/articles", getArticles)

	// Single item endpoint
	router.GET("/articles/:id", getArticle)
	router.POST("/articles", postArticle)
	router.PUT("/articles/:id", putArticle)
	router.DELETE("/articles/:id", deleteArticle)

	return router
}
