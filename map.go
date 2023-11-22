package interpolate

import (
	"fmt"
	"regexp"
	"strings"
)

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
