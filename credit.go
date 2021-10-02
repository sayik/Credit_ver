package main

import (
	"fmt"
)

var credit_card int64

func digits(credit_num int64) int
func multiply_even(credit_num int64, counts int) int
func addition(counts int, credit_num int64) int
func check_valid(total_count int) int
func brand(counts int, credit_num int64) int

func main() {
	fmt.Println("Input credit card number here :")
	fmt.Scanln(&credit_card)
	/*fmt.Println(credit_card)*/
	const counts int = digits(credit_card)
	const total_value int = multiply_even(credit_card, counts) + addition(counts, credit_card)
	if(check_valid(total_value))
		fmt.Println("The credit card is valid and it is a ")
	else
	{
		fmt.Println("Card is invalid! cut it \n")
		return 0
	}

	switch brand(counts, credit_card)
	{
	case 5:
		fmt.Println("Mastercard.\n")
	case 4:
		fmt.Println("Visa. \n")
	case 3:
		fmt.Println("American Express. \n")
	default
		break
	}



}

func digits(credit_num int64) int {

	val := 1
	for {
		credit_num/=10
		if(credit_num)
			val++
		else
			break;
	}
	return val

}

func multiply_even(credit_num int64, counts int) int {
	r := 0
	for j := counts/2; j>0; j-- {
		x int = ((credit_num%100)/10)*2
		credit_num/=100
		if x>=10
			r+=(x%10)+(x/10)
		else
			r+=x
	}
	return r
}


func addition(counts int, credit_num int64) int {
	total := 0
	for i:= counts/2; i>0; i--
	{
		total+= credit_num%10
		credit_num/=100
	}
	return total
}

func check_valid(total_count int) int {
	if !total_count%10
		return 1
	else 
		return 0
}

func brand(counts int, credit_num int64) int {
	if counts==13 {
		for j :=counts; j>1; j--
			credit_num/=10;
	}
	else if counts==16
		for j := counts; j>1; j--
			credit_num/=10
	else 
		return 0
	return credit_num
}