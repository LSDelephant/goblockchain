package main

import (
    "time"
)

type Block struct {
    Index        int
    Timestamp    string
    Data         string
    PrevHash     string
    Hash         string
    Nonce        int
    Difficulty   int
}

var Blockchain []Block
const Difficulty = 3

func InitBlockchain() {
    genesis := Block{
        Index:      0,
        Timestamp:  time.Now().String(),
        Data:       "Genesis Block",
        PrevHash:   "",
        Difficulty: Difficulty,
    }
    genesis.Hash, genesis.Nonce = ProofOfWork(genesis)
    Blockchain = append(Blockchain, genesis)
}

func AddBlock(data string) Block {
    prevBlock := Blockchain[len(Blockchain)-1]
    newBlock := Block{
        Index:      prevBlock.Index + 1,
        Timestamp:  time.Now().String(),
        Data:       data,
        PrevHash:   prevBlock.Hash,
        Difficulty: Difficulty,
    }
    newBlock.Hash, newBlock.Nonce = ProofOfWork(newBlock)
    Blockchain = append(Blockchain, newBlock)
    return newBlock
}
