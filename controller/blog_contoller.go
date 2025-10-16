package controller

import (
	"app/db"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {

	row, err := db.DB.Query("SELECT * FROM posts")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query posts"})
		return
	}

	var posts []models.Post

	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Description)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		posts = append(posts, post)
	}

	context.JSON(http.StatusOK, gin.H{"posts": posts})
}

func AddPost(context *gin.Context) {
	userId := context.MustGet("user_id").(int)

	var newPost models.Post

	err := context.BindJSON(&newPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO posts (user_id, title, description) VALUES ($1, $2, $3)", userId, newPost.Title, newPost.Description)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post added successfully"})
}
