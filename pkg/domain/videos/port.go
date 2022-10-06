package videos

import (
	"context"

	"github.com/ekszuki/graphhql-server/graph/model"
)

type Repository interface {
	Create(ctx context.Context, video *model.NewVideo) (*model.Video, error)
	FindAll(ctx context.Context) ([]*model.Video, error)
}
