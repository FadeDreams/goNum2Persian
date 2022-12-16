package goNum2Persian

import (
	"testing"
)

func TestToEnglishDigits(t *testing.T) {
	tests := []struct {
		value  string
		output string
	}{
		{"۱۲۳۴۵۶۷۸۹۰", "1234567890"},
		{"١٢٣٤٥٦٧٨٩٠", "1234567890"},
		{"٠١٢٣٤٥٦٧٨٩", "0123456789"},
		{"۰۱۲۳۴۵۶۷۸۹", "0123456789"},
		{"0123456789", "0123456789"},
		{"", ""},
		{"abcdefghijklmnopqrstuvwxyz", ""},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", ""},
		{"123ABCDEFGHIJKLMNOPQRSTUVWXYZ", ""},
	}

	for _, test := range tests {
		result := ToEnglishDigits(test.value)
		if result != test.output {
			t.Errorf("digitsFaToEn(%q) = %q; want %q", test.value, result, test.output)
		}
	}
}

func TestNum2Persian(t *testing.T) {
	tests := []struct {
		input     interface{}
		level     *int
		isOrdinal bool
		expected  string
	}{
		{123, nil, false, "یکصد و بیست و سه"},
		{-456, nil, false, "منفی چهارصد و پنجاه و شش"},
		{"12", nil, false, "دوازده"},
		{"14", nil, true, "چهاردهم"},
	}

	for _, test := range tests {
		level := 0
		result := Num2Persian(test.input, &level, test.isOrdinal)
		if result != test.expected {
			t.Errorf("Num2Persian(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
