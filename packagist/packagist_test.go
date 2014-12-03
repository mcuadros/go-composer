package packagist

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ComposerSuite struct{}

var _ = Suite(&ComposerSuite{})

func (s *ComposerSuite) TestPackagist_NewPackagist(c *C) {
	p := NewPackagist()
	r, err := p.GetPackages("composer/composer")
	c.Assert(err, IsNil)
	c.Assert(r, HasLen, 9)
}
