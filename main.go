package main

import (
	"./cmd"
	"./internal/logs"
)

func main() {
	cmd.Execute()
	logs.Logger.Error("starting")
}
