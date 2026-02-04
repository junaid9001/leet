func checkAlmostEquivalent(word1 string, word2 string) bool {
	m := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		m[word1[i]]++
		m[word2[i]]--
	}

	for _, v := range m {
		if v > 3 || v< -3{
			return false
		}
	}
	return true

}