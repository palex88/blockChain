package main

import (
	"fmt"
	"github.com/palex88/blockChain/p2"
)

func main() {

	jsonBlock := "{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}"
	block, err := p2.DecodeBlockFromJson(jsonBlock)
	if err != nil {
		panic(err)
	}
	fmt.Println(block)
	fmt.Printf("Block: %+v\n", block)
	json, err := block.EncodeToJson()
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonBlock)
	fmt.Println(json)

	fmt.Println()

	jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	bc, err := p2.DecodeChainFromJson(jsonBlockChain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d\n", bc.Length)
	fmt.Printf("chain len: %d\n", len(bc.Chain))
	fmt.Printf("chain: %v\n", bc.Chain)
	jsonString, err := bc.EncodeToJson()
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonBlockChain)
	fmt.Println(jsonString)
}