package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func printPosts() {
	//fmt.Println("Printing Posts")
	sql := "SELECT title FROM posts"
	rows, _ := db.Query(sql)
	for rows.Next() {
		//fmt.Println("User")
		post := Post{}
		rows.Scan(&(post.Title))
		fmt.Println(post)
	}
	rows.Close()

}
func createPostsTable() /**PostsDataBase*/ {
	db, _ = sql.Open("sqlite3", "posts.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	posts_table := `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        text TEXT NOT NULL);`
	query, err := db.Prepare(posts_table)
	if err != nil {
		fmt.Println(err)
		return
	}
	query.Exec()
	fmt.Println("Table for posts created successfully!")
	query.Close()
	//db.Close()
}
func insertPostinTable(post *Post) {
	db, _ := sql.Open("sqlite3", "posts.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	poststmt, _ := db.Prepare("INSERT INTO posts (title, text) VALUES(?, ?)")
	_, err := poststmt.Exec(post.Title, post.Text)
	if err != nil {
		fmt.Println(err)
		return
	}
	poststmt.Close()
	defer db.Close()
}
func getPostFromTable() *[]Post {
	db, err := sql.Open("sqlite3", "posts.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	row, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer row.Close()
	posts := []Post{}
	for row.Next() { // Iterate and fetch the records from result cursor
		post := Post{}
		row.Scan(&(post.Title), &(post.Text))
		posts = append(posts, post)
		//fmt.Println(users)
	}
	return &posts
}

func userPost(post Post) *Post {
	//fmt.Println(user)
	rows, _ := db.Query("SELECT * FROM posts WHERE title = ? AND text = ?", strings.TrimSpace(post.Title), strings.TrimSpace(post.Text))
	postUser := Post{}
	for rows.Next() { // Iterate and fetch the records from result cursor
		rows.Scan(&(postUser.Id), &(postUser.Title), &(postUser.Text))
		rows.Close()
		return &postUser
	}
	return nil

}
