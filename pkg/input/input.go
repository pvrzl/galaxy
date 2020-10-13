package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	romanPkg "galaxy/pkg/roman"
	"galaxy/pkg/romanParser"
)

type Iparser interface {
	Parse(string) error
}

var (
	parser Iparser
)

func init() {
	roman := romanPkg.New("")
	parser = romanParser.NewDefaultParser(roman)
}

// exception to prevent force quit when panic happened, so the game can be continue
func exception() {
	if r := recover(); r != nil {
		defer exception()
		fmt.Printf("%+v\n", r)
		startShell()
	}
}

// Checkinput is a function to detect wether the input come from file or stdin
func CheckInput() {
	if len(os.Args) > 1 {
		readFile(os.Args[1])
	} else {
		startShell()
	}
}

func readFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = parser.Parse(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func startShell() {
	defer exception()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmdString = strings.TrimSpace(cmdString)
		if cmdString == "exit" {
			break
		}
		if cmdString == "" {
			continue
		}
		err = parser.Parse(cmdString)
		if err != nil {
			panic(err)
		}
	}
}
