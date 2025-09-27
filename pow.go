package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strings"
)

func CalculateHash(block Block, nonce int) string {
    record := fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp, block.Data, block.PrevHash, nonce)
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

func ProofOfWork(block Block) (string, int) {
    nonce := 0
    for {
        hash := CalculateHash(block, nonce)
        if strings.HasPrefix(hash, strings.Repeat("0", block.Difficulty)) {
            return hash, nonce
        }
        nonce++
    }
}
