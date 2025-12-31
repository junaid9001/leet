func transformArray(nums []int) []int {
    for i:=0;i<len(nums);i++{
        nums[i]=nums[i]%2
    }
    slices.Sort(nums)
    return nums
}