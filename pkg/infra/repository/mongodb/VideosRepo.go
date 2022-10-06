package mongodb

import (
	"context"
	"time"

	"github.com/ekszuki/graphhql-server/graph/model"
	"github.com/ekszuki/graphhql-server/pkg/domain/videos"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type videosRepo struct {
	collection *mongo.Collection
}

func NewVideoRepo(collection *mongo.Collection) videos.Repository {
	return &videosRepo{
		collection: collection,
	}
}

func (r *videosRepo) Create(
	ctx context.Context, video *model.NewVideo,
) (*model.Video, error) {
	mVideo := model.Video{
		ID:    uuid.NewString(),
		Title: video.Title,
		URL:   video.URL,
		Author: &model.User{
			ID:   video.UserID,
			Name: video.Name,
		},
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctxTimeout, mVideo)
	if err != nil {
		return nil, err
	}

	return &mVideo, nil
}

func (r *videosRepo) FindAll(ctx context.Context) ([]*model.Video, error) {
	var list []*model.Video

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cur, err := r.collection.Find(ctxTimeout, bson.D{})
	if err != nil {
		return list, err
	}

	for cur.Next(ctxTimeout) {
		var video model.Video
		err = cur.Decode(&video)
		if err != nil {
			return nil, err
		}
		list = append(list, &video)
	}

	return list, nil
}
