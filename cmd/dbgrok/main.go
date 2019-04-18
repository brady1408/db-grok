package main

import (
	"os"

	"bitbucket.org/atlascloudapp/db-grok/cmd/dbgrok/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
