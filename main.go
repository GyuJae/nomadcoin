package main

import (
	"github.com/gyujae/nomadcoin/cli"
	"github.com/gyujae/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
