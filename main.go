package main

import (
	"fmt"
)

//うに
func calculate_number(number int) int {
	fmt.Println("Calculating...")

	//ここに計算処理を書く
	return number
}	

//にっしー
func check_number(number int) bool{
	if number > 0 {
        return true
    }
	return false
	//もしnumberが正ならtrueを返す
}


func main() {
	var number int
	fmt.Println("Hello, World!")
	fmt.Scan(&number)

	if check_number(number) {
	number =calculate_number(number)
	}
	else{
		fmt.Println("Please input positive number")
	}
}