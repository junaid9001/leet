func convertTemperature(celsius float64) []float64 {
    Kelvin:= celsius + 273.15
    Fahrenheit:= celsius * 1.80 + 32.00
    ans:=[]float64{Kelvin,Fahrenheit}
    return ans
    }