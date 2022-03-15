package practice_1_5_1

type UnionFind struct {
	id    []int
	count int

	queryCount int
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) Union(p, q int) {
	uf.queryCount++
	pid := uf.id[p]
	uf.queryCount++
	qid := uf.id[q]
	if qid == pid {
		return
	}
	for i := range uf.id {
		uf.queryCount++
		if uf.id[i] == qid {
			uf.queryCount++
			uf.id[i] = pid
		}
	}
	uf.count--
}

func (uf *UnionFind) Find(p int) int {
	uf.queryCount++
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
