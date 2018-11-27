package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chanchailee/grpc-blockchain/proto"
	"google.golang.org/grpc"
)

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot dial server: %+v", err)
	}

	client := proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock(client)
	}

	if *listFlag {
		getBlockchain(client)
	}
}

func addBlock(client proto.BlockchainClient) {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("Unable to add block : %+v", err)
	}

	log.Printf("new block hash : %s\n", block.Hash)
}

func getBlockchain(client proto.BlockchainClient) {

	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})

	if err != nil {
		log.Fatalf("Unable to get blockchain: %+v", err)
	}

	log.Printf("Blockchain Lists:\n")

	for _, b := range bc.Blocks {
		// log.Printf("hash: %s, prev block hash: %s, data : %s\n", b.Hash, b.PrevBlockHash, b.Data)
		log.Printf("\nBlock Struct: %+v\n", b)
	}
}
