package misc

import (
	"regexp"
	"strings"
)

var modifierRegex = `[._-]?(?:(stable|beta|b|RC|alpha|a|patch|pl|p)(?:[.-]?(\d+))?)?([.-]?dev)?`

func Normalize(version string) string {
	r := regexp.MustCompile(`^([^,\s]+) +as +([^,\s]+)$`)

	// ignore aliases and just assume the alias is required instead of the source
	matched := r.FindAllStringSubmatch(version, -1)
	if matched != nil {
		version = matched[0][1]
	}

	// match master-like branches
	r = regexp.MustCompile(`^(?:dev-)?(?:master|trunk|default)$`)
	if r.MatchString(strings.ToLower(version)) {
		return "9999999-dev"
	}

	if strings.HasPrefix(strings.ToLower(version), "dev-") {
		return "dev-" + version[4:len(version)]
	}

	// match classical versioning
	r = regexp.MustCompile(`(?i)^v?(\d{1,3})(\.\d+)?(\.\d+)?(\.\d+)?` + modifierRegex + `$`)

	index := 0
	matched = r.FindAllStringSubmatch(version, -1)
	if matched != nil {
		version = ""
		for _, val := range matched[0][1:5] {
			if val != "" {
				version = version + val
			} else {
				version = version + ".0"
			}
		}

		index = 5
	} else {
		// match date-based versioning
		r = regexp.MustCompile(`(?i)^v?(\d{4}(?:[.:-]?\d{2}){1,6}(?:[.:-]?\d{1,3})?)` + modifierRegex + `$`)
		matched = r.FindAllStringSubmatch(version, -1)

		if matched != nil {
			version = regexp.MustCompile(`\D`).ReplaceAllString(matched[0][1], "-")
			index = 2
		}
	}

	if index != 0 {
		if matched[0][index] != "" {
			if matched[0][index] == "stable" {
				return version
			}

			version = version + "-" + expandStability(matched[0][index])
			if matched[0][index+1] != "" {
				version = version + matched[0][index+1]
			}
		}

		if matched[0][index+2] != "" {
			version = version + "-dev"
		}

		return version
	}

	r = regexp.MustCompile(`(.*?)[.-]?dev$`)
	matched = r.FindAllStringSubmatch(version, -1)

	if matched != nil {
		return normalizeBranch(matched[0][1])
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

	if name == "master" {
		return Normalize(name)
	}

	if name == "trunk" {
		return Normalize(name)
	}

	if name == "default" {
		return Normalize(name)
	}

	//name = strings.ToLower(name)

	r := regexp.MustCompile(`^v?(\d+)(\.(?:\d+|[x*]))?(\.(?:\d+|[x*]))?(\.(?:\d+|[x*]))?$`)
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
