package core

type BlockDevice interface {
	ReadBlock(blockNum uint64) ([]byte, error)
	WriteBlock(blockNum uint64, data []byte) error
	Sync() error
	BlockSize() uint32
	TotalBlocks() uint64
}
