package main

import(
  "fmt"
  "os"
  "os/exec"
  "log"
)

func listTunnels(tunnels []Tunnel, ids []int) {
  fmt.Println("ID  ON  Host         Lport   Rport   Description")
  fmt.Println("--  --  ----         -----   -----   -----------")
  for _,tunnel := range(tunnels) {
    listTunnel(tunnel)
  }
}

var listFormat = "%02d  %s   %-10.10s   %5d   %5d   %s\n"

func listTunnel(tunnel Tunnel) {
  ON := "N"
  if checkSocket(tunnel) { ON = "Y" }
  fmt.Printf(listFormat,
    tunnel.Id,
    ON,
    tunnel.Host,
    tunnel.LocalPort,
    tunnel.RemotePort,
    tunnel.Description)
}

func openTunnels(tunnels []Tunnel, ids []int) {
  for _,tunnel := range(tunnels) {
    openTunnel(tunnel)
  }
}

func openTunnel(tunnel Tunnel) {
  if checkSocket(tunnel) {
    fmt.Printf("Tunnel id %d (%s) is already open.\n", tunnel.Id, tunnel.Description)
    return
  }
  fmt.Printf("Opening tunnel id '%d'... ", tunnel.Id)
  forward := fmt.Sprintf("%d:localhost:%d", tunnel.LocalPort, tunnel.RemotePort)
  cmd := exec.Command(ssh, "-f", "-n", "-N", "-M", "-S", getSocket(tunnel),
    "-L", forward, tunnel.Host)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  if debug { log.Printf("\nCommand:", cmd) }
  err := cmd.Run()
  if err != nil { log.Fatal(err) }
  fmt.Println("done.")
}

func closeTunnels(tunnels []Tunnel, ids []int) {
  for _,tunnel := range(tunnels) {
    closeTunnel(tunnel)
  }
}

func closeTunnel(tunnel Tunnel) {
  socket := getSocket(tunnel)
  if !checkSocket(tunnel) {
    fmt.Printf("Tunnel id %d (%s) is not open.\n", tunnel.Id, tunnel.Description)
    return
  }
  fmt.Printf("Closing tunnel id '%d'... ", tunnel.Id)
  cmd := exec.Command(ssh, "-S", socket, "-O", "exit", tunnel.Host)
  if debug { log.Printf("\nCommand:", cmd) }
  e := cmd.Run()
  if e != nil { log.Fatal(e) }
  fmt.Println("done.")
}

func checkSocket(tunnel Tunnel) bool {
  socketOk := false
  if exists(getSocket(tunnel)) {
    cmd := exec.Command(ssh, "-S", getSocket(tunnel), "-O", "check", tunnel.Host)
    err := cmd.Run()
    if err != nil {
      fmt.Println("Removing invalid socket:", getSocket(tunnel))
      os.Remove(getSocket(tunnel))
      socketOk = false
    } else {
      socketOk = true
    }
  }
  return socketOk
}
