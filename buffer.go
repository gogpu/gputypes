package gputypes

// BufferUsage describes how a buffer can be used.
//
// This is a bit flag type. Combine multiple usages with bitwise OR.
type BufferUsage uint32

const (
	// BufferUsageNone indicates no usage (invalid for most operations).
	BufferUsageNone BufferUsage = 0
	// BufferUsageMapRead allows mapping the buffer for reading.
	BufferUsageMapRead BufferUsage = 1 << iota
	// BufferUsageMapWrite allows mapping the buffer for writing.
	BufferUsageMapWrite
	// BufferUsageCopySrc allows the buffer to be a copy source.
	BufferUsageCopySrc
	// BufferUsageCopyDst allows the buffer to be a copy destination.
	BufferUsageCopyDst
	// BufferUsageIndex allows use as an index buffer.
	BufferUsageIndex
	// BufferUsageVertex allows use as a vertex buffer.
	BufferUsageVertex
	// BufferUsageUniform allows use as a uniform buffer.
	BufferUsageUniform
	// BufferUsageStorage allows use as a storage buffer.
	BufferUsageStorage
	// BufferUsageIndirect allows use for indirect draw/dispatch commands.
	BufferUsageIndirect
	// BufferUsageQueryResolve allows use for query result resolution.
	BufferUsageQueryResolve
)

// bufferUsageAll is a mask of all valid buffer usage flags.
const bufferUsageAll = BufferUsageMapRead | BufferUsageMapWrite |
	BufferUsageCopySrc | BufferUsageCopyDst |
	BufferUsageIndex | BufferUsageVertex |
	BufferUsageUniform | BufferUsageStorage |
	BufferUsageIndirect | BufferUsageQueryResolve

// Contains returns true if the usage includes the given flag.
func (u BufferUsage) Contains(flag BufferUsage) bool {
	return u&flag == flag
}

// ContainsUnknownBits returns true if the usage contains any unknown flags.
func (u BufferUsage) ContainsUnknownBits() bool {
	return u&^bufferUsageAll != 0
}

// BufferDescriptor describes a buffer.
type BufferDescriptor struct {
	// Label is an optional debug label.
	Label string
	// Size is the buffer size in bytes.
	Size uint64
	// Usage describes how the buffer will be used.
	Usage BufferUsage
	// MappedAtCreation indicates if the buffer should be mapped at creation.
	// If true, the buffer must have MapRead or MapWrite usage.
	MappedAtCreation bool
}

// BufferMapState describes the map state of a buffer.
type BufferMapState uint8

const (
	// BufferMapStateUnmapped means the buffer is not mapped.
	BufferMapStateUnmapped BufferMapState = iota
	// BufferMapStatePending means a map operation is pending.
	BufferMapStatePending
	// BufferMapStateMapped means the buffer is currently mapped.
	BufferMapStateMapped
)

// String returns the map state name.
func (s BufferMapState) String() string {
	switch s {
	case BufferMapStateUnmapped:
		return "Unmapped"
	case BufferMapStatePending:
		return "Pending"
	case BufferMapStateMapped:
		return "Mapped"
	default:
		return "Unknown"
	}
}

// MapMode describes the access mode for buffer mapping.
//
// This is a bit flag type.
type MapMode uint8

const (
	// MapModeNone indicates no mapping (invalid for map operations).
	MapModeNone MapMode = 0
	// MapModeRead maps the buffer for reading.
	MapModeRead MapMode = 1 << iota
	// MapModeWrite maps the buffer for writing.
	MapModeWrite
)

// BufferBindingType describes how a buffer is bound in a bind group.
type BufferBindingType uint8

const (
	// BufferBindingTypeUndefined is an undefined binding type (invalid).
	BufferBindingTypeUndefined BufferBindingType = iota
	// BufferBindingTypeUniform binds as a uniform buffer (read-only in shaders).
	BufferBindingTypeUniform
	// BufferBindingTypeStorage binds as a storage buffer (read-write in shaders).
	BufferBindingTypeStorage
	// BufferBindingTypeReadOnlyStorage binds as a read-only storage buffer.
	BufferBindingTypeReadOnlyStorage
)

// String returns the binding type name.
func (t BufferBindingType) String() string {
	switch t {
	case BufferBindingTypeUndefined:
		return "Undefined"
	case BufferBindingTypeUniform:
		return "Uniform"
	case BufferBindingTypeStorage:
		return "Storage"
	case BufferBindingTypeReadOnlyStorage:
		return "ReadOnlyStorage"
	default:
		return "Unknown"
	}
}

// IndexFormat describes the format of index buffer data.
type IndexFormat uint8

const (
	// IndexFormatUint16 uses 16-bit unsigned integers (max 65535 indices).
	IndexFormatUint16 IndexFormat = iota
	// IndexFormatUint32 uses 32-bit unsigned integers.
	IndexFormatUint32
)

// String returns the index format name.
func (f IndexFormat) String() string {
	switch f {
	case IndexFormatUint16:
		return "Uint16"
	case IndexFormatUint32:
		return "Uint32"
	default:
		return "Unknown"
	}
}

// Size returns the byte size of the index format.
func (f IndexFormat) Size() uint32 {
	switch f {
	case IndexFormatUint16:
		return 2
	case IndexFormatUint32:
		return 4
	default:
		return 0
	}
}
