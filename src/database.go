package main

import (
	"database/sql"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Database struct {
	path string
	db   *sql.DB
}

func newDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite", "database.sqlite")

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite Database")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id 			text PRIMARY KEY,
			title       text,
			summary     text,
			content     text,
			publishedAt text
    	);
	`)

	if err != nil {
		return nil, err
	}

	return &Database{path, db}, nil
}

func (database *Database) Exec(sql string) (sql.Result, error) {
	result, err := database.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (database *Database) NewPost(title string, summary string, content string, publishedAt string) (string, error) {
	id := uuid.New().String()
	_, err := database.Exec(fmt.Sprintf(
		`INSERT INTO posts (id, title, summary, content, publishedAt) VALUES ("%s", "%s", "%s", "%s", "%s")`,
		id, title, summary, content, publishedAt,
	))

	if err != nil {
		return "err", err
	}

	return id, nil
}

func (database *Database) GetAllPosts() ([]PostSummary, error) {
	rows, err := database.db.Query(`SELECT id, title, summary, publishedAt FROM posts`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []PostSummary
	for rows.Next() {
		var p PostSummary

		if err := rows.Scan(&p.ID, &p.Title, &p.Summary, &p.PublishedAt); err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (database *Database) GetPost(id string, toHTML bool) (*Post, error) {
	postSql, err := database.db.Query(fmt.Sprintf(`SELECT * FROM posts WHERE id="%s"`, id))

	if err != nil {
		return nil, err
	}

	defer postSql.Close()

	var post Post
	postSql.Next()
	err = postSql.Scan(&post.ID, &post.Title, &post.Summary, &post.Content, &post.PublishedAt)

	if err != nil {
		return nil, err
	}

	if toHTML {
		post.Content = string(markdown.ToHTML([]byte(post.Content), nil, nil))
	}

	return &post, nil
}
