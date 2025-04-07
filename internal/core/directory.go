package core

type DirectoryEntry struct {
	Name     string // Filename
	InodeNum uint64 // Linked inode number
	FileType uint8  // Same as inode type
}
