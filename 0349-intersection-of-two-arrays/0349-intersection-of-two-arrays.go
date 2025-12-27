func intersection(nums1 []int, nums2 []int) []int {
    var ans []int
    set:=make(map[int]int)
    for _,n:=range nums1{
        set[n]=1
    }

    for _,val:=range nums2{
        if set[val]==1{
            ans=append(ans,val)
            set[val]=2
        }
    }
    return ans
}