package main

import (
	"fmt"
)

var credit_card int64

/*func digits(credit_num int64) int
func multiply_even(credit_num int64, counts int) int
func addition(counts int, credit_num int64) int
func check_valid(total_count int) int
func brand(counts int, credit_num int64) int*/

func digits(credit_num int64) int {

	var val int = 1

	for {
		credit_num /= 10
		if credit_num > 0 {
			val++
		} else {
			break
		}
	}
	return val

}

func multiply_even(credit_num int64, counts int) int64 {
	var r int64 = 0
	for j := counts / 2; j > 0; j-- {
		x := ((credit_num % 100) / 10) * 2
		credit_num /= 100
		if x >= 10 {
			r += ((x % 10) + (x / 10))
		} else {
			r += x
		}
	}
	return r
}

func addition(counts int, credit_num int64) int64 {
	var total int64 = 0
	for i := counts / 2; i > 0; i-- {
		total += credit_num % 10
		credit_num /= 100
	}
	return total
}

func check_valid(total_count int) int {
	if !(total_count%10 == 0) {
		return 1
	} else {
		return 0
	}
}

func brand(counts int, credit_num int64) int64 {
	if counts == 13 {
		for j := counts; j > 1; j-- {
			credit_num /= 10
		}
	} else if counts == 16 {
		for j := counts; j > 1; j-- {
			credit_num /= 10
		}
	} else {
		return 0
	}
	return credit_num
}

func main() {
	fmt.Println("Input credit card number here :")
	fmt.Scanln(&credit_card)
	/*fmt.Println(credit_card)*/
	var counts int = digits(credit_card)
	var total_value int64 = multiply_even(credit_card, counts) + addition(counts, credit_card)
	//var m int = check_valid(int(total_value))
	if check_valid(int(total_value)) == 0 {
		fmt.Println("The credit card is valid and it is a ")
	} else {
		fmt.Println("Card is invalid! cut it :P \n")
		return //- why doesn't it work!! I want to break the main function if card is invalid
	}

	switch brand(counts, credit_card) {
	case 5:
		fmt.Println("Mastercard.\n")
	case 4:
		fmt.Println("Visa. \n")
	case 3:
		fmt.Println("American Express. \n")
	default:
		break

	}

}
