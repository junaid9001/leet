/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumOperations = function(nums) {
    let sor=nums.toSorted((a,b)=>a-b)
    let flt=sor.filter(item=>item!==0)
    let operation=0
    while(flt.length>0){
        let trg=flt[0];
        let mod=flt.map(item=>item-trg).filter(item=>item!==0)
        flt=mod
        operation++
    }
    return operation    
};