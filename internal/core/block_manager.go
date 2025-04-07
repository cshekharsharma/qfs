package core

import "sync"

type BlockManager struct {
	sync.Mutex
	totalBlocks uint64
	freeList    []uint64
}

func NewDataBlockManager(total uint64) *BlockManager {
	free := make([]uint64, total)
	for i := uint64(0); i < total; i++ {
		free[i] = i
	}
	return &BlockManager{
		totalBlocks: total,
		freeList:    free,
	}
}

func (bm *BlockManager) Allocate() (uint64, bool) {
	bm.Lock()
	defer bm.Unlock()
	if len(bm.freeList) == 0 {
		return 0, false
	}
	block := bm.freeList[0]
	bm.freeList = bm.freeList[1:]
	return block, true
}

func (bm *BlockManager) Free(block uint64) {
	bm.Lock()
	defer bm.Unlock()
	bm.freeList = append(bm.freeList, block)
}
