package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

const hoursPerDay = 24

type point struct {
	x, y int
}

func main() {
	fmt.Println("Welcome to Week 1 of Golang Training")
	fmt.Println("===============================================")
	fmt.Println("Enter the Number of Question you want to implement!")
	fmt.Println("You have the following Chocies")
	fmt.Println("2.1 2.2 2.5 3.1 3.2 3.3 3.4 3.5 3.6 4.3 5.0")
	var questionNo float64
	fmt.Scan(&questionNo)
	switch {
	case questionNo == 2.1:
		{

			var firVal int
			var secVal int
			fmt.Println("Enter Your First Number")
			fmt.Scan(&firVal)
			fmt.Println("Enter Your Second Number")
			fmt.Scan(&secVal)
			sum := firVal + secVal
			subtract := firVal - secVal
			multiply := firVal * secVal
			divide := firVal / secVal
			fmt.Println("The Addition of two numbers will be", sum)
			fmt.Println("The Subtraction of two numbers will be", subtract)
			fmt.Println("The Multiplication of two numbers will be", multiply)
			fmt.Println("The division of two numbers will be", divide)
		}
	case questionNo == 2.2:
		{

			p := point{1, 2}
			fmt.Printf("struct1: %v\n", p)

			fmt.Printf("struct2: %+v\n", p)

			fmt.Printf("struct3: %#v\n", p)

			fmt.Printf("type: %T\n", p)

			fmt.Printf("bool: %t\n", true)

			fmt.Printf("int: %d\n", 123)

			fmt.Printf("bin: %b\n", 14)

			fmt.Printf("char: %c\n", 33)

			fmt.Printf("hex: %x\n", 456)

			fmt.Printf("float1: %f\n", 78.9)

			fmt.Printf("float2: %e\n", 123400000.0)
			fmt.Printf("float3: %E\n", 123400000.0)

			fmt.Printf("str1: %s\n", "\"string\"")

			fmt.Printf("str2: %q\n", "\"string\"")

			fmt.Printf("str3: %x\n", "hex this")

			fmt.Printf("pointer: %p\n", &p)

			fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

			fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

			fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

			fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

			fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

			s := fmt.Sprintf("sprintf: a %s", "string")
			fmt.Println(s)
		}

	case questionNo == 2.5:
		{

			var userGuess int
			fmt.Println("Choose any number between 1 and 100")
			fmt.Scan(&userGuess)
			min, max := 1, 100
			rand.Seed(time.Now().UnixNano())
			secretNumber := rand.Intn(max-min) + min
			if userGuess == secretNumber {
				fmt.Println("You just guessed the same number as Computer")
			} else {
				fmt.Printf("Sorry Your Guess is Not same as Computer The Computer Guessed Number is %v", secretNumber)
			}

		}
	case questionNo == 3.1:
		{

			x := 5
			y := 8

			fmt.Println("x == y:", x == y)
			fmt.Println("x != y:", x != y)
			fmt.Println("x < y:", x < y)
			fmt.Println("x > y:", x > y)
			fmt.Println("x <= y:", x <= y)
			fmt.Println("x >= y:", x >= y)

		}
	case questionNo == 3.2:
		{

			str1 := "Geeks"
			str2 := "Geek"
			str3 := "GeeksforGeeks"
			str4 := "Geeks"

			result1 := str1 == str2
			result2 := str2 == str3
			result3 := str3 == str4
			result4 := str1 == str4

			fmt.Println("Result 1: ", result1)
			fmt.Println("Result 2: ", result2)
			fmt.Println("Result 3: ", result3)
			fmt.Println("Result 4: ", result4)

			result5 := str1 != str2
			result6 := str2 != str3
			result7 := str3 != str4
			result8 := str1 != str4

			fmt.Println("\nResult 5: ", result5)
			fmt.Println("Result 6: ", result6)
			fmt.Println("Result 7: ", result7)
			fmt.Println("Result 8: ", result8)

		}

	case questionNo == 3.3:
		{

			rand.Seed(time.Now().UnixNano())

			num := -5 + rand.Intn(10)

			if num > 0 {

				fmt.Println("The number is positive")
			} else if num == 0 {

				fmt.Println("The number is zero")
			} else {

				fmt.Println("The number is negative")
			}

		}

	case questionNo == 3.4:
		{

			var a bool = true
			var b bool = false
			if a && b {
				fmt.Printf("Line 1 - Condition is true\n")
			}
			if a || b {
				fmt.Printf("Line 2 - Condition is true\n")
			}

			a = false
			b = true
			if a && b {
				fmt.Printf("Line 3 - Condition is true\n")
			} else {
				fmt.Printf("Line 3 - Condition is not true\n")
			}
			if !(a && b) {
				fmt.Printf("Line 4 - Condition is true\n")
			}

		}

	case questionNo == 3.5:
		{
			var month string
			fmt.Scanln(&month)

			switch month {
			case "january", "december":
				fmt.Println("Winter.")
			case "february", "march":
				fmt.Println("Spring.")
			case "april", "may", "june":
				fmt.Println("Summer.")
			case "july", "august":
				fmt.Println("Monsoon.")
			case "september", "november":
				fmt.Println("Autumn.")
			}
		}

	case questionNo == 3.6:
		{

			var i, j int

			fmt.Println("Enter the range you want prime number!")
			var limit int
			fmt.Scan(&limit)
			for i = 2; i < limit; i++ {
				for j = 2; j <= (i / j); j++ {
					if i%j == 0 {
						break // if factor found, not prime
					}
				}
				if j > (i / j) {
					fmt.Printf("%d is prime\n", i)
				}
			}

		}
	case questionNo == 4.3:
		{
			year := 2018
			switch month := rand.Intn(12) + 1; month {
			case 2:
				day := rand.Intn(28) + 1
				fmt.Println(era, year, month, day)
			case 4, 6, 9, 11:
				day := rand.Intn(30) + 1
				fmt.Println(era, year, month, day)
			default:
				day := rand.Intn(31) + 1
				fmt.Println(era, year, month, day)
			}
		}
	case questionNo == 5.0:
		{

			fmt.Printf("%-15v %10v %6v %6v\n", "Spaceline", "Days", "Trip type", "Price")
			fmt.Println("=============================================")

			for count := 0; count < 10; count++ {

				// Random Ship Travel Speed
				rand.Seed(time.Now().UnixNano())
				minShipSpeed := 16
				maxShipSpeed := 30
				randomShipSpeed := (rand.Intn(maxShipSpeed-minShipSpeed+1) + minShipSpeed)
				shipSpeed := randomShipSpeed * 3600

				var distance = 62100000
				days := distance / shipSpeed / hoursPerDay

				// Selecting random trip type based on the days
				var tripType = ""

				if days > 40 {
					tripType = "Round-Trip"
				} else if days < 40 {
					tripType = "One-way"
				}

				// Generating random price

				minPrice := 36
				maxPrice := 50
				randomPrice := (rand.Intn(maxPrice-minPrice+1) + minPrice)

				var price int

				if shipSpeed > 100000 {
					price = randomPrice + 10
				} else if shipSpeed < 100000 {
					price = randomPrice
				}

				if tripType == "Round-Trip" {
					price = randomPrice * 2
				}
				// Selecting a random spaceline
				var spaceline = ""
				randomNumber := rand.Intn(3) + 1

				switch randomNumber {
				case 1:
					spaceline = "SpaceX"
				case 2:
					spaceline = "Virgin Galactic"
				case 3:
					spaceline = "Space Adventures"
				}

				fmt.Printf("%-18v %8v %10v $%4v %2v\n", spaceline, days, tripType, price, "mil")
			}
		}
	default:
		fmt.Println("Error ! Enter the question Number from above given Value.")
	}

}
