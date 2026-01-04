func sumFourDivisors(nums []int) int {
	var a int
	for _, val := range nums {
		temp := []int{}
		for i := 1; i*i <= val; i++ {
			if val%i == 0 {
				temp = append(temp, i)
                
				if i*i != val {
					temp = append(temp, val/i)
				}
			}
			if len(temp) > 4 {
				break
			}
		}
		if len(temp) == 4 {
			for _, v := range temp {
				a += v
			}
		}
	}
	return a
}