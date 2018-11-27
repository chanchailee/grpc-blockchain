package main

import (
	"context"
	"log"
	"net"

	"github.com/chanchailee/grpc-blockchain/blockchain"
	"github.com/chanchailee/grpc-blockchain/proto"
	"google.golang.org/grpc"
)

// Server :
type Server struct {
	Blockchain *blockchain.Blockchain
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to listen on port 8080 : %+v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

// AddBlock :
func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := blockchain.AddBlock(s.Blockchain, in.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain :
func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)

	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
