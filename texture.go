package gputypes

// TextureFormat describes the format of a texture.
//
// This enum covers all texture formats defined in the WebGPU specification,
// including uncompressed, depth/stencil, and compressed formats.
type TextureFormat uint32

const (
	// TextureFormatUndefined is an undefined format (invalid).
	// webgpu.h: WGPUTextureFormat_Undefined = 0x00000000.
	TextureFormatUndefined TextureFormat = 0x00000000

	// 8-bit formats.

	// TextureFormatR8Unorm is a single 8-bit normalized unsigned integer.
	TextureFormatR8Unorm TextureFormat = 0x00000001
	// TextureFormatR8Snorm is a single 8-bit normalized signed integer.
	TextureFormatR8Snorm TextureFormat = 0x00000002
	// TextureFormatR8Uint is a single 8-bit unsigned integer.
	TextureFormatR8Uint TextureFormat = 0x00000003
	// TextureFormatR8Sint is a single 8-bit signed integer.
	TextureFormatR8Sint TextureFormat = 0x00000004

	// 16-bit formats.

	// TextureFormatR16Unorm is a single 16-bit normalized unsigned integer.
	TextureFormatR16Unorm TextureFormat = 0x00000005
	// TextureFormatR16Snorm is a single 16-bit normalized signed integer.
	TextureFormatR16Snorm TextureFormat = 0x00000006
	// TextureFormatR16Uint is a single 16-bit unsigned integer.
	TextureFormatR16Uint TextureFormat = 0x00000007
	// TextureFormatR16Sint is a single 16-bit signed integer.
	TextureFormatR16Sint TextureFormat = 0x00000008
	// TextureFormatR16Float is a single 16-bit float.
	TextureFormatR16Float TextureFormat = 0x00000009
	// TextureFormatRG8Unorm is two 8-bit normalized unsigned integers.
	TextureFormatRG8Unorm TextureFormat = 0x0000000A
	// TextureFormatRG8Snorm is two 8-bit normalized signed integers.
	TextureFormatRG8Snorm TextureFormat = 0x0000000B
	// TextureFormatRG8Uint is two 8-bit unsigned integers.
	TextureFormatRG8Uint TextureFormat = 0x0000000C
	// TextureFormatRG8Sint is two 8-bit signed integers.
	TextureFormatRG8Sint TextureFormat = 0x0000000D

	// 32-bit formats.

	// TextureFormatR32Float is a single 32-bit float.
	TextureFormatR32Float TextureFormat = 0x0000000E
	// TextureFormatR32Uint is a single 32-bit unsigned integer.
	TextureFormatR32Uint TextureFormat = 0x0000000F
	// TextureFormatR32Sint is a single 32-bit signed integer.
	TextureFormatR32Sint TextureFormat = 0x00000010
	// TextureFormatRG16Unorm is two 16-bit normalized unsigned integers.
	TextureFormatRG16Unorm TextureFormat = 0x00000011
	// TextureFormatRG16Snorm is two 16-bit normalized signed integers.
	TextureFormatRG16Snorm TextureFormat = 0x00000012
	// TextureFormatRG16Uint is two 16-bit unsigned integers.
	TextureFormatRG16Uint TextureFormat = 0x00000013
	// TextureFormatRG16Sint is two 16-bit signed integers.
	TextureFormatRG16Sint TextureFormat = 0x00000014
	// TextureFormatRG16Float is two 16-bit floats.
	TextureFormatRG16Float TextureFormat = 0x00000015
	// TextureFormatRGBA8Unorm is four 8-bit normalized unsigned integers.
	TextureFormatRGBA8Unorm TextureFormat = 0x00000016
	// TextureFormatRGBA8UnormSrgb is four 8-bit normalized unsigned integers in sRGB.
	TextureFormatRGBA8UnormSrgb TextureFormat = 0x00000017
	// TextureFormatRGBA8Snorm is four 8-bit normalized signed integers.
	TextureFormatRGBA8Snorm TextureFormat = 0x00000018
	// TextureFormatRGBA8Uint is four 8-bit unsigned integers.
	TextureFormatRGBA8Uint TextureFormat = 0x00000019
	// TextureFormatRGBA8Sint is four 8-bit signed integers.
	TextureFormatRGBA8Sint TextureFormat = 0x0000001A
	// TextureFormatBGRA8Unorm is four 8-bit normalized unsigned integers (BGRA order).
	TextureFormatBGRA8Unorm TextureFormat = 0x0000001B
	// TextureFormatBGRA8UnormSrgb is four 8-bit normalized unsigned integers in sRGB (BGRA order).
	TextureFormatBGRA8UnormSrgb TextureFormat = 0x0000001C

	// Packed 32-bit formats.

	// TextureFormatRGB10A2Uint is a packed 10-10-10-2 unsigned integer format.
	TextureFormatRGB10A2Uint TextureFormat = 0x0000001D
	// TextureFormatRGB10A2Unorm is a packed 10-10-10-2 normalized unsigned format.
	TextureFormatRGB10A2Unorm TextureFormat = 0x0000001E
	// TextureFormatRG11B10Ufloat is a packed 11-11-10 unsigned float format.
	TextureFormatRG11B10Ufloat TextureFormat = 0x0000001F
	// TextureFormatRGB9E5Ufloat is a packed 9-9-9-5 unsigned float format.
	TextureFormatRGB9E5Ufloat TextureFormat = 0x00000020

	// 64-bit formats.

	// TextureFormatRG32Float is two 32-bit floats.
	TextureFormatRG32Float TextureFormat = 0x00000021
	// TextureFormatRG32Uint is two 32-bit unsigned integers.
	TextureFormatRG32Uint TextureFormat = 0x00000022
	// TextureFormatRG32Sint is two 32-bit signed integers.
	TextureFormatRG32Sint TextureFormat = 0x00000023
	// TextureFormatRGBA16Unorm is four 16-bit normalized unsigned integers.
	TextureFormatRGBA16Unorm TextureFormat = 0x00000024
	// TextureFormatRGBA16Snorm is four 16-bit normalized signed integers.
	TextureFormatRGBA16Snorm TextureFormat = 0x00000025
	// TextureFormatRGBA16Uint is four 16-bit unsigned integers.
	TextureFormatRGBA16Uint TextureFormat = 0x00000026
	// TextureFormatRGBA16Sint is four 16-bit signed integers.
	TextureFormatRGBA16Sint TextureFormat = 0x00000027
	// TextureFormatRGBA16Float is four 16-bit floats.
	TextureFormatRGBA16Float TextureFormat = 0x00000028

	// 128-bit formats.

	// TextureFormatRGBA32Float is four 32-bit floats.
	TextureFormatRGBA32Float TextureFormat = 0x00000029
	// TextureFormatRGBA32Uint is four 32-bit unsigned integers.
	TextureFormatRGBA32Uint TextureFormat = 0x0000002A
	// TextureFormatRGBA32Sint is four 32-bit signed integers.
	TextureFormatRGBA32Sint TextureFormat = 0x0000002B

	// Depth/stencil formats.

	// TextureFormatStencil8 is an 8-bit stencil format.
	TextureFormatStencil8 TextureFormat = 0x0000002C
	// TextureFormatDepth16Unorm is a 16-bit normalized depth format.
	TextureFormatDepth16Unorm TextureFormat = 0x0000002D
	// TextureFormatDepth24Plus is a 24-bit depth format (may be 32-bit internally).
	TextureFormatDepth24Plus TextureFormat = 0x0000002E
	// TextureFormatDepth24PlusStencil8 is a 24-bit depth + 8-bit stencil format.
	TextureFormatDepth24PlusStencil8 TextureFormat = 0x0000002F
	// TextureFormatDepth32Float is a 32-bit float depth format.
	TextureFormatDepth32Float TextureFormat = 0x00000030
	// TextureFormatDepth32FloatStencil8 is a 32-bit float depth + 8-bit stencil format.
	TextureFormatDepth32FloatStencil8 TextureFormat = 0x00000031

	// BC compressed formats (requires TextureCompressionBC feature).
	// webgpu.h: BC formats start at 0x32 (after depth/stencil formats).

	// TextureFormatBC1RGBAUnorm is BC1 RGBA normalized unsigned format.
	TextureFormatBC1RGBAUnorm TextureFormat = 0x00000032
	// TextureFormatBC1RGBAUnormSrgb is BC1 RGBA normalized unsigned sRGB format.
	TextureFormatBC1RGBAUnormSrgb TextureFormat = 0x00000033
	// TextureFormatBC2RGBAUnorm is BC2 RGBA normalized unsigned format.
	TextureFormatBC2RGBAUnorm TextureFormat = 0x00000034
	// TextureFormatBC2RGBAUnormSrgb is BC2 RGBA normalized unsigned sRGB format.
	TextureFormatBC2RGBAUnormSrgb TextureFormat = 0x00000035
	// TextureFormatBC3RGBAUnorm is BC3 RGBA normalized unsigned format.
	TextureFormatBC3RGBAUnorm TextureFormat = 0x00000036
	// TextureFormatBC3RGBAUnormSrgb is BC3 RGBA normalized unsigned sRGB format.
	TextureFormatBC3RGBAUnormSrgb TextureFormat = 0x00000037
	// TextureFormatBC4RUnorm is BC4 R normalized unsigned format.
	TextureFormatBC4RUnorm TextureFormat = 0x00000038
	// TextureFormatBC4RSnorm is BC4 R normalized signed format.
	TextureFormatBC4RSnorm TextureFormat = 0x00000039
	// TextureFormatBC5RGUnorm is BC5 RG normalized unsigned format.
	TextureFormatBC5RGUnorm TextureFormat = 0x0000003A
	// TextureFormatBC5RGSnorm is BC5 RG normalized signed format.
	TextureFormatBC5RGSnorm TextureFormat = 0x0000003B
	// TextureFormatBC6HRGBUfloat is BC6H RGB unsigned float format.
	TextureFormatBC6HRGBUfloat TextureFormat = 0x0000003C
	// TextureFormatBC6HRGBFloat is BC6H RGB signed float format.
	TextureFormatBC6HRGBFloat TextureFormat = 0x0000003D
	// TextureFormatBC7RGBAUnorm is BC7 RGBA normalized unsigned format.
	TextureFormatBC7RGBAUnorm TextureFormat = 0x0000003E
	// TextureFormatBC7RGBAUnormSrgb is BC7 RGBA normalized unsigned sRGB format.
	TextureFormatBC7RGBAUnormSrgb TextureFormat = 0x0000003F

	// ETC2 compressed formats (requires TextureCompressionETC2 feature).
	// webgpu.h: ETC2 formats start at 0x40 (after BC formats).

	// TextureFormatETC2RGB8Unorm is ETC2 RGB normalized unsigned format.
	TextureFormatETC2RGB8Unorm TextureFormat = 0x00000040
	// TextureFormatETC2RGB8UnormSrgb is ETC2 RGB normalized unsigned sRGB format.
	TextureFormatETC2RGB8UnormSrgb TextureFormat = 0x00000041
	// TextureFormatETC2RGB8A1Unorm is ETC2 RGB with 1-bit alpha normalized unsigned format.
	TextureFormatETC2RGB8A1Unorm TextureFormat = 0x00000042
	// TextureFormatETC2RGB8A1UnormSrgb is ETC2 RGB with 1-bit alpha normalized unsigned sRGB format.
	TextureFormatETC2RGB8A1UnormSrgb TextureFormat = 0x00000043
	// TextureFormatETC2RGBA8Unorm is ETC2 RGBA normalized unsigned format.
	TextureFormatETC2RGBA8Unorm TextureFormat = 0x00000044
	// TextureFormatETC2RGBA8UnormSrgb is ETC2 RGBA normalized unsigned sRGB format.
	TextureFormatETC2RGBA8UnormSrgb TextureFormat = 0x00000045
	// TextureFormatEACR11Unorm is EAC R normalized unsigned format.
	TextureFormatEACR11Unorm TextureFormat = 0x00000046
	// TextureFormatEACR11Snorm is EAC R normalized signed format.
	TextureFormatEACR11Snorm TextureFormat = 0x00000047
	// TextureFormatEACRG11Unorm is EAC RG normalized unsigned format.
	TextureFormatEACRG11Unorm TextureFormat = 0x00000048
	// TextureFormatEACRG11Snorm is EAC RG normalized signed format.
	TextureFormatEACRG11Snorm TextureFormat = 0x00000049

	// ASTC compressed formats (requires TextureCompressionASTC feature).
	// webgpu.h: ASTC formats start at 0x4A (after ETC2 formats).

	// TextureFormatASTC4x4Unorm is ASTC 4x4 normalized unsigned format.
	TextureFormatASTC4x4Unorm TextureFormat = 0x0000004A
	// TextureFormatASTC4x4UnormSrgb is ASTC 4x4 normalized unsigned sRGB format.
	TextureFormatASTC4x4UnormSrgb TextureFormat = 0x0000004B
	// TextureFormatASTC5x4Unorm is ASTC 5x4 normalized unsigned format.
	TextureFormatASTC5x4Unorm TextureFormat = 0x0000004C
	// TextureFormatASTC5x4UnormSrgb is ASTC 5x4 normalized unsigned sRGB format.
	TextureFormatASTC5x4UnormSrgb TextureFormat = 0x0000004D
	// TextureFormatASTC5x5Unorm is ASTC 5x5 normalized unsigned format.
	TextureFormatASTC5x5Unorm TextureFormat = 0x0000004E
	// TextureFormatASTC5x5UnormSrgb is ASTC 5x5 normalized unsigned sRGB format.
	TextureFormatASTC5x5UnormSrgb TextureFormat = 0x0000004F
	// TextureFormatASTC6x5Unorm is ASTC 6x5 normalized unsigned format.
	TextureFormatASTC6x5Unorm TextureFormat = 0x00000050
	// TextureFormatASTC6x5UnormSrgb is ASTC 6x5 normalized unsigned sRGB format.
	TextureFormatASTC6x5UnormSrgb TextureFormat = 0x00000051
	// TextureFormatASTC6x6Unorm is ASTC 6x6 normalized unsigned format.
	TextureFormatASTC6x6Unorm TextureFormat = 0x00000052
	// TextureFormatASTC6x6UnormSrgb is ASTC 6x6 normalized unsigned sRGB format.
	TextureFormatASTC6x6UnormSrgb TextureFormat = 0x00000053
	// TextureFormatASTC8x5Unorm is ASTC 8x5 normalized unsigned format.
	TextureFormatASTC8x5Unorm TextureFormat = 0x00000054
	// TextureFormatASTC8x5UnormSrgb is ASTC 8x5 normalized unsigned sRGB format.
	TextureFormatASTC8x5UnormSrgb TextureFormat = 0x00000055
	// TextureFormatASTC8x6Unorm is ASTC 8x6 normalized unsigned format.
	TextureFormatASTC8x6Unorm TextureFormat = 0x00000056
	// TextureFormatASTC8x6UnormSrgb is ASTC 8x6 normalized unsigned sRGB format.
	TextureFormatASTC8x6UnormSrgb TextureFormat = 0x00000057
	// TextureFormatASTC8x8Unorm is ASTC 8x8 normalized unsigned format.
	TextureFormatASTC8x8Unorm TextureFormat = 0x00000058
	// TextureFormatASTC8x8UnormSrgb is ASTC 8x8 normalized unsigned sRGB format.
	TextureFormatASTC8x8UnormSrgb TextureFormat = 0x00000059
	// TextureFormatASTC10x5Unorm is ASTC 10x5 normalized unsigned format.
	TextureFormatASTC10x5Unorm TextureFormat = 0x0000005A
	// TextureFormatASTC10x5UnormSrgb is ASTC 10x5 normalized unsigned sRGB format.
	TextureFormatASTC10x5UnormSrgb TextureFormat = 0x0000005B
	// TextureFormatASTC10x6Unorm is ASTC 10x6 normalized unsigned format.
	TextureFormatASTC10x6Unorm TextureFormat = 0x0000005C
	// TextureFormatASTC10x6UnormSrgb is ASTC 10x6 normalized unsigned sRGB format.
	TextureFormatASTC10x6UnormSrgb TextureFormat = 0x0000005D
	// TextureFormatASTC10x8Unorm is ASTC 10x8 normalized unsigned format.
	TextureFormatASTC10x8Unorm TextureFormat = 0x0000005E
	// TextureFormatASTC10x8UnormSrgb is ASTC 10x8 normalized unsigned sRGB format.
	TextureFormatASTC10x8UnormSrgb TextureFormat = 0x0000005F
	// TextureFormatASTC10x10Unorm is ASTC 10x10 normalized unsigned format.
	TextureFormatASTC10x10Unorm TextureFormat = 0x00000060
	// TextureFormatASTC10x10UnormSrgb is ASTC 10x10 normalized unsigned sRGB format.
	TextureFormatASTC10x10UnormSrgb TextureFormat = 0x00000061
	// TextureFormatASTC12x10Unorm is ASTC 12x10 normalized unsigned format.
	TextureFormatASTC12x10Unorm TextureFormat = 0x00000062
	// TextureFormatASTC12x10UnormSrgb is ASTC 12x10 normalized unsigned sRGB format.
	TextureFormatASTC12x10UnormSrgb TextureFormat = 0x00000063
	// TextureFormatASTC12x12Unorm is ASTC 12x12 normalized unsigned format.
	TextureFormatASTC12x12Unorm TextureFormat = 0x00000064
	// TextureFormatASTC12x12UnormSrgb is ASTC 12x12 normalized unsigned sRGB format.
	TextureFormatASTC12x12UnormSrgb TextureFormat = 0x00000065
)

// textureFormatNames maps each texture format to its name.
var textureFormatNames = map[TextureFormat]string{
	TextureFormatUndefined:            nameUndefined,
	TextureFormatR8Unorm:              "R8Unorm",
	TextureFormatR8Snorm:              "R8Snorm",
	TextureFormatR8Uint:               "R8Uint",
	TextureFormatR8Sint:               "R8Sint",
	TextureFormatR16Unorm:             "R16Unorm",
	TextureFormatR16Snorm:             "R16Snorm",
	TextureFormatR16Uint:              "R16Uint",
	TextureFormatR16Sint:              "R16Sint",
	TextureFormatR16Float:             "R16Float",
	TextureFormatRG8Unorm:             "RG8Unorm",
	TextureFormatRG8Snorm:             "RG8Snorm",
	TextureFormatRG8Uint:              "RG8Uint",
	TextureFormatRG8Sint:              "RG8Sint",
	TextureFormatR32Float:             "R32Float",
	TextureFormatR32Uint:              "R32Uint",
	TextureFormatR32Sint:              "R32Sint",
	TextureFormatRG16Unorm:            "RG16Unorm",
	TextureFormatRG16Snorm:            "RG16Snorm",
	TextureFormatRG16Uint:             "RG16Uint",
	TextureFormatRG16Sint:             "RG16Sint",
	TextureFormatRG16Float:            "RG16Float",
	TextureFormatRGBA8Unorm:           "RGBA8Unorm",
	TextureFormatRGBA8UnormSrgb:       "RGBA8UnormSrgb",
	TextureFormatRGBA8Snorm:           "RGBA8Snorm",
	TextureFormatRGBA8Uint:            "RGBA8Uint",
	TextureFormatRGBA8Sint:            "RGBA8Sint",
	TextureFormatBGRA8Unorm:           "BGRA8Unorm",
	TextureFormatBGRA8UnormSrgb:       "BGRA8UnormSrgb",
	TextureFormatRGB9E5Ufloat:         "RGB9E5Ufloat",
	TextureFormatRGB10A2Uint:          "RGB10A2Uint",
	TextureFormatRGB10A2Unorm:         "RGB10A2Unorm",
	TextureFormatRG11B10Ufloat:        "RG11B10Ufloat",
	TextureFormatRG32Float:            "RG32Float",
	TextureFormatRG32Uint:             "RG32Uint",
	TextureFormatRG32Sint:             "RG32Sint",
	TextureFormatRGBA16Unorm:          "RGBA16Unorm",
	TextureFormatRGBA16Snorm:          "RGBA16Snorm",
	TextureFormatRGBA16Uint:           "RGBA16Uint",
	TextureFormatRGBA16Sint:           "RGBA16Sint",
	TextureFormatRGBA16Float:          "RGBA16Float",
	TextureFormatRGBA32Float:          "RGBA32Float",
	TextureFormatRGBA32Uint:           "RGBA32Uint",
	TextureFormatRGBA32Sint:           "RGBA32Sint",
	TextureFormatStencil8:             "Stencil8",
	TextureFormatDepth16Unorm:         "Depth16Unorm",
	TextureFormatDepth24Plus:          "Depth24Plus",
	TextureFormatDepth24PlusStencil8:  "Depth24PlusStencil8",
	TextureFormatDepth32Float:         "Depth32Float",
	TextureFormatDepth32FloatStencil8: "Depth32FloatStencil8",
	TextureFormatBC1RGBAUnorm:         "BC1RGBAUnorm",
	TextureFormatBC1RGBAUnormSrgb:     "BC1RGBAUnormSrgb",
	TextureFormatBC2RGBAUnorm:         "BC2RGBAUnorm",
	TextureFormatBC2RGBAUnormSrgb:     "BC2RGBAUnormSrgb",
	TextureFormatBC3RGBAUnorm:         "BC3RGBAUnorm",
	TextureFormatBC3RGBAUnormSrgb:     "BC3RGBAUnormSrgb",
	TextureFormatBC4RUnorm:            "BC4RUnorm",
	TextureFormatBC4RSnorm:            "BC4RSnorm",
	TextureFormatBC5RGUnorm:           "BC5RGUnorm",
	TextureFormatBC5RGSnorm:           "BC5RGSnorm",
	TextureFormatBC6HRGBUfloat:        "BC6HRGBUfloat",
	TextureFormatBC6HRGBFloat:         "BC6HRGBFloat",
	TextureFormatBC7RGBAUnorm:         "BC7RGBAUnorm",
	TextureFormatBC7RGBAUnormSrgb:     "BC7RGBAUnormSrgb",
	TextureFormatETC2RGB8Unorm:        "ETC2RGB8Unorm",
	TextureFormatETC2RGB8UnormSrgb:    "ETC2RGB8UnormSrgb",
	TextureFormatETC2RGB8A1Unorm:      "ETC2RGB8A1Unorm",
	TextureFormatETC2RGB8A1UnormSrgb:  "ETC2RGB8A1UnormSrgb",
	TextureFormatETC2RGBA8Unorm:       "ETC2RGBA8Unorm",
	TextureFormatETC2RGBA8UnormSrgb:   "ETC2RGBA8UnormSrgb",
	TextureFormatEACR11Unorm:          "EACR11Unorm",
	TextureFormatEACR11Snorm:          "EACR11Snorm",
	TextureFormatEACRG11Unorm:         "EACRG11Unorm",
	TextureFormatEACRG11Snorm:         "EACRG11Snorm",
	TextureFormatASTC4x4Unorm:         "ASTC4x4Unorm",
	TextureFormatASTC4x4UnormSrgb:     "ASTC4x4UnormSrgb",
	TextureFormatASTC5x4Unorm:         "ASTC5x4Unorm",
	TextureFormatASTC5x4UnormSrgb:     "ASTC5x4UnormSrgb",
	TextureFormatASTC5x5Unorm:         "ASTC5x5Unorm",
	TextureFormatASTC5x5UnormSrgb:     "ASTC5x5UnormSrgb",
	TextureFormatASTC6x5Unorm:         "ASTC6x5Unorm",
	TextureFormatASTC6x5UnormSrgb:     "ASTC6x5UnormSrgb",
	TextureFormatASTC6x6Unorm:         "ASTC6x6Unorm",
	TextureFormatASTC6x6UnormSrgb:     "ASTC6x6UnormSrgb",
	TextureFormatASTC8x5Unorm:         "ASTC8x5Unorm",
	TextureFormatASTC8x5UnormSrgb:     "ASTC8x5UnormSrgb",
	TextureFormatASTC8x6Unorm:         "ASTC8x6Unorm",
	TextureFormatASTC8x6UnormSrgb:     "ASTC8x6UnormSrgb",
	TextureFormatASTC8x8Unorm:         "ASTC8x8Unorm",
	TextureFormatASTC8x8UnormSrgb:     "ASTC8x8UnormSrgb",
	TextureFormatASTC10x5Unorm:        "ASTC10x5Unorm",
	TextureFormatASTC10x5UnormSrgb:    "ASTC10x5UnormSrgb",
	TextureFormatASTC10x6Unorm:        "ASTC10x6Unorm",
	TextureFormatASTC10x6UnormSrgb:    "ASTC10x6UnormSrgb",
	TextureFormatASTC10x8Unorm:        "ASTC10x8Unorm",
	TextureFormatASTC10x8UnormSrgb:    "ASTC10x8UnormSrgb",
	TextureFormatASTC10x10Unorm:       "ASTC10x10Unorm",
	TextureFormatASTC10x10UnormSrgb:   "ASTC10x10UnormSrgb",
	TextureFormatASTC12x10Unorm:       "ASTC12x10Unorm",
	TextureFormatASTC12x10UnormSrgb:   "ASTC12x10UnormSrgb",
	TextureFormatASTC12x12Unorm:       "ASTC12x12Unorm",
	TextureFormatASTC12x12UnormSrgb:   "ASTC12x12UnormSrgb",
}

// String returns the name of the texture format.
func (f TextureFormat) String() string {
	if name, ok := textureFormatNames[f]; ok {
		return name
	}
	return nameUnknown
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

// BlockCopySize returns the number of bytes occupied per texel block for this format.
//
// For uncompressed formats, one block equals one texel.
// For compressed formats (BC, ETC2, ASTC), one block covers multiple texels
// (e.g. 4x4 for BC/ETC2, variable for ASTC), but the returned value is
// the byte size of that block.
//
// Returns 0 for formats whose copy size is implementation-defined or depends
// on the copied aspect (use BlockCopySizeAspect for those):
//   - Depth24Plus
//   - Depth24PlusStencil8
//   - Depth32FloatStencil8
//
// Returns 0 for unknown or invalid formats (including TextureFormatUndefined).
//
// Reference: Rust wgpu-types TextureFormat::block_copy_size().
func (f TextureFormat) BlockCopySize() uint32 {
	switch f {
	// 8-bit formats (1 byte per texel)
	case TextureFormatR8Unorm,
		TextureFormatR8Snorm,
		TextureFormatR8Uint,
		TextureFormatR8Sint,
		TextureFormatStencil8:
		return 1

	// 16-bit formats (2 bytes per texel)
	case TextureFormatR16Unorm,
		TextureFormatR16Snorm,
		TextureFormatR16Uint,
		TextureFormatR16Sint,
		TextureFormatR16Float,
		TextureFormatRG8Unorm,
		TextureFormatRG8Snorm,
		TextureFormatRG8Uint,
		TextureFormatRG8Sint,
		TextureFormatDepth16Unorm:
		return 2

	// 32-bit formats (4 bytes per texel)
	case TextureFormatR32Float,
		TextureFormatR32Uint,
		TextureFormatR32Sint,
		TextureFormatRG16Unorm,
		TextureFormatRG16Snorm,
		TextureFormatRG16Uint,
		TextureFormatRG16Sint,
		TextureFormatRG16Float,
		TextureFormatRGBA8Unorm,
		TextureFormatRGBA8UnormSrgb,
		TextureFormatRGBA8Snorm,
		TextureFormatRGBA8Uint,
		TextureFormatRGBA8Sint,
		TextureFormatBGRA8Unorm,
		TextureFormatBGRA8UnormSrgb,
		TextureFormatRGB10A2Uint,
		TextureFormatRGB10A2Unorm,
		TextureFormatRG11B10Ufloat,
		TextureFormatRGB9E5Ufloat,
		TextureFormatDepth32Float:
		return 4

	// 64-bit formats (8 bytes per texel)
	case TextureFormatRG32Float,
		TextureFormatRG32Uint,
		TextureFormatRG32Sint,
		TextureFormatRGBA16Unorm,
		TextureFormatRGBA16Snorm,
		TextureFormatRGBA16Uint,
		TextureFormatRGBA16Sint,
		TextureFormatRGBA16Float:
		return 8

	// 128-bit formats (16 bytes per texel)
	case TextureFormatRGBA32Float,
		TextureFormatRGBA32Uint,
		TextureFormatRGBA32Sint:
		return 16

	// Depth/stencil formats with implementation-defined copy size
	case TextureFormatDepth24Plus,
		TextureFormatDepth24PlusStencil8,
		TextureFormatDepth32FloatStencil8:
		return 0

	default:
		return compressedBlockCopySize(f)
	}
}

// compressedBlockCopySize returns the bytes per block of a compressed format,
// or 0 for any other format.
func compressedBlockCopySize(f TextureFormat) uint32 {
	switch f {
	// BC compressed formats — 8 bytes per 4x4 block
	case TextureFormatBC1RGBAUnorm,
		TextureFormatBC1RGBAUnormSrgb,
		TextureFormatBC4RUnorm,
		TextureFormatBC4RSnorm:
		return 8

	// BC compressed formats — 16 bytes per 4x4 block
	case TextureFormatBC2RGBAUnorm,
		TextureFormatBC2RGBAUnormSrgb,
		TextureFormatBC3RGBAUnorm,
		TextureFormatBC3RGBAUnormSrgb,
		TextureFormatBC5RGUnorm,
		TextureFormatBC5RGSnorm,
		TextureFormatBC6HRGBUfloat,
		TextureFormatBC6HRGBFloat,
		TextureFormatBC7RGBAUnorm,
		TextureFormatBC7RGBAUnormSrgb:
		return 16

	// ETC2/EAC compressed formats — 8 bytes per 4x4 block
	case TextureFormatETC2RGB8Unorm,
		TextureFormatETC2RGB8UnormSrgb,
		TextureFormatETC2RGB8A1Unorm,
		TextureFormatETC2RGB8A1UnormSrgb,
		TextureFormatEACR11Unorm,
		TextureFormatEACR11Snorm:
		return 8

	// ETC2/EAC compressed formats — 16 bytes per 4x4 block
	case TextureFormatETC2RGBA8Unorm,
		TextureFormatETC2RGBA8UnormSrgb,
		TextureFormatEACRG11Unorm,
		TextureFormatEACRG11Snorm:
		return 16

	// ASTC compressed formats — all 16 bytes per block (variable block dimensions)
	case TextureFormatASTC4x4Unorm,
		TextureFormatASTC4x4UnormSrgb,
		TextureFormatASTC5x4Unorm,
		TextureFormatASTC5x4UnormSrgb,
		TextureFormatASTC5x5Unorm,
		TextureFormatASTC5x5UnormSrgb,
		TextureFormatASTC6x5Unorm,
		TextureFormatASTC6x5UnormSrgb,
		TextureFormatASTC6x6Unorm,
		TextureFormatASTC6x6UnormSrgb,
		TextureFormatASTC8x5Unorm,
		TextureFormatASTC8x5UnormSrgb,
		TextureFormatASTC8x6Unorm,
		TextureFormatASTC8x6UnormSrgb,
		TextureFormatASTC8x8Unorm,
		TextureFormatASTC8x8UnormSrgb,
		TextureFormatASTC10x5Unorm,
		TextureFormatASTC10x5UnormSrgb,
		TextureFormatASTC10x6Unorm,
		TextureFormatASTC10x6UnormSrgb,
		TextureFormatASTC10x8Unorm,
		TextureFormatASTC10x8UnormSrgb,
		TextureFormatASTC10x10Unorm,
		TextureFormatASTC10x10UnormSrgb,
		TextureFormatASTC12x10Unorm,
		TextureFormatASTC12x10UnormSrgb,
		TextureFormatASTC12x12Unorm,
		TextureFormatASTC12x12UnormSrgb:
		return 16

	default:
		return 0
	}
}

// TextureDimension describes texture dimensions.
type TextureDimension uint32

const (
	// TextureDimensionUndefined is an undefined texture dimension (invalid).
	TextureDimensionUndefined TextureDimension = 0x00000000
	// TextureDimension1D is a 1D texture.
	TextureDimension1D TextureDimension = 0x00000001
	// TextureDimension2D is a 2D texture.
	TextureDimension2D TextureDimension = 0x00000002
	// TextureDimension3D is a 3D texture.
	TextureDimension3D TextureDimension = 0x00000003
)

// String returns the dimension name.
func (d TextureDimension) String() string {
	switch d {
	case TextureDimensionUndefined:
		return nameUndefined
	case TextureDimension1D:
		return "1D"
	case TextureDimension2D:
		return "2D"
	case TextureDimension3D:
		return "3D"
	default:
		return nameUnknown
	}
}

// TextureViewDimension describes a texture view dimension.
type TextureViewDimension uint32

const (
	// TextureViewDimensionUndefined uses the same dimension as the texture.
	TextureViewDimensionUndefined TextureViewDimension = 0x00000000
	// TextureViewDimension1D is a 1D texture view.
	TextureViewDimension1D TextureViewDimension = 0x00000001
	// TextureViewDimension2D is a 2D texture view.
	TextureViewDimension2D TextureViewDimension = 0x00000002
	// TextureViewDimension2DArray is a 2D array texture view.
	TextureViewDimension2DArray TextureViewDimension = 0x00000003
	// TextureViewDimensionCube is a cube texture view.
	TextureViewDimensionCube TextureViewDimension = 0x00000004
	// TextureViewDimensionCubeArray is a cube array texture view.
	TextureViewDimensionCubeArray TextureViewDimension = 0x00000005
	// TextureViewDimension3D is a 3D texture view.
	TextureViewDimension3D TextureViewDimension = 0x00000006
)

// String returns the view dimension name.
func (d TextureViewDimension) String() string {
	switch d {
	case TextureViewDimensionUndefined:
		return nameUndefined
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
		return nameUnknown
	}
}

// TextureAspect describes which aspects of a texture to access.
type TextureAspect uint32

const (
	// TextureAspectUndefined is an undefined texture aspect (invalid).
	TextureAspectUndefined TextureAspect = 0x00000000
	// TextureAspectAll accesses all aspects (default).
	TextureAspectAll TextureAspect = 0x00000001
	// TextureAspectStencilOnly accesses only the stencil aspect.
	TextureAspectStencilOnly TextureAspect = 0x00000002
	// TextureAspectDepthOnly accesses only the depth aspect.
	TextureAspectDepthOnly TextureAspect = 0x00000003
)

// String returns the aspect name.
func (a TextureAspect) String() string {
	switch a {
	case TextureAspectUndefined:
		return nameUndefined
	case TextureAspectAll:
		return "All"
	case TextureAspectStencilOnly:
		return "StencilOnly"
	case TextureAspectDepthOnly:
		return "DepthOnly"
	default:
		return nameUnknown
	}
}

// TextureUsage describes how a texture can be used.
//
// This is a bit flag type. Combine multiple usages with bitwise OR.
type TextureUsage uint64

const (
	// TextureUsageNone indicates no usage (invalid for most operations).
	TextureUsageNone TextureUsage = 0x0000000000000000
	// TextureUsageCopySrc allows the texture to be a copy source.
	TextureUsageCopySrc TextureUsage = 0x0000000000000001
	// TextureUsageCopyDst allows the texture to be a copy destination.
	TextureUsageCopyDst TextureUsage = 0x0000000000000002
	// TextureUsageTextureBinding allows the texture to be bound as a sampled texture.
	TextureUsageTextureBinding TextureUsage = 0x0000000000000004
	// TextureUsageStorageBinding allows the texture to be bound as a storage texture.
	TextureUsageStorageBinding TextureUsage = 0x0000000000000008
	// TextureUsageRenderAttachment allows the texture to be used as a render attachment.
	TextureUsageRenderAttachment TextureUsage = 0x0000000000000010
)

// textureUsageAll is a mask of all valid texture usage flags.
const textureUsageAll = TextureUsageCopySrc | TextureUsageCopyDst |
	TextureUsageTextureBinding | TextureUsageStorageBinding |
	TextureUsageRenderAttachment

// Contains returns true if the usage includes the given flag.
func (u TextureUsage) Contains(flag TextureUsage) bool {
	return u&flag == flag
}

// ContainsUnknownBits returns true if the usage contains any unknown flags.
func (u TextureUsage) ContainsUnknownBits() bool {
	return u&^textureUsageAll != 0
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
type TextureSampleType uint32

const (
	// TextureSampleTypeUndefined is an undefined sample type (invalid).
	TextureSampleTypeUndefined TextureSampleType = 0x00000000
	// TextureSampleTypeFloat samples as filterable floating-point.
	TextureSampleTypeFloat TextureSampleType = 0x00000001
	// TextureSampleTypeUnfilterableFloat samples as unfilterable floating-point.
	TextureSampleTypeUnfilterableFloat TextureSampleType = 0x00000002
	// TextureSampleTypeDepth samples as depth comparison.
	TextureSampleTypeDepth TextureSampleType = 0x00000003
	// TextureSampleTypeSint samples as signed integer.
	TextureSampleTypeSint TextureSampleType = 0x00000004
	// TextureSampleTypeUint samples as unsigned integer.
	TextureSampleTypeUint TextureSampleType = 0x00000005
)

// String returns the sample type name.
func (t TextureSampleType) String() string {
	switch t {
	case TextureSampleTypeUndefined:
		return nameUndefined
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
		return nameUnknown
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
