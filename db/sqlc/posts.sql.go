// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: posts.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO "posts" ("title", "body", "user_id", "status") 
VALUES ($1, $2, $3, $4) RETURNING id, title, body, user_id, status, created_at, likes
`

type CreatePostParams struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int32  `json:"user_id"`
	Status string `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.queryRow(ctx, q.createPostStmt, createPost,
		arg.Title,
		arg.Body,
		arg.UserID,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.Likes,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, title, body, user_id, status, created_at, likes FROM "posts"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.queryRow(ctx, q.getPostStmt, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.Likes,
	)
	return i, err
}

const likePost = `-- name: LikePost :exec
UPDATE "posts" SET "likes" = "likes" + 1
WHERE id = $1 RETURNING id, title, body, user_id, status, created_at, likes
`

func (q *Queries) LikePost(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.likePostStmt, likePost, id)
	return err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, body, user_id, status, created_at, likes FROM "posts"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.query(ctx, q.listPostsStmt, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
