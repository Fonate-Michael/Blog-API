package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Comment struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	PostId  int    `json:"post_id"`
	Comment string `json:"comment"`
}

type Like struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}
