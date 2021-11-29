package main

import (
	"tum/cmd"

	. "github.com/lyderic/tools"
)

func init() {
	err := CheckBinaries("ssh")
	E(err)
}

func main() {
	cmd.Execute()
}
