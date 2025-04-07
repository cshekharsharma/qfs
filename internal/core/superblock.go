package core

import "time"

type Superblock struct {
	Version         uint32    // Filesystem version
	MagicNumber     uint32    // Magic number for validation
	BlockSize       uint32    // Size of each block in bytes
	TotalBlocks     uint64    // Total blocks in the filesystem
	FreeBlocks      uint64    // Currently available blocks
	TotalInodes     uint64    // Total inodes possible
	FreeInodes      uint64    // Available inodes
	MountTime       time.Time // Last mount time
	WriteTime       time.Time // Last write time
	MountCount      uint32    // Number of mounts since last check
	MaxMountCount   uint32    // Max allowed mounts before fsck
	UUID            [16]byte  // Unique FS identifier
	VolumeName      string    // Human-readable name
	LastMountedAt   string    // Mount path last used
	FeatureCompat   uint32    // Compatible features (bitmask)
	FeatureIncompat uint32    // Incompatible features (bitmask)
	FeatureReadonly uint32    // Read-only compatible features
}
