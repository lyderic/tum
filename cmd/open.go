package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o", "start"},
	//DisableFlagsInUseLine: true,
	Short: "open tunnel [<tunnel#>]",
	Run: func(cmd *cobra.Command, args []string) {
		actionOnAll("open", args)
	},
}

func openTunnel(tunnel Tunnel) {
	if socketIsActive(tunnel) {
		fmt.Printf("Tunnel id %d (%s) is already open.\n", tunnel.Id, tunnel.Description)
		return
	}
	fmt.Printf("Opening tunnel id '%d'... ", tunnel.Id)
	forward := fmt.Sprintf("%d:localhost:%d", tunnel.LocalPort, tunnel.RemotePort)
	cmd := exec.Command("ssh",
		"-f",                    // Requests ssh to go to background just before command execution
		"-n",                    // Prevents reading from stdin
		"-N",                    // Do not execute a remote command
		"-M",                    // Places the ssh client into “master” mode for connection sharing
		"-T",                    // Disable pseudo-terminal allocation
		"-S", getSocket(tunnel), // Bind to a socket
		"-L", forward, tunnel.Host)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	Debug("\n[XeQ]:%v", cmd.Args)
	err := cmd.Run()
	if err != nil {
		return
	}
	fmt.Println("done.")
}

func init() {
	rootCmd.AddCommand(openCmd)
}