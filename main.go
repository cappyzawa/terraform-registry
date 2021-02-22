package main

import (
	"os"

	"github.com/cappyzawa/terraform-registry/internal/command"
)

var (
	version = "dev"
)

func main() {
	o := &command.Opt{
		Name:    os.Args[0],
		Args:    os.Args[1:],
		Version: version,
		Out:     os.Stdout,
		Err:     os.Stderr,
	}
	os.Exit(command.Execute(o))
}
