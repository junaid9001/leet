func containsNearbyDuplicate(nums []int, k int) bool {
    m:=make(map[int]int)

    for i,val:=range nums{
        x,ok:=m[val]
        if ok{
            r:=i-x
            if r<=k{
                return true
            }
        }

        m[val]=i
    }
    return false
}