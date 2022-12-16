package strings

import (
	"reflect"
	"testing"
)

type validIpAddressFn func(s string) []string

func testValidIpAddressFn(t *testing.T, fn validIpAddressFn) {
	for _, test := range validIpAddressTest {
		if got := fn(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%v)=%v; want %v", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestValidIpAddress(t *testing.T) { testValidIpAddressFn(t, ValidIpAddress) }

func benchmarkValidIpAddressFn(b *testing.B, fn validIpAddressFn) {
	for i := 0; i < b.N; i++ {
		fn("255255255255")
	}
}

func BenchmarkValidIpAddress(b *testing.B) { benchmarkValidIpAddressFn(b, ValidIpAddress) }

var validIpAddressTest = []struct {
	in   string
	want []string
}{
	{"255255255255", []string{"255.255.255.255"}},
	{"1111", []string{"1.1.1.1"}},
	{"11000", []string{"1.10.0.0", "11.0.0.0"}},
	{"188212", []string{"1.8.8.212", "1.8.82.12", "1.88.2.12", "1.88.21.2", "18.8.2.12", "18.8.21.2", "18.82.1.2", "188.2.1.2"}},
	{"19417168171168521851744716", []string{}},
	{"186148", []string{"1.8.6.148", "1.8.61.48", "1.86.1.48", "1.86.14.8", "18.6.1.48", "18.6.14.8", "18.61.4.8", "186.1.4.8"}},
	{"9", []string{}},
	{"71248", []string{"7.1.2.48", "7.1.24.8", "7.12.4.8", "71.2.4.8"}},
	{"13190", []string{"1.3.1.90", "1.3.19.0", "1.31.9.0", "13.1.9.0"}},
	{"153662462532101038023545131162", []string{}},
	{"17799", []string{"1.7.7.99", "1.7.79.9", "1.77.9.9", "17.7.9.9"}},
	{"42442057120149175", []string{}},
	{"", []string{}},
	{"24729", []string{"2.4.7.29", "2.4.72.9", "2.47.2.9", "24.7.2.9"}},
	{"166", []string{}},
	{"5521821011617115847", []string{}},
	{"99220136", []string{"9.9.220.136", "9.92.20.136", "9.92.201.36", "99.2.20.136", "99.2.201.36", "99.22.0.136", "99.220.1.36", "99.220.13.6"}},
	{"55321511462082542112081492055828", []string{}},
	{"1224110056181200140210179138", []string{}},
	{"227158101116", []string{"227.158.101.116"}},
	{"20019382193", []string{"200.193.82.193"}},
	{"4211636717685", []string{}},
	{"3420533572095483", []string{}},
	{"47237", []string{"4.7.2.37", "4.7.23.7", "4.72.3.7", "47.2.3.7"}},
	{"166249971302220120929229", []string{}},
	{"176784624830", []string{}},
}
