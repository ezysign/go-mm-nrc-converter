package converter

import (
	"errors"
	"regexp"
	"strings"
)

type nrcconverter interface {
	isBurmeseNumber(number string) bool
	isEnglishNumber(number string) bool
	sanitizeBrackets(bracket string) string
	convertEN2MMnumber(enNumber string) string
	convertMM2ENnumber(mmNumber string) string
	ValidateBurmeseNRC(nrcNumber string) bool
	ValidateEnglishNRC(nrcNumber string) bool
	ConvertToEnglishFormat(nrcNumber string) (string, error)
	ConvertToBurmeseFormat(nrcNumber string) (string, error)
}

type converter struct{}

func NewNRCConverter() nrcconverter {
	return &converter{}
}

func (converter *converter) isBurmeseNumber(number string) bool {
	mmNumber := "၀၁၂၃၄၅၆၇၈၉"
	return strings.Contains(mmNumber, number)

}

func (converter *converter) sanitizeBrackets(bracket string) string {
	if bracket == "(" {
		return `\(`
	}

	if bracket == ")" {
		return `\)`
	}
	return bracket
}

func (converter *converter) isEnglishNumber(number string) bool {
	mmNumber := "0123456789"
	return strings.Contains(mmNumber, number)

}

func (converter *converter) convertEN2MMnumber(number string) string {
	char := []rune(number)
	ascii := int(char[0])
	if converter.isEnglishNumber(number) == false {
		return number
	}
	return string(ascii + 4112)
}

func (converter *converter) convertMM2ENnumber(number string) string {
	char := []rune(number)
	ascii := int(char[0])
	if converter.isBurmeseNumber(number) == false {
		return number
	}
	return string(ascii - 4112)
}

func (converter *converter) ValidateEnglishNRC(nrcNumber string) bool {
	var re = regexp.MustCompile(enNRCformatRegex)
	return len(re.FindStringIndex(nrcNumber)) > 0

}

func (converter *converter) ValidateBurmeseNRC(nrcNumber string) bool {
	var re = regexp.MustCompile(mmNRCformatRegex)
	return len(re.FindStringIndex(nrcNumber)) > 0
}

func (converter *converter) ConvertToEnglishFormat(nrcNumber string) (string, error) {

	if converter.ValidateBurmeseNRC(nrcNumber) == false {
		return "", errors.New("Invalid Myanmar NRC Number")
	}
	re2 := regexp.MustCompile("")

	for typeKey, typeVal := range nrcmm2en {
		re2 = regexp.MustCompile(typeKey)
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(typeVal)))
	}

	for charKey, charVal := range mm2enChar {
		re2 = regexp.MustCompile(charKey)
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(charVal)))
	}

	for _, num := range nrcNumber {
		re2 = regexp.MustCompile(converter.sanitizeBrackets(string(int(num))))
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(converter.convertMM2ENnumber(string(int(num))))))
	}

	return nrcNumber, nil
}

func (converter *converter) ConvertToBurmeseFormat(nrcNumber string) (string, error) {
	if converter.ValidateEnglishNRC(nrcNumber) == false {
		return "", errors.New("Invalid English NRC Number")
	}
	re2 := regexp.MustCompile("")
	for _, num := range nrcNumber {

		re2 = regexp.MustCompile(converter.sanitizeBrackets(string(int(num))))
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(converter.convertEN2MMnumber(string(int(num))))))
	}

	for charKey, charVal := range en2mmChar {
		re2 = regexp.MustCompile(charKey)
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(charVal)))
	}

	for typeKey, typeVal := range nrcen2MM {
		re2 = regexp.MustCompile(typeKey)
		nrcNumber = string(re2.ReplaceAll([]byte(nrcNumber), []byte(typeVal)))
	}

	return nrcNumber, nil
}
