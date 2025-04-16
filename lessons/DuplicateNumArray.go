package lessons

func DuplNumArray(a []int)[]int{
	count := make(map[int]int)
    res := make([]int, 0)

    for _, v := range a {
        count[v]++
    }

    for num, c := range count {
        if c > 1 {
            res = append(res, num)
        }
    }

    return res
}

//{1, 2, 0, 3, 2, 1, 5, 1, 9, 6}