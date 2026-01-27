# gputypes

WebGPU type definitions for the [gogpu](https://github.com/gogpu) ecosystem.

## Overview

`gputypes` provides all WebGPU enums, structs, and constants as pure Go types with zero dependencies. It follows the pattern of [wgpu-types](https://docs.rs/wgpu-types) in the Rust ecosystem.

## Installation

```bash
go get github.com/gogpu/gputypes
```

**Requires:** Go 1.25+

## Features

- **Enums** — TextureFormat, BufferUsage, ShaderStage, PrimitiveTopology, etc.
- **Structs** — Extent3D, Origin3D, Color, Limits, Features, etc.
- **Constants** — Default limits, format properties, alignment requirements
- **Zero dependencies** — Pure Go, no external imports

## Planned Types

Based on [WebGPU spec](https://www.w3.org/TR/webgpu/) and [wgpu-types](https://docs.rs/wgpu-types):

### Texture & Sampler
- `TextureFormat` (100+ formats)
- `TextureUsage`, `TextureDimension`, `TextureViewDimension`
- `AddressMode`, `FilterMode`, `CompareFunction`

### Buffer & Binding
- `BufferUsage`, `BufferBindingType`
- `BindingType`, `BindGroupLayoutEntry`
- `ShaderStage` flags

### Pipeline
- `PrimitiveTopology`, `FrontFace`, `CullMode`
- `BlendState`, `BlendFactor`, `BlendOperation`
- `DepthStencilState`, `StencilOperation`
- `VertexFormat`, `VertexStepMode`

### Limits & Features
- `Limits` struct with all WebGPU limits
- `Features` flags for optional capabilities
- `DownlevelCapabilities` for backend restrictions

## Relationship to gpucontext

| Package | Purpose |
|---------|---------|
| `gpucontext` | Interfaces (DeviceProvider, EventSource, Registry) |
| `gputypes` | Data types (enums, structs, constants) |

Both packages have **zero dependencies** and form the shared foundation of the gogpu ecosystem.

## Ecosystem

| Package | Description |
|---------|-------------|
| [gogpu/gogpu](https://github.com/gogpu/gogpu) | Graphics framework |
| [gogpu/wgpu](https://github.com/gogpu/wgpu) | Pure Go WebGPU implementation |
| [gogpu/gg](https://github.com/gogpu/gg) | 2D graphics library |
| [gogpu/gpucontext](https://github.com/gogpu/gpucontext) | Shared interfaces |

## License

MIT License — see [LICENSE](LICENSE) for details.
