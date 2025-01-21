package main

import (
	"os"

	"github.com/drunkleen/gusers/utils"
	"github.com/drunkleen/gusers/utmpx"
)

func main() {
	utils.ParseArgsAndExitIfNeeded(os.Args)

	for _, ent := range utmpx.GetAll() {
		if len(ent.Line) < 4 || ent.Line[:4] != "pts/" {
			continue
		}

		os.Stdout.Write([]byte(ent.User + " "))
	}
	os.Stdout.Write([]byte("\n"))
}
