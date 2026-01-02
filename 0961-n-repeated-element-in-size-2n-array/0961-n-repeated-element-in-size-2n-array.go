func repeatedNTimes(nums []int) int {
    n:=len(nums)/2
    m:=make(map[int]int)
    for _,v:=range nums{
        m[v]++
        if m[v]==n{
            return v
        }
    }
    return 0
}