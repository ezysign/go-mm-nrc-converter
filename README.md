# Go-mm-nrc-Converter

[![Go Report Card](https://goreportcard.com/badge/github.com/ezysign/go-mm-nrc-converter)](https://goreportcard.com/report/github.com/ezysign/go-mm-nrc-converter) [![Build Status](https://travis-ci.org/ezysign/go-mm-nrc-converter.svg?branch=master)](https://travis-ci.org/ezysign/go-mm-nrc-converter) [![GitHub issues](https://img.shields.io/github/issues/ezysign/go-mm-nrc-converter)](https://github.com/ezysign/go-mm-nrc-converter/issues) [![GitHub forks](https://img.shields.io/github/forks/ezysign/go-mm-nrc-converter)](https://github.com/ezysign/go-mm-nrc-converter/network) [![GitHub stars](https://img.shields.io/github/stars/ezysign/go-mm-nrc-converter)](https://github.com/ezysign/go-mm-nrc-converter/stargazers) [![GitHub license](https://img.shields.io/github/license/ezysign/go-mm-nrc-converter)](https://github.com/ezysign/go-mm-nrc-converter/blob/master/LICENSE.md)

Go-mm-nrc-converter is a Burmese National Registration Card(NRC) tool written in Golang. Inspired by [NRC-Prefixer](https://github.com/greenlikeorange/NRCPrefix) and[Burther](https://github.com/vincent-paing/Burtha) for android

```go
	nrcConverter := converter.NewNRCConverter()
	nrcConverter.ConvertToEnglishFormat("၁၂/မစသ(သ)၀၈၇၆၄၅"))
	nrcConverter.ConvertToBurmeseFormat("12/MaCaSa(M)087645")

//Output: 12/MaCaSa(M)087645
//Output: ၁၂/မစသ(သ)၀၈၇၆၄၅

```

go-mm-nrc-converter supports all NRC types with the following pattern

| Burmese                         | English                   |
| ------------------------------- | ------------------------- |
| နိုင်ငံသား (နိုင်)              | Citizen (C)               |
| ဧည့်နိုင်ငံသား (ဧည့်)           | Associate Citizen (AC)    |
| နိုင်ငံသားပြုခွင့်ရသူ (ပြု)     | Naturalized Citizen (NC)  |
| နိုင်ငံသားပြုစီစစ်ခံရမည့်သူ (စ) | National Verification (V) |
| သာသနာဝင် (သ)                    | Monk (M)                  |
| သာသနာနွယ်ဝင် (သီ)               | Nun (N)                   |

## Getting Started

### Add Dependency

```bash
go get github.com/ezysign/go-mm-nrc-converter/converter
```

## Features

### Validation

go-mm-nrc-converter provide a validation of NRCNumbers in both English & Burmese.

```go
package main

import (
	"fmt"
	"github.com/ezysign/go-mm-nrc-converter/converter"
)

func  main() {
	nrcConverter := converter.NewNRCConverter()
	nrcConverter.ValidateEnglishNRC("12/MaGaTa(နိုင်)198475"))
	nrcConverter.ValidateBurmeseNRC("၁၂/မစလ(နိုင်)၀၈၇၆၄၅"))
}
//Output: false
//Output: true
```

### Conversion

go-mm-nrc-converter provide a conversion of NRCNumbers in both English & Burmese.

```go
package main

import (
	"fmt"
	"github.com/ezysign/go-mm-nrc-converter/converter"
)

func  main() {
	nrcConverter := converter.NewNRCConverter()
	nrcConverter.ConvertToEnglishFormat("၁၂/မစသ(သ)၀၈၇၆၄၅"))
	nrcConverter.ConvertToBurmeseFormat("12/MaCaSa(M)087645")
}
//Output: 12/MaCaSa(M)087645
//Output: ၁၂/မစသ(သ)၀၈၇၆၄၅
```

### Playground

https://play.golang.org/p/DxK_WotOrUy

## Contributing

1. Fork it (https://github.com/ezysign/go-mm-nrc-converter)

2) Create your feature branch (`git checkout -b my-new-feature`)

3. Commit your changes (`git commit -am 'Add some feature'`)

4) Push to the branch (`git push origin my-new-feature`)

5. Create a new Pull Request

## Licenese

```
Copyright 2020 EzySign

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
