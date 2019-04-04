package p2

import (
	"encoding/hex"
	json2 "encoding/json"
	"fmt"
	"github.com/palex88/Merkle-Patricia-Trie/p1"
	"golang.org/x/crypto/sha3"
	"time"
)

type Block struct {
	Value  p1.MerklePatriciaTrie `json:"mpt"`
	Header struct {
		Height     int32  `json:"height"`
		Time       int64  `json:"timeStamp"`
		Hash       string `json:"hash"`
		ParentHash string `json:"parentHash"'`
		Size       int32  `json:"size"`
	}
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
	block.Header.Hash = block.HashBlock()
	block.Header.Size = int32(len([]byte(fmt.Sprintf("%v", block.Value))))
}

func DecodeBlockFromJson(jsonBlock string) Block {

	data := Data{}
	err := json2.Unmarshal([]byte(jsonBlock), &data)
	if err != nil {
		panic(err)
	}

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

	return block
}

func (block *Block) EncodeToJson() string {

	data := Data{}
	data.Height = block.Header.Height
	data.Hash = block.Header.Hash
	data.Size = block.Header.Size
	data.ParentHash = block.Header.ParentHash
	data.TimeStamp = block.Header.Time
	data.Mpt = block.Value.KVPairs

	json, err := json2.Marshal(&data)
	if err != nil {
		panic(err)
	}

	return string(json)
}
