package p2

import (
	"encoding/hex"
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
	Value p1.MerklePatriciaTrie
}

func (block *Block) HashBlock() string {

	// TODO: Update mpt to have capitalized parameters
	hashStr := string(block.Header.Height) + string(block.Header.Time) + block.Header.ParentHash + block.Value.Root + string(block.Header.Size)
	bytes := sha3.Sum256([]byte(hashStr))
	hash := hex.EncodeToString(bytes[:])

	return hash
}

func (block *Block) Initial(height int32, parentHash string, value p1.MerklePatriciaTrie)  {

	block.Header.Height = height
	block.Header.ParentHash = parentHash
	block.Value = value
	block.Header.Time = time.Now().Unix()
	block.Header.Size = 0 // TODO: Needs to be updated to the proper value. Maybe check chain for last block height
	block.Header.Hash = block.HashBlock()
}

func DecodeBlockFromJson(json string) Block {
	return Block{}
}

func (block *Block) EncodeToJson() string {

	return ""
}
