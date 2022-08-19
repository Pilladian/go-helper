package helper

import (
	"errors"
	"fmt"
	"regexp"
)

// Takes input string "data" and checks for basic JSON syntax. This check only supports simple JSON without arrays.
// Example: {0: "success", "1": "success"} would match, while {0: ["fail", "fail"]} would fail.
func ValidateSimpleJSON(data string) error {
	re, _ := regexp.Compile(` \{("(\d|\w)" *: *(\d+|"\w+") *, *)*"(\d|\w)" *: *(\d+|"\w+")\} `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", data))) {
		return errors.New("data did not match json format or contained invalid characters")
	}
	return nil
}
