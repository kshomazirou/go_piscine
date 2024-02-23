package piscine 

import "ft"

func PrintComb() {

	for i:= 0; i <= 9; i++ {
		for j:= i + 1; j <= 9;j++ {
			for l := j + 1; l <= 9; l++{
			ft.PrintRune(rune(i) + '0');
			ft.PrintRune(rune(j) + '0');
			ft.PrintRune(rune(l) + '0');
			if i != 7 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
}
	ft.PrintRune('\n');
}
