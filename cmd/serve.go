package cmd

import (
	"cliProject/internal/logs"
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		startServe()
		stoppingServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServe() {
	logs.Logger.Error("starting from function serve")
}

func stoppingServe() {
	logs.Logger.Error("stoping from function serve")
}
