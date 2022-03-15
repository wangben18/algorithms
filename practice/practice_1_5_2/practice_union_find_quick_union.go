package practice_1_5_2

type UnionFind struct {
	id    []int
	count int

	queryCount int
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) Union(p, q int) {
	pFind := uf.Find(p)
	qFind := uf.Find(q)
	if pFind == qFind {
		return
	}
	uf.queryCount++
	uf.id[p] = qFind
	uf.count--
}

func (uf *UnionFind) Find(p int) int {
	for p != uf.id[p] && func() bool { uf.queryCount++; return true }() {
		uf.queryCount++
		p = uf.id[p]
	}
	return p
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func NewUnionFind(n int) (uf UnionFind) {
	uf.count = n
	uf.id = make([]int, n)
	for i := range uf.id {
		uf.id[i] = i
	}
	return
}
