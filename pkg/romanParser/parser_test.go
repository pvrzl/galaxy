package romanParser

import (
	"testing"

	convertRoman "galaxy/pkg/roman"
)

func TestCommandParser(t *testing.T) {
	roman := convertRoman.New("")
	parser := NewDefaultParser(roman)

	listTests := map[string]error{
		"bogj":        ErrInvalidCommand,
		"bogj is IGG": ErrUnknownRomanSymbols,
		"bogj is I":   nil,
		"glob is I":   nil,
		"prok is V":   nil,
		"pish is X":   nil,
	}

	for k, v := range listTests {
		if parser.Parse(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, parser.Parse(k))
		}
	}

	listTests = map[string]error{
		"bogj is jgob credits":                      ErrInvalidCommand,
		"bogj jgob hgoj is 800 credits":             ErrUnknowPrice,
		"bogj jgob bogj is 800 credits":             ErrInvalidCommand,
		"bogj bogj bogj bogj hgoj is 900 credits":   convertRoman.ErrInvalidRoman,
		"bogj jgob is 800 credits":                  nil,
		"glob prok has less Credits than pish pish": nil,
	}

	for k, v := range listTests {
		if parser.Parse(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, parser.Parse(k))
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
		if parser.Parse(k) != v {
			t.Errorf("%v should return %v, instead got : %v", k, v, parser.Parse(k))
		}
	}

}
