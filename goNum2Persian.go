package goNum2Persian

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ToEnglishDigits(value string) string {
	var arabicNumbers []string = []string{"١", "٢", "٣", "٤", "٥", "٦", "٧", "٨", "٩", "٠"}
	var persianNumbers []string = []string{"۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹", "۰"}
	var englishNumbers []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	if value == "" {
		return ""
	}

	for _, c := range value {
		if !unicode.IsDigit(c) {
			return ""
		}
	}

	for _, englishNumber := range englishNumbers {
		if strings.Contains(value, englishNumber) {
			for i := 0; i < len(englishNumbers); i++ {
				re := regexp.MustCompile(arabicNumbers[i])
				value = re.ReplaceAllString(value, englishNumbers[i])
			}
			return value
		}
	}

	for _, arabicNumber := range arabicNumbers {
		if strings.Contains(value, arabicNumber) {
			for i := 0; i < len(englishNumbers); i++ {
				re := regexp.MustCompile(arabicNumbers[i])
				value = re.ReplaceAllString(value, englishNumbers[i])
			}
			return value
		}
	}

	for _, persianNumber := range persianNumbers {
		if strings.Contains(value, persianNumber) {
			for i := 0; i < len(englishNumbers); i++ {
				re := regexp.MustCompile(persianNumbers[i])
				value = re.ReplaceAllString(value, englishNumbers[i])
			}
			return value
		}
	}
	return ""

}

func Num2Persian(input interface{}, level *int, isOrdinal ...bool) string {

	if level == nil {
		zero := 0
		level = &zero
	}

	if input == "" {
		return ""
	}

	var num int
	switch input := input.(type) {
	case int:
		num = input
	case string:
		numStr1 := ToEnglishDigits(input)
		if num2, err := strconv.Atoi(numStr1); err != nil {
			panic(err)
		}else{
			num = num2
		}
	}

	if num > 999999999999999 {
		panic("Error: number is out of range")
	}

	if num == 0 {
		if *level == 0 {
			return "صفر"
		} else {
			return ""
		}
	}

	if num < 0 {
		num = -num
		return "منفی " + Num2Persian(strconv.Itoa(num), level)
	}

	result := ""
	yekan := []string{"یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"}
	dahgan := []string{"بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"}
	sadgan := []string{"یکصد", "دویست", "سیصد", "چهارصد", "پانصد", "ششصد", "هفتصد", "هشتصد", "نهصد"}
	dah := []string{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هیجده", "نوزده"}

	if *level > 0 {
		result += " و "
		*level -= 1
	}

	tmp1 := int(*level) + 1
	if num < 10 {
		result += yekan[num-1]
	} else if num < 20 {
		result += dah[num-10]
	} else if num < 100 {
		result += dahgan[(num/10)-2] + Num2Persian(num%10, &tmp1)
	} else if num < 1000 {
		result += sadgan[(num/100)-1] + Num2Persian(num%100, &tmp1)
	} else if num < 1000000 {
		result += Num2Persian(num/1000, level) + " هزار" + Num2Persian(num%1000, &tmp1)
	} else if num < 1000000000 {
		result += Num2Persian(num/1000000, level) + " میلیون" + Num2Persian(num%1000000, &tmp1)
	} else if num < 1000000000000 {
		result += Num2Persian(num/1000000000, level) + " میلیارد" + Num2Persian(num%1000000000, &tmp1)
	} else if num < 1000000000000000 {
		result += Num2Persian(num/1000000000000, level) + " تریلیارد" + Num2Persian(num%1000000000000, &tmp1)
	}

	if len(isOrdinal) > 0 && isOrdinal[0] == true {
		return func() string {
			result, err := addOrdinalSuffix(result)
			if err != nil {
				return ""
			}
			return result
		}()
		//return result + addOrdinalSuffix
	}

	return result
}

func addOrdinalSuffix(number string) (string, error) {
	// check if the input is a string
	if number == "" {
		return "", errors.New("PersianTools: addOrdinalSuffix - The input must be string")
	}

	// check if the number ends with "ی"
	if strings.HasSuffix(number, "ی") {
		return number + " اُم", nil
	}

	// check if the number ends with "سه"
	if strings.HasSuffix(number, "سه") {
		return number[:len(number)-2] + "سوم", nil
	}

	// return the number with the suffix "م"
	return number + "م", nil
}

