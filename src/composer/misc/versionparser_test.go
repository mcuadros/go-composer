package misc

import "testing"

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
}

func TestParseConstraints(t *testing.T) {
	for in, out := range simpleConstraints {
		if x := parseConstraints(in); x[0] != out[0] {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

/*
   "*":                   [2]string{nil, nil},
   "*.*":                 [2]string{nil, nil},
   "*.x.*":               [2]string{nil, nil},
   "x.x.x.*":             [2]string{nil, nil},
*/

/*
var wildcardConstraints = map[string][]string{


}

    public function wildcardConstraints()
    {
        return array(
            "2.*":     new VersionConstraint(">": "1.9999999.9999999.9999999"): new VersionConstraint("<": "2.9999999.9999999.9999999")):
            "20.*":    new VersionConstraint(">": "19.9999999.9999999.9999999"):new VersionConstraint("<": "20.9999999.9999999.9999999")):
            "2.0.*":   new VersionConstraint(">": "1.9999999.9999999.9999999"): new VersionConstraint("<": "2.0.9999999.9999999")):
            "2.2.x":   new VersionConstraint(">": "2.1.9999999.9999999"):       new VersionConstraint("<": "2.2.9999999.9999999")):
            "2.10.x":  new VersionConstraint(">": "2.9.9999999.9999999"):       new VersionConstraint("<": "2.10.9999999.9999999")):
            "2.1.3.*": new VersionConstraint(">": "2.1.2.9999999"):             new VersionConstraint("<": "2.1.3.9999999")):
            "0.*":     null:                                                    new VersionConstraint("<": "0.9999999.9999999.9999999")):
        );
    }
    public function tildeConstraints()
    {
        return "            "~1":         new VersionConstraint(">=": "1.0.0.0"):         new VersionConstraint("<": "2.0.0.0-dev")):
            "~1.2":       new VersionConstraint(">=": "1.2.0.0"):         new VersionConstraint("<": "2.0.0.0-dev")):
            "~1.2.3":     new VersionConstraint(">=": "1.2.3.0"):         new VersionConstraint("<": "1.3.0.0-dev")):
            "~1.2.3.4":   new VersionConstraint(">=": "1.2.3.4"):         new VersionConstraint("<": "1.2.4.0-dev")):
            "~1.2-beta":  new VersionConstraint(">=": "1.2.0.0-beta"):    new VersionConstraint("<": "2.0.0.0-dev")):
            "~1.2-b2":    new VersionConstraint(">=": "1.2.0.0-beta2"):   new VersionConstraint("<": "2.0.0.0-dev")):
            "~1.2-BETA2": new VersionConstraint(">=": "1.2.0.0-beta2"):   new VersionConstraint("<": "2.0.0.0-dev")):
            "~1.2.2-dev": new VersionConstraint(">=": "1.2.2.0-dev"):     new VersionConstraint("<": "1.3.0.0-dev")):
        );
    }
*/
