package piscine

import "ft"

func PrintAlphabet() {

	for r := 'a'; r <= 'z'; r++ {
	if err := ft.PrintRune(r); err != nil {
	panic(err)
	}
	}
	ft.PrintRune('\n')
}
