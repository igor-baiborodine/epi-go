package strings

import (
	"testing"
)

func TestSnakeString(t *testing.T) {
	for _, test := range snakeStringTestData {
		if got := SnakeString(test.s); got != test.want {
			t.Errorf("SnakeString(%s) = %s; want %s", test.s, got, test.want)
		}
	}
}

var snakeStringTestData = []struct {
	s    string
	want string
}{
	{"Hello World!", "e lHloWrdlo!"},
	{"slkt", "lskt"},
	{"zbxei", "bzxie"},
	{"jizilhpyvstjbua", "ihsujzlpvtbaiyj"},
	{"qdfj", "dqfj"},
	{"zwkpifnryzdhwgoedxucoxryluwfhrwsiptrzkmwopswgvkfdj", "wfzgxxurpkpvjzkinydwoduorlwhwitzmosgkdprhecyfsrwwf"},
	{"ulcuenvuoj", "lnjucevouu"},
	{"wumzfvxmbrpit", "uvrwmfxbptzmi"},
	{"glctvq", "lqgcvt"},
	{"gcsmst", "ctgssm"},
	{"atzzywvhl", "twazyvlzh"},
	{"wkzclgyqbwwjmgyjntekj", "kgwgtwzlybwmynejcqjjk"},
	{"umfcg", "mufgc"},
	{"xnkxgzcppwpepxktytlldgxkyhfpxgtwzvfslboxzfkrmkuj", "nzwxtghgvbfkxkgcpppkyldxyfxtzflozkmuxpetlkpwsxrj"},
	{"qveha", "vqeah"},
	{"kpmvzc", "pckmzv"},
	{"ltgdo", "tlgod"},
	{"rnjhnk", "nkrjnh"},
}
