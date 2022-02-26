package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	. "github.com/lyderic/tools"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:                   "list",
	Aliases:               []string{"ls", "l"},
	DisableFlagsInUseLine: true,
	Short:                 "list tunnels",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	tunnels := loadTunnels()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"ID", "ON", "Host", "Lport", "Rport", "Description"})
	for _, tunnel := range tunnels {
		row := table.Row{
			fmt.Sprintf("%02d", tunnel.Id),
			Ternary(socketIsActive(tunnel), "Y", "N"),
			tunnel.Host,
			tunnel.LocalPort,
			tunnel.RemotePort,
			tunnel.Description,
		}
		t.AppendRow(row)
	}
	t.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
