package piscine

import "ft"

func PrintDigits() {
	for i:= 0; i<= 9; i++{
	if err := ft.PrintRune(rune(i) + '0'); err != nil {
	panic(err)
	}
	}
	ft.PrintRune('\n')
}
