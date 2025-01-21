package main

import (
	"os"

	"github.com/drunkleen/gusers/utils"
	"github.com/drunkleen/gusers/utmpx"
)

// main is the entry point for the gusers application. It parses command-line
// arguments, retrieves all logged-in user entries from the utmpx database, filters
// entries to include only those from pseudo-terminals ("pts/"), and outputs the
// usernames of these sessions to standard output.
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
