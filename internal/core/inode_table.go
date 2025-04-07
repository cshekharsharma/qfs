package core

import "sync"

type InodeTable struct {
	sync.RWMutex
	inodes map[uint64]*Inode
}

func NewInodeTable() *InodeTable {
	return &InodeTable{
		inodes: make(map[uint64]*Inode),
	}
}

func (it *InodeTable) Get(inodeNum uint64) (*Inode, bool) {
	it.RLock()
	defer it.RUnlock()
	inode, ok := it.inodes[inodeNum]
	return inode, ok
}

func (it *InodeTable) Set(inodeNum uint64, inode *Inode) {
	it.Lock()
	defer it.Unlock()
	it.inodes[inodeNum] = inode
}
