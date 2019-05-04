package util

import (
	"regexp"
	"strings"
)

func GetOrderMatch(primaryKeyName string, orderString string) (string) {
	reg := regexp.MustCompile("[0-9A-Za-z.-_]+")
	matchString := reg.FindAllString(orderString, -1)
	return matchString[0] + " desc , " + primaryKeyName + " desc"
}

func ValidSQLSting(input string) string {
	reg := regexp.MustCompile("[0-9A-Za-z_.-]+")
	matchString := reg.FindAllString(input, -1)
	return matchString[0]
}

func GetOrderSpecify(inputString string) (bool) {
	return strings.Contains(strings.ToUpper(inputString), strings.ToUpper("desc")) ||
		strings.Contains(strings.ToUpper(inputString), strings.ToUpper("asc"))
}

func GetOrderString(primaryKeyName string) (string) {
	return " concat( ? , " + primaryKeyName + " ) desc"
}
