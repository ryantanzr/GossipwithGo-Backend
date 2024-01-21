package models

import (
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	pgx "github.com/jackc/pgx/v5"
)

type Post struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ScanIntoPosts(rows *pgx.Rows) ([]*Post, error) {

	posts := []*Post{}
	if err := pgxscan.ScanAll(&posts, *rows); err != nil {
		fmt.Println("Scan row error", err)
		return posts, err
	}

	return posts, nil

}

func ScanIntoPost(rows *pgx.Rows) (*Post, error) {

	post := Post{0, "", "", ""}
	if err := pgxscan.ScanRow(&post, *rows); err != nil {
		fmt.Println("Scan row error", err)
		return &post, err
	}

	return &post, nil

}
