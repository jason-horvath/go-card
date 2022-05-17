package main

import (
	"fmt"
	"os"
	"strconv"
)

type CreditCard struct {
	Number int
	CV int
	ExpireMonth int
	ExpireYear int
}

func main() {
	
	card := CreditCard {
		Number: getArg(1),
		CV: getArg(2), 
		ExpireMonth: getArg(3),
		ExpireYear: getArg(4),
	}

	fmt.Println("CC Number", card.Number)
	fmt.Println("CC CV", card.CV)
	fmt.Println("CC Expire Month", card.ExpireMonth)
	fmt.Println("CC Expire Year", card.ExpireYear)
	fmt.Println("\n---- Results ----")
	fmt.Println("CC Luhn Valid:", card.isValidLuhn())
	fmt.Println("CC Prefix Valid:", card.isValidPrefix())
}

func getArg(index int) (int) {
	retVal := 0
	args := os.Args
	numberOrArgs := len(args[1:])
	
	if numberOrArgs >= index {
		argVal, _ := strconv.Atoi(args[index])
		retVal = argVal
	}

	return retVal
}

func (cc *CreditCard) isValidLength() (bool) {
	ccLength := len(string(cc.Number))
	var validLength bool = false

	if ccLength >=13 && ccLength >= 16 {
		validLength = true
	}

	return validLength
}

func (cc *CreditCard) isValidLuhn() (bool) {
	return (cc.getLuhnSum() % 10) == 0 
}

func (cc *CreditCard) getLuhnSum() (int) {
	num := cc.Number
	ccNumStr := strconv.Itoa(num)
	
	var sum int
	for i, _ := range ccNumStr {
		addToSum := num % 10
		if i % 2 == 1 {
			addToSum = addToSum * 2
			if addToSum > 9 {
				addToSum = addToSum % 10 + addToSum / 10
			}
		}
		sum += addToSum
		num = num / 10
	}
	
	return sum
}

func (cc *CreditCard) isValidPrefix() (bool) {
	var validPrefix bool = false
	var numStr = strconv.Itoa(cc.Number)
	if
		string(numStr[0]) == "4" ||
		string(numStr[0]) == "5" ||
		string(numStr[0]) == "6" ||
		string(numStr[0:1]) == "37" {
		validPrefix = true
	}
	

	return validPrefix
}