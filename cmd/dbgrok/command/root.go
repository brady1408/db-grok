package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command contains a cobra.Command
type Command = cobra.Command

// Run is the main or root command called in main.go
func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

// RootCmd setup the rood command struct
var RootCmd = &cobra.Command{
	Use:   "dbgrok",
	Short: "Open source, db to go tool",
	Long:  "DBGrok is an open source tool to take your database and convert it to go files.",
}

func init() {
	RootCmd.PersistentFlags().String("url", "", "Configure the database connection url.")
	RootCmd.PersistentFlags().StringP("config", "c", "dbgrok.json", "JSON file containing a persistent config")

	viper.SetEnvPrefix("dbg")
	viper.BindEnv("config")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
