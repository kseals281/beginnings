package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var colors = []string{"blue", "red", "green", "yellow"}

/*
Define a function called Play that takes no parameters. The function should use the other functions you have defined to handle the overall logic of the whirlybird.
    1) The user should be prompted to enter a question for the whirlybird.
    2) If the user entered 'exit' (not case-sensitive), the program should exit.
    3) The user should be prompted to choose a color.
    4) If there are an even number of letters in the selected color, the user should be prompted to select an even number. If there are an odd number of letters in the color, they should be prompted to choose an odd number.
    5) If the number the user selected is prime, the user should be prompted to enter an even number. If the selected number is not prime, the user should be prompted to enter an odd number.
    6) This second number the user has selected should be used to determine the fortune the user should receive.
    7) The user's original question and then the fortune (on separate lines) should be printed out.
    8) Go back to step 1.
*/

func main() {
	var userNumber int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your question or exit:\n")
	text, _ := reader.ReadString('\n')

	if strings.ToLower(text) == "exit" {

	}

	color := GetColorSelection()

	if (len(color) % 2) == 0 {
		userNumber = GetNumberSelection(true)
	} else {
		userNumber = GetNumberSelection(false)
	}

	userNumber = GetNumberSelection(IsPrime(userNumber))

	fmt.Println(GetFortune(userNumber))
}

/*
GetFortune Function

Define a function called GetFortune that htakes one parameter named fortune_num.
The parameter should be an integer between 1 and 8 (inclusive). For each valid value,
the function should return a different fortune as a string. If the parameter's value
is anything other than an integer between 1 and 8, the function should return the following string:
  'Invalid forutune number.'
*/

func GetFortune(num int) string {
	switch num {
	case 1:
		return "Fortune one"

	case 2:
		return "Fortune two"

	case 3:
		return "Fortune three"

	case 4:
		return "Fortune four"

	case 5:
		return "Fortune five"

	case 6:
		return "Fortune six"

	case 7:
		return "Fortune seven"

	case 8:
		return "Fortune eight"
	}
	return "Invalid fortune number"
}

/*
GetColorSelection Function

Define a function called GetColorSelection that takes no parameters. The function
should prompt the user to enter one of the colors defined at the beginning of the file.
The user's input should not be case-sensitive. Only the four colors defined at the
beginning of the program should be considered valid input and only valid values
should be returned. If the user enters anything other than a valid color, you should
notify the user that the input was invalid and prompt them to enter their color again.
This process should continue until the user enters a valid color.

Note: see Python documentation to find a function that will convert strings
      to lower case.
*/

func GetColorSelection() string {
	var chosen string
	fmt.Println("\nPlease choose a color", colors)
	fmt.Scan(&chosen)
	for {
		chosen = strings.ToLower(chosen)
		for i := range colors {
			if colors[i] == chosen {
				return chosen
			}
		}
		fmt.Println("Your input was invalid. Please choose a color", colors)
		fmt.Scan(&chosen)
	}
}

/*
GetNumberSelection Function

Define a function called GetNumberSelection that takes one parameter named
use_even_numbers. The parameter will be a boolean value. The function should
prompt the user to enter an interger and return their selection as an
integer. If use_even_numbers is True, the user should be prompted to enter 2,
4, 6, or 8. If use_even_numbers is False, the user should be prompted to
input 1, 3, 5, or 7. No other input values should be allowed and only valid
values should be returned. If the user enters anything other than a valid
integer, you should notify the user that the input was invalid and prompt
them to enter their number again. This process should continue until the user
enters a valid number.*/

func GetNumberSelection(use_even_numbers bool) int {
	var number string
	if use_even_numbers == true {
		evenNumbers := []string{"2", "4", "6", "8"}
		fmt.Println("\nPlease pick a number", evenNumbers)
		fmt.Scan(&number)
		for {
			for i := range evenNumbers {
				if evenNumbers[i] == number {
					return (i * 2) + 2
				}
			}
			fmt.Println("Your input was invalid. Please pick a number", evenNumbers)
			fmt.Scan(&number)
		}
	} else {
		oddNumbers := []string{"1", "3", "5", "7"}
		fmt.Println("\nPlease pick a number", oddNumbers)
		fmt.Scan(&number)
		for {
			for i := range oddNumbers {
				if oddNumbers[i] == number {
					return (i * 2) + 1
				}
			}
			fmt.Println("Your input was invalid. Please pick a number", oddNumbers)
			fmt.Scan(&number)
		}
	}
}

/*
IsPrime Function

Define a function called IsPrime that takes one parameter named x. The parameter
will be an integer. The function should return a boolean value: True if the number
is prime, False otherwise. If x is not an integer, then the function should return None.

Notes:
    - 0 and 1 are *not* prime numbers.
    - Negative numbers are never prime.
    - Even though we only use 1-8 in this program, this function should work
      for any integer.
*/

func IsPrime(x int) bool {
	test := 1
	if test >= x {
		return false
	}
	for test <= x {
		test += 1
		result := x % test
		if result == 0 && test != 1.0 && test != x {
			return false
		}
		if test == x {
			return true
		}
	}
	return false
}
