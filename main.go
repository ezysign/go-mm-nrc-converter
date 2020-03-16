package main

import (
	"github.com/ezysign/go-mm-nrc-converter/converter"
)

func main() {
	nrcConverter := converter.NewNRCConverter()
	nrcConverter.IsBurmeseNumber("၁၂/မစလ(နိုင်)၀၈၇၆၄၅")
	nrcConverter.IsEnglishNumber("12/MaGaTa(နိုင်)198475")
	nrcConverter.ValidateEnglishNRC("12/MaGaTa(နိုင်)198475")
	nrcConverter.ValidateBurmeseNRC("၁၂/မစလ(နိုင်)၀၈၇၆၄၅")
	nrcConverter.ConvertToBurmeseFormat("12/MaGaTa(နိုင်)198475")
	nrcConverter.ConvertToEnglishFormat("၁၂/မစလ(နိုင်)၀၈၇၆၄၅")
}
