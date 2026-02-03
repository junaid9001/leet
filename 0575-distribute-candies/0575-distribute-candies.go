func distributeCandies(candyType []int) int {
    m:=make(map[int]bool)
    var count int
    for _,v:= range candyType{
        if !m[v]{
            m[v]=true
            count++
            if count==len(candyType)/2{
                return count
            }
        }
    }
    return count
}