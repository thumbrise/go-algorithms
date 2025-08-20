package problems

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example1",
			nums: []int{1, 2, 3, 4},
			k:    5,
			want: 2,
		},
		{
			name: "Example2",
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
			want: 1,
		},
		{
			name: "NoPairs",
			nums: []int{1, 1, 1, 1},
			k:    10,
			want: 0,
		},
		{
			name: "AllPairs",
			nums: []int{2, 2, 3, 3, 4, 4},
			k:    6,
			want: 3,
		},
		{
			name: "SingleElement",
			nums: []int{5},
			k:    5,
			want: 0,
		},
		{
			name: "EmptyArray",
			nums: []int{},
			k:    1,
			want: 0,
		},
		{
			name: "LargeNumbers",
			nums: []int{1000000000, 1, 999999999},
			k:    1000000000,
			want: 1,
		},
		{
			name: "MultiplePairsWithDuplicates",
			nums: []int{1, 3, 2, 2, 3, 1, 4, 4},
			k:    5,
			want: 4,
		},
		{
			name: "ProvocationSet",
			nums: []int{3, 3, 3, 3},
			k:    6,
			want: 2, // Можно убрать две пары (3+3) и (3+3)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxOperations(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("maxOperations(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
