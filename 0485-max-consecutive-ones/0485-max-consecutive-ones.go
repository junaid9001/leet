func findMaxConsecutiveOnes(nums []int) int {
  var ans int 
  cur:=0
  for i:=0;i<len(nums);i++{
    if nums[i]==0{
        cur=0
        continue
    }
    if nums[i]==1{
        cur++
        if cur>ans{
            ans=cur
        }
    }
  }
  if cur>ans{
    ans=cur
  }
  return ans
}