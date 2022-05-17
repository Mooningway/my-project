package u_math

import (
	"fmt"
	"strconv"
)

func Multiply4F(num1, num2 float64) float64 {
	temp := fmt.Sprintf(`%.4f`, num1*num2)
	result, _ := strconv.ParseFloat(temp, 64)
	return result
}

func Divide4F(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	temp := fmt.Sprintf(`%.4f`, num1/num2)
	result, _ := strconv.ParseFloat(temp, 64)
	return result
}
