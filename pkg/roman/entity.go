package roman

import (
	"errors"
	"regexp"
	"strings"

	"galaxy/internal"
)

var (
	// ErrInvalidRoman occured when romanString is not valid roman numeric
	ErrInvalidRoman = errors.New("Requested number is in invalid format")
	// ErrInvalidRomanChar occured when invalid roman character found
	ErrInvalidRomanChar = errors.New("Invalid roman character")
	symbols             = []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanType hold the romanstring and roman character alias
type RomanType struct {
	romanString string
	alias       map[string]string
}

// New return predefined roman type
func New(romanString string) *RomanType {
	return &RomanType{
		romanString: romanString,
		alias:       make(map[string]string),
	}
}

// IsNil Check wether the alias is nil or not
func (r *RomanType) IsNil() bool {
	return r.alias == nil
}

// SetAlias Set alias for roman character
func (r *RomanType) SetAlias(origin string, alias string) error {
	if !internal.Contains([]string{"I", "V", "X", "L", "C", "D", "M"}, origin) {
		return ErrInvalidRomanChar
	}

	r.alias[alias] = origin
	return nil
}

// IsAlias is a func to check if the alias is already taken
func (r *RomanType) IsAlias(alias string) bool {
	_, ok := r.alias[alias]
	return ok
}

// SetValue is a func to set romanstring value
func (r *RomanType) SetValue(value string) {
	r.romanString = value

}

// Value is a func to get numeric value of romanString
func (r *RomanType) Value() (int, error) {
	if len(r.alias) > 0 {
		words := strings.Split(r.romanString, " ")
		romanString := []string{}
		for _, word := range words {
			if val, ok := r.alias[word]; ok {
				romanString = append(romanString, val)
			} else {
				romanString = append(romanString, word)
			}
		}
		r.romanString = strings.Join(romanString, "")
	}

	if !r.validation() {
		return 0, ErrInvalidRoman
	}

	c := strings.Split(r.romanString, "")
	value := r.compute(c)
	return value, nil
}

func (r *RomanType) compute(c []string) int {
	lastDigit := 1000
	latin := 0
	for _, v := range c {
		var digit int
		if val, ok := r.alias[v]; ok {
			digit = romanMap[val]
		} else {
			digit = romanMap[v]
		}
		if lastDigit < digit {
			latin -= 2 * lastDigit
		}
		lastDigit = digit
		latin += lastDigit
	}
	return latin
}

func (r *RomanType) validation() bool {
	if r.romanString == "" {
		return false
	}

	if check, _ := regexp.MatchString("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", r.romanString); check == true {
		return check
	}
	return false
}

// Reverse is a func to convert int to romanString
func Reverse(num int) string {
	if num == 0 {
		return ""
	}

	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	roman := ""
	i := 0

	for num > 0 {

		k := num / values[i]
		for j := 0; j < k; j++ {

			roman += symbols[i]

			num -= values[i]
		}
		i++
	}
	return roman
}
