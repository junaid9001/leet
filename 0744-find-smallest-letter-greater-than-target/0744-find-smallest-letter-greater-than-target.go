func nextGreatestLetter(letters []byte, target byte) byte {
    for key,v:=range letters{
        if v>target{
            return v
        }
        if key==len(letters)-1{
            return letters[0]
        }
    }
    return byte(0)
}