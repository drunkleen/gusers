package utils

import "os"

// ParseArgsAndExitIfNeeded parses the arguments and exits if it encounters
// --help or --version. If there are more than two arguments, it prints an
// error message and exits with status 1.
func ParseArgsAndExitIfNeeded(args []string) {
	if len(args) <= 1 || args[1] == "--" {
		return
	}

	findInArgs(args, "--v", "--version", printVersion)
	findInArgs(args, "--h", "--help", printHelp)

	if len(args) > 2 {
		os.Stdout.Write([]byte(
			"users: extra operand '" + args[2] + "'\n" +
				"Try 'users --help' for more information.\n"))
		os.Exit(1)
	}

	os.Exit(0)
}

// findInArgs iterates through args and checks if argSign or argSign2
// exist in the list. If either is found, the run function is executed
// and the program exits with status 1.
func findInArgs(args []string, argSign, argSign2 string, run func()) {
	for _, arg := range args {
		if arg == argSign || arg == argSign2 {
			run()
			os.Exit(1)
		}
	}
}

// printTag takes a single arg string and writes an error message to
// os.Stdout and exits with a non-zero status if the arg is invalid.
// It checks if the arg is a short option (single character) or a long
// option (multiple characters) and writes the appropriate error
// message.
func printTag(arg string) {
	if len(arg) > 2 {
		return
	}
	if arg[0] == '-' {
		os.Stdout.Write([]byte(
			"users: invalid option -- '" + arg[1:] + "'\n" +
				"Try 'users --help' for more information.\n",
		))
		return
	}
	if arg[:2] == "--" {
		os.Stdout.Write([]byte(
			"users: unrecognized option '" + arg + "'\n" +
				"Try 'users --help' for more information.\n",
		))
	}
}

// printVersion writes the version information for the users command to os.Stdout.
// It includes the version number, copyright information, license information,
// and the name of the author.
func printVersion() {
	os.Stdout.Write([]byte(
		"gusers v0.1" +
			"Copyright (C) 2025 Free Software Foundation, Inc.\n" +
			"License GPLv3+: GNU GPL version 3 or later <https://fsf.org/>.\n" +
			"This is free software: you are free to change and redistribute it.\n" +
			"There is NO WARRANTY, to the extent permitted by law.\n\n" +
			"Written by DrunkLeen.\n",
	))
}

// printHelp writes the help message for the users command to os.Stdout.
// It describes the usage, options, and provides links to the online
// documentation and the local GNU info page.
func printHelp() {
	os.Stdout.Write([]byte(
		"Usage: users [OPTION]... [FILE]\n" +
			"Output who is currently logged in according to FILE.\n" +
			"If FILE is not specified, use /var/run/utmp.  /var/log/wtmp as FILE is common.\n\n" +
			"--help        display this help and exit" +
			"--version     output version information and exit\n\n" +
			"GNU coreutils online help: <https://www.gnu.org/software/coreutils/>\n" +
			"Full documentation <https://github.com/drunkleen/gusers>\n" +
			"or available locally via: info '(coreutils) users invocation'\n",
	))
}
