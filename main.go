package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ryoga-exe/umesh/internal/builtin_commands"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	buitinCommands := map[string]func([]string) error{
		"cd":   builtin_commands.Cd,
		"exit": builtin_commands.Exit,
	}

	for {
		pwd, _ := os.Getwd()
		fmt.Print(pwd, "> ")

		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		argv := strings.Split(string(line), " ")

		buitinCommand, ok := buitinCommands[argv[0]]

		if ok {
			err = buitinCommand(argv)
		} else {
			err = execCommand(argv)
		}
	}
}

func execCommand(argv []string) (err error) {
	cmd, err := absPathWithPATH(string(argv[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	log.Printf("command %s\n", cmd)

	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

	process, err := os.StartProcess(cmd, argv, &procAttr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = process.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return nil
}

func absPathWithPATH(target string) (targetAbsPath string, err error) {
	targetFileName := filepath.Base(target)
	log.Printf("target %s\n", target)
	log.Printf("targetFileName %s\n", targetFileName)

	// if the specified string is a path
	if target != targetFileName {
		if filepath.IsAbs(target) {
			// if absolute path
			targetAbsPath = target
		} else {
			// if relative path
			targetAbsPath, err = filepath.Abs(target)
			if err != nil {
				log.Fatal(err)
			}
		}

		if Exists(targetAbsPath) {
			return targetAbsPath, nil
		} else {
			return "", fmt.Errorf("%s: no such file or directory", targetAbsPath)
		}
	}

	// if the specified string is a file name

	// find from $PATH
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		log.Printf("%s\n", path)
		targetAbsPath = filepath.Join(path, targetFileName)
		if Exists(targetAbsPath) {
			log.Printf("find in PATH %s\n", targetAbsPath)
			return targetAbsPath, nil
		}
	}
	return "", fmt.Errorf("%s: no such file or directory", targetFileName)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
