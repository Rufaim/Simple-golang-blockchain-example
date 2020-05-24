package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type BlockData interface {
	ToString() string
}

type Block struct {
	Position  uint64
	Timestamp string
	Data      BlockData
	Hash      string
	PrevHash  string
}

type writableBlock struct {
	Position  uint64
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

func calculateBlockHash(b *Block) (string, error) {
	data := string(b.Position) + b.Timestamp + b.Data.ToString() + b.PrevHash
	hash := sha256.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (b *Block) generateHash() error {
	hash, err := calculateBlockHash(b)
	if err != nil {
		return err
	}
	b.Hash = hash
	return nil
}

func (b *Block) IsGenesis() bool {
	return len(b.PrevHash) == 0 && b.Position < 0
}

func (b *Block) isValid(prevBlock *Block) bool {
	if prevBlock.Position+1 != b.Position ||
		prevBlock.Hash != b.PrevHash {
		return false
	}

	if hash, err := calculateBlockHash(b); hash != b.Hash || err != nil {
		return false
	}
	return true
}

func (b *Block) toWritable() *writableBlock {
	return &writableBlock{
		Position:  b.Position,
		Timestamp: b.Timestamp,
		Data:      b.Data.ToString(),
		Hash:      b.Hash,
		PrevHash:  b.PrevHash,
	}
}

func NewBlock(data BlockData, prevBlock *Block) (*Block, error) {
	var err error
	if len(prevBlock.Hash) == 0 {
		err = prevBlock.generateHash()
		if err != nil {
			return nil, err
		}
	}

	block := &Block{
		Position:  prevBlock.Position + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	err = block.generateHash()
	if err != nil {
		return nil, err
	}
	return block, nil
}
