func uniqueOccurrences(arr []int) bool {
    m:=make(map[int]int)
    for _,val:=range arr{
        m[val]++
    }


   seen:=make(map[int]bool)

    for _,v:=range m{
        if !seen[v]{
            seen[v]=true
            continue
        }
        return false 
    }
    return true
}