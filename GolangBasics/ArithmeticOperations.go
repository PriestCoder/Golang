package GolangBasics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3

	var result int = a + b

	fmt.Println("Addition :", result)

	result = a - b

	fmt.Println("Subtraction :", result)

	result = a * b

	fmt.Println("Multiplication :", result)

	result = a % b

	fmt.Println("Modulus :", result)

	result = a / b

	fmt.Println("Divison :", result)

	const p float64 = 22.0 / 7.0

	fmt.Println(p)

	//Overflow with signed integers
	var MaxInt int64 = 9223372036854775807 //max value that int64 can hold
    fmt.Println(MaxInt)
	MaxInt = MaxInt + 1

	fmt.Println(MaxInt) 

	//Overflow with unsigned integers 
	var uMaxInt uint64 = 18446744073709551615 // max value for uint 64 type
	fmt.Println(uMaxInt) 

	uMaxInt = uMaxInt+1 
    fmt.Println(uMaxInt)

	//Underflow with floating point numbers 
	var SmallFloat float64 = 1.0e-323 
	fmt.Println(SmallFloat)


	SmallFloat = SmallFloat/math.MaxFloat64 
	fmt.Println(SmallFloat)

}
