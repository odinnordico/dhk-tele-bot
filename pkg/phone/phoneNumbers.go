package phone

import (
	"strings"

	s "github.com/odinnordico/dhk-the-bot/pkg/system"
)

func ValidatePhone(n string) (string, bool) {
	number := strings.Replace(n, " ", "", -1)
	number = strings.Replace(number, "-", "", -1)
	number = strings.Replace(number, "_", "", -1)
	number = strings.Replace(number, " /", "", -1)
	if !strings.HasPrefix(n, "+") {
		cc := s.GetenvOrDefault("COUNTRY_CODE", "57")
		if !strings.HasPrefix(cc, "+") {
			cc = "+" + cc
		}
		number = cc + number
	}
	if len(n) < 9 {
		return number, false
	}
	return number, true
}
