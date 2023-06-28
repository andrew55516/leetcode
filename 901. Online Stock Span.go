package main

type StockSpanner struct {
	stack [][2]int
	count int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
		count: 0,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.count++
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, this.count})

	if len(this.stack) == 1 {
		return this.count
	}

	return this.stack[len(this.stack)-1][1] - this.stack[len(this.stack)-2][1]
}
