package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Task_0
//Print a chessboard with the specified dimensions of height and width, according to the example height - 4 and wigth 6:
func chessboard() {
	height := 4
	width := 6

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
var arr []string

func loadArr(file string) error {
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
func scanArrStr() {
	var str string
	var num = 0
	loadArr("Task01.txt")
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

//Task 3

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return  fib(n-2) + fib(n-1)
}
func printFib(j int)  {
	for i:=0 ; i < j; i++ {
		fmt.Println(fib(i))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task 4

func palindrome()  {
	var str string
	var num int


}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Task 5

func luckyTickets(min, max int)  {
	Min := min
	Max := max

	var easyFormula int
	var hardFormula int

	easyFormula = easyForm(Min,Max)
	hardFormula = hardForm(Min,Max)

	fmt.Println("EasyFormula = ", easyFormula)
	fmt.Println("HardFormula",hardFormula)
}


func easyForm(min, max int) int  {
	easyFormula := 0
	for i := min; i <= max; i++ {
		dig1 := i/1000
		dig2 := i%1000
		if ((dig1/100)+((dig1/10)%10)+(dig1%10)) == ((dig2/100)+((dig2/10)%10)+(dig2%10)){
			//fmt.Println(i)
			easyFormula++
		}
	}
	return easyFormula
}

func hardForm(min, max int) int  {
	hardFormula := 0
	for i := min; i <= max; i++ {
		dig1 := 0
		dig2 := 0
		if (i%10)%2 == 0 {
			dig1+=i%10
		}else {
			dig2 += i%10
		}
		if(((i/10)%10)%2) == 0{
			dig1+=(i/10)%10
		}else {
			dig2+=(i/10)%10
		}
		if(((i/100)%10)%2) == 0{
			dig1+=(i/100)%10
		}else {
			dig2+=(i/100)%10
		}
		if(((i/1000)%10)%2) == 0{
			dig1+=(i/1000)%10
		}else {
			dig2+=(i/1000)%10
		}
		if(((i/10000)%10)%2) == 0{
			dig1+=(i/10000)%10
		}else {
			dig2+=(i/10000)%10
		}
		if(((i/100000)%10)%2) == 0{
			dig1+=(i/100000)%10
		}else {
			dig2+=(i/100000)%10
		}
		if dig1== dig2 {
			//fmt.Println(i)
			hardFormula++
		}
	}
	return hardFormula
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	//	chessboard()
	//scanArrStr()
	//printFib(10)
	//luckyTickets(120123,320320)

}
