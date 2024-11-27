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
	id := context.Param("id")
	var updatedArticle models.Article
	if err := context.ShouldBindBodyWithJSON(&updatedArticle); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	query := `UPDATE articles SET title=$1, content=$2, author=$3 WHERE id=$4`
	_, err := database.DB.Exec(query, updatedArticle.Title, updatedArticle.Content, updatedArticle.Author, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})

}

// delete article by id
func deleteArticle(context *gin.Context) {
	id := context.Param("id")
	_, err := database.DB.Exec("DELETE FROM articles WHERE id=$1", id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
