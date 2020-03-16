package converter

import (
	"testing"
)

func TestBracketSanitizer(t *testing.T) {
	utils := NewNRCConverter()
	if utils.sanitizeBrackets("(") != `\(` {
		t.Log("it Should return SanitizedRegex for bracket (")
		t.Fail()
	}

	if utils.sanitizeBrackets(")") != `\)` {
		t.Log("it Should return SanitizedRegex for bracket ) ")
		t.Fail()
	}
}

func TestIsBurmesenumber(t *testing.T) {
	utils := NewNRCConverter()
	if utils.IsBurmeseNumber("၀") != true {
		t.Log("it Should return IsBurmeseNumber true")
		t.Fail()
	}

	if utils.IsBurmeseNumber("0") != false {
		t.Log("it Should return IsBurmeseNumber false")
		t.Fail()
	}
}

func TestIsEnglishnumber(t *testing.T) {
	utils := NewNRCConverter()
	if utils.IsEnglishNumber("၀") != false {
		t.Log("it Should return IsEnglishNumber true")
		t.Fail()
	}

	if utils.IsEnglishNumber("0") != true {
		t.Log("it Should return IsEnglishNumber false")
		t.Fail()
	}
}

func TestConvertEN2MMnumber(t *testing.T) {
	utils := NewNRCConverter()

	if utils.convertEN2MMnumber("0") != "၀" {
		t.Log("it Should return Myanmar Number ")
		t.Fail()
	}
	if utils.convertEN2MMnumber("1") != "၁" {
		t.Log("it Should return Myanmar Number ")
		t.Fail()
	}
}

func TestConvertMM2ENnumber(t *testing.T) {
	utils := NewNRCConverter()

	if utils.convertMM2ENnumber("၀") != "0" {
		t.Log("it Should return En Number ")
		t.Fail()
	}
	if utils.convertMM2ENnumber("၁") != "1" {
		t.Log("it Should return En Number ")
		t.Fail()
	}
}

func TestValidateENnrcString(t *testing.T) {
	utils := NewNRCConverter()

	if utils.ValidateEnglishNRC("12/MaGaTa(N)094186") != true {
		t.Log("it Should return true ")
		t.Fail()
	}

}

func TestValidateMMnrcString(t *testing.T) {
	utils := NewNRCConverter()

	if utils.ValidateBurmeseNRC("၁၂/မဂတ(နိုင်)၀၉၄၁၈၅") != true {
		t.Log("it Should return true ")
		t.Fail()
	}

}

func TestShouldRaiseErrorOnImproperNRCStrings(t *testing.T) {
	utils := NewNRCConverter()

	if _, err := utils.ConvertToBurmeseFormat("12/MaGaTa(Da)094186"); err == nil {
		t.Log("it Should return true ")
		t.Fail()
	}

	if _, err := utils.ConvertToEnglishFormat("၁၂/မဂသသ(နိုင်)၀၉၄၁၈၅"); err == nil {
		t.Log("it Should return true ")
		t.Fail()
	}

}

func TestConvertConvertToEnglishFormatnrc(t *testing.T) {
	utils := NewNRCConverter()

	if result, _ := utils.ConvertToEnglishFormat("၁၂/မဂတ(နိုင်)၀၉၄၁၈၅"); result != "12/MaGaTa(C)094185" {
		t.Log("it Should return citizen in english")
		t.Fail()
	}

	if result, _ := utils.ConvertToEnglishFormat("၁၂/မဂတ(ဧည့်)၀၉၄၁၈၅"); result != "12/MaGaTa(AC)094185" {
		t.Log("it Should return foreign citizen in english")
		t.Fail()
	}

	if result, _ := utils.ConvertToEnglishFormat("၁၂/မဂတ(ပြု)၀၉၄၁၈၅"); result != "12/MaGaTa(NC)094185" {
		t.Log("it Should return citizen in progress in english")
		t.Fail()
	}

	if result, _ := utils.ConvertToEnglishFormat("၁၂/မဂတ(သ)၀၉၄၁၈၅"); result != "12/MaGaTa(M)094185" {
		t.Log("it Should return monk in english")
		t.Fail()
	}

	if result, _ := utils.ConvertToEnglishFormat("၁၂/မဂတ(သီ)၀၉၄၁၈၅"); result != "12/MaGaTa(N)094185" {
		t.Log("it Should return nun NRC in english")
		t.Fail()
	}

}

func TestConvertConvertToBurmeseFormatnrc(t *testing.T) {
	utils := NewNRCConverter()

	if result, _ := utils.ConvertToBurmeseFormat("12/SaGaTa(C)194185"); result != "၁၂/သဂတ(နိုင်)၁၉၄၁၈၅" {
		t.Log("it Should return citizen in myanmar ")
		t.Fail()
	}

	if result, _ := utils.ConvertToBurmeseFormat("12/SaGaTa(AC)194185"); result != "၁၂/သဂတ(ဧည့်)၁၉၄၁၈၅" {
		t.Log("it Should return foreign citizen in myanmar")
		t.Fail()
	}

	if result, _ := utils.ConvertToBurmeseFormat("12/SaGaTa(NC)194185"); result != "၁၂/သဂတ(ပြု)၁၉၄၁၈၅" {
		t.Log("it Should return citizen in progress in myanmar")
		t.Fail()
	}

	if result, _ := utils.ConvertToBurmeseFormat("12/SaGaTa(M)194185"); result != "၁၂/သဂတ(သ)၁၉၄၁၈၅" {
		t.Log("it Should return monk nrc in myanmar ")
		t.Fail()
	}

	if result, _ := utils.ConvertToBurmeseFormat("12/SaGaTa(N)194185"); result != "၁၂/သဂတ(သီ)၁၉၄၁၈၅" {
		t.Log("it Should return Nun nrc in myanmar  ")
		t.Fail()
	}

}
