package dto

import (
	"time"
)

type CreateCheckpointRequest struct {
	PhoneNumber string
}

type CreatePassageRequest struct {
	CheckpointID int64
	DocumentID   int64
	Type         int64
	Time         *time.Time
}

type GetPassageRequest struct {
	PassageID int64
}

type GetCheckpointRequest struct {
	CheckpointID int64
}

type ListPassagesRequest struct {
	DocumentID int64
}

type DeletePassageRequest struct {
	PassageID int64
}

type DeleteCheckpointRequest struct {
	CheckpointID int64
}
