package interpolate

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func FromURLValues(s string, values url.Values) string {
	re := regexp.MustCompile(`%{(.*?)\}`)

	for _, sm := range re.FindAllStringSubmatch(s, -1) {
		val, ok := values[sm[1]]

		if !ok {
			val = []string{""}
		}

		s = strings.Replace(s, sm[0], fmt.Sprintf("%s", val[0]), -1)
	}

	return s
}
