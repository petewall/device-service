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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db := Connect(dbConfig)
		api := &API{
			DB:        db,
			LogOutput: cmd.OutOrStdout(),
		}

		cmd.Printf("Listening on port %d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), api.GetMux())
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
