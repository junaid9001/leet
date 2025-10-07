/**
 * @param {number} n
 * @return {string[]}
 */
var fizzBuzz = function(n) {
    let answer=[]
    for(let i=1;i<=n;i++){
        answer.push(i.toString())
    }
    for(let i=0;i<answer.length;i++){
        let ton=Number(answer[i])
        if(ton%3===0&ton%5===0){
            answer[i]="FizzBuzz"
        }else if(ton%3===0){
            answer[i]="Fizz"
        }else if(ton%5===0){
            answer[i]="Buzz"
        }else{
            answer[i]=answer[i]
        }
    }
    return answer
};