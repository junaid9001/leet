/**
 * @param {number} num
 * @return {number}
 */
var findComplement = function(num) {
    let bnr=num.toString(2).split("")
    let newer=bnr.map(item=>item==0?1:0)
    let str=newer.join('')
    let output=parseInt(str,2)
    return output
};