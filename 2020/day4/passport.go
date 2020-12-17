package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport is a passport entry
type Passport map[FieldType]string

// PassportValidator is an interface that knows how to validate a passport
type PassportValidator interface {
	GetValidFieldTypes() []FieldType
	GetMandatoryFieldTypes() []FieldType
	IsValid(passport Passport) bool
}

// SimplePassportValidator is a basic concrete validator
type SimplePassportValidator struct {
	FieldValidators map[FieldType]func(string) bool
}

// GetValidFieldTypes returns the valid field types for a basic validator
func (s SimplePassportValidator) GetValidFieldTypes() []FieldType {
	return []FieldType{
		BirthYear,
		IssueYear,
		ExpirationYear,
		Height,
		HairColor,
		EyeColor,
		PassportID,
		CountryID,
	}
}

// GetMandatoryFieldTypes returns the mandatory field types for a basic validator
func (s SimplePassportValidator) GetMandatoryFieldTypes() []FieldType {
	return []FieldType{
		BirthYear,
		IssueYear,
		ExpirationYear,
		Height,
		HairColor,
		EyeColor,
		PassportID,
	}
}

// NewSimplePassportValidator constructs a validator with predefined field validators
func NewSimplePassportValidator() SimplePassportValidator {
	validator := SimplePassportValidator{
		FieldValidators: make(map[FieldType]func(string) bool),
	}

	validator.FieldValidators[BirthYear] = func(s string) bool {
		result, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return result >= 1920 && result <= 2002
	}

	validator.FieldValidators[IssueYear] = func(s string) bool {
		result, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return result >= 2010 && result <= 2020
	}

	validator.FieldValidators[ExpirationYear] = func(s string) bool {
		result, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return result >= 2020 && result <= 2030
	}

	heightPattern := regexp.MustCompile("^(\\d+)(\\w+)$")

	validator.FieldValidators[Height] = func(s string) bool {
		match := heightPattern.FindStringSubmatch(s)
		if len(match) != 3 {
			return false
		}

		size, err := strconv.Atoi(match[1])
		if err != nil {
			return false
		}

		if match[2] == "cm" {
			return size >= 150 && size <= 193
		} else if match[2] == "in" {
			return size >= 59 && size <= 76
		}

		return false
	}

	colorPattern := regexp.MustCompile("^#[0-9a-fA-F]{6}$")

	validator.FieldValidators[HairColor] = func(s string) bool {
		return colorPattern.MatchString(s)
	}

	validator.FieldValidators[EyeColor] = func(s string) bool {
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range validColors {
			if color == s {
				return true
			}
		}
		return false
	}

	validator.FieldValidators[PassportID] = func(s string) bool {
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return len(s) == 9 && num > 0
	}

	return validator
}

// IsValid validates a passport
func (s SimplePassportValidator) IsValid(passport Passport) bool {
	for _, fieldType := range s.GetMandatoryFieldTypes() {
		fieldValue, ok := passport[fieldType]
		if !ok {
			fmt.Printf("Field not found: %s\n", fieldType)
			return false
		}

		if validator, ok := s.FieldValidators[fieldType]; ok {
			if !validator(fieldValue) {
				fmt.Printf("Validation failed for %s=%s\n", fieldType, fieldValue)
				return false
			}
		}
	}

	return true
}

// FieldType is a valid passport field type
type FieldType string

// These are the expected field types for a passport
const (
	BirthYear      FieldType = "byr"
	IssueYear      FieldType = "iyr"
	ExpirationYear FieldType = "eyr"
	Height         FieldType = "hgt"
	HairColor      FieldType = "hcl"
	EyeColor       FieldType = "ecl"
	PassportID     FieldType = "pid"
	CountryID      FieldType = "cid"
)

// ReadPassports reads passports from a file.
func ReadPassports(fileName string) ([]Passport, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	passports := []Passport{}
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			passportText := strings.Join(lines, " ")
			passport := ParsePassport(passportText)
			if err != nil {
				return nil, err
			}
			passports = append(passports, passport)
			lines = []string{}
			continue
		}

		lines = append(lines, text)
	}

	passportText := strings.Join(lines, " ")
	passport := ParsePassport(passportText)
	if err != nil {
		return nil, err
	}
	passports = append(passports, passport)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passports, nil
}

var fieldPattern = regexp.MustCompile("(\\w+):([^\\s]+)")

// ParsePassport parses a passport from text
func ParsePassport(input string) Passport {
	result := Passport{}
	matches := fieldPattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		result[FieldType(match[1])] = match[2]
	}
	return result
}
