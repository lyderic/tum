package main

import(
  "log"
  "os"
  "os/exec"
  "fmt"
  "path"
)

func getSocket(tunnel Tunnel) string {
  socket := fmt.Sprintf("%s/%s-%02d-socket", appdir, appname, tunnel.Id)
  if debug { log.Println("Socket:", socket) }
  return socket
}

func exists(socket string) bool {
  present := false
  if _,e := os.Stat(socket); e == nil {
    present = true
  }
  return present
}

func checkSSH() {
  path,e := exec.LookPath("ssh")
  if e != nil {
    log.Fatal("ssh executable not found! Aborting...")
  } else {
    ssh = path
  }
  if debug { log.Println("ssh:", ssh) }
}

func initAppDir() {
  if !exists(appdir) {
    fmt.Printf("Application directory (%s) not found! Creating... ", appdir)
    if e := os.MkdirAll(appdir, 0700); e != nil {
      fmt.Println()
      log.Fatalf("mkdir '%s' failed! %s\n", appdir, e)
    }
    fmt.Println("done.")
  }
}

func getRunning(tunnel Tunnel) bool {
  return exists(getSocket(tunnel))
}

func readSocket(tunnel Tunnel) string {
  return "'readSocket': not implemented yet!"
}

func displayVersion() {
  fmt.Printf("%s - v. %s (c) Lyderic Landry, London 2017\n", appname, version)
}

func getAppdir() string {
  return path.Join(os.Getenv("HOME"), "." + appname)
}
