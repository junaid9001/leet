/**
 * @param {number} num
 * @return {number}
 */
var maximum69Number  = function(num) {
    let arr=num.toString().split("")
    let arrnum=arr.map(item=>Number(item))
    let max=num
    for(let i=0;i<arr.length;i++){
        let copy=[...arrnum]
        if(copy[i]===9){
            copy[i]=6
        }else{
            copy[i]=9
        }
       let numm=Number(copy.join(""))
       if(numm>max){
        max=numm
       }

    }
  return max
};