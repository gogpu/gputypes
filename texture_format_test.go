package gputypes

import "testing"

// Expected values below are transcribed from the match arms of Rust wgpu-types,
// file wgpu-types/src/texture/format.rs at wgpu commit 14328f4:
// block_dimensions, is_compressed, block_copy_size, sample_type,
// components_with_aspect, remove_srgb_suffix and add_srgb_suffix.
// They are not derived from this package's implementation.

// formatRow is the expected result of every helper for one format.
//
//   - copySize: block_copy_size with the All aspect (0 = None)
//   - comps: components() (All aspect)
//   - sample: sample_type with no aspect and no features (Undefined = None)
//   - noSrgb / addSrgb: remove_srgb_suffix / add_srgb_suffix (zero = unchanged)
type formatRow struct {
	f          TextureFormat
	bw, bh     uint32
	copySize   uint32
	comps      uint8
	sample     TextureSampleType
	noSrgb     TextureFormat
	addSrgb    TextureFormat
	hasSrgbSfx bool
}

const (
	fFloat   = TextureSampleTypeFloat
	fUnfilt  = TextureSampleTypeUnfilterableFloat
	fDepth   = TextureSampleTypeDepth
	fUint    = TextureSampleTypeUint
	fSint    = TextureSampleTypeSint
	fNoneTyp = TextureSampleTypeUndefined
)

// plain is an uncompressed format without an sRGB twin.
func plain(f TextureFormat, copySize uint32, comps uint8, sample TextureSampleType) formatRow {
	return formatRow{f: f, bw: 1, bh: 1, copySize: copySize, comps: comps, sample: sample}
}

// unorm is a format whose add_srgb_suffix is srgb.
func unorm(f TextureFormat, bw, bh, copySize uint32, comps uint8, srgb TextureFormat) formatRow {
	return formatRow{f: f, bw: bw, bh: bh, copySize: copySize, comps: comps, sample: fFloat, addSrgb: srgb}
}

// srgbFmt is a format whose remove_srgb_suffix is lin.
func srgbFmt(f TextureFormat, bw, bh, copySize uint32, comps uint8, lin TextureFormat) formatRow {
	return formatRow{f: f, bw: bw, bh: bh, copySize: copySize, comps: comps, sample: fFloat, noSrgb: lin, hasSrgbSfx: true}
}

// block is a compressed format without an sRGB twin. Every such format in
// the table (BC4, BC5, BC6H, EAC) has 4x4 blocks.
func block(f TextureFormat, copySize uint32, comps uint8) formatRow {
	return formatRow{f: f, bw: 4, bh: 4, copySize: copySize, comps: comps, sample: fFloat}
}

func formatRows() []formatRow {
	rows := []formatRow{
		// 8-bit
		plain(TextureFormatR8Unorm, 1, 1, fFloat),
		plain(TextureFormatR8Snorm, 1, 1, fFloat),
		plain(TextureFormatR8Uint, 1, 1, fUint),
		plain(TextureFormatR8Sint, 1, 1, fSint),
		// 16-bit
		plain(TextureFormatR16Unorm, 2, 1, fFloat),
		plain(TextureFormatR16Snorm, 2, 1, fFloat),
		plain(TextureFormatR16Uint, 2, 1, fUint),
		plain(TextureFormatR16Sint, 2, 1, fSint),
		plain(TextureFormatR16Float, 2, 1, fFloat),
		plain(TextureFormatRG8Unorm, 2, 2, fFloat),
		plain(TextureFormatRG8Snorm, 2, 2, fFloat),
		plain(TextureFormatRG8Uint, 2, 2, fUint),
		plain(TextureFormatRG8Sint, 2, 2, fSint),
		// 32-bit
		plain(TextureFormatR32Float, 4, 1, fUnfilt),
		plain(TextureFormatR32Uint, 4, 1, fUint),
		plain(TextureFormatR32Sint, 4, 1, fSint),
		plain(TextureFormatRG16Unorm, 4, 2, fFloat),
		plain(TextureFormatRG16Snorm, 4, 2, fFloat),
		plain(TextureFormatRG16Uint, 4, 2, fUint),
		plain(TextureFormatRG16Sint, 4, 2, fSint),
		plain(TextureFormatRG16Float, 4, 2, fFloat),
		unorm(TextureFormatRGBA8Unorm, 1, 1, 4, 4, TextureFormatRGBA8UnormSrgb),
		srgbFmt(TextureFormatRGBA8UnormSrgb, 1, 1, 4, 4, TextureFormatRGBA8Unorm),
		plain(TextureFormatRGBA8Snorm, 4, 4, fFloat),
		plain(TextureFormatRGBA8Uint, 4, 4, fUint),
		plain(TextureFormatRGBA8Sint, 4, 4, fSint),
		unorm(TextureFormatBGRA8Unorm, 1, 1, 4, 4, TextureFormatBGRA8UnormSrgb),
		srgbFmt(TextureFormatBGRA8UnormSrgb, 1, 1, 4, 4, TextureFormatBGRA8Unorm),
		plain(TextureFormatRGB10A2Uint, 4, 4, fUint),
		plain(TextureFormatRGB10A2Unorm, 4, 4, fFloat),
		plain(TextureFormatRG11B10Ufloat, 4, 3, fFloat),
		plain(TextureFormatRGB9E5Ufloat, 4, 3, fFloat),
		// 64-bit
		plain(TextureFormatRG32Float, 8, 2, fUnfilt),
		plain(TextureFormatRG32Uint, 8, 2, fUint),
		plain(TextureFormatRG32Sint, 8, 2, fSint),
		plain(TextureFormatRGBA16Unorm, 8, 4, fFloat),
		plain(TextureFormatRGBA16Snorm, 8, 4, fFloat),
		plain(TextureFormatRGBA16Uint, 8, 4, fUint),
		plain(TextureFormatRGBA16Sint, 8, 4, fSint),
		plain(TextureFormatRGBA16Float, 8, 4, fFloat),
		// 128-bit
		plain(TextureFormatRGBA32Float, 16, 4, fUnfilt),
		plain(TextureFormatRGBA32Uint, 16, 4, fUint),
		plain(TextureFormatRGBA32Sint, 16, 4, fSint),
		// depth/stencil: combined formats have no copy size and no sample type without an aspect
		plain(TextureFormatStencil8, 1, 1, fUint),
		plain(TextureFormatDepth16Unorm, 2, 1, fDepth),
		plain(TextureFormatDepth24Plus, 0, 1, fDepth),
		plain(TextureFormatDepth24PlusStencil8, 0, 2, fNoneTyp),
		plain(TextureFormatDepth32Float, 4, 1, fDepth),
		plain(TextureFormatDepth32FloatStencil8, 0, 2, fNoneTyp),
		// BC
		unorm(TextureFormatBC1RGBAUnorm, 4, 4, 8, 4, TextureFormatBC1RGBAUnormSrgb),
		srgbFmt(TextureFormatBC1RGBAUnormSrgb, 4, 4, 8, 4, TextureFormatBC1RGBAUnorm),
		unorm(TextureFormatBC2RGBAUnorm, 4, 4, 16, 4, TextureFormatBC2RGBAUnormSrgb),
		srgbFmt(TextureFormatBC2RGBAUnormSrgb, 4, 4, 16, 4, TextureFormatBC2RGBAUnorm),
		unorm(TextureFormatBC3RGBAUnorm, 4, 4, 16, 4, TextureFormatBC3RGBAUnormSrgb),
		srgbFmt(TextureFormatBC3RGBAUnormSrgb, 4, 4, 16, 4, TextureFormatBC3RGBAUnorm),
		block(TextureFormatBC4RUnorm, 8, 1),
		block(TextureFormatBC4RSnorm, 8, 1),
		block(TextureFormatBC5RGUnorm, 16, 2),
		block(TextureFormatBC5RGSnorm, 16, 2),
		block(TextureFormatBC6HRGBUfloat, 16, 3),
		block(TextureFormatBC6HRGBFloat, 16, 3),
		unorm(TextureFormatBC7RGBAUnorm, 4, 4, 16, 4, TextureFormatBC7RGBAUnormSrgb),
		srgbFmt(TextureFormatBC7RGBAUnormSrgb, 4, 4, 16, 4, TextureFormatBC7RGBAUnorm),
		// ETC2 / EAC
		unorm(TextureFormatETC2RGB8Unorm, 4, 4, 8, 3, TextureFormatETC2RGB8UnormSrgb),
		srgbFmt(TextureFormatETC2RGB8UnormSrgb, 4, 4, 8, 3, TextureFormatETC2RGB8Unorm),
		unorm(TextureFormatETC2RGB8A1Unorm, 4, 4, 8, 4, TextureFormatETC2RGB8A1UnormSrgb),
		srgbFmt(TextureFormatETC2RGB8A1UnormSrgb, 4, 4, 8, 4, TextureFormatETC2RGB8A1Unorm),
		unorm(TextureFormatETC2RGBA8Unorm, 4, 4, 16, 4, TextureFormatETC2RGBA8UnormSrgb),
		srgbFmt(TextureFormatETC2RGBA8UnormSrgb, 4, 4, 16, 4, TextureFormatETC2RGBA8Unorm),
		block(TextureFormatEACR11Unorm, 8, 1),
		block(TextureFormatEACR11Snorm, 8, 1),
		block(TextureFormatEACRG11Unorm, 16, 2),
		block(TextureFormatEACRG11Snorm, 16, 2),
	}
	return append(rows, astcRows()...)
}

// astcRows lists the 14 ASTC block sizes; every ASTC format has a 16 byte
// block, four components and a Unorm/UnormSrgb pair.
func astcRows() []formatRow {
	pairs := []struct {
		lin, srgb TextureFormat
		bw, bh    uint32
	}{
		{TextureFormatASTC4x4Unorm, TextureFormatASTC4x4UnormSrgb, 4, 4},
		{TextureFormatASTC5x4Unorm, TextureFormatASTC5x4UnormSrgb, 5, 4},
		{TextureFormatASTC5x5Unorm, TextureFormatASTC5x5UnormSrgb, 5, 5},
		{TextureFormatASTC6x5Unorm, TextureFormatASTC6x5UnormSrgb, 6, 5},
		{TextureFormatASTC6x6Unorm, TextureFormatASTC6x6UnormSrgb, 6, 6},
		{TextureFormatASTC8x5Unorm, TextureFormatASTC8x5UnormSrgb, 8, 5},
		{TextureFormatASTC8x6Unorm, TextureFormatASTC8x6UnormSrgb, 8, 6},
		{TextureFormatASTC8x8Unorm, TextureFormatASTC8x8UnormSrgb, 8, 8},
		{TextureFormatASTC10x5Unorm, TextureFormatASTC10x5UnormSrgb, 10, 5},
		{TextureFormatASTC10x6Unorm, TextureFormatASTC10x6UnormSrgb, 10, 6},
		{TextureFormatASTC10x8Unorm, TextureFormatASTC10x8UnormSrgb, 10, 8},
		{TextureFormatASTC10x10Unorm, TextureFormatASTC10x10UnormSrgb, 10, 10},
		{TextureFormatASTC12x10Unorm, TextureFormatASTC12x10UnormSrgb, 12, 10},
		{TextureFormatASTC12x12Unorm, TextureFormatASTC12x12UnormSrgb, 12, 12},
	}
	rows := make([]formatRow, 0, 2*len(pairs))
	for _, p := range pairs {
		rows = append(rows,
			unorm(p.lin, p.bw, p.bh, 16, 4, p.srgb),
			srgbFmt(p.srgb, p.bw, p.bh, 16, 4, p.lin))
	}
	return rows
}

// TestTextureFormatTableComplete is the guard of every table test below: the
// formats are the contiguous values 1..N that have a String() name, and each
// has exactly one row. Adding a format without a row fails here.
func TestTextureFormatTableComplete(t *testing.T) {
	rows := formatRows()
	seen := make(map[TextureFormat]bool, len(rows))
	for _, r := range rows {
		if seen[r.f] {
			t.Errorf("%v has two rows", r.f)
		}
		seen[r.f] = true
	}
	members := 0
	for f := TextureFormat(1); f.String() != "Unknown"; f++ {
		members++
		if !seen[f] {
			t.Errorf("format %v (%#x) has no row in formatRows", f, uint32(f))
		}
	}
	if members != len(rows) {
		t.Errorf("enum has %d named formats, table %d rows", members, len(rows))
	}
	if TextureFormat(members+1).String() != "Unknown" {
		t.Errorf("format %d unexpectedly has a name; the enum is not contiguous", members+1)
	}
}

func TestTextureFormatBlockDimensions(t *testing.T) {
	for _, r := range formatRows() {
		w, h := r.f.BlockDimensions()
		if w != r.bw || h != r.bh {
			t.Errorf("%v.BlockDimensions() = %dx%d, want %dx%d", r.f, w, h, r.bw, r.bh)
		}
		wantCompressed := r.bw != 1 || r.bh != 1
		if got := r.f.IsCompressed(); got != wantCompressed {
			t.Errorf("%v.IsCompressed() = %v, want %v", r.f, got, wantCompressed)
		}
	}
	for _, f := range []TextureFormat{TextureFormatUndefined, TextureFormat(0xFFFF)} {
		if w, h := f.BlockDimensions(); w != 1 || h != 1 {
			t.Errorf("%v.BlockDimensions() = %dx%d, want 1x1", f, w, h)
		}
		if f.IsCompressed() {
			t.Errorf("%v.IsCompressed() = true", f)
		}
	}
}

func TestTextureFormatBlockCopySizeAspect(t *testing.T) {
	for _, r := range formatRows() {
		// All and Undefined: the Rust call with Some(All) / None.
		for _, a := range []TextureAspect{TextureAspectAll, TextureAspectUndefined} {
			size, ok := r.f.BlockCopySizeAspect(a)
			if size != r.copySize || ok != (r.copySize != 0) {
				t.Errorf("%v.BlockCopySizeAspect(%v) = (%d, %v), want (%d, %v)",
					r.f, a, size, ok, r.copySize, r.copySize != 0)
			}
		}
		// BlockCopySize is unchanged: the aspect-independent size, 0 where None.
		if got := r.f.BlockCopySize(); got != r.copySize {
			t.Errorf("%v.BlockCopySize() = %d, want %d", r.f, got, r.copySize)
		}
	}

	combined := []struct {
		name   string
		f      TextureFormat
		aspect TextureAspect
		size   uint32
		ok     bool
	}{
		{"D24S8 stencil", TextureFormatDepth24PlusStencil8, TextureAspectStencilOnly, 1, true},
		{"D24S8 depth", TextureFormatDepth24PlusStencil8, TextureAspectDepthOnly, 0, false},
		{"D32S8 stencil", TextureFormatDepth32FloatStencil8, TextureAspectStencilOnly, 1, true},
		{"D32S8 depth", TextureFormatDepth32FloatStencil8, TextureAspectDepthOnly, 4, true},
		{"Depth24Plus depth", TextureFormatDepth24Plus, TextureAspectDepthOnly, 0, false},
		{"Depth32Float depth", TextureFormatDepth32Float, TextureAspectDepthOnly, 4, true},
		{"Depth16Unorm depth", TextureFormatDepth16Unorm, TextureAspectDepthOnly, 2, true},
		{"Stencil8 stencil", TextureFormatStencil8, TextureAspectStencilOnly, 1, true},
		{"RGBA8 ignores aspect", TextureFormatRGBA8Unorm, TextureAspectDepthOnly, 4, true},
		{"Undefined format", TextureFormatUndefined, TextureAspectAll, 0, false},
		{"unknown format", TextureFormat(0xFFFF), TextureAspectAll, 0, false},
	}
	for _, c := range combined {
		size, ok := c.f.BlockCopySizeAspect(c.aspect)
		if size != c.size || ok != c.ok {
			t.Errorf("%s: BlockCopySizeAspect = (%d, %v), want (%d, %v)", c.name, size, ok, c.size, c.ok)
		}
	}
}

func TestTextureFormatComponents(t *testing.T) {
	for _, r := range formatRows() {
		if got := r.f.Components(); got != r.comps {
			t.Errorf("%v.Components() = %d, want %d", r.f, got, r.comps)
		}
		// The aspect only matters for the combined formats (checked below).
		if r.f != TextureFormatDepth24PlusStencil8 && r.f != TextureFormatDepth32FloatStencil8 {
			for _, a := range []TextureAspect{TextureAspectUndefined, TextureAspectDepthOnly, TextureAspectStencilOnly} {
				if got := r.f.ComponentsWithAspect(a); got != r.comps {
					t.Errorf("%v.ComponentsWithAspect(%v) = %d, want %d", r.f, a, got, r.comps)
				}
			}
		}
	}
	for _, f := range []TextureFormat{TextureFormatDepth24PlusStencil8, TextureFormatDepth32FloatStencil8} {
		want := map[TextureAspect]uint8{
			TextureAspectAll:         2,
			TextureAspectUndefined:   2,
			TextureAspectDepthOnly:   1,
			TextureAspectStencilOnly: 1,
		}
		for a, n := range want {
			if got := f.ComponentsWithAspect(a); got != n {
				t.Errorf("%v.ComponentsWithAspect(%v) = %d, want %d", f, a, got, n)
			}
		}
	}
	for _, f := range []TextureFormat{TextureFormatUndefined, TextureFormat(0xFFFF)} {
		if got := f.Components(); got != 0 {
			t.Errorf("%v.Components() = %d, want 0", f, got)
		}
	}
}

func TestTextureFormatSampleType(t *testing.T) {
	for _, r := range formatRows() {
		got, ok := r.f.SampleType(TextureAspectAll, 0)
		if got != r.sample || ok != (r.sample != fNoneTyp) {
			t.Errorf("%v.SampleType(All, none) = (%v, %v), want (%v, %v)",
				r.f, got, ok, r.sample, r.sample != fNoneTyp)
		}
	}

	// Float32Filterable turns the three 32-bit float formats filterable and
	// changes nothing else.
	f32 := map[TextureFormat]bool{
		TextureFormatR32Float: true, TextureFormatRG32Float: true, TextureFormatRGBA32Float: true,
	}
	for _, r := range formatRows() {
		got, ok := r.f.SampleType(TextureAspectAll, Features(FeatureFloat32Filterable))
		want, wantOK := r.sample, r.sample != fNoneTyp
		if f32[r.f] {
			want = fFloat
		}
		if got != want || ok != wantOK {
			t.Errorf("%v.SampleType(All, Float32Filterable) = (%v, %v), want (%v, %v)", r.f, got, ok, want, wantOK)
		}
	}

	combined := []struct {
		f      TextureFormat
		aspect TextureAspect
		want   TextureSampleType
	}{
		{TextureFormatDepth24PlusStencil8, TextureAspectDepthOnly, fDepth},
		{TextureFormatDepth24PlusStencil8, TextureAspectStencilOnly, fUint},
		{TextureFormatDepth24PlusStencil8, TextureAspectAll, fNoneTyp},
		{TextureFormatDepth24PlusStencil8, TextureAspectUndefined, fNoneTyp},
		{TextureFormatDepth32FloatStencil8, TextureAspectDepthOnly, fDepth},
		{TextureFormatDepth32FloatStencil8, TextureAspectStencilOnly, fUint},
		{TextureFormatDepth32FloatStencil8, TextureAspectAll, fNoneTyp},
		{TextureFormatDepth32FloatStencil8, TextureAspectUndefined, fNoneTyp},
		// Single-aspect formats ignore the aspect.
		{TextureFormatDepth32Float, TextureAspectStencilOnly, fDepth},
		{TextureFormatStencil8, TextureAspectDepthOnly, fUint},
		{TextureFormatRGBA8Unorm, TextureAspectDepthOnly, fFloat},
	}
	for _, c := range combined {
		got, ok := c.f.SampleType(c.aspect, 0)
		if got != c.want || ok != (c.want != fNoneTyp) {
			t.Errorf("%v.SampleType(%v) = (%v, %v), want (%v, %v)", c.f, c.aspect, got, ok, c.want, c.want != fNoneTyp)
		}
	}
	for _, f := range []TextureFormat{TextureFormatUndefined, TextureFormat(0xFFFF)} {
		if got, ok := f.SampleType(TextureAspectAll, 0); got != fNoneTyp || ok {
			t.Errorf("%v.SampleType = (%v, %v), want (Undefined, false)", f, got, ok)
		}
	}
}

func TestTextureFormatSrgbSuffix(t *testing.T) {
	for _, r := range formatRows() {
		wantRemove := r.noSrgb
		if wantRemove == 0 {
			wantRemove = r.f
		}
		if got := r.f.RemoveSrgbSuffix(); got != wantRemove {
			t.Errorf("%v.RemoveSrgbSuffix() = %v, want %v", r.f, got, wantRemove)
		}
		wantAdd := r.addSrgb
		if wantAdd == 0 {
			wantAdd = r.f
		}
		if got := r.f.AddSrgbSuffix(); got != wantAdd {
			t.Errorf("%v.AddSrgbSuffix() = %v, want %v", r.f, got, wantAdd)
		}
		// The suffix is exactly what IsSrgb reports, and both directions invert each other.
		if r.f.IsSrgb() != r.hasSrgbSfx {
			t.Errorf("%v.IsSrgb() = %v, table says %v", r.f, r.f.IsSrgb(), r.hasSrgbSfx)
		}
		if (r.f.RemoveSrgbSuffix() != r.f) != r.f.IsSrgb() {
			t.Errorf("%v: RemoveSrgbSuffix changes the format = %v, IsSrgb = %v",
				r.f, r.f.RemoveSrgbSuffix() != r.f, r.f.IsSrgb())
		}
		if r.f.IsSrgb() && r.f.RemoveSrgbSuffix().AddSrgbSuffix() != r.f {
			t.Errorf("%v: add(remove(f)) != f", r.f)
		}
	}
	for _, f := range []TextureFormat{TextureFormatUndefined, TextureFormat(0xFFFF)} {
		if f.RemoveSrgbSuffix() != f || f.AddSrgbSuffix() != f {
			t.Errorf("%v: srgb suffix helpers must return unknown formats unchanged", f)
		}
	}
}
