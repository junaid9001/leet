func checkIfPangram(sentence string) bool {
    b:=[]byte(sentence)
    for i:=97;i<=122;i++{
        if !slices.Contains(b,byte(i)){
            return false
        }
    }
    return true
}