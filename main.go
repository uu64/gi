package main

import (
	"os"

	"github.com/uu64/gi/lib/cmd"
)

func main() {
	cmd.NewCmd().Start(os.Args[1:])
}
