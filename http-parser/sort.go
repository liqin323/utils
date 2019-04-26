package httpParser

import (
	"strconv"
	"strings"
)

type S map[string]int

func ParseSort(sort string) (s S) {

	s = make(S)

	if sort == "" {
		return
	}

	sorts := strings.Split(sort, ",")
	for _, item := range sorts {

		it := strings.Split(item, "::")

		value, _err := strconv.Atoi(it[1])
		if _err != nil {
			value = 1
		}

		name := it[0]
		if name != "" {
			s[name] = value
		}
	}

	return
}
