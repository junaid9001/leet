/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {
    let emp=[]
    for(let i=0;i<nums.length;i++){
        let count=0
        for(let j=0;j<nums.length;j++){
            
            if(nums[i]>nums[j]){
                count++
            }
        }
        emp.push(count)
    }
    return emp
};