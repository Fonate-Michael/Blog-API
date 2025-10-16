package controller

import (
	"app/db"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetComment(context *gin.Context) {
	postId := context.Param("id")
	row, err := db.DB.Query("SELECT id, user_id, post_id, comment FROM comments WHERE post_id = $1", postId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comments"})
		return
	}

	var comments []models.Comment

	for row.Next() {
		var comment models.Comment
		err := row.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Comment)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		comments = append(comments, comment)
	}

	context.JSON(http.StatusOK, gin.H{"comments": comments})
}

func AddComment(context *gin.Context) {
	userId := context.MustGet("user_id").(int)
	postId := context.Param("id")

	var newComment models.Comment

	err := context.BindJSON(&newComment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO comments (user_id, post_id, comment) VALUES ($1, $2, $3)", userId, postId, newComment.Comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
