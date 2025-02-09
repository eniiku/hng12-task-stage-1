package main

import "strconv"


// IsPrime checks if a number is prime.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPerfect checks if a number is perfect.
func IsPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// IsArmstrong checks if a number is an Armstrong number.
func IsArmstrong(n int) bool {
	digits := strconv.Itoa(n)
	sum := 0
	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sum += d * d * d
	}
	return sum == n
}

// DigitSum calculates the sum of the digits of a number.
func DigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
