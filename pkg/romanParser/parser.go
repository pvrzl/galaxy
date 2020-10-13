package romanParser

import (
	"errors"
	"fmt"
	"galaxy/internal"
	"regexp"
	"strconv"
	"strings"
)

type Iroman interface {
	SetAlias(string, string) error
	SetValue(string)
	Value() (int, error)
	IsNil() bool
	IsAlias(string) bool
}

type Parser struct {
	ValidSymbols   []string
	Roman          Iroman
	ThingsPrice    map[string]float32
	Things         []string
	OperatorString []string
}

var (
	// ErrInvalidCommand Occurred when parser can't recognize the command
	ErrInvalidCommand = errors.New("I have no idea what you are talking about")
	// ErrUnknowPrice Occurred when parser can't determine unknown word price
	ErrUnknowPrice = errors.New("One unknown word please, i can't calculate two or more unknown word at once")
	// ErrUnknownRomanSymbols Occurred when parser found unknown roman symbol
	ErrUnknownRomanSymbols = errors.New("Unknown roman symbol")
)

func NewDefaultParser(roman Iroman) *Parser {
	return &Parser{
		ValidSymbols: []string{"I", "V", "X", "L", "C", "D", "M"},
		Roman:        roman,
		ThingsPrice:  make(map[string]float32),
		Things:       []string{},
		OperatorString: []string{
			" has more Credits than ",
			" has less Credits than ",
			" larger than ",
			" smaller than ",
		},
	}
}

func (p *Parser) Parse(cmdString string) error {
	cmdString = p.sanitizeCmd(cmdString)
	cmds := strings.Split(cmdString, " is ")
	if len(cmds) == 2 {
		return p.isParser(cmds)
	}

	if val, ok := internal.StringHas(cmdString, p.OperatorString); ok {
		return p.operatorParser(cmdString, val)
	}

	return ErrInvalidCommand
}

func (p *Parser) sanitizeCmd(cmdString string) string {
	cmdString = strings.TrimSpace(cmdString)
	cmdString = strings.TrimRight(cmdString, "?")
	cmdString = strings.TrimPrefix(cmdString, "Does ")
	cmdString = strings.TrimPrefix(cmdString, "Is ")
	cmdString = strings.TrimSpace(cmdString)
	singleSpacePattern := regexp.MustCompile(`\s+`)
	cmdString = singleSpacePattern.ReplaceAllString(cmdString, " ")
	return cmdString
}

func (p *Parser) operatorParser(cmdString string, operator string) (err error) {
	var val1, val2 float32
	var value int

	cmds := strings.Split(cmdString, operator)
	if len(cmds) != 2 {
		return ErrInvalidCommand
	}

	secCmds := strings.Split(cmds[1], " ")
	firstCmds := strings.Split(cmds[0], " ")
	lenFirstCmd := len(firstCmds)
	lenSecCmd := len(secCmds)

	price, ok := p.ThingsPrice[firstCmds[lenFirstCmd-1]]
	if ok {
		p.Roman.SetValue(strings.Replace(cmds[0], " "+firstCmds[lenFirstCmd-1], "", 1))
		value, err = p.Roman.Value()
		val1 = float32(value) * price
	} else {
		p.Roman.SetValue(cmds[0])
		value, err = p.Roman.Value()
		val1 = float32(value)
	}

	price, ok = p.ThingsPrice[secCmds[lenSecCmd-1]]
	if ok {
		p.Roman.SetValue(strings.Replace(cmds[1], " "+secCmds[lenSecCmd-1], "", 1))
		value, err = p.Roman.Value()
		val2 = float32(value) * price
	} else {
		p.Roman.SetValue(cmds[1])
		value, err = p.Roman.Value()
		val2 = float32(value)
	}

	if strings.Contains(operator, "more") || strings.Contains(operator, "less") {
		if val1 > val2 {
			fmt.Printf("%s has more Credits than %s\n", cmds[0], cmds[1])
		} else {
			fmt.Printf("%s has less Credits than %s\n", cmds[0], cmds[1])
		}
	}

	if strings.Contains(operator, "larger") || strings.Contains(operator, "smaller") {
		if val1 > val2 {
			fmt.Printf("%s is larger than %s\n", cmds[0], cmds[1])
		} else {
			fmt.Printf("%s is smaller than %s\n", cmds[0], cmds[1])
		}
	}

	return err
}

func (p *Parser) isParser(cmds []string) error {
	secCmds := strings.Split(cmds[1], " ")
	firstCmds := strings.Split(cmds[0], " ")
	lenFirstCmd := len(firstCmds)
	lenSecCmd := len(secCmds)
	// if it has 1 word for each parts, then it must be define the alias for roman
	if (lenFirstCmd == 1) && (lenSecCmd == 1) {
		if !internal.Contains(p.ValidSymbols, cmds[1]) {
			return ErrUnknownRomanSymbols
		}
		p.Roman.SetAlias(cmds[1], cmds[0])
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
			if !p.Roman.IsAlias(word) {
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

		p.Roman.SetValue(strings.Replace(cmds[0], " "+thing, "", 1))
		value, err := p.Roman.Value()
		if err != nil {
			return err
		}
		p.ThingsPrice[thing] = float32(i) / float32(value)
		if !internal.Contains(p.Things, thing) {
			p.Things = append(p.Things, thing)
		}
		return nil

	}

	// if the command start with "how much" or "how many credits", then it must calculate the command based on collected data,
	// also it check if the command has "things" word at last sentence, if not calculate roman as usual, else calculate roman and multiply it with "things" price
	lowerFirstCmd := strings.ToLower(cmds[0])
	if strings.HasPrefix(lowerFirstCmd, "how much") || strings.HasPrefix(lowerFirstCmd, "how many credits") {

		price, ok := p.ThingsPrice[secCmds[lenSecCmd-1]]
		if ok {
			p.Roman.SetValue(strings.Replace(cmds[1], " "+secCmds[lenSecCmd-1], "", 1))
			value, err := p.Roman.Value()
			if err != nil {
				return err
			}
			fmt.Println(cmds[1], "is", float32(value)*price)
			return nil
		} else {
			p.Roman.SetValue(cmds[1])
			value, err := p.Roman.Value()
			if err != nil {
				return err
			}
			fmt.Println(cmds[1], "is", value)
			return nil
		}

	}

	return ErrInvalidCommand
}
