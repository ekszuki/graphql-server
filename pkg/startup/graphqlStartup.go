package startup

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekszuki/graphhql-server/pkg/infra/api"
	"github.com/ekszuki/graphhql-server/pkg/infra/repository/mongodb"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type startup struct {
	ctx context.Context
}

func NewStartup(ctx context.Context) *startup {
	return &startup{
		ctx: ctx,
	}
}

func (s *startup) Initialize() {
	logCtx := log.WithFields(log.Fields{"component": "startup", "function": "Initialize"})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	logCtx.Info("Loading envs...")

	var myEnv map[string]string
	myEnv, err := godotenv.Read("app.env")
	if err != nil {
		logCtx.Fatalf("cannot load env file: %v", err)
	}

	logCtx.Info("Connection on MongoDB...")

	mongoURL := fmt.Sprintf(myEnv["MONGO_URL"], myEnv["MONGO_USER"], myEnv["MONGO_PASS"])
	mongoClient, err := mongo.Connect(
		s.ctx, options.Client().ApplyURI(mongoURL),
	)
	if err != nil {
		logCtx.Errorf("can't connect on MongoDB: %v", err)
	}

	err = mongoClient.Ping(s.ctx, nil)
	if err != nil {
		logCtx.Fatalf("cannot connect on the database: %v", err)
	}

	logCtx.Info("Initializing repositories...")
	repos := getMongoDBRepositories(mongoClient)

	logCtx.Info("Creating GraphQL server instance")
	apiServer := api.NewAPIServer(s.ctx, repos)
	apiServer.Initialize()

	<-c
}

func getMongoDBRepositories(client *mongo.Client) api.Repositories {
	videoCollection := client.Database("video_catalog").Collection("videos")
	repositories := api.Repositories{
		Videos: mongodb.NewVideoRepo(videoCollection),
	}

	return repositories
}
