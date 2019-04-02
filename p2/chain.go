package p2

import (
	json2 "encoding/json"
	"errors"
	"fmt"
)

type BlockChain struct {
	Chain  map[int32][]Block
	Length int32
}

func (bc *BlockChain) Get(height int32) (block []Block, err error) {

	if height > bc.Length {
		return make([]Block, 0), errors.New("block height does not exist")
	} else {
		return bc.Chain[height], nil
	}
}

func (bc *BlockChain) Insert(block Block) {

	height := block.Header.Height
	blockArray := bc.Chain[height]
	hash := block.Header.Hash
	for _, value := range blockArray {
		if value.Header.Hash == hash {
			return
		}
	}

	blockArray = append(blockArray, block)
}

func (bc *BlockChain) EncodeToJson() (json string, err error) {

	return "", nil
}

func DecodeChainFromJson(jsonChain string) (bc BlockChain, err error) {

	data := make([]BlockChain, 0)
	err = json2.Unmarshal([]byte(jsonChain), &data)
	if err != nil {
		panic(err)
	}

	for i, value := range data {
		fmt.Printf("i: %d\n", i)
		fmt.Printf("value: %v\n", value)
		//block := Block{}
		//block.Header.Height = value.Header.Height
		//block.Header.Time = value.Header.Time
		//block.Header.Hash = value.Header.Hash
		//block.Header.ParentHash = value.Header.ParentHash
		//block.Header.Size = value.Header.Size
		//fmt.Println(block)
	}

	return BlockChain{}, nil
}
