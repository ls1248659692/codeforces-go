package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func waysToSplit(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r := 2; r < n && 3*sum[r] <= 2*sum[n]; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sum[n]) + 1
		ans += sort.Search(r-l1, func(l int) bool { return 2*sum[l+l1] > sum[r] })
	}
	return ans % (1e9 + 7)
}
