package misc

import "testing"
import (
	"strings"
)

var versions = map[string]string{
	"1-stable":            "1.0.0.0",
	"1.0.0":               "1.0.0.0",
	"1.2.3.4":             "1.2.3.4",
	"1.0.0RC1dev":         "1.0.0.0-RC1-dev",
	"1.0.0-rC15-dev":      "1.0.0.0-RC15-dev",
	"1.0.0.RC.15-dev":     "1.0.0.0-RC15-dev",
	"1.0.0-rc1":           "1.0.0.0-RC1",
	"1.0.0.pl3-dev":       "1.0.0.0-patch3-dev",
	"1.0-dev":             "1.0.0.0-dev",
	"0":                   "0.0.0.0",
	"10.4.13-beta":        "10.4.13.0-beta",
	"10.4.13-b":           "10.4.13.0-beta",
	"10.4.13-b5":          "10.4.13.0-beta5",
	"v1.0.0":              "1.0.0.0",
	"v20100102":           "20100102",
	"2010.01":             "2010-01",
	"2010.01.02":          "2010-01-02",
	"2010-01-02":          "2010-01-02",
	"2010-01-02.5":        "2010-01-02-5",
	"20100102-203040":     "20100102-203040",
	"20100102203040-10":   "20100102203040-10",
	"20100102-203040-p1":  "20100102-203040-patch1",
	"dev-master":          "9999999-dev",
	"dev-trunk":           "9999999-dev",
	"1.x-dev":             "1.9999999.9999999.9999999-dev",
	"dev-feature-foo":     "dev-feature-foo",
	"DEV-FOOBAR":          "dev-FOOBAR",
	"dev-feature/foo":     "dev-feature/foo",
	"dev-master as 1.0.0": "9999999-dev",
}

func TestNormalize(t *testing.T) {
	for in, out := range versions {
		if x := Normalize(in); x != out {
			t.Errorf("FAIL: Normalize(%v) = %v: want %v", in, x, out)
		}
	}
}

var branches = map[string]string{
	"v1.x":      "1.9999999.9999999.9999999-dev",
	"v1.*":      "1.9999999.9999999.9999999-dev",
	"v1.0":      "1.0.9999999.9999999-dev",
	"2.0":       "2.0.9999999.9999999-dev",
	"v1.0.x":    "1.0.9999999.9999999-dev",
	"v1.0.3.*":  "1.0.3.9999999-dev",
	"v2.4.0":    "2.4.0.9999999-dev",
	"2.4.4":     "2.4.4.9999999-dev",
	"master":    "9999999-dev",
	"trunk":     "9999999-dev",
	"feature-a": "dev-feature-a",
	"FOOBAR":    "dev-FOOBAR",
}

func TestNormalizeBranch(t *testing.T) {
	for in, out := range branches {
		if x := normalizeBranch(in); x != out {
			t.Errorf("FAIL: normalizeBranch(%v) = %v: want %v", in, x, out)
		}
	}
}

var simpleConstraints = map[string][]string{
	"*":                   []string{"", ""},
	"*.*":                 []string{"", ""},
	"*.x.*":               []string{"", ""},
	"x.x.x.*":             []string{"", ""},
	"<>1.0.0":             []string{"<>", "1.0.0.0"},
	"!=1.0.0":             []string{"!=", "1.0.0.0"},
	">1.0.0":              []string{">", "1.0.0.0"},
	"<1.2.3.4":            []string{"<", "1.2.3.4-dev"},
	"<=1.2.3":             []string{"<=", "1.2.3.0"},
	">=1.2.3":             []string{">=", "1.2.3.0"},
	"=1.2.3":              []string{"=", "1.2.3.0"},
	"==1.2.3":             []string{"=", "1.2.3.0"},
	"1.2.3":               []string{"=", "1.2.3.0"},
	"=1.0":                []string{"=", "1.0.0.0"},
	"1.2.3b5":             []string{"=", "1.2.3.0-beta5"},
	">= 1.2.3":            []string{">=", "1.2.3.0"},
	">=dev-master":        []string{">=", "9999999-dev"},
	"dev-master":          []string{"=", "9999999-dev"},
	"dev-feature-a":       []string{"=", "dev-feature-a"},
	"dev-some-fix":        []string{"=", "dev-some-fix"},
	"dev-CAPS":            []string{"=", "dev-CAPS"},
	"dev-master as 1.0.0": []string{"=", "9999999-dev"},
	"<1.2.3.4-stable":     []string{"<", "1.2.3.4"},

	"<=3.0@dev":            []string{"<=", "3.0.0.0"},
	"1.0@dev":              []string{"=", "1.0.0.0"},                 //IgnoresStabilityFlag
	"1.0.x-dev#abcd123":    []string{"=", "1.0.9999999.9999999-dev"}, //IgnoresReferenceOnDevVersion
	"1.0.x-dev#trunk/@123": []string{"=", "1.0.9999999.9999999-dev"}, //IgnoresReferenceOnDevVersion
	//"1.0#abcd123":          []string{"=", "1.0.0.0"},                 //FailsOnBadReference
	//"1.0#trunk/@123":       []string{"=", "1.0.0.0"},                 //FailsOnBadReference
}

func TestParseConstraints(t *testing.T) {
	for in, out := range simpleConstraints {
		if x := parseConstraints(in); strings.Join(x, "-") != strings.Join(out, "-") {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var wildcardConstraints = map[string][]string{
	"2.*":     []string{">", "1.9999999.9999999.9999999", "<", "2.9999999.9999999.9999999"},
	"20.*":    []string{">", "19.9999999.9999999.9999999", "<", "20.9999999.9999999.9999999"},
	"2.0.*":   []string{">", "1.9999999.9999999.9999999", "<", "2.0.9999999.9999999"},
	"2.2.x":   []string{">", "2.1.9999999.9999999", "<", "2.2.9999999.9999999"},
	"2.10.x":  []string{">", "2.9.9999999.9999999", "<", "2.10.9999999.9999999"},
	"2.1.3.*": []string{">", "2.1.2.9999999", "<", "2.1.3.9999999"},
	"0.*":     []string{"", "", "<", "0.9999999.9999999.9999999"},
}

func TestParseConstraintsWildcardConstraints(t *testing.T) {
	for in, out := range wildcardConstraints {
		if x := parseConstraints(in); strings.Join(x, "-") != strings.Join(out, "-") {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var tildeConstraints = map[string][]string{
	"~1":         []string{">=", "1.0.0.0", "<", "2.0.0.0-dev"},
	"~1.2":       []string{">=", "1.2.0.0", "<", "2.0.0.0-dev"},
	"~1.2.3":     []string{">=", "1.2.3.0", "<", "1.3.0.0-dev"},
	"~1.2.3.4":   []string{">=", "1.2.3.4", "<", "1.2.4.0-dev"},
	"~1.2-beta":  []string{">=", "1.2.0.0-beta", "<", "2.0.0.0-dev"},
	"~1.2-b2":    []string{">=", "1.2.0.0-beta2", "<", "2.0.0.0-dev"},
	"~1.2-BETA2": []string{">=", "1.2.0.0-beta2", "<", "2.0.0.0-dev"},
	"~1.2.2-dev": []string{">=", "1.2.2.0-dev", "<", "1.3.0.0-dev"},
}

func TestParseConstraintsTildeConstraints(t *testing.T) {
	for in, out := range tildeConstraints {
		if x := parseConstraints(in); strings.Join(x, "-") != strings.Join(out, "-") {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var multiConstraints = map[string][]string{
	">2.0,<=3.0":            []string{">", "2.0.0.0", "<=", "3.0.0.0"},
	">2.0@stable,<=3.0@dev": []string{">", "2.0.0.0", "<=", "3.0.0.0-dev"},
}

func TestParseConstraintsMultiConstraints(t *testing.T) {
	for in, out := range multiConstraints {
		if x := parseConstraints(in); strings.Join(x, "-") != strings.Join(out, "-") {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}
