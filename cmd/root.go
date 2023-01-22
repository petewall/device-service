package cmd

import (
	"fmt"
	"net/http"
	"os"

	. "github.com/petewall/device-service/v2/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "device-service",
	Short: "A service for managing device records",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Connecting to database %s:%d...\n", viper.GetString("db.host"), viper.GetInt("db.port"))
		db := Connect(&DBConfig{
			Host: viper.GetString("db.host"),
			Port: viper.GetInt("db.port"),
		})
		cmd.Println("Connected to database")

		api := &API{
			DB:        db,
			LogOutput: cmd.OutOrStdout(),
		}

		port := viper.GetInt("port")
		cmd.Printf("Listening on port %d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), api.GetHttpHandler())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int("port", 5050, "Port to listen on")
	_ = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	_ = viper.BindEnv("port", "PORT")

	rootCmd.Flags().String("db-host", "", "DB host")
	_ = viper.BindPFlag("db.host", rootCmd.Flags().Lookup("db-host"))
	_ = viper.BindEnv("db.host", "DB_HOST")

	rootCmd.Flags().Int("db-port", 6379, "DB port")
	_ = viper.BindPFlag("db.port", rootCmd.Flags().Lookup("db-port"))
	_ = viper.BindEnv("db.port", "DB_PORT")

	rootCmd.SetOut(rootCmd.OutOrStdout())
}
