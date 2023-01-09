package cmd

import (
	"go-learning/router"

	"github.com/spf13/cobra"
)

// server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var testType string
var testNum int

// add command
func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&testType, "type", "t", "hackrank", "test type")
	serverCmd.Flags().IntVarP(&testNum, "num", "n", 1, "test number")
}

func start() {
	router.SetRouter(testType, testNum)
}
