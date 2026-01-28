package gputypes

// AddressMode describes texture coordinate addressing.
//
// Determines how texture coordinates outside [0.0, 1.0] are handled.
type AddressMode uint8

const (
	// AddressModeClampToEdge clamps coordinates to the edge texel.
	AddressModeClampToEdge AddressMode = iota
	// AddressModeRepeat repeats the texture (wrapping).
	AddressModeRepeat
	// AddressModeMirrorRepeat repeats with mirroring.
	AddressModeMirrorRepeat
)

// String returns the address mode name.
func (m AddressMode) String() string {
	switch m {
	case AddressModeClampToEdge:
		return "ClampToEdge"
	case AddressModeRepeat:
		return "Repeat"
	case AddressModeMirrorRepeat:
		return "MirrorRepeat"
	default:
		return "Unknown"
	}
}

// FilterMode describes texture filtering.
//
// Used for magnification and minification filters.
type FilterMode uint8

const (
	// FilterModeNearest uses nearest-neighbor filtering (pixelated).
	FilterModeNearest FilterMode = iota
	// FilterModeLinear uses linear interpolation (smooth).
	FilterModeLinear
)

// String returns the filter mode name.
func (m FilterMode) String() string {
	switch m {
	case FilterModeNearest:
		return "Nearest"
	case FilterModeLinear:
		return "Linear"
	default:
		return "Unknown"
	}
}

// MipmapFilterMode describes mipmap filtering.
type MipmapFilterMode uint8

const (
	// MipmapFilterModeNearest selects the nearest mip level.
	MipmapFilterModeNearest MipmapFilterMode = iota
	// MipmapFilterModeLinear interpolates between mip levels.
	MipmapFilterModeLinear
)

// String returns the mipmap filter mode name.
func (m MipmapFilterMode) String() string {
	switch m {
	case MipmapFilterModeNearest:
		return "Nearest"
	case MipmapFilterModeLinear:
		return "Linear"
	default:
		return "Unknown"
	}
}

// CompareFunction describes a comparison function.
//
// Used for depth testing and sampler comparison.
type CompareFunction uint8

const (
	// CompareFunctionUndefined is undefined (no comparison).
	CompareFunctionUndefined CompareFunction = iota
	// CompareFunctionNever always fails.
	CompareFunctionNever
	// CompareFunctionLess passes if source < destination.
	CompareFunctionLess
	// CompareFunctionEqual passes if source == destination.
	CompareFunctionEqual
	// CompareFunctionLessEqual passes if source <= destination.
	CompareFunctionLessEqual
	// CompareFunctionGreater passes if source > destination.
	CompareFunctionGreater
	// CompareFunctionNotEqual passes if source != destination.
	CompareFunctionNotEqual
	// CompareFunctionGreaterEqual passes if source >= destination.
	CompareFunctionGreaterEqual
	// CompareFunctionAlways always passes.
	CompareFunctionAlways
)

// String returns the compare function name.
func (f CompareFunction) String() string {
	switch f {
	case CompareFunctionUndefined:
		return "Undefined"
	case CompareFunctionNever:
		return "Never"
	case CompareFunctionLess:
		return "Less"
	case CompareFunctionEqual:
		return "Equal"
	case CompareFunctionLessEqual:
		return "LessEqual"
	case CompareFunctionGreater:
		return "Greater"
	case CompareFunctionNotEqual:
		return "NotEqual"
	case CompareFunctionGreaterEqual:
		return "GreaterEqual"
	case CompareFunctionAlways:
		return "Always"
	default:
		return "Unknown"
	}
}

// SamplerDescriptor describes a sampler.
type SamplerDescriptor struct {
	// Label is an optional debug label.
	Label string
	// AddressModeU is the U (X) coordinate addressing mode.
	AddressModeU AddressMode
	// AddressModeV is the V (Y) coordinate addressing mode.
	AddressModeV AddressMode
	// AddressModeW is the W (Z) coordinate addressing mode.
	AddressModeW AddressMode
	// MagFilter is the magnification filter (texture appears larger).
	MagFilter FilterMode
	// MinFilter is the minification filter (texture appears smaller).
	MinFilter FilterMode
	// MipmapFilter is the mipmap selection filter.
	MipmapFilter MipmapFilterMode
	// LodMinClamp is the minimum level of detail (0.0 = base level).
	LodMinClamp float32
	// LodMaxClamp is the maximum level of detail.
	LodMaxClamp float32
	// Compare is the comparison function for depth sampling (Undefined = none).
	Compare CompareFunction
	// MaxAnisotropy is the maximum anisotropic filtering level (1-16, 1 = disabled).
	MaxAnisotropy uint16
}

// DefaultSamplerDescriptor returns a sampler descriptor with sensible defaults.
//
// The default uses:
//   - ClampToEdge addressing
//   - Nearest filtering
//   - No mipmap filtering
//   - No comparison
//   - No anisotropic filtering
func DefaultSamplerDescriptor() SamplerDescriptor {
	return SamplerDescriptor{
		AddressModeU:  AddressModeClampToEdge,
		AddressModeV:  AddressModeClampToEdge,
		AddressModeW:  AddressModeClampToEdge,
		MagFilter:     FilterModeNearest,
		MinFilter:     FilterModeNearest,
		MipmapFilter:  MipmapFilterModeNearest,
		LodMinClamp:   0.0,
		LodMaxClamp:   32.0,
		Compare:       CompareFunctionUndefined,
		MaxAnisotropy: 1,
	}
}

// LinearSamplerDescriptor returns a sampler descriptor with linear filtering.
//
// Suitable for smooth texture sampling.
func LinearSamplerDescriptor() SamplerDescriptor {
	return SamplerDescriptor{
		AddressModeU:  AddressModeClampToEdge,
		AddressModeV:  AddressModeClampToEdge,
		AddressModeW:  AddressModeClampToEdge,
		MagFilter:     FilterModeLinear,
		MinFilter:     FilterModeLinear,
		MipmapFilter:  MipmapFilterModeLinear,
		LodMinClamp:   0.0,
		LodMaxClamp:   32.0,
		Compare:       CompareFunctionUndefined,
		MaxAnisotropy: 1,
	}
}

// SamplerBindingType describes how a sampler is bound in a bind group.
type SamplerBindingType uint8

const (
	// SamplerBindingTypeUndefined is undefined (invalid).
	SamplerBindingTypeUndefined SamplerBindingType = iota
	// SamplerBindingTypeFiltering supports filtered texture sampling.
	SamplerBindingTypeFiltering
	// SamplerBindingTypeNonFiltering does not support filtering.
	SamplerBindingTypeNonFiltering
	// SamplerBindingTypeComparison is for depth comparison sampling.
	SamplerBindingTypeComparison
)

// String returns the sampler binding type name.
func (t SamplerBindingType) String() string {
	switch t {
	case SamplerBindingTypeUndefined:
		return "Undefined"
	case SamplerBindingTypeFiltering:
		return "Filtering"
	case SamplerBindingTypeNonFiltering:
		return "NonFiltering"
	case SamplerBindingTypeComparison:
		return "Comparison"
	default:
		return "Unknown"
	}
}
