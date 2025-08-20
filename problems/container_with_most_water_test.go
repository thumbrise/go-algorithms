package problems

import "testing"

func Test_MaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Two ones",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Symmetric descending and ascending",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "Small heights",
			height: []int{1, 2, 1},
			want:   2,
		},
		{
			name:   "All zeros",
			height: []int{0, 0, 0, 0},
			want:   0,
		},
		{
			name:   "Mixed heights example 1",
			height: []int{1, 3, 2, 5, 25, 24, 5},
			want:   24,
		},
		{
			name:   "Mixed heights example 2",
			height: []int{2, 3, 4, 5, 18, 17, 6},
			want:   17,
		},
		{
			name:   "All equal heights",
			height: []int{5, 5, 5, 5, 5},
			want:   20,
		},
		{
			name:   "Increasing then decreasing",
			height: []int{1, 2, 4, 3},
			want:   4,
		},
		{
			name:   "Peaks and troughs",
			height: []int{1, 2, 1, 3, 1, 2},
			want:   8,
		},
		{
			name:   "Descending heights",
			height: []int{10, 9, 8, 7, 6, 5},
			want:   25,
		},
		{
			name:   "Optimization test",
			height: []int{5, 1, 10, 1, 5},
			want:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
