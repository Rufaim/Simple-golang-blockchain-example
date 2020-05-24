package blockchain

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data BlockData) error {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block, err := NewBlock(data, prevBlock)
	if err != nil {
		return err
	}
	if block.isValid(prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
	return nil
}

func (bc *Blockchain) GetWritiableBlocks() []*writableBlock {
	bcLen := len(bc.blocks)
	wblocks := make([]*writableBlock, bcLen-1)
	for i := 1; i < bcLen; i++ {
		wblocks[i-1] = bc.blocks[i].toWritable()
	}
	return wblocks
}

func NewBlockchain() *Blockchain {
	genesisBlock, err := newGenesisBlock()
	if err != nil {
		panic("Error genesis implementation")
	}
	return &Blockchain{[]*Block{genesisBlock}}
}
