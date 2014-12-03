package packagist

import (
	. "gopkg.in/check.v1"
)

func (s *ComposerSuite) TestPackages_Get(c *C) {
	expected := &Package{Version: "1.2.3"}
	pckgs := Packages{
		"1.0.0": &Package{Version: "1.0.0"},
		"1.2.1": &Package{Version: "1.2.1"},
		"1.2.3": expected,
		"1.2.2": &Package{Version: "1.2.2"},
		"2.0.0": &Package{Version: "2.0.0"},
	}

	r := pckgs.Get("1.2.*")
	c.Assert(r, Equals, expected)
}
