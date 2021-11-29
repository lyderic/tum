package cmd

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/lyderic/tools"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:                   "edit",
	Aliases:               []string{"e", "vim"},
	DisableFlagsInUseLine: true,
	Short:                 "edit",
	Run: func(cmd *cobra.Command, args []string) {
		editConfigfile()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func editConfigfile() {
	if !PathExists(CONFIG_FILE) {
		E(fmt.Errorf("Configuration file not found: %q\n", CONFIG_FILE))
		return
	}
	cmd := exec.Command("vim", CONFIG_FILE)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Run()
	E(err)
}
