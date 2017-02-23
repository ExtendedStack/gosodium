package argon2

/* Argon2 internal constants */
const (
	// Version of the algorithm
	versionNumber = 0x10

	// Memory block size in bytes
	blockSize     = 1024
	qwordsInBlock = blockSize / 8

	// Number of pseudo-random values generated by one call to Blake
	// in Argon2i to generate reference block positions
	addressesInBlock = 128

	// Pre-hashing digest length and its extension
	prehashDigestLength = 64
	prehashSeedLength   = 72
)

/* Argon2 internal data types */

// Structure for the 1KB memory block implemented as 128 64-bit words.
// Memory blocks can be copied, XORed. Internal words can be accessed by []
// (no bounds checking).
type block [qwordsInBlock]uint64

// Argon2 instance
type instance struct {
	memory        []block // Memory pointer
	passes        uint32  // Number of passes
	memoryBlocks  uint32  // Number of blocks in memory
	segmentLength uint32
	laneLength    uint32
	lanes         uint32
	threads       uint32
	variant       Variant
}

// Argon2 position: where we construct the block right now. Used to
// distribute work between threads.
type position struct {
	pass  uint32
	lane  uint32
	slice uint8
	index uint32
}