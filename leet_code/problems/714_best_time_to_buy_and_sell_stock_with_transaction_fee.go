package problems

import "fmt"

func maxProfitWithFee(prices []int, fee int) int {
	type User struct {
		id int
	}
	p := make(map[int]User)
	p[1] = User{1}
	p[2] = User{2}
	p[3] = User{3}
	for _, pp := range p {
		pp.id = 10
	}
	fmt.Println(p)
	return 0
}
