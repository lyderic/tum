/* Class Tunnel */

package cmd

import (
	"os"

	. "github.com/lyderic/tools"
	"gopkg.in/yaml.v2"
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
