package main

import (
	"context"

	"github.com/ekszuki/graphhql-server/pkg/startup"
)

func main() {
	ctx := context.Background()
	sup := startup.NewStartup(ctx)
	sup.Initialize()
}
