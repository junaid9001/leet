func findMissingElements(nums []int) []int {
    var res []int
 slices.Sort(nums)

    for i:=nums[0];i<=nums[len(nums)-1];i++{
        if !slices.Contains(nums,i){
            res=append(res,i)
        }
    }
    return res
}