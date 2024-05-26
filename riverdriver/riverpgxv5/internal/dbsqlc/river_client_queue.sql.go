// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: river_client_queue.sql

package dbsqlc

import (
	"context"
	"time"
)

const clientQueueCreateOrSetUpdatedAtMany = `-- name: ClientQueueCreateOrSetUpdatedAtMany :one
INSERT INTO river_client_queue (
    metadata,
    name,
    paused_at,
    river_client_id,
    updated_at
) VALUES (
    coalesce($1::jsonb, '{}'),
    unnest($2::text[]),
    coalesce($3::timestamptz, NULL),
    $4,
    coalesce($5::timestamptz, now())
) ON CONFLICT (name) DO UPDATE
SET
    updated_at = coalesce($5::timestamptz, now())
RETURNING river_client_id, name, created_at, max_workers, metadata, num_jobs_completed, num_jobs_running, updated_at
`

type ClientQueueCreateOrSetUpdatedAtManyParams struct {
	Metadata      []byte
	Name          []string
	PausedAt      *time.Time
	RiverClientID string
	UpdatedAt     *time.Time
}

func (q *Queries) ClientQueueCreateOrSetUpdatedAtMany(ctx context.Context, db DBTX, arg *ClientQueueCreateOrSetUpdatedAtManyParams) (*RiverClientQueue, error) {
	row := db.QueryRow(ctx, clientQueueCreateOrSetUpdatedAtMany,
		arg.Metadata,
		arg.Name,
		arg.PausedAt,
		arg.RiverClientID,
		arg.UpdatedAt,
	)
	var i RiverClientQueue
	err := row.Scan(
		&i.RiverClientID,
		&i.Name,
		&i.CreatedAt,
		&i.MaxWorkers,
		&i.Metadata,
		&i.NumJobsCompleted,
		&i.NumJobsRunning,
		&i.UpdatedAt,
	)
	return &i, err
}