package error

import (
	"fmt"
	"os"
)

const errMsg = `{"level": %q, "msg": %q, "type": %T, "value": "%v" %s}
`

func NoLoggerFatal(i any, err error, a ...any) {
	args := argsToString(a)

	fmt.Fprintf(os.Stderr, errMsg, "fatal", err, i, i, args)
	os.Exit(1)
}

func NoLoggerError(i any, err error, a ...any) {
	args := argsToString(a)

	fmt.Fprintf(os.Stderr, errMsg, "error", err, i, i, args)
}

func argsToString(a ...any) string {
	args := ""
	la := len(a)
	for i := 0; i < len(a); i++ {
		args += fmt.Sprintf("%q", a[i])
		args += ": "
		if i+1 > la {
			args += `""`
		} else {
			i++
			args += fmt.Sprintf("%q", a[i])
		}
	}
	if len(args) > 0 {
		args = ", " + args
	}
	return args
}
