package gputypes

// TextureFormat describes the format of a texture.
//
// This enum covers all texture formats defined in the WebGPU specification,
// including uncompressed, depth/stencil, and compressed formats.
type TextureFormat uint32

const (
	// TextureFormatUndefined is an undefined format (invalid).
	TextureFormatUndefined TextureFormat = iota

	// 8-bit formats
	// TextureFormatR8Unorm is a single 8-bit normalized unsigned integer.
	TextureFormatR8Unorm
	// TextureFormatR8Snorm is a single 8-bit normalized signed integer.
	TextureFormatR8Snorm
	// TextureFormatR8Uint is a single 8-bit unsigned integer.
	TextureFormatR8Uint
	// TextureFormatR8Sint is a single 8-bit signed integer.
	TextureFormatR8Sint

	// 16-bit formats
	// TextureFormatR16Uint is a single 16-bit unsigned integer.
	TextureFormatR16Uint
	// TextureFormatR16Sint is a single 16-bit signed integer.
	TextureFormatR16Sint
	// TextureFormatR16Float is a single 16-bit float.
	TextureFormatR16Float
	// TextureFormatRG8Unorm is two 8-bit normalized unsigned integers.
	TextureFormatRG8Unorm
	// TextureFormatRG8Snorm is two 8-bit normalized signed integers.
	TextureFormatRG8Snorm
	// TextureFormatRG8Uint is two 8-bit unsigned integers.
	TextureFormatRG8Uint
	// TextureFormatRG8Sint is two 8-bit signed integers.
	TextureFormatRG8Sint

	// 32-bit formats
	// TextureFormatR32Uint is a single 32-bit unsigned integer.
	TextureFormatR32Uint
	// TextureFormatR32Sint is a single 32-bit signed integer.
	TextureFormatR32Sint
	// TextureFormatR32Float is a single 32-bit float.
	TextureFormatR32Float
	// TextureFormatRG16Uint is two 16-bit unsigned integers.
	TextureFormatRG16Uint
	// TextureFormatRG16Sint is two 16-bit signed integers.
	TextureFormatRG16Sint
	// TextureFormatRG16Float is two 16-bit floats.
	TextureFormatRG16Float
	// TextureFormatRGBA8Unorm is four 8-bit normalized unsigned integers.
	TextureFormatRGBA8Unorm
	// TextureFormatRGBA8UnormSrgb is four 8-bit normalized unsigned integers in sRGB.
	TextureFormatRGBA8UnormSrgb
	// TextureFormatRGBA8Snorm is four 8-bit normalized signed integers.
	TextureFormatRGBA8Snorm
	// TextureFormatRGBA8Uint is four 8-bit unsigned integers.
	TextureFormatRGBA8Uint
	// TextureFormatRGBA8Sint is four 8-bit signed integers.
	TextureFormatRGBA8Sint
	// TextureFormatBGRA8Unorm is four 8-bit normalized unsigned integers (BGRA order).
	TextureFormatBGRA8Unorm
	// TextureFormatBGRA8UnormSrgb is four 8-bit normalized unsigned integers in sRGB (BGRA order).
	TextureFormatBGRA8UnormSrgb

	// Packed 32-bit formats
	// TextureFormatRGB9E5Ufloat is a packed 9-9-9-5 unsigned float format.
	TextureFormatRGB9E5Ufloat
	// TextureFormatRGB10A2Uint is a packed 10-10-10-2 unsigned integer format.
	TextureFormatRGB10A2Uint
	// TextureFormatRGB10A2Unorm is a packed 10-10-10-2 normalized unsigned format.
	TextureFormatRGB10A2Unorm
	// TextureFormatRG11B10Ufloat is a packed 11-11-10 unsigned float format.
	TextureFormatRG11B10Ufloat

	// 64-bit formats
	// TextureFormatRG32Uint is two 32-bit unsigned integers.
	TextureFormatRG32Uint
	// TextureFormatRG32Sint is two 32-bit signed integers.
	TextureFormatRG32Sint
	// TextureFormatRG32Float is two 32-bit floats.
	TextureFormatRG32Float
	// TextureFormatRGBA16Uint is four 16-bit unsigned integers.
	TextureFormatRGBA16Uint
	// TextureFormatRGBA16Sint is four 16-bit signed integers.
	TextureFormatRGBA16Sint
	// TextureFormatRGBA16Float is four 16-bit floats.
	TextureFormatRGBA16Float

	// 128-bit formats
	// TextureFormatRGBA32Uint is four 32-bit unsigned integers.
	TextureFormatRGBA32Uint
	// TextureFormatRGBA32Sint is four 32-bit signed integers.
	TextureFormatRGBA32Sint
	// TextureFormatRGBA32Float is four 32-bit floats.
	TextureFormatRGBA32Float

	// Depth/stencil formats
	// TextureFormatStencil8 is an 8-bit stencil format.
	TextureFormatStencil8
	// TextureFormatDepth16Unorm is a 16-bit normalized depth format.
	TextureFormatDepth16Unorm
	// TextureFormatDepth24Plus is a 24-bit depth format (may be 32-bit internally).
	TextureFormatDepth24Plus
	// TextureFormatDepth24PlusStencil8 is a 24-bit depth + 8-bit stencil format.
	TextureFormatDepth24PlusStencil8
	// TextureFormatDepth32Float is a 32-bit float depth format.
	TextureFormatDepth32Float
	// TextureFormatDepth32FloatStencil8 is a 32-bit float depth + 8-bit stencil format.
	TextureFormatDepth32FloatStencil8

	// BC compressed formats (requires TextureCompressionBC feature)
	// TextureFormatBC1RGBAUnorm is BC1 RGBA normalized unsigned format.
	TextureFormatBC1RGBAUnorm
	// TextureFormatBC1RGBAUnormSrgb is BC1 RGBA normalized unsigned sRGB format.
	TextureFormatBC1RGBAUnormSrgb
	// TextureFormatBC2RGBAUnorm is BC2 RGBA normalized unsigned format.
	TextureFormatBC2RGBAUnorm
	// TextureFormatBC2RGBAUnormSrgb is BC2 RGBA normalized unsigned sRGB format.
	TextureFormatBC2RGBAUnormSrgb
	// TextureFormatBC3RGBAUnorm is BC3 RGBA normalized unsigned format.
	TextureFormatBC3RGBAUnorm
	// TextureFormatBC3RGBAUnormSrgb is BC3 RGBA normalized unsigned sRGB format.
	TextureFormatBC3RGBAUnormSrgb
	// TextureFormatBC4RUnorm is BC4 R normalized unsigned format.
	TextureFormatBC4RUnorm
	// TextureFormatBC4RSnorm is BC4 R normalized signed format.
	TextureFormatBC4RSnorm
	// TextureFormatBC5RGUnorm is BC5 RG normalized unsigned format.
	TextureFormatBC5RGUnorm
	// TextureFormatBC5RGSnorm is BC5 RG normalized signed format.
	TextureFormatBC5RGSnorm
	// TextureFormatBC6HRGBUfloat is BC6H RGB unsigned float format.
	TextureFormatBC6HRGBUfloat
	// TextureFormatBC6HRGBFloat is BC6H RGB signed float format.
	TextureFormatBC6HRGBFloat
	// TextureFormatBC7RGBAUnorm is BC7 RGBA normalized unsigned format.
	TextureFormatBC7RGBAUnorm
	// TextureFormatBC7RGBAUnormSrgb is BC7 RGBA normalized unsigned sRGB format.
	TextureFormatBC7RGBAUnormSrgb

	// ETC2 compressed formats (requires TextureCompressionETC2 feature)
	// TextureFormatETC2RGB8Unorm is ETC2 RGB normalized unsigned format.
	TextureFormatETC2RGB8Unorm
	// TextureFormatETC2RGB8UnormSrgb is ETC2 RGB normalized unsigned sRGB format.
	TextureFormatETC2RGB8UnormSrgb
	// TextureFormatETC2RGB8A1Unorm is ETC2 RGB with 1-bit alpha normalized unsigned format.
	TextureFormatETC2RGB8A1Unorm
	// TextureFormatETC2RGB8A1UnormSrgb is ETC2 RGB with 1-bit alpha normalized unsigned sRGB format.
	TextureFormatETC2RGB8A1UnormSrgb
	// TextureFormatETC2RGBA8Unorm is ETC2 RGBA normalized unsigned format.
	TextureFormatETC2RGBA8Unorm
	// TextureFormatETC2RGBA8UnormSrgb is ETC2 RGBA normalized unsigned sRGB format.
	TextureFormatETC2RGBA8UnormSrgb
	// TextureFormatEACR11Unorm is EAC R normalized unsigned format.
	TextureFormatEACR11Unorm
	// TextureFormatEACR11Snorm is EAC R normalized signed format.
	TextureFormatEACR11Snorm
	// TextureFormatEACRG11Unorm is EAC RG normalized unsigned format.
	TextureFormatEACRG11Unorm
	// TextureFormatEACRG11Snorm is EAC RG normalized signed format.
	TextureFormatEACRG11Snorm

	// ASTC compressed formats (requires TextureCompressionASTC feature)
	// TextureFormatASTC4x4Unorm is ASTC 4x4 normalized unsigned format.
	TextureFormatASTC4x4Unorm
	// TextureFormatASTC4x4UnormSrgb is ASTC 4x4 normalized unsigned sRGB format.
	TextureFormatASTC4x4UnormSrgb
	// TextureFormatASTC5x4Unorm is ASTC 5x4 normalized unsigned format.
	TextureFormatASTC5x4Unorm
	// TextureFormatASTC5x4UnormSrgb is ASTC 5x4 normalized unsigned sRGB format.
	TextureFormatASTC5x4UnormSrgb
	// TextureFormatASTC5x5Unorm is ASTC 5x5 normalized unsigned format.
	TextureFormatASTC5x5Unorm
	// TextureFormatASTC5x5UnormSrgb is ASTC 5x5 normalized unsigned sRGB format.
	TextureFormatASTC5x5UnormSrgb
	// TextureFormatASTC6x5Unorm is ASTC 6x5 normalized unsigned format.
	TextureFormatASTC6x5Unorm
	// TextureFormatASTC6x5UnormSrgb is ASTC 6x5 normalized unsigned sRGB format.
	TextureFormatASTC6x5UnormSrgb
	// TextureFormatASTC6x6Unorm is ASTC 6x6 normalized unsigned format.
	TextureFormatASTC6x6Unorm
	// TextureFormatASTC6x6UnormSrgb is ASTC 6x6 normalized unsigned sRGB format.
	TextureFormatASTC6x6UnormSrgb
	// TextureFormatASTC8x5Unorm is ASTC 8x5 normalized unsigned format.
	TextureFormatASTC8x5Unorm
	// TextureFormatASTC8x5UnormSrgb is ASTC 8x5 normalized unsigned sRGB format.
	TextureFormatASTC8x5UnormSrgb
	// TextureFormatASTC8x6Unorm is ASTC 8x6 normalized unsigned format.
	TextureFormatASTC8x6Unorm
	// TextureFormatASTC8x6UnormSrgb is ASTC 8x6 normalized unsigned sRGB format.
	TextureFormatASTC8x6UnormSrgb
	// TextureFormatASTC8x8Unorm is ASTC 8x8 normalized unsigned format.
	TextureFormatASTC8x8Unorm
	// TextureFormatASTC8x8UnormSrgb is ASTC 8x8 normalized unsigned sRGB format.
	TextureFormatASTC8x8UnormSrgb
	// TextureFormatASTC10x5Unorm is ASTC 10x5 normalized unsigned format.
	TextureFormatASTC10x5Unorm
	// TextureFormatASTC10x5UnormSrgb is ASTC 10x5 normalized unsigned sRGB format.
	TextureFormatASTC10x5UnormSrgb
	// TextureFormatASTC10x6Unorm is ASTC 10x6 normalized unsigned format.
	TextureFormatASTC10x6Unorm
	// TextureFormatASTC10x6UnormSrgb is ASTC 10x6 normalized unsigned sRGB format.
	TextureFormatASTC10x6UnormSrgb
	// TextureFormatASTC10x8Unorm is ASTC 10x8 normalized unsigned format.
	TextureFormatASTC10x8Unorm
	// TextureFormatASTC10x8UnormSrgb is ASTC 10x8 normalized unsigned sRGB format.
	TextureFormatASTC10x8UnormSrgb
	// TextureFormatASTC10x10Unorm is ASTC 10x10 normalized unsigned format.
	TextureFormatASTC10x10Unorm
	// TextureFormatASTC10x10UnormSrgb is ASTC 10x10 normalized unsigned sRGB format.
	TextureFormatASTC10x10UnormSrgb
	// TextureFormatASTC12x10Unorm is ASTC 12x10 normalized unsigned format.
	TextureFormatASTC12x10Unorm
	// TextureFormatASTC12x10UnormSrgb is ASTC 12x10 normalized unsigned sRGB format.
	TextureFormatASTC12x10UnormSrgb
	// TextureFormatASTC12x12Unorm is ASTC 12x12 normalized unsigned format.
	TextureFormatASTC12x12Unorm
	// TextureFormatASTC12x12UnormSrgb is ASTC 12x12 normalized unsigned sRGB format.
	TextureFormatASTC12x12UnormSrgb
)

// String returns the name of the texture format.
func (f TextureFormat) String() string {
	switch f {
	case TextureFormatUndefined:
		return "Undefined"
	case TextureFormatR8Unorm:
		return "R8Unorm"
	case TextureFormatR8Snorm:
		return "R8Snorm"
	case TextureFormatR8Uint:
		return "R8Uint"
	case TextureFormatR8Sint:
		return "R8Sint"
	case TextureFormatR16Uint:
		return "R16Uint"
	case TextureFormatR16Sint:
		return "R16Sint"
	case TextureFormatR16Float:
		return "R16Float"
	case TextureFormatRG8Unorm:
		return "RG8Unorm"
	case TextureFormatRG8Snorm:
		return "RG8Snorm"
	case TextureFormatRG8Uint:
		return "RG8Uint"
	case TextureFormatRG8Sint:
		return "RG8Sint"
	case TextureFormatR32Uint:
		return "R32Uint"
	case TextureFormatR32Sint:
		return "R32Sint"
	case TextureFormatR32Float:
		return "R32Float"
	case TextureFormatRG16Uint:
		return "RG16Uint"
	case TextureFormatRG16Sint:
		return "RG16Sint"
	case TextureFormatRG16Float:
		return "RG16Float"
	case TextureFormatRGBA8Unorm:
		return "RGBA8Unorm"
	case TextureFormatRGBA8UnormSrgb:
		return "RGBA8UnormSrgb"
	case TextureFormatRGBA8Snorm:
		return "RGBA8Snorm"
	case TextureFormatRGBA8Uint:
		return "RGBA8Uint"
	case TextureFormatRGBA8Sint:
		return "RGBA8Sint"
	case TextureFormatBGRA8Unorm:
		return "BGRA8Unorm"
	case TextureFormatBGRA8UnormSrgb:
		return "BGRA8UnormSrgb"
	case TextureFormatRGB9E5Ufloat:
		return "RGB9E5Ufloat"
	case TextureFormatRGB10A2Uint:
		return "RGB10A2Uint"
	case TextureFormatRGB10A2Unorm:
		return "RGB10A2Unorm"
	case TextureFormatRG11B10Ufloat:
		return "RG11B10Ufloat"
	case TextureFormatRG32Uint:
		return "RG32Uint"
	case TextureFormatRG32Sint:
		return "RG32Sint"
	case TextureFormatRG32Float:
		return "RG32Float"
	case TextureFormatRGBA16Uint:
		return "RGBA16Uint"
	case TextureFormatRGBA16Sint:
		return "RGBA16Sint"
	case TextureFormatRGBA16Float:
		return "RGBA16Float"
	case TextureFormatRGBA32Uint:
		return "RGBA32Uint"
	case TextureFormatRGBA32Sint:
		return "RGBA32Sint"
	case TextureFormatRGBA32Float:
		return "RGBA32Float"
	case TextureFormatStencil8:
		return "Stencil8"
	case TextureFormatDepth16Unorm:
		return "Depth16Unorm"
	case TextureFormatDepth24Plus:
		return "Depth24Plus"
	case TextureFormatDepth24PlusStencil8:
		return "Depth24PlusStencil8"
	case TextureFormatDepth32Float:
		return "Depth32Float"
	case TextureFormatDepth32FloatStencil8:
		return "Depth32FloatStencil8"
	default:
		return "Unknown"
	}
}

// IsDepthStencil returns true if this is a depth or stencil format.
func (f TextureFormat) IsDepthStencil() bool {
	switch f {
	case TextureFormatStencil8,
		TextureFormatDepth16Unorm,
		TextureFormatDepth24Plus,
		TextureFormatDepth24PlusStencil8,
		TextureFormatDepth32Float,
		TextureFormatDepth32FloatStencil8:
		return true
	default:
		return false
	}
}

// HasDepth returns true if this format has a depth component.
func (f TextureFormat) HasDepth() bool {
	switch f {
	case TextureFormatDepth16Unorm,
		TextureFormatDepth24Plus,
		TextureFormatDepth24PlusStencil8,
		TextureFormatDepth32Float,
		TextureFormatDepth32FloatStencil8:
		return true
	default:
		return false
	}
}

// HasStencil returns true if this format has a stencil component.
func (f TextureFormat) HasStencil() bool {
	switch f {
	case TextureFormatStencil8,
		TextureFormatDepth24PlusStencil8,
		TextureFormatDepth32FloatStencil8:
		return true
	default:
		return false
	}
}

// IsSrgb returns true if this is an sRGB format.
func (f TextureFormat) IsSrgb() bool {
	switch f {
	case TextureFormatRGBA8UnormSrgb,
		TextureFormatBGRA8UnormSrgb,
		TextureFormatBC1RGBAUnormSrgb,
		TextureFormatBC2RGBAUnormSrgb,
		TextureFormatBC3RGBAUnormSrgb,
		TextureFormatBC7RGBAUnormSrgb,
		TextureFormatETC2RGB8UnormSrgb,
		TextureFormatETC2RGB8A1UnormSrgb,
		TextureFormatETC2RGBA8UnormSrgb,
		TextureFormatASTC4x4UnormSrgb,
		TextureFormatASTC5x4UnormSrgb,
		TextureFormatASTC5x5UnormSrgb,
		TextureFormatASTC6x5UnormSrgb,
		TextureFormatASTC6x6UnormSrgb,
		TextureFormatASTC8x5UnormSrgb,
		TextureFormatASTC8x6UnormSrgb,
		TextureFormatASTC8x8UnormSrgb,
		TextureFormatASTC10x5UnormSrgb,
		TextureFormatASTC10x6UnormSrgb,
		TextureFormatASTC10x8UnormSrgb,
		TextureFormatASTC10x10UnormSrgb,
		TextureFormatASTC12x10UnormSrgb,
		TextureFormatASTC12x12UnormSrgb:
		return true
	default:
		return false
	}
}

// TextureDimension describes texture dimensions.
type TextureDimension uint8

const (
	// TextureDimension1D is a 1D texture.
	TextureDimension1D TextureDimension = iota
	// TextureDimension2D is a 2D texture.
	TextureDimension2D
	// TextureDimension3D is a 3D texture.
	TextureDimension3D
)

// String returns the dimension name.
func (d TextureDimension) String() string {
	switch d {
	case TextureDimension1D:
		return "1D"
	case TextureDimension2D:
		return "2D"
	case TextureDimension3D:
		return "3D"
	default:
		return "Unknown"
	}
}

// TextureViewDimension describes a texture view dimension.
type TextureViewDimension uint8

const (
	// TextureViewDimensionUndefined uses the same dimension as the texture.
	TextureViewDimensionUndefined TextureViewDimension = iota
	// TextureViewDimension1D is a 1D texture view.
	TextureViewDimension1D
	// TextureViewDimension2D is a 2D texture view.
	TextureViewDimension2D
	// TextureViewDimension2DArray is a 2D array texture view.
	TextureViewDimension2DArray
	// TextureViewDimensionCube is a cube texture view.
	TextureViewDimensionCube
	// TextureViewDimensionCubeArray is a cube array texture view.
	TextureViewDimensionCubeArray
	// TextureViewDimension3D is a 3D texture view.
	TextureViewDimension3D
)

// String returns the view dimension name.
func (d TextureViewDimension) String() string {
	switch d {
	case TextureViewDimensionUndefined:
		return "Undefined"
	case TextureViewDimension1D:
		return "1D"
	case TextureViewDimension2D:
		return "2D"
	case TextureViewDimension2DArray:
		return "2DArray"
	case TextureViewDimensionCube:
		return "Cube"
	case TextureViewDimensionCubeArray:
		return "CubeArray"
	case TextureViewDimension3D:
		return "3D"
	default:
		return "Unknown"
	}
}

// TextureAspect describes which aspects of a texture to access.
type TextureAspect uint8

const (
	// TextureAspectAll accesses all aspects (default).
	TextureAspectAll TextureAspect = iota
	// TextureAspectStencilOnly accesses only the stencil aspect.
	TextureAspectStencilOnly
	// TextureAspectDepthOnly accesses only the depth aspect.
	TextureAspectDepthOnly
)

// String returns the aspect name.
func (a TextureAspect) String() string {
	switch a {
	case TextureAspectAll:
		return "All"
	case TextureAspectStencilOnly:
		return "StencilOnly"
	case TextureAspectDepthOnly:
		return "DepthOnly"
	default:
		return "Unknown"
	}
}

// TextureUsage describes how a texture can be used.
//
// This is a bit flag type. Combine multiple usages with bitwise OR.
type TextureUsage uint32

const (
	// TextureUsageNone indicates no usage (invalid for most operations).
	TextureUsageNone TextureUsage = 0
	// TextureUsageCopySrc allows the texture to be a copy source.
	TextureUsageCopySrc TextureUsage = 1 << iota
	// TextureUsageCopyDst allows the texture to be a copy destination.
	TextureUsageCopyDst
	// TextureUsageTextureBinding allows the texture to be bound as a sampled texture.
	TextureUsageTextureBinding
	// TextureUsageStorageBinding allows the texture to be bound as a storage texture.
	TextureUsageStorageBinding
	// TextureUsageRenderAttachment allows the texture to be used as a render attachment.
	TextureUsageRenderAttachment
)

// Contains returns true if the usage includes the given flag.
func (u TextureUsage) Contains(flag TextureUsage) bool {
	return u&flag == flag
}

// TextureDescriptor describes a texture.
type TextureDescriptor struct {
	// Label is an optional debug label.
	Label string
	// Size is the texture size.
	Size Extent3D
	// MipLevelCount is the number of mip levels (1 for no mipmapping).
	MipLevelCount uint32
	// SampleCount is the number of samples (1 for non-multisampled).
	SampleCount uint32
	// Dimension is the texture dimension.
	Dimension TextureDimension
	// Format is the texture format.
	Format TextureFormat
	// Usage describes how the texture will be used.
	Usage TextureUsage
	// ViewFormats lists compatible view formats (optional).
	ViewFormats []TextureFormat
}

// TextureViewDescriptor describes a texture view.
type TextureViewDescriptor struct {
	// Label is an optional debug label.
	Label string
	// Format is the view format (defaults to texture format if Undefined).
	Format TextureFormat
	// Dimension is the view dimension (defaults to match texture if Undefined).
	Dimension TextureViewDimension
	// Aspect specifies which aspect to view.
	Aspect TextureAspect
	// BaseMipLevel is the first mip level accessible to the view.
	BaseMipLevel uint32
	// MipLevelCount is the number of mip levels accessible (0 for all remaining).
	MipLevelCount uint32
	// BaseArrayLayer is the first array layer accessible to the view.
	BaseArrayLayer uint32
	// ArrayLayerCount is the number of array layers accessible (0 for all remaining).
	ArrayLayerCount uint32
}

// TextureSampleType describes the sample type of a texture.
type TextureSampleType uint8

const (
	// TextureSampleTypeFloat samples as filterable floating-point.
	TextureSampleTypeFloat TextureSampleType = iota
	// TextureSampleTypeUnfilterableFloat samples as unfilterable floating-point.
	TextureSampleTypeUnfilterableFloat
	// TextureSampleTypeDepth samples as depth comparison.
	TextureSampleTypeDepth
	// TextureSampleTypeSint samples as signed integer.
	TextureSampleTypeSint
	// TextureSampleTypeUint samples as unsigned integer.
	TextureSampleTypeUint
)

// String returns the sample type name.
func (t TextureSampleType) String() string {
	switch t {
	case TextureSampleTypeFloat:
		return "Float"
	case TextureSampleTypeUnfilterableFloat:
		return "UnfilterableFloat"
	case TextureSampleTypeDepth:
		return "Depth"
	case TextureSampleTypeSint:
		return "Sint"
	case TextureSampleTypeUint:
		return "Uint"
	default:
		return "Unknown"
	}
}

// ImageCopyTexture describes a texture copy source or destination.
type ImageCopyTexture struct {
	// Texture is a handle to the texture (implementation-specific).
	Texture uintptr
	// MipLevel is the mip level to copy.
	MipLevel uint32
	// Origin is the origin of the copy region in the texture.
	Origin Origin3D
	// Aspect is the aspect of the texture to copy.
	Aspect TextureAspect
}

// TextureDataLayout describes the layout of texture data in memory.
type TextureDataLayout struct {
	// Offset is the offset in bytes from the start of the data.
	Offset uint64
	// BytesPerRow is the number of bytes per row of texture data.
	// Must be a multiple of 256 for buffer-to-texture copies.
	BytesPerRow uint32
	// RowsPerImage is the number of rows per image for 3D textures.
	RowsPerImage uint32
}
