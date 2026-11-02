package gputypes

import "testing"

// TestFeatureNamesComplete checks that every member of the Feature const block
// has a name and that the table holds nothing else.
func TestFeatureNamesComplete(t *testing.T) {
	members := 0
	for f := Feature(1); f < featureEnd; f <<= 1 {
		members++
		if f.String() == "Unknown" {
			t.Errorf("feature bit %#x has no name in featureNames", uint64(f))
		}
	}
	if len(featureNames) != members {
		t.Errorf("featureNames has %d entries, the const block %d members", len(featureNames), members)
	}
	if got := featureEnd.String(); got != "Unknown" {
		t.Errorf("featureEnd.String() = %q, want Unknown", got)
	}
	if got := (FeatureShaderF16 | FeatureShaderInt64).String(); got != "Unknown" {
		t.Errorf("a combined value has String() = %q, want Unknown", got)
	}
}

// TestShaderCapabilityFeatures checks the shader capability members: each has
// its own bit after FeatureRayTracingPipelines (existing values never move)
// and a name.
func TestShaderCapabilityFeatures(t *testing.T) {
	members := []struct {
		f    Feature
		name string
	}{
		{FeaturePrimitiveIndex, "PrimitiveIndex"},
		{FeatureClipDistances, "ClipDistances"},
		{FeatureShaderEarlyDepthTest, "ShaderEarlyDepthTest"},
		{FeatureShaderInt64, "ShaderInt64"},
		{FeatureShaderInt64AtomicMinMax, "ShaderInt64AtomicMinMax"},
		{FeatureShaderInt64AtomicAllOps, "ShaderInt64AtomicAllOps"},
		{FeatureShaderFloat32Atomic, "ShaderFloat32Atomic"},
		{FeatureShaderI16, "ShaderI16"},
		{FeatureShaderBarycentrics, "ShaderBarycentrics"},
		{FeatureShaderPerVertex, "ShaderPerVertex"},
		{FeatureSubgroupVertex, "SubgroupVertex"},
		{FeatureShaderDrawIndex, "ShaderDrawIndex"},
	}
	if FeatureRayTracingPipelines != 1<<24 {
		t.Fatalf("FeatureRayTracingPipelines = %#x, want 1<<24: an existing member moved", uint64(FeatureRayTracingPipelines))
	}
	var seen Features
	for i, m := range members {
		if want := Feature(1) << (25 + i); m.f != want {
			t.Errorf("%s = %#x, want %#x", m.name, uint64(m.f), uint64(want))
		}
		if got := m.f.String(); got != m.name {
			t.Errorf("String() = %q, want %q", got, m.name)
		}
		if seen.Contains(m.f) {
			t.Errorf("%s reuses a bit", m.name)
		}
		seen.Insert(m.f)
	}
}
