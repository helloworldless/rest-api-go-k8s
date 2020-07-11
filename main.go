package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	DbHost      = "db"
	DbUser      = "postgres-dev"
	DbPassword  = "not-for-prod"
	DbName      = "dev"
	DbMigration = `CREATE TABLE IF NOT EXISTS posts (
						id serial PRIMARY KEY,
						author text NOT NULL,
						content text NOT NULL,
						created_at timestamp with time zone DEFAULT current_timestamp)`
)

type Post struct {
	Author    string    `json:"author" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func GetPosts() ([]Post, error) {
	const query = `SELECT author, content, created_at FROM posts ORDER BY created_at DESC LIMIT 100`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]Post, 0)

	for rows.Next() {
		var author string
		var content string
		var createdAt time.Time
		err := rows.Scan(&author, &content, &createdAt)
		if err != nil {
			return nil, err
		}
		results = append(results, Post{author, content, createdAt})
	}

	return results, nil
}

func SavePost(post Post) error {
	const query = `INSERT INTO posts(author, content, created_at) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, post.Author, post.Content, post.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.Default()
	r.GET("/posts", func(context *gin.Context) {

		posts, err := GetPosts()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Status": "Internal Server Error" + err.Error()})
			return
		}

		context.JSON(http.StatusOK, posts)
	})

	r.POST("/posts", func(context *gin.Context) {
		var body Post
		err := context.BindJSON(&body)
		if err != nil {
			log.Println("Failed to serialize request body", err.Error())
			context.JSON(http.StatusUnprocessableEntity, gin.H{"Status": "Invalid body"})
			return
		}

		body.CreatedAt = time.Now()

		if err := SavePost(body); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Status": "Internal Server Error" + err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"Status": "OK"})
	})

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if _, err := db.Query(DbMigration); err != nil {
		log.Println("failed to run migrations", err.Error())
		return
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

	log.Println("Server started")

}
