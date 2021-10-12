package main

import (
	"fmt"
	"github.com/ezysign/go-mm-nrc-converter/converter"
)

func main() {
	nrcConverter := converter.NewNRCConverter()

	fmt.Println(nrcConverter.ValidateEnglishNRC("12/MaGaTa(နိုင်)198475"))
	fmt.Println(nrcConverter.ValidateBurmeseNRC("၁၂/မစလ(နိုင်)၀၈၇၆၄၅"))
	fmt.Println(nrcConverter.ValidateEnglishNRC("12/UKaMa(C)123456"))

	fmt.Println(nrcConverter.ConvertToEnglishFormat("၁၂/မစသ(သ)၀၈၇၆၄၅"))
	fmt.Println(nrcConverter.ConvertToBurmeseFormat("12/MaCaSa(M)087645"))

	fmt.Println(nrcConverter.ConvertToBurmeseFormat("12/UKaMa(C)123456"))
	fmt.Println(nrcConverter.ConvertToBurmeseFormat("12/AuKaMa(C)123456"))
	fmt.Println(nrcConverter.ConvertToEnglishFormat("၁၂/ဥကမ(နိုင်)၁၂၃၄၅၆"))

}
