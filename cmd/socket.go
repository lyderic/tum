package cmd

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/lyderic/tools"
)

type Socket struct {
	Path   string
	Exists bool
}

func socketIsActive(tunnel Tunnel) (ok bool) {
	socket := tunnel.getTheSocket()
	if socket.exists() {
		cmd := exec.Command("ssh", "-S", socket.Path, "-O", "check", tunnel.Host)
		Debug("[XeQ] %v\n", cmd.Args)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Removing invalid socket:", socket)
			os.Remove(socket.Path)
			ok = false
		} else {
			return true
		}
	} else {
		Debug("Socket not found: %q\n", socket)
	}
	return
}

func (socket Socket) exists() bool {
	return PathExists(socket.Path)
}

/*
func exists(socket string) bool {
	return PathExists(socket)
}
*/

func getRunning(tunnel Tunnel) bool {
	return tunnel.getSocket().Exists
}
