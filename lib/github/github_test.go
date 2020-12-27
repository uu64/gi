package github

import (
	"context"
	"reflect"
	"testing"

	"github.com/uu64/gi/lib/gi"
)

func TestListAllContents(t *testing.T) {
	expect := []*gi.Content{
		&gi.Content{
			Path: ".gitmodules",
			Type: gi.CtFile,
		},
		&gi.Content{
			Path: "LICENSE",
			Type: gi.CtFile,
		},
		&gi.Content{
			Path: "README.md",
			Type: gi.CtFile,
		},
		&gi.Content{
			Path: "testdocument.txt",
			Type: gi.CtFile,
		},
		&gi.Content{
			Path: "docs",
			Type: gi.CtDirectory,
		},
		&gi.Content{
			Path: "ghapi-test",
			Type: gi.CtSubmodule,
		},
		&gi.Content{
			Path: "docs/testdocument.txt",
			Type: gi.CtFile,
		},
	}

	gh := New()
	if gh == nil {
		t.Fatal("sum(1,2)shouldbe3,butdoesn'tmatch")
	}

	t.Run("can list all contents.", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		ref := "main"
		path := "/"

		contents := gh.ListAllContents(ctx, owner, repo, ref, path)
		reflect.DeepEqual(expect, contents)
	})
}
