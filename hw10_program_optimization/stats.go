package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	var user User

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &user); err != nil {
			return nil, err
		}

		matched := strings.HasSuffix(user.Email, domain)
		if matched {
			emailSlice := strings.SplitN(user.Email, "@", 2)
			if len(emailSlice) != 2 {
				return nil, fmt.Errorf("invalid email: %q", user.Email)
			}
			resultString := strings.ToLower(emailSlice[1])
			result[resultString]++
		}
	}
	return result, nil
}
