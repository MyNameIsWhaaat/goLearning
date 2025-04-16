package main

import (
	// "sync"
	"exampleGo/Semaphore/lessons"
	"fmt"
)

func main() {

	// semaphore:= lessons.Semaphore{
	// 	C: make(chan struct{}, 5),
	// }

	// var wg sync.WaitGroup

	// lessons.Sem(&wg, &semaphore)

	// res := lessons.Calc()
	// fmt.Println(res)

	// mas := []int{2, 7, 9 ,10, 45, 6}

	//even, odd := lessons.EvenOdd(mas)

	// fmt.Print(even, odd)
	a := []int{23, 3, 2, 1, 6, 7, 9}
	b := []int{6, 2, 4, 23, 23, 9}

	fmt.Printf("%v\n", lessons.Intersection(a, b))

	a1 := []int{1, 2, 0, 3, 2, 1, 5, 1, 9, 6}
	fmt.Printf("%v\n", lessons.DuplNumArray(a1))
}
