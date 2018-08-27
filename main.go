package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	_, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}
