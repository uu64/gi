package main

import (
	"github.com/uu64/gi/lib/cmd"
	"github.com/uu64/gi/lib/config"
)

func main() {
	cmd.Start(config.Get())
}
