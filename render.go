package gputypes

// LoadOp describes the load operation for an attachment.
type LoadOp uint8

const (
	// LoadOpClear clears the attachment to a specified value.
	LoadOpClear LoadOp = iota
	// LoadOpLoad loads the existing contents.
	LoadOpLoad
)

// String returns the load operation name.
func (op LoadOp) String() string {
	switch op {
	case LoadOpClear:
		return "Clear"
	case LoadOpLoad:
		return "Load"
	default:
		return "Unknown"
	}
}

// StoreOp describes the store operation for an attachment.
type StoreOp uint8

const (
	// StoreOpDiscard discards the contents (for performance when not needed).
	StoreOpDiscard StoreOp = iota
	// StoreOpStore stores the contents to memory.
	StoreOpStore
)

// String returns the store operation name.
func (op StoreOp) String() string {
	switch op {
	case StoreOpDiscard:
		return "Discard"
	case StoreOpStore:
		return "Store"
	default:
		return "Unknown"
	}
}

// BlendFactor describes a blend factor for color blending.
type BlendFactor uint8

const (
	// BlendFactorZero uses 0.
	BlendFactorZero BlendFactor = iota
	// BlendFactorOne uses 1.
	BlendFactorOne
	// BlendFactorSrc uses the source color.
	BlendFactorSrc
	// BlendFactorOneMinusSrc uses (1 - source color).
	BlendFactorOneMinusSrc
	// BlendFactorSrcAlpha uses the source alpha.
	BlendFactorSrcAlpha
	// BlendFactorOneMinusSrcAlpha uses (1 - source alpha).
	BlendFactorOneMinusSrcAlpha
	// BlendFactorDst uses the destination color.
	BlendFactorDst
	// BlendFactorOneMinusDst uses (1 - destination color).
	BlendFactorOneMinusDst
	// BlendFactorDstAlpha uses the destination alpha.
	BlendFactorDstAlpha
	// BlendFactorOneMinusDstAlpha uses (1 - destination alpha).
	BlendFactorOneMinusDstAlpha
	// BlendFactorSrcAlphaSaturated uses min(source alpha, 1 - destination alpha).
	BlendFactorSrcAlphaSaturated
	// BlendFactorConstant uses the constant blend color.
	BlendFactorConstant
	// BlendFactorOneMinusConstant uses (1 - constant blend color).
	BlendFactorOneMinusConstant
)

// String returns the blend factor name.
func (f BlendFactor) String() string {
	switch f {
	case BlendFactorZero:
		return "Zero"
	case BlendFactorOne:
		return "One"
	case BlendFactorSrc:
		return "Src"
	case BlendFactorOneMinusSrc:
		return "OneMinusSrc"
	case BlendFactorSrcAlpha:
		return "SrcAlpha"
	case BlendFactorOneMinusSrcAlpha:
		return "OneMinusSrcAlpha"
	case BlendFactorDst:
		return "Dst"
	case BlendFactorOneMinusDst:
		return "OneMinusDst"
	case BlendFactorDstAlpha:
		return "DstAlpha"
	case BlendFactorOneMinusDstAlpha:
		return "OneMinusDstAlpha"
	case BlendFactorSrcAlphaSaturated:
		return "SrcAlphaSaturated"
	case BlendFactorConstant:
		return "Constant"
	case BlendFactorOneMinusConstant:
		return "OneMinusConstant"
	default:
		return "Unknown"
	}
}

// BlendOperation describes a blend operation.
type BlendOperation uint8

const (
	// BlendOperationAdd computes src + dst.
	BlendOperationAdd BlendOperation = iota
	// BlendOperationSubtract computes src - dst.
	BlendOperationSubtract
	// BlendOperationReverseSubtract computes dst - src.
	BlendOperationReverseSubtract
	// BlendOperationMin computes min(src, dst).
	BlendOperationMin
	// BlendOperationMax computes max(src, dst).
	BlendOperationMax
)

// String returns the blend operation name.
func (op BlendOperation) String() string {
	switch op {
	case BlendOperationAdd:
		return "Add"
	case BlendOperationSubtract:
		return "Subtract"
	case BlendOperationReverseSubtract:
		return "ReverseSubtract"
	case BlendOperationMin:
		return "Min"
	case BlendOperationMax:
		return "Max"
	default:
		return "Unknown"
	}
}

// BlendComponent describes blending for a single color component (RGB or alpha).
type BlendComponent struct {
	// SrcFactor is the source blend factor.
	SrcFactor BlendFactor
	// DstFactor is the destination blend factor.
	DstFactor BlendFactor
	// Operation is the blend operation.
	Operation BlendOperation
}

// BlendState describes color blending for a render target.
type BlendState struct {
	// Color describes RGB channel blending.
	Color BlendComponent
	// Alpha describes alpha channel blending.
	Alpha BlendComponent
}

// BlendStateReplace returns a blend state that replaces the destination.
func BlendStateReplace() BlendState {
	return BlendState{
		Color: BlendComponent{
			SrcFactor: BlendFactorOne,
			DstFactor: BlendFactorZero,
			Operation: BlendOperationAdd,
		},
		Alpha: BlendComponent{
			SrcFactor: BlendFactorOne,
			DstFactor: BlendFactorZero,
			Operation: BlendOperationAdd,
		},
	}
}

// BlendStateAlpha returns a standard alpha blending state.
//
// This is the most common blend state for transparent rendering.
func BlendStateAlpha() BlendState {
	return BlendState{
		Color: BlendComponent{
			SrcFactor: BlendFactorSrcAlpha,
			DstFactor: BlendFactorOneMinusSrcAlpha,
			Operation: BlendOperationAdd,
		},
		Alpha: BlendComponent{
			SrcFactor: BlendFactorOne,
			DstFactor: BlendFactorOneMinusSrcAlpha,
			Operation: BlendOperationAdd,
		},
	}
}

// BlendStatePremultiplied returns a blend state for premultiplied alpha.
func BlendStatePremultiplied() BlendState {
	return BlendState{
		Color: BlendComponent{
			SrcFactor: BlendFactorOne,
			DstFactor: BlendFactorOneMinusSrcAlpha,
			Operation: BlendOperationAdd,
		},
		Alpha: BlendComponent{
			SrcFactor: BlendFactorOne,
			DstFactor: BlendFactorOneMinusSrcAlpha,
			Operation: BlendOperationAdd,
		},
	}
}

// ColorWriteMask describes which color channels to write.
//
// This is a bit flag type.
type ColorWriteMask uint8

const (
	// ColorWriteMaskNone writes no channels.
	ColorWriteMaskNone ColorWriteMask = 0
	// ColorWriteMaskRed writes the red channel.
	ColorWriteMaskRed ColorWriteMask = 1 << iota
	// ColorWriteMaskGreen writes the green channel.
	ColorWriteMaskGreen
	// ColorWriteMaskBlue writes the blue channel.
	ColorWriteMaskBlue
	// ColorWriteMaskAlpha writes the alpha channel.
	ColorWriteMaskAlpha
	// ColorWriteMaskAll writes all channels.
	ColorWriteMaskAll = ColorWriteMaskRed | ColorWriteMaskGreen | ColorWriteMaskBlue | ColorWriteMaskAlpha
)

// ColorTargetState describes a color target in a render pipeline.
type ColorTargetState struct {
	// Format is the texture format of the target.
	Format TextureFormat
	// Blend describes color blending (nil for no blending).
	Blend *BlendState
	// WriteMask specifies which color channels to write.
	WriteMask ColorWriteMask
}

// PrimitiveTopology describes how vertices form primitives.
type PrimitiveTopology uint8

const (
	// PrimitiveTopologyPointList renders each vertex as a point.
	PrimitiveTopologyPointList PrimitiveTopology = iota
	// PrimitiveTopologyLineList renders pairs of vertices as lines.
	PrimitiveTopologyLineList
	// PrimitiveTopologyLineStrip renders connected lines.
	PrimitiveTopologyLineStrip
	// PrimitiveTopologyTriangleList renders groups of 3 vertices as triangles.
	PrimitiveTopologyTriangleList
	// PrimitiveTopologyTriangleStrip renders connected triangles.
	PrimitiveTopologyTriangleStrip
)

// String returns the topology name.
func (t PrimitiveTopology) String() string {
	switch t {
	case PrimitiveTopologyPointList:
		return "PointList"
	case PrimitiveTopologyLineList:
		return "LineList"
	case PrimitiveTopologyLineStrip:
		return "LineStrip"
	case PrimitiveTopologyTriangleList:
		return "TriangleList"
	case PrimitiveTopologyTriangleStrip:
		return "TriangleStrip"
	default:
		return "Unknown"
	}
}

// FrontFace describes the front face winding order.
type FrontFace uint8

const (
	// FrontFaceCCW treats counter-clockwise vertices as front-facing.
	FrontFaceCCW FrontFace = iota
	// FrontFaceCW treats clockwise vertices as front-facing.
	FrontFaceCW
)

// String returns the front face name.
func (f FrontFace) String() string {
	switch f {
	case FrontFaceCCW:
		return "CCW"
	case FrontFaceCW:
		return "CW"
	default:
		return "Unknown"
	}
}

// CullMode describes which faces to cull.
type CullMode uint8

const (
	// CullModeNone culls no faces.
	CullModeNone CullMode = iota
	// CullModeFront culls front faces.
	CullModeFront
	// CullModeBack culls back faces.
	CullModeBack
)

// String returns the cull mode name.
func (m CullMode) String() string {
	switch m {
	case CullModeNone:
		return "None"
	case CullModeFront:
		return "Front"
	case CullModeBack:
		return "Back"
	default:
		return "Unknown"
	}
}

// PrimitiveState describes primitive assembly state.
type PrimitiveState struct {
	// Topology is the primitive topology.
	Topology PrimitiveTopology
	// StripIndexFormat is the index format for strip topologies (nil for non-strip).
	StripIndexFormat *IndexFormat
	// FrontFace is the front face winding order.
	FrontFace FrontFace
	// CullMode specifies which faces to cull.
	CullMode CullMode
	// UnclippedDepth enables unclipped depth (requires feature).
	UnclippedDepth bool
}

// DefaultPrimitiveState returns a primitive state with common defaults.
func DefaultPrimitiveState() PrimitiveState {
	return PrimitiveState{
		Topology:       PrimitiveTopologyTriangleList,
		FrontFace:      FrontFaceCCW,
		CullMode:       CullModeNone,
		UnclippedDepth: false,
	}
}

// MultisampleState describes multisampling state.
type MultisampleState struct {
	// Count is the number of samples per pixel (1, 2, 4, 8, or 16).
	Count uint32
	// Mask is the sample mask (all bits set = all samples).
	Mask uint64
	// AlphaToCoverageEnabled enables alpha-to-coverage.
	AlphaToCoverageEnabled bool
}

// DefaultMultisampleState returns a multisample state with no multisampling.
func DefaultMultisampleState() MultisampleState {
	return MultisampleState{
		Count:                  1,
		Mask:                   0xFFFFFFFF,
		AlphaToCoverageEnabled: false,
	}
}

// StencilOperation describes a stencil operation.
type StencilOperation uint8

const (
	// StencilOperationKeep keeps the current value.
	StencilOperationKeep StencilOperation = iota
	// StencilOperationZero sets to zero.
	StencilOperationZero
	// StencilOperationReplace replaces with reference value.
	StencilOperationReplace
	// StencilOperationInvert inverts all bits.
	StencilOperationInvert
	// StencilOperationIncrementClamp increments and clamps to maximum.
	StencilOperationIncrementClamp
	// StencilOperationDecrementClamp decrements and clamps to zero.
	StencilOperationDecrementClamp
	// StencilOperationIncrementWrap increments and wraps to zero.
	StencilOperationIncrementWrap
	// StencilOperationDecrementWrap decrements and wraps to maximum.
	StencilOperationDecrementWrap
)

// String returns the stencil operation name.
func (op StencilOperation) String() string {
	switch op {
	case StencilOperationKeep:
		return "Keep"
	case StencilOperationZero:
		return "Zero"
	case StencilOperationReplace:
		return "Replace"
	case StencilOperationInvert:
		return "Invert"
	case StencilOperationIncrementClamp:
		return "IncrementClamp"
	case StencilOperationDecrementClamp:
		return "DecrementClamp"
	case StencilOperationIncrementWrap:
		return "IncrementWrap"
	case StencilOperationDecrementWrap:
		return "DecrementWrap"
	default:
		return "Unknown"
	}
}

// StencilFaceState describes stencil operations for a face.
type StencilFaceState struct {
	// Compare is the comparison function.
	Compare CompareFunction
	// FailOp is the operation on stencil test failure.
	FailOp StencilOperation
	// DepthFailOp is the operation on depth test failure.
	DepthFailOp StencilOperation
	// PassOp is the operation on both tests passing.
	PassOp StencilOperation
}

// DefaultStencilFaceState returns a stencil face state that always passes and keeps.
func DefaultStencilFaceState() StencilFaceState {
	return StencilFaceState{
		Compare:     CompareFunctionAlways,
		FailOp:      StencilOperationKeep,
		DepthFailOp: StencilOperationKeep,
		PassOp:      StencilOperationKeep,
	}
}

// DepthStencilState describes depth and stencil state.
type DepthStencilState struct {
	// Format is the depth/stencil texture format.
	Format TextureFormat
	// DepthWriteEnabled enables depth writing.
	DepthWriteEnabled bool
	// DepthCompare is the depth comparison function.
	DepthCompare CompareFunction
	// StencilFront describes front face stencil state.
	StencilFront StencilFaceState
	// StencilBack describes back face stencil state.
	StencilBack StencilFaceState
	// StencilReadMask is the mask for stencil reads.
	StencilReadMask uint32
	// StencilWriteMask is the mask for stencil writes.
	StencilWriteMask uint32
	// DepthBias is the constant depth bias.
	DepthBias int32
	// DepthBiasSlopeScale is the slope-based depth bias.
	DepthBiasSlopeScale float32
	// DepthBiasClamp is the maximum depth bias.
	DepthBiasClamp float32
}

// DefaultDepthStencilState returns a depth-stencil state with depth testing enabled.
func DefaultDepthStencilState(format TextureFormat) DepthStencilState {
	return DepthStencilState{
		Format:              format,
		DepthWriteEnabled:   true,
		DepthCompare:        CompareFunctionLess,
		StencilFront:        DefaultStencilFaceState(),
		StencilBack:         DefaultStencilFaceState(),
		StencilReadMask:     0xFFFFFFFF,
		StencilWriteMask:    0xFFFFFFFF,
		DepthBias:           0,
		DepthBiasSlopeScale: 0,
		DepthBiasClamp:      0,
	}
}

// RenderPassColorAttachment describes a color attachment for a render pass.
type RenderPassColorAttachment struct {
	// View is the texture view to render to (implementation-specific handle).
	View uintptr
	// ResolveTarget is the texture view for multisample resolve (0 if none).
	ResolveTarget uintptr
	// LoadOp describes how to load the attachment.
	LoadOp LoadOp
	// StoreOp describes how to store the attachment.
	StoreOp StoreOp
	// ClearValue is the clear color (used when LoadOp is Clear).
	ClearValue Color
}

// RenderPassDepthStencilAttachment describes a depth-stencil attachment.
type RenderPassDepthStencilAttachment struct {
	// View is the texture view (implementation-specific handle).
	View uintptr
	// DepthLoadOp describes how to load depth.
	DepthLoadOp LoadOp
	// DepthStoreOp describes how to store depth.
	DepthStoreOp StoreOp
	// DepthClearValue is the clear depth value.
	DepthClearValue float32
	// DepthReadOnly indicates if depth is read-only.
	DepthReadOnly bool
	// StencilLoadOp describes how to load stencil.
	StencilLoadOp LoadOp
	// StencilStoreOp describes how to store stencil.
	StencilStoreOp StoreOp
	// StencilClearValue is the clear stencil value.
	StencilClearValue uint32
	// StencilReadOnly indicates if stencil is read-only.
	StencilReadOnly bool
}

// RenderPassDescriptor describes a render pass.
type RenderPassDescriptor struct {
	// Label is an optional debug label.
	Label string
	// ColorAttachments are the color attachments.
	ColorAttachments []RenderPassColorAttachment
	// DepthStencilAttachment is the depth-stencil attachment (nil if none).
	DepthStencilAttachment *RenderPassDepthStencilAttachment
}
