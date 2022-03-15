package base

type UnionFind interface {
	Count() int
	Connected(p, q int) bool
	Find(p int) int
	Union(p, q int)
}
