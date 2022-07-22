package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

/*
	len在这里不会冗余吗？
	参数名在函数前是结构函数的语法：https://www.cnblogs.com/sunylat/p/6384721.html
*/
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("PrevHash: %s\n", block.PrevBlockHash)
		fmt.Printf("CurrHash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

/*
	go语言函数名大小写：小写代表方法只能在当前包使用，是私有方法，大写代表公有方法。
*/
