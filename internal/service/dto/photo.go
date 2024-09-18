package dto

import "course/internal/model"

type CreatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

type CreatePhotoKeyRequest struct {
	DocumentID *model.DocumentID
	Key        *model.PhotoKey
}

type GetPhotoRequest struct {
	DocumentID int64
}

type UpdatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

type UpdatePhotoKeyRequest struct {
	DocumentID *model.DocumentID
	Key        *model.PhotoKey
}

type DeletePhotoRequest struct {
	DocumentID int64
}
