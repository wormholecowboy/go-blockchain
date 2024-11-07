package blockchain

import "fmt"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
  pow := NewProof(block)
  nonce, hash := pow.Run()
  fmt.Printf("hash from create block: %x\n", hash)

  block.Hash = hash[:]
  block.Nonce = nonce

	return block
}
