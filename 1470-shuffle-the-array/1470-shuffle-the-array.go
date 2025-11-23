func shuffle(nums []int, n int) []int {
    ans:=make([]int,len(nums))

    for i:=0;i<n;i++{
        ans[2*i]=nums[i]
        ans[2*i+1]=nums[i+n]
    }
    return ans
}