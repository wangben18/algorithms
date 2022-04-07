package quick_union_weighted_compressed

type UnionFind struct {
	id    []int
	count int
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
	uf.id[p] = qFind
	uf.count--
}

func (uf *UnionFind) Find(p int) int {
	idNeedCompress := make([]int, 0)
	for p != uf.id[p] {
		p = uf.id[p]
		idNeedCompress = append(idNeedCompress, p)
	}
	for _, id := range idNeedCompress {
		uf.id[id] = p
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
