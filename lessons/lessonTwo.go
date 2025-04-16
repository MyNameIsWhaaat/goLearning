 package lessons

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for value := range ch {
// 		fmt.Printf("go %d: %d\n", id, value)
// 	}
// }


// func main() {
	
// 	valuesChan := make(chan int, 30)

// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		valuesChan <- i
// 	}

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go worker(i, valuesChan, &wg)
// 	}

// 	close(valuesChan)

// 	wg.Wait()
// }
