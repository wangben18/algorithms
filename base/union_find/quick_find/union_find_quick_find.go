package quick_find

type UnionFind struct {
	id    []int
	count int
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) Union(p, q int) {
	pid := uf.id[p]
	qid := uf.id[q]
	if qid == pid {
		return
	}
	for i := range uf.id {
		if uf.id[i] == qid {
			uf.id[i] = pid
		}
	}
	uf.count--
}

func (uf *UnionFind) Find(p int) int {
	return uf.id[p]
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
