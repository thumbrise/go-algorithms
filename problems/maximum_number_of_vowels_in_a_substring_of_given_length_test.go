package problems

import (
	"testing"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"Example1", "abciiidef", 3, 3},
		{"Example2", "aeiou", 2, 2},
		{"Example3", "leetcode", 3, 2},
		{"NoVowels", "rhythms", 4, 0},
		{"SingleCharVowel", "a", 1, 1},
		{"AllConsonants", "bcbcbbb", 1, 0},
		{"AllVowels", "aeiaeioo", 5, 5},
		{"MixedCase", "abcdeiouu", 4, 4},
		{"ShortSubstring", "xyzabcde", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxVowels(tt.s, tt.k)
			if got != tt.expected {
				t.Errorf("maxVowels(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
