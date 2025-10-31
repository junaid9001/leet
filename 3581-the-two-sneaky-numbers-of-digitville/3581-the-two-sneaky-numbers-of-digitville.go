func getSneakyNumbers(nums []int) []int {
    var output []int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]==nums[j]{
                output=append(output,nums[i])
                break;
            }
        }
    }
    return output
}