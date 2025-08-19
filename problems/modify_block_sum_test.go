package problems

import (
	"bytes"
	"testing"
)

// Задача:
// Дана последовательность байтов data []byte, представляющая поток событий.
// Требуется обнаружить подряд идущие одинаковые байты (блоки),
// и заменить каждый такой блок на один байт — сумму этого байта и длины блока modulo 256.
// Например, блок из байтов [5,5,5] (байт=5, длина=3) заменяется на один байт (5+3)%256 = 8.
// Если блок длиной 1 — просто оставляем этот байт без изменений.
//
// Необходимо сделать преобразование на месте (in-place) и за один проход (one-pass).
// Вернуть новую длину среза после преобразования.
//
// Эта задача использует выслеживание паттернов (блоков повторов),
// обработку потока, формирование новых данных, аналогично исходной,
// но с другой логикой формирования выходных значений.

func Test_ModifyBlocksSum(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		wantOut []byte
		wantLen int
	}{
		{
			name:    "empty",
			in:      []byte(""),
			wantOut: []byte(""),
			wantLen: 0,
		},
		{
			name:    "single byte",
			in:      []byte("a"),
			wantOut: []byte("a"),
			wantLen: 1,
		},
		{
			name:    "no repeated bytes",
			in:      []byte("abc"),
			wantOut: []byte("abc"),
			wantLen: 3,
		},
		{
			name:    "double bytes",
			in:      []byte("aabb"),
			wantOut: []byte{('a' + 2) % 256, ('b' + 2) % 256},
			wantLen: 2,
		},
		{
			name: "mixed repeats",
			in:   []byte("aaabcddddd"),
			wantOut: []byte{
				('a' + 3) % 256, 'b', 'c', ('d' + 5) % 256,
			},
			wantLen: 4,
		},
		{
			name: "single repeats and singles",
			in:   []byte("aabbc"),
			wantOut: []byte{
				('a' + 2) % 256, ('b' + 2) % 256, 'c',
			},
			wantLen: 3,
		},
		{
			name:    "long repeat",
			in:      []byte("eeeeeeeeee"), // 10 times 'e'
			wantOut: []byte{('e' + 10) % 256},
			wantLen: 1,
		},
		{
			name:    "single at end",
			in:      []byte("fffgh"),
			wantOut: []byte{('f' + 3) % 256, 'g', 'h'},
			wantLen: 3,
		},
		{
			name:    "no repeats at all",
			in:      []byte("xyz"),
			wantOut: []byte("xyz"),
			wantLen: 3,
		},
		{
			name: "alternating repeats",
			in:   []byte("aabbccddeeff"),
			wantOut: []byte{
				('a' + 2) % 256, ('b' + 2) % 256, ('c' + 2) % 256,
				('d' + 2) % 256, ('e' + 2) % 256, ('f' + 2) % 256,
			},
			wantLen: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inCopy := make([]byte, len(tt.in))
			copy(inCopy, tt.in)

			gotLen := modifyBlocksSum(inCopy)
			if gotLen != tt.wantLen {
				t.Errorf("modifyBlocksSum() length = %v, want %v", gotLen, tt.wantLen)
			}
			if !bytes.Equal(inCopy[:gotLen], tt.wantOut) {
				t.Errorf("original=%v\nmodifyBlocksSum() output = %v, want %v", tt.in, inCopy[:gotLen], tt.wantOut)
			}
		})
	}
}

// Функция modifyBlocksSum должна реализовывать описанный алгоритм:
// на месте заменить блоки одинаковых байтов на сумму байта и длины блока (mod 256),
// одиночные — оставлять без изменений,
// возвращать новую длину.
