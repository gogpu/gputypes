# Changelog

All notable changes to gputypes will be documented in this file.

## [v0.7.0] - 2026-08-30

### Added

- **DownlevelCapabilities** (ADR-071) — tracks backend capabilities for graceful degradation on non-conformant adapters (GLES 3.0, CPU software rasterizer). The W3C WebGPU spec excludes non-conformant adapters from `requestAdapter()` — as a native Go library, we degrade gracefully instead of refusing to run. Matches Rust wgpu-types `DownlevelCapabilities`.
  - `DownlevelFlags` type (`uint32`) with 27 flag constants at explicit `1 << N` bit positions matching Rust wgpu-types (`limits.rs:1102-1246`). NOT iota — binary Rust compatibility.
  - `DownlevelCapabilities` struct with 3 fields: `Flags DownlevelFlags`, `Limits DownlevelLimits`, `ShaderModel ShaderModel` (matches Rust `limits.rs:1056-1063`)
  - `DownlevelLimits` empty struct (reserved for future, matches Rust `limits.rs:1044`)
  - `ShaderModel` named type with `ShaderModelSm2` (2), `ShaderModelSm4` (4), `ShaderModelSm5` (5) constants
  - `Contains()` method on `DownlevelFlags`
  - `DownlevelFlagsAll()` — all 27 bits set
  - `DownlevelFlagsCompliant()` — all flags minus AnisotropicFiltering (WebGPU compliance, matches Rust `limits.rs:1249-1257`)
  - `DefaultDownlevelCapabilities()` — all flags + Sm5 (matches Rust Default impl, `limits.rs:1065-1072`)
  - `IsWebGPUCompliant()` method — checks compliant flags + default limits + ShaderModel >= Sm5 (matches Rust `limits.rs:1075-1085`)
  - `String()` methods on `DownlevelFlags` and `ShaderModel`
  - 379 lines of tests covering all bit positions, Contains, All, Compliant, IsWebGPUCompliant (8 sub-cases), ShaderModel ordering

### Changed

- **`DownlevelLimits()` renamed to `DownlevelDefaultLimits()`** — avoids name conflict with the new `DownlevelLimits` struct

## [v0.6.0] - 2026-08-27

### Added

- **Ray Tracing types** (experimental, matching Rust wgpu `EXPERIMENTAL_RAY_QUERY`):
  - `ray_tracing.go`: `AccelerationStructureUpdateMode`, `AccelerationStructureCopyMode`, `AccelerationStructureType` enums
  - `AccelerationStructureFlags` bitflags (7 flags: AllowUpdate, AllowCompaction, PreferFastTrace, PreferFastBuild, LowMemory, UseTransform, AllowRayHitVertexReturn)
  - `AccelerationStructureGeometryFlags` bitflags (Opaque, NoDuplicateAnyHitInvocation)
  - `BlasTriangleGeometrySizeDescriptor`, `BlasAABBGeometrySizeDescriptor`, `BlasGeometrySizeDescriptors`
  - `CreateBlasDescriptor`, `CreateTlasDescriptor`
  - `AccelerationStructureBindingLayout` (for bind group layouts)
  - Constants: `AABBGeometryMinStride` (24), `TransformBufferAlignment` (16), `InstanceBufferAlignment` (16)
- **Feature flags** (bits 20-24): `FeatureRayQuery`, `FeatureRayHitVertexReturn`, `FeatureExtendedASVertexFormats`, `FeatureASBindingArray`, `FeatureRayTracingPipelines`
- **Buffer usage flags** (bits 11-14): `BufferUsageAccelerationStructureScratch`, `BufferUsageBlasInput`, `BufferUsageTlasInput`, `BufferUsageAccelerationStructureQuery`
- **Limits fields** (8): `MaxBlasPrimitiveCount`, `MaxBlasGeometryCount`, `MaxTlasInstanceCount`, `MaxAccelerationStructuresPerShaderStage`, `MaxBuffersAndAccelerationStructuresPerShaderStage`, `MaxBindingArrayAccelerationStructureElementsPerShaderStage`, `MaxRayDispatchCount`, `MaxRayRecursionDepth`
- **BindGroupLayoutEntry**: `AccelerationStructure *AccelerationStructureBindingLayout` field
- **Tests**: feature bit position verification, no-overlap with existing bits, buffer usage hex values, flags Contains(), constants validation

**Note:** Ray tracing is experimental — not in the W3C WebGPU spec (Milestone 4+, blocked on bindless). Matches Rust wgpu's experimental approach. Feature-gated: all RT operations require `FeatureRayQuery` to be enabled on the device.

## [v0.5.2] - 2026-08-11

### Fixed

- **`Features.Contains`** — now requires all queried bits, not just any overlap (#6, @besmpl)
  - Aligns with `BufferUsage`, `TextureUsage`, `ShaderStage` (all use all-bits containment)
  - Matches Rust `bitflags::contains` semantics
  - Zero-bit queries return true (empty set is subset of everything)

## [v0.5.1] - 2026-06-28

### Fixed

- **`TextureFormat.BlockCopySize() uint32`** — canonical bytes-per-texel-block method on TextureFormat. Eliminates 3 independent buggy implementations across wgpu, gogpu, and gg (each had different wrong values — e.g., RG16 returned 8 instead of 4 in gogpu). Covers all 87 defined formats, verified against Rust wgpu-types `block_copy_size`. Returns 0 for implementation-defined formats (Depth24Plus). Table-driven tests with completeness guard.

## [v0.5.0] - 2026-04-21

### Changed (BREAKING)

- **PrimitiveTopology**: zero value is now `PrimitiveTopologyTriangleList` (was `PrimitiveTopologyUndefined`). Enum values renumbered: TriangleList=0, PointList=1, LineList=2, LineStrip=3, TriangleStrip=4. Removed `PrimitiveTopologyUndefined`.
- **FrontFace**: zero value is now `FrontFaceCCW` (was `FrontFaceUndefined`). Enum values renumbered: CCW=0, CW=1. Removed `FrontFaceUndefined`.
- **CullMode**: zero value is now `CullModeNone` (was `CullModeUndefined`). Enum values renumbered: None=0, Front=1, Back=2. Removed `CullModeUndefined`.
- **`DefaultPrimitiveState()`** now returns `PrimitiveState{}` — the zero value IS the WebGPU spec default. Function kept for explicit documentation and parity with other Default*State helpers.

**Why:** Go zero value of `PrimitiveState{}` is now a fully valid WebGPU-spec-default configuration. No normalization pass needed. Eliminates the class of bugs where `Undefined` sentinel values leak into HAL backends.

**Migration:** downstream repos (wgpu, gogpu, gg) must update HAL enum→API mapping tables. `*Undefined` constants no longer exist — remove any references or `switch` cases.

## [v0.4.0] - 2026-04-03

### Added

- **BlendComponent.UsesConstant()** — returns true if the blend component uses
  `BlendFactorConstant` or `BlendFactorOneMinusConstant` in either `SrcFactor` or
  `DstFactor`. Used by wgpu for draw-time validation that `SetBlendConstant()` has
  been called when the pipeline requires it. Matches Rust wgpu-types
  `BlendComponent::uses_constant()`.

## [v0.3.0] - 2026-03-10

### Added

- **TextureUsage.ContainsUnknownBits()** — returns true if the usage contains
  any unknown flags. Follows the same pattern as `BufferUsage.ContainsUnknownBits()`.
  Used by wgpu core validation layer for texture descriptor validation.

## [v0.2.0] - 2026-01-29

### Changed

- **webgpu.h compliance**: All enum values now use explicit hex constants matching the official WebGPU C header specification
  - TextureFormat: 97 formats with spec-compliant values (0x00000000 - 0x00000060)
  - BufferUsage: Explicit bit flags matching WebGPU spec
  - TextureUsage: Explicit bit flags matching WebGPU spec
  - LoadOp, StoreOp: Spec-compliant values
  - BlendFactor, BlendOperation: Spec-compliant values
  - AddressMode, FilterMode: Spec-compliant values
  - VertexFormat: 31 formats with spec-compliant values
  - PresentMode, CompositeAlphaMode: Spec-compliant values

### Migration

Values changed from iota-based to explicit webgpu.h values. This ensures:
- Binary compatibility with wgpu-native and other WebGPU implementations
- Correct serialization/deserialization of GPU descriptors
- Interoperability with C/Rust WebGPU code

If you were relying on specific numeric values, they may have changed. Use the named constants.

## [v0.1.0] - 2026-01-29

### Added

- Initial release with WebGPU type definitions
- **Texture types**: TextureFormat (97 formats including BC, ETC2, ASTC), TextureUsage, TextureDimension, TextureViewDimension, TextureAspect, TextureSampleType
- **Buffer types**: BufferUsage, BufferBindingType, BufferDescriptor, IndexFormat
- **Sampler types**: AddressMode, FilterMode, MipmapFilterMode, CompareFunction, SamplerDescriptor
- **Render types**: LoadOp, StoreOp, BlendState, BlendFactor, BlendOperation, ColorWriteMask, RenderPassDescriptor
- **Shader types**: ShaderStage, ShaderSource (WGSL, SPIR-V, GLSL)
- **Vertex types**: VertexFormat (31 formats), VertexStepMode, VertexBufferLayout
- **Binding types**: BindGroupLayoutEntry, BindGroupEntry, BufferBindingLayout, TextureBindingLayout
- **Adapter types**: DeviceType, Backend, PowerPreference, AdapterInfo
- **Surface types**: PresentMode, CompositeAlphaMode, SurfaceConfiguration
- **Limits & Features**: Full WebGPU limits struct, feature flags
- **Geometry types**: Extent3D, Origin3D, Color
