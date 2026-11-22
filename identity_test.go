package gputypes

import "testing"

// TestDeviceUUIDString pins the text form and the NVML rendering of the same
// bytes. The vector is the deviceUUID vulkaninfo reports for the NVIDIA
// CMP 30HX on the development host, whose nvidia-smi UUID is
// "GPU-288bbc7b-41e0-9cf2-8ff0-5918509a0240" — the empirical byte-order check
// behind the NVMLString contract.
func TestDeviceUUIDString(t *testing.T) {
	u := DeviceUUID{0x28, 0x8b, 0xbc, 0x7b, 0x41, 0xe0, 0x9c, 0xf2, 0x8f, 0xf0, 0x59, 0x18, 0x50, 0x9a, 0x02, 0x40}
	if got, want := u.String(), "288bbc7b-41e0-9cf2-8ff0-5918509a0240"; got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
	if got, want := u.NVMLString(), "GPU-288bbc7b-41e0-9cf2-8ff0-5918509a0240"; got != want {
		t.Fatalf("NVMLString() = %q, want %q", got, want)
	}
	if u.IsZero() {
		t.Fatal("a non-zero UUID reports IsZero")
	}

	var zero DeviceUUID
	if got, want := zero.String(), "00000000-0000-0000-0000-000000000000"; got != want {
		t.Fatalf("zero String() = %q, want %q", got, want)
	}
	if !zero.IsZero() {
		t.Fatal("the all-zero UUID does not report IsZero")
	}
}

// TestDeviceUUIDComparable pins that DeviceUUID compares with ==, so callers
// can use it as a map key and detect collisions by equality.
func TestDeviceUUIDComparable(t *testing.T) {
	a := DeviceUUID{1}
	b := DeviceUUID{1}
	c := DeviceUUID{2}
	if a != b || a == c {
		t.Fatal("DeviceUUID equality does not compare byte content")
	}
	seen := map[DeviceUUID]int{a: 1}
	if _, ok := seen[b]; !ok {
		t.Fatal("a DeviceUUID with equal bytes is not found in a map")
	}
}

// TestPCIBusIDString pins both renderings. The first vector is the CMP 30HX on
// the development host: Vulkan reports domain 0, bus 4, device 0, function 0,
// and nvidia-smi prints "00000000:04:00.0".
func TestPCIBusIDString(t *testing.T) {
	tests := []struct {
		id         PCIBusID
		wantStr    string
		wantNVML   string
		wantIsZero bool
	}{
		{PCIBusID{Domain: 0, Bus: 4, Device: 0, Function: 0}, "0000:04:00.0", "00000000:04:00.0", false},
		{PCIBusID{Domain: 0xFFFF, Bus: 0xAB, Device: 0xCD, Function: 0xA}, "ffff:ab:cd.a", "0000FFFF:AB:CD.A", false},
		{PCIBusID{Domain: 0x12345678, Bus: 1, Device: 2, Function: 3}, "12345678:01:02.3", "12345678:01:02.3", false},
		{PCIBusID{}, "0000:00:00.0", "00000000:00:00.0", true},
	}
	for _, tt := range tests {
		if got := tt.id.String(); got != tt.wantStr {
			t.Errorf("%+v String() = %q, want %q", tt.id, got, tt.wantStr)
		}
		if got := tt.id.NVMLString(); got != tt.wantNVML {
			t.Errorf("%+v NVMLString() = %q, want %q", tt.id, got, tt.wantNVML)
		}
		if got := tt.id.IsZero(); got != tt.wantIsZero {
			t.Errorf("%+v IsZero() = %v, want %v", tt.id, got, tt.wantIsZero)
		}
	}
}

// TestAdapterIdentityIsZero pins the aggregate zero semantics: every field
// absent is the normal "no identity" answer, and any single reported field
// makes the value non-zero.
func TestAdapterIdentityIsZero(t *testing.T) {
	var id AdapterIdentity
	if !id.IsZero() {
		t.Fatal("the zero AdapterIdentity does not report IsZero")
	}

	tests := []AdapterIdentity{
		{DeviceUUID: DeviceUUID{1}},
		{PCIBusID: PCIBusID{Bus: 1}},
		{AdapterLUID: 1},
		{MetalRegistryID: 1},
	}
	for i, tt := range tests {
		if tt.IsZero() {
			t.Errorf("case %d: a single reported field still reports IsZero: %+v", i, tt)
		}
	}
}
