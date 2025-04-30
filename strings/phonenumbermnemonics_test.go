package strings

import (
	"reflect"
	"testing"
)

func TestComputeMnemonics(t *testing.T) {
	for _, test := range ComputeMnemonicsTestData {
		if got := ComputeMnemonics(test.n); !reflect.DeepEqual(got, test.want) {
			t.Errorf("ComputeMnemonics(%s) = %v; want %v", test.n, got, test.want)
		}
	}
}

var ComputeMnemonicsTestData = []struct {
	n    string
	want []string
}{
	{"0", []string{"0"}},
	{"1", []string{"1"}},
	{"2", []string{"A", "B", "C"}},
	{"9", []string{"W", "X", "Y", "Z"}},
	{"19", []string{"1W", "1X", "1Y", "1Z"}},
	{"2804", []string{"AT0G", "AT0H", "AT0I", "AU0G", "AU0H", "AU0I", "AV0G", "AV0H", "AV0I", "BT0G", "BT0H", "BT0I", "BU0G", "BU0H", "BU0I", "BV0G", "BV0H", "BV0I", "CT0G", "CT0H", "CT0I", "CU0G", "CU0H", "CU0I", "CV0G", "CV0H", "CV0I"}},
	{"11203", []string{"11A0D", "11A0E", "11A0F", "11B0D", "11B0E", "11B0F", "11C0D", "11C0E", "11C0F"}},
	{"470418", []string{"GP0G1T", "GP0G1U", "GP0G1V", "GP0H1T", "GP0H1U", "GP0H1V", "GP0I1T", "GP0I1U", "GP0I1V", "GQ0G1T", "GQ0G1U", "GQ0G1V", "GQ0H1T", "GQ0H1U", "GQ0H1V", "GQ0I1T", "GQ0I1U", "GQ0I1V", "GR0G1T", "GR0G1U", "GR0G1V", "GR0H1T", "GR0H1U", "GR0H1V", "GR0I1T", "GR0I1U", "GR0I1V", "GS0G1T", "GS0G1U", "GS0G1V", "GS0H1T", "GS0H1U", "GS0H1V", "GS0I1T", "GS0I1U", "GS0I1V", "HP0G1T", "HP0G1U", "HP0G1V", "HP0H1T", "HP0H1U", "HP0H1V", "HP0I1T", "HP0I1U", "HP0I1V", "HQ0G1T", "HQ0G1U", "HQ0G1V", "HQ0H1T", "HQ0H1U", "HQ0H1V", "HQ0I1T", "HQ0I1U", "HQ0I1V", "HR0G1T", "HR0G1U", "HR0G1V", "HR0H1T", "HR0H1U", "HR0H1V", "HR0I1T", "HR0I1U", "HR0I1V", "HS0G1T", "HS0G1U", "HS0G1V", "HS0H1T", "HS0H1U", "HS0H1V", "HS0I1T", "HS0I1U", "HS0I1V", "IP0G1T", "IP0G1U", "IP0G1V", "IP0H1T", "IP0H1U", "IP0H1V", "IP0I1T", "IP0I1U", "IP0I1V", "IQ0G1T", "IQ0G1U", "IQ0G1V", "IQ0H1T", "IQ0H1U", "IQ0H1V", "IQ0I1T", "IQ0I1U", "IQ0I1V", "IR0G1T", "IR0G1U", "IR0G1V", "IR0H1T", "IR0H1U", "IR0H1V", "IR0I1T", "IR0I1U", "IR0I1V", "IS0G1T", "IS0G1U", "IS0G1V", "IS0H1T", "IS0H1U", "IS0H1V", "IS0I1T", "IS0I1U", "IS0I1V"}},
	{"173816", []string{"1PDT1M", "1PDT1N", "1PDT1O", "1PDU1M", "1PDU1N", "1PDU1O", "1PDV1M", "1PDV1N", "1PDV1O", "1PET1M", "1PET1N", "1PET1O", "1PEU1M", "1PEU1N", "1PEU1O", "1PEV1M", "1PEV1N", "1PEV1O", "1PFT1M", "1PFT1N", "1PFT1O", "1PFU1M", "1PFU1N", "1PFU1O", "1PFV1M", "1PFV1N", "1PFV1O", "1QDT1M", "1QDT1N", "1QDT1O", "1QDU1M", "1QDU1N", "1QDU1O", "1QDV1M", "1QDV1N", "1QDV1O", "1QET1M", "1QET1N", "1QET1O", "1QEU1M", "1QEU1N", "1QEU1O", "1QEV1M", "1QEV1N", "1QEV1O", "1QFT1M", "1QFT1N", "1QFT1O", "1QFU1M", "1QFU1N", "1QFU1O", "1QFV1M", "1QFV1N", "1QFV1O", "1RDT1M", "1RDT1N", "1RDT1O", "1RDU1M", "1RDU1N", "1RDU1O", "1RDV1M", "1RDV1N", "1RDV1O", "1RET1M", "1RET1N", "1RET1O", "1REU1M", "1REU1N", "1REU1O", "1REV1M", "1REV1N", "1REV1O", "1RFT1M", "1RFT1N", "1RFT1O", "1RFU1M", "1RFU1N", "1RFU1O", "1RFV1M", "1RFV1N", "1RFV1O", "1SDT1M", "1SDT1N", "1SDT1O", "1SDU1M", "1SDU1N", "1SDU1O", "1SDV1M", "1SDV1N", "1SDV1O", "1SET1M", "1SET1N", "1SET1O", "1SEU1M", "1SEU1N", "1SEU1O", "1SEV1M", "1SEV1N", "1SEV1O", "1SFT1M", "1SFT1N", "1SFT1O", "1SFU1M", "1SFU1N", "1SFU1O", "1SFV1M", "1SFV1N", "1SFV1O"}},
}
