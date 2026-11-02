package gputypes

// Format helpers ported from Rust wgpu-types (wgpu-types/src/texture/format.rs,
// commit 14328f4). Together with BlockCopySize, IsSrgb and the depth/stencil
// predicates in texture.go they form the helper set that texture-copy and
// view-format validation needs.
//
// gputypes has no multi-planar formats (NV12, P010) and no plane aspects, so
// the plane arms of the Rust matches have no counterpart here. The Go zero
// value TextureAspectUndefined is treated like TextureAspectAll, the way
// Rust treats a missing aspect (None) and TextureAspect::All alike.

// BlockDimensions returns the width and height in texels of one texel block
// (https://gpuweb.github.io/gpuweb/#texel-block).
//
// Uncompressed formats have a block of 1x1. BC and ETC2/EAC formats use 4x4
// blocks, ASTC formats the block size in their name. TextureFormatUndefined
// and unknown values return 1x1 as well, so a caller dividing by the result
// never divides by zero.
//
// Reference: Rust wgpu-types TextureFormat::block_dimensions().
func (f TextureFormat) BlockDimensions() (width, height uint32) {
	switch f {
	case TextureFormatBC1RGBAUnorm, TextureFormatBC1RGBAUnormSrgb,
		TextureFormatBC2RGBAUnorm, TextureFormatBC2RGBAUnormSrgb,
		TextureFormatBC3RGBAUnorm, TextureFormatBC3RGBAUnormSrgb,
		TextureFormatBC4RUnorm, TextureFormatBC4RSnorm,
		TextureFormatBC5RGUnorm, TextureFormatBC5RGSnorm,
		TextureFormatBC6HRGBUfloat, TextureFormatBC6HRGBFloat,
		TextureFormatBC7RGBAUnorm, TextureFormatBC7RGBAUnormSrgb,
		TextureFormatETC2RGB8Unorm, TextureFormatETC2RGB8UnormSrgb,
		TextureFormatETC2RGB8A1Unorm, TextureFormatETC2RGB8A1UnormSrgb,
		TextureFormatETC2RGBA8Unorm, TextureFormatETC2RGBA8UnormSrgb,
		TextureFormatEACR11Unorm, TextureFormatEACR11Snorm,
		TextureFormatEACRG11Unorm, TextureFormatEACRG11Snorm,
		TextureFormatASTC4x4Unorm, TextureFormatASTC4x4UnormSrgb:
		return 4, 4
	default:
		if w, h, ok := astcBlockDimensions(f); ok {
			return w, h
		}
		return 1, 1
	}
}

// astcBlockDimensions returns the block size of the ASTC formats other than
// 4x4 (which BlockDimensions handles together with BC and ETC2).
func astcBlockDimensions(f TextureFormat) (width, height uint32, ok bool) {
	switch f {
	case TextureFormatASTC5x4Unorm, TextureFormatASTC5x4UnormSrgb:
		return 5, 4, true
	case TextureFormatASTC5x5Unorm, TextureFormatASTC5x5UnormSrgb:
		return 5, 5, true
	case TextureFormatASTC6x5Unorm, TextureFormatASTC6x5UnormSrgb:
		return 6, 5, true
	case TextureFormatASTC6x6Unorm, TextureFormatASTC6x6UnormSrgb:
		return 6, 6, true
	case TextureFormatASTC8x5Unorm, TextureFormatASTC8x5UnormSrgb:
		return 8, 5, true
	case TextureFormatASTC8x6Unorm, TextureFormatASTC8x6UnormSrgb:
		return 8, 6, true
	case TextureFormatASTC8x8Unorm, TextureFormatASTC8x8UnormSrgb:
		return 8, 8, true
	case TextureFormatASTC10x5Unorm, TextureFormatASTC10x5UnormSrgb:
		return 10, 5, true
	case TextureFormatASTC10x6Unorm, TextureFormatASTC10x6UnormSrgb:
		return 10, 6, true
	case TextureFormatASTC10x8Unorm, TextureFormatASTC10x8UnormSrgb:
		return 10, 8, true
	case TextureFormatASTC10x10Unorm, TextureFormatASTC10x10UnormSrgb:
		return 10, 10, true
	case TextureFormatASTC12x10Unorm, TextureFormatASTC12x10UnormSrgb:
		return 12, 10, true
	case TextureFormatASTC12x12Unorm, TextureFormatASTC12x12UnormSrgb:
		return 12, 12, true
	default:
		return 0, 0, false
	}
}

// IsCompressed returns true for block-compressed formats (BC, ETC2/EAC,
// ASTC): the formats whose texel block is larger than 1x1.
//
// Reference: Rust wgpu-types TextureFormat::is_compressed().
func (f TextureFormat) IsCompressed() bool {
	w, h := f.BlockDimensions()
	return w != 1 || h != 1
}

// BlockCopySizeAspect returns the number of bytes one texel block occupies
// during an image copy when only the given aspect is copied (the texel block
// copy footprint, https://gpuweb.github.io/gpuweb/#texel-block-copy-footprint).
//
// For every format except the combined depth/stencil formats the aspect does
// not matter and the result equals BlockCopySize. For the combined formats:
//
//   - Depth24PlusStencil8: StencilOnly is 1; DepthOnly is not copyable
//   - Depth32FloatStencil8: StencilOnly is 1, DepthOnly is 4
//
// ok is false (and size 0) where Rust returns None: a combined depth/stencil
// format with TextureAspectAll or TextureAspectUndefined, Depth24Plus, and
// Depth24PlusStencil8 with DepthOnly. It is also false for
// TextureFormatUndefined and unknown values.
//
// Reference: Rust wgpu-types TextureFormat::block_copy_size(Some(aspect)).
func (f TextureFormat) BlockCopySizeAspect(aspect TextureAspect) (size uint32, ok bool) {
	switch f {
	case TextureFormatDepth24PlusStencil8:
		if aspect == TextureAspectStencilOnly {
			return 1, true
		}
		return 0, false
	case TextureFormatDepth32FloatStencil8:
		switch aspect {
		case TextureAspectStencilOnly:
			return 1, true
		case TextureAspectDepthOnly:
			return 4, true
		default:
			return 0, false
		}
	default:
		size = f.BlockCopySize()
		return size, size != 0
	}
}

// SampleType returns the sample type a texture of this format has when
// sampled through the given aspect. features are the device features: only
// FeatureFloat32Filterable matters, it makes the 32-bit float formats
// filterable. Pass the zero Features for a device without features.
//
// ok is false (and the type TextureSampleTypeUndefined) where Rust returns
// None: a combined depth/stencil format with TextureAspectAll or
// TextureAspectUndefined, TextureFormatUndefined and unknown values.
//
// Reference: Rust wgpu-types TextureFormat::sample_type().
func (f TextureFormat) SampleType(aspect TextureAspect, features Features) (TextureSampleType, bool) {
	switch f {
	case TextureFormatR32Float, TextureFormatRG32Float, TextureFormatRGBA32Float:
		if features.Contains(FeatureFloat32Filterable) {
			return TextureSampleTypeFloat, true
		}
		return TextureSampleTypeUnfilterableFloat, true
	case TextureFormatStencil8:
		return TextureSampleTypeUint, true
	case TextureFormatDepth16Unorm, TextureFormatDepth24Plus, TextureFormatDepth32Float:
		return TextureSampleTypeDepth, true
	case TextureFormatDepth24PlusStencil8, TextureFormatDepth32FloatStencil8:
		switch aspect {
		case TextureAspectDepthOnly:
			return TextureSampleTypeDepth, true
		case TextureAspectStencilOnly:
			return TextureSampleTypeUint, true
		default:
			return TextureSampleTypeUndefined, false
		}
	}
	return f.colorSampleType()
}

// colorSampleType is the aspect-independent part of SampleType for the color
// formats other than the 32-bit floats.
func (f TextureFormat) colorSampleType() (TextureSampleType, bool) {
	switch f {
	case TextureFormatR8Uint, TextureFormatRG8Uint, TextureFormatRGBA8Uint,
		TextureFormatR16Uint, TextureFormatRG16Uint, TextureFormatRGBA16Uint,
		TextureFormatR32Uint, TextureFormatRG32Uint, TextureFormatRGBA32Uint,
		TextureFormatRGB10A2Uint:
		return TextureSampleTypeUint, true
	case TextureFormatR8Sint, TextureFormatRG8Sint, TextureFormatRGBA8Sint,
		TextureFormatR16Sint, TextureFormatRG16Sint, TextureFormatRGBA16Sint,
		TextureFormatR32Sint, TextureFormatRG32Sint, TextureFormatRGBA32Sint:
		return TextureSampleTypeSint, true
	}
	if !f.isKnown() {
		return TextureSampleTypeUndefined, false
	}
	// Every other defined format (normalized, 16-bit float, packed float,
	// BC, ETC2/EAC, ASTC) samples as filterable float.
	return TextureSampleTypeFloat, true
}

// Components returns the number of components of this format. For combined
// depth/stencil formats it counts both aspects; use ComponentsWithAspect for
// a single aspect.
//
// Returns 0 for TextureFormatUndefined and unknown values.
//
// Reference: Rust wgpu-types TextureFormat::components().
func (f TextureFormat) Components() uint8 {
	return f.ComponentsWithAspect(TextureAspectAll)
}

// ComponentsWithAspect returns the number of components of this format when
// only the given aspect is considered. The aspect matters only for the
// combined depth/stencil formats: DepthOnly and StencilOnly have one
// component, All and Undefined have two.
//
// Returns 0 for TextureFormatUndefined and unknown values.
//
// Reference: Rust wgpu-types TextureFormat::components_with_aspect().
func (f TextureFormat) ComponentsWithAspect(aspect TextureAspect) uint8 {
	switch f {
	case TextureFormatR8Unorm, TextureFormatR8Snorm, TextureFormatR8Uint, TextureFormatR8Sint,
		TextureFormatR16Unorm, TextureFormatR16Snorm, TextureFormatR16Uint, TextureFormatR16Sint,
		TextureFormatR16Float, TextureFormatR32Uint, TextureFormatR32Sint, TextureFormatR32Float,
		TextureFormatStencil8, TextureFormatDepth16Unorm, TextureFormatDepth24Plus,
		TextureFormatDepth32Float,
		TextureFormatBC4RUnorm, TextureFormatBC4RSnorm,
		TextureFormatEACR11Unorm, TextureFormatEACR11Snorm:
		return 1
	case TextureFormatRG8Unorm, TextureFormatRG8Snorm, TextureFormatRG8Uint, TextureFormatRG8Sint,
		TextureFormatRG16Unorm, TextureFormatRG16Snorm, TextureFormatRG16Uint, TextureFormatRG16Sint,
		TextureFormatRG16Float, TextureFormatRG32Uint, TextureFormatRG32Sint, TextureFormatRG32Float,
		TextureFormatBC5RGUnorm, TextureFormatBC5RGSnorm,
		TextureFormatEACRG11Unorm, TextureFormatEACRG11Snorm:
		return 2
	case TextureFormatRGB9E5Ufloat, TextureFormatRG11B10Ufloat,
		TextureFormatBC6HRGBUfloat, TextureFormatBC6HRGBFloat,
		TextureFormatETC2RGB8Unorm, TextureFormatETC2RGB8UnormSrgb:
		return 3
	case TextureFormatDepth24PlusStencil8, TextureFormatDepth32FloatStencil8:
		if aspect == TextureAspectDepthOnly || aspect == TextureAspectStencilOnly {
			return 1
		}
		return 2
	}
	if !f.isKnown() {
		return 0
	}
	// Every other defined format (RGBA and BGRA variants, RGB10A2, BC1/2/3/7,
	// ETC2 with alpha, ASTC) has four components.
	return 4
}

// isKnown reports whether f is a format this package defines: everything
// except TextureFormatUndefined and values outside the enum.
func (f TextureFormat) isKnown() bool {
	return f != TextureFormatUndefined && f.String() != "Unknown"
}

// RemoveSrgbSuffix changes the *UnormSrgb formats to the matching *Unorm
// format, which performs only linear scaling. All other formats are returned
// unchanged.
//
// Reference: Rust wgpu-types TextureFormat::remove_srgb_suffix().
func (f TextureFormat) RemoveSrgbSuffix() TextureFormat {
	switch f {
	case TextureFormatRGBA8UnormSrgb:
		return TextureFormatRGBA8Unorm
	case TextureFormatBGRA8UnormSrgb:
		return TextureFormatBGRA8Unorm
	case TextureFormatBC1RGBAUnormSrgb:
		return TextureFormatBC1RGBAUnorm
	case TextureFormatBC2RGBAUnormSrgb:
		return TextureFormatBC2RGBAUnorm
	case TextureFormatBC3RGBAUnormSrgb:
		return TextureFormatBC3RGBAUnorm
	case TextureFormatBC7RGBAUnormSrgb:
		return TextureFormatBC7RGBAUnorm
	case TextureFormatETC2RGB8UnormSrgb:
		return TextureFormatETC2RGB8Unorm
	case TextureFormatETC2RGB8A1UnormSrgb:
		return TextureFormatETC2RGB8A1Unorm
	case TextureFormatETC2RGBA8UnormSrgb:
		return TextureFormatETC2RGBA8Unorm
	}
	return f.removeAstcSrgb()
}

// removeAstcSrgb is the ASTC part of RemoveSrgbSuffix. ASTC formats are
// consecutive Unorm/UnormSrgb pairs, but the mapping is spelled out so a
// future renumbering cannot silently break it.
func (f TextureFormat) removeAstcSrgb() TextureFormat {
	switch f {
	case TextureFormatASTC4x4UnormSrgb:
		return TextureFormatASTC4x4Unorm
	case TextureFormatASTC5x4UnormSrgb:
		return TextureFormatASTC5x4Unorm
	case TextureFormatASTC5x5UnormSrgb:
		return TextureFormatASTC5x5Unorm
	case TextureFormatASTC6x5UnormSrgb:
		return TextureFormatASTC6x5Unorm
	case TextureFormatASTC6x6UnormSrgb:
		return TextureFormatASTC6x6Unorm
	case TextureFormatASTC8x5UnormSrgb:
		return TextureFormatASTC8x5Unorm
	case TextureFormatASTC8x6UnormSrgb:
		return TextureFormatASTC8x6Unorm
	case TextureFormatASTC8x8UnormSrgb:
		return TextureFormatASTC8x8Unorm
	case TextureFormatASTC10x5UnormSrgb:
		return TextureFormatASTC10x5Unorm
	case TextureFormatASTC10x6UnormSrgb:
		return TextureFormatASTC10x6Unorm
	case TextureFormatASTC10x8UnormSrgb:
		return TextureFormatASTC10x8Unorm
	case TextureFormatASTC10x10UnormSrgb:
		return TextureFormatASTC10x10Unorm
	case TextureFormatASTC12x10UnormSrgb:
		return TextureFormatASTC12x10Unorm
	case TextureFormatASTC12x12UnormSrgb:
		return TextureFormatASTC12x12Unorm
	}
	return f
}

// AddSrgbSuffix changes the *Unorm formats that have an sRGB variant to the
// matching *UnormSrgb format, which decodes sRGB when read in shaders and
// encodes it when written as a render target. All other formats are returned
// unchanged.
//
// Reference: Rust wgpu-types TextureFormat::add_srgb_suffix().
func (f TextureFormat) AddSrgbSuffix() TextureFormat {
	switch f {
	case TextureFormatRGBA8Unorm:
		return TextureFormatRGBA8UnormSrgb
	case TextureFormatBGRA8Unorm:
		return TextureFormatBGRA8UnormSrgb
	case TextureFormatBC1RGBAUnorm:
		return TextureFormatBC1RGBAUnormSrgb
	case TextureFormatBC2RGBAUnorm:
		return TextureFormatBC2RGBAUnormSrgb
	case TextureFormatBC3RGBAUnorm:
		return TextureFormatBC3RGBAUnormSrgb
	case TextureFormatBC7RGBAUnorm:
		return TextureFormatBC7RGBAUnormSrgb
	case TextureFormatETC2RGB8Unorm:
		return TextureFormatETC2RGB8UnormSrgb
	case TextureFormatETC2RGB8A1Unorm:
		return TextureFormatETC2RGB8A1UnormSrgb
	case TextureFormatETC2RGBA8Unorm:
		return TextureFormatETC2RGBA8UnormSrgb
	}
	return f.addAstcSrgb()
}

// addAstcSrgb is the ASTC part of AddSrgbSuffix.
func (f TextureFormat) addAstcSrgb() TextureFormat {
	switch f {
	case TextureFormatASTC4x4Unorm:
		return TextureFormatASTC4x4UnormSrgb
	case TextureFormatASTC5x4Unorm:
		return TextureFormatASTC5x4UnormSrgb
	case TextureFormatASTC5x5Unorm:
		return TextureFormatASTC5x5UnormSrgb
	case TextureFormatASTC6x5Unorm:
		return TextureFormatASTC6x5UnormSrgb
	case TextureFormatASTC6x6Unorm:
		return TextureFormatASTC6x6UnormSrgb
	case TextureFormatASTC8x5Unorm:
		return TextureFormatASTC8x5UnormSrgb
	case TextureFormatASTC8x6Unorm:
		return TextureFormatASTC8x6UnormSrgb
	case TextureFormatASTC8x8Unorm:
		return TextureFormatASTC8x8UnormSrgb
	case TextureFormatASTC10x5Unorm:
		return TextureFormatASTC10x5UnormSrgb
	case TextureFormatASTC10x6Unorm:
		return TextureFormatASTC10x6UnormSrgb
	case TextureFormatASTC10x8Unorm:
		return TextureFormatASTC10x8UnormSrgb
	case TextureFormatASTC10x10Unorm:
		return TextureFormatASTC10x10UnormSrgb
	case TextureFormatASTC12x10Unorm:
		return TextureFormatASTC12x10UnormSrgb
	case TextureFormatASTC12x12Unorm:
		return TextureFormatASTC12x12UnormSrgb
	}
	return f
}
