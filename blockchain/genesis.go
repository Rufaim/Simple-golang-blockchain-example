package blockchain

import "time"

type genesisBlockData struct{}

func (gbd *genesisBlockData) ToString() string {
	return ""
}

func newGenesisBlock() (*Block, error) {
	block := &Block{
		Position:  0,
		Timestamp: time.Now().String(),
		Data:      &genesisBlockData{},
		PrevHash:  "",
	}
	err := block.generateHash()
	if err != nil {
		return nil, err
	}
	return block, nil
}
