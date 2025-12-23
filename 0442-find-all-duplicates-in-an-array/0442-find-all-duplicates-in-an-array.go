func findDuplicates(nums []int) []int {
    var ans []int
    set:=make(map[int]int)
    for _,val:=range nums{
        set[val]++
        if set[val]>1{
            ans=append(ans,val)
        }
    }
    return ans

}