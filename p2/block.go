package p2

import (
	"encoding/hex"
	json2 "encoding/json"
	"github.com/palex88/Merkle-Patricia-Trie/p1"
	"golang.org/x/crypto/sha3"
	"time"
)

type Header struct {
	Height     int32
	Time       int64
	Hash       string
	ParentHash string
	Size       int32
}

type Block struct {
	Header Header
	Value  p1.MerklePatriciaTrie
}

type Data struct {
	Height     int32             `json:"height"`
	TimeStamp  int64             `json:"timeStamp"`
	Hash       string            `json:"hash"`
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	Mpt        map[string]string `json:"mpt"`
}

func (block *Block) HashBlock() string {

	// TODO: Update mpt to have capitalized parameters
	hashStr := string(block.Header.Height) + string(block.Header.Time) + block.Header.ParentHash + block.Value.Root + string(block.Header.Size)
	bytes := sha3.Sum256([]byte(hashStr))
	hash := hex.EncodeToString(bytes[:])

	return hash
}

func (block *Block) Initial(height int32, parentHash string, value p1.MerklePatriciaTrie) {

	block.Header.Height = height
	block.Header.ParentHash = parentHash
	block.Value = value
	block.Header.Time = time.Now().Unix()
	block.Header.Size = 0 // TODO: Needs to be updated to the proper value. Maybe check chain for last block height
	block.Header.Hash = block.HashBlock()
}

func DecodeBlockFromJson(jsonBlock string) Block {

	data := make([]Data, 0)
	err := json2.Unmarshal([]byte(jsonBlock), &data)
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
	return Block{}
}

func (block *Block) EncodeToJson() string {
	json, err := json2.Marshal(&block)
	if err != nil {
		panic(err)
	}

	return string(json)
}
