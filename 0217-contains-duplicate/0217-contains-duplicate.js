/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let map=new Set(nums)
    if (map.size<nums.length){
        return true
    }else{return false}

};