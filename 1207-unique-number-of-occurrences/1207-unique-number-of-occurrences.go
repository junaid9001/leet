func uniqueOccurrences(arr []int) bool {
    m:=make(map[int]int)
    count:=make(map[int]bool)
    for _,val:=range arr{
        m[val]++
    }
    fmt.Println(m)

    for _,v:=range m{
        _,ok:=count[v]
        if !ok{
            count[v]=true
            continue
        }
        return false
    }
    return true
}