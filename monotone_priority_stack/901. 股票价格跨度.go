package main

import "fmt"

type StockSpanner struct {
	Prices []int
	Stack  [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Prices: make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	this.Prices = append(this.Prices, price)

	var (
		cnt = 1
	)

	fmt.Println("当前数据: ", price, "移除栈顶数据: ", this.Stack)

	// 每次都重新计算会超时

	// 0 -> stackIdx
	// 1 -> 历史cnt

	for len(this.Stack) > 0 && price >= this.Prices[this.Stack[len(this.Stack)-1][0]] {
		// 每移除一次，计算上移除的那个所带的cnt
		cnt += this.Stack[len(this.Stack)-1][1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}

	this.Stack = append(this.Stack, [2]int{
		len(this.Prices) - 1, cnt,
	})

	return cnt
}

func main() {
	var stockSpanner StockSpanner = Constructor()
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(80))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(70))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(75))
	fmt.Println(stockSpanner.Next(85))
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
