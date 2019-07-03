package main

import "fmt"

func main() {
	var i int
	/*fmt.Println("Enter number of strings")
	fmt.Scan(n)*/
	var a []string
	a[0] = "aaa"
	a[1] = "lll"
	a[2] = "uuu"
	a[3] = "bbb"
	a[4] = "hhh"
	/*for i:=0; i<n; i++{
		fmt.Println("Enter strings: ")
		fmt.Scan(a[i])
	}*/
	MergeSort(a, 0, 4)
	for i = 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func Merge(a []string, low int, high int, mid int) {
	var i, j, k int;
	i = low
	k = 0
	j = mid + 1
	var t []string

	for (i <= mid && j <= high) {
		if (a[i] < a[j]) {
			t[k] = a[i]
			k++
			i++
		} else {
			t[k] = a[j]
			k++
			j++
		}
	}
	for i <= mid {
		t[k] = a[i]
		k++
		i++
	}
	for (j <= high) {
		t[k] = a[j]
		k++
		j++
	}
	for i = low; i <= high; i++ {
		a[i] = t[i-low]
	}

}

func MergeSort(a []string, low int, high int) {
	var mid int
	if low < high {
		mid = (low + high) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid+1, high)
		Merge(a, low, high, mid)
	}
}
