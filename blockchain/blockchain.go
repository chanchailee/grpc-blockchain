package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block :
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain :
type Blockchain struct {
	Blocks []*Block
}

// NewBlock :
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	SetHash(block)

	return block
}

// SetHash :
func SetHash(b *Block) {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// AddBlock :
func AddBlock(bc *Blockchain, data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// NewBlockchain :
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock :
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}
