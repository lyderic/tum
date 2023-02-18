package cmd

import "os"

const (
	VERSION = "0.3.1"
	APPNAME = "tum"
)

var (
	CONFIG_FILE = os.Getenv("HOME") + "/.tum.config"
)
