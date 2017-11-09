/* Class Tunnel */

package main

import(
  "fmt"
)

type Tunnel struct {
  Id           int
  Description  string  `json:"description"`
  Host         string  `json:"host"`
  LocalPort    int     `json:"localPort"`
  RemotePort   int     `json:"remotePort"`
}

func (t *Tunnel) getId() int {
  return t.Id
}

func (t *Tunnel) toString() string {
  return fmt.Sprintf("%02d %s %d %d '%s'",
    t.Id, t.Host, t.LocalPort, t.RemotePort, t.Description)
}
