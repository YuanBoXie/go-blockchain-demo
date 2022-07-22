package main

import "go-blockchain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Alice")
	bc.SendData("Send 1 EOS to Bob")
	bc.Print()
}
