package main

import (
  "fmt"
  "time"
)


func main() {
  bc := NewBlockChain()

  bc.AddBlock("acc1 1btc to acc2")
  bc.AddBlock("acc2 5btc to acc1")

  for _, block := range bc.blocks {
    fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Printf("Transaction Time: %s\n", time.Unix(block.Timestamp, 0))
  }
}
