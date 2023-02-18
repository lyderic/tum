/* Class Tunnel */

package cmd

import (
	"fmt"
	"os"

	. "github.com/lyderic/tools"
	"gopkg.in/yaml.v3"
)

type Tunnel struct {
	Id          int    `yaml:"id"`
	Description string `yaml:"description"`
	Host        string `yaml:"host"`
	LocalPort   int    `yaml:"localPort"`
	RemotePort  int    `yaml:"remotePort"`
}

func loadTunnels() (tunnels []Tunnel) {
	raw, err := os.ReadFile(CONFIG_FILE)
	E(err)
	err = yaml.Unmarshal(raw, &tunnels)
	return
}

func (tunnel Tunnel) getSocket() (socket Socket) {
	socket.Path = fmt.Sprintf("%s/%s-%02d-socket",
		os.Getenv("XDG_RUNTIME_DIR"), APPNAME, tunnel.Id)
	socket.Exists = PathExists(socket.Path)
	Debug("socket: %s\n", socket)
	return
}

func (tunnel Tunnel) getSocketPath() (path string) {
	path = fmt.Sprintf("%s/%s-%02d-socket",
		os.Getenv("XDG_RUNTIME_DIR"), APPNAME, tunnel.Id)
	Debug("socket path: %s\n", path)
	return
}

func (tunnel Tunnel) getTheSocket() (socket Socket) {
	socket.Path = fmt.Sprintf("%s/%s-%02d-socket",
		os.Getenv("XDG_RUNTIME_DIR"), APPNAME, tunnel.Id)
	Debug("socket path: %s\n", socket.Path)
	return
}
