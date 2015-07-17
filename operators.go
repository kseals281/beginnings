package main

import("fmt"; "math")

func main() {
    a, b, c, d, e, f, h := 3, 4, 5, 2.155, 8, 9876543, 150
    x, y, z := "amazing", "this", "is"

    fmt.Println(problemOne(float64(a)))
    fmt.Println(problemTwo(a, b))
    fmt.Println(problemThree(x, y, z))
    fmt.Println(problemFour(a, b))
    fmt.Println(problemFive(b))
    fmt.Println(problemSix(a, b))
    fmt.Println(problemSeven(a, b, c))

    x, y = problemEight(x, y)
    fmt.Println(x, y)

    fmt.Println(problemNine(c))
    fmt.Println(problemTen(d))
    fmt.Println(problemEleven(a, b, c, d))
    fmt.Println(problemTwelve(a, b, c, e))
    fmt.Println(bonusA(f, 3))
    fmt.Println(bonusB(h))
}

func problemOne(a float64) float64 {
    //Problem number one
    //Cube 'a' and return the result
    fmt.Println("Problem One")
    return math.Pow(a, 3)
}

func problemTwo(a, b int) (bool, string) {
    //Task: Implement xor (abbreviation for 'exclusive or'). Xor should be true if a
    //and b are not the same. Return the value of a xor b.
    fmt.Println("\nProblem Two")
    if a == b {
	return false, "Equal"
    } else {
	return true, "Not equal"
    }
}

func problemThree(a, b, c string) string {
    //Task: Combine a, b, and c to form a sentence (string) and return the result.
    fmt.Println("\nProblem Three")
    return (b + " " + a + " " + c)
}

func problemFour(a, b int) int {
    //Task: Calculate the remainder of the division of the numerator by the
    //denominator and return the result.
    fmt.Println("\nProblem Four")
    return (a % b)
}

func problemFive(b int) float64 {
    //Task: Calculate the square root of 'x' and return the result.
    fmt.Println("\nProblem Five")
    return math.Sqrt(float64(b))
}

func problemSix(exp, base int) float64 {
    //Task: Use one of the shift operators (left or right) to multiply base by 2^exp
    //and return the result.
    fmt.Println("\nProblem Six")
    return float64(base) * math.Exp2(float64(exp))
}

func problemSeven(a, b, c int) bool {
    //Task: Check to see if a is between b and c. Return True if a is between b and
    //c, but False otherwise.
    fmt.Println("\nProblem Seven")
    return (c < a && a < b) || (b < a && a < c)
}

func problemEight(value1, value2 string) (string, string) {
    //Task: Swap value1 and value2. Given value1 and value2 above make sure that you
    //swap value1 to be value2, and value2 to be value1. Return value1 + value2
    fmt.Println("\nProblem Eight")
    //var value1, value2 string = value2, value1
    return value2, value1
}

func problemNine(num int) int {
    //Task: Return num rounded to the nearest multiple of 10 (5 rounds up).
    fmt.Println("\nProblem Nine")
    num += 5
    num -= num % 10
    return num
}

func problemTen(num float64) float64 {
    //Task: Return the thousandth digit in num. If there is no thousandth digit,
    //return 0.
    fmt.Println("\nProblem Ten", num)
    return math.Mod(num, 0.01)
}

func problemEleven(top1, bottom1, top2 int, bottom2 float64) bool {
    //Task:  Given two ranges a1-a2 and b1-b2, determine if the two ranges overlap.
    //Store a boolean representing the outcome in a variable named output.
    fmt.Println("\nProblem Eleven")
    if (float64(top1) > bottom2) && (top1 < top2) || (top2 > bottom1) && (top2 < top1) {
	return true
    }
    return false
}

func problemTwelve(x1, y1, x2, y2 int) float64 {
    //Task: Calculate the distance between two points (x1, y1) and (x2, y2) and
    //store it in output. If you don't remember your high school math, feel
    //free to Google the distance formula.
    fmt.Println("\nProblem Twelve")
    return math.Sqrt(math.Pow(float64(x2 - x1), 2) + math.Pow(float64(y2 - y1), 2))
}

func bonusA(num, n int) int {
    //Task: Return the nth digit in num. (The rightmost digit is at n = 1)
    fmt.Println("\nBonus A")
    return int((num / int(math.Pow(float64(10), float64(n - 1)))) % 10)
}

func bonusB(min int) (days, hours, minutes int) {
    //Task: Convert the given time (which represents total minutes) into
    //days/hours/minutes. Store your results in variables named days, hours,
    //and minutes and return a tuple containing those values.
    fmt.Println("\nBonus B")
    days = min / 1440
    hours = (min % 1440) / 60
    minutes = (min % 1440) % 60
    return days, hours, minutes
}
