func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[string]*struct {
		start int
		end   int
	})
    ans:=-1

	for key, val := range s {
		_, ok := m[string(val)]
		if !ok {
			m[string(val)] = &struct {
				start int
				end   int
			}{start: key}
		} else {
			exis:=m[string(val)]
            exis.end=key
		}
	}

    for _,val:=range m{
        if val.end-val.start-1>ans{
            ans=val.end-val.start-1
        }
    }
    return ans
}