package gputypes

import "fmt"

// DeviceUUID is a 16-byte universally unique identifier for a physical GPU, as
// Vulkan reports it (VkPhysicalDeviceIDProperties::deviceUUID, core since
// Vulkan 1.1). The all-zero value means the backend did not report one.
//
// Promise: the Vulkan specification requires the value to be immutable for a
// given device "across instances, processes, driver APIs, driver versions, and
// system reboots"; it is per device, not per model. It is not a board serial:
// it may change if the hardware is moved between slots while the system is
// powered off, and conformance testing cannot prove that two devices never
// collide, so a caller matching adapters must treat all-zero as unknown and
// guard against collisions (fall back to PCIBusID, then to the enumeration
// index as a tie-breaker).
//
// NVMLString renders the value the way NVIDIA's NVML does ("GPU-" plus the
// dashed text form). That is a rendering of the same 128 bits — on NVIDIA
// hardware Vulkan and CUDA report the same bytes, so the string equals
// `nvidia-smi`'s UUID — not a claim that the adapter is an NVML device: every
// Vulkan adapter's deviceUUID renders this way.
type DeviceUUID [16]byte

// String returns the lowercase dashed 8-4-4-4-12 text form of the UUID.
func (u DeviceUUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}

// NVMLString returns the UUID in the form NVML uses: "GPU-" followed by the
// dashed text form. It is a rendering only; see DeviceUUID.
func (u DeviceUUID) NVMLString() string {
	return "GPU-" + u.String()
}

// IsZero reports whether the UUID is the all-zero value, which means the
// backend did not report identity.
func (u DeviceUUID) IsZero() bool {
	return u == DeviceUUID{}
}

// PCIBusID is a PCI slot locator: the domain, bus, device and function numbers
// of the slot a GPU occupies, as Vulkan reports through VK_EXT_pci_bus_info.
// The all-zero value means the backend did not report one; a device genuinely
// at 0000:00:00.0 cannot be distinguished from that.
//
// Promise: the locator is persistent while the card stays in its slot, and it
// is per slot — it distinguishes cards that are otherwise identical (same
// model, same vendor and device IDs). The extension is optional: a zero value
// is a normal answer, not an error.
type PCIBusID struct {
	// Domain is the PCI domain (segment) number.
	Domain uint32
	// Bus is the PCI bus number.
	Bus uint32
	// Device is the device number on the bus.
	Device uint32
	// Function is the function number on the device.
	Function uint32
}

// String returns the lowercase 4-digit-domain form "dddd:bb:dd.f" — the form
// Rust wgpu renders (e.g. "0000:01:00.0").
func (p PCIBusID) String() string {
	return fmt.Sprintf("%04x:%02x:%02x.%x", p.Domain, p.Bus, p.Device, p.Function)
}

// NVMLString returns the form NVML uses: an 8-digit zero-padded uppercase
// domain, two-digit bus and device, and the function digit (e.g.
// "00000000:01:00.0"). As with DeviceUUID, this is a rendering, not a claim
// that the adapter is managed by NVML, and callers comparing against
// driver-provided strings should compare case-insensitively.
func (p PCIBusID) NVMLString() string {
	return fmt.Sprintf("%08X:%02X:%02X.%X", p.Domain, p.Bus, p.Device, p.Function)
}

// IsZero reports whether the locator is the all-zero value, which means the
// backend did not report PCI bus information.
func (p PCIBusID) IsZero() bool {
	return p == PCIBusID{}
}

// AdapterIdentity carries the hardware identity a backend can report for a
// GPU. It is a native extension beyond the WebGPU specification, which defines
// no per-device identifier: every field may be absent, and the all-zero value
// of a field means "this backend did not report it" — never an error.
//
// The fields have different lifetimes; the promise is per identifier:
//
//   - DeviceUUID is persistent per device (Vulkan deviceUUID).
//   - PCIBusID is persistent per slot (VK_EXT_pci_bus_info).
//   - AdapterLUID is session-local: Windows changes a LUID on reboot, driver
//     update, and adapter disable or disconnect, so it matches within one boot
//     only. On Windows, Vulkan's deviceLUID equals the DXGI adapter LUID.
//   - MetalRegistryID is instance-local: an IORegistry handle that is not
//     constant across OS reboots, so it identifies a GPU only within the
//     running process.
//
// A backend fills the fields it can supply: GLES, the software backend and the
// browser report all zeroes, and so does the wgpu-native binding until it
// exposes identity data.
type AdapterIdentity struct {
	// DeviceUUID is the persistent per-device identifier, or zero.
	DeviceUUID DeviceUUID
	// PCIBusID is the persistent per-slot locator, or zero.
	PCIBusID PCIBusID
	// AdapterLUID is the session-local 64-bit adapter LUID, or zero.
	// It is the LUID's numeric value: uint64(uint32(HighPart))<<32 or-ed with
	// uint64(LowPart), which is also Vulkan's 8 deviceLUID bytes read
	// little-endian.
	AdapterLUID uint64
	// MetalRegistryID is the instance-local Metal MTLDevice registry ID, or
	// zero.
	MetalRegistryID uint64
}

// IsZero reports whether no backend reported any identity data.
func (id AdapterIdentity) IsZero() bool {
	return id.DeviceUUID.IsZero() && id.PCIBusID.IsZero() &&
		id.AdapterLUID == 0 && id.MetalRegistryID == 0
}
