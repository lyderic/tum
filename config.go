package main

import(
  "io/ioutil"
  "path"
  "encoding/json"
  "log"
  "os"
  "os/exec"
)

func loadTunnelsDefinitions() []Tunnel{
  var tunnels []Tunnel
  if debug { log.Println("configfile:", getConfigFile()) }
  content,e := ioutil.ReadFile(getConfigFile())
  if e != nil { log.Fatal(e) }
  e = json.Unmarshal(content, &tunnels)
  if e != nil {
    log.Fatalf("Configuration parsing error! %s\n", e)
  }
  if debug { log.Println("Tunnels as in config file:", tunnels) }
  for i := 0 ; i < len(tunnels) ; i++ {
    tunnels[i].Id = i + 1
  }
  if debug { log.Println("Tunnels after setting 'Id' field:", tunnels) }
  return tunnels
}

func mapTunnelDefinitions() map[int]Tunnel {
  m := make(map[int]Tunnel)
  tunnels := loadTunnelsDefinitions()
  for i := 0 ; i < len(tunnels) ; i++ {
    m[i+1] = tunnels[i]
  }
  return m
}

func editConfigfile() {
  cmd := exec.Command("vim", getConfigFile())
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}

func getConfigFile() string {
  return path.Join(appdir, path.Base(appname) + ".json")
}
