package packagist

import (
	. "gopkg.in/check.v1"
)

func (s *ComposerSuite) TestNewResolver_Resolve(c *C) {
	r := NewResolver()
	err := r.Resolve("mongator/mongator", "dev-master")
	c.Assert(err, IsNil)
	c.Assert(r.Packages, HasLen, 4)
}
