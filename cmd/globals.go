package cmd

import "os"

const (
	VERSION = "0.2.8"
	APPNAME = "tum"
)

var (
	CONFIG_FILE = os.Getenv("HOME") + "/.tum.config"
)
