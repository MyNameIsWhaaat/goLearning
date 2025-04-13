package lessons
//Написать семафор
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

func Sem(wg *sync.WaitGroup, s *Semaphore){
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int){
			s.Acquire()

			time.Sleep(time.Duration(rand.Intn(3)) + time.Second)

			s.Release()

			fmt.Printf("Горутина %d что-то там сделала\n", id)

			wg.Done()

		}(i)
	}

	wg.Wait()
	fmt.Println("Вроде все все сделали")
}



