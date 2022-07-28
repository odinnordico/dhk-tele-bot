package system

import (
	"os"

	e "github.com/odinnordico/dhk-the-bot/pkg/error"
)

func Must[T any](t T, err error) T {
	if err != nil {
		e.NoLoggerFatal(t, err)
	}
	return t
}

func GetEnvOrDefault(k, d string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	os.Setenv(k, d)
	return d
}
