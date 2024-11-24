package api

import (
	"github.com/gin-gonic/gin"
)

// CRUD Handlers
func getArticle(context *gin.Context) {
	context.JSON(200, gin.H{"message": "GET article"})
}

func postArticle(context *gin.Context) {
	context.JSON(201, gin.H{"message": "POST article"})
}

func putArticle(context *gin.Context) {
	context.JSON(202, gin.H{"message": "PUT article"})
}

func deleteArticle(context *gin.Context) {
	context.JSON(203, gin.H{"message": "DELETE article"})
}
