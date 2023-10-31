package builtin_commands

import (
	"fmt"
	"log"
	"os"
)

func Cd(argv []string) (err error) {
	var dir string
	argc := len(argv)
	if argc == 1 {
		dir, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
	} else if argc == 2 {
		dir = argv[1]
	} else {
		return fmt.Errorf("%s: too many arguments", "cd")
	}
	return os.Chdir(dir)
}
