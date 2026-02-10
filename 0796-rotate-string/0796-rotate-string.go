func rotateString(s string, goal string) bool {
    if len(s)!=len(goal){
        return false
    }
    if s==goal{
        return true
    }
	b := []byte(s)

	for i := 0; i < len(s); i++ {
		first:=b[0]
		b = append(b[1:],byte(first))
		if string(b) == goal {
			return true
		}
	}
	return false
}