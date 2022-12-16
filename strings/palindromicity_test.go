package strings

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

type palindromicityFn func(s string) bool

func testPalindromicityFn(t *testing.T, fn palindromicityFn) {
	for _, test := range palindromicityTest {
		if got := fn(test.in); got != test.want {
			t.Errorf("%v(%v)=%v; want %v", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestPalindromicity(t *testing.T) { testPalindromicityFn(t, Palindromicity) }

func benchmarkPalindromicityFn(b *testing.B, fn palindromicityFn) {
	b.StopTimer()
	s := make([]byte, math.MaxInt16)
	rand.Seed(time.Now().UnixNano())

	var min = 33 // exclude space and control characters
	var max = math.MaxInt8 - 1
	var i, j = 0, len(s) - 1

	for i <= j {
		s[i] = byte(rand.Intn(max-min) + min)
		s[j] = s[i]
		i++
		j--
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(string(s))
	}
}

func BenchmarkPalindromicity(b *testing.B) { benchmarkPalindromicityFn(b, Palindromicity) }

var palindromicityTest = []struct {
	in   string
	want bool
}{
	{"", true},
	{"a", true},
	{"A man, a plan, a canal, Panama.", true},
	{"Able was I, ere I saw Elba!", true},
	{"Ray a Ray", false},
	{"'MK!|", false},
	{"+N}y!:ehFUfX'!k6'!A_[[pp]78!@)!?is{1A7|~:cXzNC\"})\\F}wi\"t3\\^qhf$}v_(>vzrM\\WYdDU3Nie#}`6pKCqFw}S!dM6e{$e-guAZ]", false},
	{"GZ>\\", false},
	{"^4W)3Fv", false},
	{"pN:%\"<@WU}l2q\\w~/7Sa:boa7", false},
	{"=^9[uW?24XFV2a='KqH>_KrsXoU1~r,WrqNIRG{E\"$4:<<g_=O", false},
	{"=", true},
	{"?a2zJAU2OZ*I9W", false},
	{",<N]/d%-magQQO@SP(b!h(v~!U7`Ck,QOphRYdl)q+sDro'U]^'^=i", false},
	{"rqa9EH6[!V~k2fmeF8hB1(`[bdQ!f}'vd5_~w3Smy~(&*5b[#Nj5XF-cGs_kztKJIBkK4+>D.)L!p*YL=5Ipf,yNRH>7H:#,&N+3j*CWKo:7L#", false},
	{"f=_rzHgs", false},
	{"V\"t", false},
	{"8XtTp<A}4YAmP&/3j#\\~C!-((g<", false},
	{"U_pu-dndb~^tj*\"Pw%a_1a`rw2&7VgR^l(j\"10A];re\\=", false},
	{"7}y", false},
	{"[{q4n", false},
	{"H>>+ZvTI5)YPR/(D\"!cyTx/BPYvG2|eiw45]=7h*7'L", false},
	{"+", true},
}
