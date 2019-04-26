package httpParser

import "strings"

type F map[string]string

func ParseFilter(filter string) (f F) {

	f = make(F)

	if filter == "" {
		return
	}

	filterSet := strings.Split(filter, ",")
	for _, v := range filterSet {
		item := strings.Split(v, "::")
		if item[0] != "" && item[1] != "" {
			f[item[0]] = item[1]
		}
	}

	return
}
