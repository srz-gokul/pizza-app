package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the cmd (executable)
var RootCmd = &cobra.Command{
	Use:   "pizza-app",
	Short: "pizza app implementation on Golang",
}

const (
	// postgres
	postgresDBHost            = "POSTGRES_HOST"
	postgresDBHostDefault     = "localhost"
	postgresDBPort            = "POSTGRES_PORT"
	postgresDBPortDefault     = 5432
	postgresDBUserName        = "POSTGRES_USER"
	postgresDBUserNameDefault = "user"
	postgresDBPassword        = "POSTGRES_PASSWORD"
	postgresDBPasswordDefault = "pass"
	postgresDBName            = "POSTGRES_DB"
	postgresDBNameDefault     = "sample"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.BindEnv(PortEvar)
	viper.SetDefault(PortEvar, portDefault)

	viper.BindEnv(postgresDBHost)
	viper.SetDefault(postgresDBHost, postgresDBHostDefault)
	viper.BindEnv(postgresDBPort)
	viper.SetDefault(postgresDBPort, postgresDBPortDefault)
	viper.BindEnv(postgresDBUserName)
	viper.SetDefault(postgresDBUserName, postgresDBUserNameDefault)
	viper.BindEnv(postgresDBPassword)
	viper.SetDefault(postgresDBPassword, postgresDBPasswordDefault)
	viper.BindEnv(postgresDBName)
	viper.SetDefault(postgresDBName, postgresDBNameDefault)
}

// Execute is the entry into the CLI, executing the root CMD.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
