package words

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name, input string
		want        int
	}{
		{"empty", "", 0},
		{"mixed whitespace", " Go\tworks\nwell ", 3},
		{"punctuation", "hello, world!", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.input); got != tc.want {
				t.Errorf("Count(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	for b.Loop() {
		Count(" Go\tworks\nwell ")
	}
}
