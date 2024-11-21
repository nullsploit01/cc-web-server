package cmd

import (
	"os"

	"github.com/nullsploit01/cc-web-server/internal"
	"github.com/spf13/cobra"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "ccws [flags]",
	Short: "A simple web server built with Go",
	Long: `ccws is a lightweight web server built with Go. 
It serves static files and demonstrates handling HTTP requests. 

Examples:
  # Start the server on the default port (8080):
  ccws

  # Start the server on a custom port:
  ccws --port 9090
  ccws -p 9090`,

	Run: func(cmd *cobra.Command, args []string) {
		s := internal.InitServer(port)
		s.StartServer()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "The port to listen on")
}
