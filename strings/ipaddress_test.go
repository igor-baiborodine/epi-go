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

var validIpAddressTest = []struct {
	in   string
	want []string
}{
	{"1111", []string{"1.1.1.1"}},
}
