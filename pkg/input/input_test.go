package input

import (
	"testing"

	"github.com/homina/galaxy/pkg/convertRoman"
)

func TestCommandParser(t *testing.T) {
	roman = convertRoman.New("")
	thingsPrice = make(map[string]float32)
	listTests := map[string]error{
		"bogj":        ErrInvalidCommand,
		"bogj is IGG": ErrUnknownRomanSymbols,
		"bogj is I":   nil,
	}

	for k, v := range listTests {
		if commandParser(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, commandParser(k))
		}
	}

	listTests = map[string]error{
		"bogj is jgob credits":                    ErrInvalidCommand,
		"bogj jgob hgoj is 800 credits":           ErrUnknowPrice,
		"bogj jgob bogj is 800 credits":           ErrInvalidCommand,
		"bogj bogj bogj bogj hgoj is 900 credits": convertRoman.ErrInvalidRoman,
		"bogj jgob is 800 credits":                nil,
	}

	for k, v := range listTests {
		if commandParser(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, commandParser(k))
		}
	}

	listTests = map[string]error{
		"how much is bogj bogj":                nil,
		"how much is bogj jgob":                nil,
		"how much is bogj bogj bogj bogj jgob": convertRoman.ErrInvalidRoman,
		"how much is bogj hyog":                convertRoman.ErrInvalidRoman,
		"how much is bogj bogj bogj bogj hyog": convertRoman.ErrInvalidRoman,
		"how much is":                          ErrInvalidCommand,
		"how much is ha he heheha":             convertRoman.ErrInvalidRoman,
		"hmm apa is gimana ya":                 ErrInvalidCommand,
	}

	for k, v := range listTests {
		if commandParser(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, commandParser(k))
		}
	}

}
