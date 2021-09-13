package main

import "strconv"

func fizzBuzz(number int) []string {
	res := []string{}
	for i := 1; i <= number; i++ {
		word := ""
		if i%3 == 0 {
			word += "Fizz"
		}
		if i%5 == 0 {
			word += "Buzz"
		}
		if len(word) == 0 {
			word = strconv.Itoa(i)
		}
		res = append(res, word)
	}
	return res
}
