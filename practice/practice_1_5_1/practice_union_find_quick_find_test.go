package practice_1_5_1

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	testcases := []struct {
		p, q int
	}{
		{9, 0}, {3, 4}, {5, 8}, {7, 2}, {2, 1}, {5, 7}, {0, 3}, {4, 2},
	}
	uf := NewUnionFind(10)
	for _, testcase := range testcases {
		if uf.Connected(testcase.p, testcase.q) {
			continue
		}
		uf.Union(testcase.p, testcase.q)
		fmt.Println(testcase.p, testcase.q)
		fmt.Println("id:", uf.id)
	}
	fmt.Println(uf.count, "components")
	fmt.Println(uf.queryCount, "queries")
}
