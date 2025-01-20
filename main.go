package main

import (
	"fmt"

	"github.com/drunkleen/gusers/utmpx"
)

func main() {
	for _, ent := range utmpx.GetAll() {
		if len(ent.Line) < 4 || ent.Line[:4] != "pts/" {
			continue
		}

		fmt.Printf("%v ", ent.User)
	}
	fmt.Println()
}
