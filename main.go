package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		pwd, _ := os.Getwd()
		fmt.Print(pwd, "> ")

		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		args := strings.Split(string(line), " ")
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

		process, err := os.StartProcess(args[0], args, &procAttr)
		if err != nil {
			log.Fatal(err)
		}

		if _, err = process.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
