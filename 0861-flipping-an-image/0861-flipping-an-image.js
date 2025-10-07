/**
 * @param {number[][]} image
 * @return {number[][]}
 */
var flipAndInvertImage = function(image) {
    let output=[]
    for(let i=0;i<image.length;i++){
        let pi=image[i].reverse()
        let bye=pi.map(item=>item===0?1:0)
        output.push(bye)
    }
    return output
};