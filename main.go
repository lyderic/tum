package main

import(
  "fmt"
  "flag"
  "log"
  "strconv"
)

func init() {
  log.SetFlags(log.Lshortfile)
  checkSSH()
}

func main() {

  debugPtr := flag.Bool("debug", false, "show debugging information")
  appdirPtr := flag.String("appdir", getAppdir(), "application `dir`ectory")
  flag.Usage = func() {
    displayVersion()
    fmt.Printf("Usage: %s <option> <action>\n\n", appname)
    flag.PrintDefaults()
    fmt.Println()
    fmt.Println(`Action is one of:
  [e]dit                 edit tunnel definitions
  [l]ist <id> <id> ...   list tunnels
  [o]pen <id> <id> ...   open tunnels
  [c]lose <id> <id> ...  close tunnels
`)
  }
  flag.Parse()

  debug = *debugPtr
  appdir = *appdirPtr

  if debug { log.Println("flag args:", flag.Args()) }
  if debug { log.Println("appdir:", appdir) }
  if debug { log.Println("appname:", appname) }

  initAppDir()
  allTunnels := loadTunnelsDefinitions()

  action := flag.Arg(0)
  if debug { log.Println("action:", action) }

  var ids []int
  if len(flag.Args()) > 0 {
    for idx,arg := range(flag.Args()) {
      if idx == 0 { continue }
      if id,err := strconv.Atoi(arg); err == nil {
        ids = append(ids, id)
      } else {
        fmt.Printf("%s: not a numeric ID. Ignored...\n", arg)
      }
    }
  }
  if debug { log.Println("IDs:", ids) }

  var tunnels []Tunnel
  if len(ids) == 0 {
    tunnels = allTunnels
  } else {
    for _,id := range(ids) {
      for _,tunnel := range(allTunnels) {
        if tunnel.Id == id {
          tunnels = append(tunnels, tunnel)
        }
      }
    }
  }
  if debug { log.Println("Selected tunnels:", tunnels) }

  switch action {
    case "e", "edit":   editConfigfile()
    case "l", "list":   listTunnels(tunnels, ids)
    case "o", "open":   openTunnels(tunnels, ids)
    case "c", "close":  closeTunnels(tunnels, ids)
    default:            listTunnels(tunnels, ids)
  }
}
