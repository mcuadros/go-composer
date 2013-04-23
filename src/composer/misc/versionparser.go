package misc

import (
	"regexp"
	"strconv"
	"strings"
)

var modifierRegex = `[._-]?(?:(stable|beta|b|RC|alpha|a|patch|pl|p)(?:[.-]?(\d+))?)?([.-]?dev)?`

func Normalize(version string) string {

	// ignore aliases and just assume the alias is required instead of the source
	result := RegFind(`^([^,\s]+) +as +([^,\s]+)$`, version)
	if result != nil {
		version = result[1]
	}

	// match master-like branches
	r := regexp.MustCompile(`^(?:dev-)?(?:master|trunk|default)$`)
	if r.MatchString(strings.ToLower(version)) {
		return "9999999-dev"
	}

	if strings.HasPrefix(strings.ToLower(version), "dev-") {
		return "dev-" + version[4:len(version)]
	}

	index := 0

	// match classical versioning
	result = RegFind(`(?i)^v?(\d{1,3})(\.\d+)?(\.\d+)?(\.\d+)?`+modifierRegex+`$`, version)
	if result != nil {
		version = ""
		for _, val := range result[1:5] {
			if val != "" {
				version = version + val
			} else {
				version = version + ".0"
			}
		}

		index = 5
	} else {
		// match date-based versioning
		result = RegFind(`(?i)^v?(\d{4}(?:[.:-]?\d{2}){1,6}(?:[.:-]?\d{1,3})?)`+modifierRegex+`$`, version)
		if result != nil {
			version = regexp.MustCompile(`\D`).ReplaceAllString(result[1], "-")
			index = 2
		}
	}

	if index != 0 {
		if result[index] != "" {
			if result[index] == "stable" {
				return version
			}

			version = version + "-" + expandStability(result[index])
			if result[index+1] != "" {
				version = version + result[index+1]
			}
		}

		if result[index+2] != "" {
			version = version + "-dev"
		}

		return version
	}

	result = RegFind(`(.*?)[.-]?dev$`, version)
	if result != nil {
		return normalizeBranch(result[1])
	}

	return version
}

func expandStability(stability string) string {
	stability = strings.ToLower(stability)

	switch stability {
	case "a":
		return "alpha"
	case "b":
		return "beta"
	case "p":
		return "patch"
	case "pl":
		return "patch"
	case "rc":
		return "RC"
	}

	return stability
}

func normalizeBranch(name string) string {
	name = strings.Trim(name, " ")

	if name == "master" || name == "trunk" || name == "default" {
		return Normalize(name)
	}

	r := regexp.MustCompile(`(?i)^v?(\d+)(\.(?:\d+|[x*]))?(\.(?:\d+|[x*]))?(\.(?:\d+|[x*]))?$`)
	replace := strings.NewReplacer("*", "9999999", "x", "9999999")

	matched := r.FindAllStringSubmatch(name, -1)
	if matched != nil {
		name = ""
		for _, val := range matched[0][1:5] {
			if val != "" {
				name = name + replace.Replace(val)
			} else {
				name = name + ".9999999"
			}
		}

		return name + "-dev"

	}

	return "dev-" + name
}

func parseConstraints(constraint string) []string {
	result := RegFind(`(?i)^([^,\s]*?)@(stable|RC|beta|alpha|dev)$`, constraint)
	if result != nil {
		constraint = result[1]
		if constraint == "" {
			constraint = "*"
		}
	}

	result = RegFind(`(?i)^(dev-[^,\s@]+?|[^,\s@]+?\.x-dev)#.+$`, constraint)
	if result != nil {
		if result[1] != "" {
			constraint = result[1]
		}
	}

	constraints := RegSplit(`\s*,\s*`, strings.Trim(constraint, " "))

	if len(constraints) > 1 {
		output := make([]string, 0)
		for _, part := range constraints {
			output = append(output, parseConstraint(part)...)
		}

		return output
	} else if len(constraints) == 0 {
		return []string{"000000", constraint}
	} else {
		return parseConstraint(constraints[0])
	}

	return []string{"222333", constraint}
}

func parseConstraint(constraint string) []string {

	stabilityModifier := ""

	result := RegFind(`(?i)^([^,\s]+?)@(stable|RC|beta|alpha|dev)$`, constraint)
	if result != nil {
		constraint = result[1]
		if result[2] != "stable" {
			stabilityModifier = result[2]
		}
	}

	result = RegFind(`^[x*](\.[x*])*$`, constraint)
	if result != nil {
		return []string{"", ""}
	}

	highVersion := ""
	lowVersion := ""

	result = RegFind(`(?i)^~(\d+)(?:\.(\d+))?(?:\.(\d+))?(?:\.(\d+))?`+modifierRegex+`?$`, constraint)
	if result != nil {
		if len(result) > 4 && result[4] != "" {
			last, _ := strconv.Atoi(result[3])
			highVersion = result[1] + "." + result[2] + "." + strconv.Itoa(last+1) + ".0-dev"
			lowVersion = result[1] + "." + result[2] + "." + result[3] + "." + result[4]
		} else if len(result) > 3 && result[3] != "" {
			last, _ := strconv.Atoi(result[2])
			highVersion = result[1] + "." + strconv.Itoa(last+1) + ".0.0-dev"
			lowVersion = result[1] + "." + result[2] + "." + result[3] + ".0"
		} else {
			last, _ := strconv.Atoi(result[1])
			highVersion = strconv.Itoa(last+1) + ".0.0.0-dev"
			if len(result) > 2 && result[2] != "" {
				lowVersion = result[1] + "." + result[2] + ".0.0"
			} else {
				lowVersion = result[1] + ".0.0.0"
			}
		}

		if len(result) > 5 && result[5] != "" {
			lowVersion = lowVersion + "-" + expandStability(result[5])

		}

		if len(result) > 6 && result[6] != "" {
			lowVersion = lowVersion + result[6]
		}

		if len(result) > 7 && result[7] != "" {
			lowVersion = lowVersion + "-dev"
		}

		//fmt.Printf("----------------------------->= %s\n", lowVersion)
		//fmt.Printf("-----------------------------< %s\n", highVersion)
		return []string{">=", lowVersion, "<", highVersion}
	}

	result = RegFind(`^(\d+)(?:\.(\d+))?(?:\.(\d+))?\.[x*]$`, constraint)
	if result != nil {
		if len(result) > 3 && result[3] != "" {
			highVersion = result[1] + "." + result[2] + "." + result[3] + ".9999999"
			if result[3] == "0" {
				last, _ := strconv.Atoi(result[2])
				lowVersion = result[1] + "." + strconv.Itoa(last-1) + ".9999999.9999999"
			} else {
				last, _ := strconv.Atoi(result[3])
				lowVersion = result[1] + "." + result[2] + "." + strconv.Itoa(last-1) + ".9999999"
			}

		} else if len(result) > 2 && result[2] != "" {
			highVersion = result[1] + "." + result[2] + ".9999999.9999999"
			if result[2] == "0" {
				last, _ := strconv.Atoi(result[1])
				lowVersion = strconv.Itoa(last-1) + ".9999999.9999999.9999999"
			} else {
				last, _ := strconv.Atoi(result[2])
				lowVersion = result[1] + "." + strconv.Itoa(last-1) + ".9999999.9999999"
			}

		} else {
			highVersion = result[1] + ".9999999.9999999.9999999"
			if result[1] == "0" {
				return []string{"", "", "<", highVersion}
			} else {
				last, _ := strconv.Atoi(result[1])
				lowVersion = strconv.Itoa(last-1) + ".9999999.9999999.9999999"
			}
		}

		//fmt.Printf("-----------------------------> %s\n", lowVersion)
		//fmt.Printf("-----------------------------< %s\n", highVersion)
		return []string{">", lowVersion, "<", highVersion}
	}

	// match operators constraints
	result = RegFind(`^(<>|!=|>=?|<=?|==?)?\s*(.*)`, constraint)
	if result != nil {
		version := Normalize(result[2])

		if stabilityModifier != "" && parseStability(version) == "stable" {
			version = version + "-" + stabilityModifier
		} else if result[1] == "<" {
			match := RegFind(`(?i)-stable$`, result[2])
			if match == nil {
				version = version + "-dev"
			}
		}

		if len(result) > 1 && result[1] != "" {
			//fmt.Printf("JODER:%s\n", strings.Join(result, "/"))
			return []string{result[1], version}
		} else {
			return []string{"=", version}
		}
	}

	return []string{constraint, stabilityModifier}
}

func parseStability(version string) string {
	version = regexp.MustCompile(`(?i)#.+$`).ReplaceAllString(version, " ")
	version = strings.ToLower(version)

	if strings.HasPrefix(version, "dev-") || strings.HasSuffix(version, "-dev") {
		return "dev"
	}

	result := RegFind(`(?i)^v?(\d{1,3})(\.\d+)?(\.\d+)?(\.\d+)?`+modifierRegex+`$`, version)
	if result != nil {
		if len(result) > 3 {
			return "dev"
		}
	}

	if result[1] != "" {
		if "beta" == result[1] || "b" == result[1] {
			return "beta"
		}
		if "alpha" == result[1] || "a" == result[1] {
			return "alpha"
		}
		if "rc" == result[1] {
			return "RC"
		}
	}

	return "stable"
}

func RegFind(pattern string, subject string) []string {
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(subject, -1)

	if matched != nil {
		return matched[0]
	}

	return nil
}

func RegSplit(pattern string, subject string) []string {
	reg := regexp.MustCompile(pattern)
	indexes := reg.FindAllStringIndex(subject, -1)

	laststart := 0
	result := make([]string, len(indexes)+1)

	for i, element := range indexes {
		result[i] = subject[laststart:element[0]]
		laststart = element[1]
	}

	result[len(indexes)] = subject[laststart:len(subject)]
	return result
}
