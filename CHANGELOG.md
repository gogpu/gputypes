# Changelog

All notable changes to gputypes will be documented in this file.

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
