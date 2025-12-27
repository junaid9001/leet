func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	var ans []int
	for _, val := range nums {
		count[val]++
	}

	for i := 0; i < k; i++ {
		var temp int
		for key, val := range count {
			if val > count[temp] {
				temp = key
			}
		}
        ans=append(ans,temp)
        delete(count,temp)

	}
    return ans
}