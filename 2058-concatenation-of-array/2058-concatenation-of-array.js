/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getConcatenation = function(nums) {
    let newarr=[]
    for(let i=0;i<nums.length;i++){
        newarr.push(nums[i])
    }
        for(let i=0;i<nums.length;i++){
        newarr.push(nums[i])
    }
    return newarr
};