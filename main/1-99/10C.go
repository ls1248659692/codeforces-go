package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF10C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ans := int64(0)
	cnt := [9]int64{}
	for i := 1; i <= n; i++ {
		cnt[i%9]++
		ans -= int64(n / i)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ans += cnt[i] * cnt[j] * cnt[i*j%9]
		}
	}
	Fprint(out, ans)
}

//func main() { CF10C(os.Stdin, os.Stdout) }
