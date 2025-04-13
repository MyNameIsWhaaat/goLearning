package main

import (
	// "sync"
	// "exampleGo/Semaphore/lessons"
	"fmt"
)

func intersection(a, b []int) []int{
    c := make([]int, 0)   
    lenA := len(a)
    lenB := len(b)

    for i:=0; i<=lenA; i++{
        for j:=0; j<=lenB; j++{
            if a[i] == b[j]{
                c = append(c, a[i])
                break
            } 
        }       
    }

    return c
}




func main(){

	// semaphore:= lessons.Semaphore{
	// 	C: make(chan struct{}, 5),
	// }

	// var wg sync.WaitGroup

	// lessons.Sem(&wg, &semaphore)

	// res := lessons.Calc()
	// fmt.Println(res)
	
	// mas := []int{2, 7, 9 ,10, 45, 6}

	// even, odd := lessons.EvenOdd(mas)

	// fmt.Print(even, odd)

	    //         0   1  2  3
		a := []int{23, 3, 2, 1}

		//         0 1  2  3
		b := []int{6, 2, 4, 23}
		// [2, 23]
	
		fmt.Printf("%v\n", intersection(a, b))
}