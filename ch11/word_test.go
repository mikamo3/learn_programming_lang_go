package word

import "testing"
import "time"
import "math/rand"

func TestPalidrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"kayak", true},
		{"foobar", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalidrome(%q) = %v`, test.input, got)
		}
	}
}
func TestJpPalidrome(t *testing.T) {
	input := "ほげほ"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalidrome(%q) = false`, input)
	}
}
func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalidrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalidrome(%q) = false", p)
		}
	}
}
