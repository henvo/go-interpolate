package interpolate

import (
	"fmt"
	"regexp"
	"strings"
)

// FromMap will perform string interpolation on a given string
// and map with string keys and interface{} value.
// In order to be able to interpolate variables with values
// of the given map the keys to the values need to be wrapped
// with a percentage sign and curly bracketts `%{}`
func FromMap(s string, m map[string]interface{}) string {
	re := regexp.MustCompile(`%{(.*?)\}`)

	for _, sm := range re.FindAllStringSubmatch(s, -1) {
		val, ok := m[sm[1]]

		if !ok {
			val = ""
		}

		s = strings.Replace(s, sm[0], fmt.Sprintf("%s", val), -1)
	}

	return s
}
