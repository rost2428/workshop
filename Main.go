package main

import (
	"bufio"
	"fmt"
	"os"
)

//Task 0
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
var arr[]string

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


func main() {
	//	chessboard()
	loadArr("E:\\Golang\\pkg\\mod\\github.com\\rost2428\\!go!lang!prog@v0.0.0-20210928123024-e712cf29b280\\Worckshop\\Task01.txt")
	for _, s := range arr {
		fmt.Println(s)
	}
}
