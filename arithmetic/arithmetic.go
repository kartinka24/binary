package arithmetic

import "github.com/binary/convert"

// Divide - частное от деления в любой системе счисления
func Divide(first, second []int, base int) []int {
	var result []int

	return result
}

// Difference - разность двух чисел любых систем счисления
func Difference(first, second []int, base int) []int {
	return convert.DecTo((convert.ToDec(first, base) - convert.ToDec(second, base)), base)
}

// Multiply - произведение двух чисел любых систем счисления
func Multiply(first, second []int, base int) []int {
	return convert.DecTo((convert.ToDec(first, base) * convert.ToDec(second, base)), base)
}

// Sum - сумма двух чисел любых систем счисления
func Sum(first, second []int, base int) []int {
	return convert.DecTo((convert.ToDec(first, base) + convert.ToDec(second, base)), base)
}

// *************************************************
// ****************** ПРОБА ПЕРА *******************
// *************************************************

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
	}
	for i := len(result) - 1; i >= 0; i-- {
		reversResult = append(reversResult, result[i])
	}

	return reversResult
}
