package main

import (
	"fmt"
)
const DISCORDID = os.Getenv("DISCORDID")
const DISCOR_TOKEN = os.Getenv("DISCORD_TOKEN")

//うに
func calculate_number(number int) int {
	fmt.Println("Calculating...")

	number *=2

	//ここに計算処理を書く
	return number
}	

//にっしー
func check_number(number int) bool{
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

	fmt.Println("hello");
}
