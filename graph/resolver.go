package graph

import (
	"github.com/ekszuki/graphhql-server/pkg/domain/videos"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	VideoRepo videos.Repository
}
