package builtin_commands

import "os"

func Exit(argv []string) (err error) {
	os.Exit(0)
	return nil
}
