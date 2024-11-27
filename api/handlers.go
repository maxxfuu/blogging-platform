package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blogging-platform/database"
	"blogging-platform/models"
)

// fetch all articles
func getArticle(context *gin.Context) {
	var articles []models.Article

	err := database.DB.Select(&articles, "SELECT * FROM articles")
	if err != nil {
		// Status Code, Server Error: 500
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}
	// Status Code, OK: 200
	context.JSON(http.StatusOK, articles) // return fetched articles
}

// create a new article
func postArticle(context *gin.Context) {
	var newArticle models.Article

	err1 := context.ShouldBindJSON(&newArticle)
	if err1 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Server unable to process a request due to client error"})
		return
	}

	query := `INSERT INTO articles (title, content, author, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
	err2 := database.DB.QueryRow(query, newArticle.Title, newArticle.Content, newArticle.Author).Scan(&newArticle.ID)
	if err2 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}
	context.JSON(http.StatusCreated, newArticle)
}

// update an existing article
func putArticle(context *gin.Context) {

}

// delete article by id
func deleteArticle(context *gin.Context) {

}
