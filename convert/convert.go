package convert

import (
	"math"
)

// BinToDec - перевод двоичного числа в десятичную систему счисления
func BinToDec(binary []int) int {
	var (
		decimal int
		pow     int
	)

	pow = len(binary) - 1
	for i := 0; i < len(binary); i++ {
		decimal += binary[i] * int(math.Pow(2, float64(pow)))
		pow--
	}

	return decimal
}

// ToDec - перевод из любой системы счисления в 10-тичную
func ToDec(number []int, base int) int {
	var (
		decimal int
		pow     int
	)

	pow = len(number) - 1
	for i := 0; i < len(number); i++ {
		decimal += number[i] * int(math.Pow(float64(base), float64(pow)))
		pow--
	}

	return decimal
}

// DecTo - перевод из 10-ой системы счисления в любую
func DecTo(decimal, base int) []int {
	var (
		result  []int
		remains []int
	)

	for {
		remains = append(remains, decimal%base)
		decimal /= base
		if decimal == 0 {
			break
		}
	}

	for i := len(remains) - 1; i >= 0; i-- {
		result = append(result, remains[i])
	}

	return result
}

// **************************************************
// ******************* ПРОБА ПЕРА *******************
// **************************************************

// DecToBin - перевод десятичного числа в двоичную систему счисления
func DecToBin(decimal int) []int {
	var (
		binary  []int // двоичный результат числа
		remains []int // для остатков
	)

	for {
		remains = append(remains, decimal%2)
		decimal /= 2

		if decimal == 0 {
			break
		}
	}

	for i := len(remains) - 1; i >= 0; i-- {
		binary = append(binary, remains[i])
	}

	return binary
}

// DecToThree - перевод десятичного числа в троичную систему счисления
func DecToThree(decimal int) []int {
	var (
		result  []int
		remains []int
	)

	for {
		remains = append(remains, decimal%3)
		decimal /= 3
		if decimal == 0 {
			break
		}
	}

	for i := len(remains) - 1; i >= 0; i-- {
		result = append(result, remains[i])
	}

	return result
}
