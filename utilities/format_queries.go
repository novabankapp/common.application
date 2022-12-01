package utilities

import (
	"fmt"
	noSql "github.com/novabankapp/common.data/repositories/base/cassandra"
	"regexp"
	"strings"
)

func MakeQueries(queries []map[string]string, field, compare, value string) []map[string]string {

	m := make(map[string]string)
	m[noSql.COLUMN] = field
	m[noSql.COMPARE] = compare
	m[noSql.VALUE] = value
	queries = append(queries, m)
	return queries
}
func FormatPhonePasswordResetMessage(pin string, expiryDate string) string {
	return fmt.Sprintf("Your Password reset pin is %s and will expire after %s", pin, expiryDate)
}
func FormatPhoneLoginMessage(pin string, expiryDate string) string {
	return fmt.Sprintf("Your Password login pin is %s and will expire after %s", pin, expiryDate)
}
func FormatEmailPasswordResetMessage(hash string, expiryDate string) string {
	return fmt.Sprintf("Your Password reset pin is %s and will expire after %s", hash, expiryDate)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
