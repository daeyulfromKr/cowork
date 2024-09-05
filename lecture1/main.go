package main

import (
	"fmt"
)

func IsPrime(n int) (bool, string) {
	if n <= 0 || n == 1 {
		return false, "negative value and 0 ,1 is not a prime number\n"
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d that is not a prime number\n", n)
		}
	}
	return true, fmt.Sprintf("%d that is a prime number\n", n)
}
