package problems

import (
	"bytes"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		wantOut []byte
		wantLen int
	}{
		{
			name:    "basic",
			in:      []byte("aaabeccccccccccccdddd"),
			wantOut: []byte("a3bec12d4"),
			wantLen: 9,
		},
		{
			name:    "single",
			in:      []byte("a"),
			wantOut: []byte("a"),
			wantLen: 1,
		},
		{
			name:    "different ending",
			in:      []byte("aabbccdde"),
			wantOut: []byte("a2b2c2d2e"),
			wantLen: 9,
		},

		// Дополнительные кейсы для тренировки

		{
			name:    "empty input",
			in:      []byte(""),
			wantOut: []byte(""),
			wantLen: 0,
		},
		{
			name:    "all unique",
			in:      []byte("abcde"),
			wantOut: []byte("abcde"),
			wantLen: 5,
		},
		{
			name:    "all same chars",
			in:      []byte("xxxxx"),
			wantOut: []byte("x5"),
			wantLen: 2,
		},
		{
			name:    "pairs only",
			in:      []byte("aabbcc"),
			wantOut: []byte("a2b2c2"),
			wantLen: 6,
		},
		{
			name:    "mixed repeats",
			in:      []byte("aaabccccddde"),
			wantOut: []byte("a3bc4d3e"),
			wantLen: 8,
		},
		{
			name:    "long repeat (multi-digit)",
			in:      []byte("zzzzzzzzzzz"), // 11 times 'z'
			wantOut: []byte("z11"),
			wantLen: 3,
		},
		{
			name:    "single doubles",
			in:      []byte("mmnnoo"),
			wantOut: []byte("m2n2o2"),
			wantLen: 6,
		},
		{
			name:    "mixed singles and repeats",
			in:      []byte("aabccddde"),
			wantOut: []byte("a2bc2d3e"),
			wantLen: 8,
		},
		{
			name:    "two repeats only at start",
			in:      []byte("qrrst"),
			wantOut: []byte("qr2st"),
			wantLen: 5,
		},
		{
			name:    "very long repeat block",
			in:      bytes.Repeat([]byte{'x'}, 100+5), // 105 times 'x'
			wantOut: append([]byte("x105")),           // первый символ + "105"
			wantLen: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём копию входных данных, чтобы не менять оригинал в тестах
			inputCopy := make([]byte, len(tt.in))
			copy(inputCopy, tt.in)

			gotLen := compress(inputCopy)
			if gotLen != tt.wantLen {
				t.Errorf("compress() = %v, wantLen %v", gotLen, tt.wantLen)
			}
			if bytes.Compare(inputCopy[:gotLen], tt.wantOut) != 0 {
				t.Errorf("original = %v\ncompress() = %v, wantOut %v", string(tt.in), string(inputCopy[:gotLen]), string(tt.wantOut))
			}
		})
	}
}
