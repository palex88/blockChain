package p2

import (
	json2 "encoding/json"
	"errors"
	"strings"
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

	if bc.Chain == nil {
		bc.Chain = make(map[int32][]Block)
	}

	blockArray := bc.Chain[height]
	hash := block.Header.Hash
	for _, value := range blockArray {
		if value.Header.Hash == hash {
			return
		}
	}

	blockArray = append(blockArray, block)
	bc.Chain[height] = blockArray
	if bc.Length < height {
		bc.Length = height
	}
}

func (bc *BlockChain) EncodeToJson() (json string, err error) {

	var blocks []string

	for _, val := range bc.Chain {
		//fmt.Printf("i: %d\n", i)
		//fmt.Printf("val: %v\n", val)
		for _, block := range val {
			jsonBlock, err := block.EncodeToJson()
			if err != nil {
				return "", errors.New("chain cannot be encoded to json")
			}
			blocks = append(blocks, jsonBlock)
		}
	}
	blocksString := "["
	blocksString += strings.Join(blocks, ",")
	blocksString += "]"
	return blocksString, nil
}

func DecodeChainFromJson(jsonChain string) (bc BlockChain, err error) {

	bc = BlockChain{}

	value := make([]Data, 0)
	err = json2.Unmarshal([]byte(jsonChain), &value)
	if err != nil {
		return bc, err
	}

	for _, data := range value {
		block := Block{}
		block.Header.Height = data.Height
		block.Header.Time = data.TimeStamp
		block.Header.Hash = data.Hash
		block.Header.ParentHash = data.ParentHash
		block.Header.Size = data.Size
		block.Value.Initial()
		for key, value := range data.Mpt {
			block.Value.Insert(key, value)
		}
		bc.Insert(block)
	}

	return bc, nil
}
