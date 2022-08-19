package helper

import "fmt"

// Takes two strings "expectation" and "actual" as input and returns the pretty-print version.
func Prettify(expectation string, actual string) string {
	return fmt.Sprintf("\n\t[ WANTED ] %s\n\t[ ACTUAL ] %s\n\n", expectation, actual)
}
