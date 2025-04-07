package fs

import (
	"github.com/cshekharsharma/qfs/internal/core"
)

type QuantaFS struct {
	MountPath   string
	Debug       bool
	Superblock  *core.Superblock
	InodeTable  *core.InodeTable
	BlockMgr    *core.BlockManager
	BlockDevice core.BlockDevice
}
