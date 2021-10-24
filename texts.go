package texts

import (
	"strconv"
	"time"
)

func sensor(v interface{}) ([]rune, string) {
	var output []rune
	var types string
	switch t := v.(type) {
	case int:
		types = "int"
		str := strconv.Itoa(t)
		output = []rune(str)
	case string:
		types = "string"
		output = []rune(t)
	case float64:
		types = "float"
		str := strconv.FormatFloat(t, 'f', -1, 64)
		output = []rune(str)
	case bool:
		types = "bool"
		str := strconv.FormatBool(t)
		output = []rune(str)
	}

	return output, types
}

// All types of values can be entered.
// Print the entered values one letter at a time.
func Text(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 35)
	}
	println()
}

// Like the existing Text() function,
// You can output the entered values one by one and enter all types of values.
// But Not change the line.
func NotLineText(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 35)
	}
}

// Print characters a little slow than the existing Text() function.
func SlowText(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 70)
	}
	println()
}

// prints letters very slow.
func SoSlowText(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 200)
	}
	println()
}

// Print characters a little faster than the existing Text() function.
func FastText(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 20)
	}
	println()
}

// prints letters very quickly.
func SoFastText(v interface{}) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * 1)
	}
	println()
}

// Set it at the speed the user wants.
// Calculate the speed in the form of millisecond.
func WhatTimeText(v interface{}, timer int) {
	s, _ := sensor(v)
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
		time.Sleep(time.Millisecond * time.Duration(timer))
	}
	println()
}

// It's the same as the Text() function, but you can receive multiple values at once.
func Texts(v ...interface{}) {
	for i := 0; i < len(v); i++ {
		s, _ := sensor(v[i])
		for i := 0; i < len(s); i++ {
			print(string(s[i]))
			time.Sleep(time.Millisecond * 35)
		}
	}
	println()
}

func OptionTexts(v []interface{}, timer int, check string) {
	if check == "on" {
		for i := 0; i < len(v); i++ {
			s, _ := sensor(v[i])
			for i := 0; i < len(s); i++ {
				print(string(s[i]))
				time.Sleep(time.Millisecond * time.Duration(timer))
			}
		}
		println()
	} else if check == "off" {
		for i := 0; i < len(v); i++ {
			s, _ := sensor(v[i])
			for i := 0; i < len(s); i++ {
				print(string(s[i]))
				time.Sleep(time.Millisecond * time.Duration(timer))
			}
		}
	} else {
		println("Incorrect input.")
	}
}
