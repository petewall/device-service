package cmd

import (
	"fmt"
	"net/http"
	"os"

	. "github.com/petewall/device-service/v2/lib"
	"github.com/spf13/cobra"
)

var port int
var dbConfig *DBConfig = &DBConfig{}

var rootCmd = &cobra.Command{
	Use:   "device-service",
	Short: "A service for managing device records",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := Connect(dbConfig)
		api := &API{
			DB:        db,
			LogOutput: cmd.OutOrStdout(),
		}

		cmd.Printf("Listening on port %d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), api.GetMux())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&port, "port", 5050, "Port to listen on")
	rootCmd.Flags().StringVar(&dbConfig.Host, "db-host", "localhost", "DB host")
	rootCmd.Flags().IntVar(&dbConfig.Port, "db-port", 6379, "DB port")
}
