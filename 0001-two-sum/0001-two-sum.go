func twoSum(nums []int, target int) []int {
	m := map[int]int{}


	for key, val := range nums {
		comp := target - val
		_, ok := m[comp]
		if ok {
			return []int{key, m[comp]}
		} else {
			m[val] = key
		}
	}
	return nil
}