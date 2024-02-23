package piscine 

import "ft"

func PrintReverseAlphabet() {
	for r := 'z'; r >= 'a'; r-- {
		if err := ft.PrintRune(r); err != nil {
		panic(err);
	}
	}
	ft.PrintRune('\n')
}
