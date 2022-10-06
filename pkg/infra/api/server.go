package api

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ekszuki/graphhql-server/graph"
	"github.com/ekszuki/graphhql-server/graph/generated"
	"github.com/ekszuki/graphhql-server/pkg/domain/videos"
	"github.com/sirupsen/logrus"
)

type Repositories struct {
	Videos videos.Repository
}

type apiserver struct {
	ctx          context.Context
	server       *http.Server
	Repositories Repositories
}

func NewAPIServer(
	ctx context.Context,
	repositories Repositories,
) *apiserver {
	return &apiserver{
		ctx:          ctx,
		Repositories: repositories,
	}
}

func (s *apiserver) Initialize() {
	logCtx := logrus.WithFields(logrus.Fields{
		"component": "graph/infra/api/server.go",
		"function":  "Initialize",
	})

	defaultPort := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	addr := ":" + defaultPort
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{
				VideoRepo: s.Repositories.Videos,
			},
		}))

	s.server = &http.Server{
		Addr:    addr,
		Handler: srv,
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logCtx.Infof("connect to http://localhost:%s/ for GraphQL playground", port)

	go func() {
		err := http.ListenAndServe(addr, nil)
		if !errors.As(err, &http.ErrServerClosed) {
			logCtx.Errorf("cannot start server api: %v", err)
		}
	}()

}
