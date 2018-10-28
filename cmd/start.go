package cmd

import (
	"../internal/logs"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command",
	Run: func(cmd *cobra.Command, args []string) {
		// Print the logs via logger from internal
		logs.Logger.Println("start called inline")
		// example of many functions which should use the logs...
		start()
		stopping()

	},
}

func init() {
	logs.NewLogger() // FIXME: You get a Logger from function NewLogger(), now you need to use it.
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start() {
	logs.Logger.Error("starting from function start")
}

func stopping() {
	logs.Logger.Error("stoping from function start")
}
