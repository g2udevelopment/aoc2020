package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

func main() {
	lines, _ := util.ReadLines("./input-day4")
	var passports []Passport
	var buffer string
	for _, line := range lines {
		if line == "" {
			passport := parsePassport(strings.TrimLeft(buffer, " "))
			passports = append(passports, passport)
			buffer = ""
		} else {

			buffer = buffer + " " + line
		}

	}
	passport := parsePassport(strings.TrimLeft(buffer, " ")) //Last one
	passports = append(passports, passport)

	var validCount int
	for _, passport := range passports {
		//fmt.Printf("passport: %v \n", passport)
		if validPassport(passport) {
			validCount++
		}
	}
	fmt.Printf("Number of valid passwords: %v", validCount)
}

func validPassport(p Passport) bool {
	attribs := []struct {
		key        string
		validation func(string) bool
	}{{"byr", isValidByr}, {"iyr", isValidIyr}, {"eyr", isValidEyr}, {"hgt", IsValidHeight}, {"hcl", IsValidHC}, {"ecl", IsValidEye}, {"pid", IsValidPID}, {"cid", IsValidCID}}
	for _, att := range attribs {
		val, _ := p[att.key]
		if !att.validation(val) {
			return false
		}
	}
	return true
}

func parsePassport(p string) Passport {

	var buffer string
	var passport Passport = make(Passport)
	//fmt.Printf("buffer: %v \n", p)
	for _, char := range p {
		if string(char) == " " {
			attrib := strings.Split(buffer, ":")
			passport[attrib[0]] = attrib[1]
			buffer = ""
		} else {
			buffer = buffer + string(char)

		}
	}
	attrib := strings.Split(buffer, ":") //Last one
	passport[attrib[0]] = attrib[1]
	return passport
}

///////Validation functions
func isValidNotEmpty(value string) bool {
	return len(value) > 0
}

func isValidByr(value string) bool {

	birth, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return birth >= 1920 && birth <= 2002
}

func isValidIyr(value string) bool {

	issue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return issue >= 2010 && issue <= 2020
}

func isValidEyr(value string) bool {

	exp, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return exp >= 2020 && exp <= 2030
}

func IsValidHeight(value string) bool {

	r, _ := regexp.Compile("^([0-9]{2,3})(in|cm)$")
	m := r.FindStringSubmatch(value)

	if len(m) == 0 {
		return false
	}
	unit := m[2]
	num, err := strconv.Atoi(m[1])
	if err != nil {
		return false
	}

	if unit == "cm" {
		return num >= 150 && num <= 193
	} else if unit == "in" {
		return num >= 59 && num <= 76
	}
	return false
}

func IsValidHC(value string) bool {

	r, _ := regexp.Compile("^#[0-9a-f]{6}$")
	return r.MatchString(value)
}

func IsValidEye(value string) bool {

	r, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	return r.MatchString(value)
}

func IsValidPID(value string) bool {

	r, _ := regexp.Compile("^[0-9]{9}$")
	return r.MatchString(value)
}

func IsValidCID(value string) bool {
	return true
}
