package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
	"unicode/utf8"
)

type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsiusTwo(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvinThree(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {

	fmt.Println("Welcome to Week 2 of Golang Training")
	fmt.Println("===============================================")
	fmt.Println("Enter the Number of Question you want to implement!")
	fmt.Println("You have the following Chocies")
	fmt.Println("6.2 6.3 6.4 7.1 7.2 7.3 8.1 8.2 8.3")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("9.1 9.2 9.3 9.4 9.5 10.1 10.2 10.3 10.4 10.5")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("11.0 12.2 13.1 13.2 13.3")
	var questionNo float64
	fmt.Scan(&questionNo)

	switch {
	case questionNo == 6.2:
		{
			third := 1.0 / 3
			fmt.Println(third)
			fmt.Printf("%v\n", third)
			fmt.Printf("%f\n", third)
			fmt.Printf("%.3f\n", third)
			fmt.Printf("%6.2f\n", third)
			fmt.Printf("%06.2f\n", third)
		}
	case questionNo == 6.3:
		{
			celsius := 21.0
			fmt.Print((celsius/5.0*9.0)+32, "º F\n")
			fahrenheit := (celsius * 9.0 / 5.0) + 32.0
			fmt.Print(fahrenheit, "º F")
		}
	case questionNo == 6.4:
		{
			piggyBank := 0.1
			piggyBank += 0.2
			fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
		}
	case questionNo == 7.1:
		{
			var value int = -12
			fmt.Println(value)
		}
	case questionNo == 7.2:
		{
			a := "text"
			fmt.Printf("Type %T for %[1]v\n", a)
			var red, green, blue uint8 = 0x00, 0x8d, 0xd5
			fmt.Printf("%x %x %x", red, green, blue)
		}
	case questionNo == 7.3:
		{
			var red uint8 = 255
			red++
			fmt.Println(red)
			var number int8 = 127
			number++
			fmt.Println(number)
			fmt.Println("Second Part")
			future := time.Unix(12622780800, 0)
			fmt.Println(future)
		}
	case questionNo == 8.1:
		{
			const lightSpeed = 299792 // km/s
			const secondsPerDay = 86400
			var distance int64 = 41.3e12
			fmt.Println("Alpha Centauri is", distance, "km away.")
			days := distance / lightSpeed / secondsPerDay
			fmt.Println("That is", days, "days of travel at light speed.")
		}
	case questionNo == 8.2:
		{
			lightSpeed := big.NewInt(299792)
			secondsPerDay := big.NewInt(86400)
			distance := new(big.Int)
			distance.SetString("24000000000000000000", 10)
			fmt.Println("Andromeda Galaxy is", distance, "km away.")
			seconds := new(big.Int)
			seconds.Div(distance, lightSpeed)
			days := new(big.Int)
			days.Div(seconds, secondsPerDay)
			fmt.Println("That is", days, "days of travel at light speed.")
		}
	case questionNo == 8.3:
		{
			const distance = 24000000000000000000
			const lightSpeed = 299792
			const secondsPerDay = 86400
			const days = distance / lightSpeed / secondsPerDay
			fmt.Println("Andromeda Galaxy is", days, "light days away.")
		}
	case questionNo == 9.1:
		{
			peace := "peace be upon you\nupon you be peace"
			fmt.Println(peace)
			fmt.Println(`strings can span multiple lines with the \n escape sequence`)
		}
	case questionNo == 9.2:
		{
			var pi rune = 960
			var alpha rune = 940
			var omega rune = 969
			var bang byte = 33
			fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
		}
	case questionNo == 9.3:
		{
			message := "shalom"
			c := message[5]
			fmt.Printf("%c\n", c)
		}
	case questionNo == 9.4:
		{
			message := "shalom"
			for i := 0; i < 6; i++ {
				c := message[i]
				fmt.Printf("%c\n", c)
			}
		}
	case questionNo == 9.5:
		{
			question := "¿Cómo estás?"
			fmt.Println(len(question), "bytes")
			fmt.Println(utf8.RuneCountInString(question), "runes")
			c, size := utf8.DecodeRuneInString(question)
			fmt.Printf("First rune: %c %v bytes", c, size)
		}
	case questionNo == 10.1:
		{
			age := 41
			marsDays := 687
			earthDays := 365.2425
			fmt.Println("I am", age*int(earthDays)/marsDays, "years old on Mars.")
		}
	case questionNo == 10.2:
		{
			age := 41
			marsAge := float64(age)
			marsDays := 687.0
			earthDays := 365.2425
			marsAge = marsAge * earthDays / marsDays
			fmt.Println("I am", marsAge, "years old on Mars.")
		}
	case questionNo == 10.3:
		{
			var bh float64 = 32767
			var h = int16(bh)
			fmt.Println(h)
		}
	case questionNo == 10.4:
		{
			var pi rune = 960
			var alpha rune = 940
			var omega rune = 969
			var bang byte = 33
			fmt.Print(string(pi), string(alpha), string(omega), string(bang))
		}
	case questionNo == 10.5:
		{
			launch := false
			launchText := fmt.Sprintf("%v", launch)
			fmt.Println("Ready for launch:", launchText)
			var yesNo string
			if launch {
				yesNo = "yes"
			} else {
				yesNo = "no"
			}
			fmt.Println("Ready for launch:", yesNo)
		}
	case questionNo == 11.0:
		{
			fmt.Println("Part 1 of  The Vigen?re Cipher Program")
			fmt.Println("======================================")
			cipherText := "CSOITEUIWUIZNSROCNKFD"
			keyword := "GOLANG"
			message := ""
			keyIndex := 0
			for i := 0; i < len(cipherText); i++ {
				// A=0, B=1, ... Z=25
				c := cipherText[i] - 'A'
				k := keyword[keyIndex] - 'A'
				// cipher letter - key letter
				c = (c-k+26)%26 + 'A'
				message += string(c)
				// increment keyIndex
				keyIndex++
				keyIndex %= len(keyword)
			}
			fmt.Println(message)
			fmt.Println("======================================")
			fmt.Println("Part 2 of  The Vigen?re Cipher Program")
			fmt.Println("======================================")
			messageTwo := "your message goes here"
			keywordTwo := "golang"
			keyIndexTwo := 0
			cipherTextTwo := ""
			messageTwo = strings.ToUpper(strings.Replace(messageTwo, " ", "", -1))
			keywordTwo = strings.ToUpper(strings.Replace(keywordTwo, " ", "", -1))
			for i := 0; i < len(messageTwo); i++ {
				c := messageTwo[i]
				if c >= 'A' && c <= 'Z' {
					// A=0, B=1, ... Z=25
					c -= 'A'
					k := keywordTwo[keyIndexTwo] - 'A'
					// cipher letter + key letter
					c = (c+k)%26 + 'A'
					// increment keyIndex
					keyIndexTwo++
					keyIndexTwo %= len(keywordTwo)
				}
				cipherTextTwo += string(c)
			}
			fmt.Println(cipherTextTwo)
		}
	case questionNo == 12.2:
		{
			kelvin := 294.0
			celsius := kelvinToCelsius(kelvin)
			fmt.Print(kelvin, "º K is ", celsius, "º C")
		}
	case questionNo == 13.1:
		{
			type celsius float64
			var temperature celsius = 20
			fmt.Println(temperature)
		}
	case questionNo == 13.2:
		{
			var k kelvin = 294.0
			c := kelvinToCelsiusTwo(k)
			fmt.Print(k, "º K is ", c, "º C")
		}
	case questionNo == 13.3:
		{
			var c celsius = 127.0
			k := celsiusToKelvinThree(c)
			fmt.Print(c, "º C is ", k, "º K")
		}
	default:
		{
			fmt.Println("Error ! Enter the question Number from above given Value.")
		}
	}
}
