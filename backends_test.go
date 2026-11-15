package gputypes

import (
	"testing"
	"unsafe"
)

// TestBackendValues pins every Backend value and name. 0..5 match Rust
// wgpu-types' Backend; BackendSoftware is our own 6. A renumbering changes
// what adapters and logs report, so it fails here first.
func TestBackendValues(t *testing.T) {
	tests := []struct {
		name  string
		value Backend
		want  uint8
		text  string
	}{
		{"BackendEmpty", BackendEmpty, 0, "Empty"},
		{"BackendVulkan", BackendVulkan, 1, "Vulkan"},
		{"BackendMetal", BackendMetal, 2, "Metal"},
		{"BackendDX12", BackendDX12, 3, "DX12"},
		{"BackendGL", BackendGL, 4, "GL"},
		{"BackendBrowserWebGPU", BackendBrowserWebGPU, 5, "BrowserWebGPU"},
		{"BackendSoftware", BackendSoftware, 6, "Software"},
	}
	for _, tt := range tests {
		if uint8(tt.value) != tt.want {
			t.Errorf("%s = %d, want %d", tt.name, uint8(tt.value), tt.want)
		}
		if got := tt.value.String(); got != tt.text {
			t.Errorf("%s.String() = %q, want %q", tt.name, got, tt.text)
		}
	}
}

// TestBackendsFlags pins the set: the type is a u32, every flag is
// 1 << BackendX (the invariant Contains relies on), Contains agrees with each
// flag, and the aggregate masks are exact.
func TestBackendsFlags(t *testing.T) {
	if got := unsafe.Sizeof(Backends(0)); got != 4 {
		t.Fatalf("Backends is %d bytes, want 4 (u32 like Rust wgpu-types)", got)
	}

	pairs := []struct {
		name    string
		backend Backend
		flag    Backends
	}{
		{"Vulkan", BackendVulkan, BackendsVulkan},
		{"Metal", BackendMetal, BackendsMetal},
		{"DX12", BackendDX12, BackendsDX12},
		{"GL", BackendGL, BackendsGL},
		{"BrowserWebGPU", BackendBrowserWebGPU, BackendsBrowserWebGPU},
		{"Software", BackendSoftware, BackendsSoftware},
	}
	union := BackendsNone
	for _, p := range pairs {
		if want := Backends(1) << p.backend; p.flag != want {
			t.Errorf("Backends%s = %d, want 1<<Backend%s = %d", p.name, p.flag, p.name, want)
		}
		if !p.flag.Contains(p.backend) {
			t.Errorf("Backends%s.Contains(Backend%s) = false, want true", p.name, p.name)
		}
		if p.flag.Contains(BackendEmpty) {
			t.Errorf("Backends%s.Contains(BackendEmpty) = true, want false", p.name)
		}
		union |= p.flag
	}

	if BackendsSoftware != 64 {
		t.Errorf("BackendsSoftware = %d, want 64", BackendsSoftware)
	}
	if got, want := BackendsAll, BackendsPrimary|BackendsSecondary; got != want {
		t.Errorf("BackendsAll = %d, want %d", got, want)
	}
	if BackendsAll.Contains(BackendSoftware) {
		t.Error("BackendsAll contains software: software is selected explicitly, never by the default set")
	}
	if !BackendsAll.Contains(BackendVulkan) || !BackendsAll.Contains(BackendGL) {
		t.Error("BackendsAll misses a hardware backend")
	}
	if BackendsNone.Contains(BackendSoftware) {
		t.Error("BackendsNone contains software")
	}
	if union == BackendsAll {
		t.Error("the union of every flag equals BackendsAll: BackendsSoftware must stay outside it")
	}
}
