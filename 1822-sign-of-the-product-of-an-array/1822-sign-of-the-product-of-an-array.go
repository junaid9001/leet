func arraySign(nums []int) int {
	sign := 1
	for _, val := range nums {
        if val==0{
            return 0
        }
        if val<0{
            sign =-sign
        }
	}
    return sign
}

