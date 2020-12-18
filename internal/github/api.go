package main

import (
	"context"

	"github.com/google/go-github/v33/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	repos, _, _ := client.Repositories.List(ctx, "uu64", nil)
	for _, repo := range repos {
		println(*repo.Name)
	}
}
