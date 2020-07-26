package tdd

import (
	"github.com/stretchr/testify/assert"
	"go-search-history/inspect"
	"testing"
)

func TestInspect(t *testing.T) {
	assert.Equal(t, inspect.Inquiry{}, inspect.Inspect(""))

	assert.Equal(t, inspect.Inquiry{
		"1",
		true,
		false,
		false,
		"number",
		1}, inspect.Inspect("1"))

	assert.Equal(t, inspect.Inquiry{
		"123",
		true,
		false,
		false,
		"number",
		3}, inspect.Inspect("123"))

	assert.Equal(t, inspect.Inquiry{
		"",
		false,
		false,
		false,
		"",
		0}, inspect.Inspect(""))

	assert.Equal(t, inspect.Inquiry{
		"abc",
		false,
		true,
		false,
		"letters",
		3}, inspect.Inspect("abc"))
}

func TestIsLetter(t *testing.T) {
	assert.Equal(t, true, inspect.IsLetter("asd"))
	assert.Equal(t, false, inspect.IsLetter(""))
	assert.Equal(t, false, inspect.IsLetter("1"))
	assert.Equal(t, false, inspect.IsLetter(" a "))
	assert.Equal(t, false, inspect.IsLetter(" abc1 "))
	assert.Equal(t, false, inspect.IsLetter(" &/()=?»|!_:;ª^`* "))
}

func TestContainsCharacterAndNumber(t *testing.T) {
	assert.Equal(t, true, inspect.ContainsCharacterOrNumber("a1"))
	assert.Equal(t, false, inspect.ContainsCharacterOrNumber(""))
	assert.Equal(t, true, inspect.ContainsCharacterOrNumber("1"))
	assert.Equal(t, true, inspect.ContainsCharacterOrNumber("abc"))
	assert.Equal(t, true, inspect.ContainsCharacterOrNumber(" abc "))
	assert.Equal(t, true, inspect.ContainsCharacterOrNumber(" Rei 1"))
	assert.Equal(t, false, inspect.ContainsCharacterOrNumber(" &/()=?»|!_:;ª^`* "))
}

func TestIsNumeric(t *testing.T) {
	assert.Equal(t, false, inspect.IsNumeric("a1"))
	assert.Equal(t, false, inspect.IsNumeric(""))
	assert.Equal(t, true, inspect.IsNumeric("1"))
	assert.Equal(t, false, inspect.IsNumeric("abc"))
	assert.Equal(t, false, inspect.IsNumeric(" abc "))
	assert.Equal(t, false, inspect.IsNumeric(" Rei 1"))
	assert.Equal(t, false, inspect.IsNumeric(" &/()=?»|!_:;ª^`* "))
}

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, inspect.IsEmpty(""))
	assert.Equal(t, false, inspect.IsEmpty("a"))
	assert.Equal(t, false, inspect.IsEmpty(" "))
}

func TestIsSpecialCharacter(t *testing.T) {
	assert.Equal(t, false, inspect.IsSpecialCharacter(""))
	assert.Equal(t, false, inspect.IsSpecialCharacter("a"))
	assert.Equal(t, true, inspect.IsSpecialCharacter(" "))
	assert.Equal(t, true, inspect.IsSpecialCharacter("/()=?»|!"))
}

func TestRetrieveTypeLanguage(t *testing.T) {
	assert.Equal(t, "none", inspect.RetrieveTypeLanguage(""))
	assert.Equal(t, "letters", inspect.RetrieveTypeLanguage("abcde"))
	assert.Equal(t, "numbers", inspect.RetrieveTypeLanguage("123"))
}