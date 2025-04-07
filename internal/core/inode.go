package core

import "time"

const (
	InodeTypeFile = iota
	InodeTypeDirectory
	InodeTypeSymlink
)

type Inode struct {
	InodeNumber   uint64            // Unique identifier
	Type          uint8             // Type: file, directory, symlink
	Permissions   uint16            // Unix-style permissions
	UID           uint32            // Owner user ID
	GID           uint32            // Owner group ID
	Size          uint64            // File size in bytes
	CreatedAt     time.Time         // Creation time
	ModifiedAt    time.Time         // Last modification
	AccessedAt    time.Time         // Last access
	LinkCount     uint16            // Number of hard links
	BlockPointers []uint64          // Direct/indirect block pointers
	ExtendedAttrs map[string]string // Optional xattrs
	Flags         uint32            // Immutable, append-only, etc.
}
