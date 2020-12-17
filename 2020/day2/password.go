package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Password represents a password entry.
type Password struct {
	TimesAtLeast       int
	TimesAtMost        int
	VerificationLetter rune
	Password           string
}

var passwordPattern = regexp.MustCompile("^(\\d+)-(\\d+)\\s?(\\w{1}):\\s?(\\w+)$")

// ReadPasswords reads Password entries from a file.
func ReadPasswords(fileName string) ([]Password, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passwords := []Password{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := passwordPattern.FindAllStringSubmatch(line, -1)
		if len(matches) > 1 {
			return nil, fmt.Errorf("line does not match expected pattern: %s", line)
		}

		timesAtLeast, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return nil, fmt.Errorf("cannot convert to integer: %s, line: %s", matches[0][1], line)
		}

		timesAtMost, err := strconv.Atoi(matches[0][2])
		if err != nil {
			return nil, fmt.Errorf("cannot convert to integer: %s, line: %s", matches[0][2], line)
		}

		passwords = append(passwords, Password{
			TimesAtLeast:       timesAtLeast,
			TimesAtMost:        timesAtMost,
			VerificationLetter: []rune(matches[0][3])[0],
			Password:           matches[0][4],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
}

// IsValid validates a password according to the old rules
func (p Password) IsValid() bool {
	var count int
	for _, i := range []rune(p.Password) {
		if i == p.VerificationLetter {
			count++
		}
	}
	return count >= p.TimesAtLeast && count <= p.TimesAtMost
}

// IsValidNew validates a password according to the new rules
func (p Password) IsValidNew() bool {
	var count int
	runes := []rune(p.Password)

	if runes[p.TimesAtLeast-1] == p.VerificationLetter {
		count++
	}
	if runes[p.TimesAtMost-1] == p.VerificationLetter {
		count++
	}

	return count == 1
}
