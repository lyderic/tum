package cmd

import "os"

const (
	VERSION = "0.2.3"
	APPNAME = "tum"
)

var (
	ssh         string
	CONFIG_FILE = os.Getenv("HOME") + "/.tum.config"
)
