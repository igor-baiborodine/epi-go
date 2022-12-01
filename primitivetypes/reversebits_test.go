package primitivetypes

import "testing"

type reverseBitsFn func(uint64) uint64

func testReverseBitsFn(t *testing.T, fn reverseBitsFn) {
	for _, test := range reverseBitsTestData {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%#.16x) = %#.16x; want %#.16x", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestReverseBits(t *testing.T)          { testReverseBitsFn(t, ReverseBits) }
func TestReverseBitsReverse64(t *testing.T) { testReverseBitsFn(t, ReverseBitsReverse64) }

func benchReverseBitsFn(b *testing.B, fn reverseBitsFn) {
	for i := 0; i < b.N; i++ {
		fn(0xffffffff00000000)
	}
}

func BenchmarkReverseBits(b *testing.B)          { benchReverseBitsFn(b, ReverseBits) }
func BenchmarkReverseBitsReverse64(b *testing.B) { benchReverseBitsFn(b, ReverseBitsReverse64) }

var reverseBitsTestData = []struct {
	in   uint64
	want uint64
}{
	{1351510410656, 405942121183313920},
	{98370191902452400, 960580743675689600},
	{13958, 7019985919163760640},
	{425765368052, 3391185800900116480},
	{3336380, 4402070723661660160},
	{2064762, 6809996790444589056},
	{499135728, 1093252898353053696},
	{136480640450232, 2125109716022329344},
	{212863968, 562964453230903296},
	{156310, 7585820991096356864},
	{13428, 3327034224719953920},
	{1291173549264, 804754521417842688},
	{3052800818, 5535046730858364928},
	{528, 594475150812905472},
	{6, 6917529027641081856},
	{970, 6034823500676464640},
	{4919940, 2398397699016622080},
	{3470851185816786062, 8154430653829436428},
	{32410056252, 4353947100631793664},
	{11907696334052670, 8987399632668742656},
}
