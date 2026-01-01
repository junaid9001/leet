func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var a int
    for _,v:=range hours{
        if v>=target{
            a++
        }
    }
    return a
}