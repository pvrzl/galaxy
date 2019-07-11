package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/homina/galaxy/internal"
	"github.com/homina/galaxy/pkg/convertRoman"
)

type Iroman interface {
	SetAlias(string, string) (convertRoman.RomanType, error)
	SetValue(string) convertRoman.RomanType
	Value() (int, error)
	IsNil() bool
	IsAlias(string) bool
}

var (
	// ErrInvalidCommand Occurred when parser can't recognize the command
	ErrInvalidCommand = errors.New("I have no idea what you are talking about")
	// ErrUnknowPrice Occurred when parser can't determine unknown word price
	ErrUnknowPrice = errors.New("One unknown word please, i can't calculate two or more unknown word at once")
	// ErrUnknownRomanSymbols Occurred when parser found unknown roman symbol
	ErrUnknownRomanSymbols = errors.New("Unknown roman symbol")
	validSymbols           = []string{"I", "V", "X", "L", "C", "D", "M"}
	roman                  Iroman
	thingsPrice            map[string]float32
	things                 = []string{}
)

// Checkinput is a function to detect wether the input come from file or stdin
func CheckInput() {
	if thingsPrice == nil {
		thingsPrice = make(map[string]float32)
	}
	if roman == nil {
		roman = convertRoman.New("")
	}
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
		err = commandParser(scanner.Text())
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
		err = commandParser(cmdString)
		if err != nil {
			panic(err)
		}
	}
}

// exception to prevent force quit when panic happened, so the game can be continue
func exception() {
	if r := recover(); r != nil {
		defer exception()
		fmt.Printf("%+v\n", r)
		startShell()
	}
}

func commandParser(cmdString string) error {
	// split the command to two parts, it always has two parts pattern, no more no less
	cmdString = strings.TrimRight(cmdString, "?")
	cmdString = strings.TrimSpace(cmdString)
	cmds := strings.Split(cmdString, " is ")
	if len(cmds) != 2 {
		return ErrInvalidCommand
	}

	secCmds := strings.Split(cmds[1], " ")
	firstCmds := strings.Split(cmds[0], " ")
	lenFirstCmd := len(firstCmds)
	lenSecCmd := len(secCmds)
	// if it has 1 word for each parts, then it must be define the alias for roman
	if (lenFirstCmd == 1) && (lenSecCmd == 1) {
		if !internal.Contains(validSymbols, cmds[1]) {
			return ErrUnknownRomanSymbols
		}
		roman, _ = roman.SetAlias(cmds[1], cmds[0])
		return nil
	}

	// if it has numeral value and "credits" word in second part, it must be define "things" price (eg: metal), and also we must check the unknown word from the first part, make sure the unknown word only occured once and in the last part of sentence
	if lenSecCmd == 2 && strings.HasSuffix(strings.ToLower(cmds[1]), "credits") {

		i, err := strconv.Atoi(secCmds[0])
		if err != nil {
			return ErrInvalidCommand
		}

		count := 0
		thing := ""
		for _, word := range firstCmds {
			if !roman.IsAlias(word) {
				count++
				thing = word
			}
		}

		if count != 1 {
			return ErrUnknowPrice
		}

		if firstCmds[lenFirstCmd-1] != thing {
			return ErrInvalidCommand
		}

		roman = roman.SetValue(strings.Replace(cmds[0], " "+thing, "", 1))
		value, err := roman.Value()
		if err != nil {
			return err
		}
		thingsPrice[thing] = float32(i) / float32(value)
		if !internal.Contains(things, thing) {
			things = append(things, thing)
		}
		return nil

	}

	// if the command start with "how much" or "how many credits", then it must calculate the command based on collected data,
	// also it check if the command has "things" word at last sentence, if not calculate roman as usual, else calculate roman and multiply it with "things" price
	lowerFirstCmd := strings.ToLower(cmds[0])
	if strings.HasPrefix(lowerFirstCmd, "how much") || strings.HasPrefix(lowerFirstCmd, "how many credits") {

		price, ok := thingsPrice[secCmds[lenSecCmd-1]]
		if ok {
			roman = roman.SetValue(strings.Replace(cmds[1], " "+secCmds[lenSecCmd-1], "", 1))
			value, err := roman.Value()
			if err != nil {
				return err
			}
			fmt.Println(cmds[1], "is", float32(value)*price)
			return nil
		} else {
			roman = roman.SetValue(cmds[1])
			value, err := roman.Value()
			if err != nil {
				return err
			}
			fmt.Println(cmds[1], "is", value)
			return nil
		}

	}

	return ErrInvalidCommand
}
