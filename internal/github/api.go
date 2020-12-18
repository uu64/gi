package main

import (
	"context"

	"github.com/google/go-github/v33/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	repos, resp, err := client.Repositories.ListByOrg(ctx, "github", nil)
	println(repos)
	println(resp)
	println(err)
}
