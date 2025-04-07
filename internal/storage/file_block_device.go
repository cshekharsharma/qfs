package storage

import (
	"os"
)

type FileBlockDevice struct {
	file      *os.File
	blockSize uint32
	total     uint64
}

func NewFileBlockDevice(path string, blockSize uint32, totalBlocks uint64) (*FileBlockDevice, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &FileBlockDevice{file, blockSize, totalBlocks}, nil
}
