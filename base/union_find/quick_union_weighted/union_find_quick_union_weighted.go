package quick_union_weighted

type UnionFind struct {
	id     []int
	weight []int
	count  int
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
	if uf.weight[pFind] > uf.weight[qFind] {
		uf.id[q] = pFind
		uf.weight[pFind] += uf.weight[qFind]
	} else {
		uf.id[p] = qFind
		uf.weight[qFind] += uf.weight[pFind]
	}
	uf.count--
}

func (uf *UnionFind) Find(p int) int {
	for p != uf.id[p] {
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
	uf.weight = make([]int, n)
	return
}
