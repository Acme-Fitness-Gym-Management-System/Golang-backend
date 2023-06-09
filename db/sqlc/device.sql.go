// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: device.sql

package db

import (
	"context"
)

const createDevice = `-- name: CreateDevice :one
insert into device  (description ,status)
VALUES
    ($1,$2) RETURNING id, description, status
`

type CreateDeviceParams struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) (Device, error) {
	row := q.db.QueryRowContext(ctx, createDevice, arg.Description, arg.Status)
	var i Device
	err := row.Scan(&i.ID, &i.Description, &i.Status)
	return i, err
}

const getAllDevices = `-- name: GetAllDevices :many
SELECT id, description, status FROM device
`

func (q *Queries) GetAllDevices(ctx context.Context) ([]Device, error) {
	rows, err := q.db.QueryContext(ctx, getAllDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Device{}
	for rows.Next() {
		var i Device
		if err := rows.Scan(&i.ID, &i.Description, &i.Status); err != nil {
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

const getDevice = `-- name: GetDevice :one
SELECT id, description, status FROM device
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDevice(ctx context.Context, id int64) (Device, error) {
	row := q.db.QueryRowContext(ctx, getDevice, id)
	var i Device
	err := row.Scan(&i.ID, &i.Description, &i.Status)
	return i, err
}
