package cmd

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/lyderic/tools"
)

func socketIsActive(tunnel Tunnel) (ok bool) {
	socket := getSocket(tunnel)
	if PathExists(socket) {
		cmd := exec.Command("ssh", "-S", socket, "-O", "check", tunnel.Host)
		Debug("[XeQ] %v\n", cmd.Args)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Removing invalid socket:", socket)
			os.Remove(getSocket(tunnel))
			ok = false
		} else {
			ok = true
		}
	} else {
		Debug("Socket not found: %q\n", socket)
	}
	return
}

func getSocket(tunnel Tunnel) (socket string) {
	socket = fmt.Sprintf("%s/%s-%02d-socket",
		os.Getenv("XDG_RUNTIME_DIR"), APPNAME, tunnel.Id)
	Debug("socket: %s\n", socket)
	return
}

func exists(socket string) bool {
	return PathExists(socket)
}

func getRunning(tunnel Tunnel) bool {
	return exists(getSocket(tunnel))
}
