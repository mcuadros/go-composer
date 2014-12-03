package packagist

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mcuadros/go-version"
)

type Resolver struct {
	packagist *Packagist
	Packages  map[string]*Package
	sync.WaitGroup
}

func NewResolver() *Resolver {
	return &Resolver{
		packagist: NewPackagist(),
		Packages:  make(map[string]*Package, 0),
	}
}

func (r *Resolver) Resolve(name, constraint string) error {
	pckg := r.GetPackage(name, constraint)
	if pckg == nil {
		panic(fmt.Errorf("unable to find %s -> %s", name, constraint))
	}

	r.addPackage(pckg)
	return nil
}

func (r *Resolver) addPackage(pckg *Package) error {
	r.Packages[pckg.Name] = pckg

	deps, err := r.resolvePackageDependencies(pckg)
	if err != nil {
		return err
	}

	for _, pckg := range deps {
		r.addPackage(pckg)
	}

	return nil
}

func (r *Resolver) resolvePackageDependencies(pckg *Package) ([]*Package, error) {
	pckgs := make([]*Package, 0)

	for require, constraint := range pckg.Require {
		if p, ok := r.Packages[require]; ok {
			if r.isConflict(p, constraint) {
				//panic(fmt.Errorf("conflict %s -> %s <%s>", p, require, constraint))
			}

			continue
		}

		if require == "php" {
			continue
		}

		if strings.HasPrefix(require, "ext-") {
			continue
		}

		if constraint == "self.version" {
			constraint = pckg.Version
		}

		if p := r.GetPackage(require, constraint); p != nil {
			pckgs = append(pckgs, p)
		}
	}

	return pckgs, nil
}

func (r *Resolver) isConflict(pckg *Package, constraint string) bool {
	c := version.NewConstrainGroupFromString(constraint)

	return !c.Match(pckg.Version)
}

func (r *Resolver) GetPackage(name, constraint string) *Package {
	pckgs, err := r.packagist.GetPackages(name)
	if err != nil {
		return nil
	}

	return pckgs.Get(constraint)
}
