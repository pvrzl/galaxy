package convertRoman

import (
	"strings"
	"testing"
)

func TestNewRoman(t *testing.T) {
	str := "hello"
	var roman interface{}
	roman = New(str)
	data, ok := roman.(RomanType)
	if !ok {
		t.Errorf("new func should return roman type")
	}

	if data.romanString != str {
		t.Errorf(`new func should have "` + str + `" value when called with "` + str + `"`)
	}
}

func TestIsNil(t *testing.T) {
	roman := RomanType{}
	if !roman.IsNil() {
		t.Errorf(`roman alias should be nil if not invoked by New func`)
	}

	roman = New("")
	if roman.IsNil() {
		t.Errorf(`roman alias should not be nil if already invoked by New func`)
	}
}

func TestValidRoman(t *testing.T) {
	testLists := map[string]bool{
		"":      false,
		"XXX69": false,
		"ML":    true,
		"ml":    false,
	}
	// you can continue to make test list if you want

	for k, v := range testLists {
		roman := New(k)
		if roman.validation() != v {
			t.Errorf("the validation func should return %v when called with %v", v, k)
		}
	}

}

func TestReverse(t *testing.T) {
	testLists := map[string]int{
		"":           0,
		"XXV":        25,
		"LXXVII":     77,
		"CXXXVIII":   138,
		"CDXVIII":    418,
		"CMLXXXVIII": 988,
	}

	for k, v := range testLists {
		if k != Reverse(v) {
			t.Errorf("the to Roman func should return %v when called with %v", k, v)
		}
	}
}

func TestSetValue(t *testing.T) {
	roman := New("")
	if roman.romanString != "" {
		t.Errorf("the roman string should be empty if invoked by empty params ")
	}
	roman = roman.SetValue("halo")
	if roman.romanString != "halo" {
		t.Errorf("the roman string should be halo if set by setvalue fn")
	}
}

func TestAlias(t *testing.T) {
	roman := New("MM")
	_, err := roman.SetAlias("F", "Fck")
	if err != ErrInvalidRomanChar {
		t.Errorf("set alias should return error when called with invalid roman char")
	}

	_, err = roman.SetAlias("I", "Idiot")
	if err != nil {
		t.Errorf("set alias should not return error when called with valid roman char")
	}

	if roman.alias["Idiot"] != "I" {
		t.Errorf("set alias should set Idiot as alias to I")
	}
}

func TestIsAlias(t *testing.T) {
	roman := New("")
	if roman.IsAlias("huha") {
		t.Errorf("is alias should return false, because huha not set yet")
	}

	roman.SetAlias("I", "huha")
	if !roman.IsAlias("huha") {
		t.Errorf("now is alias should return true, because huha already set before")
	}
}

func TestCompute(t *testing.T) {
	testLists := map[string]int{
		"XXV":        25,
		"LXXVII":     77,
		"CXXXVIII":   138,
		"CDXVIII":    418,
		"CMLXXXVIII": 988,
	}

	for k, v := range testLists {
		roman := New(k)
		value := roman.compute(strings.Split(k, ""))
		if value != v {
			t.Errorf("the compute func should return %v when called with %v", v, k)
		}
	}

	// convert with alias
	str := "glob glob fyuh"
	roman := New(str)
	roman, _ = roman.SetAlias("X", "glob")
	roman, _ = roman.SetAlias("V", "fyuh")
	roman, _ = roman.SetAlias("C", "huha")

	value := roman.compute(strings.Split(str, " "))
	if value != 25 {
		t.Errorf("the compute with alias func should return %v when called with %v, got %v", 25, "glob glob fyuh", value)

	}

	roman.romanString = "huha glob fyuh"
	value = roman.compute(strings.Split(roman.romanString, " "))
	if value != 115 {
		t.Errorf("the compute with alias func should return %v when called with %v, got %v", 115, "huha glob fyuh", value)

	}

}

func TestValue(t *testing.T) {
	testListsInvalid := map[string]error{
		"":               ErrInvalidRoman,
		"XXX69":          ErrInvalidRoman,
		"glob glob fyuh": ErrInvalidRoman,
		"globglobfyuh":   ErrInvalidRoman,
		"XXV":            nil,
	}

	for k, v := range testListsInvalid {
		roman := New(k)
		_, err := roman.Value()
		if err != v {
			t.Errorf("the value func should return error %v when called with %v", v, k)
		}
	}

	testLists := map[string]int{
		"XXV":        25,
		"LXXVII":     77,
		"CXXXVIII":   138,
		"CDXVIII":    418,
		"CMLXXXVIII": 988,
	}

	for k, v := range testLists {
		roman := New(k)
		value, _ := roman.Value()
		if value != v {
			t.Errorf("the value func should return %v when called with %v", v, k)
		}
	}

	testListsWithAlias := map[string]int{
		"glob glob fyuh":       25,
		"L glob glob fyuh I I": 77,
	}

	for k, v := range testListsWithAlias {
		roman := New(k)
		roman, _ = roman.SetAlias("X", "glob")
		roman, _ = roman.SetAlias("V", "fyuh")
		roman, _ = roman.SetAlias("C", "huha")
		value, _ := roman.Value()
		if value != v {
			t.Errorf("the value func with alias should return %v when called with %v", v, k)
		}
	}

	testListWithAliasInvalid := map[string]error{
		"fyuh fyuh": ErrInvalidRoman,
	}

	for k, v := range testListWithAliasInvalid {
		roman := New(k)
		roman, _ = roman.SetAlias("V", "fyuh")
		_, err := roman.Value()
		if err != v {
			t.Errorf("the value func with alias which has invalid format should return %v when called with %v", v, k)
		}

	}
}
