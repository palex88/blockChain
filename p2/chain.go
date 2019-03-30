package p2

type BlockChain struct {
	Chain map[int32] []Block
	Length int32
}

func (chain *BlockChain) Get()  {
	
}

func (chain *BlockChain) Insert()  {
	
}

func (chain *BlockChain) EncodeToJson() (json string, err error) {

	return "", nil
}

func DecodeChainFromJson(json string) (bc BlockChain, err error) {

	return BlockChain{}, nil
}
