/**
 * @param {number} num
 * @return {boolean}
 */
var isSameAfterReversals = function(num) {
    let numstr=num.toString()
    let numrev1=[...numstr].reverse()
    let fresh=[...numrev1]

    for(let i=0;i<numrev1.length;i++){
        if(numrev1[i]!=="0"){
            break;
        }
        fresh.splice(0,1)
    }

    numrev2=Number(fresh.reverse().join(""))

    if(numrev2===num){
        return true
    }
    return false

    

};