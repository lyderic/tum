package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:     "close",
	Aliases: []string{"c", "stop"},
	//DisableFlagsInUseLine: true,
	Short: "close tunnel [<tunnel#>]",
	Run: func(cmd *cobra.Command, args []string) {
		actionOnAll("close", args)
	},
}

func closeTunnel(tunnel Tunnel) {
	socket := getSocket(tunnel)
	if !socketIsActive(tunnel) {
		fmt.Printf("Tunnel id %d (%s) is not open.\n", tunnel.Id, tunnel.Description)
		return
	}
	fmt.Printf("Closing tunnel id '%d'... ", tunnel.Id)
	cmd := exec.Command("ssh", "-S", socket, "-O", "exit", tunnel.Host)
	Debug("\nCommand: %v", cmd)
	e := cmd.Run()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("done.")
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
