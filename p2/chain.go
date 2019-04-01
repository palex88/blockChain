package p2

import (
	json2 "encoding/json"
	"errors"
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

	data := make([]Data, 0)
	err = json2.Unmarshal([]byte(jsonChain), &data)
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		block := Block{}
		block.Header.Height = value.Height
		block.Header.Time = value.TimeStamp
		block.Header.Hash = value.Hash
		block.Header.ParentHash = value.ParentHash
		block.Header.Size = value.Size
	}

	return BlockChain{}, nil
}
