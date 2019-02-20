package p2

import "github.com/palex88/Merkle-Patricia-Trie/p1"

type Header struct {
	Height     int32
	Time       int64
	Hash       string
	ParentHash string
	Size       int32
}

type Block struct {
	Header Header
	Mpt p1.MerklePatriciaTrie
}

func (block *Block) Initial()  {

}

func (block *Block) DecodeFromJson()  {
	
}

func (block *Block) EncodeToJson() {

}


