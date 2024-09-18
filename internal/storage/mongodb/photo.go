package mongodb

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/mongodb"
)

type photoDataStorageImpl struct {
	*mongodb.MongoDB
}

func NewPhotoDataStorage(db *mongodb.MongoDB) storage.PhotoDataStorage {
	return &photoDataStorageImpl{db}
}

func (p *photoDataStorageImpl) Save(ctx context.Context, request *dto.CreatePhotoRequest) (*model.PhotoKey, error) {
	documentID := strconv.Itoa(int(request.DocumentID))
	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{
		Key:   documentID,
		Value: len(request.Data)},
	})

	objectID, err := p.Bucket.UploadFromStream(
		fmt.Sprintf("%s_%d.jpg", documentID, time.Now().Unix()),
		bytes.NewReader(request.Data),
		uploadOpts,
	)
	if err != nil {
		return nil, err
	}

	return model.ToPhotoKey(objectID.Hex()), nil
}

func (p *photoDataStorageImpl) Get(ctx context.Context, key *model.PhotoKey) ([]byte, error) {
	id, err := primitive.ObjectIDFromHex(key.String())
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	if _, err = p.Bucket.DownloadToStream(id, buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *photoDataStorageImpl) Delete(ctx context.Context, key *model.PhotoKey) error {
	id, err := primitive.ObjectIDFromHex(key.String())
	if err != nil {
		return err
	}

	return p.Bucket.Delete(id)
}
