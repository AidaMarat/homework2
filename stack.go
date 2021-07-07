package main

import "fmt"

func main() {
	st := stack{}

	for i := 0; i < 10; i++ {
		st.push(i)
	}
	fmt.Println(st)
	fmt.Println(st.pop())
	fmt.Println(st.peek())
	fmt.Println(st)

}

type stack struct {
	values []int
}

//it is method, not function
func (st *stack) push(value int) {
	st.values = append(st.values, value)
}

func (st *stack) pop() int {
	node := st.values[len(st.values)-1]
	st.values = st.values[:len(st.values)-1]
	return node
}

func (st *stack) peek() int {
	node := st.values[len(st.values)-1]
	return node
}

//possible:
/*func (st *stack) push() (int, int){

}*/
