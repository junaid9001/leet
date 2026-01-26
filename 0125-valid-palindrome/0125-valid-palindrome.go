func isPalindrome(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		if !unicode.IsLetter(r[i])&&!unicode.IsDigit(r[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(r[j]) &&!unicode.IsDigit(r[j]){
			j--
			continue
		}
		if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
			return false
		}else{
            i++
            j--
        }
	}
    return true
}