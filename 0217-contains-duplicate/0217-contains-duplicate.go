func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if !m[v] {
			m[v]=true
            continue
		}
        return true

	}
	return false
}