package convert

import "fmt"

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

// SumHardMode - сумма двух бинарных чисел (ручной режим)
func SumHardMode(first, second []int) []int {
	var (
		reversResult []int // конечный результат - перевернутый
		result       []int // результат
		lenfirst     int   // длина первого числа
		lensecond    int   // длина второго числа
		memory       int   // память
	)

	lenfirst = len(first)
	lensecond = len(second)

	// Если первое число длиннее второго
	if lenfirst >= lensecond {
		j := lensecond - 1
		for i := lenfirst - 1; i >= 0; i-- {
			//если существует второе число
			if j >= 0 {
				if res := first[i] + second[j]; res <= 1 {
					if memory > 0 {
						if resMemory := res + memory; resMemory <= 1 {
							result = append(result, resMemory)
							memory--
						} else {
							result = append(result, 0)
						}
					} else {
						result = append(result, res)
					}
				} else {
					if memory > 0 {
						result = append(result, 1)
					} else {
						memory++
						result = append(result, 0)
					}
				}
				j--
			} else { //Если второе число кончилось
				if memory > 0 {
					if res := first[i] + memory; res <= 1 {
						result = append(result, res)
						memory--
					} else {
						result = append(result, 0)
					}
				} else {
					result = append(result, first[i])
				}
			}
		}
		if memory > 0 {
			result = append(result, memory)
			memory--
		}
		fmt.Println()
	} else { // Если второе число длиннее первого
		j := lenfirst - 1
		for i := lensecond - 1; i >= 0; i-- {
			if j >= 0 {
				if res := second[i] + first[j]; res <= 1 {
					if memory > 0 {
						if resMemmory := res + memory; resMemmory <= 1 {
							result = append(result, resMemmory)
							memory--
						} else {
							result = append(result, 0)
						}

					} else {
						result = append(result, res)
					}
				} else {
					if memory > 0 {
						result = append(result, 1)
					} else {
						memory++
						result = append(result, 0)

					}

				}
				j--
			} else {
				if memory > 0 {
					if res := second[i] + memory; res <= 1 {
						result = append(result, res)
						memory--
					} else {
						result = append(result, 0)
					}
				} else {
					result = append(result, second[i])
				}
			}
		}
		if memory > 0 {
			result = append(result, memory)
			memory--
		}
		for i := len(result) - 1; i >= 0; i-- {
			reversResult = append(reversResult, result[i])
		}
	}
	return reversResult
}
