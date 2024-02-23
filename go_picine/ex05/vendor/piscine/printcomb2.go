package piscine 

import "ft"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1;j < 100;j++ {
		PrintNumber(i)
		ft.PrintRune(' ')
		PrintNumber(j)
		if i != 98 {	
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		}
	}
	ft.PrintRune('\n');
}

func PrintNumber(num int) {
	ft.PrintRune(rune('0' + num / 10))
	ft.PrintRune(rune('0' + num % 10))
}
