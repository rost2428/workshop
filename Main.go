package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arr []string

func loadArr(file string) error {
	arr = nil
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		arr = append(arr, scan.Text())
	}
	return scan.Err()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_0
//Print a chessboard
func chessboard(h, w int) {
	height := h
	width := w

	for i := 0; i < height; i++ {
		for j := 0; j < width*2; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					fmt.Print("^")
				} else {
					fmt.Print(" ")
				}
			} else {
				if j%2 == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("^")
				}
			}
		}
		fmt.Println()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_1

func splitNumbers(file string) {
	var str string
	var num = 0
	loadArr(file)
	for _, str = range arr {
	}
	arrStr := strings.Split(str, ",")
	for _, i := range arrStr {
		d, err := strconv.Atoi(i)
		if err == nil {
			if (d > 0) && (d%2 == 0) {
				fmt.Println(d)
				num++
			}
		}
	}
	fmt.Println("Number of positive even numbers = ", num)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_2

func Task_2(fail string) {
	loadArr(fail)
	var str string
	var strNum string
	for _, str = range arr {
	}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			continue
		}
		strNum += string(str[i])
	}
	var card string
	fmt.Println("Enter bank card number")
	fmt.Scan(&card)
	strings.TrimRight(card, "/n")
	if len(card) == 16 {
		if card == strNum {
			fmt.Println(str[:14], "**** - is true")

		}else {fmt.Println("Your card is wrong")}
	}else {fmt.Println("Your card is wrong")}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_3

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}
func printFib(j int) {
	for i := 0; i < j; i++ {
		fmt.Println(fib(i))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_4

func palindrome(file string) {
	var str string

	loadArr(file)
	//var ch string
	palindrome := "0"
	for _, s := range arr {
		str += s
	}
	for i := 0; i < len(str); i++ {
		if i < 1 {
			continue
		}
		if str[i] == str[i-1] {
			palindrome = srchPalind_1((i - 1), i, str)
			fmt.Println(palindrome)
		}
		if (i + 1) != len(str) {
			if str[i+1] == str[i-1] {
				palindrome = srchPalind_2((i - 1), (i + 1), str)
				fmt.Println(palindrome)
			}
		}
	}
	if palindrome == "0" {
		fmt.Println("palindrome = ", palindrome)
	}
}

func srchPalind_1(nL, nR int, str string) string {
	var palind string
	for i := 0; ; i++ {
		if (nL - i) == 0 {
			for j := 0; j <= (nR + i); j++ {
				palind += string(str[j])
			}
			break
		}
		if (nR + i) == len(str) {
			for j := (nL - i); j <= (nR + i); j++ {
				palind += string(str[j])
			}
			break
		}
		if (str[nL-i]) != (str[nR+i]) {
			for j := (nL - i + 1); j <= (nR + i - 1); j++ {
				palind += string(str[j])
			}
			break
		}
	}
	return palind
}

func srchPalind_2(nL, nR int, str string) string {
	var palind string
	for i := 0; ; i++ {
		if (nL - i) == 0 {
			for j := 0; j <= (nR + i); j++ {
				palind += string(str[j])
			}
			break
		}
		if (nR + i) == len(str) {
			for j := (nL - i); j <= (nR + i); j++ {
				palind += string(str[j])
			}
			break
		}
		if (str[nL-i]) != (str[nR+i]) {
			for j := (nL - i + 1); j <= (nR + i - 1); j++ {
				palind += string(str[j])
			}
			break
		}
	}
	return palind
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task_5

func luckyTickets(min, max int) {
	Min := min
	Max := max

	var easyFormula int
	var hardFormula int

	easyFormula = easyForm(Min, Max)
	hardFormula = hardForm(Min, Max)

	fmt.Println("EasyFormula = ", easyFormula)
	fmt.Println("HardFormula", hardFormula)
}

func easyForm(min, max int) int {
	easyFormula := 0
	for i := min; i <= max; i++ {
		dig1 := i / 1000
		dig2 := i % 1000
		if ((dig1 / 100) + ((dig1 / 10) % 10) + (dig1 % 10)) == ((dig2 / 100) + ((dig2 / 10) % 10) + (dig2 % 10)) {
			//fmt.Println(i)
			easyFormula++
		}
	}
	return easyFormula
}

func hardForm(min, max int) int {
	hardFormula := 0
	for i := min; i <= max; i++ {
		dig1 := 0
		dig2 := 0
		if (i%10)%2 == 0 {
			dig1 += i % 10
		} else {
			dig2 += i % 10
		}
		if (((i / 10) % 10) % 2) == 0 {
			dig1 += (i / 10) % 10
		} else {
			dig2 += (i / 10) % 10
		}
		if (((i / 100) % 10) % 2) == 0 {
			dig1 += (i / 100) % 10
		} else {
			dig2 += (i / 100) % 10
		}
		if (((i / 1000) % 10) % 2) == 0 {
			dig1 += (i / 1000) % 10
		} else {
			dig2 += (i / 1000) % 10
		}
		if (((i / 10000) % 10) % 2) == 0 {
			dig1 += (i / 10000) % 10
		} else {
			dig2 += (i / 10000) % 10
		}
		if (((i / 100000) % 10) % 2) == 0 {
			dig1 += (i / 100000) % 10
		} else {
			dig2 += (i / 100000) % 10
		}
		if dig1 == dig2 {
			//fmt.Println(i)
			hardFormula++
		}
	}
	return hardFormula
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	//fmt.Println("///////////////////////////////////////////////////////////////")
	//chessboard(4, 6)
	//fmt.Println("///////////////////////////////////////////////////////////////")
	//splitNumbers("Task01.txt")
	//fmt.Println("///////////////////////////////////////////////////////////////")
	//Task_2("Task02.txt")
	//fmt.Println("///////////////////////////////////////////////////////////////")
	//printFib(10)
	//fmt.Println("///////////////////////////////////////////////////////////////")
	//palindrome("Task04.txt")
	fmt.Println("///////////////////////////////////////////////////////////////")
	luckyTickets(120123, 320320)
}
