package cmd

import (
	"fmt"
	"strconv"

	. "github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func Debug(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		Cyan(format, args...)
	}
}

func Verbose(format string, args ...interface{}) {
	if viper.GetBool("verbose") {
		fmt.Printf(format, args...)
	}
}

func actionOnAll(action string, args []string) {
	tunnels := loadTunnels()
	if len(args) == 0 {
		n := len(tunnels)
		for i := 1; i < n+1; i++ {
			args = append(args, strconv.Itoa(i))
		}
	}
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		E(err)
		for _, tunnel := range tunnels {
			if tunnel.Id == id {
				switch action {
				case "open":
					openTunnel(tunnel)
				case "close":
					closeTunnel(tunnel)
				default:
					E(fmt.Errorf("Invalid action: %s\n", action))
				}
			}
		}
	}
}
