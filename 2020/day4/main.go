package main

import (
	"fmt"
	"goaoc/utils"
	"strconv"
	"strings"
	"unicode"
)

var reqFields map[string]bool

func init() {
	reqFields = map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
}

func allFields(currFields map[string]bool) bool {
	for k := range reqFields {
		if _, ex := currFields[k]; !ex {
			return false
		}
	}
	return true
}

func allFields2(currFields map[string]string) bool {
	for k := range reqFields {
		if _, ex := currFields[k]; !ex {
			return false
		}
	}
	return true
}

func part1() {
	lines := utils.ReadLinesAsString("input.txt")

	currFields := map[string]bool{}
	nValid := 0
	for _, line := range lines {
		if line == "" {
			if allFields(currFields) {
				nValid++
			}
			currFields = map[string]bool{}
		} else {
			parts := strings.Split(line, " ")
			for _, p := range parts {
				kv := strings.Split(p, ":")
				currFields[kv[0]] = true
			}
		}
	}
	if allFields(currFields) {
		nValid++
	}
	fmt.Println(nValid)
}

func doesFollowRules(rules map[string]func(string) bool, passport map[string]string) bool {
	for k, v := range passport {
		if _, ex := rules[k]; !ex {
			if k == "cid" {
				continue
			}
			return false
		}
		validator := rules[k]
		if !validator(v) {
			return false
		}
	}
	return true
}

func part2() {
	lines := utils.ReadLinesAsString("input.txt")
	fieldRules := map[string]func(str string) bool{
		"byr": func(str string) bool {
			val, _ := strconv.Atoi(str)
			return val >= 1920 && val <= 2002
		},
		"iyr": func(str string) bool {
			val, _ := strconv.Atoi(str)
			return val >= 2010 && val <= 2020
		},
		"eyr": func(str string) bool {
			val, _ := strconv.Atoi(str)
			return val >= 2020 && val <= 2030
		},
		"hgt": func(str string) bool {
			val, err := strconv.Atoi(str[:len(str)-2])
			if err != nil {
				return false
			}
			unit := str[len(str)-2:]
			if unit == "cm" {
				return val >= 150 && val <= 193
			} else if unit == "in" {
				return val >= 59 && val <= 76
			}
			return false
		},
		"hcl": func(str string) bool {
			if str[0] != '#' {
				return false
			}
			if len(str[1:]) != 6 {
				return false
			}
			for _, r := range str[1:] {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					return false
				}
			}
			return true
		},
		"ecl": func(str string) bool {
			eyeColors := map[string]bool{
				"amb": true,
				"blu": true,
				"brn": true,
				"gry": true,
				"grn": true,
				"hzl": true,
				"oth": true,
			}
			if _, ex := eyeColors[str]; !ex {
				return false
			}
			return true
		},
		"pid": func(str string) bool {
			if len(str) != 9 {
				return false
			}
			for _, r := range str {
				if r < '0' || r > '9' {
					return false
				}
			}
			return true
		},
	}
	currFields := map[string]string{}
	nValid := 0
	for _, line := range lines {
		if line == "" {
			if allFields2(currFields) && doesFollowRules(fieldRules, currFields) {
				nValid++
			}
			currFields = make(map[string]string)
		} else {
			parts := strings.Split(line, " ")
			for _, p := range parts {
				kv := strings.Split(p, ":")
				currFields[kv[0]] = kv[1]
			}
		}
	}
	if allFields2(currFields) && doesFollowRules(fieldRules, currFields) {
		nValid++
	}
	fmt.Println(nValid)
}

func main() {
	part1()
	part2()
}
