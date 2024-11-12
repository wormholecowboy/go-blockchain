package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type Blockchain struct {
  LastHash []byte
  Database *badger.DB
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
  var lastHash []byte

  opts := badger.DefaultOptions
  opts.Dir = dbPath
  opts.ValueDir = dbPath
  
  db, err := badger.Open(opts)
  Handle(err)

  err = db.Update(func(txn *badger.Txn) error {
    if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
      fmt.Println("No blockchain found")
      genesis := Genesis()
      fmt.Println("Genesis proved")
      err = txn.Set(genesis.Hash, genesis.Serialize())
      Handle(err)
      err = txn.Set([]byte("lh"), genesis.Hash)

      lastHash = genesis.Hash
      return err
    } else {
      item, err := txn.Get([]byte("lh"))
      Handle(err)
      lastHash, err = item.Value()
      return err
    }
  })

  Handle(err)

  blockchain := Blockchain{lastHash, db}
  return &blockchain
}
