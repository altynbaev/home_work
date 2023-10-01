package hw10programoptimization

import (
	"bufio"
	"io"
	"regexp"
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

	regExp, err := regexp.Compile("\\." + domain)
	if err != nil {
		return result, err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &user); err != nil {
			return nil, err
		}

		matched := regExp.MatchString(user.Email)
		if err != nil {
			return nil, err
		}
		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
