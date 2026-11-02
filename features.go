package gputypes

// Feature represents a WebGPU feature flag.
//
// Features are optional capabilities that may not be available on all GPUs.
// Check Features.Contains() to determine if a feature is supported.
type Feature uint64

const (
	// FeatureDepthClipControl enables depth clip control (unclipped depth).
	FeatureDepthClipControl Feature = 1 << iota
	// FeatureDepth32FloatStencil8 enables the depth32float-stencil8 texture format.
	FeatureDepth32FloatStencil8
	// FeatureTextureCompressionBC enables BC texture compression formats.
	FeatureTextureCompressionBC
	// FeatureTextureCompressionETC2 enables ETC2 texture compression formats.
	FeatureTextureCompressionETC2
	// FeatureTextureCompressionASTC enables ASTC texture compression formats.
	FeatureTextureCompressionASTC
	// FeatureIndirectFirstInstance enables first instance parameter in indirect draws.
	FeatureIndirectFirstInstance
	// FeatureShaderF16 enables f16 (half-precision float) in shaders.
	FeatureShaderF16
	// FeatureRG11B10UfloatRenderable enables RG11B10Ufloat as a render target format.
	FeatureRG11B10UfloatRenderable
	// FeatureBGRA8UnormStorage enables BGRA8Unorm as a storage texture format.
	FeatureBGRA8UnormStorage
	// FeatureFloat32Filterable enables filtering of float32 textures.
	FeatureFloat32Filterable
	// FeatureTimestampQuery enables timestamp queries.
	FeatureTimestampQuery
	// FeaturePipelineStatisticsQuery enables pipeline statistics queries.
	FeaturePipelineStatisticsQuery
	// FeatureMultiDrawIndirect enables multi-draw indirect commands.
	FeatureMultiDrawIndirect
	// FeatureMultiDrawIndirectCount enables multi-draw indirect with count.
	FeatureMultiDrawIndirectCount
	// FeaturePushConstants enables push constants (non-standard extension).
	FeaturePushConstants
	// FeatureTextureAdapterSpecificFormatFeatures enables adapter-specific texture formats.
	FeatureTextureAdapterSpecificFormatFeatures
	// FeatureShaderFloat64 enables f64 (double-precision float) in shaders.
	FeatureShaderFloat64
	// FeatureVertexAttribute64bit enables 64-bit vertex attributes.
	FeatureVertexAttribute64bit
	// FeatureSubgroupOperations enables subgroup operations in shaders.
	FeatureSubgroupOperations
	// FeatureSubgroupBarrier enables subgroup barriers in shaders.
	FeatureSubgroupBarrier

	// FeatureRayQuery enables inline ray queries in compute and fragment shaders.
	// Experimental — not in W3C WebGPU spec (Milestone 4+).
	FeatureRayQuery
	// FeatureRayHitVertexReturn enables triangle vertex data retrieval on ray hit.
	FeatureRayHitVertexReturn
	// FeatureExtendedASVertexFormats enables Float32x2, Float16x2, Float16x4
	// as valid vertex formats for acceleration structure geometry.
	FeatureExtendedASVertexFormats
	// FeatureASBindingArray enables arrays of acceleration structures in shaders.
	FeatureASBindingArray
	// FeatureRayTracingPipelines enables dedicated ray tracing pipelines.
	FeatureRayTracingPipelines

	// Shader capability features. Each one lets shaders use a naga capability
	// that Rust wgpu gates on the feature of the same name
	// (wgpu-naga-bridge features_to_naga_capabilities). Values are appended;
	// existing members are never renumbered.

	// FeaturePrimitiveIndex enables @builtin(primitive_index) (WGSL
	// `enable primitive_index;`). Rust: PRIMITIVE_INDEX.
	FeaturePrimitiveIndex
	// FeatureClipDistances enables @builtin(clip_distances) (WGSL
	// `enable clip_distances;`). Rust: CLIP_DISTANCES.
	FeatureClipDistances
	// FeatureShaderEarlyDepthTest enables @early_depth_test. Rust:
	// SHADER_EARLY_DEPTH_TEST.
	FeatureShaderEarlyDepthTest
	// FeatureShaderInt64 enables i64 and u64 in shaders. Rust: SHADER_INT64.
	FeatureShaderInt64
	// FeatureShaderInt64AtomicMinMax enables atomicMin and atomicMax on 64-bit
	// integers in storage buffers. Rust: SHADER_INT64_ATOMIC_MIN_MAX.
	FeatureShaderInt64AtomicMinMax
	// FeatureShaderInt64AtomicAllOps enables every atomic operation on 64-bit
	// integers. Rust: SHADER_INT64_ATOMIC_ALL_OPS.
	FeatureShaderInt64AtomicAllOps
	// FeatureShaderFloat32Atomic enables atomic add and exchange on f32 in
	// storage buffers. Rust: SHADER_FLOAT32_ATOMIC.
	FeatureShaderFloat32Atomic
	// FeatureShaderI16 enables i16 and u16 in shaders (WGSL
	// `enable wgpu_int16;`). Rust: SHADER_I16.
	FeatureShaderI16
	// FeatureShaderBarycentrics enables @builtin(barycentric) in fragment
	// shaders. Rust: SHADER_BARYCENTRICS.
	FeatureShaderBarycentrics
	// FeatureShaderPerVertex enables per-vertex fragment inputs (WGSL
	// `enable wgpu_per_vertex;`). Rust: SHADER_PER_VERTEX.
	FeatureShaderPerVertex
	// FeatureSubgroupVertex enables subgroup operations in vertex shaders.
	// Rust: SUBGROUP_VERTEX.
	FeatureSubgroupVertex
	// FeatureShaderDrawIndex enables @builtin(draw_index) (WGSL
	// `enable draw_index;`). Rust: SHADER_DRAW_INDEX.
	FeatureShaderDrawIndex

	// featureEnd is one bit past the last member. It is not a feature; tests
	// use it to check that every member has a name.
	featureEnd
)

// featureNames maps each feature to its name. A member added to the const
// block without an entry here fails TestFeatureNamesComplete.
var featureNames = map[Feature]string{
	FeatureDepthClipControl:                     "DepthClipControl",
	FeatureDepth32FloatStencil8:                 "Depth32FloatStencil8",
	FeatureTextureCompressionBC:                 "TextureCompressionBC",
	FeatureTextureCompressionETC2:               "TextureCompressionETC2",
	FeatureTextureCompressionASTC:               "TextureCompressionASTC",
	FeatureIndirectFirstInstance:                "IndirectFirstInstance",
	FeatureShaderF16:                            "ShaderF16",
	FeatureRG11B10UfloatRenderable:              "RG11B10UfloatRenderable",
	FeatureBGRA8UnormStorage:                    "BGRA8UnormStorage",
	FeatureFloat32Filterable:                    "Float32Filterable",
	FeatureTimestampQuery:                       "TimestampQuery",
	FeaturePipelineStatisticsQuery:              "PipelineStatisticsQuery",
	FeatureMultiDrawIndirect:                    "MultiDrawIndirect",
	FeatureMultiDrawIndirectCount:               "MultiDrawIndirectCount",
	FeaturePushConstants:                        "PushConstants",
	FeatureTextureAdapterSpecificFormatFeatures: "TextureAdapterSpecificFormatFeatures",
	FeatureShaderFloat64:                        "ShaderFloat64",
	FeatureVertexAttribute64bit:                 "VertexAttribute64bit",
	FeatureSubgroupOperations:                   "SubgroupOperations",
	FeatureSubgroupBarrier:                      "SubgroupBarrier",
	FeatureRayQuery:                             "RayQuery",
	FeatureRayHitVertexReturn:                   "RayHitVertexReturn",
	FeatureExtendedASVertexFormats:              "ExtendedASVertexFormats",
	FeatureASBindingArray:                       "ASBindingArray",
	FeatureRayTracingPipelines:                  "RayTracingPipelines",
	FeaturePrimitiveIndex:                       "PrimitiveIndex",
	FeatureClipDistances:                        "ClipDistances",
	FeatureShaderEarlyDepthTest:                 "ShaderEarlyDepthTest",
	FeatureShaderInt64:                          "ShaderInt64",
	FeatureShaderInt64AtomicMinMax:              "ShaderInt64AtomicMinMax",
	FeatureShaderInt64AtomicAllOps:              "ShaderInt64AtomicAllOps",
	FeatureShaderFloat32Atomic:                  "ShaderFloat32Atomic",
	FeatureShaderI16:                            "ShaderI16",
	FeatureShaderBarycentrics:                   "ShaderBarycentrics",
	FeatureShaderPerVertex:                      "ShaderPerVertex",
	FeatureSubgroupVertex:                       "SubgroupVertex",
	FeatureShaderDrawIndex:                      "ShaderDrawIndex",
}

// String returns the feature name, or "Unknown" for a value that is not a
// single defined feature.
func (f Feature) String() string {
	if name, ok := featureNames[f]; ok {
		return name
	}
	return "Unknown"
}

// Features is a set of feature flags.
type Features uint64

// Contains checks if the feature set contains a specific feature.
func (f Features) Contains(feature Feature) bool {
	return f&Features(feature) == Features(feature)
}

// ContainsAll checks if the feature set contains all specified features.
func (f Features) ContainsAll(other Features) bool {
	return f&other == other
}

// Insert adds a feature to the set.
func (f *Features) Insert(feature Feature) {
	*f |= Features(feature)
}

// Remove removes a feature from the set.
func (f *Features) Remove(feature Feature) {
	*f &^= Features(feature)
}

// Intersect returns features common to both sets.
func (f Features) Intersect(other Features) Features {
	return f & other
}

// Union returns all features from both sets.
func (f Features) Union(other Features) Features {
	return f | other
}

// IsEmpty returns true if no features are set.
func (f Features) IsEmpty() bool {
	return f == 0
}

// Count returns the number of features in the set.
func (f Features) Count() int {
	count := 0
	for v := f; v != 0; v &= v - 1 {
		count++
	}
	return count
}
