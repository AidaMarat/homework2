package main

import "fmt"

func main() {
	dq := dequeue{}

	for i := 0; i < 10; i++ {
		dq.push_back(i)
	}
	fmt.Println("Test main: ")
	fmt.Println(dq)
	fmt.Println(dq.pop_back())
	fmt.Println(dq.pop_up())
	dq.push_up(22)
	fmt.Println(dq)
	fmt.Println(dq.peek(false))
	fmt.Println(dq.peek(true))
}

type dequeue struct {
	values []int
}

func (dq *dequeue) push_back(value int) {
	dq.values = append(dq.values, value)
}

func (dq *dequeue) pop_back() int {
	node := dq.values[len(dq.values)-1]
	dq.values = dq.values[:len(dq.values)-1]
	return node
}

func (dq *dequeue) push_up(value int) {
	var n [100]int
	for i := 0; i < len(dq.values); i++ {
		n[i+1] = dq.values[i]
	}
	dq.values = dq.values[:len(dq.values)+1]
	n[0] = value
	for i := 0; i < len(dq.values); i++ {
		dq.values[i] = n[i]
	}
}

func (dq *dequeue) pop_up() int {
	node := dq.values[0]
	var n [100]int
	for i := 1; i < len(dq.values); i++ {
		n[i-1] = dq.values[i]
	}
	for i := 0; i < len(dq.values)-1; i++ {
		dq.values[i] = n[i]
	}
	dq.values = dq.values[:len(dq.values)-1]
	return node
}

func (dq *dequeue) peek(value bool) int {
	var node int
	if value == true {
		node = dq.values[len(dq.values)-1]
	} else {
		node = dq.values[0]
	}
	return node
}
