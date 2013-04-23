package misc

import (
	"strconv"
	"strings"
)

var specialForms = map[string]int{
	"dev":   -6,
	"alpha": -5,
	"a":     -5,
	"beta":  -4,
	"b":     -4,
	"RC":    -3,
	"rc":    -3,
	"#":     -2,
	"p":     1,
	"pl":    1,
}

func prepVersion(version string) []string {
	if len(version) == 0 {
		return []string{""}
	}

	version = RegReplace(`[_\-+]`, ".", version)
	version = RegReplace(`([^.\d]+)`, ".$1.", version)
	version = RegReplace(`\.{2,}`, ".", version)

	return strings.Split(version, ".")
}

func numVersion(value string) int {
	if value == "" {
		return 0
	}

	if number, err := strconv.Atoi(value); err == nil {
		return number
	}

	if special, ok := specialForms[value]; ok {
		return special
	}

	return -7
}

func CompareVersionSimple(version1, version2 string) int {
	compare := 0
	x := 0
	r := 0
	l := 0

	v1 := prepVersion(version1)
	v2 := prepVersion(version2)
	len1 := len(v1)
	len2 := len(v2)

	if len1 > len2 {
		x = len1
	} else {
		x = len2
	}

	for i := 0; i < x; i++ {
		if i < len1 && i < len2 {
			if v1[i] == v2[i] {
				continue
			}
		}

		r = 0
		if i < len1 {
			r = numVersion(v1[i])
		}

		l = 0
		if i < len2 {
			l = numVersion(v2[i])
		}

		if r < l {
			compare = -1
			break
		} else if r > l {
			compare = 1
			break
		}
	}

	return compare
}

func CompareVersion(version1, version2, operator string) bool {
	compare := CompareVersionSimple(version1, version2)

	switch {
	case operator == ">" || operator == "gt":
		if compare > 0 {
			return true
		} else {
			return false
		}
	case operator == ">=" || operator == "ge":
		if compare >= 0 {
			return true
		} else {
			return false
		}
	case operator == "<=" || operator == "le":
		if compare <= 0 {
			return true
		} else {
			return false
		}
	case operator == "==" || operator == "=" || operator == "eq":
		if compare == 0 {
			return true
		} else {
			return false
		}
	case operator == "<>" || operator == "!=" || operator == "ne":
		if compare != 0 {
			return true
		} else {
			return false
		}
	case operator == "" || operator == "<" || operator == "lt":
		if compare < 0 {
			return true
		} else {
			return false
		}
	}

	return false
}
