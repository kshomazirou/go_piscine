package piscine

import "ft"

func IntToString(num int) string {
	if num == 0 {
		return "0"
	} else if num == 1 {
		return "1"
	} else if num == 2 {
		return "2"
	} else if num == 3 {
		return "3"
	} else if num == 4 {
		return "4"
	} else if num == 5 {
		return "5"
	} else if num == 6 {
		return "6"
	} else if num == 7 {
		return "7"
	} else if num == 8 {
		return "8"
	} else {
		return "9"
	}
}

func ft_itoa(num int) string {
	var result string

	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = IntToString(num % 10) + result
		num /= 10
	}
	return result
}

func ft_putstr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}


func PrintComb(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					ft_putstr(", ")
				}
				ft_putstr(result + ft_itoa(i))
				*count++
			} else {
				PrintComb(n - 1, i, result + ft_itoa(i), count)
			}
		}
	}

}

func PrintCombN(n int) {
	var count int = 0;
	for i := 0; i < 10; i++ {
		if n > 1 { 
			PrintComb(n - 1, i, ft_itoa(i), &count)
		} else {
			if count > 0 {
				ft_putstr(", ")
			}
			ft_putstr(ft_itoa(i))
			count++
		}
	}
	ft_putstr("\n")
}
