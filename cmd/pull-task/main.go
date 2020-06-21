package main

import (
	"github.com/tpeetz/pull-task/cmd"
)

var version = "0.0.1"

func main() {
	cmd.Execute(version)
}
