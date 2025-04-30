package strings

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	for _, test := range reverseWordsTestData {
		if got := ReverseWords(test.s); got != test.want {
			t.Errorf("ReverseWords(%s) = %s; want %s", test.s, got, test.want)
		}
	}
}

var reverseWordsTestData = []struct {
	s    string
	want string
}{
	{"Alice likes Bob", "Bob likes Alice"},
	{"ram is costly", "costly is ram"},
	{"bS SP1u5  r6ao", "r6ao  SP1u5 bS"},
	{"RrMRw", "RrMRw"},
	{"J  h CfRq", "CfRq h  J"},
	{"p 6fEZcNVZ W  vklZO 8TX nuM  RS xB4rJu1  Io ZASeLvxc Rl DP  A2v9Yf  TM 5B 7 O X6BHu mwSio iIfVoC", "iIfVoC mwSio X6BHu O 7 5B TM  A2v9Yf  DP Rl ZASeLvxc Io  xB4rJu1 RS  nuM 8TX vklZO  W 6fEZcNVZ p"},
	{"V", "V"},
	{"KuAY rz7bx 1w4I dlLdfcypbVc A  VE ", " VE  A dlLdfcypbVc 1w4I rz7bx KuAY"},
	{"VSIL8X1lj", "VSIL8X1lj"},
	{"8Ro3s41mVF E", "E 8Ro3s41mVF"},
	{"a", "a"},
	{"pACYKogmFP n", "n pACYKogmFP"},
	{"uvQJxnjH3BZ9iJ", "uvQJxnjH3BZ9iJ"},
	{"  9ZoAzq", "9ZoAzq  "},
	{"JTV3cWuG", "JTV3cWuG"},
	{"HN 84 ", " 84 HN"},
}
