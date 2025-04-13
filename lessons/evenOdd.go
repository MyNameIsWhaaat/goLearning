package lessons

func EvenOdd(input [] int) (even []int, odd []int){
	for _, v := range input{
		if (v % 2 == 0){
			even = append(even, v)
		} else{
			odd = append(odd, v)
		}
	}

	return even, odd
}