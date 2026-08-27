# Roadmap

## Current: v0.6.0

Ray tracing types (experimental). Feature flags, limits, buffer usage, geometry descriptors, binding layout. Matches Rust wgpu `EXPERIMENTAL_RAY_QUERY`.

## Released

### v0.6.0 (2026-08-27)
- Ray tracing: `AccelerationStructureFlags`, `BlasTriangleGeometrySizeDescriptor`, `CreateBlasDescriptor`, `CreateTlasDescriptor`, `AccelerationStructureBindingLayout`
- 5 feature bits (20-24): RayQuery, RayHitVertexReturn, ExtendedASVertexFormats, ASBindingArray, RayTracingPipelines
- 4 buffer usage bits (11-14): AccelerationStructureScratch, BlasInput, TlasInput, AccelerationStructureQuery
- 8 limits fields for RT hardware capabilities
- Tests for bit positions, no-overlap, flags Contains()

### v0.5.2 (2026-08-11)
- `Features.Contains()` fix — all-bits containment semantics

### v0.5.1 (2026-06-28)
- `TextureFormat.BlockCopySize()` — canonical bytes-per-texel-block

### v0.5.0 (2026-04-21)
- **BREAKING:** PrimitiveTopology, FrontFace, CullMode renumbered — zero value = WebGPU spec default
- Removed `*Undefined` sentinel constants
- `DefaultPrimitiveState()` now returns `PrimitiveState{}` (zero value)

### v0.4.0 (2026-04-03)
- `BlendComponent.UsesConstant()` for blend constant validation

### v0.3.0 (2026-03-10)
- `TextureUsage.ContainsUnknownBits()` for validation

### v0.2.0 (2026-01-29)
- webgpu.h spec-compliant enum values
- Binary compatibility with wgpu-native

### v0.1.0 (2026-01-29)
- Initial release with core WebGPU types

## Planned

### v1.0.0
- Stable API matching WebGPU spec
- Full coverage of all WebGPU types
- Comprehensive documentation

## Design Principles

1. **Zero dependencies** — Always
2. **WebGPU spec compliance** — Follow W3C WebGPU exactly
3. **Go-idiomatic zero values** — Zero value of any type should be the spec default
4. **Ecosystem compatibility** — Works with all gogpu projects
