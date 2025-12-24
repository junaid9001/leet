func largestOddNumber(num string) string {
    for i:=len(num)-1;i>=0;i--{
        dig:=num[i] - '0'

        if dig%2!=0{
            return num[:i+1]
        }
    }
    return ""
}