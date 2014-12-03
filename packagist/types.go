package packagist

import (
	"fmt"

	"github.com/mcuadros/go-version"
)

type Package struct {
	Name              string
	Description       string
	Keywords          []string
	Homepage          string
	Version           string
	VersionNormalized string `json:"version_normalized"`
	Source            *Source
	Require           map[string]string
	RequireDev        map[string]string `json:""require-dev"`
}

func (p *Package) String() string {
	return fmt.Sprintf("%s <%s>", p.Name, p.Version)
}

type Packages map[string]*Package

func (p Packages) Get(constraint string) *Package {
	c := version.NewConstrainGroupFromString(constraint)

	versions := p.getSortedVersions()
	for i := len(versions) - 1; i >= 0; i-- {
		if c.Match(versions[i]) {
			return p[versions[i]]
		}
	}

	return nil
}

func (p Packages) getSortedVersions() []string {
	versions := p.getVersions()
	version.Sort(versions)
	return versions
}

func (p Packages) getVersions() []string {
	versions := make([]string, len(p))
	var i int
	for v, _ := range p {
		versions[i] = v
		i++
	}

	return versions
}

type Source struct {
	Type      string
	Url       string
	Reference string
}
