package cmd

import (
	"github.com/spf13/cobra"

	svr "github.com/dcwangmit01/goapi/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "Launches the example webserver on https://localhost:10080",
	Run: func(cmd *cobra.Command, args []string) {
		svr.StartServer()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}