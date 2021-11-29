package cmd

import (
	"os"

	"github.com/spf13/cobra"

	. "github.com/lyderic/tools"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Version:               VERSION,
	Use:                   "tum",
	Short:                 "tum application",
	DisableFlagsInUseLine: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Redln(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("debug", "", false, "show debugging information")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "be verbose")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	Debug("%v\n", viper.AllSettings())
}
