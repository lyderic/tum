package cmd

import "os"

const (
	VERSION = "0.3.2"
	APPNAME = "tum"
)

var (
	CONFIG_FILE = os.Getenv("HOME") + "/.tum.config"
)
