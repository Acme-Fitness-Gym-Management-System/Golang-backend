// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: classcatalogue.sql

package db

import (
	"context"
)

const createClassCatalogue = `-- name: CreateClassCatalogue :one
INSERT INTO classcatalogue (courseid,userid)
VALUES ($1,$2) RETURNING id, userid, courseid
`

type CreateClassCatalogueParams struct {
	Courseid int64 `json:"courseid"`
	Userid   int64 `json:"userid"`
}

func (q *Queries) CreateClassCatalogue(ctx context.Context, arg CreateClassCatalogueParams) (Classcatalogue, error) {
	row := q.db.QueryRowContext(ctx, createClassCatalogue, arg.Courseid, arg.Userid)
	var i Classcatalogue
	err := row.Scan(&i.ID, &i.Userid, &i.Courseid)
	return i, err
}

const getClassEnrolment = `-- name: GetClassEnrolment :many
SELECT userid FROM classcatalogue
WHERE courseid = $1
`

func (q *Queries) GetClassEnrolment(ctx context.Context, courseid int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getClassEnrolment, courseid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var userid int64
		if err := rows.Scan(&userid); err != nil {
			return nil, err
		}
		items = append(items, userid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserClass = `-- name: GetUserClass :many
SELECT id, userid, courseid FROM classcatalogue
WHERE userid = $1
`

func (q *Queries) GetUserClass(ctx context.Context, userid int64) ([]Classcatalogue, error) {
	rows, err := q.db.QueryContext(ctx, getUserClass, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Classcatalogue{}
	for rows.Next() {
		var i Classcatalogue
		if err := rows.Scan(&i.ID, &i.Userid, &i.Courseid); err != nil {
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
