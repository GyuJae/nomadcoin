package main

import (
	"github.com/gyujae/nomadcoin/cli"
	"github.com/gyujae/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
	// blockchain.Blockchain().AddBlock("first")
	// blockchain.Blockchain().AddBlock("second")
	// blockchain.Blockchain().AddBlock("thrid")
}
