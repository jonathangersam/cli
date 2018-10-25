package main

import (
	"cliProject/cmd"
	"cliProject/internal/logs"
)

func main() {
	cmd.Execute()
	logs.Logger.Error("starting")
}
