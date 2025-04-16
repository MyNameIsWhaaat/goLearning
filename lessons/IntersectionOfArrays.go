package lessons

func Intersection(a, b []int) []int{
    res := make([]int, 0, 0)
	arrayElements := make(map[int]bool)
    
	for _, value := range a{
		arrayElements[value] = true
	}

	for _, value := range b{
		if arrayElements[value]{
			res = append(res, value)
			delete(arrayElements, value)
		}
	}

    return res
}
