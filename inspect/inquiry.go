package inspect

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const letters = "letters"
const number = "number"
const letterAndNumber = "lettersAndNumbers"
const lettersAndNumbersAndSpecialChar = "lettersAndNumbersAndSpecialChar"
const specialCharacters = "specialChar"
const none = "none"

type Inquiry struct {
	Text                      string
	ContainsNumbers           bool
	ContainsCharacters        bool
	ContainsStrangeCharacters bool
	LanguageType              string
	QtyOfCharacters           int
}

func Inspect(text string) Inquiry {
	var inquiry Inquiry

	if len(text) > 0 {
		text := strings.Trim(text, " ")

		typeLang := RetrieveTypeLanguage(text)

		inquiry = Inquiry{
			Text:                      text,
			ContainsNumbers:           true,
			ContainsCharacters:        false,
			ContainsStrangeCharacters: false,
			LanguageType:              typeLang,
			QtyOfCharacters:           len(text),
		}
	}
	return inquiry
}

func IsLetter(text string) bool {
	if IsEmpty(text) { return false }

	for _, r := range text {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func ContainsCharacterOrNumber(text string) bool {
	if IsEmpty(text) { return false }
	_, err := regexp.MatchString(`^.*[a-zA-Z][0-9].*$`, text)
	return err == nil
}

func IsNumeric(text string) bool {
	_, err := strconv.ParseFloat(text, 64)
	return  err == nil
}

func IsEmpty(text string) bool {
	if len(text) == 0 { return true } else { return false }
}

func IsSpecialCharacter(text string) bool {
	if IsEmpty(text) { return false }
	matched, _ := regexp.MatchString(`[[:cntrl:][:punct:][:space:]]|[»«ª*¨¨]`, text)
	return matched
}

func RetrieveTypeLanguage(text string) string{
	if(IsEmpty(text)){
		return none
	} else if(IsLetter(text) && IsNumeric(text)){
		return letterAndNumber
	} else if(IsNumeric(text)){
		return number
	}

	return none
}