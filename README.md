# goNum2Persian
This tool will convert Persian number words into their numerical counterparts in Golang.
## Installation
```
go get github.com/fadedreams/goNum2Persian
```

## Example usage

```
package main
import (
	"fmt"
	"github.com/fadedreams/goNum2Persian"
)

func main() {
	fmt.Println(goNum2Persian.Num2Persian("5678", nil)) //یکصد و یازده
	fmt.Println(goNum2Persian.Num2Persian("٥٦٧٨", nil)) //پنج هزار و ششصد و هفتاد و هشت
	//set IsOrdinal to true
	fmt.Println(goNum2Persian.Num2Persian("٥٦٧٨", nil, true)) //پنج هزار و ششصد و هفتاد و هشتم

	//converts Persian/Arabic to English Digits
	fmt.Println(goNum2Persian.ToEnglishDigits("۴۵۶"))
	fmt.Println(goNum2Persian.ToEnglishDigits("٥٦٧٨"))
}

```

