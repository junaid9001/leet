func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	var ans []int

	for key, val := range nums {
		comp := target - val
		_, ok := m[comp]
		if ok {
			ans = append(ans, []int{key, m[comp]}...)
			return ans
		} else {
			m[val] = key
		}
	}
	return nil
}